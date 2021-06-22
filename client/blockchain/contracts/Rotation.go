// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relay

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RelayABI is the input ABI used to generate the binding from.
const RelayABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"InactiveRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RotationStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorsRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addNewValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executeRotation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemovedValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInactiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isInactive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockRotation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addInactiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeInactiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Relay is an auto generated Go binding around an Ethereum contract.
type Relay struct {
	RelayCaller     // Read-only binding to the contract
	RelayTransactor // Write-only binding to the contract
	RelayFilterer   // Log filterer for contract events
}

// RelayCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelaySession struct {
	Contract     *Relay            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayCallerSession struct {
	Contract *RelayCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RelayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayTransactorSession struct {
	Contract     *RelayTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayRaw struct {
	Contract *Relay // Generic contract binding to access the raw methods on
}

// RelayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayCallerRaw struct {
	Contract *RelayCaller // Generic read-only contract binding to access the raw methods on
}

// RelayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayTransactorRaw struct {
	Contract *RelayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelay creates a new instance of Relay, bound to a specific deployed contract.
func NewRelay(address common.Address, backend bind.ContractBackend) (*Relay, error) {
	contract, err := bindRelay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relay{RelayCaller: RelayCaller{contract: contract}, RelayTransactor: RelayTransactor{contract: contract}, RelayFilterer: RelayFilterer{contract: contract}}, nil
}

// NewRelayCaller creates a new read-only instance of Relay, bound to a specific deployed contract.
func NewRelayCaller(address common.Address, caller bind.ContractCaller) (*RelayCaller, error) {
	contract, err := bindRelay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayCaller{contract: contract}, nil
}

// NewRelayTransactor creates a new write-only instance of Relay, bound to a specific deployed contract.
func NewRelayTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayTransactor, error) {
	contract, err := bindRelay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayTransactor{contract: contract}, nil
}

// NewRelayFilterer creates a new log filterer instance of Relay, bound to a specific deployed contract.
func NewRelayFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayFilterer, error) {
	contract, err := bindRelay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayFilterer{contract: contract}, nil
}

// bindRelay binds a generic wrapper to an already deployed contract.
func bindRelay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relay *RelayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relay.Contract.RelayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relay *RelayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.Contract.RelayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relay *RelayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relay.Contract.RelayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relay *RelayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relay *RelayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relay *RelayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relay.Contract.contract.Transact(opts, method, params...)
}

// GetBlockRotation is a free data retrieval call binding the contract method 0x6388549a.
//
// Solidity: function getBlockRotation() view returns(uint256)
func (_Relay *RelayCaller) GetBlockRotation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getBlockRotation")
	return *ret0, err
}

// GetBlockRotation is a free data retrieval call binding the contract method 0x6388549a.
//
// Solidity: function getBlockRotation() view returns(uint256)
func (_Relay *RelaySession) GetBlockRotation() (*big.Int, error) {
	return _Relay.Contract.GetBlockRotation(&_Relay.CallOpts)
}

// GetBlockRotation is a free data retrieval call binding the contract method 0x6388549a.
//
// Solidity: function getBlockRotation() view returns(uint256)
func (_Relay *RelayCallerSession) GetBlockRotation() (*big.Int, error) {
	return _Relay.Contract.GetBlockRotation(&_Relay.CallOpts)
}

// GetInactiveValidators is a free data retrieval call binding the contract method 0x636ec582.
//
// Solidity: function getInactiveValidators() view returns(address[])
func (_Relay *RelayCaller) GetInactiveValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getInactiveValidators")
	return *ret0, err
}

// GetInactiveValidators is a free data retrieval call binding the contract method 0x636ec582.
//
// Solidity: function getInactiveValidators() view returns(address[])
func (_Relay *RelaySession) GetInactiveValidators() ([]common.Address, error) {
	return _Relay.Contract.GetInactiveValidators(&_Relay.CallOpts)
}

// GetInactiveValidators is a free data retrieval call binding the contract method 0x636ec582.
//
// Solidity: function getInactiveValidators() view returns(address[])
func (_Relay *RelayCallerSession) GetInactiveValidators() ([]common.Address, error) {
	return _Relay.Contract.GetInactiveValidators(&_Relay.CallOpts)
}

// GetNewValidators is a free data retrieval call binding the contract method 0xa3ee944a.
//
// Solidity: function getNewValidators() view returns(address[])
func (_Relay *RelayCaller) GetNewValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getNewValidators")
	return *ret0, err
}

