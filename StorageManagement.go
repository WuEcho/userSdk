// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// DasKeySetInfo is an auto generated low-level Go binding around an user-defined struct.
type DasKeySetInfo struct {
	RequiredAmountOfSignatures *big.Int
	Addrs                      []common.Address
}

// StorageManagementMetaData contains all meta data concerning the StorageManagement contract.
var StorageManagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dasKey\",\"type\":\"bytes32\"}],\"name\":\"DasKeySetInfoRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"DASKEYSETINFO\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredAmountOfSignatures\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"internalType\":\"structDasKeySetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requiredAmountOfSignatures\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addrs\",\"type\":\"address[]\"}],\"name\":\"SetValidKeyset\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dasKeySetInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredAmountOfSignatures\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domiconNode\",\"outputs\":[{\"internalType\":\"contractDomiconNode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDomiconNode\",\"name\":\"_domiconNode\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StorageManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageManagementMetaData.ABI instead.
var StorageManagementABI = StorageManagementMetaData.ABI

// StorageManagement is an auto generated Go binding around an Ethereum contract.
type StorageManagement struct {
	StorageManagementCaller     // Read-only binding to the contract
	StorageManagementTransactor // Write-only binding to the contract
	StorageManagementFilterer   // Log filterer for contract events
}

// StorageManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageManagementSession struct {
	Contract     *StorageManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageManagementCallerSession struct {
	Contract *StorageManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StorageManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageManagementTransactorSession struct {
	Contract     *StorageManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StorageManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageManagementRaw struct {
	Contract *StorageManagement // Generic contract binding to access the raw methods on
}

// StorageManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageManagementCallerRaw struct {
	Contract *StorageManagementCaller // Generic read-only contract binding to access the raw methods on
}

// StorageManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageManagementTransactorRaw struct {
	Contract *StorageManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageManagement creates a new instance of StorageManagement, bound to a specific deployed contract.
func NewStorageManagement(address common.Address, backend bind.ContractBackend) (*StorageManagement, error) {
	contract, err := bindStorageManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageManagement{StorageManagementCaller: StorageManagementCaller{contract: contract}, StorageManagementTransactor: StorageManagementTransactor{contract: contract}, StorageManagementFilterer: StorageManagementFilterer{contract: contract}}, nil
}

// NewStorageManagementCaller creates a new read-only instance of StorageManagement, bound to a specific deployed contract.
func NewStorageManagementCaller(address common.Address, caller bind.ContractCaller) (*StorageManagementCaller, error) {
	contract, err := bindStorageManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageManagementCaller{contract: contract}, nil
}

// NewStorageManagementTransactor creates a new write-only instance of StorageManagement, bound to a specific deployed contract.
func NewStorageManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageManagementTransactor, error) {
	contract, err := bindStorageManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageManagementTransactor{contract: contract}, nil
}

// NewStorageManagementFilterer creates a new log filterer instance of StorageManagement, bound to a specific deployed contract.
func NewStorageManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageManagementFilterer, error) {
	contract, err := bindStorageManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageManagementFilterer{contract: contract}, nil
}

// bindStorageManagement binds a generic wrapper to an already deployed contract.
func bindStorageManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageManagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageManagement *StorageManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageManagement.Contract.StorageManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageManagement *StorageManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageManagement.Contract.StorageManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageManagement *StorageManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageManagement.Contract.StorageManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageManagement *StorageManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageManagement *StorageManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageManagement *StorageManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageManagement.Contract.contract.Transact(opts, method, params...)
}

// DASKEYSETINFO is a free data retrieval call binding the contract method 0x9d28029b.
//
// Solidity: function DASKEYSETINFO(bytes32 _key) view returns((uint256,address[]))
func (_StorageManagement *StorageManagementCaller) DASKEYSETINFO(opts *bind.CallOpts, _key [32]byte) (DasKeySetInfo, error) {
	var out []interface{}
	err := _StorageManagement.contract.Call(opts, &out, "DASKEYSETINFO", _key)

	if err != nil {
		return *new(DasKeySetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DasKeySetInfo)).(*DasKeySetInfo)

	return out0, err

}

// DASKEYSETINFO is a free data retrieval call binding the contract method 0x9d28029b.
//
// Solidity: function DASKEYSETINFO(bytes32 _key) view returns((uint256,address[]))
func (_StorageManagement *StorageManagementSession) DASKEYSETINFO(_key [32]byte) (DasKeySetInfo, error) {
	return _StorageManagement.Contract.DASKEYSETINFO(&_StorageManagement.CallOpts, _key)
}

