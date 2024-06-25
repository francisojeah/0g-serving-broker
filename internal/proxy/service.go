package proxy

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/0glabs/0g-data-retrieve-agent/internal/contract"
	"github.com/0glabs/0g-data-retrieve-agent/internal/errors"
	"github.com/0glabs/0g-data-retrieve-agent/internal/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
)

func (p *Proxy) AddHTTPRoute(route string, targetURL string) {
	//TODO: Add a URL validation
	_, exists := p.serviceRoutes[route]

	p.serviceRoutesLock.Lock()
	p.serviceRoutes[route] = true
	p.serviceRoutesLock.Unlock()

	if exists {
		return
	}

	h := func(c *gin.Context) {
		p.proxyHTTPRequest(c, route, targetURL)
	}
	p.serviceGroup.Any(route+"/*any", h)
}

func (p *Proxy) DeleteRoute(route string) {
	p.serviceRoutesLock.Lock()
	p.serviceRoutes[route] = false
	p.serviceRoutesLock.Unlock()
}

func (p *Proxy) validateRequest(dbReq model.Request) error {
	cReq := contract.Request{}
	if err := cReq.ConvertFromDB(dbReq); err != nil {
		return errors.Wrap(err, "convert request from db schema to contract schema")
	}
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	// TODO: Verify the following fields in the request header:
	//  - inputToken matches the number of input tokens in the request body.
	//  - previousOutputCount matches the number of tokens returned in the previous response.
	//  - previousSignature matches the signature of the previous request.
	//  - nonce is greater than the nonce of the previous request.

	pass, err := p.contract.Verify(callOpts, cReq)
	if err != nil {
		return errors.Wrap(err, "verify request")
	}
	if !pass {
		return errors.New("invalid request")
	}

	return nil
}

func (p *Proxy) proxyHTTPRequest(c *gin.Context, route, targetURL string) {
	dbReq, err := getRequest(c.Request)
	if err != nil {
		errors.Response(c, err)
		return
	}
	if err := p.validateRequest(*dbReq); err != nil {
		errors.Response(c, err)
		return
	}
	if ret := p.db.Create(&dbReq); ret.Error != nil {
		errors.Response(c, err)
		return
	}

	client := &http.Client{}
	targetRoute := strings.TrimPrefix(c.Request.RequestURI, servicePrefix+"/"+route)
	req, err := http.NewRequest(c.Request.Method, targetURL+targetRoute, c.Request.Body)
	if err != nil {
		errors.Response(c, err)
		return
	}

	for k, v := range c.Request.Header {
		if _, ok := requestMetaData[k]; !ok {
			req.Header.Set(k, v[0])
			continue
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		errors.Response(c, err)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}
	c.Writer.WriteHeader(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errors.Response(c, err)
		return
	}
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func getRequest(req *http.Request) (*model.Request, error) {
	dbReq := &model.Request{}
	for k := range requestMetaData {
		values := req.Header.Values(k)
		if len(values) == 0 {
			return nil, errors.Wrapf(errors.New("missing Header"), "%s", k)
		}
		switch k {
		case "Address":
			dbReq.UserAddress = values[0]
		case "Nonce":
			dbReq.Nonce = values[0]
		case "Name":
			dbReq.Name = values[0]
		case "Token-Count":
			dbReq.InputCount = values[0]
		case "Previous-Output-Token-Count":
			dbReq.PreviousOutputCount = values[0]
		case "Previous-Signature":
			dbReq.PreviousSignature = values[0]
		case "Signature":
			dbReq.Signature = values[0]
		case "Created-At":
			dbReq.CreatedAt = values[0]
		}
	}
	return dbReq, nil
}
