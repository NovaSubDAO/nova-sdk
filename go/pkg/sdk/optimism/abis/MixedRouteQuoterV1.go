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

// IMixedRouteQuoterV1QuoteExactInputSingleV2Params is an auto generated low-level Go binding around an user-defined struct.
type IMixedRouteQuoterV1QuoteExactInputSingleV2Params struct {
	TokenIn  common.Address
	TokenOut common.Address
	Stable   bool
	AmountIn *big.Int
}

// IMixedRouteQuoterV1QuoteExactInputSingleV3Params is an auto generated low-level Go binding around an user-defined struct.
type IMixedRouteQuoterV1QuoteExactInputSingleV3Params struct {
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	TickSpacing       *big.Int
	SqrtPriceLimitX96 *big.Int
}

// MixedRouteQuoterV1MetaData contains all meta data concerning the MixedRouteQuoterV1 contract.
var MixedRouteQuoterV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"v3SqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"v3InitializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"v3SwapGasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"internalType\":\"structIMixedRouteQuoterV1.QuoteExactInputSingleV2Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingleV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIMixedRouteQuoterV1.QuoteExactInputSingleV3Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingleV3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MixedRouteQuoterV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use MixedRouteQuoterV1MetaData.ABI instead.
var MixedRouteQuoterV1ABI = MixedRouteQuoterV1MetaData.ABI

// MixedRouteQuoterV1 is an auto generated Go binding around an Ethereum contract.
type MixedRouteQuoterV1 struct {
	MixedRouteQuoterV1Caller     // Read-only binding to the contract
	MixedRouteQuoterV1Transactor // Write-only binding to the contract
	MixedRouteQuoterV1Filterer   // Log filterer for contract events
}

