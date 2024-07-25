package proxy

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	constant "github.com/0glabs/0g-serving-agent/common/const"
	"github.com/0glabs/0g-serving-agent/common/contract"
	"github.com/0glabs/0g-serving-agent/common/errors"
	commonModel "github.com/0glabs/0g-serving-agent/common/model"
	"github.com/0glabs/0g-serving-agent/extractor"
	"github.com/0glabs/0g-serving-agent/extractor/chatbot"
	"github.com/0glabs/0g-serving-agent/provider/internal/ctrl"
	"github.com/0glabs/0g-serving-agent/provider/model"
)

type Proxy struct {
	db       *gorm.DB
	ctrl     *ctrl.Ctrl
	contract *contract.ServingContract
	address  string

	serviceRoutes     map[string]bool
	serviceRoutesLock sync.RWMutex
	serviceGroup      *gin.RouterGroup
}

func New(db *gorm.DB, ctrl *ctrl.Ctrl, router *gin.Engine, c *contract.ServingContract, address string) *Proxy {
	p := &Proxy{
		db:            db,
		ctrl:          ctrl,
		contract:      c,
		address:       address,
		serviceRoutes: make(map[string]bool),
		serviceGroup:  router.Group(constant.ServicePrefix),
	}

	p.serviceGroup.Use(p.routeFilterMiddleware)
	return p
}

func (p *Proxy) Start() error {
	tx := p.db.Model(model.Service{})
	services := []model.Service{}
	if ret := tx.Find(&services); ret.Error != nil {
		return ret.Error
	}
	for _, svc := range services {
		switch svc.Type {
		case "RPC":
			// TODO: Add p.AddRPCRoute
		case "chatbot":
			p.AddHTTPRoute(svc.Name, svc.URL, svc.Type)
		default:
			return errors.New("invalid service type")
		}
	}
	return nil
}

func (p *Proxy) routeFilterMiddleware(ctx *gin.Context) {
	route := strings.TrimPrefix(ctx.Request.URL.Path, constant.ServicePrefix+"/")
	segments := strings.Split(route, "/")
	if len(segments) == 0 || segments[0] == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	p.serviceRoutesLock.RLock()
	valid, exists := p.serviceRoutes[segments[0]]
	p.serviceRoutesLock.RUnlock()
	if !exists || !valid {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.Next()
}

func (p *Proxy) AddHTTPRoute(route, targetURL, svcType string) {
	//TODO: Add a URL validation
	_, exists := p.serviceRoutes[route]

	p.serviceRoutesLock.Lock()
	p.serviceRoutes[route] = true
	p.serviceRoutesLock.Unlock()

	if exists {
		return
	}

	h := func(c *gin.Context) {
		p.proxyHTTPRequest(c, route, targetURL, svcType)
	}
	p.serviceGroup.Any(route+"/*any", h)
}

func (p *Proxy) DeleteRoute(route string) {
	p.serviceRoutesLock.Lock()
	p.serviceRoutes[route] = false
	p.serviceRoutesLock.Unlock()
}

func (p *Proxy) proxyHTTPRequest(ctx *gin.Context, route, targetURL, svcType string) {
	reqV := requestValidator{
		db:       p.db,
		contract: p.contract,
		request:  &commonModel.Request{},
	}

	var extractor extractor.ProviderReqRespExtractor
	switch svcType {
	case "chatbot":
		extractor = &chatbot.ProviderChatBot{}
	default:
		requestError(ctx, errors.New("unknown service type"), "prepare request extractor")
		return
	}
	reqV.extractor = extractor

	err := reqV.backFillMetadata(ctx, p.address)
	if err != nil {
		requestError(ctx, err, "get request metadata")
		return
	}

	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		errors.Response(ctx, err)
		requestError(ctx, err, "read request body")
		return
	}

	err = reqV.validate(reqBody)
	if err != nil {
		requestError(ctx, err, "verify request")
		return
	}

	// request should be stored even if an error occurs in the proxied service.
	if ret := p.db.Create(reqV.request); ret.Error != nil {
		requestError(ctx, err, "store request")
		return
	}

	req, err := p.prepareRequest(ctx, targetURL, route, reqBody)
	if err != nil {
		requestError(ctx, err, "prepare request")
		return
	}

	p.processRequest(ctx, req, &reqV)
}

