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

// VelodromeMetaData contains all meta data concerning the Velodrome contract.
var VelodromeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"slot0\",\"inputs\":[],\"outputs\":[{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"observationIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"observationCardinality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"observationCardinalityNext\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"unlocked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"amount1\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// VelodromeABI is the input ABI used to generate the binding from.
// Deprecated: Use VelodromeMetaData.ABI instead.
var VelodromeABI = VelodromeMetaData.ABI

// Velodrome is an auto generated Go binding around an Ethereum contract.
type Velodrome struct {
	VelodromeCaller     // Read-only binding to the contract
	VelodromeTransactor // Write-only binding to the contract
	VelodromeFilterer   // Log filterer for contract events
}

// VelodromeCaller is an auto generated read-only Go binding around an Ethereum contract.
type VelodromeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VelodromeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VelodromeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VelodromeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VelodromeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VelodromeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VelodromeSession struct {
	Contract     *Velodrome        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VelodromeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VelodromeCallerSession struct {
	Contract *VelodromeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VelodromeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VelodromeTransactorSession struct {
	Contract     *VelodromeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VelodromeRaw is an auto generated low-level Go binding around an Ethereum contract.
type VelodromeRaw struct {
	Contract *Velodrome // Generic contract binding to access the raw methods on
}

// VelodromeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VelodromeCallerRaw struct {
	Contract *VelodromeCaller // Generic read-only contract binding to access the raw methods on
}

// VelodromeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VelodromeTransactorRaw struct {
	Contract *VelodromeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVelodrome creates a new instance of Velodrome, bound to a specific deployed contract.
func NewVelodrome(address common.Address, backend bind.ContractBackend) (*Velodrome, error) {
	contract, err := bindVelodrome(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Velodrome{VelodromeCaller: VelodromeCaller{contract: contract}, VelodromeTransactor: VelodromeTransactor{contract: contract}, VelodromeFilterer: VelodromeFilterer{contract: contract}}, nil
}

// NewVelodromeCaller creates a new read-only instance of Velodrome, bound to a specific deployed contract.
func NewVelodromeCaller(address common.Address, caller bind.ContractCaller) (*VelodromeCaller, error) {
	contract, err := bindVelodrome(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VelodromeCaller{contract: contract}, nil
}

// NewVelodromeTransactor creates a new write-only instance of Velodrome, bound to a specific deployed contract.
func NewVelodromeTransactor(address common.Address, transactor bind.ContractTransactor) (*VelodromeTransactor, error) {
	contract, err := bindVelodrome(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VelodromeTransactor{contract: contract}, nil
}

// NewVelodromeFilterer creates a new log filterer instance of Velodrome, bound to a specific deployed contract.
func NewVelodromeFilterer(address common.Address, filterer bind.ContractFilterer) (*VelodromeFilterer, error) {
	contract, err := bindVelodrome(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VelodromeFilterer{contract: contract}, nil
}

// bindVelodrome binds a generic wrapper to an already deployed contract.
func bindVelodrome(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VelodromeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Velodrome *VelodromeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Velodrome.Contract.VelodromeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Velodrome *VelodromeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Velodrome.Contract.VelodromeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Velodrome *VelodromeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Velodrome.Contract.VelodromeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Velodrome *VelodromeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Velodrome.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Velodrome *VelodromeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Velodrome.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Velodrome *VelodromeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Velodrome.Contract.contract.Transact(opts, method, params...)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, bool unlocked)
func (_Velodrome *VelodromeCaller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _Velodrome.contract.Call(opts, &out, "slot0")

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
func (_Velodrome *VelodromeSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	Unlocked                   bool
}, error) {
	return _Velodrome.Contract.Slot0(&_Velodrome.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, bool unlocked)
func (_Velodrome *VelodromeCallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	Unlocked                   bool
}, error) {
	return _Velodrome.Contract.Slot0(&_Velodrome.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Velodrome *VelodromeCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Velodrome.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Velodrome *VelodromeSession) Token0() (common.Address, error) {
	return _Velodrome.Contract.Token0(&_Velodrome.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Velodrome *VelodromeCallerSession) Token0() (common.Address, error) {
	return _Velodrome.Contract.Token0(&_Velodrome.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Velodrome *VelodromeCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Velodrome.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Velodrome *VelodromeSession) Token1() (common.Address, error) {
	return _Velodrome.Contract.Token1(&_Velodrome.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Velodrome *VelodromeCallerSession) Token1() (common.Address, error) {
	return _Velodrome.Contract.Token1(&_Velodrome.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_Velodrome *VelodromeTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _Velodrome.contract.Transact(opts, "swap", recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_Velodrome *VelodromeSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _Velodrome.Contract.Swap(&_Velodrome.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_Velodrome *VelodromeTransactorSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _Velodrome.Contract.Swap(&_Velodrome.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}