// GetNewValidators is a free data retrieval call binding the contract method 0xa3ee944a.
//
// Solidity: function getNewValidators() view returns(address[])
func (_Relay *RelaySession) GetNewValidators() ([]common.Address, error) {
	return _Relay.Contract.GetNewValidators(&_Relay.CallOpts)
}

// GetNewValidators is a free data retrieval call binding the contract method 0xa3ee944a.
//
// Solidity: function getNewValidators() view returns(address[])
func (_Relay *RelayCallerSession) GetNewValidators() ([]common.Address, error) {
	return _Relay.Contract.GetNewValidators(&_Relay.CallOpts)
}

// GetRemovedValidators is a free data retrieval call binding the contract method 0x8211cd54.
//
// Solidity: function getRemovedValidators() view returns(address[])
func (_Relay *RelayCaller) GetRemovedValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getRemovedValidators")
	return *ret0, err
}

// GetRemovedValidators is a free data retrieval call binding the contract method 0x8211cd54.
//
// Solidity: function getRemovedValidators() view returns(address[])
func (_Relay *RelaySession) GetRemovedValidators() ([]common.Address, error) {
	return _Relay.Contract.GetRemovedValidators(&_Relay.CallOpts)
}

// GetRemovedValidators is a free data retrieval call binding the contract method 0x8211cd54.
//
// Solidity: function getRemovedValidators() view returns(address[])
func (_Relay *RelayCallerSession) GetRemovedValidators() ([]common.Address, error) {
	return _Relay.Contract.GetRemovedValidators(&_Relay.CallOpts)
}

// IsInactive is a free data retrieval call binding the contract method 0x5fcbc896.
//
// Solidity: function isInactive(address validator) view returns(bool)
func (_Relay *RelayCaller) IsInactive(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "isInactive", validator)
	return *ret0, err
}

// IsInactive is a free data retrieval call binding the contract method 0x5fcbc896.
//
// Solidity: function isInactive(address validator) view returns(bool)
func (_Relay *RelaySession) IsInactive(validator common.Address) (bool, error) {
	return _Relay.Contract.IsInactive(&_Relay.CallOpts, validator)
}

// IsInactive is a free data retrieval call binding the contract method 0x5fcbc896.
//
// Solidity: function isInactive(address validator) view returns(bool)
func (_Relay *RelayCallerSession) IsInactive(validator common.Address) (bool, error) {
	return _Relay.Contract.IsInactive(&_Relay.CallOpts, validator)
}

// AddInactiveValidator is a paid mutator transaction binding the contract method 0x6f4b62a6.
//
// Solidity: function addInactiveValidator(address validator) returns(bool)
func (_Relay *RelayTransactor) AddInactiveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "addInactiveValidator", validator)
}

// AddInactiveValidator is a paid mutator transaction binding the contract method 0x6f4b62a6.
//
// Solidity: function addInactiveValidator(address validator) returns(bool)
func (_Relay *RelaySession) AddInactiveValidator(validator common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddInactiveValidator(&_Relay.TransactOpts, validator)
}

// AddInactiveValidator is a paid mutator transaction binding the contract method 0x6f4b62a6.
//
// Solidity: function addInactiveValidator(address validator) returns(bool)
func (_Relay *RelayTransactorSession) AddInactiveValidator(validator common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddInactiveValidator(&_Relay.TransactOpts, validator)
}

// AddNewValidator is a paid mutator transaction binding the contract method 0xde2a0b38.
//
// Solidity: function addNewValidator(address validator) returns()
func (_Relay *RelayTransactor) AddNewValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "addNewValidator", validator)
}

// AddNewValidator is a paid mutator transaction binding the contract method 0xde2a0b38.
//
// Solidity: function addNewValidator(address validator) returns()
func (_Relay *RelaySession) AddNewValidator(validator common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddNewValidator(&_Relay.TransactOpts, validator)
}

// AddNewValidator is a paid mutator transaction binding the contract method 0xde2a0b38.
//
// Solidity: function addNewValidator(address validator) returns()
func (_Relay *RelayTransactorSession) AddNewValidator(validator common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddNewValidator(&_Relay.TransactOpts, validator)
}

// ExecuteRotation is a paid mutator transaction binding the contract method 0x8866982f.
//
// Solidity: function executeRotation() returns()
func (_Relay *RelayTransactor) ExecuteRotation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "executeRotation")
}

