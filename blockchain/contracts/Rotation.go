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
const RelayABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RotationSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addOldValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addNewValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOldValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getNewValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBlock\",\"type\":\"uint256\"}],\"name\":\"setBlockRotation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// GetOldValidators is a free data retrieval call binding the contract method 0xb74adcb7.
//
// Solidity: function getOldValidators() view returns(address[])
func (_Relay *RelayCaller) GetOldValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getOldValidators")
	return *ret0, err
}

// GetOldValidators is a free data retrieval call binding the contract method 0xb74adcb7.
//
// Solidity: function getOldValidators() view returns(address[])
func (_Relay *RelaySession) GetOldValidators() ([]common.Address, error) {
	return _Relay.Contract.GetOldValidators(&_Relay.CallOpts)
}

// GetOldValidators is a free data retrieval call binding the contract method 0xb74adcb7.
//
// Solidity: function getOldValidators() view returns(address[])
func (_Relay *RelayCallerSession) GetOldValidators() ([]common.Address, error) {
	return _Relay.Contract.GetOldValidators(&_Relay.CallOpts)
}

// ValidatorIsActive is a free data retrieval call binding the contract method 0x92c65ecc.
//
// Solidity: function validatorIsActive(address validator) view returns(bool)
func (_Relay *RelayCaller) ValidatorIsActive(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "validatorIsActive", validator)
	return *ret0, err
}

// ValidatorIsActive is a free data retrieval call binding the contract method 0x92c65ecc.
//
// Solidity: function validatorIsActive(address validator) view returns(bool)
func (_Relay *RelaySession) ValidatorIsActive(validator common.Address) (bool, error) {
	return _Relay.Contract.ValidatorIsActive(&_Relay.CallOpts, validator)
}

// ValidatorIsActive is a free data retrieval call binding the contract method 0x92c65ecc.
//
// Solidity: function validatorIsActive(address validator) view returns(bool)
func (_Relay *RelayCallerSession) ValidatorIsActive(validator common.Address) (bool, error) {
	return _Relay.Contract.ValidatorIsActive(&_Relay.CallOpts, validator)
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

// AddOldValidator is a paid mutator transaction binding the contract method 0xe6d7f036.
//
// Solidity: function addOldValidator(address validator) returns()
func (_Relay *RelayTransactor) AddOldValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "addOldValidator", validator)
}

// AddOldValidator is a paid mutator transaction binding the contract method 0xe6d7f036.
//
// Solidity: function addOldValidator(address validator) returns()
func (_Relay *RelaySession) AddOldValidator(validator common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddOldValidator(&_Relay.TransactOpts, validator)
}

// AddOldValidator is a paid mutator transaction binding the contract method 0xe6d7f036.
//
// Solidity: function addOldValidator(address validator) returns()
func (_Relay *RelayTransactorSession) AddOldValidator(validator common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddOldValidator(&_Relay.TransactOpts, validator)
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

// SetBlockRotation is a paid mutator transaction binding the contract method 0x77d98b0a.
//
// Solidity: function setBlockRotation(uint256 newBlock) returns(bool)
func (_Relay *RelayTransactor) SetBlockRotation(opts *bind.TransactOpts, newBlock *big.Int) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "setBlockRotation", newBlock)
}

// SetBlockRotation is a paid mutator transaction binding the contract method 0x77d98b0a.
//
// Solidity: function setBlockRotation(uint256 newBlock) returns(bool)
func (_Relay *RelaySession) SetBlockRotation(newBlock *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.SetBlockRotation(&_Relay.TransactOpts, newBlock)
}

// SetBlockRotation is a paid mutator transaction binding the contract method 0x77d98b0a.
//
// Solidity: function setBlockRotation(uint256 newBlock) returns(bool)
func (_Relay *RelayTransactorSession) SetBlockRotation(newBlock *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.SetBlockRotation(&_Relay.TransactOpts, newBlock)
}

// RelayRotationSetIterator is returned from FilterRotationSet and is used to iterate over the raw logs and unpacked data for RotationSet events raised by the Relay contract.
type RelayRotationSetIterator struct {
	Event *RelayRotationSet // Event containing the contract specifics and raw log

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
func (it *RelayRotationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRotationSet)
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
		it.Event = new(RelayRotationSet)
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
func (it *RelayRotationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRotationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRotationSet represents a RotationSet event raised by the Relay contract.
type RelayRotationSet struct {
	Seter       common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRotationSet is a free log retrieval operation binding the contract event 0xcb8834c2494939a9e9ca97ab0bf01dfaf5baa4146efb9ebeb03f706c5d85a49e.
//
// Solidity: event RotationSet(address indexed seter, uint256 blockNumber)
func (_Relay *RelayFilterer) FilterRotationSet(opts *bind.FilterOpts, seter []common.Address) (*RelayRotationSetIterator, error) {

	var seterRule []interface{}
	for _, seterItem := range seter {
		seterRule = append(seterRule, seterItem)
	}

	logs, sub, err := _Relay.contract.FilterLogs(opts, "RotationSet", seterRule)
	if err != nil {
		return nil, err
	}
	return &RelayRotationSetIterator{contract: _Relay.contract, event: "RotationSet", logs: logs, sub: sub}, nil
}

// WatchRotationSet is a free log subscription operation binding the contract event 0xcb8834c2494939a9e9ca97ab0bf01dfaf5baa4146efb9ebeb03f706c5d85a49e.
//
// Solidity: event RotationSet(address indexed seter, uint256 blockNumber)
func (_Relay *RelayFilterer) WatchRotationSet(opts *bind.WatchOpts, sink chan<- *RelayRotationSet, seter []common.Address) (event.Subscription, error) {

	var seterRule []interface{}
	for _, seterItem := range seter {
		seterRule = append(seterRule, seterItem)
	}

	logs, sub, err := _Relay.contract.WatchLogs(opts, "RotationSet", seterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRotationSet)
				if err := _Relay.contract.UnpackLog(event, "RotationSet", log); err != nil {
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

// ParseRotationSet is a log parse operation binding the contract event 0xcb8834c2494939a9e9ca97ab0bf01dfaf5baa4146efb9ebeb03f706c5d85a49e.
//
// Solidity: event RotationSet(address indexed seter, uint256 blockNumber)
func (_Relay *RelayFilterer) ParseRotationSet(log types.Log) (*RelayRotationSet, error) {
	event := new(RelayRotationSet)
	if err := _Relay.contract.UnpackLog(event, "RotationSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
