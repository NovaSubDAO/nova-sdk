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

// VelodromePoolMetaData contains all meta data concerning the VelodromePool contract.
var VelodromePoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"slot0\",\"inputs\":[],\"outputs\":[{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"observationIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"observationCardinality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"observationCardinalityNext\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"unlocked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"amount1\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// VelodromePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use VelodromePoolMetaData.ABI instead.
var VelodromePoolABI = VelodromePoolMetaData.ABI

// VelodromePool is an auto generated Go binding around an Ethereum contract.
type VelodromePool struct {
	VelodromePoolCaller     // Read-only binding to the contract
	VelodromePoolTransactor // Write-only binding to the contract
	VelodromePoolFilterer   // Log filterer for contract events
}

// VelodromePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type VelodromePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VelodromePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VelodromePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VelodromePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VelodromePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VelodromePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VelodromePoolSession struct {
	Contract     *VelodromePool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VelodromePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VelodromePoolCallerSession struct {
	Contract *VelodromePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VelodromePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VelodromePoolTransactorSession struct {
	Contract     *VelodromePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VelodromePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type VelodromePoolRaw struct {
	Contract *VelodromePool // Generic contract binding to access the raw methods on
}

// VelodromePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VelodromePoolCallerRaw struct {
	Contract *VelodromePoolCaller // Generic read-only contract binding to access the raw methods on
}

// VelodromePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VelodromePoolTransactorRaw struct {
	Contract *VelodromePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVelodromePool creates a new instance of VelodromePool, bound to a specific deployed contract.
func NewVelodromePool(address common.Address, backend bind.ContractBackend) (*VelodromePool, error) {
	contract, err := bindVelodromePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VelodromePool{VelodromePoolCaller: VelodromePoolCaller{contract: contract}, VelodromePoolTransactor: VelodromePoolTransactor{contract: contract}, VelodromePoolFilterer: VelodromePoolFilterer{contract: contract}}, nil
}

// NewVelodromePoolCaller creates a new read-only instance of VelodromePool, bound to a specific deployed contract.
func NewVelodromePoolCaller(address common.Address, caller bind.ContractCaller) (*VelodromePoolCaller, error) {
	contract, err := bindVelodromePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VelodromePoolCaller{contract: contract}, nil
}

// NewVelodromePoolTransactor creates a new write-only instance of VelodromePool, bound to a specific deployed contract.
func NewVelodromePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*VelodromePoolTransactor, error) {
	contract, err := bindVelodromePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VelodromePoolTransactor{contract: contract}, nil
}

// NewVelodromePoolFilterer creates a new log filterer instance of VelodromePool, bound to a specific deployed contract.
func NewVelodromePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*VelodromePoolFilterer, error) {
	contract, err := bindVelodromePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VelodromePoolFilterer{contract: contract}, nil
}

// bindVelodromePool binds a generic wrapper to an already deployed contract.
func bindVelodromePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VelodromePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VelodromePool *VelodromePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VelodromePool.Contract.VelodromePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VelodromePool *VelodromePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VelodromePool.Contract.VelodromePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VelodromePool *VelodromePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VelodromePool.Contract.VelodromePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VelodromePool *VelodromePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VelodromePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VelodromePool *VelodromePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VelodromePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VelodromePool *VelodromePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VelodromePool.Contract.contract.Transact(opts, method, params...)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, bool unlocked)
func (_VelodromePool *VelodromePoolCaller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _VelodromePool.contract.Call(opts, &out, "slot0")

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
func (_VelodromePool *VelodromePoolSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	Unlocked                   bool
}, error) {
	return _VelodromePool.Contract.Slot0(&_VelodromePool.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, bool unlocked)
func (_VelodromePool *VelodromePoolCallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	Unlocked                   bool
}, error) {
	return _VelodromePool.Contract.Slot0(&_VelodromePool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_VelodromePool *VelodromePoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VelodromePool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_VelodromePool *VelodromePoolSession) Token0() (common.Address, error) {
	return _VelodromePool.Contract.Token0(&_VelodromePool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_VelodromePool *VelodromePoolCallerSession) Token0() (common.Address, error) {
	return _VelodromePool.Contract.Token0(&_VelodromePool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_VelodromePool *VelodromePoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VelodromePool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_VelodromePool *VelodromePoolSession) Token1() (common.Address, error) {
	return _VelodromePool.Contract.Token1(&_VelodromePool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_VelodromePool *VelodromePoolCallerSession) Token1() (common.Address, error) {
	return _VelodromePool.Contract.Token1(&_VelodromePool.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_VelodromePool *VelodromePoolTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _VelodromePool.contract.Transact(opts, "swap", recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_VelodromePool *VelodromePoolSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _VelodromePool.Contract.Swap(&_VelodromePool.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_VelodromePool *VelodromePoolTransactorSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _VelodromePool.Contract.Swap(&_VelodromePool.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}
