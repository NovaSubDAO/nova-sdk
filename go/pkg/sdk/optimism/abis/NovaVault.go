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

// NovaVaultMetaData contains all meta data concerning the NovaVault contract.
var NovaVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_sDAI\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stables\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"novaAdapters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_novaAdapters\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"stable\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"referral\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"stable\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdapterApproval\",\"inputs\":[{\"name\":\"stable\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"adapter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Referral\",\"inputs\":[{\"name\":\"referral\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// NovaVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use NovaVaultMetaData.ABI instead.
var NovaVaultABI = NovaVaultMetaData.ABI

// NovaVault is an auto generated Go binding around an Ethereum contract.
type NovaVault struct {
	NovaVaultCaller     // Read-only binding to the contract
	NovaVaultTransactor // Write-only binding to the contract
	NovaVaultFilterer   // Log filterer for contract events
}

// NovaVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type NovaVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NovaVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NovaVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NovaVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NovaVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NovaVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NovaVaultSession struct {
	Contract     *NovaVault        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NovaVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NovaVaultCallerSession struct {
	Contract *NovaVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NovaVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NovaVaultTransactorSession struct {
	Contract     *NovaVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NovaVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type NovaVaultRaw struct {
	Contract *NovaVault // Generic contract binding to access the raw methods on
}

// NovaVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NovaVaultCallerRaw struct {
	Contract *NovaVaultCaller // Generic read-only contract binding to access the raw methods on
}

// NovaVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NovaVaultTransactorRaw struct {
	Contract *NovaVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNovaVault creates a new instance of NovaVault, bound to a specific deployed contract.
func NewNovaVault(address common.Address, backend bind.ContractBackend) (*NovaVault, error) {
	contract, err := bindNovaVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NovaVault{NovaVaultCaller: NovaVaultCaller{contract: contract}, NovaVaultTransactor: NovaVaultTransactor{contract: contract}, NovaVaultFilterer: NovaVaultFilterer{contract: contract}}, nil
}

// NewNovaVaultCaller creates a new read-only instance of NovaVault, bound to a specific deployed contract.
func NewNovaVaultCaller(address common.Address, caller bind.ContractCaller) (*NovaVaultCaller, error) {
	contract, err := bindNovaVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NovaVaultCaller{contract: contract}, nil
}

// NewNovaVaultTransactor creates a new write-only instance of NovaVault, bound to a specific deployed contract.
func NewNovaVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*NovaVaultTransactor, error) {
	contract, err := bindNovaVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NovaVaultTransactor{contract: contract}, nil
}

// NewNovaVaultFilterer creates a new log filterer instance of NovaVault, bound to a specific deployed contract.
func NewNovaVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*NovaVaultFilterer, error) {
	contract, err := bindNovaVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NovaVaultFilterer{contract: contract}, nil
}

// bindNovaVault binds a generic wrapper to an already deployed contract.
func bindNovaVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NovaVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NovaVault *NovaVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NovaVault.Contract.NovaVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NovaVault *NovaVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NovaVault.Contract.NovaVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NovaVault *NovaVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NovaVault.Contract.NovaVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NovaVault *NovaVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NovaVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NovaVault *NovaVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NovaVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NovaVault *NovaVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NovaVault.Contract.contract.Transact(opts, method, params...)
}

