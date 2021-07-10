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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Id     common.Address
	Status uint8
}

// RelayABI is the input ABI used to generate the binding from.
const RelayABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIRotation.State\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"ChangeValidatorStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"InactiveRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NewValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RotationStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorToActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorToInactive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorsRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"enumIRotation.State\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIRotation.Validator\",\"name\":\"_validator\",\"type\":\"tuple\"}],\"name\":\"addInactiveValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"}],\"name\":\"addActiveValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executeRotation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemovedValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"enumIRotation.State\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIRotation.Validator[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"enumIRotation.State\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIRotation.Validator[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"},{\"internalType\":\"enumIRotation.State\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"}],\"name\":\"isInactive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeValidatorsRound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockRotation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_id\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Relay *RelayCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "DEFAULT_ADMIN_ROLE")
	return *ret0, err
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Relay *RelaySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Relay.Contract.DEFAULTADMINROLE(&_Relay.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Relay *RelayCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Relay.Contract.DEFAULTADMINROLE(&_Relay.CallOpts)
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
// Solidity: function getRemovedValidators() view returns((address,uint8)[])
func (_Relay *RelayCaller) GetRemovedValidators(opts *bind.CallOpts) ([]Struct0, error) {
	var (
		ret0 = new([]Struct0)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getRemovedValidators")
	return *ret0, err
}

// GetRemovedValidators is a free data retrieval call binding the contract method 0x8211cd54.
//
// Solidity: function getRemovedValidators() view returns((address,uint8)[])
func (_Relay *RelaySession) GetRemovedValidators() ([]Struct0, error) {
	return _Relay.Contract.GetRemovedValidators(&_Relay.CallOpts)
}

// GetRemovedValidators is a free data retrieval call binding the contract method 0x8211cd54.
//
// Solidity: function getRemovedValidators() view returns((address,uint8)[])
func (_Relay *RelayCallerSession) GetRemovedValidators() ([]Struct0, error) {
	return _Relay.Contract.GetRemovedValidators(&_Relay.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Relay *RelayCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getRoleAdmin", role)
	return *ret0, err
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Relay *RelaySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Relay.Contract.GetRoleAdmin(&_Relay.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Relay *RelayCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Relay.Contract.GetRoleAdmin(&_Relay.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Relay *RelayCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getRoleMember", role, index)
	return *ret0, err
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Relay *RelaySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Relay.Contract.GetRoleMember(&_Relay.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Relay *RelayCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Relay.Contract.GetRoleMember(&_Relay.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Relay *RelayCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getRoleMemberCount", role)
	return *ret0, err
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Relay *RelaySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Relay.Contract.GetRoleMemberCount(&_Relay.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Relay *RelayCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Relay.Contract.GetRoleMemberCount(&_Relay.CallOpts, role)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint8)[])
func (_Relay *RelayCaller) GetValidators(opts *bind.CallOpts) ([]Struct0, error) {
	var (
		ret0 = new([]Struct0)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint8)[])
func (_Relay *RelaySession) GetValidators() ([]Struct0, error) {
	return _Relay.Contract.GetValidators(&_Relay.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint8)[])
func (_Relay *RelayCallerSession) GetValidators() ([]Struct0, error) {
	return _Relay.Contract.GetValidators(&_Relay.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Relay *RelayCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "hasRole", role, account)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Relay *RelaySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Relay.Contract.HasRole(&_Relay.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Relay *RelayCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Relay.Contract.HasRole(&_Relay.CallOpts, role, account)
}

// IsInactive is a free data retrieval call binding the contract method 0x5fcbc896.
//
// Solidity: function isInactive(address _id) view returns(bool)
func (_Relay *RelayCaller) IsInactive(opts *bind.CallOpts, _id common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "isInactive", _id)
	return *ret0, err
}

// IsInactive is a free data retrieval call binding the contract method 0x5fcbc896.
//
// Solidity: function isInactive(address _id) view returns(bool)
func (_Relay *RelaySession) IsInactive(_id common.Address) (bool, error) {
	return _Relay.Contract.IsInactive(&_Relay.CallOpts, _id)
}

// IsInactive is a free data retrieval call binding the contract method 0x5fcbc896.
//
// Solidity: function isInactive(address _id) view returns(bool)
func (_Relay *RelayCallerSession) IsInactive(_id common.Address) (bool, error) {
	return _Relay.Contract.IsInactive(&_Relay.CallOpts, _id)
}

// AddActiveValidator is a paid mutator transaction binding the contract method 0xa885a232.
//
// Solidity: function addActiveValidator(address _id) returns()
func (_Relay *RelayTransactor) AddActiveValidator(opts *bind.TransactOpts, _id common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "addActiveValidator", _id)
}

// AddActiveValidator is a paid mutator transaction binding the contract method 0xa885a232.
//
// Solidity: function addActiveValidator(address _id) returns()
func (_Relay *RelaySession) AddActiveValidator(_id common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddActiveValidator(&_Relay.TransactOpts, _id)
}

// AddActiveValidator is a paid mutator transaction binding the contract method 0xa885a232.
//
// Solidity: function addActiveValidator(address _id) returns()
func (_Relay *RelayTransactorSession) AddActiveValidator(_id common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddActiveValidator(&_Relay.TransactOpts, _id)
}

// AddInactiveValidator is a paid mutator transaction binding the contract method 0xd74675a8.
//
// Solidity: function addInactiveValidator((address,uint8) _validator) returns()
func (_Relay *RelayTransactor) AddInactiveValidator(opts *bind.TransactOpts, _validator Struct0) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "addInactiveValidator", _validator)
}

// AddInactiveValidator is a paid mutator transaction binding the contract method 0xd74675a8.
//
// Solidity: function addInactiveValidator((address,uint8) _validator) returns()
func (_Relay *RelaySession) AddInactiveValidator(_validator Struct0) (*types.Transaction, error) {
	return _Relay.Contract.AddInactiveValidator(&_Relay.TransactOpts, _validator)
}

// AddInactiveValidator is a paid mutator transaction binding the contract method 0xd74675a8.
//
// Solidity: function addInactiveValidator((address,uint8) _validator) returns()
func (_Relay *RelayTransactorSession) AddInactiveValidator(_validator Struct0) (*types.Transaction, error) {
	return _Relay.Contract.AddInactiveValidator(&_Relay.TransactOpts, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _id) returns(bool)
func (_Relay *RelayTransactor) AddValidator(opts *bind.TransactOpts, _id common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "addValidator", _id)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _id) returns(bool)
func (_Relay *RelaySession) AddValidator(_id common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddValidator(&_Relay.TransactOpts, _id)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _id) returns(bool)
func (_Relay *RelayTransactorSession) AddValidator(_id common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddValidator(&_Relay.TransactOpts, _id)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Relay *RelaySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.GrantRole(&_Relay.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.GrantRole(&_Relay.TransactOpts, role, account)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _id) returns(bool)
func (_Relay *RelayTransactor) RemoveValidator(opts *bind.TransactOpts, _id common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "removeValidator", _id)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _id) returns(bool)
func (_Relay *RelaySession) RemoveValidator(_id common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RemoveValidator(&_Relay.TransactOpts, _id)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _id) returns(bool)
func (_Relay *RelayTransactorSession) RemoveValidator(_id common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RemoveValidator(&_Relay.TransactOpts, _id)
}

// RemoveValidatorsRound is a paid mutator transaction binding the contract method 0xe97b11a5.
//
// Solidity: function removeValidatorsRound() returns(bool)
func (_Relay *RelayTransactor) RemoveValidatorsRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "removeValidatorsRound")
}

// RemoveValidatorsRound is a paid mutator transaction binding the contract method 0xe97b11a5.
//
// Solidity: function removeValidatorsRound() returns(bool)
func (_Relay *RelaySession) RemoveValidatorsRound() (*types.Transaction, error) {
	return _Relay.Contract.RemoveValidatorsRound(&_Relay.TransactOpts)
}

// RemoveValidatorsRound is a paid mutator transaction binding the contract method 0xe97b11a5.
//
// Solidity: function removeValidatorsRound() returns(bool)
func (_Relay *RelayTransactorSession) RemoveValidatorsRound() (*types.Transaction, error) {
	return _Relay.Contract.RemoveValidatorsRound(&_Relay.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Relay *RelaySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RenounceRole(&_Relay.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RenounceRole(&_Relay.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Relay *RelaySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RevokeRole(&_Relay.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RevokeRole(&_Relay.TransactOpts, role, account)
}

// SetStatus is a paid mutator transaction binding the contract method 0x278e07ce.
//
// Solidity: function setStatus(address _id, uint8 _status) returns(bool)
func (_Relay *RelayTransactor) SetStatus(opts *bind.TransactOpts, _id common.Address, _status uint8) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "setStatus", _id, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x278e07ce.
//
// Solidity: function setStatus(address _id, uint8 _status) returns(bool)
func (_Relay *RelaySession) SetStatus(_id common.Address, _status uint8) (*types.Transaction, error) {
	return _Relay.Contract.SetStatus(&_Relay.TransactOpts, _id, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x278e07ce.
//
// Solidity: function setStatus(address _id, uint8 _status) returns(bool)
func (_Relay *RelayTransactorSession) SetStatus(_id common.Address, _status uint8) (*types.Transaction, error) {
	return _Relay.Contract.SetStatus(&_Relay.TransactOpts, _id, _status)
}

// RelayChangeValidatorStatusIterator is returned from FilterChangeValidatorStatus and is used to iterate over the raw logs and unpacked data for ChangeValidatorStatus events raised by the Relay contract.
type RelayChangeValidatorStatusIterator struct {
	Event *RelayChangeValidatorStatus // Event containing the contract specifics and raw log

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
func (it *RelayChangeValidatorStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayChangeValidatorStatus)
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
		it.Event = new(RelayChangeValidatorStatus)
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
func (it *RelayChangeValidatorStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayChangeValidatorStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayChangeValidatorStatus represents a ChangeValidatorStatus event raised by the Relay contract.
type RelayChangeValidatorStatus struct {
	Arg0 common.Address
	Arg1 uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChangeValidatorStatus is a free log retrieval operation binding the contract event 0xbc42ff7934f72b78426fd37c740647e45892d6f0b9449bb07019a916dcb59eb9.
//
// Solidity: event ChangeValidatorStatus(address arg0, uint8 arg1)
func (_Relay *RelayFilterer) FilterChangeValidatorStatus(opts *bind.FilterOpts) (*RelayChangeValidatorStatusIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "ChangeValidatorStatus")
	if err != nil {
		return nil, err
	}
	return &RelayChangeValidatorStatusIterator{contract: _Relay.contract, event: "ChangeValidatorStatus", logs: logs, sub: sub}, nil
}

// WatchChangeValidatorStatus is a free log subscription operation binding the contract event 0xbc42ff7934f72b78426fd37c740647e45892d6f0b9449bb07019a916dcb59eb9.
//
// Solidity: event ChangeValidatorStatus(address arg0, uint8 arg1)
func (_Relay *RelayFilterer) WatchChangeValidatorStatus(opts *bind.WatchOpts, sink chan<- *RelayChangeValidatorStatus) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "ChangeValidatorStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayChangeValidatorStatus)
				if err := _Relay.contract.UnpackLog(event, "ChangeValidatorStatus", log); err != nil {
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

// ParseChangeValidatorStatus is a log parse operation binding the contract event 0xbc42ff7934f72b78426fd37c740647e45892d6f0b9449bb07019a916dcb59eb9.
//
// Solidity: event ChangeValidatorStatus(address arg0, uint8 arg1)
func (_Relay *RelayFilterer) ParseChangeValidatorStatus(log types.Log) (*RelayChangeValidatorStatus, error) {
	event := new(RelayChangeValidatorStatus)
	if err := _Relay.contract.UnpackLog(event, "ChangeValidatorStatus", log); err != nil {
		return nil, err
	}
	return event, nil
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

// RelayNewValidatorAddedIterator is returned from FilterNewValidatorAdded and is used to iterate over the raw logs and unpacked data for NewValidatorAdded events raised by the Relay contract.
type RelayNewValidatorAddedIterator struct {
	Event *RelayNewValidatorAdded // Event containing the contract specifics and raw log

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
func (it *RelayNewValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayNewValidatorAdded)
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
		it.Event = new(RelayNewValidatorAdded)
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
func (it *RelayNewValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayNewValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayNewValidatorAdded represents a NewValidatorAdded event raised by the Relay contract.
type RelayNewValidatorAdded struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewValidatorAdded is a free log retrieval operation binding the contract event 0x6245c36f74f343491cc1bc8543e19c698cc3f8954c805a119182eaf321069db0.
//
// Solidity: event NewValidatorAdded(address arg0)
func (_Relay *RelayFilterer) FilterNewValidatorAdded(opts *bind.FilterOpts) (*RelayNewValidatorAddedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "NewValidatorAdded")
	if err != nil {
		return nil, err
	}
	return &RelayNewValidatorAddedIterator{contract: _Relay.contract, event: "NewValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchNewValidatorAdded is a free log subscription operation binding the contract event 0x6245c36f74f343491cc1bc8543e19c698cc3f8954c805a119182eaf321069db0.
//
// Solidity: event NewValidatorAdded(address arg0)
func (_Relay *RelayFilterer) WatchNewValidatorAdded(opts *bind.WatchOpts, sink chan<- *RelayNewValidatorAdded) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "NewValidatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayNewValidatorAdded)
				if err := _Relay.contract.UnpackLog(event, "NewValidatorAdded", log); err != nil {
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

// ParseNewValidatorAdded is a log parse operation binding the contract event 0x6245c36f74f343491cc1bc8543e19c698cc3f8954c805a119182eaf321069db0.
//
// Solidity: event NewValidatorAdded(address arg0)
func (_Relay *RelayFilterer) ParseNewValidatorAdded(log types.Log) (*RelayNewValidatorAdded, error) {
	event := new(RelayNewValidatorAdded)
	if err := _Relay.contract.UnpackLog(event, "NewValidatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Relay contract.
type RelayRoleAdminChangedIterator struct {
	Event *RelayRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RelayRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRoleAdminChanged)
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
		it.Event = new(RelayRoleAdminChanged)
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
func (it *RelayRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRoleAdminChanged represents a RoleAdminChanged event raised by the Relay contract.
type RelayRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Relay *RelayFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RelayRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Relay.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RelayRoleAdminChangedIterator{contract: _Relay.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Relay *RelayFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RelayRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Relay.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRoleAdminChanged)
				if err := _Relay.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Relay *RelayFilterer) ParseRoleAdminChanged(log types.Log) (*RelayRoleAdminChanged, error) {
	event := new(RelayRoleAdminChanged)
	if err := _Relay.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Relay contract.
type RelayRoleGrantedIterator struct {
	Event *RelayRoleGranted // Event containing the contract specifics and raw log

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
func (it *RelayRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRoleGranted)
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
		it.Event = new(RelayRoleGranted)
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
func (it *RelayRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRoleGranted represents a RoleGranted event raised by the Relay contract.
type RelayRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Relay *RelayFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RelayRoleGrantedIterator, error) {

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

	logs, sub, err := _Relay.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RelayRoleGrantedIterator{contract: _Relay.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Relay *RelayFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RelayRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Relay.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRoleGranted)
				if err := _Relay.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Relay *RelayFilterer) ParseRoleGranted(log types.Log) (*RelayRoleGranted, error) {
	event := new(RelayRoleGranted)
	if err := _Relay.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Relay contract.
type RelayRoleRevokedIterator struct {
	Event *RelayRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RelayRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRoleRevoked)
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
		it.Event = new(RelayRoleRevoked)
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
func (it *RelayRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRoleRevoked represents a RoleRevoked event raised by the Relay contract.
type RelayRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Relay *RelayFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RelayRoleRevokedIterator, error) {

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

	logs, sub, err := _Relay.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RelayRoleRevokedIterator{contract: _Relay.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Relay *RelayFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RelayRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Relay.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRoleRevoked)
				if err := _Relay.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Relay *RelayFilterer) ParseRoleRevoked(log types.Log) (*RelayRoleRevoked, error) {
	event := new(RelayRoleRevoked)
	if err := _Relay.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// RelayValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the Relay contract.
type RelayValidatorRemovedIterator struct {
	Event *RelayValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *RelayValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayValidatorRemoved)
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
		it.Event = new(RelayValidatorRemoved)
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
func (it *RelayValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayValidatorRemoved represents a ValidatorRemoved event raised by the Relay contract.
type RelayValidatorRemoved struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address arg0)
func (_Relay *RelayFilterer) FilterValidatorRemoved(opts *bind.FilterOpts) (*RelayValidatorRemovedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return &RelayValidatorRemovedIterator{contract: _Relay.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address arg0)
func (_Relay *RelayFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *RelayValidatorRemoved) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "ValidatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayValidatorRemoved)
				if err := _Relay.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address arg0)
func (_Relay *RelayFilterer) ParseValidatorRemoved(log types.Log) (*RelayValidatorRemoved, error) {
	event := new(RelayValidatorRemoved)
	if err := _Relay.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayValidatorToActiveIterator is returned from FilterValidatorToActive and is used to iterate over the raw logs and unpacked data for ValidatorToActive events raised by the Relay contract.
type RelayValidatorToActiveIterator struct {
	Event *RelayValidatorToActive // Event containing the contract specifics and raw log

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
func (it *RelayValidatorToActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayValidatorToActive)
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
		it.Event = new(RelayValidatorToActive)
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
func (it *RelayValidatorToActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayValidatorToActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayValidatorToActive represents a ValidatorToActive event raised by the Relay contract.
type RelayValidatorToActive struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorToActive is a free log retrieval operation binding the contract event 0x90c846181d4b5c8392b6c51c830ca30f1e53d95ab40de786e6bb32fc32dadd5e.
//
// Solidity: event ValidatorToActive(address validator)
func (_Relay *RelayFilterer) FilterValidatorToActive(opts *bind.FilterOpts) (*RelayValidatorToActiveIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "ValidatorToActive")
	if err != nil {
		return nil, err
	}
	return &RelayValidatorToActiveIterator{contract: _Relay.contract, event: "ValidatorToActive", logs: logs, sub: sub}, nil
}

// WatchValidatorToActive is a free log subscription operation binding the contract event 0x90c846181d4b5c8392b6c51c830ca30f1e53d95ab40de786e6bb32fc32dadd5e.
//
// Solidity: event ValidatorToActive(address validator)
func (_Relay *RelayFilterer) WatchValidatorToActive(opts *bind.WatchOpts, sink chan<- *RelayValidatorToActive) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "ValidatorToActive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayValidatorToActive)
				if err := _Relay.contract.UnpackLog(event, "ValidatorToActive", log); err != nil {
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

// ParseValidatorToActive is a log parse operation binding the contract event 0x90c846181d4b5c8392b6c51c830ca30f1e53d95ab40de786e6bb32fc32dadd5e.
//
// Solidity: event ValidatorToActive(address validator)
func (_Relay *RelayFilterer) ParseValidatorToActive(log types.Log) (*RelayValidatorToActive, error) {
	event := new(RelayValidatorToActive)
	if err := _Relay.contract.UnpackLog(event, "ValidatorToActive", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayValidatorToInactiveIterator is returned from FilterValidatorToInactive and is used to iterate over the raw logs and unpacked data for ValidatorToInactive events raised by the Relay contract.
type RelayValidatorToInactiveIterator struct {
	Event *RelayValidatorToInactive // Event containing the contract specifics and raw log

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
func (it *RelayValidatorToInactiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayValidatorToInactive)
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
		it.Event = new(RelayValidatorToInactive)
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
func (it *RelayValidatorToInactiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayValidatorToInactiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayValidatorToInactive represents a ValidatorToInactive event raised by the Relay contract.
type RelayValidatorToInactive struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorToInactive is a free log retrieval operation binding the contract event 0x2281c855c3591681b043b8cb763e2066e173fdf139dc00244bfddeeba4292250.
//
// Solidity: event ValidatorToInactive(address validator)
func (_Relay *RelayFilterer) FilterValidatorToInactive(opts *bind.FilterOpts) (*RelayValidatorToInactiveIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "ValidatorToInactive")
	if err != nil {
		return nil, err
	}
	return &RelayValidatorToInactiveIterator{contract: _Relay.contract, event: "ValidatorToInactive", logs: logs, sub: sub}, nil
}

// WatchValidatorToInactive is a free log subscription operation binding the contract event 0x2281c855c3591681b043b8cb763e2066e173fdf139dc00244bfddeeba4292250.
//
// Solidity: event ValidatorToInactive(address validator)
func (_Relay *RelayFilterer) WatchValidatorToInactive(opts *bind.WatchOpts, sink chan<- *RelayValidatorToInactive) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "ValidatorToInactive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayValidatorToInactive)
				if err := _Relay.contract.UnpackLog(event, "ValidatorToInactive", log); err != nil {
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

// ParseValidatorToInactive is a log parse operation binding the contract event 0x2281c855c3591681b043b8cb763e2066e173fdf139dc00244bfddeeba4292250.
//
// Solidity: event ValidatorToInactive(address validator)
func (_Relay *RelayFilterer) ParseValidatorToInactive(log types.Log) (*RelayValidatorToInactive, error) {
	event := new(RelayValidatorToInactive)
	if err := _Relay.contract.UnpackLog(event, "ValidatorToInactive", log); err != nil {
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