// MixedRouteQuoterV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type MixedRouteQuoterV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixedRouteQuoterV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MixedRouteQuoterV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixedRouteQuoterV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MixedRouteQuoterV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixedRouteQuoterV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MixedRouteQuoterV1Session struct {
	Contract     *MixedRouteQuoterV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MixedRouteQuoterV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MixedRouteQuoterV1CallerSession struct {
	Contract *MixedRouteQuoterV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MixedRouteQuoterV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MixedRouteQuoterV1TransactorSession struct {
	Contract     *MixedRouteQuoterV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MixedRouteQuoterV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type MixedRouteQuoterV1Raw struct {
	Contract *MixedRouteQuoterV1 // Generic contract binding to access the raw methods on
}

// MixedRouteQuoterV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MixedRouteQuoterV1CallerRaw struct {
	Contract *MixedRouteQuoterV1Caller // Generic read-only contract binding to access the raw methods on
}

// MixedRouteQuoterV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MixedRouteQuoterV1TransactorRaw struct {
	Contract *MixedRouteQuoterV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMixedRouteQuoterV1 creates a new instance of MixedRouteQuoterV1, bound to a specific deployed contract.
func NewMixedRouteQuoterV1(address common.Address, backend bind.ContractBackend) (*MixedRouteQuoterV1, error) {
	contract, err := bindMixedRouteQuoterV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MixedRouteQuoterV1{MixedRouteQuoterV1Caller: MixedRouteQuoterV1Caller{contract: contract}, MixedRouteQuoterV1Transactor: MixedRouteQuoterV1Transactor{contract: contract}, MixedRouteQuoterV1Filterer: MixedRouteQuoterV1Filterer{contract: contract}}, nil
}

// NewMixedRouteQuoterV1Caller creates a new read-only instance of MixedRouteQuoterV1, bound to a specific deployed contract.
func NewMixedRouteQuoterV1Caller(address common.Address, caller bind.ContractCaller) (*MixedRouteQuoterV1Caller, error) {
	contract, err := bindMixedRouteQuoterV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MixedRouteQuoterV1Caller{contract: contract}, nil
}

// NewMixedRouteQuoterV1Transactor creates a new write-only instance of MixedRouteQuoterV1, bound to a specific deployed contract.
func NewMixedRouteQuoterV1Transactor(address common.Address, transactor bind.ContractTransactor) (*MixedRouteQuoterV1Transactor, error) {
	contract, err := bindMixedRouteQuoterV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MixedRouteQuoterV1Transactor{contract: contract}, nil
}

// NewMixedRouteQuoterV1Filterer creates a new log filterer instance of MixedRouteQuoterV1, bound to a specific deployed contract.
func NewMixedRouteQuoterV1Filterer(address common.Address, filterer bind.ContractFilterer) (*MixedRouteQuoterV1Filterer, error) {
	contract, err := bindMixedRouteQuoterV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MixedRouteQuoterV1Filterer{contract: contract}, nil
}

// bindMixedRouteQuoterV1 binds a generic wrapper to an already deployed contract.
func bindMixedRouteQuoterV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MixedRouteQuoterV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MixedRouteQuoterV1.Contract.MixedRouteQuoterV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MixedRouteQuoterV1.Contract.MixedRouteQuoterV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MixedRouteQuoterV1.Contract.MixedRouteQuoterV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MixedRouteQuoterV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MixedRouteQuoterV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MixedRouteQuoterV1.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Caller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MixedRouteQuoterV1.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Session) WETH9() (common.Address, error) {
	return _MixedRouteQuoterV1.Contract.WETH9(&_MixedRouteQuoterV1.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1CallerSession) WETH9() (common.Address, error) {
	return _MixedRouteQuoterV1.Contract.WETH9(&_MixedRouteQuoterV1.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MixedRouteQuoterV1.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Session) Factory() (common.Address, error) {
	return _MixedRouteQuoterV1.Contract.Factory(&_MixedRouteQuoterV1.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1CallerSession) Factory() (common.Address, error) {
	return _MixedRouteQuoterV1.Contract.Factory(&_MixedRouteQuoterV1.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Caller) FactoryV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MixedRouteQuoterV1.contract.Call(opts, &out, "factoryV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Session) FactoryV2() (common.Address, error) {
	return _MixedRouteQuoterV1.Contract.FactoryV2(&_MixedRouteQuoterV1.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1CallerSession) FactoryV2() (common.Address, error) {
	return _MixedRouteQuoterV1.Contract.FactoryV2(&_MixedRouteQuoterV1.CallOpts)
}

// QuoteExactInputSingleV2 is a free data retrieval call binding the contract method 0xc550b186.
//
// Solidity: function quoteExactInputSingleV2((address,address,bool,uint256) params) view returns(uint256 amountOut)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Caller) QuoteExactInputSingleV2(opts *bind.CallOpts, params IMixedRouteQuoterV1QuoteExactInputSingleV2Params) (*big.Int, error) {
	var out []interface{}
	err := _MixedRouteQuoterV1.contract.Call(opts, &out, "quoteExactInputSingleV2", params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteExactInputSingleV2 is a free data retrieval call binding the contract method 0xc550b186.
//
// Solidity: function quoteExactInputSingleV2((address,address,bool,uint256) params) view returns(uint256 amountOut)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Session) QuoteExactInputSingleV2(params IMixedRouteQuoterV1QuoteExactInputSingleV2Params) (*big.Int, error) {
	return _MixedRouteQuoterV1.Contract.QuoteExactInputSingleV2(&_MixedRouteQuoterV1.CallOpts, params)
}

// QuoteExactInputSingleV2 is a free data retrieval call binding the contract method 0xc550b186.
//
// Solidity: function quoteExactInputSingleV2((address,address,bool,uint256) params) view returns(uint256 amountOut)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1CallerSession) QuoteExactInputSingleV2(params IMixedRouteQuoterV1QuoteExactInputSingleV2Params) (*big.Int, error) {
	return _MixedRouteQuoterV1.Contract.QuoteExactInputSingleV2(&_MixedRouteQuoterV1.CallOpts, params)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Caller) UniswapV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _MixedRouteQuoterV1.contract.Call(opts, &out, "uniswapV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _MixedRouteQuoterV1.Contract.UniswapV3SwapCallback(&_MixedRouteQuoterV1.CallOpts, amount0Delta, amount1Delta, path)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1CallerSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _MixedRouteQuoterV1.Contract.UniswapV3SwapCallback(&_MixedRouteQuoterV1.CallOpts, amount0Delta, amount1Delta, path)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] v3SqrtPriceX96AfterList, uint32[] v3InitializedTicksCrossedList, uint256 v3SwapGasEstimate)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Transactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _MixedRouteQuoterV1.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] v3SqrtPriceX96AfterList, uint32[] v3InitializedTicksCrossedList, uint256 v3SwapGasEstimate)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Session) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _MixedRouteQuoterV1.Contract.QuoteExactInput(&_MixedRouteQuoterV1.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] v3SqrtPriceX96AfterList, uint32[] v3InitializedTicksCrossedList, uint256 v3SwapGasEstimate)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1TransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _MixedRouteQuoterV1.Contract.QuoteExactInput(&_MixedRouteQuoterV1.TransactOpts, path, amountIn)
}

// QuoteExactInputSingleV3 is a paid mutator transaction binding the contract method 0x891e50c6.
//
// Solidity: function quoteExactInputSingleV3((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Transactor) QuoteExactInputSingleV3(opts *bind.TransactOpts, params IMixedRouteQuoterV1QuoteExactInputSingleV3Params) (*types.Transaction, error) {
	return _MixedRouteQuoterV1.contract.Transact(opts, "quoteExactInputSingleV3", params)
}

// QuoteExactInputSingleV3 is a paid mutator transaction binding the contract method 0x891e50c6.
//
// Solidity: function quoteExactInputSingleV3((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1Session) QuoteExactInputSingleV3(params IMixedRouteQuoterV1QuoteExactInputSingleV3Params) (*types.Transaction, error) {
	return _MixedRouteQuoterV1.Contract.QuoteExactInputSingleV3(&_MixedRouteQuoterV1.TransactOpts, params)
}

// QuoteExactInputSingleV3 is a paid mutator transaction binding the contract method 0x891e50c6.
//
// Solidity: function quoteExactInputSingleV3((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_MixedRouteQuoterV1 *MixedRouteQuoterV1TransactorSession) QuoteExactInputSingleV3(params IMixedRouteQuoterV1QuoteExactInputSingleV3Params) (*types.Transaction, error) {
	return _MixedRouteQuoterV1.Contract.QuoteExactInputSingleV3(&_MixedRouteQuoterV1.TransactOpts, params)
}
