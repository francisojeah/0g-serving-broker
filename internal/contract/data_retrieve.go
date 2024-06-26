// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Request is an auto generated low-level Go binding around an user-defined struct.
type Request struct {
	UserAddress         common.Address
	Nonce               *big.Int
	Name                string
	InputCount          *big.Int
	PreviousOutputCount *big.Int
	PreviousSignature   []byte
	Signature           []byte
	CreatedAt           *big.Int
}

// RequestTrace is an auto generated low-level Go binding around an user-defined struct.
type RequestTrace struct {
	Requests []Request
}

// DataRetrieveMetaData contains all meta data concerning the DataRetrieve contract.
var DataRetrieveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RefundInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RefundLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RefundProcessed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ServiceNotexist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"UserAccountNotexists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RefundRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"service\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"service\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"ServiceUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"inputPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"addOrUpdateService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"depositFund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllServices\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"inputPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"outputPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"names\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"updatedAts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUserAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getService\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inputPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getUserAccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_locktime\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"processRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"inputCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"previousOutputCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"previousSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structRequest[]\",\"name\":\"requests\",\"type\":\"tuple[]\"}],\"internalType\":\"structRequestTrace[]\",\"name\":\"traces\",\"type\":\"tuple[]\"}],\"name\":\"settleFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_locktime\",\"type\":\"uint256\"}],\"name\":\"updateLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"inputCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"previousOutputCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"previousSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DataRetrieveABI is the input ABI used to generate the binding from.
// Deprecated: Use DataRetrieveMetaData.ABI instead.
var DataRetrieveABI = DataRetrieveMetaData.ABI

// DataRetrieve is an auto generated Go binding around an Ethereum contract.
type DataRetrieve struct {
	DataRetrieveCaller     // Read-only binding to the contract
	DataRetrieveTransactor // Write-only binding to the contract
	DataRetrieveFilterer   // Log filterer for contract events
}

// DataRetrieveCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataRetrieveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataRetrieveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataRetrieveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataRetrieveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataRetrieveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataRetrieveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataRetrieveSession struct {
	Contract     *DataRetrieve     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataRetrieveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataRetrieveCallerSession struct {
	Contract *DataRetrieveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DataRetrieveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataRetrieveTransactorSession struct {
	Contract     *DataRetrieveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DataRetrieveRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataRetrieveRaw struct {
	Contract *DataRetrieve // Generic contract binding to access the raw methods on
}

// DataRetrieveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataRetrieveCallerRaw struct {
	Contract *DataRetrieveCaller // Generic read-only contract binding to access the raw methods on
}

// DataRetrieveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataRetrieveTransactorRaw struct {
	Contract *DataRetrieveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataRetrieve creates a new instance of DataRetrieve, bound to a specific deployed contract.
func NewDataRetrieve(address common.Address, backend bind.ContractBackend) (*DataRetrieve, error) {
	contract, err := bindDataRetrieve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataRetrieve{DataRetrieveCaller: DataRetrieveCaller{contract: contract}, DataRetrieveTransactor: DataRetrieveTransactor{contract: contract}, DataRetrieveFilterer: DataRetrieveFilterer{contract: contract}}, nil
}

// NewDataRetrieveCaller creates a new read-only instance of DataRetrieve, bound to a specific deployed contract.
func NewDataRetrieveCaller(address common.Address, caller bind.ContractCaller) (*DataRetrieveCaller, error) {
	contract, err := bindDataRetrieve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataRetrieveCaller{contract: contract}, nil
}

// NewDataRetrieveTransactor creates a new write-only instance of DataRetrieve, bound to a specific deployed contract.
func NewDataRetrieveTransactor(address common.Address, transactor bind.ContractTransactor) (*DataRetrieveTransactor, error) {
	contract, err := bindDataRetrieve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataRetrieveTransactor{contract: contract}, nil
}

// NewDataRetrieveFilterer creates a new log filterer instance of DataRetrieve, bound to a specific deployed contract.
func NewDataRetrieveFilterer(address common.Address, filterer bind.ContractFilterer) (*DataRetrieveFilterer, error) {
	contract, err := bindDataRetrieve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataRetrieveFilterer{contract: contract}, nil
}

// bindDataRetrieve binds a generic wrapper to an already deployed contract.
func bindDataRetrieve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DataRetrieveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataRetrieve *DataRetrieveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataRetrieve.Contract.DataRetrieveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataRetrieve *DataRetrieveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataRetrieve.Contract.DataRetrieveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataRetrieve *DataRetrieveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataRetrieve.Contract.DataRetrieveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataRetrieve *DataRetrieveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataRetrieve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataRetrieve *DataRetrieveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataRetrieve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataRetrieve *DataRetrieveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataRetrieve.Contract.contract.Transact(opts, method, params...)
}

