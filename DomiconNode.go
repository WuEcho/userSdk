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

// DomiconNodeNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type DomiconNodeNodeInfo struct {
	Url             string
	Name            string
	StakedTokens    *big.Int
	Location        string
	MaxStorageSpace *big.Int
	Addr            common.Address
}

// DomicNodeMetaData contains all meta data concerning the DomicNode contract.
var DomicNodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedTokens\",\"type\":\"uint256\"}],\"name\":\"BroadcastNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedTokens\",\"type\":\"uint256\"}],\"name\":\"StorageNode\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"IsNodeBroadcast\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakedTokens\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxStorageSpace\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structDomiconNode.NodeInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"RegisterBroadcastNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakedTokens\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxStorageSpace\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structDomiconNode.NodeInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"RegisterStorageNodeList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetDom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"broadcastNodeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"broadcastingNodes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakedTokens\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxStorageSpace\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"storageNodeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"storageNodes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakedTokens\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxStorageSpace\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DomicNodeABI is the input ABI used to generate the binding from.
// Deprecated: Use DomicNodeMetaData.ABI instead.
var DomicNodeABI = DomicNodeMetaData.ABI

// DomicNode is an auto generated Go binding around an Ethereum contract.
type DomicNode struct {
	DomicNodeCaller     // Read-only binding to the contract
	DomicNodeTransactor // Write-only binding to the contract
	DomicNodeFilterer   // Log filterer for contract events
}

// DomicNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DomicNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomicNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DomicNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomicNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DomicNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomicNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DomicNodeSession struct {
	Contract     *DomicNode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DomicNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DomicNodeCallerSession struct {
	Contract *DomicNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DomicNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DomicNodeTransactorSession struct {
	Contract     *DomicNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DomicNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DomicNodeRaw struct {
	Contract *DomicNode // Generic contract binding to access the raw methods on
}

// DomicNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DomicNodeCallerRaw struct {
	Contract *DomicNodeCaller // Generic read-only contract binding to access the raw methods on
}

// DomicNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DomicNodeTransactorRaw struct {
	Contract *DomicNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDomicNode creates a new instance of DomicNode, bound to a specific deployed contract.
func NewDomicNode(address common.Address, backend bind.ContractBackend) (*DomicNode, error) {
	contract, err := bindDomicNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DomicNode{DomicNodeCaller: DomicNodeCaller{contract: contract}, DomicNodeTransactor: DomicNodeTransactor{contract: contract}, DomicNodeFilterer: DomicNodeFilterer{contract: contract}}, nil
}

// NewDomicNodeCaller creates a new read-only instance of DomicNode, bound to a specific deployed contract.
func NewDomicNodeCaller(address common.Address, caller bind.ContractCaller) (*DomicNodeCaller, error) {
	contract, err := bindDomicNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DomicNodeCaller{contract: contract}, nil
}

// NewDomicNodeTransactor creates a new write-only instance of DomicNode, bound to a specific deployed contract.
func NewDomicNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*DomicNodeTransactor, error) {
	contract, err := bindDomicNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DomicNodeTransactor{contract: contract}, nil
}

// NewDomicNodeFilterer creates a new log filterer instance of DomicNode, bound to a specific deployed contract.
func NewDomicNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*DomicNodeFilterer, error) {
	contract, err := bindDomicNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DomicNodeFilterer{contract: contract}, nil
}

// bindDomicNode binds a generic wrapper to an already deployed contract.
func bindDomicNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DomicNodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomicNode *DomicNodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomicNode.Contract.DomicNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomicNode *DomicNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomicNode.Contract.DomicNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomicNode *DomicNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomicNode.Contract.DomicNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomicNode *DomicNodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomicNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomicNode *DomicNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomicNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomicNode *DomicNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomicNode.Contract.contract.Transact(opts, method, params...)
}

