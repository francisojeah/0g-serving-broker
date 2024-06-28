package proxy

import (
	"fmt"
	"strings"

	"github.com/0glabs/0g-data-retrieve-agent/internal/contract"
	"github.com/0glabs/0g-data-retrieve-agent/internal/errors"
	"github.com/0glabs/0g-data-retrieve-agent/internal/model"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type chatBotRequest model.Request

// Generate used by the user agent to generate the next request,
// the data is sourced from the proxied request, previous response,
// and the signature from those data.
func (c *chatBotRequest) generate(reqBody map[string]interface{}, key, provider string) error {
	// TODO: Get metadata from DB instead of mock

	// For Chatbot: Read the request body to extract the `message` field
	message, ok := reqBody["message"].(string)
	if !ok || message == "" {
		return errors.New("Missing or invalid message field")

	}
	c.InputCount = fmt.Sprintf("%d", len(strings.Fields(message)))

	cReq := contract.Request{}
	if err := cReq.ConvertFromDB(model.Request(*c)); err != nil {
		return nil
	}

	sig, err := cReq.GetSignature(key, provider)
	if err != nil {
		return nil
	}
	c.Signature = hexutil.Encode(sig)

	return nil
}

// Called by the user agent, extract metadata from the request for subsequent signing.

func validate(dbReq model.Request, provider string) (bool, error) {
	// TODO: Verify the following fields in the request header:
	//  - inputToken matches the number of input tokens in the request body.
	//  - previousOutputCount matches the number of tokens returned in the previous response.
	//  - nonce is greater than the nonce of the previous request.

	cReq := contract.Request{}
	if err := cReq.ConvertFromDB(dbReq); err != nil {
		return false, errors.Wrap(err, "convert request from db schema to contract schema")
	}

	// https://github.com/ethereum/go-ethereum/issues/19751#issuecomment-504900739
	// Transform yellow paper V from 27/28 to 0/1
	if cReq.Signature[64] == 27 || cReq.Signature[64] == 28 {
		cReq.Signature[64] -= 27
	}

	prefixedHash, err := cReq.GetMessage(provider)
	if err != nil {
		return false, errors.Wrap(err, "Get Message")
	}

	recovered, err := crypto.SigToPub(prefixedHash.Bytes(), cReq.Signature)
	if err != nil {
		return false, errors.Wrap(err, "SigToPub")
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	return recoveredAddr == cReq.UserAddress, nil
}
