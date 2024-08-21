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

// LibSwapSwapData is an auto generated low-level Go binding around an user-defined struct.
type LibSwapSwapData struct {
	CallTo           common.Address
	ApproveTo        common.Address
	SendingAssetId   common.Address
	ReceivingAssetId common.Address
	FromAmount       *big.Int
	CallData         []byte
	RequiresDeposit  bool
}

// NovaVaultV2MetaData contains all meta data concerning the NovaVaultV2 contract.
var NovaVaultV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sDAI\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swapFacet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ContractNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GenericSwapFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetId\",\"type\":\"address\"}],\"name\":\"InvalidAssetId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Referral\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"addDex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"requiresDeposit\",\"type\":\"bool\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"setFunctionApprovalBySignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"callTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sendingAssetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receivingAssetId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"requiresDeposit\",\"type\":\"bool\"}],\"internalType\":\"structLibSwap.SwapData[]\",\"name\":\"swapData\",\"type\":\"tuple[]\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NovaVaultV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use NovaVaultV2MetaData.ABI instead.
var NovaVaultV2ABI = NovaVaultV2MetaData.ABI

// NovaVaultV2 is an auto generated Go binding around an Ethereum contract.
type NovaVaultV2 struct {
	NovaVaultV2Caller     // Read-only binding to the contract
	NovaVaultV2Transactor // Write-only binding to the contract
	NovaVaultV2Filterer   // Log filterer for contract events
}

// NovaVaultV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type NovaVaultV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NovaVaultV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type NovaVaultV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NovaVaultV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NovaVaultV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NovaVaultV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NovaVaultV2Session struct {
	Contract     *NovaVaultV2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NovaVaultV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NovaVaultV2CallerSession struct {
	Contract *NovaVaultV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NovaVaultV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NovaVaultV2TransactorSession struct {
	Contract     *NovaVaultV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NovaVaultV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type NovaVaultV2Raw struct {
	Contract *NovaVaultV2 // Generic contract binding to access the raw methods on
}

// NovaVaultV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NovaVaultV2CallerRaw struct {
	Contract *NovaVaultV2Caller // Generic read-only contract binding to access the raw methods on
}

// NovaVaultV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NovaVaultV2TransactorRaw struct {
	Contract *NovaVaultV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewNovaVaultV2 creates a new instance of NovaVaultV2, bound to a specific deployed contract.
func NewNovaVaultV2(address common.Address, backend bind.ContractBackend) (*NovaVaultV2, error) {
	contract, err := bindNovaVaultV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NovaVaultV2{NovaVaultV2Caller: NovaVaultV2Caller{contract: contract}, NovaVaultV2Transactor: NovaVaultV2Transactor{contract: contract}, NovaVaultV2Filterer: NovaVaultV2Filterer{contract: contract}}, nil
}

// NewNovaVaultV2Caller creates a new read-only instance of NovaVaultV2, bound to a specific deployed contract.
func NewNovaVaultV2Caller(address common.Address, caller bind.ContractCaller) (*NovaVaultV2Caller, error) {
	contract, err := bindNovaVaultV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NovaVaultV2Caller{contract: contract}, nil
}

// NewNovaVaultV2Transactor creates a new write-only instance of NovaVaultV2, bound to a specific deployed contract.
func NewNovaVaultV2Transactor(address common.Address, transactor bind.ContractTransactor) (*NovaVaultV2Transactor, error) {
	contract, err := bindNovaVaultV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NovaVaultV2Transactor{contract: contract}, nil
}

// NewNovaVaultV2Filterer creates a new log filterer instance of NovaVaultV2, bound to a specific deployed contract.
func NewNovaVaultV2Filterer(address common.Address, filterer bind.ContractFilterer) (*NovaVaultV2Filterer, error) {
	contract, err := bindNovaVaultV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NovaVaultV2Filterer{contract: contract}, nil
}

// bindNovaVaultV2 binds a generic wrapper to an already deployed contract.
func bindNovaVaultV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NovaVaultV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NovaVaultV2 *NovaVaultV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NovaVaultV2.Contract.NovaVaultV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NovaVaultV2 *NovaVaultV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.NovaVaultV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NovaVaultV2 *NovaVaultV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.NovaVaultV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NovaVaultV2 *NovaVaultV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NovaVaultV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NovaVaultV2 *NovaVaultV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NovaVaultV2 *NovaVaultV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NovaVaultV2 *NovaVaultV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NovaVaultV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NovaVaultV2 *NovaVaultV2Session) Owner() (common.Address, error) {
	return _NovaVaultV2.Contract.Owner(&_NovaVaultV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NovaVaultV2 *NovaVaultV2CallerSession) Owner() (common.Address, error) {
	return _NovaVaultV2.Contract.Owner(&_NovaVaultV2.CallOpts)
}

// AddDex is a paid mutator transaction binding the contract method 0x536db266.
//
// Solidity: function addDex(address _contract) returns()
func (_NovaVaultV2 *NovaVaultV2Transactor) AddDex(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _NovaVaultV2.contract.Transact(opts, "addDex", _contract)
}

// AddDex is a paid mutator transaction binding the contract method 0x536db266.
//
// Solidity: function addDex(address _contract) returns()
func (_NovaVaultV2 *NovaVaultV2Session) AddDex(_contract common.Address) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.AddDex(&_NovaVaultV2.TransactOpts, _contract)
}

