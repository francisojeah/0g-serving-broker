package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
)

//go:generate go run ./gen

type DataRetrieveContract struct {
	*Contract
	*DataRetrieve
}

func NewDataRetrieveContract(dataRetrieveAddress common.Address, clientWithSigner *web3go.Client, customGasPrice, customGasLimit uint64) (*DataRetrieveContract, error) {
	backend, signer := clientWithSigner.ToClientForContract()

	contract, err := newContract(clientWithSigner, signer, customGasPrice, customGasLimit)
	if err != nil {
		return nil, err
	}

	dataRetrieve, err := NewDataRetrieve(dataRetrieveAddress, backend)
	if err != nil {
		return nil, err
	}

	return &DataRetrieveContract{contract, dataRetrieve}, nil
}

type Contract struct {
	client  *web3go.Client
	account common.Address
	signer  bind.SignerFn

	customGasPrice uint64
	customGasLimit uint64
}

func newContract(clientWithSigner *web3go.Client, signerFn bind.SignerFn, customGasPrice, customGasLimit uint64) (*Contract, error) {
	signer, err := defaultSigner(clientWithSigner)
	if err != nil {
		return nil, err
	}

	return &Contract{
		client:         clientWithSigner,
		account:        signer.Address(),
		signer:         signerFn,
		customGasPrice: customGasPrice,
		customGasLimit: customGasLimit,
	}, nil
}

func (c *Contract) CreateTransactOpts() *bind.TransactOpts {
	var gasPrice *big.Int
	if c.customGasPrice > 0 {
		gasPrice = new(big.Int).SetUint64(c.customGasPrice)
	}

	return &bind.TransactOpts{
		From:     c.account,
		GasPrice: gasPrice,
		GasLimit: c.customGasLimit,
		Signer:   c.signer,
	}
}