func (p *Proxy) prepareRequest(ctx *gin.Context, targetURL, route string, reqBody []byte) (*http.Request, error) {
	targetRoute := strings.TrimPrefix(ctx.Request.RequestURI, constant.ServicePrefix+"/"+route)
	if targetRoute != "/" {
		targetURL += targetRoute
	}
	req, err := http.NewRequest(ctx.Request.Method, targetURL, io.NopCloser(bytes.NewBuffer(reqBody)))
	if err != nil {
		errors.Response(ctx, errors.Wrap(err, "provider proxy: prepare request for the proxied service"))
		return nil, err
	}

	for k, v := range ctx.Request.Header {
		if _, ok := constant.RequestMetaData[k]; !ok {
			req.Header.Set(k, v[0])
			continue
		}
	}
	return req, nil
}

func (p *Proxy) processRequest(ctx *gin.Context, req *http.Request, reqV *requestValidator) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errors.Response(ctx, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseError(ctx, reqV.extractor.ErrMsg(resp.Body), "read error from proxied service")
		return
	}

	for k, v := range resp.Header {
		ctx.Writer.Header()[k] = v
	}
	ctx.Writer.WriteHeader(resp.StatusCode)

	fee := reqV.getUnsettleFee()
	account := model.User{
		User:             reqV.request.UserAddress,
		LastRequestNonce: &reqV.request.Nonce,
		UnsettledFee:     &fee,
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/event-stream") {
		p.handleResponse(ctx, resp, reqV.extractor, account)
	} else {
		p.handleStreamResponse(ctx, resp, reqV.extractor, account)
	}
}

func (p *Proxy) handleResponse(ctx *gin.Context, resp *http.Response, extractor extractor.ProviderReqRespExtractor, account model.User) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		responseError(ctx, err, "read from body")
		return
	}

	contentEncoding := resp.Header.Get("Content-Encoding")
	outputContent, err := extractor.GetRespContent(body, contentEncoding)
	if err != nil {
		responseError(ctx, err, "extract content")
		return
	}

	outputCount, err := extractor.GetOutputCount([][]byte{outputContent})
	if err != nil {
		responseError(ctx, err, "extract count")
		return
	}

	account.LastResponseTokenCount = &outputCount
	if err = p.ctrl.UpdateUserAccount(account.User, account); err != nil {
		responseError(ctx, err, "update user account in db")
		return
	}

	ctx.Data(http.StatusOK, resp.Header.Get("Content-Type"), body)
}

func (p *Proxy) handleStreamResponse(ctx *gin.Context, resp *http.Response, extractor extractor.ProviderReqRespExtractor, account model.User) {
	ctx.Stream(func(w io.Writer) bool {
		var chunkBuf bytes.Buffer
		var output [][]byte
		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					return false
				}
				responseError(ctx, err, "read from body")
				return false
			}

			chunkBuf.WriteString(line)
			if line == "\n" || line == "\r\n" {
				_, err := w.Write(chunkBuf.Bytes())
				if err != nil {
					responseError(ctx, err, "write to stream")
					return false
				}

				encoding := resp.Header.Get("Content-Encoding")
				content, err := extractor.GetRespContent(chunkBuf.Bytes(), encoding)
				if err != nil {
					responseError(ctx, err, "extract content")
					return false
				}

				completed, err := extractor.StreamCompleted(content)
				if err != nil {
					responseError(ctx, err, "check stream completed")
					return false
				}
				if completed {
					outputCount, err := extractor.GetOutputCount(output)
					if err != nil {
						responseError(ctx, err, "extract output count")
						return false
					}

					account.LastResponseTokenCount = &outputCount
					err = p.ctrl.UpdateUserAccount(account.User, account)
					if err != nil {
						responseError(ctx, err, "update user account in db")
						return false
					}
				}
				output = append(output, content)
				ctx.Writer.Flush()
				chunkBuf.Reset()
			}
		}
	})
}

func responseError(ctx *gin.Context, err error, context string) {
	errors.Response(ctx, errors.Wrap(err, "provider proxy: handle proxied service response, "+context))
}

func requestError(ctx *gin.Context, err error, context string) {
	errors.Response(ctx, errors.Wrap(err, "provider proxy: handle proxied service request, "+context))
}
