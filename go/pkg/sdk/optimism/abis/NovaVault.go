// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimismNovaVault

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

// OptimismNovaVaultMetaData contains all meta data concerning the OptimismNovaVault contract.
var OptimismNovaVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_sDAI\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stables\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"novaAdapters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_novaAdapters\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"stable\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"referral\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"stable\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdapterApproval\",\"inputs\":[{\"name\":\"stable\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"adapter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Referral\",\"inputs\":[{\"name\":\"referral\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// OptimismNovaVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use OptimismNovaVaultMetaData.ABI instead.
var OptimismNovaVaultABI = OptimismNovaVaultMetaData.ABI

// OptimismNovaVault is an auto generated Go binding around an Ethereum contract.
type OptimismNovaVault struct {
	OptimismNovaVaultCaller     // Read-only binding to the contract
	OptimismNovaVaultTransactor // Write-only binding to the contract
	OptimismNovaVaultFilterer   // Log filterer for contract events
}

// OptimismNovaVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptimismNovaVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismNovaVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptimismNovaVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismNovaVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptimismNovaVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismNovaVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptimismNovaVaultSession struct {
	Contract     *OptimismNovaVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OptimismNovaVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptimismNovaVaultCallerSession struct {
	Contract *OptimismNovaVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OptimismNovaVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptimismNovaVaultTransactorSession struct {
	Contract     *OptimismNovaVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OptimismNovaVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptimismNovaVaultRaw struct {
	Contract *OptimismNovaVault // Generic contract binding to access the raw methods on
}

// OptimismNovaVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptimismNovaVaultCallerRaw struct {
	Contract *OptimismNovaVaultCaller // Generic read-only contract binding to access the raw methods on
}

// OptimismNovaVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptimismNovaVaultTransactorRaw struct {
	Contract *OptimismNovaVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptimismNovaVault creates a new instance of OptimismNovaVault, bound to a specific deployed contract.
func NewOptimismNovaVault(address common.Address, backend bind.ContractBackend) (*OptimismNovaVault, error) {
	contract, err := bindOptimismNovaVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismNovaVault{OptimismNovaVaultCaller: OptimismNovaVaultCaller{contract: contract}, OptimismNovaVaultTransactor: OptimismNovaVaultTransactor{contract: contract}, OptimismNovaVaultFilterer: OptimismNovaVaultFilterer{contract: contract}}, nil
}

// NewOptimismNovaVaultCaller creates a new read-only instance of OptimismNovaVault, bound to a specific deployed contract.
func NewOptimismNovaVaultCaller(address common.Address, caller bind.ContractCaller) (*OptimismNovaVaultCaller, error) {
	contract, err := bindOptimismNovaVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismNovaVaultCaller{contract: contract}, nil
}

// NewOptimismNovaVaultTransactor creates a new write-only instance of OptimismNovaVault, bound to a specific deployed contract.
func NewOptimismNovaVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismNovaVaultTransactor, error) {
	contract, err := bindOptimismNovaVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismNovaVaultTransactor{contract: contract}, nil
}

// NewOptimismNovaVaultFilterer creates a new log filterer instance of OptimismNovaVault, bound to a specific deployed contract.
func NewOptimismNovaVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismNovaVaultFilterer, error) {
	contract, err := bindOptimismNovaVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismNovaVaultFilterer{contract: contract}, nil
}

// bindOptimismNovaVault binds a generic wrapper to an already deployed contract.
func bindOptimismNovaVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismNovaVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismNovaVault *OptimismNovaVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismNovaVault.Contract.OptimismNovaVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismNovaVault *OptimismNovaVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismNovaVault.Contract.OptimismNovaVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismNovaVault *OptimismNovaVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismNovaVault.Contract.OptimismNovaVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismNovaVault *OptimismNovaVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismNovaVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismNovaVault *OptimismNovaVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismNovaVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismNovaVault *OptimismNovaVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismNovaVault.Contract.contract.Transact(opts, method, params...)
}

