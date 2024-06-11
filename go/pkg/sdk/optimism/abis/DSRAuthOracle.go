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

// IDSROraclePotData is an auto generated low-level Go binding around an user-defined struct.
type IDSROraclePotData struct {
	Dsr *big.Int
	Chi *big.Int
	Rho *big.Int
}

// DSRAuthOracleMetaData contains all meta data concerning the DSRAuthOracle contract.
var DSRAuthOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxDSR\",\"type\":\"uint256\"}],\"name\":\"SetMaxDSR\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"dsr\",\"type\":\"uint96\"},{\"internalType\":\"uint120\",\"name\":\"chi\",\"type\":\"uint120\"},{\"internalType\":\"uint40\",\"name\":\"rho\",\"type\":\"uint40\"}],\"indexed\":false,\"internalType\":\"structIDSROracle.PotData\",\"name\":\"nextData\",\"type\":\"tuple\"}],\"name\":\"SetPotData\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DATA_PROVIDER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getConversionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConversionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getConversionRateBinomialApprox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConversionRateBinomialApprox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getConversionRateLinearApprox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConversionRateLinearApprox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDSR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPotData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"dsr\",\"type\":\"uint96\"},{\"internalType\":\"uint120\",\"name\":\"chi\",\"type\":\"uint120\"},{\"internalType\":\"uint40\",\"name\":\"rho\",\"type\":\"uint40\"}],\"internalType\":\"structIDSROracle.PotData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRho\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDSR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDSR\",\"type\":\"uint256\"}],\"name\":\"setMaxDSR\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"dsr\",\"type\":\"uint96\"},{\"internalType\":\"uint120\",\"name\":\"chi\",\"type\":\"uint120\"},{\"internalType\":\"uint40\",\"name\":\"rho\",\"type\":\"uint40\"}],\"internalType\":\"structIDSROracle.PotData\",\"name\":\"nextData\",\"type\":\"tuple\"}],\"name\":\"setPotData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DSRAuthOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use DSRAuthOracleMetaData.ABI instead.
var DSRAuthOracleABI = DSRAuthOracleMetaData.ABI

// DSRAuthOracle is an auto generated Go binding around an Ethereum contract.
type DSRAuthOracle struct {
	DSRAuthOracleCaller     // Read-only binding to the contract
	DSRAuthOracleTransactor // Write-only binding to the contract
	DSRAuthOracleFilterer   // Log filterer for contract events
}

// DSRAuthOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSRAuthOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSRAuthOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSRAuthOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSRAuthOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSRAuthOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSRAuthOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSRAuthOracleSession struct {
	Contract     *DSRAuthOracle    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSRAuthOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSRAuthOracleCallerSession struct {
	Contract *DSRAuthOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DSRAuthOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSRAuthOracleTransactorSession struct {
	Contract     *DSRAuthOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DSRAuthOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSRAuthOracleRaw struct {
	Contract *DSRAuthOracle // Generic contract binding to access the raw methods on
}

// DSRAuthOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSRAuthOracleCallerRaw struct {
	Contract *DSRAuthOracleCaller // Generic read-only contract binding to access the raw methods on
}

// DSRAuthOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSRAuthOracleTransactorRaw struct {
	Contract *DSRAuthOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSRAuthOracle creates a new instance of DSRAuthOracle, bound to a specific deployed contract.
func NewDSRAuthOracle(address common.Address, backend bind.ContractBackend) (*DSRAuthOracle, error) {
	contract, err := bindDSRAuthOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSRAuthOracle{DSRAuthOracleCaller: DSRAuthOracleCaller{contract: contract}, DSRAuthOracleTransactor: DSRAuthOracleTransactor{contract: contract}, DSRAuthOracleFilterer: DSRAuthOracleFilterer{contract: contract}}, nil
}

// NewDSRAuthOracleCaller creates a new read-only instance of DSRAuthOracle, bound to a specific deployed contract.
func NewDSRAuthOracleCaller(address common.Address, caller bind.ContractCaller) (*DSRAuthOracleCaller, error) {
	contract, err := bindDSRAuthOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSRAuthOracleCaller{contract: contract}, nil
}

// NewDSRAuthOracleTransactor creates a new write-only instance of DSRAuthOracle, bound to a specific deployed contract.
func NewDSRAuthOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*DSRAuthOracleTransactor, error) {
	contract, err := bindDSRAuthOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSRAuthOracleTransactor{contract: contract}, nil
}

