// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimismMixedRouteQuoterV1

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

// OptimismMixedRouteQuoterV1MetaData contains all meta data concerning the OptimismMixedRouteQuoterV1 contract.
var OptimismMixedRouteQuoterV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"v3SqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"v3InitializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"v3SwapGasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"internalType\":\"structIMixedRouteQuoterV1.QuoteExactInputSingleV2Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingleV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIMixedRouteQuoterV1.QuoteExactInputSingleV3Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingleV3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OptimismMixedRouteQuoterV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use OptimismMixedRouteQuoterV1MetaData.ABI instead.
var OptimismMixedRouteQuoterV1ABI = OptimismMixedRouteQuoterV1MetaData.ABI

// OptimismMixedRouteQuoterV1 is an auto generated Go binding around an Ethereum contract.
type OptimismMixedRouteQuoterV1 struct {
	OptimismMixedRouteQuoterV1Caller     // Read-only binding to the contract
	OptimismMixedRouteQuoterV1Transactor // Write-only binding to the contract
	OptimismMixedRouteQuoterV1Filterer   // Log filterer for contract events
}

// OptimismMixedRouteQuoterV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type OptimismMixedRouteQuoterV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismMixedRouteQuoterV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OptimismMixedRouteQuoterV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismMixedRouteQuoterV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptimismMixedRouteQuoterV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismMixedRouteQuoterV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptimismMixedRouteQuoterV1Session struct {
	Contract     *OptimismMixedRouteQuoterV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OptimismMixedRouteQuoterV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptimismMixedRouteQuoterV1CallerSession struct {
	Contract *OptimismMixedRouteQuoterV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// OptimismMixedRouteQuoterV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptimismMixedRouteQuoterV1TransactorSession struct {
	Contract     *OptimismMixedRouteQuoterV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// OptimismMixedRouteQuoterV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type OptimismMixedRouteQuoterV1Raw struct {
	Contract *OptimismMixedRouteQuoterV1 // Generic contract binding to access the raw methods on
}

// OptimismMixedRouteQuoterV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptimismMixedRouteQuoterV1CallerRaw struct {
	Contract *OptimismMixedRouteQuoterV1Caller // Generic read-only contract binding to access the raw methods on
}

// OptimismMixedRouteQuoterV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptimismMixedRouteQuoterV1TransactorRaw struct {
	Contract *OptimismMixedRouteQuoterV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOptimismMixedRouteQuoterV1 creates a new instance of OptimismMixedRouteQuoterV1, bound to a specific deployed contract.
func NewOptimismMixedRouteQuoterV1(address common.Address, backend bind.ContractBackend) (*OptimismMixedRouteQuoterV1, error) {
	contract, err := bindOptimismMixedRouteQuoterV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismMixedRouteQuoterV1{OptimismMixedRouteQuoterV1Caller: OptimismMixedRouteQuoterV1Caller{contract: contract}, OptimismMixedRouteQuoterV1Transactor: OptimismMixedRouteQuoterV1Transactor{contract: contract}, OptimismMixedRouteQuoterV1Filterer: OptimismMixedRouteQuoterV1Filterer{contract: contract}}, nil
}

// NewOptimismMixedRouteQuoterV1Caller creates a new read-only instance of OptimismMixedRouteQuoterV1, bound to a specific deployed contract.
func NewOptimismMixedRouteQuoterV1Caller(address common.Address, caller bind.ContractCaller) (*OptimismMixedRouteQuoterV1Caller, error) {
	contract, err := bindOptimismMixedRouteQuoterV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismMixedRouteQuoterV1Caller{contract: contract}, nil
}

// NewOptimismMixedRouteQuoterV1Transactor creates a new write-only instance of OptimismMixedRouteQuoterV1, bound to a specific deployed contract.
func NewOptimismMixedRouteQuoterV1Transactor(address common.Address, transactor bind.ContractTransactor) (*OptimismMixedRouteQuoterV1Transactor, error) {
	contract, err := bindOptimismMixedRouteQuoterV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismMixedRouteQuoterV1Transactor{contract: contract}, nil
}

// NewOptimismMixedRouteQuoterV1Filterer creates a new log filterer instance of OptimismMixedRouteQuoterV1, bound to a specific deployed contract.
func NewOptimismMixedRouteQuoterV1Filterer(address common.Address, filterer bind.ContractFilterer) (*OptimismMixedRouteQuoterV1Filterer, error) {
	contract, err := bindOptimismMixedRouteQuoterV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismMixedRouteQuoterV1Filterer{contract: contract}, nil
}

