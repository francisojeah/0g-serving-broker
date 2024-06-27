package contract

import (
	"bytes"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/0glabs/0g-data-retrieve-agent/internal/errors"
	"github.com/0glabs/0g-data-retrieve-agent/internal/model"
)

func (r Request) ConvertToDb() model.Request {
	ret := model.Request{
		CreatedAt:           r.CreatedAt.String(),
		UserAddress:         r.UserAddress.String(),
		Nonce:               r.Nonce.String(),
		Name:                r.Name,
		InputCount:          r.InputCount.String(),
		PreviousOutputCount: r.PreviousOutputCount.String(),
		PreviousSignature:   hexutil.Encode(r.PreviousSignature),
		Signature:           hexutil.Encode(r.Signature),
	}

	return ret
}

func (r Request) GetMessage(serviceProviderAddress string) (common.Hash, error) {
	buf := new(bytes.Buffer)
	buf.Write(common.HexToAddress(serviceProviderAddress).Bytes())
	buf.Write(r.UserAddress.Bytes())
	buf.Write([]byte(r.Name))
	buf.Write(common.LeftPadBytes(r.InputCount.Bytes(), 32))
	buf.Write(common.LeftPadBytes(r.PreviousOutputCount.Bytes(), 32))
	buf.Write(r.PreviousSignature)
	buf.Write(common.LeftPadBytes(r.Nonce.Bytes(), 32))
	buf.Write(common.LeftPadBytes(r.CreatedAt.Bytes(), 32))

	msg := crypto.Keccak256Hash(buf.Bytes())
	prefixedMsg := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), msg.Bytes())

	return prefixedMsg, nil
}

func (r Request) GetSignature(keyHex string, provide string) ([]byte, error) {
	key, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		return nil, err
	}

	msg, err := r.GetMessage(provide)
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(msg.Bytes(), key)
	if err != nil {
		return nil, err
	}

	// https://github.com/ethereum/go-ethereum/issues/19751#issuecomment-504900739
	if sig[64] == 0 || sig[64] == 1 {
		sig[64] += 27
	}

	return sig, nil
}

func (r *Request) ConvertFromDB(req model.Request) error {
	userAddress := common.HexToAddress(req.UserAddress)
	nonce, ok := new(big.Int).SetString(req.Nonce, 10)
	if !ok {
		return errors.Wrapf(errors.New("invalid Nonce"), "converted from %s", req.Nonce)
	}
	inputCount, ok := new(big.Int).SetString(req.InputCount, 10)
	if !ok {
		return errors.Wrapf(errors.New("invalid InputCount"), "converted from %s", req.InputCount)
	}
	previousOutputCount, ok := new(big.Int).SetString(req.PreviousOutputCount, 10)
	if !ok {
		return errors.Wrapf(errors.New("invalid PreviousOutputCount"), "converted from %s", req.PreviousOutputCount)
	}
	previousSignature, err := hexutil.Decode(req.PreviousSignature)
	if err != nil {
		return errors.Wrapf(err, "convert PreviousSignature %s", req.PreviousSignature)
	}
	createdAt, err := time.Parse(time.RFC3339, req.CreatedAt)
	if err != nil {
		return errors.Wrapf(err, "convert createdAt %s", req.CreatedAt)
	}
	var signature []byte
	if req.Signature != "" {
		signature, err = hexutil.Decode(req.Signature)
		if err != nil {
			return errors.Wrapf(err, "convert signature from request: %s", req.Signature)
		}
	}

	r.UserAddress = userAddress
	r.Nonce = nonce
	r.Name = req.Name
	r.InputCount = inputCount
	r.PreviousOutputCount = previousOutputCount
	r.PreviousSignature = previousSignature
	r.Signature = signature
	r.CreatedAt = big.NewInt(createdAt.Unix())

	return nil
}
