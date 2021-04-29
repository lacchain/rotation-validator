// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RotationStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ValidatorsRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addOldValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addNewValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executeRotation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOldValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockRotation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetBlockRotation is a free data retrieval call binding the contract method 0x6388549a.
//
// Solidity: function getBlockRotation() view returns(uint256)
func (_Contract *ContractCaller) GetBlockRotation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getBlockRotation")
	return *ret0, err
}

// GetBlockRotation is a free data retrieval call binding the contract method 0x6388549a.
//
// Solidity: function getBlockRotation() view returns(uint256)
func (_Contract *ContractSession) GetBlockRotation() (*big.Int, error) {
	return _Contract.Contract.GetBlockRotation(&_Contract.CallOpts)
}

// GetBlockRotation is a free data retrieval call binding the contract method 0x6388549a.
//
// Solidity: function getBlockRotation() view returns(uint256)
func (_Contract *ContractCallerSession) GetBlockRotation() (*big.Int, error) {
	return _Contract.Contract.GetBlockRotation(&_Contract.CallOpts)
}

// GetNewValidators is a free data retrieval call binding the contract method 0xa3ee944a.
//
// Solidity: function getNewValidators() view returns(address[])
func (_Contract *ContractCaller) GetNewValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getNewValidators")
	return *ret0, err
}

// GetNewValidators is a free data retrieval call binding the contract method 0xa3ee944a.
//
// Solidity: function getNewValidators() view returns(address[])
func (_Contract *ContractSession) GetNewValidators() ([]common.Address, error) {
	return _Contract.Contract.GetNewValidators(&_Contract.CallOpts)
}

// GetNewValidators is a free data retrieval call binding the contract method 0xa3ee944a.
//
// Solidity: function getNewValidators() view returns(address[])
func (_Contract *ContractCallerSession) GetNewValidators() ([]common.Address, error) {
	return _Contract.Contract.GetNewValidators(&_Contract.CallOpts)
}

// GetOldValidators is a free data retrieval call binding the contract method 0xb74adcb7.
//
// Solidity: function getOldValidators() view returns(address[])
func (_Contract *ContractCaller) GetOldValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getOldValidators")
	return *ret0, err
}

// GetOldValidators is a free data retrieval call binding the contract method 0xb74adcb7.
//
// Solidity: function getOldValidators() view returns(address[])
func (_Contract *ContractSession) GetOldValidators() ([]common.Address, error) {
	return _Contract.Contract.GetOldValidators(&_Contract.CallOpts)
}

// GetOldValidators is a free data retrieval call binding the contract method 0xb74adcb7.
//
// Solidity: function getOldValidators() view returns(address[])
func (_Contract *ContractCallerSession) GetOldValidators() ([]common.Address, error) {
	return _Contract.Contract.GetOldValidators(&_Contract.CallOpts)
}

// AddNewValidator is a paid mutator transaction binding the contract method 0xde2a0b38.
//
// Solidity: function addNewValidator(address validator) returns()
func (_Contract *ContractTransactor) AddNewValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addNewValidator", validator)
}

// AddNewValidator is a paid mutator transaction binding the contract method 0xde2a0b38.
//
// Solidity: function addNewValidator(address validator) returns()
func (_Contract *ContractSession) AddNewValidator(validator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddNewValidator(&_Contract.TransactOpts, validator)
}

// AddNewValidator is a paid mutator transaction binding the contract method 0xde2a0b38.
//
// Solidity: function addNewValidator(address validator) returns()
func (_Contract *ContractTransactorSession) AddNewValidator(validator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddNewValidator(&_Contract.TransactOpts, validator)
}

// AddOldValidator is a paid mutator transaction binding the contract method 0xe6d7f036.
//
// Solidity: function addOldValidator(address validator) returns()
func (_Contract *ContractTransactor) AddOldValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addOldValidator", validator)
}

// AddOldValidator is a paid mutator transaction binding the contract method 0xe6d7f036.
//
// Solidity: function addOldValidator(address validator) returns()
func (_Contract *ContractSession) AddOldValidator(validator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddOldValidator(&_Contract.TransactOpts, validator)
}

// AddOldValidator is a paid mutator transaction binding the contract method 0xe6d7f036.
//
// Solidity: function addOldValidator(address validator) returns()
func (_Contract *ContractTransactorSession) AddOldValidator(validator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddOldValidator(&_Contract.TransactOpts, validator)
}

// ExecuteRotation is a paid mutator transaction binding the contract method 0x8866982f.
//
// Solidity: function executeRotation() returns()
func (_Contract *ContractTransactor) ExecuteRotation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "executeRotation")
}

// ExecuteRotation is a paid mutator transaction binding the contract method 0x8866982f.
//
// Solidity: function executeRotation() returns()
func (_Contract *ContractSession) ExecuteRotation() (*types.Transaction, error) {
	return _Contract.Contract.ExecuteRotation(&_Contract.TransactOpts)
}