// ExecuteRotation is a paid mutator transaction binding the contract method 0x8866982f.
//
// Solidity: function executeRotation() returns()
func (_Relay *RelaySession) ExecuteRotation() (*types.Transaction, error) {
	return _Relay.Contract.ExecuteRotation(&_Relay.TransactOpts)
}

// ExecuteRotation is a paid mutator transaction binding the contract method 0x8866982f.
//
// Solidity: function executeRotation() returns()
func (_Relay *RelayTransactorSession) ExecuteRotation() (*types.Transaction, error) {
	return _Relay.Contract.ExecuteRotation(&_Relay.TransactOpts)
}

// RemoveInactiveValidator is a paid mutator transaction binding the contract method 0x7d8f5b22.
//
// Solidity: function removeInactiveValidator(address validator) returns(bool)
func (_Relay *RelayTransactor) RemoveInactiveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "removeInactiveValidator", validator)
}

// RemoveInactiveValidator is a paid mutator transaction binding the contract method 0x7d8f5b22.
//
// Solidity: function removeInactiveValidator(address validator) returns(bool)
func (_Relay *RelaySession) RemoveInactiveValidator(validator common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RemoveInactiveValidator(&_Relay.TransactOpts, validator)
}

// RemoveInactiveValidator is a paid mutator transaction binding the contract method 0x7d8f5b22.
//
// Solidity: function removeInactiveValidator(address validator) returns(bool)
func (_Relay *RelayTransactorSession) RemoveInactiveValidator(validator common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RemoveInactiveValidator(&_Relay.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_Relay *RelayTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_Relay *RelaySession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RemoveValidator(&_Relay.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_Relay *RelayTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RemoveValidator(&_Relay.TransactOpts, validator)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0xea179f4a.
//
// Solidity: function removeValidators() returns(bool)
func (_Relay *RelayTransactor) RemoveValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "removeValidators")
}

// RemoveValidators is a paid mutator transaction binding the contract method 0xea179f4a.
//
// Solidity: function removeValidators() returns(bool)
func (_Relay *RelaySession) RemoveValidators() (*types.Transaction, error) {
	return _Relay.Contract.RemoveValidators(&_Relay.TransactOpts)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0xea179f4a.
//
// Solidity: function removeValidators() returns(bool)
func (_Relay *RelayTransactorSession) RemoveValidators() (*types.Transaction, error) {
	return _Relay.Contract.RemoveValidators(&_Relay.TransactOpts)
}

// RelayInactiveRemovedIterator is returned from FilterInactiveRemoved and is used to iterate over the raw logs and unpacked data for InactiveRemoved events raised by the Relay contract.
type RelayInactiveRemovedIterator struct {
	Event *RelayInactiveRemoved // Event containing the contract specifics and raw log

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
func (it *RelayInactiveRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayInactiveRemoved)
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
		it.Event = new(RelayInactiveRemoved)
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
func (it *RelayInactiveRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayInactiveRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayInactiveRemoved represents a InactiveRemoved event raised by the Relay contract.
type RelayInactiveRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInactiveRemoved is a free log retrieval operation binding the contract event 0xdad573457756ca43591f41c85fd94695f7bdde556c99989995f1705725b641c1.
//
// Solidity: event InactiveRemoved(address validator)
func (_Relay *RelayFilterer) FilterInactiveRemoved(opts *bind.FilterOpts) (*RelayInactiveRemovedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "InactiveRemoved")
	if err != nil {
		return nil, err
	}
	return &RelayInactiveRemovedIterator{contract: _Relay.contract, event: "InactiveRemoved", logs: logs, sub: sub}, nil
}

// WatchInactiveRemoved is a free log subscription operation binding the contract event 0xdad573457756ca43591f41c85fd94695f7bdde556c99989995f1705725b641c1.
//
// Solidity: event InactiveRemoved(address validator)
func (_Relay *RelayFilterer) WatchInactiveRemoved(opts *bind.WatchOpts, sink chan<- *RelayInactiveRemoved) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "InactiveRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayInactiveRemoved)
				if err := _Relay.contract.UnpackLog(event, "InactiveRemoved", log); err != nil {
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

// ParseInactiveRemoved is a log parse operation binding the contract event 0xdad573457756ca43591f41c85fd94695f7bdde556c99989995f1705725b641c1.
//
// Solidity: event InactiveRemoved(address validator)
func (_Relay *RelayFilterer) ParseInactiveRemoved(log types.Log) (*RelayInactiveRemoved, error) {
	event := new(RelayInactiveRemoved)
	if err := _Relay.contract.UnpackLog(event, "InactiveRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayRotationStartedIterator is returned from FilterRotationStarted and is used to iterate over the raw logs and unpacked data for RotationStarted events raised by the Relay contract.
type RelayRotationStartedIterator struct {
	Event *RelayRotationStarted // Event containing the contract specifics and raw log

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
func (it *RelayRotationStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRotationStarted)
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
		it.Event = new(RelayRotationStarted)
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
func (it *RelayRotationStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRotationStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRotationStarted represents a RotationStarted event raised by the Relay contract.
type RelayRotationStarted struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRotationStarted is a free log retrieval operation binding the contract event 0xdb9b0d65d2a0d71eac0a89db8f448ac3e7f63fb21a8ef93d2dc84ce9b701057b.
//
// Solidity: event RotationStarted(uint256 blockNumber)
func (_Relay *RelayFilterer) FilterRotationStarted(opts *bind.FilterOpts) (*RelayRotationStartedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "RotationStarted")
	if err != nil {
		return nil, err
	}
	return &RelayRotationStartedIterator{contract: _Relay.contract, event: "RotationStarted", logs: logs, sub: sub}, nil
}

// WatchRotationStarted is a free log subscription operation binding the contract event 0xdb9b0d65d2a0d71eac0a89db8f448ac3e7f63fb21a8ef93d2dc84ce9b701057b.
//
// Solidity: event RotationStarted(uint256 blockNumber)
func (_Relay *RelayFilterer) WatchRotationStarted(opts *bind.WatchOpts, sink chan<- *RelayRotationStarted) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "RotationStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRotationStarted)
				if err := _Relay.contract.UnpackLog(event, "RotationStarted", log); err != nil {
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

// ParseRotationStarted is a log parse operation binding the contract event 0xdb9b0d65d2a0d71eac0a89db8f448ac3e7f63fb21a8ef93d2dc84ce9b701057b.
//
// Solidity: event RotationStarted(uint256 blockNumber)
func (_Relay *RelayFilterer) ParseRotationStarted(log types.Log) (*RelayRotationStarted, error) {
	event := new(RelayRotationStarted)
	if err := _Relay.contract.UnpackLog(event, "RotationStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayValidatorsRemovedIterator is returned from FilterValidatorsRemoved and is used to iterate over the raw logs and unpacked data for ValidatorsRemoved events raised by the Relay contract.
type RelayValidatorsRemovedIterator struct {
	Event *RelayValidatorsRemoved // Event containing the contract specifics and raw log

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
func (it *RelayValidatorsRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayValidatorsRemoved)
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
		it.Event = new(RelayValidatorsRemoved)
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
func (it *RelayValidatorsRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayValidatorsRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayValidatorsRemoved represents a ValidatorsRemoved event raised by the Relay contract.
type RelayValidatorsRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorsRemoved is a free log retrieval operation binding the contract event 0xe90a97bf186846f4ff7224144c9e11765229bf9f6ababdeb136bf7677584f14b.
//
// Solidity: event ValidatorsRemoved(address validator)
func (_Relay *RelayFilterer) FilterValidatorsRemoved(opts *bind.FilterOpts) (*RelayValidatorsRemovedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "ValidatorsRemoved")
	if err != nil {
		return nil, err
	}
	return &RelayValidatorsRemovedIterator{contract: _Relay.contract, event: "ValidatorsRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorsRemoved is a free log subscription operation binding the contract event 0xe90a97bf186846f4ff7224144c9e11765229bf9f6ababdeb136bf7677584f14b.
//
// Solidity: event ValidatorsRemoved(address validator)
func (_Relay *RelayFilterer) WatchValidatorsRemoved(opts *bind.WatchOpts, sink chan<- *RelayValidatorsRemoved) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "ValidatorsRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayValidatorsRemoved)
				if err := _Relay.contract.UnpackLog(event, "ValidatorsRemoved", log); err != nil {
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

// ParseValidatorsRemoved is a log parse operation binding the contract event 0xe90a97bf186846f4ff7224144c9e11765229bf9f6ababdeb136bf7677584f14b.
//
// Solidity: event ValidatorsRemoved(address validator)
func (_Relay *RelayFilterer) ParseValidatorsRemoved(log types.Log) (*RelayValidatorsRemoved, error) {
	event := new(RelayValidatorsRemoved)
	if err := _Relay.contract.UnpackLog(event, "ValidatorsRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
