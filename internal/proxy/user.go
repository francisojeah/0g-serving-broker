package proxy

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/0glabs/0g-data-retrieve-agent/internal/errors"
	"github.com/gin-gonic/gin"
)

func (p *Proxy) GetData(c *gin.Context, url, name, provider, key string) {
	client := &http.Client{}

	var reqBody map[string]interface{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		errors.Response(c, err)
		return
	}

	cbReq := chatBotRequest{
		CreatedAt:           time.Now().Format(time.RFC3339),
		UserAddress:         p.address,
		Nonce:               "12345678910111213",
		ServiceName:         name,
		PreviousOutputCount: "0",
	}
	if err := cbReq.generate(reqBody, key, provider); err != nil {
		errors.Response(c, err)
		return
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		errors.Response(c, err)
		return
	}

	route := url + servicePrefix + "/" + name
	req, err := http.NewRequest(c.Request.Method, route, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		errors.Response(c, err)
		return
	}

	req.Header.Set("Token-Count", cbReq.InputCount)
	req.Header.Set("Address", cbReq.UserAddress)
	req.Header.Set("Service-Name", cbReq.ServiceName)
	req.Header.Set("Previous-Output-Token-Count", cbReq.PreviousOutputCount)
	req.Header.Set("Created-At", cbReq.CreatedAt)
	req.Header.Set("Nonce", cbReq.Nonce)
	req.Header.Set("Previous-Signature", cbReq.PreviousSignature)
	req.Header.Set("Signature", cbReq.Signature)

	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
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

	// TODO: Get output token count from resp.Body

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errors.Response(c, err)
		return
	}
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