// NewDSRAuthOracleFilterer creates a new log filterer instance of DSRAuthOracle, bound to a specific deployed contract.
func NewDSRAuthOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*DSRAuthOracleFilterer, error) {
	contract, err := bindDSRAuthOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSRAuthOracleFilterer{contract: contract}, nil
}

// bindDSRAuthOracle binds a generic wrapper to an already deployed contract.
func bindDSRAuthOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DSRAuthOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSRAuthOracle *DSRAuthOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DSRAuthOracle.Contract.DSRAuthOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSRAuthOracle *DSRAuthOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.DSRAuthOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSRAuthOracle *DSRAuthOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.DSRAuthOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSRAuthOracle *DSRAuthOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DSRAuthOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSRAuthOracle *DSRAuthOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSRAuthOracle *DSRAuthOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.contract.Transact(opts, method, params...)
}

// DATAPROVIDERROLE is a free data retrieval call binding the contract method 0x3aeae31d.
//
// Solidity: function DATA_PROVIDER_ROLE() view returns(bytes32)
func (_DSRAuthOracle *DSRAuthOracleCaller) DATAPROVIDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "DATA_PROVIDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DATAPROVIDERROLE is a free data retrieval call binding the contract method 0x3aeae31d.
//
// Solidity: function DATA_PROVIDER_ROLE() view returns(bytes32)
func (_DSRAuthOracle *DSRAuthOracleSession) DATAPROVIDERROLE() ([32]byte, error) {
	return _DSRAuthOracle.Contract.DATAPROVIDERROLE(&_DSRAuthOracle.CallOpts)
}

// DATAPROVIDERROLE is a free data retrieval call binding the contract method 0x3aeae31d.
//
// Solidity: function DATA_PROVIDER_ROLE() view returns(bytes32)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) DATAPROVIDERROLE() ([32]byte, error) {
	return _DSRAuthOracle.Contract.DATAPROVIDERROLE(&_DSRAuthOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DSRAuthOracle *DSRAuthOracleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DSRAuthOracle *DSRAuthOracleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DSRAuthOracle.Contract.DEFAULTADMINROLE(&_DSRAuthOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DSRAuthOracle.Contract.DEFAULTADMINROLE(&_DSRAuthOracle.CallOpts)
}

// GetAPR is a free data retrieval call binding the contract method 0xc89d5b8b.
//
// Solidity: function getAPR() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetAPR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getAPR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAPR is a free data retrieval call binding the contract method 0xc89d5b8b.
//
// Solidity: function getAPR() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) GetAPR() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetAPR(&_DSRAuthOracle.CallOpts)
}

// GetAPR is a free data retrieval call binding the contract method 0xc89d5b8b.
//
// Solidity: function getAPR() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetAPR() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetAPR(&_DSRAuthOracle.CallOpts)
}

// GetChi is a free data retrieval call binding the contract method 0x0416c1a5.
//
// Solidity: function getChi() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetChi(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getChi")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChi is a free data retrieval call binding the contract method 0x0416c1a5.
//
// Solidity: function getChi() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) GetChi() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetChi(&_DSRAuthOracle.CallOpts)
}

// GetChi is a free data retrieval call binding the contract method 0x0416c1a5.
//
// Solidity: function getChi() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetChi() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetChi(&_DSRAuthOracle.CallOpts)
}

// GetConversionRate is a free data retrieval call binding the contract method 0x6e5b6b28.
//
// Solidity: function getConversionRate(uint256 timestamp) view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetConversionRate(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getConversionRate", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConversionRate is a free data retrieval call binding the contract method 0x6e5b6b28.
//
// Solidity: function getConversionRate(uint256 timestamp) view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) GetConversionRate(timestamp *big.Int) (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRate(&_DSRAuthOracle.CallOpts, timestamp)
}