// AddDex is a paid mutator transaction binding the contract method 0x536db266.
//
// Solidity: function addDex(address _contract) returns()
func (_NovaVaultV2 *NovaVaultV2TransactorSession) AddDex(_contract common.Address) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.AddDex(&_NovaVaultV2.TransactOpts, _contract)
}

// Deposit is a paid mutator transaction binding the contract method 0x8b7cb3c3.
//
// Solidity: function deposit((address,address,address,address,uint256,bytes,bool)[] swapData, uint16 referral) payable returns(bool, uint256)
func (_NovaVaultV2 *NovaVaultV2Transactor) Deposit(opts *bind.TransactOpts, swapData []LibSwapSwapData, referral uint16) (*types.Transaction, error) {
	return _NovaVaultV2.contract.Transact(opts, "deposit", swapData, referral)
}

// Deposit is a paid mutator transaction binding the contract method 0x8b7cb3c3.
//
// Solidity: function deposit((address,address,address,address,uint256,bytes,bool)[] swapData, uint16 referral) payable returns(bool, uint256)
func (_NovaVaultV2 *NovaVaultV2Session) Deposit(swapData []LibSwapSwapData, referral uint16) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.Deposit(&_NovaVaultV2.TransactOpts, swapData, referral)
}

// Deposit is a paid mutator transaction binding the contract method 0x8b7cb3c3.
//
// Solidity: function deposit((address,address,address,address,uint256,bytes,bool)[] swapData, uint16 referral) payable returns(bool, uint256)
func (_NovaVaultV2 *NovaVaultV2TransactorSession) Deposit(swapData []LibSwapSwapData, referral uint16) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.Deposit(&_NovaVaultV2.TransactOpts, swapData, referral)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NovaVaultV2 *NovaVaultV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NovaVaultV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NovaVaultV2 *NovaVaultV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _NovaVaultV2.Contract.RenounceOwnership(&_NovaVaultV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NovaVaultV2 *NovaVaultV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NovaVaultV2.Contract.RenounceOwnership(&_NovaVaultV2.TransactOpts)
}

// SetFunctionApprovalBySignature is a paid mutator transaction binding the contract method 0xc561c148.
//
// Solidity: function setFunctionApprovalBySignature(bytes4 _selector) returns()
func (_NovaVaultV2 *NovaVaultV2Transactor) SetFunctionApprovalBySignature(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _NovaVaultV2.contract.Transact(opts, "setFunctionApprovalBySignature", _selector)
}

// SetFunctionApprovalBySignature is a paid mutator transaction binding the contract method 0xc561c148.
//
// Solidity: function setFunctionApprovalBySignature(bytes4 _selector) returns()
func (_NovaVaultV2 *NovaVaultV2Session) SetFunctionApprovalBySignature(_selector [4]byte) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.SetFunctionApprovalBySignature(&_NovaVaultV2.TransactOpts, _selector)
}

