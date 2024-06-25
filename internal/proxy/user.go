package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/0glabs/0g-data-retrieve-agent/internal/contract"
	"github.com/0glabs/0g-data-retrieve-agent/internal/errors"
	"github.com/0glabs/0g-data-retrieve-agent/internal/model"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
)

func (p *Proxy) GetData(c *gin.Context, url, route, provider, key string) {
	client := &http.Client{}

	var reqBody map[string]interface{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// For Chatbot: Read the request body to extract the `message` field
	message, ok := reqBody["message"].(string)
	if !ok || message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid message field"})
		return
	}
	tokenCount := len(strings.Fields(message))

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request body"})
		return
	}
	req, err := http.NewRequest(c.Request.Method, url+servicePrefix+"/"+route, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// TODO: Get metadata from DB instead of mock
	dbReq := model.Request{
		CreatedAt:           time.Now().Format(time.RFC3339),
		UserAddress:         p.address,
		Nonce:               "12345",
		Name:                route,
		InputCount:          fmt.Sprintf("%d", tokenCount),
		PreviousOutputCount: "0",
		PreviousSignature:   "0x0000000000000000000000000000000000000000",
	}

	cReq := contract.Request{}
	if err := cReq.ConvertFromDB(dbReq); err != nil {
		errors.Response(c, err)
		return
	}

	sig, err := cReq.GetSignature(key, provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to form contract request signature"})
		return
	}
	dbReq.Signature = hexutil.Encode(sig)

	req.Header.Set("Token-Count", dbReq.InputCount)
	req.Header.Set("Address", dbReq.UserAddress)
	req.Header.Set("Name", dbReq.Name)
	req.Header.Set("Previous-Output-Token-Count", dbReq.PreviousOutputCount)
	req.Header.Set("Created-At", dbReq.CreatedAt)
	req.Header.Set("Nonce", dbReq.Nonce)
	req.Header.Set("Previous-Signature", dbReq.PreviousSignature)
	req.Header.Set("Signature", dbReq.Signature)

	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