// ExecuteRotation is a paid mutator transaction binding the contract method 0x8866982f.
//
// Solidity: function executeRotation() returns()
func (_Contract *ContractTransactorSession) ExecuteRotation() (*types.Transaction, error) {
	return _Contract.Contract.ExecuteRotation(&_Contract.TransactOpts)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0xea179f4a.
//
// Solidity: function removeValidators() returns(bool)
func (_Contract *ContractTransactor) RemoveValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeValidators")
}

// RemoveValidators is a paid mutator transaction binding the contract method 0xea179f4a.
//
// Solidity: function removeValidators() returns(bool)
func (_Contract *ContractSession) RemoveValidators() (*types.Transaction, error) {
	return _Contract.Contract.RemoveValidators(&_Contract.TransactOpts)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0xea179f4a.
//
// Solidity: function removeValidators() returns(bool)
func (_Contract *ContractTransactorSession) RemoveValidators() (*types.Transaction, error) {
	return _Contract.Contract.RemoveValidators(&_Contract.TransactOpts)
}

// ContractRotationStartedIterator is returned from FilterRotationStarted and is used to iterate over the raw logs and unpacked data for RotationStarted events raised by the Contract contract.
type ContractRotationStartedIterator struct {
	Event *ContractRotationStarted // Event containing the contract specifics and raw log

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
func (it *ContractRotationStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRotationStarted)
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
		it.Event = new(ContractRotationStarted)
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
func (it *ContractRotationStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRotationStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRotationStarted represents a RotationStarted event raised by the Contract contract.
type ContractRotationStarted struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRotationStarted is a free log retrieval operation binding the contract event 0xdb9b0d65d2a0d71eac0a89db8f448ac3e7f63fb21a8ef93d2dc84ce9b701057b.
//
// Solidity: event RotationStarted(uint256 blockNumber)
func (_Contract *ContractFilterer) FilterRotationStarted(opts *bind.FilterOpts) (*ContractRotationStartedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RotationStarted")
	if err != nil {
		return nil, err
	}
	return &ContractRotationStartedIterator{contract: _Contract.contract, event: "RotationStarted", logs: logs, sub: sub}, nil
}

// WatchRotationStarted is a free log subscription operation binding the contract event 0xdb9b0d65d2a0d71eac0a89db8f448ac3e7f63fb21a8ef93d2dc84ce9b701057b.
//
// Solidity: event RotationStarted(uint256 blockNumber)
func (_Contract *ContractFilterer) WatchRotationStarted(opts *bind.WatchOpts, sink chan<- *ContractRotationStarted) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RotationStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRotationStarted)
				if err := _Contract.contract.UnpackLog(event, "RotationStarted", log); err != nil {
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
func (_Contract *ContractFilterer) ParseRotationStarted(log types.Log) (*ContractRotationStarted, error) {
	event := new(ContractRotationStarted)
	if err := _Contract.contract.UnpackLog(event, "RotationStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractValidatorsRemovedIterator is returned from FilterValidatorsRemoved and is used to iterate over the raw logs and unpacked data for ValidatorsRemoved events raised by the Contract contract.
type ContractValidatorsRemovedIterator struct {
	Event *ContractValidatorsRemoved // Event containing the contract specifics and raw log

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
func (it *ContractValidatorsRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorsRemoved)
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
		it.Event = new(ContractValidatorsRemoved)
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
func (it *ContractValidatorsRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorsRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorsRemoved represents a ValidatorsRemoved event raised by the Contract contract.
type ContractValidatorsRemoved struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValidatorsRemoved is a free log retrieval operation binding the contract event 0xe90a97bf186846f4ff7224144c9e11765229bf9f6ababdeb136bf7677584f14b.
//
// Solidity: event ValidatorsRemoved(address sender)
func (_Contract *ContractFilterer) FilterValidatorsRemoved(opts *bind.FilterOpts) (*ContractValidatorsRemovedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorsRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractValidatorsRemovedIterator{contract: _Contract.contract, event: "ValidatorsRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorsRemoved is a free log subscription operation binding the contract event 0xe90a97bf186846f4ff7224144c9e11765229bf9f6ababdeb136bf7677584f14b.
//
// Solidity: event ValidatorsRemoved(address sender)
func (_Contract *ContractFilterer) WatchValidatorsRemoved(opts *bind.WatchOpts, sink chan<- *ContractValidatorsRemoved) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorsRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorsRemoved)
				if err := _Contract.contract.UnpackLog(event, "ValidatorsRemoved", log); err != nil {
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
// Solidity: event ValidatorsRemoved(address sender)
func (_Contract *ContractFilterer) ParseValidatorsRemoved(log types.Log) (*ContractValidatorsRemoved, error) {
	event := new(ContractValidatorsRemoved)
	if err := _Contract.contract.UnpackLog(event, "ValidatorsRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