// SetFunctionApprovalBySignature is a paid mutator transaction binding the contract method 0xc561c148.
//
// Solidity: function setFunctionApprovalBySignature(bytes4 _selector) returns()
func (_NovaVaultV2 *NovaVaultV2TransactorSession) SetFunctionApprovalBySignature(_selector [4]byte) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.SetFunctionApprovalBySignature(&_NovaVaultV2.TransactOpts, _selector)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NovaVaultV2 *NovaVaultV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NovaVaultV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NovaVaultV2 *NovaVaultV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.TransferOwnership(&_NovaVaultV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NovaVaultV2 *NovaVaultV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.TransferOwnership(&_NovaVaultV2.TransactOpts, newOwner)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_NovaVaultV2 *NovaVaultV2Transactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _NovaVaultV2.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_NovaVaultV2 *NovaVaultV2Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.UniswapV3SwapCallback(&_NovaVaultV2.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_NovaVaultV2 *NovaVaultV2TransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.UniswapV3SwapCallback(&_NovaVaultV2.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7d4b9878.
//
// Solidity: function withdraw(uint16 referral, (address,address,address,address,uint256,bytes,bool)[] swapData) returns(bool, uint256)
func (_NovaVaultV2 *NovaVaultV2Transactor) Withdraw(opts *bind.TransactOpts, referral uint16, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _NovaVaultV2.contract.Transact(opts, "withdraw", referral, swapData)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7d4b9878.
//
// Solidity: function withdraw(uint16 referral, (address,address,address,address,uint256,bytes,bool)[] swapData) returns(bool, uint256)
func (_NovaVaultV2 *NovaVaultV2Session) Withdraw(referral uint16, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.Withdraw(&_NovaVaultV2.TransactOpts, referral, swapData)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7d4b9878.
//
// Solidity: function withdraw(uint16 referral, (address,address,address,address,uint256,bytes,bool)[] swapData) returns(bool, uint256)
func (_NovaVaultV2 *NovaVaultV2TransactorSession) Withdraw(referral uint16, swapData []LibSwapSwapData) (*types.Transaction, error) {
	return _NovaVaultV2.Contract.Withdraw(&_NovaVaultV2.TransactOpts, referral, swapData)
}

// NovaVaultV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NovaVaultV2 contract.
type NovaVaultV2OwnershipTransferredIterator struct {
	Event *NovaVaultV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NovaVaultV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NovaVaultV2OwnershipTransferred)
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
		it.Event = new(NovaVaultV2OwnershipTransferred)
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
func (it *NovaVaultV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NovaVaultV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NovaVaultV2OwnershipTransferred represents a OwnershipTransferred event raised by the NovaVaultV2 contract.
type NovaVaultV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NovaVaultV2 *NovaVaultV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NovaVaultV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NovaVaultV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NovaVaultV2OwnershipTransferredIterator{contract: _NovaVaultV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NovaVaultV2 *NovaVaultV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NovaVaultV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NovaVaultV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NovaVaultV2OwnershipTransferred)
				if err := _NovaVaultV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NovaVaultV2 *NovaVaultV2Filterer) ParseOwnershipTransferred(log types.Log) (*NovaVaultV2OwnershipTransferred, error) {
	event := new(NovaVaultV2OwnershipTransferred)
	if err := _NovaVaultV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NovaVaultV2ReferralIterator is returned from FilterReferral and is used to iterate over the raw logs and unpacked data for Referral events raised by the NovaVaultV2 contract.
type NovaVaultV2ReferralIterator struct {
	Event *NovaVaultV2Referral // Event containing the contract specifics and raw log

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
func (it *NovaVaultV2ReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NovaVaultV2Referral)
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
		it.Event = new(NovaVaultV2Referral)
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
func (it *NovaVaultV2ReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NovaVaultV2ReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NovaVaultV2Referral represents a Referral event raised by the NovaVaultV2 contract.
type NovaVaultV2Referral struct {
	Referral  uint16
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReferral is a free log retrieval operation binding the contract event 0x16902e34d01e8d5f80ce64939920a1390ff67f2546ae43ae72ac482033300968.
//
// Solidity: event Referral(uint16 referral, address indexed depositor, uint256 amount)
func (_NovaVaultV2 *NovaVaultV2Filterer) FilterReferral(opts *bind.FilterOpts, depositor []common.Address) (*NovaVaultV2ReferralIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _NovaVaultV2.contract.FilterLogs(opts, "Referral", depositorRule)
	if err != nil {
		return nil, err
	}
	return &NovaVaultV2ReferralIterator{contract: _NovaVaultV2.contract, event: "Referral", logs: logs, sub: sub}, nil
}

// WatchReferral is a free log subscription operation binding the contract event 0x16902e34d01e8d5f80ce64939920a1390ff67f2546ae43ae72ac482033300968.
//
// Solidity: event Referral(uint16 referral, address indexed depositor, uint256 amount)
func (_NovaVaultV2 *NovaVaultV2Filterer) WatchReferral(opts *bind.WatchOpts, sink chan<- *NovaVaultV2Referral, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _NovaVaultV2.contract.WatchLogs(opts, "Referral", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NovaVaultV2Referral)
				if err := _NovaVaultV2.contract.UnpackLog(event, "Referral", log); err != nil {
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

// ParseReferral is a log parse operation binding the contract event 0x16902e34d01e8d5f80ce64939920a1390ff67f2546ae43ae72ac482033300968.
//
// Solidity: event Referral(uint16 referral, address indexed depositor, uint256 amount)
func (_NovaVaultV2 *NovaVaultV2Filterer) ParseReferral(log types.Log) (*NovaVaultV2Referral, error) {
	event := new(NovaVaultV2Referral)
	if err := _NovaVaultV2.contract.UnpackLog(event, "Referral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