// bindOptimismMixedRouteQuoterV1 binds a generic wrapper to an already deployed contract.
func bindOptimismMixedRouteQuoterV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismMixedRouteQuoterV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismMixedRouteQuoterV1.Contract.OptimismMixedRouteQuoterV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismMixedRouteQuoterV1.Contract.OptimismMixedRouteQuoterV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismMixedRouteQuoterV1.Contract.OptimismMixedRouteQuoterV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismMixedRouteQuoterV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismMixedRouteQuoterV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismMixedRouteQuoterV1.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Caller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismMixedRouteQuoterV1.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Session) WETH9() (common.Address, error) {
	return _OptimismMixedRouteQuoterV1.Contract.WETH9(&_OptimismMixedRouteQuoterV1.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1CallerSession) WETH9() (common.Address, error) {
	return _OptimismMixedRouteQuoterV1.Contract.WETH9(&_OptimismMixedRouteQuoterV1.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismMixedRouteQuoterV1.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Session) Factory() (common.Address, error) {
	return _OptimismMixedRouteQuoterV1.Contract.Factory(&_OptimismMixedRouteQuoterV1.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1CallerSession) Factory() (common.Address, error) {
	return _OptimismMixedRouteQuoterV1.Contract.Factory(&_OptimismMixedRouteQuoterV1.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Caller) FactoryV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismMixedRouteQuoterV1.contract.Call(opts, &out, "factoryV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Session) FactoryV2() (common.Address, error) {
	return _OptimismMixedRouteQuoterV1.Contract.FactoryV2(&_OptimismMixedRouteQuoterV1.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1CallerSession) FactoryV2() (common.Address, error) {
	return _OptimismMixedRouteQuoterV1.Contract.FactoryV2(&_OptimismMixedRouteQuoterV1.CallOpts)
}

// QuoteExactInputSingleV2 is a free data retrieval call binding the contract method 0xc550b186.
//
// Solidity: function quoteExactInputSingleV2((address,address,bool,uint256) params) view returns(uint256 amountOut)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Caller) QuoteExactInputSingleV2(opts *bind.CallOpts, params IMixedRouteQuoterV1QuoteExactInputSingleV2Params) (*big.Int, error) {
	var out []interface{}
	err := _OptimismMixedRouteQuoterV1.contract.Call(opts, &out, "quoteExactInputSingleV2", params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteExactInputSingleV2 is a free data retrieval call binding the contract method 0xc550b186.
//
// Solidity: function quoteExactInputSingleV2((address,address,bool,uint256) params) view returns(uint256 amountOut)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Session) QuoteExactInputSingleV2(params IMixedRouteQuoterV1QuoteExactInputSingleV2Params) (*big.Int, error) {
	return _OptimismMixedRouteQuoterV1.Contract.QuoteExactInputSingleV2(&_OptimismMixedRouteQuoterV1.CallOpts, params)
}

// QuoteExactInputSingleV2 is a free data retrieval call binding the contract method 0xc550b186.
//
// Solidity: function quoteExactInputSingleV2((address,address,bool,uint256) params) view returns(uint256 amountOut)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1CallerSession) QuoteExactInputSingleV2(params IMixedRouteQuoterV1QuoteExactInputSingleV2Params) (*big.Int, error) {
	return _OptimismMixedRouteQuoterV1.Contract.QuoteExactInputSingleV2(&_OptimismMixedRouteQuoterV1.CallOpts, params)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Caller) UniswapV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _OptimismMixedRouteQuoterV1.contract.Call(opts, &out, "uniswapV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _OptimismMixedRouteQuoterV1.Contract.UniswapV3SwapCallback(&_OptimismMixedRouteQuoterV1.CallOpts, amount0Delta, amount1Delta, path)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1CallerSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _OptimismMixedRouteQuoterV1.Contract.UniswapV3SwapCallback(&_OptimismMixedRouteQuoterV1.CallOpts, amount0Delta, amount1Delta, path)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] v3SqrtPriceX96AfterList, uint32[] v3InitializedTicksCrossedList, uint256 v3SwapGasEstimate)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Transactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _OptimismMixedRouteQuoterV1.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] v3SqrtPriceX96AfterList, uint32[] v3InitializedTicksCrossedList, uint256 v3SwapGasEstimate)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Session) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _OptimismMixedRouteQuoterV1.Contract.QuoteExactInput(&_OptimismMixedRouteQuoterV1.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] v3SqrtPriceX96AfterList, uint32[] v3InitializedTicksCrossedList, uint256 v3SwapGasEstimate)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1TransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _OptimismMixedRouteQuoterV1.Contract.QuoteExactInput(&_OptimismMixedRouteQuoterV1.TransactOpts, path, amountIn)
}

// QuoteExactInputSingleV3 is a paid mutator transaction binding the contract method 0x891e50c6.
//
// Solidity: function quoteExactInputSingleV3((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Transactor) QuoteExactInputSingleV3(opts *bind.TransactOpts, params IMixedRouteQuoterV1QuoteExactInputSingleV3Params) (*types.Transaction, error) {
	return _OptimismMixedRouteQuoterV1.contract.Transact(opts, "quoteExactInputSingleV3", params)
}

// QuoteExactInputSingleV3 is a paid mutator transaction binding the contract method 0x891e50c6.
//
// Solidity: function quoteExactInputSingleV3((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1Session) QuoteExactInputSingleV3(params IMixedRouteQuoterV1QuoteExactInputSingleV3Params) (*types.Transaction, error) {
	return _OptimismMixedRouteQuoterV1.Contract.QuoteExactInputSingleV3(&_OptimismMixedRouteQuoterV1.TransactOpts, params)
}

// QuoteExactInputSingleV3 is a paid mutator transaction binding the contract method 0x891e50c6.
//
// Solidity: function quoteExactInputSingleV3((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_OptimismMixedRouteQuoterV1 *OptimismMixedRouteQuoterV1TransactorSession) QuoteExactInputSingleV3(params IMixedRouteQuoterV1QuoteExactInputSingleV3Params) (*types.Transaction, error) {
	return _OptimismMixedRouteQuoterV1.Contract.QuoteExactInputSingleV3(&_OptimismMixedRouteQuoterV1.TransactOpts, params)
}
