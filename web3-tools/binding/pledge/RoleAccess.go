// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pledge

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

// RoleAccessMetaData contains all meta data concerning the RoleAccess contract.
var RoleAccessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[],\"name\":\"allAdmins\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"}]",
}

// RoleAccessABI is the input ABI used to generate the binding from.
// Deprecated: Use RoleAccessMetaData.ABI instead.
var RoleAccessABI = RoleAccessMetaData.ABI

// RoleAccess is an auto generated Go binding around an Ethereum contract.
type RoleAccess struct {
	RoleAccessCaller     // Read-only binding to the contract
	RoleAccessTransactor // Write-only binding to the contract
	RoleAccessFilterer   // Log filterer for contract events
}

// RoleAccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleAccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleAccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleAccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleAccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleAccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleAccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleAccessSession struct {
	Contract     *RoleAccess       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleAccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleAccessCallerSession struct {
	Contract *RoleAccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RoleAccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleAccessTransactorSession struct {
	Contract     *RoleAccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RoleAccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleAccessRaw struct {
	Contract *RoleAccess // Generic contract binding to access the raw methods on
}

// RoleAccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleAccessCallerRaw struct {
	Contract *RoleAccessCaller // Generic read-only contract binding to access the raw methods on
}

// RoleAccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleAccessTransactorRaw struct {
	Contract *RoleAccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoleAccess creates a new instance of RoleAccess, bound to a specific deployed contract.
func NewRoleAccess(address common.Address, backend bind.ContractBackend) (*RoleAccess, error) {
	contract, err := bindRoleAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoleAccess{RoleAccessCaller: RoleAccessCaller{contract: contract}, RoleAccessTransactor: RoleAccessTransactor{contract: contract}, RoleAccessFilterer: RoleAccessFilterer{contract: contract}}, nil
}

// NewRoleAccessCaller creates a new read-only instance of RoleAccess, bound to a specific deployed contract.
func NewRoleAccessCaller(address common.Address, caller bind.ContractCaller) (*RoleAccessCaller, error) {
	contract, err := bindRoleAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleAccessCaller{contract: contract}, nil
}

// NewRoleAccessTransactor creates a new write-only instance of RoleAccess, bound to a specific deployed contract.
func NewRoleAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleAccessTransactor, error) {
	contract, err := bindRoleAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleAccessTransactor{contract: contract}, nil
}

// NewRoleAccessFilterer creates a new log filterer instance of RoleAccess, bound to a specific deployed contract.
func NewRoleAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleAccessFilterer, error) {
	contract, err := bindRoleAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleAccessFilterer{contract: contract}, nil
}