// NovaAdapters is a free data retrieval call binding the contract method 0xf2ff2812.
//
// Solidity: function _novaAdapters(address ) view returns(address)
func (_OptimismNovaVault *OptimismNovaVaultCaller) NovaAdapters(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _OptimismNovaVault.contract.Call(opts, &out, "_novaAdapters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NovaAdapters is a free data retrieval call binding the contract method 0xf2ff2812.
//
// Solidity: function _novaAdapters(address ) view returns(address)
func (_OptimismNovaVault *OptimismNovaVaultSession) NovaAdapters(arg0 common.Address) (common.Address, error) {
	return _OptimismNovaVault.Contract.NovaAdapters(&_OptimismNovaVault.CallOpts, arg0)
}

// NovaAdapters is a free data retrieval call binding the contract method 0xf2ff2812.
//
// Solidity: function _novaAdapters(address ) view returns(address)
func (_OptimismNovaVault *OptimismNovaVaultCallerSession) NovaAdapters(arg0 common.Address) (common.Address, error) {
	return _OptimismNovaVault.Contract.NovaAdapters(&_OptimismNovaVault.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address stable, uint256 assets, uint16 referral) returns(bool, uint256)
func (_OptimismNovaVault *OptimismNovaVaultTransactor) Deposit(opts *bind.TransactOpts, stable common.Address, assets *big.Int, referral uint16) (*types.Transaction, error) {
	return _OptimismNovaVault.contract.Transact(opts, "deposit", stable, assets, referral)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address stable, uint256 assets, uint16 referral) returns(bool, uint256)
func (_OptimismNovaVault *OptimismNovaVaultSession) Deposit(stable common.Address, assets *big.Int, referral uint16) (*types.Transaction, error) {
	return _OptimismNovaVault.Contract.Deposit(&_OptimismNovaVault.TransactOpts, stable, assets, referral)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address stable, uint256 assets, uint16 referral) returns(bool, uint256)
func (_OptimismNovaVault *OptimismNovaVaultTransactorSession) Deposit(stable common.Address, assets *big.Int, referral uint16) (*types.Transaction, error) {
	return _OptimismNovaVault.Contract.Deposit(&_OptimismNovaVault.TransactOpts, stable, assets, referral)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address stable, uint256 shares) returns(bool, uint256)
func (_OptimismNovaVault *OptimismNovaVaultTransactor) Withdraw(opts *bind.TransactOpts, stable common.Address, shares *big.Int) (*types.Transaction, error) {
	return _OptimismNovaVault.contract.Transact(opts, "withdraw", stable, shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address stable, uint256 shares) returns(bool, uint256)
func (_OptimismNovaVault *OptimismNovaVaultSession) Withdraw(stable common.Address, shares *big.Int) (*types.Transaction, error) {
	return _OptimismNovaVault.Contract.Withdraw(&_OptimismNovaVault.TransactOpts, stable, shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address stable, uint256 shares) returns(bool, uint256)
func (_OptimismNovaVault *OptimismNovaVaultTransactorSession) Withdraw(stable common.Address, shares *big.Int) (*types.Transaction, error) {
	return _OptimismNovaVault.Contract.Withdraw(&_OptimismNovaVault.TransactOpts, stable, shares)
}

// OptimismNovaVaultAdapterApprovalIterator is returned from FilterAdapterApproval and is used to iterate over the raw logs and unpacked data for AdapterApproval events raised by the OptimismNovaVault contract.
type OptimismNovaVaultAdapterApprovalIterator struct {
	Event *OptimismNovaVaultAdapterApproval // Event containing the contract specifics and raw log

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
func (it *OptimismNovaVaultAdapterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismNovaVaultAdapterApproval)
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
		it.Event = new(OptimismNovaVaultAdapterApproval)
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
func (it *OptimismNovaVaultAdapterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimismNovaVaultAdapterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimismNovaVaultAdapterApproval represents a AdapterApproval event raised by the OptimismNovaVault contract.
type OptimismNovaVaultAdapterApproval struct {
	Stable  common.Address
	Adapter common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdapterApproval is a free log retrieval operation binding the contract event 0x2a03353a3cafefd9bd1e60174ed856d38b75dc30ae89e43387bd92c4689a814b.
//
// Solidity: event AdapterApproval(address stable, address adapter)
func (_OptimismNovaVault *OptimismNovaVaultFilterer) FilterAdapterApproval(opts *bind.FilterOpts) (*OptimismNovaVaultAdapterApprovalIterator, error) {

	logs, sub, err := _OptimismNovaVault.contract.FilterLogs(opts, "AdapterApproval")
	if err != nil {
		return nil, err
	}
	return &OptimismNovaVaultAdapterApprovalIterator{contract: _OptimismNovaVault.contract, event: "AdapterApproval", logs: logs, sub: sub}, nil
}

// WatchAdapterApproval is a free log subscription operation binding the contract event 0x2a03353a3cafefd9bd1e60174ed856d38b75dc30ae89e43387bd92c4689a814b.
//
// Solidity: event AdapterApproval(address stable, address adapter)
func (_OptimismNovaVault *OptimismNovaVaultFilterer) WatchAdapterApproval(opts *bind.WatchOpts, sink chan<- *OptimismNovaVaultAdapterApproval) (event.Subscription, error) {

	logs, sub, err := _OptimismNovaVault.contract.WatchLogs(opts, "AdapterApproval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimismNovaVaultAdapterApproval)
				if err := _OptimismNovaVault.contract.UnpackLog(event, "AdapterApproval", log); err != nil {
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

// ParseAdapterApproval is a log parse operation binding the contract event 0x2a03353a3cafefd9bd1e60174ed856d38b75dc30ae89e43387bd92c4689a814b.
//
// Solidity: event AdapterApproval(address stable, address adapter)
func (_OptimismNovaVault *OptimismNovaVaultFilterer) ParseAdapterApproval(log types.Log) (*OptimismNovaVaultAdapterApproval, error) {
	event := new(OptimismNovaVaultAdapterApproval)
	if err := _OptimismNovaVault.contract.UnpackLog(event, "AdapterApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimismNovaVaultReferralIterator is returned from FilterReferral and is used to iterate over the raw logs and unpacked data for Referral events raised by the OptimismNovaVault contract.
type OptimismNovaVaultReferralIterator struct {
	Event *OptimismNovaVaultReferral // Event containing the contract specifics and raw log

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
func (it *OptimismNovaVaultReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismNovaVaultReferral)
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
		it.Event = new(OptimismNovaVaultReferral)
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
func (it *OptimismNovaVaultReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimismNovaVaultReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimismNovaVaultReferral represents a Referral event raised by the OptimismNovaVault contract.
type OptimismNovaVaultReferral struct {
	Referral  uint16
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReferral is a free log retrieval operation binding the contract event 0x16902e34d01e8d5f80ce64939920a1390ff67f2546ae43ae72ac482033300968.
//
// Solidity: event Referral(uint16 referral, address indexed depositor, uint256 amount)
func (_OptimismNovaVault *OptimismNovaVaultFilterer) FilterReferral(opts *bind.FilterOpts, depositor []common.Address) (*OptimismNovaVaultReferralIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _OptimismNovaVault.contract.FilterLogs(opts, "Referral", depositorRule)
	if err != nil {
		return nil, err
	}
	return &OptimismNovaVaultReferralIterator{contract: _OptimismNovaVault.contract, event: "Referral", logs: logs, sub: sub}, nil
}

// WatchReferral is a free log subscription operation binding the contract event 0x16902e34d01e8d5f80ce64939920a1390ff67f2546ae43ae72ac482033300968.
//
// Solidity: event Referral(uint16 referral, address indexed depositor, uint256 amount)
func (_OptimismNovaVault *OptimismNovaVaultFilterer) WatchReferral(opts *bind.WatchOpts, sink chan<- *OptimismNovaVaultReferral, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _OptimismNovaVault.contract.WatchLogs(opts, "Referral", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimismNovaVaultReferral)
				if err := _OptimismNovaVault.contract.UnpackLog(event, "Referral", log); err != nil {
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
func (_OptimismNovaVault *OptimismNovaVaultFilterer) ParseReferral(log types.Log) (*OptimismNovaVaultReferral, error) {
	event := new(OptimismNovaVaultReferral)
	if err := _OptimismNovaVault.contract.UnpackLog(event, "Referral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
