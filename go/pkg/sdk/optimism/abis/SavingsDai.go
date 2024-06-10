// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimismContracts

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

// SavingsDaiMetaData contains all meta data concerning the SavingsDai contract.
var SavingsDaiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REMOTE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SavingsDaiABI is the input ABI used to generate the binding from.
// Deprecated: Use SavingsDaiMetaData.ABI instead.
var SavingsDaiABI = SavingsDaiMetaData.ABI

// SavingsDai is an auto generated Go binding around an Ethereum contract.
type SavingsDai struct {
	SavingsDaiCaller     // Read-only binding to the contract
	SavingsDaiTransactor // Write-only binding to the contract
	SavingsDaiFilterer   // Log filterer for contract events
}

// SavingsDaiCaller is an auto generated read-only Go binding around an Ethereum contract.
type SavingsDaiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavingsDaiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SavingsDaiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavingsDaiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SavingsDaiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavingsDaiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SavingsDaiSession struct {
	Contract     *SavingsDai       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SavingsDaiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SavingsDaiCallerSession struct {
	Contract *SavingsDaiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SavingsDaiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SavingsDaiTransactorSession struct {
	Contract     *SavingsDaiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SavingsDaiRaw is an auto generated low-level Go binding around an Ethereum contract.
type SavingsDaiRaw struct {
	Contract *SavingsDai // Generic contract binding to access the raw methods on
}

// SavingsDaiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SavingsDaiCallerRaw struct {
	Contract *SavingsDaiCaller // Generic read-only contract binding to access the raw methods on
}

// SavingsDaiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SavingsDaiTransactorRaw struct {
	Contract *SavingsDaiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSavingsDai creates a new instance of SavingsDai, bound to a specific deployed contract.
func NewSavingsDai(address common.Address, backend bind.ContractBackend) (*SavingsDai, error) {
	contract, err := bindSavingsDai(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SavingsDai{SavingsDaiCaller: SavingsDaiCaller{contract: contract}, SavingsDaiTransactor: SavingsDaiTransactor{contract: contract}, SavingsDaiFilterer: SavingsDaiFilterer{contract: contract}}, nil
}

// NewSavingsDaiCaller creates a new read-only instance of SavingsDai, bound to a specific deployed contract.
func NewSavingsDaiCaller(address common.Address, caller bind.ContractCaller) (*SavingsDaiCaller, error) {
	contract, err := bindSavingsDai(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SavingsDaiCaller{contract: contract}, nil
}

// NewSavingsDaiTransactor creates a new write-only instance of SavingsDai, bound to a specific deployed contract.
func NewSavingsDaiTransactor(address common.Address, transactor bind.ContractTransactor) (*SavingsDaiTransactor, error) {
	contract, err := bindSavingsDai(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SavingsDaiTransactor{contract: contract}, nil
}

// NewSavingsDaiFilterer creates a new log filterer instance of SavingsDai, bound to a specific deployed contract.
func NewSavingsDaiFilterer(address common.Address, filterer bind.ContractFilterer) (*SavingsDaiFilterer, error) {
	contract, err := bindSavingsDai(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SavingsDaiFilterer{contract: contract}, nil
}

// bindSavingsDai binds a generic wrapper to an already deployed contract.
func bindSavingsDai(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SavingsDaiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SavingsDai *SavingsDaiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SavingsDai.Contract.SavingsDaiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SavingsDai *SavingsDaiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SavingsDai.Contract.SavingsDaiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SavingsDai *SavingsDaiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SavingsDai.Contract.SavingsDaiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SavingsDai *SavingsDaiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SavingsDai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SavingsDai *SavingsDaiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SavingsDai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SavingsDai *SavingsDaiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SavingsDai.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_SavingsDai *SavingsDaiCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_SavingsDai *SavingsDaiSession) BRIDGE() (common.Address, error) {
	return _SavingsDai.Contract.BRIDGE(&_SavingsDai.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_SavingsDai *SavingsDaiCallerSession) BRIDGE() (common.Address, error) {
	return _SavingsDai.Contract.BRIDGE(&_SavingsDai.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_SavingsDai *SavingsDaiCaller) REMOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "REMOTE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_SavingsDai *SavingsDaiSession) REMOTETOKEN() (common.Address, error) {
	return _SavingsDai.Contract.REMOTETOKEN(&_SavingsDai.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_SavingsDai *SavingsDaiCallerSession) REMOTETOKEN() (common.Address, error) {
	return _SavingsDai.Contract.REMOTETOKEN(&_SavingsDai.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SavingsDai *SavingsDaiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SavingsDai *SavingsDaiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SavingsDai.Contract.Allowance(&_SavingsDai.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SavingsDai *SavingsDaiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SavingsDai.Contract.Allowance(&_SavingsDai.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SavingsDai *SavingsDaiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SavingsDai *SavingsDaiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SavingsDai.Contract.BalanceOf(&_SavingsDai.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SavingsDai *SavingsDaiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SavingsDai.Contract.BalanceOf(&_SavingsDai.CallOpts, account)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SavingsDai *SavingsDaiCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SavingsDai *SavingsDaiSession) Bridge() (common.Address, error) {
	return _SavingsDai.Contract.Bridge(&_SavingsDai.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SavingsDai *SavingsDaiCallerSession) Bridge() (common.Address, error) {
	return _SavingsDai.Contract.Bridge(&_SavingsDai.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SavingsDai *SavingsDaiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SavingsDai *SavingsDaiSession) Decimals() (uint8, error) {
	return _SavingsDai.Contract.Decimals(&_SavingsDai.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SavingsDai *SavingsDaiCallerSession) Decimals() (uint8, error) {
	return _SavingsDai.Contract.Decimals(&_SavingsDai.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_SavingsDai *SavingsDaiCaller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_SavingsDai *SavingsDaiSession) L1Token() (common.Address, error) {
	return _SavingsDai.Contract.L1Token(&_SavingsDai.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_SavingsDai *SavingsDaiCallerSession) L1Token() (common.Address, error) {
	return _SavingsDai.Contract.L1Token(&_SavingsDai.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_SavingsDai *SavingsDaiCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_SavingsDai *SavingsDaiSession) L2Bridge() (common.Address, error) {
	return _SavingsDai.Contract.L2Bridge(&_SavingsDai.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_SavingsDai *SavingsDaiCallerSession) L2Bridge() (common.Address, error) {
	return _SavingsDai.Contract.L2Bridge(&_SavingsDai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SavingsDai *SavingsDaiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SavingsDai *SavingsDaiSession) Name() (string, error) {
	return _SavingsDai.Contract.Name(&_SavingsDai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SavingsDai *SavingsDaiCallerSession) Name() (string, error) {
	return _SavingsDai.Contract.Name(&_SavingsDai.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_SavingsDai *SavingsDaiCaller) RemoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "remoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_SavingsDai *SavingsDaiSession) RemoteToken() (common.Address, error) {
	return _SavingsDai.Contract.RemoteToken(&_SavingsDai.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_SavingsDai *SavingsDaiCallerSession) RemoteToken() (common.Address, error) {
	return _SavingsDai.Contract.RemoteToken(&_SavingsDai.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_SavingsDai *SavingsDaiCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_SavingsDai *SavingsDaiSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SavingsDai.Contract.SupportsInterface(&_SavingsDai.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_SavingsDai *SavingsDaiCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SavingsDai.Contract.SupportsInterface(&_SavingsDai.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SavingsDai *SavingsDaiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SavingsDai *SavingsDaiSession) Symbol() (string, error) {
	return _SavingsDai.Contract.Symbol(&_SavingsDai.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SavingsDai *SavingsDaiCallerSession) Symbol() (string, error) {
	return _SavingsDai.Contract.Symbol(&_SavingsDai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SavingsDai *SavingsDaiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SavingsDai *SavingsDaiSession) TotalSupply() (*big.Int, error) {
	return _SavingsDai.Contract.TotalSupply(&_SavingsDai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SavingsDai *SavingsDaiCallerSession) TotalSupply() (*big.Int, error) {
	return _SavingsDai.Contract.TotalSupply(&_SavingsDai.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SavingsDai *SavingsDaiCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SavingsDai.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SavingsDai *SavingsDaiSession) Version() (string, error) {
	return _SavingsDai.Contract.Version(&_SavingsDai.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SavingsDai *SavingsDaiCallerSession) Version() (string, error) {
	return _SavingsDai.Contract.Version(&_SavingsDai.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SavingsDai *SavingsDaiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SavingsDai *SavingsDaiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.Approve(&_SavingsDai.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SavingsDai *SavingsDaiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.Approve(&_SavingsDai.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_SavingsDai *SavingsDaiTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_SavingsDai *SavingsDaiSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.Burn(&_SavingsDai.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_SavingsDai *SavingsDaiTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.Burn(&_SavingsDai.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SavingsDai *SavingsDaiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SavingsDai.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SavingsDai *SavingsDaiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.DecreaseAllowance(&_SavingsDai.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SavingsDai *SavingsDaiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.DecreaseAllowance(&_SavingsDai.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SavingsDai *SavingsDaiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SavingsDai.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SavingsDai *SavingsDaiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.IncreaseAllowance(&_SavingsDai.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SavingsDai *SavingsDaiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.IncreaseAllowance(&_SavingsDai.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_SavingsDai *SavingsDaiTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_SavingsDai *SavingsDaiSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.Mint(&_SavingsDai.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_SavingsDai *SavingsDaiTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.Mint(&_SavingsDai.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SavingsDai *SavingsDaiTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SavingsDai *SavingsDaiSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.Transfer(&_SavingsDai.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SavingsDai *SavingsDaiTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.Transfer(&_SavingsDai.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SavingsDai *SavingsDaiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SavingsDai *SavingsDaiSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.TransferFrom(&_SavingsDai.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SavingsDai *SavingsDaiTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SavingsDai.Contract.TransferFrom(&_SavingsDai.TransactOpts, from, to, amount)
}

// SavingsDaiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SavingsDai contract.
type SavingsDaiApprovalIterator struct {
	Event *SavingsDaiApproval // Event containing the contract specifics and raw log

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
func (it *SavingsDaiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SavingsDaiApproval)
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
		it.Event = new(SavingsDaiApproval)
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
func (it *SavingsDaiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SavingsDaiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SavingsDaiApproval represents a Approval event raised by the SavingsDai contract.
type SavingsDaiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SavingsDai *SavingsDaiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SavingsDaiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SavingsDai.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SavingsDaiApprovalIterator{contract: _SavingsDai.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SavingsDai *SavingsDaiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SavingsDaiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SavingsDai.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SavingsDaiApproval)
				if err := _SavingsDai.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SavingsDai *SavingsDaiFilterer) ParseApproval(log types.Log) (*SavingsDaiApproval, error) {
	event := new(SavingsDaiApproval)
	if err := _SavingsDai.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SavingsDaiBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the SavingsDai contract.
type SavingsDaiBurnIterator struct {
	Event *SavingsDaiBurn // Event containing the contract specifics and raw log

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
func (it *SavingsDaiBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SavingsDaiBurn)
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
		it.Event = new(SavingsDaiBurn)
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
func (it *SavingsDaiBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SavingsDaiBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SavingsDaiBurn represents a Burn event raised by the SavingsDai contract.
type SavingsDaiBurn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_SavingsDai *SavingsDaiFilterer) FilterBurn(opts *bind.FilterOpts, account []common.Address) (*SavingsDaiBurnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SavingsDai.contract.FilterLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return &SavingsDaiBurnIterator{contract: _SavingsDai.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_SavingsDai *SavingsDaiFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *SavingsDaiBurn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SavingsDai.contract.WatchLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SavingsDaiBurn)
				if err := _SavingsDai.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_SavingsDai *SavingsDaiFilterer) ParseBurn(log types.Log) (*SavingsDaiBurn, error) {
	event := new(SavingsDaiBurn)
	if err := _SavingsDai.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SavingsDaiMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the SavingsDai contract.
type SavingsDaiMintIterator struct {
	Event *SavingsDaiMint // Event containing the contract specifics and raw log

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
func (it *SavingsDaiMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SavingsDaiMint)
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
		it.Event = new(SavingsDaiMint)
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
func (it *SavingsDaiMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SavingsDaiMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SavingsDaiMint represents a Mint event raised by the SavingsDai contract.
type SavingsDaiMint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_SavingsDai *SavingsDaiFilterer) FilterMint(opts *bind.FilterOpts, account []common.Address) (*SavingsDaiMintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SavingsDai.contract.FilterLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return &SavingsDaiMintIterator{contract: _SavingsDai.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_SavingsDai *SavingsDaiFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *SavingsDaiMint, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SavingsDai.contract.WatchLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SavingsDaiMint)
				if err := _SavingsDai.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_SavingsDai *SavingsDaiFilterer) ParseMint(log types.Log) (*SavingsDaiMint, error) {
	event := new(SavingsDaiMint)
	if err := _SavingsDai.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SavingsDaiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SavingsDai contract.
type SavingsDaiTransferIterator struct {
	Event *SavingsDaiTransfer // Event containing the contract specifics and raw log

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
func (it *SavingsDaiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SavingsDaiTransfer)
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
		it.Event = new(SavingsDaiTransfer)
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
func (it *SavingsDaiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SavingsDaiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SavingsDaiTransfer represents a Transfer event raised by the SavingsDai contract.
type SavingsDaiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SavingsDai *SavingsDaiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SavingsDaiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SavingsDai.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SavingsDaiTransferIterator{contract: _SavingsDai.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SavingsDai *SavingsDaiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SavingsDaiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SavingsDai.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SavingsDaiTransfer)
				if err := _SavingsDai.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SavingsDai *SavingsDaiFilterer) ParseTransfer(log types.Log) (*SavingsDaiTransfer, error) {
	event := new(SavingsDaiTransfer)
	if err := _SavingsDai.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