// DASKEYSETINFO is a free data retrieval call binding the contract method 0x9d28029b.
//
// Solidity: function DASKEYSETINFO(bytes32 _key) view returns((uint256,address[]))
func (_StorageManagement *StorageManagementCallerSession) DASKEYSETINFO(_key [32]byte) (DasKeySetInfo, error) {
	return _StorageManagement.Contract.DASKEYSETINFO(&_StorageManagement.CallOpts, _key)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(uint256 requiredAmountOfSignatures)
func (_StorageManagement *StorageManagementCaller) DasKeySetInfo(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StorageManagement.contract.Call(opts, &out, "dasKeySetInfo", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(uint256 requiredAmountOfSignatures)
func (_StorageManagement *StorageManagementSession) DasKeySetInfo(arg0 [32]byte) (*big.Int, error) {
	return _StorageManagement.Contract.DasKeySetInfo(&_StorageManagement.CallOpts, arg0)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(uint256 requiredAmountOfSignatures)
func (_StorageManagement *StorageManagementCallerSession) DasKeySetInfo(arg0 [32]byte) (*big.Int, error) {
	return _StorageManagement.Contract.DasKeySetInfo(&_StorageManagement.CallOpts, arg0)
}

// DomiconNode is a free data retrieval call binding the contract method 0x5fa4ad36.
//
// Solidity: function domiconNode() view returns(address)
func (_StorageManagement *StorageManagementCaller) DomiconNode(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StorageManagement.contract.Call(opts, &out, "domiconNode")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DomiconNode is a free data retrieval call binding the contract method 0x5fa4ad36.
//
// Solidity: function domiconNode() view returns(address)
func (_StorageManagement *StorageManagementSession) DomiconNode() (common.Address, error) {
	return _StorageManagement.Contract.DomiconNode(&_StorageManagement.CallOpts)
}

// DomiconNode is a free data retrieval call binding the contract method 0x5fa4ad36.
//
// Solidity: function domiconNode() view returns(address)
func (_StorageManagement *StorageManagementCallerSession) DomiconNode() (common.Address, error) {
	return _StorageManagement.Contract.DomiconNode(&_StorageManagement.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StorageManagement *StorageManagementCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StorageManagement.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StorageManagement *StorageManagementSession) Version() (string, error) {
	return _StorageManagement.Contract.Version(&_StorageManagement.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StorageManagement *StorageManagementCallerSession) Version() (string, error) {
	return _StorageManagement.Contract.Version(&_StorageManagement.CallOpts)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0x1304232a.
//
// Solidity: function SetValidKeyset(uint256 _requiredAmountOfSignatures, address[] _addrs) returns(bytes32 ksHash)
func (_StorageManagement *StorageManagementTransactor) SetValidKeyset(opts *bind.TransactOpts, _requiredAmountOfSignatures *big.Int, _addrs []common.Address) (*types.Transaction, error) {
	return _StorageManagement.contract.Transact(opts, "SetValidKeyset", _requiredAmountOfSignatures, _addrs)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0x1304232a.
//
// Solidity: function SetValidKeyset(uint256 _requiredAmountOfSignatures, address[] _addrs) returns(bytes32 ksHash)
func (_StorageManagement *StorageManagementSession) SetValidKeyset(_requiredAmountOfSignatures *big.Int, _addrs []common.Address) (*types.Transaction, error) {
	return _StorageManagement.Contract.SetValidKeyset(&_StorageManagement.TransactOpts, _requiredAmountOfSignatures, _addrs)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0x1304232a.
//
// Solidity: function SetValidKeyset(uint256 _requiredAmountOfSignatures, address[] _addrs) returns(bytes32 ksHash)
func (_StorageManagement *StorageManagementTransactorSession) SetValidKeyset(_requiredAmountOfSignatures *big.Int, _addrs []common.Address) (*types.Transaction, error) {
	return _StorageManagement.Contract.SetValidKeyset(&_StorageManagement.TransactOpts, _requiredAmountOfSignatures, _addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _domiconNode) returns()
func (_StorageManagement *StorageManagementTransactor) Initialize(opts *bind.TransactOpts, _domiconNode common.Address) (*types.Transaction, error) {
	return _StorageManagement.contract.Transact(opts, "initialize", _domiconNode)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _domiconNode) returns()
func (_StorageManagement *StorageManagementSession) Initialize(_domiconNode common.Address) (*types.Transaction, error) {
	return _StorageManagement.Contract.Initialize(&_StorageManagement.TransactOpts, _domiconNode)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _domiconNode) returns()
func (_StorageManagement *StorageManagementTransactorSession) Initialize(_domiconNode common.Address) (*types.Transaction, error) {
	return _StorageManagement.Contract.Initialize(&_StorageManagement.TransactOpts, _domiconNode)
}

// StorageManagementDasKeySetInfoRegisteredIterator is returned from FilterDasKeySetInfoRegistered and is used to iterate over the raw logs and unpacked data for DasKeySetInfoRegistered events raised by the StorageManagement contract.
type StorageManagementDasKeySetInfoRegisteredIterator struct {
	Event *StorageManagementDasKeySetInfoRegistered // Event containing the contract specifics and raw log

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
func (it *StorageManagementDasKeySetInfoRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageManagementDasKeySetInfoRegistered)
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
		it.Event = new(StorageManagementDasKeySetInfoRegistered)
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
func (it *StorageManagementDasKeySetInfoRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageManagementDasKeySetInfoRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageManagementDasKeySetInfoRegistered represents a DasKeySetInfoRegistered event raised by the StorageManagement contract.
type StorageManagementDasKeySetInfoRegistered struct {
	User   common.Address
	DasKey [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDasKeySetInfoRegistered is a free log retrieval operation binding the contract event 0xa8ac0184865dbe2726c1e3b32895afd3de3e76e132fb84911e48ed5f13fa6440.
//
// Solidity: event DasKeySetInfoRegistered(address indexed user, bytes32 dasKey)
func (_StorageManagement *StorageManagementFilterer) FilterDasKeySetInfoRegistered(opts *bind.FilterOpts, user []common.Address) (*StorageManagementDasKeySetInfoRegisteredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StorageManagement.contract.FilterLogs(opts, "DasKeySetInfoRegistered", userRule)
	if err != nil {
		return nil, err
	}
	return &StorageManagementDasKeySetInfoRegisteredIterator{contract: _StorageManagement.contract, event: "DasKeySetInfoRegistered", logs: logs, sub: sub}, nil
}

// WatchDasKeySetInfoRegistered is a free log subscription operation binding the contract event 0xa8ac0184865dbe2726c1e3b32895afd3de3e76e132fb84911e48ed5f13fa6440.
//
// Solidity: event DasKeySetInfoRegistered(address indexed user, bytes32 dasKey)
func (_StorageManagement *StorageManagementFilterer) WatchDasKeySetInfoRegistered(opts *bind.WatchOpts, sink chan<- *StorageManagementDasKeySetInfoRegistered, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StorageManagement.contract.WatchLogs(opts, "DasKeySetInfoRegistered", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageManagementDasKeySetInfoRegistered)
				if err := _StorageManagement.contract.UnpackLog(event, "DasKeySetInfoRegistered", log); err != nil {
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

// ParseDasKeySetInfoRegistered is a log parse operation binding the contract event 0xa8ac0184865dbe2726c1e3b32895afd3de3e76e132fb84911e48ed5f13fa6440.
//
// Solidity: event DasKeySetInfoRegistered(address indexed user, bytes32 dasKey)
func (_StorageManagement *StorageManagementFilterer) ParseDasKeySetInfoRegistered(log types.Log) (*StorageManagementDasKeySetInfoRegistered, error) {
	event := new(StorageManagementDasKeySetInfoRegistered)
	if err := _StorageManagement.contract.UnpackLog(event, "DasKeySetInfoRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageManagementInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StorageManagement contract.
type StorageManagementInitializedIterator struct {
	Event *StorageManagementInitialized // Event containing the contract specifics and raw log

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
func (it *StorageManagementInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageManagementInitialized)
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
		it.Event = new(StorageManagementInitialized)
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
func (it *StorageManagementInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageManagementInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageManagementInitialized represents a Initialized event raised by the StorageManagement contract.
type StorageManagementInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StorageManagement *StorageManagementFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageManagementInitializedIterator, error) {

	logs, sub, err := _StorageManagement.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageManagementInitializedIterator{contract: _StorageManagement.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StorageManagement *StorageManagementFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageManagementInitialized) (event.Subscription, error) {

	logs, sub, err := _StorageManagement.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageManagementInitialized)
				if err := _StorageManagement.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StorageManagement *StorageManagementFilterer) ParseInitialized(log types.Log) (*StorageManagementInitialized, error) {
	event := new(StorageManagementInitialized)
	if err := _StorageManagement.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
