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

// DomiconCommitmentDaDetails is an auto generated low-level Go binding around an user-defined struct.
type DomiconCommitmentDaDetails struct {
	Timestamp      *big.Int
	HashSignatures [32]byte
}

// PairingG1Point is an auto generated low-level Go binding around an user-defined struct.
type PairingG1Point struct {
	X *big.Int
	Y *big.Int
}

// DomicCommitMetaData contains all meta data concerning the DomicCommit contract.
var DomicCommitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPairing.G1Point\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dasKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"SendDACommitment\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"COMMITMENTS\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hashSignatures\",\"type\":\"bytes32\"}],\"internalType\":\"structDomiconCommitment.DaDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"COMMITMENTS\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hashSignatures\",\"type\":\"bytes32\"}],\"internalType\":\"structDomiconCommitment.DaDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"SetDom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_length\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_dasKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structPairing.G1Point\",\"name\":\"_commitment\",\"type\":\"tuple\"}],\"name\":\"SubmitCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committeeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"daDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hashSignatures\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domiconNode\",\"outputs\":[{\"internalType\":\"contractDomiconNode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUserCommitments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structPairing.G1Point\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"indices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDomiconNode\",\"name\":\"_domiconNode\",\"type\":\"address\"},{\"internalType\":\"contractStorageManagement\",\"name\":\"_storageManagement\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageManagement\",\"outputs\":[{\"internalType\":\"contractStorageManagement\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userCommitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DomicCommitABI is the input ABI used to generate the binding from.
// Deprecated: Use DomicCommitMetaData.ABI instead.
var DomicCommitABI = DomicCommitMetaData.ABI

// DomicCommit is an auto generated Go binding around an Ethereum contract.
type DomicCommit struct {
	DomicCommitCaller     // Read-only binding to the contract
	DomicCommitTransactor // Write-only binding to the contract
	DomicCommitFilterer   // Log filterer for contract events
}