// GetConversionRate is a free data retrieval call binding the contract method 0x6e5b6b28.
//
// Solidity: function getConversionRate(uint256 timestamp) view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetConversionRate(timestamp *big.Int) (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRate(&_DSRAuthOracle.CallOpts, timestamp)
}

// GetConversionRate0 is a free data retrieval call binding the contract method 0xf36089ec.
//
// Solidity: function getConversionRate() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetConversionRate0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getConversionRate0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConversionRate0 is a free data retrieval call binding the contract method 0xf36089ec.
//
// Solidity: function getConversionRate() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) GetConversionRate0() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRate0(&_DSRAuthOracle.CallOpts)
}

// GetConversionRate0 is a free data retrieval call binding the contract method 0xf36089ec.
//
// Solidity: function getConversionRate() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetConversionRate0() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRate0(&_DSRAuthOracle.CallOpts)
}

// GetConversionRateBinomialApprox is a free data retrieval call binding the contract method 0x018de541.
//
// Solidity: function getConversionRateBinomialApprox(uint256 timestamp) view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetConversionRateBinomialApprox(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getConversionRateBinomialApprox", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConversionRateBinomialApprox is a free data retrieval call binding the contract method 0x018de541.
//
// Solidity: function getConversionRateBinomialApprox(uint256 timestamp) view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) GetConversionRateBinomialApprox(timestamp *big.Int) (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRateBinomialApprox(&_DSRAuthOracle.CallOpts, timestamp)
}

// GetConversionRateBinomialApprox is a free data retrieval call binding the contract method 0x018de541.
//
// Solidity: function getConversionRateBinomialApprox(uint256 timestamp) view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetConversionRateBinomialApprox(timestamp *big.Int) (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRateBinomialApprox(&_DSRAuthOracle.CallOpts, timestamp)
}

// GetConversionRateBinomialApprox0 is a free data retrieval call binding the contract method 0x96376040.
//
// Solidity: function getConversionRateBinomialApprox() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetConversionRateBinomialApprox0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getConversionRateBinomialApprox0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConversionRateBinomialApprox0 is a free data retrieval call binding the contract method 0x96376040.
//
// Solidity: function getConversionRateBinomialApprox() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) GetConversionRateBinomialApprox0() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRateBinomialApprox0(&_DSRAuthOracle.CallOpts)
}

// GetConversionRateBinomialApprox0 is a free data retrieval call binding the contract method 0x96376040.
//
// Solidity: function getConversionRateBinomialApprox() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetConversionRateBinomialApprox0() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRateBinomialApprox0(&_DSRAuthOracle.CallOpts)
}

// GetConversionRateLinearApprox is a free data retrieval call binding the contract method 0x01bfbc12.
//
// Solidity: function getConversionRateLinearApprox(uint256 timestamp) view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetConversionRateLinearApprox(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getConversionRateLinearApprox", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConversionRateLinearApprox is a free data retrieval call binding the contract method 0x01bfbc12.
//
// Solidity: function getConversionRateLinearApprox(uint256 timestamp) view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) GetConversionRateLinearApprox(timestamp *big.Int) (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRateLinearApprox(&_DSRAuthOracle.CallOpts, timestamp)
}

// GetConversionRateLinearApprox is a free data retrieval call binding the contract method 0x01bfbc12.
//
// Solidity: function getConversionRateLinearApprox(uint256 timestamp) view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetConversionRateLinearApprox(timestamp *big.Int) (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRateLinearApprox(&_DSRAuthOracle.CallOpts, timestamp)
}

// GetConversionRateLinearApprox0 is a free data retrieval call binding the contract method 0x3dd36646.
//
// Solidity: function getConversionRateLinearApprox() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetConversionRateLinearApprox0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getConversionRateLinearApprox0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConversionRateLinearApprox0 is a free data retrieval call binding the contract method 0x3dd36646.
//
// Solidity: function getConversionRateLinearApprox() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) GetConversionRateLinearApprox0() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRateLinearApprox0(&_DSRAuthOracle.CallOpts)
}

// GetConversionRateLinearApprox0 is a free data retrieval call binding the contract method 0x3dd36646.
//
// Solidity: function getConversionRateLinearApprox() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetConversionRateLinearApprox0() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetConversionRateLinearApprox0(&_DSRAuthOracle.CallOpts)
}

