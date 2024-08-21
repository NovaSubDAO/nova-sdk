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

// IVelodromeMetaData contains all meta data concerning the IVelodrome contract.
var IVelodromeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"slot0\",\"inputs\":[],\"outputs\":[{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"observationIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"observationCardinality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"observationCardinalityNext\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"unlocked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"amount1\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// IVelodromeABI is the input ABI used to generate the binding from.
// Deprecated: Use IVelodromeMetaData.ABI instead.
var IVelodromeABI = IVelodromeMetaData.ABI

// IVelodrome is an auto generated Go binding around an Ethereum contract.
type IVelodrome struct {
	IVelodromeCaller     // Read-only binding to the contract
	IVelodromeTransactor // Write-only binding to the contract
	IVelodromeFilterer   // Log filterer for contract events
}

// IVelodromeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVelodromeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVelodromeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVelodromeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVelodromeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVelodromeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVelodromeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVelodromeSession struct {
	Contract     *IVelodrome // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IVelodromeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVelodromeCallerSession struct {
	Contract *IVelodromeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IVelodromeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVelodromeTransactorSession struct {
	Contract     *IVelodromeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IVelodromeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVelodromeRaw struct {
	Contract *IVelodrome // Generic contract binding to access the raw methods on
}

// IVelodromeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVelodromeCallerRaw struct {
	Contract *IVelodromeCaller // Generic read-only contract binding to access the raw methods on
}

// IVelodromeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVelodromeTransactorRaw struct {
	Contract *IVelodromeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVelodrome creates a new instance of IVelodrome, bound to a specific deployed contract.
func NewIVelodrome(address common.Address, backend bind.ContractBackend) (*IVelodrome, error) {
	contract, err := bindIVelodrome(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVelodrome{IVelodromeCaller: IVelodromeCaller{contract: contract}, IVelodromeTransactor: IVelodromeTransactor{contract: contract}, IVelodromeFilterer: IVelodromeFilterer{contract: contract}}, nil
}

// NewIVelodromeCaller creates a new read-only instance of IVelodrome, bound to a specific deployed contract.
func NewIVelodromeCaller(address common.Address, caller bind.ContractCaller) (*IVelodromeCaller, error) {
	contract, err := bindIVelodrome(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVelodromeCaller{contract: contract}, nil
}

// NewIVelodromeTransactor creates a new write-only instance of IVelodrome, bound to a specific deployed contract.
func NewIVelodromeTransactor(address common.Address, transactor bind.ContractTransactor) (*IVelodromeTransactor, error) {
	contract, err := bindIVelodrome(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVelodromeTransactor{contract: contract}, nil
}

// NewIVelodromeFilterer creates a new log filterer instance of IVelodrome, bound to a specific deployed contract.
func NewIVelodromeFilterer(address common.Address, filterer bind.ContractFilterer) (*IVelodromeFilterer, error) {
	contract, err := bindIVelodrome(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVelodromeFilterer{contract: contract}, nil
}

// bindIVelodrome binds a generic wrapper to an already deployed contract.
func bindIVelodrome(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVelodromeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVelodrome *IVelodromeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVelodrome.Contract.IVelodromeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVelodrome *IVelodromeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVelodrome.Contract.IVelodromeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVelodrome *IVelodromeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVelodrome.Contract.IVelodromeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVelodrome *IVelodromeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVelodrome.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVelodrome *IVelodromeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVelodrome.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVelodrome *IVelodromeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVelodrome.Contract.contract.Transact(opts, method, params...)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, bool unlocked)
func (_IVelodrome *IVelodromeCaller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "slot0")

	outstruct := new(struct {
		SqrtPriceX96               *big.Int
		Tick                       *big.Int
		ObservationIndex           uint16
		ObservationCardinality     uint16
		ObservationCardinalityNext uint16
		Unlocked                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationCardinality = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationCardinalityNext = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Unlocked = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, bool unlocked)
func (_IVelodrome *IVelodromeSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	Unlocked                   bool
}, error) {
	return _IVelodrome.Contract.Slot0(&_IVelodrome.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, bool unlocked)
func (_IVelodrome *IVelodromeCallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	Unlocked                   bool
}, error) {
	return _IVelodrome.Contract.Slot0(&_IVelodrome.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IVelodrome *IVelodromeCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IVelodrome *IVelodromeSession) Token0() (common.Address, error) {
	return _IVelodrome.Contract.Token0(&_IVelodrome.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_IVelodrome *IVelodromeCallerSession) Token0() (common.Address, error) {
	return _IVelodrome.Contract.Token0(&_IVelodrome.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IVelodrome *IVelodromeCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVelodrome.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IVelodrome *IVelodromeSession) Token1() (common.Address, error) {
	return _IVelodrome.Contract.Token1(&_IVelodrome.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_IVelodrome *IVelodromeCallerSession) Token1() (common.Address, error) {
	return _IVelodrome.Contract.Token1(&_IVelodrome.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_IVelodrome *IVelodromeTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _IVelodrome.contract.Transact(opts, "swap", recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_IVelodrome *IVelodromeSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _IVelodrome.Contract.Swap(&_IVelodrome.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_IVelodrome *IVelodromeTransactorSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _IVelodrome.Contract.Swap(&_IVelodrome.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}