// bindRoleAccess binds a generic wrapper to an already deployed contract.
func bindRoleAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RoleAccessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleAccess *RoleAccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleAccess.Contract.RoleAccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleAccess *RoleAccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleAccess.Contract.RoleAccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleAccess *RoleAccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleAccess.Contract.RoleAccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleAccess *RoleAccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleAccess.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleAccess *RoleAccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleAccess.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleAccess *RoleAccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleAccess.Contract.contract.Transact(opts, method, params...)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_RoleAccess *RoleAccessCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RoleAccess.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_RoleAccess *RoleAccessSession) AllAdmins() ([]common.Address, error) {
	return _RoleAccess.Contract.AllAdmins(&_RoleAccess.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_RoleAccess *RoleAccessCallerSession) AllAdmins() ([]common.Address, error) {
	return _RoleAccess.Contract.AllAdmins(&_RoleAccess.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_RoleAccess *RoleAccessCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _RoleAccess.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_RoleAccess *RoleAccessSession) IsAdmin(account common.Address) (bool, error) {
	return _RoleAccess.Contract.IsAdmin(&_RoleAccess.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_RoleAccess *RoleAccessCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _RoleAccess.Contract.IsAdmin(&_RoleAccess.CallOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_RoleAccess *RoleAccessTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RoleAccess.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_RoleAccess *RoleAccessSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _RoleAccess.Contract.AddAdmin(&_RoleAccess.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_RoleAccess *RoleAccessTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _RoleAccess.Contract.AddAdmin(&_RoleAccess.TransactOpts, account)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_RoleAccess *RoleAccessTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _RoleAccess.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_RoleAccess *RoleAccessSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _RoleAccess.Contract.BatchAddAdmin(&_RoleAccess.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_RoleAccess *RoleAccessTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _RoleAccess.Contract.BatchAddAdmin(&_RoleAccess.TransactOpts, amounts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_RoleAccess *RoleAccessTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RoleAccess.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_RoleAccess *RoleAccessSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _RoleAccess.Contract.RemoveAdmin(&_RoleAccess.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_RoleAccess *RoleAccessTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _RoleAccess.Contract.RemoveAdmin(&_RoleAccess.TransactOpts, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_RoleAccess *RoleAccessTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleAccess.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_RoleAccess *RoleAccessSession) RenounceAdmin() (*types.Transaction, error) {
	return _RoleAccess.Contract.RenounceAdmin(&_RoleAccess.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_RoleAccess *RoleAccessTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _RoleAccess.Contract.RenounceAdmin(&_RoleAccess.TransactOpts)
}

// RoleAccessAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the RoleAccess contract.
type RoleAccessAdminAddedIterator struct {
	Event *RoleAccessAdminAdded // Event containing the contract specifics and raw log

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
func (it *RoleAccessAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleAccessAdminAdded)
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
		it.Event = new(RoleAccessAdminAdded)
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
func (it *RoleAccessAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleAccessAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleAccessAdminAdded represents a AdminAdded event raised by the RoleAccess contract.
type RoleAccessAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address account)
func (_RoleAccess *RoleAccessFilterer) FilterAdminAdded(opts *bind.FilterOpts) (*RoleAccessAdminAddedIterator, error) {

	logs, sub, err := _RoleAccess.contract.FilterLogs(opts, "AdminAdded")
	if err != nil {
		return nil, err
	}
	return &RoleAccessAdminAddedIterator{contract: _RoleAccess.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address account)
func (_RoleAccess *RoleAccessFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *RoleAccessAdminAdded) (event.Subscription, error) {

	logs, sub, err := _RoleAccess.contract.WatchLogs(opts, "AdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleAccessAdminAdded)
				if err := _RoleAccess.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// ParseAdminAdded is a log parse operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address account)
func (_RoleAccess *RoleAccessFilterer) ParseAdminAdded(log types.Log) (*RoleAccessAdminAdded, error) {
	event := new(RoleAccessAdminAdded)
	if err := _RoleAccess.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleAccessAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the RoleAccess contract.
type RoleAccessAdminRemovedIterator struct {
	Event *RoleAccessAdminRemoved // Event containing the contract specifics and raw log

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
func (it *RoleAccessAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleAccessAdminRemoved)
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
		it.Event = new(RoleAccessAdminRemoved)
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
func (it *RoleAccessAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleAccessAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleAccessAdminRemoved represents a AdminRemoved event raised by the RoleAccess contract.
type RoleAccessAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address account)
func (_RoleAccess *RoleAccessFilterer) FilterAdminRemoved(opts *bind.FilterOpts) (*RoleAccessAdminRemovedIterator, error) {

	logs, sub, err := _RoleAccess.contract.FilterLogs(opts, "AdminRemoved")
	if err != nil {
		return nil, err
	}
	return &RoleAccessAdminRemovedIterator{contract: _RoleAccess.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address account)
func (_RoleAccess *RoleAccessFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *RoleAccessAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _RoleAccess.contract.WatchLogs(opts, "AdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleAccessAdminRemoved)
				if err := _RoleAccess.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address account)
func (_RoleAccess *RoleAccessFilterer) ParseAdminRemoved(log types.Log) (*RoleAccessAdminRemoved, error) {
	event := new(RoleAccessAdminRemoved)
	if err := _RoleAccess.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