// DomicCommitCaller is an auto generated read-only Go binding around an Ethereum contract.
type DomicCommitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomicCommitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DomicCommitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomicCommitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DomicCommitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomicCommitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DomicCommitSession struct {
	Contract     *DomicCommit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DomicCommitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DomicCommitCallerSession struct {
	Contract *DomicCommitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DomicCommitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DomicCommitTransactorSession struct {
	Contract     *DomicCommitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DomicCommitRaw is an auto generated low-level Go binding around an Ethereum contract.
type DomicCommitRaw struct {
	Contract *DomicCommit // Generic contract binding to access the raw methods on
}

// DomicCommitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DomicCommitCallerRaw struct {
	Contract *DomicCommitCaller // Generic read-only contract binding to access the raw methods on
}

// DomicCommitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DomicCommitTransactorRaw struct {
	Contract *DomicCommitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDomicCommit creates a new instance of DomicCommit, bound to a specific deployed contract.
func NewDomicCommit(address common.Address, backend bind.ContractBackend) (*DomicCommit, error) {
	contract, err := bindDomicCommit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DomicCommit{DomicCommitCaller: DomicCommitCaller{contract: contract}, DomicCommitTransactor: DomicCommitTransactor{contract: contract}, DomicCommitFilterer: DomicCommitFilterer{contract: contract}}, nil
}

// NewDomicCommitCaller creates a new read-only instance of DomicCommit, bound to a specific deployed contract.
func NewDomicCommitCaller(address common.Address, caller bind.ContractCaller) (*DomicCommitCaller, error) {
	contract, err := bindDomicCommit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DomicCommitCaller{contract: contract}, nil
}

// NewDomicCommitTransactor creates a new write-only instance of DomicCommit, bound to a specific deployed contract.
func NewDomicCommitTransactor(address common.Address, transactor bind.ContractTransactor) (*DomicCommitTransactor, error) {
	contract, err := bindDomicCommit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DomicCommitTransactor{contract: contract}, nil
}

// NewDomicCommitFilterer creates a new log filterer instance of DomicCommit, bound to a specific deployed contract.
func NewDomicCommitFilterer(address common.Address, filterer bind.ContractFilterer) (*DomicCommitFilterer, error) {
	contract, err := bindDomicCommit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DomicCommitFilterer{contract: contract}, nil
}

// bindDomicCommit binds a generic wrapper to an already deployed contract.
func bindDomicCommit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DomicCommitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomicCommit *DomicCommitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomicCommit.Contract.DomicCommitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomicCommit *DomicCommitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomicCommit.Contract.DomicCommitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomicCommit *DomicCommitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomicCommit.Contract.DomicCommitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomicCommit *DomicCommitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomicCommit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomicCommit *DomicCommitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomicCommit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomicCommit *DomicCommitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomicCommit.Contract.contract.Transact(opts, method, params...)
}

// COMMITMENTS is a free data retrieval call binding the contract method 0xe6e4b404.
//
// Solidity: function COMMITMENTS(uint256 _nonce) view returns((uint256,bytes32))
func (_DomicCommit *DomicCommitCaller) COMMITMENTS(opts *bind.CallOpts, _nonce *big.Int) (DomiconCommitmentDaDetails, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "COMMITMENTS", _nonce)

	if err != nil {
		return *new(DomiconCommitmentDaDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(DomiconCommitmentDaDetails)).(*DomiconCommitmentDaDetails)

	return out0, err

}

// COMMITMENTS is a free data retrieval call binding the contract method 0xe6e4b404.
//
// Solidity: function COMMITMENTS(uint256 _nonce) view returns((uint256,bytes32))
func (_DomicCommit *DomicCommitSession) COMMITMENTS(_nonce *big.Int) (DomiconCommitmentDaDetails, error) {
	return _DomicCommit.Contract.COMMITMENTS(&_DomicCommit.CallOpts, _nonce)
}

// COMMITMENTS is a free data retrieval call binding the contract method 0xe6e4b404.
//
// Solidity: function COMMITMENTS(uint256 _nonce) view returns((uint256,bytes32))
func (_DomicCommit *DomicCommitCallerSession) COMMITMENTS(_nonce *big.Int) (DomiconCommitmentDaDetails, error) {
	return _DomicCommit.Contract.COMMITMENTS(&_DomicCommit.CallOpts, _nonce)
}

// COMMITMENTS0 is a free data retrieval call binding the contract method 0xea9b6adf.
//
// Solidity: function COMMITMENTS(address _user, uint256 _index) view returns((uint256,bytes32))
func (_DomicCommit *DomicCommitCaller) COMMITMENTS0(opts *bind.CallOpts, _user common.Address, _index *big.Int) (DomiconCommitmentDaDetails, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "COMMITMENTS0", _user, _index)

	if err != nil {
		return *new(DomiconCommitmentDaDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(DomiconCommitmentDaDetails)).(*DomiconCommitmentDaDetails)

	return out0, err

}

// COMMITMENTS0 is a free data retrieval call binding the contract method 0xea9b6adf.
//
// Solidity: function COMMITMENTS(address _user, uint256 _index) view returns((uint256,bytes32))
func (_DomicCommit *DomicCommitSession) COMMITMENTS0(_user common.Address, _index *big.Int) (DomiconCommitmentDaDetails, error) {
	return _DomicCommit.Contract.COMMITMENTS0(&_DomicCommit.CallOpts, _user, _index)
}

// COMMITMENTS0 is a free data retrieval call binding the contract method 0xea9b6adf.
//
// Solidity: function COMMITMENTS(address _user, uint256 _index) view returns((uint256,bytes32))
func (_DomicCommit *DomicCommitCallerSession) COMMITMENTS0(_user common.Address, _index *big.Int) (DomiconCommitmentDaDetails, error) {
	return _DomicCommit.Contract.COMMITMENTS0(&_DomicCommit.CallOpts, _user, _index)
}

// DOM is a free data retrieval call binding the contract method 0x6a57f6b1.
//
// Solidity: function DOM() view returns(address)
func (_DomicCommit *DomicCommitCaller) DOM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "DOM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DOM is a free data retrieval call binding the contract method 0x6a57f6b1.
//
// Solidity: function DOM() view returns(address)
func (_DomicCommit *DomicCommitSession) DOM() (common.Address, error) {
	return _DomicCommit.Contract.DOM(&_DomicCommit.CallOpts)
}

// DOM is a free data retrieval call binding the contract method 0x6a57f6b1.
//
// Solidity: function DOM() view returns(address)
func (_DomicCommit *DomicCommitCallerSession) DOM() (common.Address, error) {
	return _DomicCommit.Contract.DOM(&_DomicCommit.CallOpts)
}

// Commitments is a free data retrieval call binding the contract method 0x49ce8997.
//
// Solidity: function commitments(uint256 ) view returns(uint256 X, uint256 Y)
func (_DomicCommit *DomicCommitCaller) Commitments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "commitments", arg0)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Commitments is a free data retrieval call binding the contract method 0x49ce8997.
//
// Solidity: function commitments(uint256 ) view returns(uint256 X, uint256 Y)
func (_DomicCommit *DomicCommitSession) Commitments(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _DomicCommit.Contract.Commitments(&_DomicCommit.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x49ce8997.
//
// Solidity: function commitments(uint256 ) view returns(uint256 X, uint256 Y)
func (_DomicCommit *DomicCommitCallerSession) Commitments(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _DomicCommit.Contract.Commitments(&_DomicCommit.CallOpts, arg0)
}

// CommitteeRoot is a free data retrieval call binding the contract method 0x931c64c6.
//
// Solidity: function committeeRoot() view returns(bytes32)
func (_DomicCommit *DomicCommitCaller) CommitteeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "committeeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommitteeRoot is a free data retrieval call binding the contract method 0x931c64c6.
//
// Solidity: function committeeRoot() view returns(bytes32)
func (_DomicCommit *DomicCommitSession) CommitteeRoot() ([32]byte, error) {
	return _DomicCommit.Contract.CommitteeRoot(&_DomicCommit.CallOpts)
}

// CommitteeRoot is a free data retrieval call binding the contract method 0x931c64c6.
//
// Solidity: function committeeRoot() view returns(bytes32)
func (_DomicCommit *DomicCommitCallerSession) CommitteeRoot() ([32]byte, error) {
	return _DomicCommit.Contract.CommitteeRoot(&_DomicCommit.CallOpts)
}

// DaDetails is a free data retrieval call binding the contract method 0x3f995b74.
//
// Solidity: function daDetails(bytes32 ) view returns(uint256 timestamp, bytes32 hashSignatures)
func (_DomicCommit *DomicCommitCaller) DaDetails(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Timestamp      *big.Int
	HashSignatures [32]byte
}, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "daDetails", arg0)

	outstruct := new(struct {
		Timestamp      *big.Int
		HashSignatures [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HashSignatures = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// DaDetails is a free data retrieval call binding the contract method 0x3f995b74.
//
// Solidity: function daDetails(bytes32 ) view returns(uint256 timestamp, bytes32 hashSignatures)
func (_DomicCommit *DomicCommitSession) DaDetails(arg0 [32]byte) (struct {
	Timestamp      *big.Int
	HashSignatures [32]byte
}, error) {
	return _DomicCommit.Contract.DaDetails(&_DomicCommit.CallOpts, arg0)
}

// DaDetails is a free data retrieval call binding the contract method 0x3f995b74.
//
// Solidity: function daDetails(bytes32 ) view returns(uint256 timestamp, bytes32 hashSignatures)
func (_DomicCommit *DomicCommitCallerSession) DaDetails(arg0 [32]byte) (struct {
	Timestamp      *big.Int
	HashSignatures [32]byte
}, error) {
	return _DomicCommit.Contract.DaDetails(&_DomicCommit.CallOpts, arg0)
}

// DomiconNode is a free data retrieval call binding the contract method 0x5fa4ad36.
//
// Solidity: function domiconNode() view returns(address)
func (_DomicCommit *DomicCommitCaller) DomiconNode(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "domiconNode")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DomiconNode is a free data retrieval call binding the contract method 0x5fa4ad36.
//
// Solidity: function domiconNode() view returns(address)
func (_DomicCommit *DomicCommitSession) DomiconNode() (common.Address, error) {
	return _DomicCommit.Contract.DomiconNode(&_DomicCommit.CallOpts)
}

// DomiconNode is a free data retrieval call binding the contract method 0x5fa4ad36.
//
// Solidity: function domiconNode() view returns(address)
func (_DomicCommit *DomicCommitCallerSession) DomiconNode() (common.Address, error) {
	return _DomicCommit.Contract.DomiconNode(&_DomicCommit.CallOpts)
}

// GetUserCommitments is a free data retrieval call binding the contract method 0x6471cf22.
//
// Solidity: function getUserCommitments(address _user, uint256 _index) view returns((uint256,uint256))
func (_DomicCommit *DomicCommitCaller) GetUserCommitments(opts *bind.CallOpts, _user common.Address, _index *big.Int) (PairingG1Point, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "getUserCommitments", _user, _index)

	if err != nil {
		return *new(PairingG1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(PairingG1Point)).(*PairingG1Point)

	return out0, err

}

// GetUserCommitments is a free data retrieval call binding the contract method 0x6471cf22.
//
// Solidity: function getUserCommitments(address _user, uint256 _index) view returns((uint256,uint256))
func (_DomicCommit *DomicCommitSession) GetUserCommitments(_user common.Address, _index *big.Int) (PairingG1Point, error) {
	return _DomicCommit.Contract.GetUserCommitments(&_DomicCommit.CallOpts, _user, _index)
}

// GetUserCommitments is a free data retrieval call binding the contract method 0x6471cf22.
//
// Solidity: function getUserCommitments(address _user, uint256 _index) view returns((uint256,uint256))
func (_DomicCommit *DomicCommitCallerSession) GetUserCommitments(_user common.Address, _index *big.Int) (PairingG1Point, error) {
	return _DomicCommit.Contract.GetUserCommitments(&_DomicCommit.CallOpts, _user, _index)
}

// Indices is a free data retrieval call binding the contract method 0x5063e207.
//
// Solidity: function indices(address ) view returns(uint256)
func (_DomicCommit *DomicCommitCaller) Indices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "indices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Indices is a free data retrieval call binding the contract method 0x5063e207.
//
// Solidity: function indices(address ) view returns(uint256)
func (_DomicCommit *DomicCommitSession) Indices(arg0 common.Address) (*big.Int, error) {
	return _DomicCommit.Contract.Indices(&_DomicCommit.CallOpts, arg0)
}

// Indices is a free data retrieval call binding the contract method 0x5063e207.
//
// Solidity: function indices(address ) view returns(uint256)
func (_DomicCommit *DomicCommitCallerSession) Indices(arg0 common.Address) (*big.Int, error) {
	return _DomicCommit.Contract.Indices(&_DomicCommit.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_DomicCommit *DomicCommitCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_DomicCommit *DomicCommitSession) Nonce() (*big.Int, error) {
	return _DomicCommit.Contract.Nonce(&_DomicCommit.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_DomicCommit *DomicCommitCallerSession) Nonce() (*big.Int, error) {
	return _DomicCommit.Contract.Nonce(&_DomicCommit.CallOpts)
}

// StorageManagement is a free data retrieval call binding the contract method 0x98bd4dee.
//
// Solidity: function storageManagement() view returns(address)
func (_DomicCommit *DomicCommitCaller) StorageManagement(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "storageManagement")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageManagement is a free data retrieval call binding the contract method 0x98bd4dee.
//
// Solidity: function storageManagement() view returns(address)
func (_DomicCommit *DomicCommitSession) StorageManagement() (common.Address, error) {
	return _DomicCommit.Contract.StorageManagement(&_DomicCommit.CallOpts)
}

// StorageManagement is a free data retrieval call binding the contract method 0x98bd4dee.
//
// Solidity: function storageManagement() view returns(address)
func (_DomicCommit *DomicCommitCallerSession) StorageManagement() (common.Address, error) {
	return _DomicCommit.Contract.StorageManagement(&_DomicCommit.CallOpts)
}

// UserCommitments is a free data retrieval call binding the contract method 0x9a546848.
//
// Solidity: function userCommitments(address , uint256 ) view returns(uint256 X, uint256 Y)
func (_DomicCommit *DomicCommitCaller) UserCommitments(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "userCommitments", arg0, arg1)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserCommitments is a free data retrieval call binding the contract method 0x9a546848.
//
// Solidity: function userCommitments(address , uint256 ) view returns(uint256 X, uint256 Y)
func (_DomicCommit *DomicCommitSession) UserCommitments(arg0 common.Address, arg1 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _DomicCommit.Contract.UserCommitments(&_DomicCommit.CallOpts, arg0, arg1)
}

// UserCommitments is a free data retrieval call binding the contract method 0x9a546848.
//
// Solidity: function userCommitments(address , uint256 ) view returns(uint256 X, uint256 Y)
func (_DomicCommit *DomicCommitCallerSession) UserCommitments(arg0 common.Address, arg1 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _DomicCommit.Contract.UserCommitments(&_DomicCommit.CallOpts, arg0, arg1)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DomicCommit *DomicCommitCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DomicCommit.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DomicCommit *DomicCommitSession) Version() (string, error) {
	return _DomicCommit.Contract.Version(&_DomicCommit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DomicCommit *DomicCommitCallerSession) Version() (string, error) {
	return _DomicCommit.Contract.Version(&_DomicCommit.CallOpts)
}

// SetDom is a paid mutator transaction binding the contract method 0x6d881989.
//
// Solidity: function SetDom(address _addr) returns()
func (_DomicCommit *DomicCommitTransactor) SetDom(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _DomicCommit.contract.Transact(opts, "SetDom", _addr)
}

// SetDom is a paid mutator transaction binding the contract method 0x6d881989.
//
// Solidity: function SetDom(address _addr) returns()
func (_DomicCommit *DomicCommitSession) SetDom(_addr common.Address) (*types.Transaction, error) {
	return _DomicCommit.Contract.SetDom(&_DomicCommit.TransactOpts, _addr)
}

// SetDom is a paid mutator transaction binding the contract method 0x6d881989.
//
// Solidity: function SetDom(address _addr) returns()
func (_DomicCommit *DomicCommitTransactorSession) SetDom(_addr common.Address) (*types.Transaction, error) {
	return _DomicCommit.Contract.SetDom(&_DomicCommit.TransactOpts, _addr)
}

// SubmitCommitment is a paid mutator transaction binding the contract method 0x259d8cdd.
//
// Solidity: function SubmitCommitment(uint64 _index, uint64 _length, bytes32 _dasKey, bytes[] _signatures, (uint256,uint256) _commitment) returns()
func (_DomicCommit *DomicCommitTransactor) SubmitCommitment(opts *bind.TransactOpts, _index uint64, _length uint64, _dasKey [32]byte, _signatures [][]byte, _commitment PairingG1Point) (*types.Transaction, error) {
	return _DomicCommit.contract.Transact(opts, "SubmitCommitment", _index, _length, _dasKey, _signatures, _commitment)
}

// SubmitCommitment is a paid mutator transaction binding the contract method 0x259d8cdd.
//
// Solidity: function SubmitCommitment(uint64 _index, uint64 _length, bytes32 _dasKey, bytes[] _signatures, (uint256,uint256) _commitment) returns()
func (_DomicCommit *DomicCommitSession) SubmitCommitment(_index uint64, _length uint64, _dasKey [32]byte, _signatures [][]byte, _commitment PairingG1Point) (*types.Transaction, error) {
	return _DomicCommit.Contract.SubmitCommitment(&_DomicCommit.TransactOpts, _index, _length, _dasKey, _signatures, _commitment)
}

// SubmitCommitment is a paid mutator transaction binding the contract method 0x259d8cdd.
//
// Solidity: function SubmitCommitment(uint64 _index, uint64 _length, bytes32 _dasKey, bytes[] _signatures, (uint256,uint256) _commitment) returns()
func (_DomicCommit *DomicCommitTransactorSession) SubmitCommitment(_index uint64, _length uint64, _dasKey [32]byte, _signatures [][]byte, _commitment PairingG1Point) (*types.Transaction, error) {
	return _DomicCommit.Contract.SubmitCommitment(&_DomicCommit.TransactOpts, _index, _length, _dasKey, _signatures, _commitment)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _domiconNode, address _storageManagement) returns()
func (_DomicCommit *DomicCommitTransactor) Initialize(opts *bind.TransactOpts, _domiconNode common.Address, _storageManagement common.Address) (*types.Transaction, error) {
	return _DomicCommit.contract.Transact(opts, "initialize", _domiconNode, _storageManagement)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _domiconNode, address _storageManagement) returns()
func (_DomicCommit *DomicCommitSession) Initialize(_domiconNode common.Address, _storageManagement common.Address) (*types.Transaction, error) {
	return _DomicCommit.Contract.Initialize(&_DomicCommit.TransactOpts, _domiconNode, _storageManagement)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _domiconNode, address _storageManagement) returns()
func (_DomicCommit *DomicCommitTransactorSession) Initialize(_domiconNode common.Address, _storageManagement common.Address) (*types.Transaction, error) {
	return _DomicCommit.Contract.Initialize(&_DomicCommit.TransactOpts, _domiconNode, _storageManagement)
}

// DomicCommitInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DomicCommit contract.
type DomicCommitInitializedIterator struct {
	Event *DomicCommitInitialized // Event containing the contract specifics and raw log

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
func (it *DomicCommitInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomicCommitInitialized)
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
		it.Event = new(DomicCommitInitialized)
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
func (it *DomicCommitInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomicCommitInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomicCommitInitialized represents a Initialized event raised by the DomicCommit contract.
type DomicCommitInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DomicCommit *DomicCommitFilterer) FilterInitialized(opts *bind.FilterOpts) (*DomicCommitInitializedIterator, error) {

	logs, sub, err := _DomicCommit.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DomicCommitInitializedIterator{contract: _DomicCommit.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DomicCommit *DomicCommitFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DomicCommitInitialized) (event.Subscription, error) {

	logs, sub, err := _DomicCommit.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomicCommitInitialized)
				if err := _DomicCommit.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DomicCommit *DomicCommitFilterer) ParseInitialized(log types.Log) (*DomicCommitInitialized, error) {
	event := new(DomicCommitInitialized)
	if err := _DomicCommit.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomicCommitSendDACommitmentIterator is returned from FilterSendDACommitment and is used to iterate over the raw logs and unpacked data for SendDACommitment events raised by the DomicCommit contract.
type DomicCommitSendDACommitmentIterator struct {
	Event *DomicCommitSendDACommitment // Event containing the contract specifics and raw log

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
func (it *DomicCommitSendDACommitmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomicCommitSendDACommitment)
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
		it.Event = new(DomicCommitSendDACommitment)
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
func (it *DomicCommitSendDACommitmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomicCommitSendDACommitmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomicCommitSendDACommitment represents a SendDACommitment event raised by the DomicCommit contract.
type DomicCommitSendDACommitment struct {
	User       common.Address
	Commitment PairingG1Point
	Timestamp  *big.Int
	Nonce      *big.Int
	Index      *big.Int
	Len        *big.Int
	Root       [32]byte
	DasKey     [32]byte
	Signatures [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendDACommitment is a free log retrieval operation binding the contract event 0xfb3a62c71e56d41dd4a93b72df383612d4377bd1d2c2466402c09092931815a9.
//
// Solidity: event SendDACommitment(address user, (uint256,uint256) commitment, uint256 timestamp, uint256 nonce, uint256 index, uint256 len, bytes32 root, bytes32 dasKey, bytes[] signatures)
func (_DomicCommit *DomicCommitFilterer) FilterSendDACommitment(opts *bind.FilterOpts) (*DomicCommitSendDACommitmentIterator, error) {

	logs, sub, err := _DomicCommit.contract.FilterLogs(opts, "SendDACommitment")
	if err != nil {
		return nil, err
	}
	return &DomicCommitSendDACommitmentIterator{contract: _DomicCommit.contract, event: "SendDACommitment", logs: logs, sub: sub}, nil
}

// WatchSendDACommitment is a free log subscription operation binding the contract event 0xfb3a62c71e56d41dd4a93b72df383612d4377bd1d2c2466402c09092931815a9.
//
// Solidity: event SendDACommitment(address user, (uint256,uint256) commitment, uint256 timestamp, uint256 nonce, uint256 index, uint256 len, bytes32 root, bytes32 dasKey, bytes[] signatures)
func (_DomicCommit *DomicCommitFilterer) WatchSendDACommitment(opts *bind.WatchOpts, sink chan<- *DomicCommitSendDACommitment) (event.Subscription, error) {

	logs, sub, err := _DomicCommit.contract.WatchLogs(opts, "SendDACommitment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomicCommitSendDACommitment)
				if err := _DomicCommit.contract.UnpackLog(event, "SendDACommitment", log); err != nil {
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

// ParseSendDACommitment is a log parse operation binding the contract event 0xfb3a62c71e56d41dd4a93b72df383612d4377bd1d2c2466402c09092931815a9.
//
// Solidity: event SendDACommitment(address user, (uint256,uint256) commitment, uint256 timestamp, uint256 nonce, uint256 index, uint256 len, bytes32 root, bytes32 dasKey, bytes[] signatures)
func (_DomicCommit *DomicCommitFilterer) ParseSendDACommitment(log types.Log) (*DomicCommitSendDACommitment, error) {
	event := new(DomicCommitSendDACommitment)
	if err := _DomicCommit.contract.UnpackLog(event, "SendDACommitment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