// GetDSR is a free data retrieval call binding the contract method 0x92b3dd6a.
//
// Solidity: function getDSR() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetDSR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getDSR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDSR is a free data retrieval call binding the contract method 0x92b3dd6a.
//
// Solidity: function getDSR() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) GetDSR() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetDSR(&_DSRAuthOracle.CallOpts)
}

// GetDSR is a free data retrieval call binding the contract method 0x92b3dd6a.
//
// Solidity: function getDSR() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetDSR() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetDSR(&_DSRAuthOracle.CallOpts)
}

// GetPotData is a free data retrieval call binding the contract method 0xccebbbcf.
//
// Solidity: function getPotData() view returns((uint96,uint120,uint40))
func (_DSRAuthOracle *DSRAuthOracleCaller) GetPotData(opts *bind.CallOpts) (IDSROraclePotData, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getPotData")

	if err != nil {
		return *new(IDSROraclePotData), err
	}

	out0 := *abi.ConvertType(out[0], new(IDSROraclePotData)).(*IDSROraclePotData)

	return out0, err

}

// GetPotData is a free data retrieval call binding the contract method 0xccebbbcf.
//
// Solidity: function getPotData() view returns((uint96,uint120,uint40))
func (_DSRAuthOracle *DSRAuthOracleSession) GetPotData() (IDSROraclePotData, error) {
	return _DSRAuthOracle.Contract.GetPotData(&_DSRAuthOracle.CallOpts)
}

// GetPotData is a free data retrieval call binding the contract method 0xccebbbcf.
//
// Solidity: function getPotData() view returns((uint96,uint120,uint40))
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetPotData() (IDSROraclePotData, error) {
	return _DSRAuthOracle.Contract.GetPotData(&_DSRAuthOracle.CallOpts)
}

// GetRho is a free data retrieval call binding the contract method 0xb4c7ec78.
//
// Solidity: function getRho() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetRho(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getRho")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRho is a free data retrieval call binding the contract method 0xb4c7ec78.
//
// Solidity: function getRho() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) GetRho() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetRho(&_DSRAuthOracle.CallOpts)
}