// GetAllServices is a free data retrieval call binding the contract method 0x21fe0f30.
//
// Solidity: function getAllServices() view returns(address[] addresses, uint256[] inputPrices, uint256[] outputPrices, string[] urls, string[] names, uint256[] updatedAts)
func (_DataRetrieve *DataRetrieveCaller) GetAllServices(opts *bind.CallOpts) (struct {
	Addresses    []common.Address
	InputPrices  []*big.Int
	OutputPrices []*big.Int
	Urls         []string
	Names        []string
	UpdatedAts   []*big.Int
}, error) {
	var out []interface{}
	err := _DataRetrieve.contract.Call(opts, &out, "getAllServices")

	outstruct := new(struct {
		Addresses    []common.Address
		InputPrices  []*big.Int
		OutputPrices []*big.Int
		Urls         []string
		Names        []string
		UpdatedAts   []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.InputPrices = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.OutputPrices = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Urls = *abi.ConvertType(out[3], new([]string)).(*[]string)
	outstruct.Names = *abi.ConvertType(out[4], new([]string)).(*[]string)
	outstruct.UpdatedAts = *abi.ConvertType(out[5], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAllServices is a free data retrieval call binding the contract method 0x21fe0f30.
//
// Solidity: function getAllServices() view returns(address[] addresses, uint256[] inputPrices, uint256[] outputPrices, string[] urls, string[] names, uint256[] updatedAts)
func (_DataRetrieve *DataRetrieveSession) GetAllServices() (struct {
	Addresses    []common.Address
	InputPrices  []*big.Int
	OutputPrices []*big.Int
	Urls         []string
	Names        []string
	UpdatedAts   []*big.Int
}, error) {
	return _DataRetrieve.Contract.GetAllServices(&_DataRetrieve.CallOpts)
}

// GetAllServices is a free data retrieval call binding the contract method 0x21fe0f30.
//
// Solidity: function getAllServices() view returns(address[] addresses, uint256[] inputPrices, uint256[] outputPrices, string[] urls, string[] names, uint256[] updatedAts)
func (_DataRetrieve *DataRetrieveCallerSession) GetAllServices() (struct {
	Addresses    []common.Address
	InputPrices  []*big.Int
	OutputPrices []*big.Int
	Urls         []string
	Names        []string
	UpdatedAts   []*big.Int
}, error) {
	return _DataRetrieve.Contract.GetAllServices(&_DataRetrieve.CallOpts)
}

// GetAllUserAccounts is a free data retrieval call binding the contract method 0x05756753.
//
// Solidity: function getAllUserAccounts() view returns(address[], address[], uint256[])
func (_DataRetrieve *DataRetrieveCaller) GetAllUserAccounts(opts *bind.CallOpts) ([]common.Address, []common.Address, []*big.Int, error) {
	var out []interface{}
	err := _DataRetrieve.contract.Call(opts, &out, "getAllUserAccounts")

	if err != nil {
		return *new([]common.Address), *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, err

}

// GetAllUserAccounts is a free data retrieval call binding the contract method 0x05756753.
//
// Solidity: function getAllUserAccounts() view returns(address[], address[], uint256[])
func (_DataRetrieve *DataRetrieveSession) GetAllUserAccounts() ([]common.Address, []common.Address, []*big.Int, error) {
	return _DataRetrieve.Contract.GetAllUserAccounts(&_DataRetrieve.CallOpts)
}

// GetAllUserAccounts is a free data retrieval call binding the contract method 0x05756753.
//
// Solidity: function getAllUserAccounts() view returns(address[], address[], uint256[])
func (_DataRetrieve *DataRetrieveCallerSession) GetAllUserAccounts() ([]common.Address, []common.Address, []*big.Int, error) {
	return _DataRetrieve.Contract.GetAllUserAccounts(&_DataRetrieve.CallOpts)
}

// GetService is a free data retrieval call binding the contract method 0x0e61d158.
//
// Solidity: function getService(address provider, string name) view returns(uint256 inputPrice, uint256 outputPrice, string url, uint256 updatedAt)
func (_DataRetrieve *DataRetrieveCaller) GetService(opts *bind.CallOpts, provider common.Address, name string) (struct {
	InputPrice  *big.Int
	OutputPrice *big.Int
	Url         string
	UpdatedAt   *big.Int
}, error) {
	var out []interface{}
	err := _DataRetrieve.contract.Call(opts, &out, "getService", provider, name)

	outstruct := new(struct {
		InputPrice  *big.Int
		OutputPrice *big.Int
		Url         string
		UpdatedAt   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InputPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OutputPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Url = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetService is a free data retrieval call binding the contract method 0x0e61d158.
//
// Solidity: function getService(address provider, string name) view returns(uint256 inputPrice, uint256 outputPrice, string url, uint256 updatedAt)
func (_DataRetrieve *DataRetrieveSession) GetService(provider common.Address, name string) (struct {
	InputPrice  *big.Int
	OutputPrice *big.Int
	Url         string
	UpdatedAt   *big.Int
}, error) {
	return _DataRetrieve.Contract.GetService(&_DataRetrieve.CallOpts, provider, name)
}

// GetService is a free data retrieval call binding the contract method 0x0e61d158.
//
// Solidity: function getService(address provider, string name) view returns(uint256 inputPrice, uint256 outputPrice, string url, uint256 updatedAt)
func (_DataRetrieve *DataRetrieveCallerSession) GetService(provider common.Address, name string) (struct {
	InputPrice  *big.Int
	OutputPrice *big.Int
	Url         string
	UpdatedAt   *big.Int
}, error) {
	return _DataRetrieve.Contract.GetService(&_DataRetrieve.CallOpts, provider, name)
}

// GetUserAccountBalance is a free data retrieval call binding the contract method 0xa418313c.
//
// Solidity: function getUserAccountBalance(address user, address provider) view returns(uint256)
func (_DataRetrieve *DataRetrieveCaller) GetUserAccountBalance(opts *bind.CallOpts, user common.Address, provider common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DataRetrieve.contract.Call(opts, &out, "getUserAccountBalance", user, provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAccountBalance is a free data retrieval call binding the contract method 0xa418313c.
//
// Solidity: function getUserAccountBalance(address user, address provider) view returns(uint256)
func (_DataRetrieve *DataRetrieveSession) GetUserAccountBalance(user common.Address, provider common.Address) (*big.Int, error) {
	return _DataRetrieve.Contract.GetUserAccountBalance(&_DataRetrieve.CallOpts, user, provider)
}

// GetUserAccountBalance is a free data retrieval call binding the contract method 0xa418313c.
//
// Solidity: function getUserAccountBalance(address user, address provider) view returns(uint256)
func (_DataRetrieve *DataRetrieveCallerSession) GetUserAccountBalance(user common.Address, provider common.Address) (*big.Int, error) {
	return _DataRetrieve.Contract.GetUserAccountBalance(&_DataRetrieve.CallOpts, user, provider)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_DataRetrieve *DataRetrieveCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataRetrieve.contract.Call(opts, &out, "lockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_DataRetrieve *DataRetrieveSession) LockTime() (*big.Int, error) {
	return _DataRetrieve.Contract.LockTime(&_DataRetrieve.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_DataRetrieve *DataRetrieveCallerSession) LockTime() (*big.Int, error) {
	return _DataRetrieve.Contract.LockTime(&_DataRetrieve.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataRetrieve *DataRetrieveCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataRetrieve.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataRetrieve *DataRetrieveSession) Owner() (common.Address, error) {
	return _DataRetrieve.Contract.Owner(&_DataRetrieve.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataRetrieve *DataRetrieveCallerSession) Owner() (common.Address, error) {
	return _DataRetrieve.Contract.Owner(&_DataRetrieve.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x994dd7da.
//
// Solidity: function verify((address,uint256,string,uint256,uint256,bytes,bytes,uint256) request) view returns(bool)
func (_DataRetrieve *DataRetrieveCaller) Verify(opts *bind.CallOpts, request Request) (bool, error) {
	var out []interface{}
	err := _DataRetrieve.contract.Call(opts, &out, "verify", request)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x994dd7da.
//
// Solidity: function verify((address,uint256,string,uint256,uint256,bytes,bytes,uint256) request) view returns(bool)
func (_DataRetrieve *DataRetrieveSession) Verify(request Request) (bool, error) {
	return _DataRetrieve.Contract.Verify(&_DataRetrieve.CallOpts, request)
}

// Verify is a free data retrieval call binding the contract method 0x994dd7da.
//
// Solidity: function verify((address,uint256,string,uint256,uint256,bytes,bytes,uint256) request) view returns(bool)
func (_DataRetrieve *DataRetrieveCallerSession) Verify(request Request) (bool, error) {
	return _DataRetrieve.Contract.Verify(&_DataRetrieve.CallOpts, request)
}

// AddOrUpdateService is a paid mutator transaction binding the contract method 0xf1b6bec8.
//
// Solidity: function addOrUpdateService(string name, uint256 inputPrice, uint256 outputPrice, string url) returns()
func (_DataRetrieve *DataRetrieveTransactor) AddOrUpdateService(opts *bind.TransactOpts, name string, inputPrice *big.Int, outputPrice *big.Int, url string) (*types.Transaction, error) {
	return _DataRetrieve.contract.Transact(opts, "addOrUpdateService", name, inputPrice, outputPrice, url)
}

// AddOrUpdateService is a paid mutator transaction binding the contract method 0xf1b6bec8.
//
// Solidity: function addOrUpdateService(string name, uint256 inputPrice, uint256 outputPrice, string url) returns()
func (_DataRetrieve *DataRetrieveSession) AddOrUpdateService(name string, inputPrice *big.Int, outputPrice *big.Int, url string) (*types.Transaction, error) {
	return _DataRetrieve.Contract.AddOrUpdateService(&_DataRetrieve.TransactOpts, name, inputPrice, outputPrice, url)
}

// AddOrUpdateService is a paid mutator transaction binding the contract method 0xf1b6bec8.
//
// Solidity: function addOrUpdateService(string name, uint256 inputPrice, uint256 outputPrice, string url) returns()
func (_DataRetrieve *DataRetrieveTransactorSession) AddOrUpdateService(name string, inputPrice *big.Int, outputPrice *big.Int, url string) (*types.Transaction, error) {
	return _DataRetrieve.Contract.AddOrUpdateService(&_DataRetrieve.TransactOpts, name, inputPrice, outputPrice, url)
}

// DepositFund is a paid mutator transaction binding the contract method 0xe12d4a52.
//
// Solidity: function depositFund(address provider) payable returns()
func (_DataRetrieve *DataRetrieveTransactor) DepositFund(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _DataRetrieve.contract.Transact(opts, "depositFund", provider)
}

// DepositFund is a paid mutator transaction binding the contract method 0xe12d4a52.
//
// Solidity: function depositFund(address provider) payable returns()
func (_DataRetrieve *DataRetrieveSession) DepositFund(provider common.Address) (*types.Transaction, error) {
	return _DataRetrieve.Contract.DepositFund(&_DataRetrieve.TransactOpts, provider)
}

// DepositFund is a paid mutator transaction binding the contract method 0xe12d4a52.
//
// Solidity: function depositFund(address provider) payable returns()
func (_DataRetrieve *DataRetrieveTransactorSession) DepositFund(provider common.Address) (*types.Transaction, error) {
	return _DataRetrieve.Contract.DepositFund(&_DataRetrieve.TransactOpts, provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _locktime) returns()
func (_DataRetrieve *DataRetrieveTransactor) Initialize(opts *bind.TransactOpts, _locktime *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.contract.Transact(opts, "initialize", _locktime)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _locktime) returns()
func (_DataRetrieve *DataRetrieveSession) Initialize(_locktime *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.Contract.Initialize(&_DataRetrieve.TransactOpts, _locktime)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _locktime) returns()
func (_DataRetrieve *DataRetrieveTransactorSession) Initialize(_locktime *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.Contract.Initialize(&_DataRetrieve.TransactOpts, _locktime)
}

// ProcessRefund is a paid mutator transaction binding the contract method 0x3878c5ef.
//
// Solidity: function processRefund(address provider, uint256 index) returns()
func (_DataRetrieve *DataRetrieveTransactor) ProcessRefund(opts *bind.TransactOpts, provider common.Address, index *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.contract.Transact(opts, "processRefund", provider, index)
}

// ProcessRefund is a paid mutator transaction binding the contract method 0x3878c5ef.
//
// Solidity: function processRefund(address provider, uint256 index) returns()
func (_DataRetrieve *DataRetrieveSession) ProcessRefund(provider common.Address, index *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.Contract.ProcessRefund(&_DataRetrieve.TransactOpts, provider, index)
}

// ProcessRefund is a paid mutator transaction binding the contract method 0x3878c5ef.
//
// Solidity: function processRefund(address provider, uint256 index) returns()
func (_DataRetrieve *DataRetrieveTransactorSession) ProcessRefund(provider common.Address, index *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.Contract.ProcessRefund(&_DataRetrieve.TransactOpts, provider, index)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string name) returns()
func (_DataRetrieve *DataRetrieveTransactor) RemoveService(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _DataRetrieve.contract.Transact(opts, "removeService", name)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string name) returns()
func (_DataRetrieve *DataRetrieveSession) RemoveService(name string) (*types.Transaction, error) {
	return _DataRetrieve.Contract.RemoveService(&_DataRetrieve.TransactOpts, name)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string name) returns()
func (_DataRetrieve *DataRetrieveTransactorSession) RemoveService(name string) (*types.Transaction, error) {
	return _DataRetrieve.Contract.RemoveService(&_DataRetrieve.TransactOpts, name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataRetrieve *DataRetrieveTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataRetrieve.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataRetrieve *DataRetrieveSession) RenounceOwnership() (*types.Transaction, error) {
	return _DataRetrieve.Contract.RenounceOwnership(&_DataRetrieve.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DataRetrieve *DataRetrieveTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DataRetrieve.Contract.RenounceOwnership(&_DataRetrieve.TransactOpts)
}

// RequestRefund is a paid mutator transaction binding the contract method 0x99652de7.
//
// Solidity: function requestRefund(address provider, uint256 amount) returns()
func (_DataRetrieve *DataRetrieveTransactor) RequestRefund(opts *bind.TransactOpts, provider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.contract.Transact(opts, "requestRefund", provider, amount)
}

// RequestRefund is a paid mutator transaction binding the contract method 0x99652de7.
//
// Solidity: function requestRefund(address provider, uint256 amount) returns()
func (_DataRetrieve *DataRetrieveSession) RequestRefund(provider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.Contract.RequestRefund(&_DataRetrieve.TransactOpts, provider, amount)
}

// RequestRefund is a paid mutator transaction binding the contract method 0x99652de7.
//
// Solidity: function requestRefund(address provider, uint256 amount) returns()
func (_DataRetrieve *DataRetrieveTransactorSession) RequestRefund(provider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.Contract.RequestRefund(&_DataRetrieve.TransactOpts, provider, amount)
}

// SettleFees is a paid mutator transaction binding the contract method 0x805edbef.
//
// Solidity: function settleFees(((address,uint256,string,uint256,uint256,bytes,bytes,uint256)[])[] traces) returns()
func (_DataRetrieve *DataRetrieveTransactor) SettleFees(opts *bind.TransactOpts, traces []RequestTrace) (*types.Transaction, error) {
	return _DataRetrieve.contract.Transact(opts, "settleFees", traces)
}

// SettleFees is a paid mutator transaction binding the contract method 0x805edbef.
//
// Solidity: function settleFees(((address,uint256,string,uint256,uint256,bytes,bytes,uint256)[])[] traces) returns()
func (_DataRetrieve *DataRetrieveSession) SettleFees(traces []RequestTrace) (*types.Transaction, error) {
	return _DataRetrieve.Contract.SettleFees(&_DataRetrieve.TransactOpts, traces)
}

// SettleFees is a paid mutator transaction binding the contract method 0x805edbef.
//
// Solidity: function settleFees(((address,uint256,string,uint256,uint256,bytes,bytes,uint256)[])[] traces) returns()
func (_DataRetrieve *DataRetrieveTransactorSession) SettleFees(traces []RequestTrace) (*types.Transaction, error) {
	return _DataRetrieve.Contract.SettleFees(&_DataRetrieve.TransactOpts, traces)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DataRetrieve *DataRetrieveTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DataRetrieve.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DataRetrieve *DataRetrieveSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DataRetrieve.Contract.TransferOwnership(&_DataRetrieve.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DataRetrieve *DataRetrieveTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DataRetrieve.Contract.TransferOwnership(&_DataRetrieve.TransactOpts, newOwner)
}

// UpdateLockTime is a paid mutator transaction binding the contract method 0xfbfa4e11.
//
// Solidity: function updateLockTime(uint256 _locktime) returns()
func (_DataRetrieve *DataRetrieveTransactor) UpdateLockTime(opts *bind.TransactOpts, _locktime *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.contract.Transact(opts, "updateLockTime", _locktime)
}

// UpdateLockTime is a paid mutator transaction binding the contract method 0xfbfa4e11.
//
// Solidity: function updateLockTime(uint256 _locktime) returns()
func (_DataRetrieve *DataRetrieveSession) UpdateLockTime(_locktime *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.Contract.UpdateLockTime(&_DataRetrieve.TransactOpts, _locktime)
}

// UpdateLockTime is a paid mutator transaction binding the contract method 0xfbfa4e11.
//
// Solidity: function updateLockTime(uint256 _locktime) returns()
func (_DataRetrieve *DataRetrieveTransactorSession) UpdateLockTime(_locktime *big.Int) (*types.Transaction, error) {
	return _DataRetrieve.Contract.UpdateLockTime(&_DataRetrieve.TransactOpts, _locktime)
}

// DataRetrieveBalanceUpdatedIterator is returned from FilterBalanceUpdated and is used to iterate over the raw logs and unpacked data for BalanceUpdated events raised by the DataRetrieve contract.
type DataRetrieveBalanceUpdatedIterator struct {
	Event *DataRetrieveBalanceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataRetrieveBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRetrieveBalanceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataRetrieveBalanceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataRetrieveBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRetrieveBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRetrieveBalanceUpdated represents a BalanceUpdated event raised by the DataRetrieve contract.
type DataRetrieveBalanceUpdated struct {
	User     common.Address
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBalanceUpdated is a free log retrieval operation binding the contract event 0x2047d1633ff7768462ae07d28cb16e484203bfd6d85ce832494270ebcd9081a2.
//
// Solidity: event BalanceUpdated(address indexed user, address indexed provider, uint256 amount)
func (_DataRetrieve *DataRetrieveFilterer) FilterBalanceUpdated(opts *bind.FilterOpts, user []common.Address, provider []common.Address) (*DataRetrieveBalanceUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _DataRetrieve.contract.FilterLogs(opts, "BalanceUpdated", userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &DataRetrieveBalanceUpdatedIterator{contract: _DataRetrieve.contract, event: "BalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchBalanceUpdated is a free log subscription operation binding the contract event 0x2047d1633ff7768462ae07d28cb16e484203bfd6d85ce832494270ebcd9081a2.
//
// Solidity: event BalanceUpdated(address indexed user, address indexed provider, uint256 amount)
func (_DataRetrieve *DataRetrieveFilterer) WatchBalanceUpdated(opts *bind.WatchOpts, sink chan<- *DataRetrieveBalanceUpdated, user []common.Address, provider []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _DataRetrieve.contract.WatchLogs(opts, "BalanceUpdated", userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRetrieveBalanceUpdated)
				if err := _DataRetrieve.contract.UnpackLog(event, "BalanceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBalanceUpdated is a log parse operation binding the contract event 0x2047d1633ff7768462ae07d28cb16e484203bfd6d85ce832494270ebcd9081a2.
//
// Solidity: event BalanceUpdated(address indexed user, address indexed provider, uint256 amount)
func (_DataRetrieve *DataRetrieveFilterer) ParseBalanceUpdated(log types.Log) (*DataRetrieveBalanceUpdated, error) {
	event := new(DataRetrieveBalanceUpdated)
	if err := _DataRetrieve.contract.UnpackLog(event, "BalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataRetrieveInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DataRetrieve contract.
type DataRetrieveInitializedIterator struct {
	Event *DataRetrieveInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataRetrieveInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRetrieveInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataRetrieveInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataRetrieveInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRetrieveInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRetrieveInitialized represents a Initialized event raised by the DataRetrieve contract.
type DataRetrieveInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DataRetrieve *DataRetrieveFilterer) FilterInitialized(opts *bind.FilterOpts) (*DataRetrieveInitializedIterator, error) {

	logs, sub, err := _DataRetrieve.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DataRetrieveInitializedIterator{contract: _DataRetrieve.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DataRetrieve *DataRetrieveFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DataRetrieveInitialized) (event.Subscription, error) {

	logs, sub, err := _DataRetrieve.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRetrieveInitialized)
				if err := _DataRetrieve.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DataRetrieve *DataRetrieveFilterer) ParseInitialized(log types.Log) (*DataRetrieveInitialized, error) {
	event := new(DataRetrieveInitialized)
	if err := _DataRetrieve.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataRetrieveOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DataRetrieve contract.
type DataRetrieveOwnershipTransferredIterator struct {
	Event *DataRetrieveOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataRetrieveOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRetrieveOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataRetrieveOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataRetrieveOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRetrieveOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRetrieveOwnershipTransferred represents a OwnershipTransferred event raised by the DataRetrieve contract.
type DataRetrieveOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DataRetrieve *DataRetrieveFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DataRetrieveOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DataRetrieve.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DataRetrieveOwnershipTransferredIterator{contract: _DataRetrieve.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DataRetrieve *DataRetrieveFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DataRetrieveOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DataRetrieve.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRetrieveOwnershipTransferred)
				if err := _DataRetrieve.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DataRetrieve *DataRetrieveFilterer) ParseOwnershipTransferred(log types.Log) (*DataRetrieveOwnershipTransferred, error) {
	event := new(DataRetrieveOwnershipTransferred)
	if err := _DataRetrieve.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataRetrieveRefundProcessedIterator is returned from FilterRefundProcessed and is used to iterate over the raw logs and unpacked data for RefundProcessed events raised by the DataRetrieve contract.
type DataRetrieveRefundProcessedIterator struct {
	Event *DataRetrieveRefundProcessed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataRetrieveRefundProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRetrieveRefundProcessed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataRetrieveRefundProcessed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataRetrieveRefundProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRetrieveRefundProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRetrieveRefundProcessed represents a RefundProcessed event raised by the DataRetrieve contract.
type DataRetrieveRefundProcessed struct {
	User     common.Address
	Provider common.Address
	Index    *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefundProcessed is a free log retrieval operation binding the contract event 0x6de2d5fa0b9cac9e48f300a350314e447648d158a5067a56513a46396f1b638a.
//
// Solidity: event RefundProcessed(address indexed user, address indexed provider, uint256 indexed index, uint256 amount)
func (_DataRetrieve *DataRetrieveFilterer) FilterRefundProcessed(opts *bind.FilterOpts, user []common.Address, provider []common.Address, index []*big.Int) (*DataRetrieveRefundProcessedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _DataRetrieve.contract.FilterLogs(opts, "RefundProcessed", userRule, providerRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &DataRetrieveRefundProcessedIterator{contract: _DataRetrieve.contract, event: "RefundProcessed", logs: logs, sub: sub}, nil
}

// WatchRefundProcessed is a free log subscription operation binding the contract event 0x6de2d5fa0b9cac9e48f300a350314e447648d158a5067a56513a46396f1b638a.
//
// Solidity: event RefundProcessed(address indexed user, address indexed provider, uint256 indexed index, uint256 amount)
func (_DataRetrieve *DataRetrieveFilterer) WatchRefundProcessed(opts *bind.WatchOpts, sink chan<- *DataRetrieveRefundProcessed, user []common.Address, provider []common.Address, index []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _DataRetrieve.contract.WatchLogs(opts, "RefundProcessed", userRule, providerRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRetrieveRefundProcessed)
				if err := _DataRetrieve.contract.UnpackLog(event, "RefundProcessed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundProcessed is a log parse operation binding the contract event 0x6de2d5fa0b9cac9e48f300a350314e447648d158a5067a56513a46396f1b638a.
//
// Solidity: event RefundProcessed(address indexed user, address indexed provider, uint256 indexed index, uint256 amount)
func (_DataRetrieve *DataRetrieveFilterer) ParseRefundProcessed(log types.Log) (*DataRetrieveRefundProcessed, error) {
	event := new(DataRetrieveRefundProcessed)
	if err := _DataRetrieve.contract.UnpackLog(event, "RefundProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataRetrieveRefundRequestedIterator is returned from FilterRefundRequested and is used to iterate over the raw logs and unpacked data for RefundRequested events raised by the DataRetrieve contract.
type DataRetrieveRefundRequestedIterator struct {
	Event *DataRetrieveRefundRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataRetrieveRefundRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRetrieveRefundRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataRetrieveRefundRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataRetrieveRefundRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRetrieveRefundRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRetrieveRefundRequested represents a RefundRequested event raised by the DataRetrieve contract.
type DataRetrieveRefundRequested struct {
	User      common.Address
	Provider  common.Address
	Index     *big.Int
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundRequested is a free log retrieval operation binding the contract event 0xbca585bae445c22bd427d9aa104c4f02cbfcb713869df122bbb797c43356d2e5.
//
// Solidity: event RefundRequested(address indexed user, address indexed provider, uint256 indexed index, uint256 amount, uint256 timestamp)
func (_DataRetrieve *DataRetrieveFilterer) FilterRefundRequested(opts *bind.FilterOpts, user []common.Address, provider []common.Address, index []*big.Int) (*DataRetrieveRefundRequestedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _DataRetrieve.contract.FilterLogs(opts, "RefundRequested", userRule, providerRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &DataRetrieveRefundRequestedIterator{contract: _DataRetrieve.contract, event: "RefundRequested", logs: logs, sub: sub}, nil
}

// WatchRefundRequested is a free log subscription operation binding the contract event 0xbca585bae445c22bd427d9aa104c4f02cbfcb713869df122bbb797c43356d2e5.
//
// Solidity: event RefundRequested(address indexed user, address indexed provider, uint256 indexed index, uint256 amount, uint256 timestamp)
func (_DataRetrieve *DataRetrieveFilterer) WatchRefundRequested(opts *bind.WatchOpts, sink chan<- *DataRetrieveRefundRequested, user []common.Address, provider []common.Address, index []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _DataRetrieve.contract.WatchLogs(opts, "RefundRequested", userRule, providerRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRetrieveRefundRequested)
				if err := _DataRetrieve.contract.UnpackLog(event, "RefundRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundRequested is a log parse operation binding the contract event 0xbca585bae445c22bd427d9aa104c4f02cbfcb713869df122bbb797c43356d2e5.
//
// Solidity: event RefundRequested(address indexed user, address indexed provider, uint256 indexed index, uint256 amount, uint256 timestamp)
func (_DataRetrieve *DataRetrieveFilterer) ParseRefundRequested(log types.Log) (*DataRetrieveRefundRequested, error) {
	event := new(DataRetrieveRefundRequested)
	if err := _DataRetrieve.contract.UnpackLog(event, "RefundRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataRetrieveServiceRemovedIterator is returned from FilterServiceRemoved and is used to iterate over the raw logs and unpacked data for ServiceRemoved events raised by the DataRetrieve contract.
type DataRetrieveServiceRemovedIterator struct {
	Event *DataRetrieveServiceRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataRetrieveServiceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRetrieveServiceRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataRetrieveServiceRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataRetrieveServiceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRetrieveServiceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRetrieveServiceRemoved represents a ServiceRemoved event raised by the DataRetrieve contract.
type DataRetrieveServiceRemoved struct {
	Service common.Address
	Name    common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterServiceRemoved is a free log retrieval operation binding the contract event 0x68026479739e3662c0651578523384b94455e79bfb701ce111a3164591ceba73.
//
// Solidity: event ServiceRemoved(address indexed service, string indexed name)
func (_DataRetrieve *DataRetrieveFilterer) FilterServiceRemoved(opts *bind.FilterOpts, service []common.Address, name []string) (*DataRetrieveServiceRemovedIterator, error) {

	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _DataRetrieve.contract.FilterLogs(opts, "ServiceRemoved", serviceRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &DataRetrieveServiceRemovedIterator{contract: _DataRetrieve.contract, event: "ServiceRemoved", logs: logs, sub: sub}, nil
}

// WatchServiceRemoved is a free log subscription operation binding the contract event 0x68026479739e3662c0651578523384b94455e79bfb701ce111a3164591ceba73.
//
// Solidity: event ServiceRemoved(address indexed service, string indexed name)
func (_DataRetrieve *DataRetrieveFilterer) WatchServiceRemoved(opts *bind.WatchOpts, sink chan<- *DataRetrieveServiceRemoved, service []common.Address, name []string) (event.Subscription, error) {

	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _DataRetrieve.contract.WatchLogs(opts, "ServiceRemoved", serviceRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRetrieveServiceRemoved)
				if err := _DataRetrieve.contract.UnpackLog(event, "ServiceRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceRemoved is a log parse operation binding the contract event 0x68026479739e3662c0651578523384b94455e79bfb701ce111a3164591ceba73.
//
// Solidity: event ServiceRemoved(address indexed service, string indexed name)
func (_DataRetrieve *DataRetrieveFilterer) ParseServiceRemoved(log types.Log) (*DataRetrieveServiceRemoved, error) {
	event := new(DataRetrieveServiceRemoved)
	if err := _DataRetrieve.contract.UnpackLog(event, "ServiceRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataRetrieveServiceUpdatedIterator is returned from FilterServiceUpdated and is used to iterate over the raw logs and unpacked data for ServiceUpdated events raised by the DataRetrieve contract.
type DataRetrieveServiceUpdatedIterator struct {
	Event *DataRetrieveServiceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DataRetrieveServiceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataRetrieveServiceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DataRetrieveServiceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DataRetrieveServiceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataRetrieveServiceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataRetrieveServiceUpdated represents a ServiceUpdated event raised by the DataRetrieve contract.
type DataRetrieveServiceUpdated struct {
	Service     common.Address
	Name        common.Hash
	InputPrice  *big.Int
	OutputPrice *big.Int
	Url         string
	UpdatedAt   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceUpdated is a free log retrieval operation binding the contract event 0xbbd2877472c4e0f3c638d6ed0d10153cebb87e1f44aca504cec5cc64abefc002.
//
// Solidity: event ServiceUpdated(address indexed service, string indexed name, uint256 inputPrice, uint256 outputPrice, string url, uint256 updatedAt)
func (_DataRetrieve *DataRetrieveFilterer) FilterServiceUpdated(opts *bind.FilterOpts, service []common.Address, name []string) (*DataRetrieveServiceUpdatedIterator, error) {

	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _DataRetrieve.contract.FilterLogs(opts, "ServiceUpdated", serviceRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &DataRetrieveServiceUpdatedIterator{contract: _DataRetrieve.contract, event: "ServiceUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceUpdated is a free log subscription operation binding the contract event 0xbbd2877472c4e0f3c638d6ed0d10153cebb87e1f44aca504cec5cc64abefc002.
//
// Solidity: event ServiceUpdated(address indexed service, string indexed name, uint256 inputPrice, uint256 outputPrice, string url, uint256 updatedAt)
func (_DataRetrieve *DataRetrieveFilterer) WatchServiceUpdated(opts *bind.WatchOpts, sink chan<- *DataRetrieveServiceUpdated, service []common.Address, name []string) (event.Subscription, error) {

	var serviceRule []interface{}
	for _, serviceItem := range service {
		serviceRule = append(serviceRule, serviceItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _DataRetrieve.contract.WatchLogs(opts, "ServiceUpdated", serviceRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataRetrieveServiceUpdated)
				if err := _DataRetrieve.contract.UnpackLog(event, "ServiceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceUpdated is a log parse operation binding the contract event 0xbbd2877472c4e0f3c638d6ed0d10153cebb87e1f44aca504cec5cc64abefc002.
//
// Solidity: event ServiceUpdated(address indexed service, string indexed name, uint256 inputPrice, uint256 outputPrice, string url, uint256 updatedAt)
func (_DataRetrieve *DataRetrieveFilterer) ParseServiceUpdated(log types.Log) (*DataRetrieveServiceUpdated, error) {
	event := new(DataRetrieveServiceUpdated)
	if err := _DataRetrieve.contract.UnpackLog(event, "ServiceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