// NovaAdapters is a free data retrieval call binding the contract method 0xf2ff2812.
//
// Solidity: function _novaAdapters(address ) view returns(address)
func (_NovaVault *NovaVaultCaller) NovaAdapters(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _NovaVault.contract.Call(opts, &out, "_novaAdapters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NovaAdapters is a free data retrieval call binding the contract method 0xf2ff2812.
//
// Solidity: function _novaAdapters(address ) view returns(address)
func (_NovaVault *NovaVaultSession) NovaAdapters(arg0 common.Address) (common.Address, error) {
	return _NovaVault.Contract.NovaAdapters(&_NovaVault.CallOpts, arg0)
}

// NovaAdapters is a free data retrieval call binding the contract method 0xf2ff2812.
//
// Solidity: function _novaAdapters(address ) view returns(address)
func (_NovaVault *NovaVaultCallerSession) NovaAdapters(arg0 common.Address) (common.Address, error) {
	return _NovaVault.Contract.NovaAdapters(&_NovaVault.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address stable, uint256 assets, uint16 referral) returns(bool, uint256)
func (_NovaVault *NovaVaultTransactor) Deposit(opts *bind.TransactOpts, stable common.Address, assets *big.Int, referral uint16) (*types.Transaction, error) {
	return _NovaVault.contract.Transact(opts, "deposit", stable, assets, referral)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address stable, uint256 assets, uint16 referral) returns(bool, uint256)
func (_NovaVault *NovaVaultSession) Deposit(stable common.Address, assets *big.Int, referral uint16) (*types.Transaction, error) {
	return _NovaVault.Contract.Deposit(&_NovaVault.TransactOpts, stable, assets, referral)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address stable, uint256 assets, uint16 referral) returns(bool, uint256)
func (_NovaVault *NovaVaultTransactorSession) Deposit(stable common.Address, assets *big.Int, referral uint16) (*types.Transaction, error) {
	return _NovaVault.Contract.Deposit(&_NovaVault.TransactOpts, stable, assets, referral)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address stable, uint256 shares) returns(bool, uint256)
func (_NovaVault *NovaVaultTransactor) Withdraw(opts *bind.TransactOpts, stable common.Address, shares *big.Int) (*types.Transaction, error) {
	return _NovaVault.contract.Transact(opts, "withdraw", stable, shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address stable, uint256 shares) returns(bool, uint256)
func (_NovaVault *NovaVaultSession) Withdraw(stable common.Address, shares *big.Int) (*types.Transaction, error) {
	return _NovaVault.Contract.Withdraw(&_NovaVault.TransactOpts, stable, shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address stable, uint256 shares) returns(bool, uint256)
func (_NovaVault *NovaVaultTransactorSession) Withdraw(stable common.Address, shares *big.Int) (*types.Transaction, error) {
	return _NovaVault.Contract.Withdraw(&_NovaVault.TransactOpts, stable, shares)
}

// NovaVaultAdapterApprovalIterator is returned from FilterAdapterApproval and is used to iterate over the raw logs and unpacked data for AdapterApproval events raised by the NovaVault contract.
type NovaVaultAdapterApprovalIterator struct {
	Event *NovaVaultAdapterApproval // Event containing the contract specifics and raw log

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
func (it *NovaVaultAdapterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NovaVaultAdapterApproval)
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
		it.Event = new(NovaVaultAdapterApproval)
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
func (it *NovaVaultAdapterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NovaVaultAdapterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NovaVaultAdapterApproval represents a AdapterApproval event raised by the NovaVault contract.
type NovaVaultAdapterApproval struct {
	Stable  common.Address
	Adapter common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdapterApproval is a free log retrieval operation binding the contract event 0x2a03353a3cafefd9bd1e60174ed856d38b75dc30ae89e43387bd92c4689a814b.
//
// Solidity: event AdapterApproval(address stable, address adapter)
func (_NovaVault *NovaVaultFilterer) FilterAdapterApproval(opts *bind.FilterOpts) (*NovaVaultAdapterApprovalIterator, error) {

	logs, sub, err := _NovaVault.contract.FilterLogs(opts, "AdapterApproval")
	if err != nil {
		return nil, err
	}
	return &NovaVaultAdapterApprovalIterator{contract: _NovaVault.contract, event: "AdapterApproval", logs: logs, sub: sub}, nil
}

// WatchAdapterApproval is a free log subscription operation binding the contract event 0x2a03353a3cafefd9bd1e60174ed856d38b75dc30ae89e43387bd92c4689a814b.
//
// Solidity: event AdapterApproval(address stable, address adapter)
func (_NovaVault *NovaVaultFilterer) WatchAdapterApproval(opts *bind.WatchOpts, sink chan<- *NovaVaultAdapterApproval) (event.Subscription, error) {

	logs, sub, err := _NovaVault.contract.WatchLogs(opts, "AdapterApproval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NovaVaultAdapterApproval)
				if err := _NovaVault.contract.UnpackLog(event, "AdapterApproval", log); err != nil {
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
func (_NovaVault *NovaVaultFilterer) ParseAdapterApproval(log types.Log) (*NovaVaultAdapterApproval, error) {
	event := new(NovaVaultAdapterApproval)
	if err := _NovaVault.contract.UnpackLog(event, "AdapterApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NovaVaultReferralIterator is returned from FilterReferral and is used to iterate over the raw logs and unpacked data for Referral events raised by the NovaVault contract.
type NovaVaultReferralIterator struct {
	Event *NovaVaultReferral // Event containing the contract specifics and raw log

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
func (it *NovaVaultReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NovaVaultReferral)
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
		it.Event = new(NovaVaultReferral)
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
func (it *NovaVaultReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NovaVaultReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NovaVaultReferral represents a Referral event raised by the NovaVault contract.
type NovaVaultReferral struct {
	Referral  uint16
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReferral is a free log retrieval operation binding the contract event 0x16902e34d01e8d5f80ce64939920a1390ff67f2546ae43ae72ac482033300968.
//
// Solidity: event Referral(uint16 referral, address indexed depositor, uint256 amount)
func (_NovaVault *NovaVaultFilterer) FilterReferral(opts *bind.FilterOpts, depositor []common.Address) (*NovaVaultReferralIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _NovaVault.contract.FilterLogs(opts, "Referral", depositorRule)
	if err != nil {
		return nil, err
	}
	return &NovaVaultReferralIterator{contract: _NovaVault.contract, event: "Referral", logs: logs, sub: sub}, nil
}

// WatchReferral is a free log subscription operation binding the contract event 0x16902e34d01e8d5f80ce64939920a1390ff67f2546ae43ae72ac482033300968.
//
// Solidity: event Referral(uint16 referral, address indexed depositor, uint256 amount)
func (_NovaVault *NovaVaultFilterer) WatchReferral(opts *bind.WatchOpts, sink chan<- *NovaVaultReferral, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _NovaVault.contract.WatchLogs(opts, "Referral", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NovaVaultReferral)
				if err := _NovaVault.contract.UnpackLog(event, "Referral", log); err != nil {
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
func (_NovaVault *NovaVaultFilterer) ParseReferral(log types.Log) (*NovaVaultReferral, error) {
	event := new(NovaVaultReferral)
	if err := _NovaVault.contract.UnpackLog(event, "Referral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