// GetRho is a free data retrieval call binding the contract method 0xb4c7ec78.
//
// Solidity: function getRho() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetRho() (*big.Int, error) {
	return _DSRAuthOracle.Contract.GetRho(&_DSRAuthOracle.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DSRAuthOracle *DSRAuthOracleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DSRAuthOracle *DSRAuthOracleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DSRAuthOracle.Contract.GetRoleAdmin(&_DSRAuthOracle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DSRAuthOracle.Contract.GetRoleAdmin(&_DSRAuthOracle.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DSRAuthOracle *DSRAuthOracleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DSRAuthOracle *DSRAuthOracleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DSRAuthOracle.Contract.HasRole(&_DSRAuthOracle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DSRAuthOracle.Contract.HasRole(&_DSRAuthOracle.CallOpts, role, account)
}

// MaxDSR is a free data retrieval call binding the contract method 0xab62f3c7.
//
// Solidity: function maxDSR() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCaller) MaxDSR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "maxDSR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDSR is a free data retrieval call binding the contract method 0xab62f3c7.
//
// Solidity: function maxDSR() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleSession) MaxDSR() (*big.Int, error) {
	return _DSRAuthOracle.Contract.MaxDSR(&_DSRAuthOracle.CallOpts)
}

// MaxDSR is a free data retrieval call binding the contract method 0xab62f3c7.
//
// Solidity: function maxDSR() view returns(uint256)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) MaxDSR() (*big.Int, error) {
	return _DSRAuthOracle.Contract.MaxDSR(&_DSRAuthOracle.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DSRAuthOracle *DSRAuthOracleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DSRAuthOracle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DSRAuthOracle *DSRAuthOracleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DSRAuthOracle.Contract.SupportsInterface(&_DSRAuthOracle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DSRAuthOracle *DSRAuthOracleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DSRAuthOracle.Contract.SupportsInterface(&_DSRAuthOracle.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DSRAuthOracle *DSRAuthOracleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DSRAuthOracle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DSRAuthOracle *DSRAuthOracleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.GrantRole(&_DSRAuthOracle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DSRAuthOracle *DSRAuthOracleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.GrantRole(&_DSRAuthOracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DSRAuthOracle *DSRAuthOracleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DSRAuthOracle.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DSRAuthOracle *DSRAuthOracleSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.RenounceRole(&_DSRAuthOracle.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DSRAuthOracle *DSRAuthOracleTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.RenounceRole(&_DSRAuthOracle.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DSRAuthOracle *DSRAuthOracleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DSRAuthOracle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DSRAuthOracle *DSRAuthOracleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.RevokeRole(&_DSRAuthOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DSRAuthOracle *DSRAuthOracleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.RevokeRole(&_DSRAuthOracle.TransactOpts, role, account)
}

// SetMaxDSR is a paid mutator transaction binding the contract method 0xbe13c1e3.
//
// Solidity: function setMaxDSR(uint256 _maxDSR) returns()
func (_DSRAuthOracle *DSRAuthOracleTransactor) SetMaxDSR(opts *bind.TransactOpts, _maxDSR *big.Int) (*types.Transaction, error) {
	return _DSRAuthOracle.contract.Transact(opts, "setMaxDSR", _maxDSR)
}

// SetMaxDSR is a paid mutator transaction binding the contract method 0xbe13c1e3.
//
// Solidity: function setMaxDSR(uint256 _maxDSR) returns()
func (_DSRAuthOracle *DSRAuthOracleSession) SetMaxDSR(_maxDSR *big.Int) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.SetMaxDSR(&_DSRAuthOracle.TransactOpts, _maxDSR)
}

// SetMaxDSR is a paid mutator transaction binding the contract method 0xbe13c1e3.
//
// Solidity: function setMaxDSR(uint256 _maxDSR) returns()
func (_DSRAuthOracle *DSRAuthOracleTransactorSession) SetMaxDSR(_maxDSR *big.Int) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.SetMaxDSR(&_DSRAuthOracle.TransactOpts, _maxDSR)
}

// SetPotData is a paid mutator transaction binding the contract method 0x58aff2ef.
//
// Solidity: function setPotData((uint96,uint120,uint40) nextData) returns()
func (_DSRAuthOracle *DSRAuthOracleTransactor) SetPotData(opts *bind.TransactOpts, nextData IDSROraclePotData) (*types.Transaction, error) {
	return _DSRAuthOracle.contract.Transact(opts, "setPotData", nextData)
}

// SetPotData is a paid mutator transaction binding the contract method 0x58aff2ef.
//
// Solidity: function setPotData((uint96,uint120,uint40) nextData) returns()
func (_DSRAuthOracle *DSRAuthOracleSession) SetPotData(nextData IDSROraclePotData) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.SetPotData(&_DSRAuthOracle.TransactOpts, nextData)
}

// SetPotData is a paid mutator transaction binding the contract method 0x58aff2ef.
//
// Solidity: function setPotData((uint96,uint120,uint40) nextData) returns()
func (_DSRAuthOracle *DSRAuthOracleTransactorSession) SetPotData(nextData IDSROraclePotData) (*types.Transaction, error) {
	return _DSRAuthOracle.Contract.SetPotData(&_DSRAuthOracle.TransactOpts, nextData)
}

// DSRAuthOracleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DSRAuthOracle contract.
type DSRAuthOracleRoleAdminChangedIterator struct {
	Event *DSRAuthOracleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DSRAuthOracleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSRAuthOracleRoleAdminChanged)
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
		it.Event = new(DSRAuthOracleRoleAdminChanged)
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
func (it *DSRAuthOracleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSRAuthOracleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSRAuthOracleRoleAdminChanged represents a RoleAdminChanged event raised by the DSRAuthOracle contract.
type DSRAuthOracleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DSRAuthOracle *DSRAuthOracleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DSRAuthOracleRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DSRAuthOracle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DSRAuthOracleRoleAdminChangedIterator{contract: _DSRAuthOracle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DSRAuthOracle *DSRAuthOracleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DSRAuthOracleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DSRAuthOracle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSRAuthOracleRoleAdminChanged)
				if err := _DSRAuthOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DSRAuthOracle *DSRAuthOracleFilterer) ParseRoleAdminChanged(log types.Log) (*DSRAuthOracleRoleAdminChanged, error) {
	event := new(DSRAuthOracleRoleAdminChanged)
	if err := _DSRAuthOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DSRAuthOracleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DSRAuthOracle contract.
type DSRAuthOracleRoleGrantedIterator struct {
	Event *DSRAuthOracleRoleGranted // Event containing the contract specifics and raw log

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
func (it *DSRAuthOracleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSRAuthOracleRoleGranted)
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
		it.Event = new(DSRAuthOracleRoleGranted)
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
func (it *DSRAuthOracleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSRAuthOracleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSRAuthOracleRoleGranted represents a RoleGranted event raised by the DSRAuthOracle contract.
type DSRAuthOracleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DSRAuthOracle *DSRAuthOracleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DSRAuthOracleRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DSRAuthOracle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DSRAuthOracleRoleGrantedIterator{contract: _DSRAuthOracle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DSRAuthOracle *DSRAuthOracleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DSRAuthOracleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DSRAuthOracle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSRAuthOracleRoleGranted)
				if err := _DSRAuthOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DSRAuthOracle *DSRAuthOracleFilterer) ParseRoleGranted(log types.Log) (*DSRAuthOracleRoleGranted, error) {
	event := new(DSRAuthOracleRoleGranted)
	if err := _DSRAuthOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DSRAuthOracleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DSRAuthOracle contract.
type DSRAuthOracleRoleRevokedIterator struct {
	Event *DSRAuthOracleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DSRAuthOracleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSRAuthOracleRoleRevoked)
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
		it.Event = new(DSRAuthOracleRoleRevoked)
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
func (it *DSRAuthOracleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSRAuthOracleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSRAuthOracleRoleRevoked represents a RoleRevoked event raised by the DSRAuthOracle contract.
type DSRAuthOracleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DSRAuthOracle *DSRAuthOracleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DSRAuthOracleRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DSRAuthOracle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DSRAuthOracleRoleRevokedIterator{contract: _DSRAuthOracle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DSRAuthOracle *DSRAuthOracleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DSRAuthOracleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DSRAuthOracle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSRAuthOracleRoleRevoked)
				if err := _DSRAuthOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DSRAuthOracle *DSRAuthOracleFilterer) ParseRoleRevoked(log types.Log) (*DSRAuthOracleRoleRevoked, error) {
	event := new(DSRAuthOracleRoleRevoked)
	if err := _DSRAuthOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DSRAuthOracleSetMaxDSRIterator is returned from FilterSetMaxDSR and is used to iterate over the raw logs and unpacked data for SetMaxDSR events raised by the DSRAuthOracle contract.
type DSRAuthOracleSetMaxDSRIterator struct {
	Event *DSRAuthOracleSetMaxDSR // Event containing the contract specifics and raw log

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
func (it *DSRAuthOracleSetMaxDSRIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSRAuthOracleSetMaxDSR)
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
		it.Event = new(DSRAuthOracleSetMaxDSR)
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
func (it *DSRAuthOracleSetMaxDSRIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSRAuthOracleSetMaxDSRIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSRAuthOracleSetMaxDSR represents a SetMaxDSR event raised by the DSRAuthOracle contract.
type DSRAuthOracleSetMaxDSR struct {
	MaxDSR *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMaxDSR is a free log retrieval operation binding the contract event 0xf756eadc4b6b9531a0ec2131cd12edbe371b141ab6f78c183c66398f14f5e3f7.
//
// Solidity: event SetMaxDSR(uint256 maxDSR)
func (_DSRAuthOracle *DSRAuthOracleFilterer) FilterSetMaxDSR(opts *bind.FilterOpts) (*DSRAuthOracleSetMaxDSRIterator, error) {

	logs, sub, err := _DSRAuthOracle.contract.FilterLogs(opts, "SetMaxDSR")
	if err != nil {
		return nil, err
	}
	return &DSRAuthOracleSetMaxDSRIterator{contract: _DSRAuthOracle.contract, event: "SetMaxDSR", logs: logs, sub: sub}, nil
}

// WatchSetMaxDSR is a free log subscription operation binding the contract event 0xf756eadc4b6b9531a0ec2131cd12edbe371b141ab6f78c183c66398f14f5e3f7.
//
// Solidity: event SetMaxDSR(uint256 maxDSR)
func (_DSRAuthOracle *DSRAuthOracleFilterer) WatchSetMaxDSR(opts *bind.WatchOpts, sink chan<- *DSRAuthOracleSetMaxDSR) (event.Subscription, error) {

	logs, sub, err := _DSRAuthOracle.contract.WatchLogs(opts, "SetMaxDSR")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSRAuthOracleSetMaxDSR)
				if err := _DSRAuthOracle.contract.UnpackLog(event, "SetMaxDSR", log); err != nil {
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

// ParseSetMaxDSR is a log parse operation binding the contract event 0xf756eadc4b6b9531a0ec2131cd12edbe371b141ab6f78c183c66398f14f5e3f7.
//
// Solidity: event SetMaxDSR(uint256 maxDSR)
func (_DSRAuthOracle *DSRAuthOracleFilterer) ParseSetMaxDSR(log types.Log) (*DSRAuthOracleSetMaxDSR, error) {
	event := new(DSRAuthOracleSetMaxDSR)
	if err := _DSRAuthOracle.contract.UnpackLog(event, "SetMaxDSR", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DSRAuthOracleSetPotDataIterator is returned from FilterSetPotData and is used to iterate over the raw logs and unpacked data for SetPotData events raised by the DSRAuthOracle contract.
type DSRAuthOracleSetPotDataIterator struct {
	Event *DSRAuthOracleSetPotData // Event containing the contract specifics and raw log

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
func (it *DSRAuthOracleSetPotDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSRAuthOracleSetPotData)
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
		it.Event = new(DSRAuthOracleSetPotData)
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
func (it *DSRAuthOracleSetPotDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSRAuthOracleSetPotDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSRAuthOracleSetPotData represents a SetPotData event raised by the DSRAuthOracle contract.
type DSRAuthOracleSetPotData struct {
	NextData IDSROraclePotData
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPotData is a free log retrieval operation binding the contract event 0x90e88bcfdf9cf7d631c6b9801cc9f39ad830d176408f246b9db0dc2b3aaa832b.
//
// Solidity: event SetPotData((uint96,uint120,uint40) nextData)
func (_DSRAuthOracle *DSRAuthOracleFilterer) FilterSetPotData(opts *bind.FilterOpts) (*DSRAuthOracleSetPotDataIterator, error) {

	logs, sub, err := _DSRAuthOracle.contract.FilterLogs(opts, "SetPotData")
	if err != nil {
		return nil, err
	}
	return &DSRAuthOracleSetPotDataIterator{contract: _DSRAuthOracle.contract, event: "SetPotData", logs: logs, sub: sub}, nil
}

// WatchSetPotData is a free log subscription operation binding the contract event 0x90e88bcfdf9cf7d631c6b9801cc9f39ad830d176408f246b9db0dc2b3aaa832b.
//
// Solidity: event SetPotData((uint96,uint120,uint40) nextData)
func (_DSRAuthOracle *DSRAuthOracleFilterer) WatchSetPotData(opts *bind.WatchOpts, sink chan<- *DSRAuthOracleSetPotData) (event.Subscription, error) {

	logs, sub, err := _DSRAuthOracle.contract.WatchLogs(opts, "SetPotData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSRAuthOracleSetPotData)
				if err := _DSRAuthOracle.contract.UnpackLog(event, "SetPotData", log); err != nil {
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

// ParseSetPotData is a log parse operation binding the contract event 0x90e88bcfdf9cf7d631c6b9801cc9f39ad830d176408f246b9db0dc2b3aaa832b.
//
// Solidity: event SetPotData((uint96,uint120,uint40) nextData)
func (_DSRAuthOracle *DSRAuthOracleFilterer) ParseSetPotData(log types.Log) (*DSRAuthOracleSetPotData, error) {
	event := new(DSRAuthOracleSetPotData)
	if err := _DSRAuthOracle.contract.UnpackLog(event, "SetPotData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