// DOM is a free data retrieval call binding the contract method 0x6a57f6b1.
//
// Solidity: function DOM() view returns(address)
func (_DomicNode *DomicNodeCaller) DOM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DomicNode.contract.Call(opts, &out, "DOM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DOM is a free data retrieval call binding the contract method 0x6a57f6b1.
//
// Solidity: function DOM() view returns(address)
func (_DomicNode *DomicNodeSession) DOM() (common.Address, error) {
	return _DomicNode.Contract.DOM(&_DomicNode.CallOpts)
}

// DOM is a free data retrieval call binding the contract method 0x6a57f6b1.
//
// Solidity: function DOM() view returns(address)
func (_DomicNode *DomicNodeCallerSession) DOM() (common.Address, error) {
	return _DomicNode.Contract.DOM(&_DomicNode.CallOpts)
}

// IsNodeBroadcast is a free data retrieval call binding the contract method 0xc0f2acea.
//
// Solidity: function IsNodeBroadcast(address addr) view returns(bool)
func (_DomicNode *DomicNodeCaller) IsNodeBroadcast(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _DomicNode.contract.Call(opts, &out, "IsNodeBroadcast", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNodeBroadcast is a free data retrieval call binding the contract method 0xc0f2acea.
//
// Solidity: function IsNodeBroadcast(address addr) view returns(bool)
func (_DomicNode *DomicNodeSession) IsNodeBroadcast(addr common.Address) (bool, error) {
	return _DomicNode.Contract.IsNodeBroadcast(&_DomicNode.CallOpts, addr)
}

// IsNodeBroadcast is a free data retrieval call binding the contract method 0xc0f2acea.
//
// Solidity: function IsNodeBroadcast(address addr) view returns(bool)
func (_DomicNode *DomicNodeCallerSession) IsNodeBroadcast(addr common.Address) (bool, error) {
	return _DomicNode.Contract.IsNodeBroadcast(&_DomicNode.CallOpts, addr)
}

// BroadcastNodeList is a free data retrieval call binding the contract method 0x56f7be58.
//
// Solidity: function broadcastNodeList(uint256 ) view returns(address)
func (_DomicNode *DomicNodeCaller) BroadcastNodeList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DomicNode.contract.Call(opts, &out, "broadcastNodeList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BroadcastNodeList is a free data retrieval call binding the contract method 0x56f7be58.
//
// Solidity: function broadcastNodeList(uint256 ) view returns(address)
func (_DomicNode *DomicNodeSession) BroadcastNodeList(arg0 *big.Int) (common.Address, error) {
	return _DomicNode.Contract.BroadcastNodeList(&_DomicNode.CallOpts, arg0)
}

// BroadcastNodeList is a free data retrieval call binding the contract method 0x56f7be58.
//
// Solidity: function broadcastNodeList(uint256 ) view returns(address)
func (_DomicNode *DomicNodeCallerSession) BroadcastNodeList(arg0 *big.Int) (common.Address, error) {
	return _DomicNode.Contract.BroadcastNodeList(&_DomicNode.CallOpts, arg0)
}

// BroadcastingNodes is a free data retrieval call binding the contract method 0xcac5800a.
//
// Solidity: function broadcastingNodes(address ) view returns(string url, string name, uint256 stakedTokens, string location, uint256 maxStorageSpace, address addr)
func (_DomicNode *DomicNodeCaller) BroadcastingNodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Url             string
	Name            string
	StakedTokens    *big.Int
	Location        string
	MaxStorageSpace *big.Int
	Addr            common.Address
}, error) {
	var out []interface{}
	err := _DomicNode.contract.Call(opts, &out, "broadcastingNodes", arg0)

	outstruct := new(struct {
		Url             string
		Name            string
		StakedTokens    *big.Int
		Location        string
		MaxStorageSpace *big.Int
		Addr            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Url = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.StakedTokens = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Location = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.MaxStorageSpace = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Addr = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BroadcastingNodes is a free data retrieval call binding the contract method 0xcac5800a.
//
// Solidity: function broadcastingNodes(address ) view returns(string url, string name, uint256 stakedTokens, string location, uint256 maxStorageSpace, address addr)
func (_DomicNode *DomicNodeSession) BroadcastingNodes(arg0 common.Address) (struct {
	Url             string
	Name            string
	StakedTokens    *big.Int
	Location        string
	MaxStorageSpace *big.Int
	Addr            common.Address
}, error) {
	return _DomicNode.Contract.BroadcastingNodes(&_DomicNode.CallOpts, arg0)
}

// BroadcastingNodes is a free data retrieval call binding the contract method 0xcac5800a.
//
// Solidity: function broadcastingNodes(address ) view returns(string url, string name, uint256 stakedTokens, string location, uint256 maxStorageSpace, address addr)
func (_DomicNode *DomicNodeCallerSession) BroadcastingNodes(arg0 common.Address) (struct {
	Url             string
	Name            string
	StakedTokens    *big.Int
	Location        string
	MaxStorageSpace *big.Int
	Addr            common.Address
}, error) {
	return _DomicNode.Contract.BroadcastingNodes(&_DomicNode.CallOpts, arg0)
}

// StorageNodeList is a free data retrieval call binding the contract method 0x7667d104.
//
// Solidity: function storageNodeList(uint256 ) view returns(address)
func (_DomicNode *DomicNodeCaller) StorageNodeList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DomicNode.contract.Call(opts, &out, "storageNodeList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageNodeList is a free data retrieval call binding the contract method 0x7667d104.
//
// Solidity: function storageNodeList(uint256 ) view returns(address)
func (_DomicNode *DomicNodeSession) StorageNodeList(arg0 *big.Int) (common.Address, error) {
	return _DomicNode.Contract.StorageNodeList(&_DomicNode.CallOpts, arg0)
}

// StorageNodeList is a free data retrieval call binding the contract method 0x7667d104.
//
// Solidity: function storageNodeList(uint256 ) view returns(address)
func (_DomicNode *DomicNodeCallerSession) StorageNodeList(arg0 *big.Int) (common.Address, error) {
	return _DomicNode.Contract.StorageNodeList(&_DomicNode.CallOpts, arg0)
}

// StorageNodes is a free data retrieval call binding the contract method 0x8118eb33.
//
// Solidity: function storageNodes(address ) view returns(string url, string name, uint256 stakedTokens, string location, uint256 maxStorageSpace, address addr)
func (_DomicNode *DomicNodeCaller) StorageNodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Url             string
	Name            string
	StakedTokens    *big.Int
	Location        string
	MaxStorageSpace *big.Int
	Addr            common.Address
}, error) {
	var out []interface{}
	err := _DomicNode.contract.Call(opts, &out, "storageNodes", arg0)

	outstruct := new(struct {
		Url             string
		Name            string
		StakedTokens    *big.Int
		Location        string
		MaxStorageSpace *big.Int
		Addr            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Url = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.StakedTokens = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Location = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.MaxStorageSpace = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Addr = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StorageNodes is a free data retrieval call binding the contract method 0x8118eb33.
//
// Solidity: function storageNodes(address ) view returns(string url, string name, uint256 stakedTokens, string location, uint256 maxStorageSpace, address addr)
func (_DomicNode *DomicNodeSession) StorageNodes(arg0 common.Address) (struct {
	Url             string
	Name            string
	StakedTokens    *big.Int
	Location        string
	MaxStorageSpace *big.Int
	Addr            common.Address
}, error) {
	return _DomicNode.Contract.StorageNodes(&_DomicNode.CallOpts, arg0)
}

// StorageNodes is a free data retrieval call binding the contract method 0x8118eb33.
//
// Solidity: function storageNodes(address ) view returns(string url, string name, uint256 stakedTokens, string location, uint256 maxStorageSpace, address addr)
func (_DomicNode *DomicNodeCallerSession) StorageNodes(arg0 common.Address) (struct {
	Url             string
	Name            string
	StakedTokens    *big.Int
	Location        string
	MaxStorageSpace *big.Int
	Addr            common.Address
}, error) {
	return _DomicNode.Contract.StorageNodes(&_DomicNode.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DomicNode *DomicNodeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DomicNode.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DomicNode *DomicNodeSession) Version() (string, error) {
	return _DomicNode.Contract.Version(&_DomicNode.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DomicNode *DomicNodeCallerSession) Version() (string, error) {
	return _DomicNode.Contract.Version(&_DomicNode.CallOpts)
}

// RegisterBroadcastNode is a paid mutator transaction binding the contract method 0x3ef98a18.
//
// Solidity: function RegisterBroadcastNode((string,string,uint256,string,uint256,address) info) returns()
func (_DomicNode *DomicNodeTransactor) RegisterBroadcastNode(opts *bind.TransactOpts, info DomiconNodeNodeInfo) (*types.Transaction, error) {
	return _DomicNode.contract.Transact(opts, "RegisterBroadcastNode", info)
}

// RegisterBroadcastNode is a paid mutator transaction binding the contract method 0x3ef98a18.
//
// Solidity: function RegisterBroadcastNode((string,string,uint256,string,uint256,address) info) returns()
func (_DomicNode *DomicNodeSession) RegisterBroadcastNode(info DomiconNodeNodeInfo) (*types.Transaction, error) {
	return _DomicNode.Contract.RegisterBroadcastNode(&_DomicNode.TransactOpts, info)
}

// RegisterBroadcastNode is a paid mutator transaction binding the contract method 0x3ef98a18.
//
// Solidity: function RegisterBroadcastNode((string,string,uint256,string,uint256,address) info) returns()
func (_DomicNode *DomicNodeTransactorSession) RegisterBroadcastNode(info DomiconNodeNodeInfo) (*types.Transaction, error) {
	return _DomicNode.Contract.RegisterBroadcastNode(&_DomicNode.TransactOpts, info)
}

// RegisterStorageNodeList is a paid mutator transaction binding the contract method 0xce3c7750.
//
// Solidity: function RegisterStorageNodeList((string,string,uint256,string,uint256,address) info) returns()
func (_DomicNode *DomicNodeTransactor) RegisterStorageNodeList(opts *bind.TransactOpts, info DomiconNodeNodeInfo) (*types.Transaction, error) {
	return _DomicNode.contract.Transact(opts, "RegisterStorageNodeList", info)
}

// RegisterStorageNodeList is a paid mutator transaction binding the contract method 0xce3c7750.
//
// Solidity: function RegisterStorageNodeList((string,string,uint256,string,uint256,address) info) returns()
func (_DomicNode *DomicNodeSession) RegisterStorageNodeList(info DomiconNodeNodeInfo) (*types.Transaction, error) {
	return _DomicNode.Contract.RegisterStorageNodeList(&_DomicNode.TransactOpts, info)
}

// RegisterStorageNodeList is a paid mutator transaction binding the contract method 0xce3c7750.
//
// Solidity: function RegisterStorageNodeList((string,string,uint256,string,uint256,address) info) returns()
func (_DomicNode *DomicNodeTransactorSession) RegisterStorageNodeList(info DomiconNodeNodeInfo) (*types.Transaction, error) {
	return _DomicNode.Contract.RegisterStorageNodeList(&_DomicNode.TransactOpts, info)
}

// SetDom is a paid mutator transaction binding the contract method 0x6d881989.
//
// Solidity: function SetDom(address addr) returns()
func (_DomicNode *DomicNodeTransactor) SetDom(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _DomicNode.contract.Transact(opts, "SetDom", addr)
}

// SetDom is a paid mutator transaction binding the contract method 0x6d881989.
//
// Solidity: function SetDom(address addr) returns()
func (_DomicNode *DomicNodeSession) SetDom(addr common.Address) (*types.Transaction, error) {
	return _DomicNode.Contract.SetDom(&_DomicNode.TransactOpts, addr)
}

// SetDom is a paid mutator transaction binding the contract method 0x6d881989.
//
// Solidity: function SetDom(address addr) returns()
func (_DomicNode *DomicNodeTransactorSession) SetDom(addr common.Address) (*types.Transaction, error) {
	return _DomicNode.Contract.SetDom(&_DomicNode.TransactOpts, addr)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DomicNode *DomicNodeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomicNode.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DomicNode *DomicNodeSession) Initialize() (*types.Transaction, error) {
	return _DomicNode.Contract.Initialize(&_DomicNode.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DomicNode *DomicNodeTransactorSession) Initialize() (*types.Transaction, error) {
	return _DomicNode.Contract.Initialize(&_DomicNode.TransactOpts)
}

// DomicNodeBroadcastNodeIterator is returned from FilterBroadcastNode and is used to iterate over the raw logs and unpacked data for BroadcastNode events raised by the DomicNode contract.
type DomicNodeBroadcastNodeIterator struct {
	Event *DomicNodeBroadcastNode // Event containing the contract specifics and raw log

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
func (it *DomicNodeBroadcastNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomicNodeBroadcastNode)
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
		it.Event = new(DomicNodeBroadcastNode)
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
func (it *DomicNodeBroadcastNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomicNodeBroadcastNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomicNodeBroadcastNode represents a BroadcastNode event raised by the DomicNode contract.
type DomicNodeBroadcastNode struct {
	Add          common.Address
	Url          string
	Name         string
	StakedTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBroadcastNode is a free log retrieval operation binding the contract event 0xf81ce16a7ccf3a5a010dfa9ea629627f1144fc81731e9d33059eb7bf82616815.
//
// Solidity: event BroadcastNode(address indexed add, string url, string name, uint256 stakedTokens)
func (_DomicNode *DomicNodeFilterer) FilterBroadcastNode(opts *bind.FilterOpts, add []common.Address) (*DomicNodeBroadcastNodeIterator, error) {

	var addRule []interface{}
	for _, addItem := range add {
		addRule = append(addRule, addItem)
	}

	logs, sub, err := _DomicNode.contract.FilterLogs(opts, "BroadcastNode", addRule)
	if err != nil {
		return nil, err
	}
	return &DomicNodeBroadcastNodeIterator{contract: _DomicNode.contract, event: "BroadcastNode", logs: logs, sub: sub}, nil
}

// WatchBroadcastNode is a free log subscription operation binding the contract event 0xf81ce16a7ccf3a5a010dfa9ea629627f1144fc81731e9d33059eb7bf82616815.
//
// Solidity: event BroadcastNode(address indexed add, string url, string name, uint256 stakedTokens)
func (_DomicNode *DomicNodeFilterer) WatchBroadcastNode(opts *bind.WatchOpts, sink chan<- *DomicNodeBroadcastNode, add []common.Address) (event.Subscription, error) {

	var addRule []interface{}
	for _, addItem := range add {
		addRule = append(addRule, addItem)
	}

	logs, sub, err := _DomicNode.contract.WatchLogs(opts, "BroadcastNode", addRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomicNodeBroadcastNode)
				if err := _DomicNode.contract.UnpackLog(event, "BroadcastNode", log); err != nil {
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

// ParseBroadcastNode is a log parse operation binding the contract event 0xf81ce16a7ccf3a5a010dfa9ea629627f1144fc81731e9d33059eb7bf82616815.
//
// Solidity: event BroadcastNode(address indexed add, string url, string name, uint256 stakedTokens)
func (_DomicNode *DomicNodeFilterer) ParseBroadcastNode(log types.Log) (*DomicNodeBroadcastNode, error) {
	event := new(DomicNodeBroadcastNode)
	if err := _DomicNode.contract.UnpackLog(event, "BroadcastNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomicNodeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DomicNode contract.
type DomicNodeInitializedIterator struct {
	Event *DomicNodeInitialized // Event containing the contract specifics and raw log

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
func (it *DomicNodeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomicNodeInitialized)
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
		it.Event = new(DomicNodeInitialized)
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
func (it *DomicNodeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomicNodeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomicNodeInitialized represents a Initialized event raised by the DomicNode contract.
type DomicNodeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DomicNode *DomicNodeFilterer) FilterInitialized(opts *bind.FilterOpts) (*DomicNodeInitializedIterator, error) {

	logs, sub, err := _DomicNode.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DomicNodeInitializedIterator{contract: _DomicNode.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DomicNode *DomicNodeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DomicNodeInitialized) (event.Subscription, error) {

	logs, sub, err := _DomicNode.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomicNodeInitialized)
				if err := _DomicNode.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DomicNode *DomicNodeFilterer) ParseInitialized(log types.Log) (*DomicNodeInitialized, error) {
	event := new(DomicNodeInitialized)
	if err := _DomicNode.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomicNodeStorageNodeIterator is returned from FilterStorageNode and is used to iterate over the raw logs and unpacked data for StorageNode events raised by the DomicNode contract.
type DomicNodeStorageNodeIterator struct {
	Event *DomicNodeStorageNode // Event containing the contract specifics and raw log

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
func (it *DomicNodeStorageNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomicNodeStorageNode)
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
		it.Event = new(DomicNodeStorageNode)
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
func (it *DomicNodeStorageNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomicNodeStorageNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomicNodeStorageNode represents a StorageNode event raised by the DomicNode contract.
type DomicNodeStorageNode struct {
	Add          common.Address
	Url          string
	Name         string
	StakedTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStorageNode is a free log retrieval operation binding the contract event 0xf9d8c2f358d69a5e1c66d7fd47740827d4a975987462ad9a863ea0ad5b055f82.
//
// Solidity: event StorageNode(address indexed add, string url, string name, uint256 stakedTokens)
func (_DomicNode *DomicNodeFilterer) FilterStorageNode(opts *bind.FilterOpts, add []common.Address) (*DomicNodeStorageNodeIterator, error) {

	var addRule []interface{}
	for _, addItem := range add {
		addRule = append(addRule, addItem)
	}

	logs, sub, err := _DomicNode.contract.FilterLogs(opts, "StorageNode", addRule)
	if err != nil {
		return nil, err
	}
	return &DomicNodeStorageNodeIterator{contract: _DomicNode.contract, event: "StorageNode", logs: logs, sub: sub}, nil
}

// WatchStorageNode is a free log subscription operation binding the contract event 0xf9d8c2f358d69a5e1c66d7fd47740827d4a975987462ad9a863ea0ad5b055f82.
//
// Solidity: event StorageNode(address indexed add, string url, string name, uint256 stakedTokens)
func (_DomicNode *DomicNodeFilterer) WatchStorageNode(opts *bind.WatchOpts, sink chan<- *DomicNodeStorageNode, add []common.Address) (event.Subscription, error) {

	var addRule []interface{}
	for _, addItem := range add {
		addRule = append(addRule, addItem)
	}

	logs, sub, err := _DomicNode.contract.WatchLogs(opts, "StorageNode", addRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomicNodeStorageNode)
				if err := _DomicNode.contract.UnpackLog(event, "StorageNode", log); err != nil {
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

// ParseStorageNode is a log parse operation binding the contract event 0xf9d8c2f358d69a5e1c66d7fd47740827d4a975987462ad9a863ea0ad5b055f82.
//
// Solidity: event StorageNode(address indexed add, string url, string name, uint256 stakedTokens)
func (_DomicNode *DomicNodeFilterer) ParseStorageNode(log types.Log) (*DomicNodeStorageNode, error) {
	event := new(DomicNodeStorageNode)
	if err := _DomicNode.contract.UnpackLog(event, "StorageNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
