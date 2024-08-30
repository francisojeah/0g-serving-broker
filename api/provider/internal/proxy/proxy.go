package proxy

import (
	"io"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"

	constant "github.com/0glabs/0g-serving-agent/common/const"
	"github.com/0glabs/0g-serving-agent/common/errors"
	"github.com/0glabs/0g-serving-agent/extractor"
	"github.com/0glabs/0g-serving-agent/extractor/chatbot"
	"github.com/0glabs/0g-serving-agent/extractor/zgstorage"
	"github.com/0glabs/0g-serving-agent/provider/internal/ctrl"
)

type Proxy struct {
	ctrl *ctrl.Ctrl

	serviceRoutes     map[string]bool
	serviceRoutesLock sync.RWMutex
	serviceTargets    map[string]string
	serviceTypes      map[string]string
	serviceGroup      *gin.RouterGroup
}

func New(ctrl *ctrl.Ctrl, engine *gin.Engine) *Proxy {
	p := &Proxy{
		ctrl:           ctrl,
		serviceRoutes:  make(map[string]bool),
		serviceTargets: make(map[string]string),
		serviceTypes:   make(map[string]string),
		serviceGroup:   engine.Group(constant.ServicePrefix),
	}

	p.serviceGroup.Use(p.routeFilterMiddleware)
	return p
}

func (p *Proxy) Start() error {
	svcs, err := p.ctrl.ListService()
	if err != nil {
		return errors.Wrap(err, "Provider: start proxy service, list service in db")
	}
	for _, svc := range svcs {
		switch svc.Type {
		case "zgStorage", "chatbot":
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
		handleAgentError(ctx, errors.New("route is invalid"), "route filter middleware")
		return
	}

	p.serviceRoutesLock.RLock()
	valid, exists := p.serviceRoutes[segments[0]]
	p.serviceRoutesLock.RUnlock()
	if !exists {
		handleAgentError(ctx, errors.New("route is not exist"), "route filter middleware")
		return
	}
	if !valid {
		handleAgentError(ctx, errors.New("route is deleted"), "route filter middleware")
		return
	}
	ctx.Next()
}

func (p *Proxy) AddHTTPRoute(route, targetURL, svcType string) {
	//TODO: Add a URL validation
	_, exists := p.serviceRoutes[route]

	p.serviceRoutesLock.Lock()
	p.serviceRoutes[route] = true
	p.serviceTargets[route] = targetURL
	p.serviceTypes[route] = svcType
	p.serviceRoutesLock.Unlock()

	if exists {
		return
	}

	h := func(ctx *gin.Context) {
		p.proxyHTTPRequest(ctx, route)
	}
	p.serviceGroup.Any(route+"/*any", h)
}

func (p *Proxy) DeleteRoute(route string) {
	p.serviceRoutesLock.Lock()
	p.serviceRoutes[route] = false
	delete(p.serviceTargets, route)
	delete(p.serviceTypes, route)
	p.serviceRoutesLock.Unlock()
}

func (p *Proxy) UpdateRoute(route string, newTargetURL, newSvcType string) error {
	//TODO: Add a URL validation
	valid, exists := p.serviceRoutes[route]
	if !exists {
		return errors.New("route is not exist")
	}
	if !valid {
		return errors.New("route is deleted")
	}

	p.serviceRoutesLock.Lock()
	p.serviceRoutes[route] = true
	p.serviceTargets[route] = newTargetURL
	p.serviceTypes[route] = newSvcType
	p.serviceRoutesLock.Unlock()

	return nil
}

func (p *Proxy) proxyHTTPRequest(ctx *gin.Context, route string) {
	p.serviceRoutesLock.RLock()
	targetURL := p.serviceTargets[route]
	svcType := p.serviceTypes[route]
	p.serviceRoutesLock.RUnlock()

	var extractor extractor.ProviderReqRespExtractor
	switch svcType {
	case "zgStorage":
		extractor = &zgstorage.ProviderZgStorage{}
	case "chatbot":
		extractor = &chatbot.ProviderChatBot{}
	default:
		handleAgentError(ctx, errors.New("unknown service type"), "prepare request extractor")
		return
	}
	svc, err := p.ctrl.GetService(route)
	if err != nil {
		handleAgentError(ctx, err, "get service")
		return
	}
	req, err := p.ctrl.GetFromHTTPRequest(ctx)
	if err != nil {
		handleAgentError(ctx, err, "get model.request from HTTP request")
		return
	}
	fee := svc.InputPrice*req.InputCount + svc.OutputPrice*req.PreviousOutputCount
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		handleAgentError(ctx, err, "read request body")
		return
	}
	inputCount, err := extractor.GetInputCount(reqBody)
	if err != nil {
		handleAgentError(ctx, err, "get input count")
		return
	}
	if err := p.ctrl.ValidateRequest(ctx, req, fee, inputCount); err != nil {
		handleAgentError(ctx, err, "validate request")
		return
	}
	if err := p.ctrl.CreateRequest(req); err != nil {
		handleAgentError(ctx, err, "create request")
		return
	}

	httpReq, err := p.ctrl.PrepareHTTPRequest(ctx, targetURL, route, reqBody)
	if err != nil {
		handleAgentError(ctx, err, "prepare HTTP request")
		return
	}
	p.ctrl.ProcessHTTPRequest(ctx, httpReq, req, extractor, fee)
}

func handleAgentError(ctx *gin.Context, err error, context string) {
	info := "Provider proxy: handle proxied service"
	if context != "" {
		info += (", " + context)
	}
	errors.Response(ctx, errors.Wrap(err, info))
}
