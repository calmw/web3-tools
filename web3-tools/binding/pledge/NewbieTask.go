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

// NewbieTaskMetaData contains all meta data concerning the NewbieTask contract.
var NewbieTaskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"checkSignatureError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"claimType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"taskStatusError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimType\",\"type\":\"uint256\"}],\"name\":\"ClaimLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimType_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"adminSetClaimType\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"faceAuthAddress_\",\"type\":\"address\"}],\"name\":\"adminSetFaceAuth\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bindingAddress_\",\"type\":\"address\"}],\"name\":\"adminSetIntoBinding\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"intoEarlyBirdAddress_\",\"type\":\"address\"}],\"name\":\"adminSetIntoEarlyBird\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"socialMiningAddress_\",\"type\":\"address\"}],\"name\":\"adminSetIntoSocialMining\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxReward_\",\"type\":\"uint256\"}],\"name\":\"adminSetMaxReward\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pledgePoolAddress_\",\"type\":\"address\"}],\"name\":\"adminSetPledgePool\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"serverAddress_\",\"type\":\"address\"}],\"name\":\"adminSetServerAddress\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskTotalRewardThreshold_\",\"type\":\"uint256\"}],\"name\":\"adminSetTaskTotalRewardThreshold\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toxAddress_\",\"type\":\"address\"}],\"name\":\"adminSetTox\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawLimitAddress_\",\"type\":\"address\"}],\"name\":\"adminSetWithdrawLimit\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[],\"name\":\"allAdmins\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"allTaskFinished\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[],\"name\":\"binding\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"contractIntoBinding\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"taskStatus\",\"type\":\"bool\"}],\"name\":\"checkSignature\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"taskStatus\",\"type\":\"bool\"}],\"name\":\"claim\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimRecord\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"taskStatus\",\"type\":\"bool\"}],\"name\":\"claimV2\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[],\"name\":\"faceAuth\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"contractFaceAuth\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"claimType\",\"type\":\"uint256\"}],\"name\":\"getTaskStatus\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\"},{\"inputs\":[],\"name\":\"initialize\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[],\"name\":\"intoEarlyBird\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"contractIntoEarlyBird\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\"},{\"inputs\":[],\"name\":\"maxReward\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"},{\"inputs\":[],\"name\":\"pledgePool\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"contractPledgePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"inputs\":[],\"name\":\"serverAddress\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"},{\"inputs\":[],\"name\":\"socialMining\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"contractIntoSocialMining\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"taskReward\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"},{\"inputs\":[],\"name\":\"taskTotalReward\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"},{\"inputs\":[],\"name\":\"taskTotalRewardThreshold\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"},{\"inputs\":[],\"name\":\"tox\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userReward\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"},{\"inputs\":[],\"name\":\"withdrawLimit\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"contractIntoWithdrawLimit\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// NewbieTaskABI is the input ABI used to generate the binding from.
// Deprecated: Use NewbieTaskMetaData.ABI instead.
var NewbieTaskABI = NewbieTaskMetaData.ABI

// NewbieTask is an auto generated Go binding around an Ethereum contract.
type NewbieTask struct {
	NewbieTaskCaller     // Read-only binding to the contract
	NewbieTaskTransactor // Write-only binding to the contract
	NewbieTaskFilterer   // Log filterer for contract events
}

// NewbieTaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type NewbieTaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewbieTaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NewbieTaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewbieTaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NewbieTaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewbieTaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NewbieTaskSession struct {
	Contract     *NewbieTask       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NewbieTaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NewbieTaskCallerSession struct {
	Contract *NewbieTaskCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NewbieTaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NewbieTaskTransactorSession struct {
	Contract     *NewbieTaskTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NewbieTaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type NewbieTaskRaw struct {
	Contract *NewbieTask // Generic contract binding to access the raw methods on
}

// NewbieTaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NewbieTaskCallerRaw struct {
	Contract *NewbieTaskCaller // Generic read-only contract binding to access the raw methods on
}

// NewbieTaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NewbieTaskTransactorRaw struct {
	Contract *NewbieTaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNewbieTask creates a new instance of NewbieTask, bound to a specific deployed contract.
func NewNewbieTask(address common.Address, backend bind.ContractBackend) (*NewbieTask, error) {
	contract, err := bindNewbieTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NewbieTask{NewbieTaskCaller: NewbieTaskCaller{contract: contract}, NewbieTaskTransactor: NewbieTaskTransactor{contract: contract}, NewbieTaskFilterer: NewbieTaskFilterer{contract: contract}}, nil
}

// NewNewbieTaskCaller creates a new read-only instance of NewbieTask, bound to a specific deployed contract.
func NewNewbieTaskCaller(address common.Address, caller bind.ContractCaller) (*NewbieTaskCaller, error) {
	contract, err := bindNewbieTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NewbieTaskCaller{contract: contract}, nil
}

// NewNewbieTaskTransactor creates a new write-only instance of NewbieTask, bound to a specific deployed contract.
func NewNewbieTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*NewbieTaskTransactor, error) {
	contract, err := bindNewbieTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NewbieTaskTransactor{contract: contract}, nil
}

// NewNewbieTaskFilterer creates a new log filterer instance of NewbieTask, bound to a specific deployed contract.
func NewNewbieTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*NewbieTaskFilterer, error) {
	contract, err := bindNewbieTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NewbieTaskFilterer{contract: contract}, nil
}

// bindNewbieTask binds a generic wrapper to an already deployed contract.
func bindNewbieTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NewbieTaskMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NewbieTask *NewbieTaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NewbieTask.Contract.NewbieTaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NewbieTask *NewbieTaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewbieTask.Contract.NewbieTaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NewbieTask *NewbieTaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NewbieTask.Contract.NewbieTaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NewbieTask *NewbieTaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NewbieTask.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NewbieTask *NewbieTaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewbieTask.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NewbieTask *NewbieTaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NewbieTask.Contract.contract.Transact(opts, method, params...)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_NewbieTask *NewbieTaskCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_NewbieTask *NewbieTaskSession) AllAdmins() ([]common.Address, error) {
	return _NewbieTask.Contract.AllAdmins(&_NewbieTask.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_NewbieTask *NewbieTaskCallerSession) AllAdmins() ([]common.Address, error) {
	return _NewbieTask.Contract.AllAdmins(&_NewbieTask.CallOpts)
}

// AllTaskFinished is a free data retrieval call binding the contract method 0xcfde5d0b.
//
// Solidity: function allTaskFinished(address user_) view returns(bool)
func (_NewbieTask *NewbieTaskCaller) AllTaskFinished(opts *bind.CallOpts, user_ common.Address) (bool, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "allTaskFinished", user_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllTaskFinished is a free data retrieval call binding the contract method 0xcfde5d0b.
//
// Solidity: function allTaskFinished(address user_) view returns(bool)
func (_NewbieTask *NewbieTaskSession) AllTaskFinished(user_ common.Address) (bool, error) {
	return _NewbieTask.Contract.AllTaskFinished(&_NewbieTask.CallOpts, user_)
}

// AllTaskFinished is a free data retrieval call binding the contract method 0xcfde5d0b.
//
// Solidity: function allTaskFinished(address user_) view returns(bool)
func (_NewbieTask *NewbieTaskCallerSession) AllTaskFinished(user_ common.Address) (bool, error) {
	return _NewbieTask.Contract.AllTaskFinished(&_NewbieTask.CallOpts, user_)
}

// Binding is a free data retrieval call binding the contract method 0x480a51a6.
//
// Solidity: function binding() view returns(address)
func (_NewbieTask *NewbieTaskCaller) Binding(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "binding")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Binding is a free data retrieval call binding the contract method 0x480a51a6.
//
// Solidity: function binding() view returns(address)
func (_NewbieTask *NewbieTaskSession) Binding() (common.Address, error) {
	return _NewbieTask.Contract.Binding(&_NewbieTask.CallOpts)
}

// Binding is a free data retrieval call binding the contract method 0x480a51a6.
//
// Solidity: function binding() view returns(address)
func (_NewbieTask *NewbieTaskCallerSession) Binding() (common.Address, error) {
	return _NewbieTask.Contract.Binding(&_NewbieTask.CallOpts)
}

// CheckSignature is a free data retrieval call binding the contract method 0x5ba33ee6.
//
// Solidity: function checkSignature(bytes signature, address sender, uint256 chainId, uint256 claimType, bool taskStatus) view returns(bool)
func (_NewbieTask *NewbieTaskCaller) CheckSignature(opts *bind.CallOpts, signature []byte, sender common.Address, chainId *big.Int, claimType *big.Int, taskStatus bool) (bool, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "checkSignature", signature, sender, chainId, claimType, taskStatus)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckSignature is a free data retrieval call binding the contract method 0x5ba33ee6.
//
// Solidity: function checkSignature(bytes signature, address sender, uint256 chainId, uint256 claimType, bool taskStatus) view returns(bool)
func (_NewbieTask *NewbieTaskSession) CheckSignature(signature []byte, sender common.Address, chainId *big.Int, claimType *big.Int, taskStatus bool) (bool, error) {
	return _NewbieTask.Contract.CheckSignature(&_NewbieTask.CallOpts, signature, sender, chainId, claimType, taskStatus)
}

// CheckSignature is a free data retrieval call binding the contract method 0x5ba33ee6.
//
// Solidity: function checkSignature(bytes signature, address sender, uint256 chainId, uint256 claimType, bool taskStatus) view returns(bool)
func (_NewbieTask *NewbieTaskCallerSession) CheckSignature(signature []byte, sender common.Address, chainId *big.Int, claimType *big.Int, taskStatus bool) (bool, error) {
	return _NewbieTask.Contract.CheckSignature(&_NewbieTask.CallOpts, signature, sender, chainId, claimType, taskStatus)
}

// ClaimRecord is a free data retrieval call binding the contract method 0xb00060f2.
//
// Solidity: function claimRecord(address , uint256 ) view returns(uint256)
func (_NewbieTask *NewbieTaskCaller) ClaimRecord(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "claimRecord", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimRecord is a free data retrieval call binding the contract method 0xb00060f2.
//
// Solidity: function claimRecord(address , uint256 ) view returns(uint256)
func (_NewbieTask *NewbieTaskSession) ClaimRecord(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NewbieTask.Contract.ClaimRecord(&_NewbieTask.CallOpts, arg0, arg1)
}

// ClaimRecord is a free data retrieval call binding the contract method 0xb00060f2.
//
// Solidity: function claimRecord(address , uint256 ) view returns(uint256)
func (_NewbieTask *NewbieTaskCallerSession) ClaimRecord(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NewbieTask.Contract.ClaimRecord(&_NewbieTask.CallOpts, arg0, arg1)
}

// FaceAuth is a free data retrieval call binding the contract method 0x7f0b4af7.
//
// Solidity: function faceAuth() view returns(address)
func (_NewbieTask *NewbieTaskCaller) FaceAuth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "faceAuth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FaceAuth is a free data retrieval call binding the contract method 0x7f0b4af7.
//
// Solidity: function faceAuth() view returns(address)
func (_NewbieTask *NewbieTaskSession) FaceAuth() (common.Address, error) {
	return _NewbieTask.Contract.FaceAuth(&_NewbieTask.CallOpts)
}

// FaceAuth is a free data retrieval call binding the contract method 0x7f0b4af7.
//
// Solidity: function faceAuth() view returns(address)
func (_NewbieTask *NewbieTaskCallerSession) FaceAuth() (common.Address, error) {
	return _NewbieTask.Contract.FaceAuth(&_NewbieTask.CallOpts)
}

// GetTaskStatus is a free data retrieval call binding the contract method 0x7e333eb2.
//
// Solidity: function getTaskStatus(address user_, uint256 claimType) view returns(bool)
func (_NewbieTask *NewbieTaskCaller) GetTaskStatus(opts *bind.CallOpts, user_ common.Address, claimType *big.Int) (bool, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "getTaskStatus", user_, claimType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetTaskStatus is a free data retrieval call binding the contract method 0x7e333eb2.
//
// Solidity: function getTaskStatus(address user_, uint256 claimType) view returns(bool)
func (_NewbieTask *NewbieTaskSession) GetTaskStatus(user_ common.Address, claimType *big.Int) (bool, error) {
	return _NewbieTask.Contract.GetTaskStatus(&_NewbieTask.CallOpts, user_, claimType)
}

// GetTaskStatus is a free data retrieval call binding the contract method 0x7e333eb2.
//
// Solidity: function getTaskStatus(address user_, uint256 claimType) view returns(bool)
func (_NewbieTask *NewbieTaskCallerSession) GetTaskStatus(user_ common.Address, claimType *big.Int) (bool, error) {
	return _NewbieTask.Contract.GetTaskStatus(&_NewbieTask.CallOpts, user_, claimType)
}

// IntoEarlyBird is a free data retrieval call binding the contract method 0x46dbc4a9.
//
// Solidity: function intoEarlyBird() view returns(address)
func (_NewbieTask *NewbieTaskCaller) IntoEarlyBird(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "intoEarlyBird")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IntoEarlyBird is a free data retrieval call binding the contract method 0x46dbc4a9.
//
// Solidity: function intoEarlyBird() view returns(address)
func (_NewbieTask *NewbieTaskSession) IntoEarlyBird() (common.Address, error) {
	return _NewbieTask.Contract.IntoEarlyBird(&_NewbieTask.CallOpts)
}

// IntoEarlyBird is a free data retrieval call binding the contract method 0x46dbc4a9.
//
// Solidity: function intoEarlyBird() view returns(address)
func (_NewbieTask *NewbieTaskCallerSession) IntoEarlyBird() (common.Address, error) {
	return _NewbieTask.Contract.IntoEarlyBird(&_NewbieTask.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_NewbieTask *NewbieTaskCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_NewbieTask *NewbieTaskSession) IsAdmin(account common.Address) (bool, error) {
	return _NewbieTask.Contract.IsAdmin(&_NewbieTask.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_NewbieTask *NewbieTaskCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _NewbieTask.Contract.IsAdmin(&_NewbieTask.CallOpts, account)
}

// MaxReward is a free data retrieval call binding the contract method 0x66a78e6c.
//
// Solidity: function maxReward() view returns(uint256)
func (_NewbieTask *NewbieTaskCaller) MaxReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "maxReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReward is a free data retrieval call binding the contract method 0x66a78e6c.
//
// Solidity: function maxReward() view returns(uint256)
func (_NewbieTask *NewbieTaskSession) MaxReward() (*big.Int, error) {
	return _NewbieTask.Contract.MaxReward(&_NewbieTask.CallOpts)
}

// MaxReward is a free data retrieval call binding the contract method 0x66a78e6c.
//
// Solidity: function maxReward() view returns(uint256)
func (_NewbieTask *NewbieTaskCallerSession) MaxReward() (*big.Int, error) {
	return _NewbieTask.Contract.MaxReward(&_NewbieTask.CallOpts)
}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_NewbieTask *NewbieTaskCaller) PledgePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "pledgePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_NewbieTask *NewbieTaskSession) PledgePool() (common.Address, error) {
	return _NewbieTask.Contract.PledgePool(&_NewbieTask.CallOpts)
}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_NewbieTask *NewbieTaskCallerSession) PledgePool() (common.Address, error) {
	return _NewbieTask.Contract.PledgePool(&_NewbieTask.CallOpts)
}

// ServerAddress is a free data retrieval call binding the contract method 0xdb420fe3.
//
// Solidity: function serverAddress() view returns(address)
func (_NewbieTask *NewbieTaskCaller) ServerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "serverAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServerAddress is a free data retrieval call binding the contract method 0xdb420fe3.
//
// Solidity: function serverAddress() view returns(address)
func (_NewbieTask *NewbieTaskSession) ServerAddress() (common.Address, error) {
	return _NewbieTask.Contract.ServerAddress(&_NewbieTask.CallOpts)
}

// ServerAddress is a free data retrieval call binding the contract method 0xdb420fe3.
//
// Solidity: function serverAddress() view returns(address)
func (_NewbieTask *NewbieTaskCallerSession) ServerAddress() (common.Address, error) {
	return _NewbieTask.Contract.ServerAddress(&_NewbieTask.CallOpts)
}

// SocialMining is a free data retrieval call binding the contract method 0xe909868f.
//
// Solidity: function socialMining() view returns(address)
func (_NewbieTask *NewbieTaskCaller) SocialMining(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "socialMining")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SocialMining is a free data retrieval call binding the contract method 0xe909868f.
//
// Solidity: function socialMining() view returns(address)
func (_NewbieTask *NewbieTaskSession) SocialMining() (common.Address, error) {
	return _NewbieTask.Contract.SocialMining(&_NewbieTask.CallOpts)
}

// SocialMining is a free data retrieval call binding the contract method 0xe909868f.
//
// Solidity: function socialMining() view returns(address)
func (_NewbieTask *NewbieTaskCallerSession) SocialMining() (common.Address, error) {
	return _NewbieTask.Contract.SocialMining(&_NewbieTask.CallOpts)
}

// TaskReward is a free data retrieval call binding the contract method 0x710e2281.
//
// Solidity: function taskReward(uint256 ) view returns(uint256)
func (_NewbieTask *NewbieTaskCaller) TaskReward(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "taskReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskReward is a free data retrieval call binding the contract method 0x710e2281.
//
// Solidity: function taskReward(uint256 ) view returns(uint256)
func (_NewbieTask *NewbieTaskSession) TaskReward(arg0 *big.Int) (*big.Int, error) {
	return _NewbieTask.Contract.TaskReward(&_NewbieTask.CallOpts, arg0)
}

// TaskReward is a free data retrieval call binding the contract method 0x710e2281.
//
// Solidity: function taskReward(uint256 ) view returns(uint256)
func (_NewbieTask *NewbieTaskCallerSession) TaskReward(arg0 *big.Int) (*big.Int, error) {
	return _NewbieTask.Contract.TaskReward(&_NewbieTask.CallOpts, arg0)
}

// TaskTotalReward is a free data retrieval call binding the contract method 0x58b496d5.
//
// Solidity: function taskTotalReward() view returns(uint256)
func (_NewbieTask *NewbieTaskCaller) TaskTotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "taskTotalReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskTotalReward is a free data retrieval call binding the contract method 0x58b496d5.
//
// Solidity: function taskTotalReward() view returns(uint256)
func (_NewbieTask *NewbieTaskSession) TaskTotalReward() (*big.Int, error) {
	return _NewbieTask.Contract.TaskTotalReward(&_NewbieTask.CallOpts)
}

// TaskTotalReward is a free data retrieval call binding the contract method 0x58b496d5.
//
// Solidity: function taskTotalReward() view returns(uint256)
func (_NewbieTask *NewbieTaskCallerSession) TaskTotalReward() (*big.Int, error) {
	return _NewbieTask.Contract.TaskTotalReward(&_NewbieTask.CallOpts)
}

// TaskTotalRewardThreshold is a free data retrieval call binding the contract method 0x30339e00.
//
// Solidity: function taskTotalRewardThreshold() view returns(uint256)
func (_NewbieTask *NewbieTaskCaller) TaskTotalRewardThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "taskTotalRewardThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskTotalRewardThreshold is a free data retrieval call binding the contract method 0x30339e00.
//
// Solidity: function taskTotalRewardThreshold() view returns(uint256)
func (_NewbieTask *NewbieTaskSession) TaskTotalRewardThreshold() (*big.Int, error) {
	return _NewbieTask.Contract.TaskTotalRewardThreshold(&_NewbieTask.CallOpts)
}

// TaskTotalRewardThreshold is a free data retrieval call binding the contract method 0x30339e00.
//
// Solidity: function taskTotalRewardThreshold() view returns(uint256)
func (_NewbieTask *NewbieTaskCallerSession) TaskTotalRewardThreshold() (*big.Int, error) {
	return _NewbieTask.Contract.TaskTotalRewardThreshold(&_NewbieTask.CallOpts)
}

// Tox is a free data retrieval call binding the contract method 0x46c811fd.
//
// Solidity: function tox() view returns(address)
func (_NewbieTask *NewbieTaskCaller) Tox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "tox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tox is a free data retrieval call binding the contract method 0x46c811fd.
//
// Solidity: function tox() view returns(address)
func (_NewbieTask *NewbieTaskSession) Tox() (common.Address, error) {
	return _NewbieTask.Contract.Tox(&_NewbieTask.CallOpts)
}

// Tox is a free data retrieval call binding the contract method 0x46c811fd.
//
// Solidity: function tox() view returns(address)
func (_NewbieTask *NewbieTaskCallerSession) Tox() (common.Address, error) {
	return _NewbieTask.Contract.Tox(&_NewbieTask.CallOpts)
}

// UserReward is a free data retrieval call binding the contract method 0x000a74be.
//
// Solidity: function userReward(address ) view returns(uint256)
func (_NewbieTask *NewbieTaskCaller) UserReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "userReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserReward is a free data retrieval call binding the contract method 0x000a74be.
//
// Solidity: function userReward(address ) view returns(uint256)
func (_NewbieTask *NewbieTaskSession) UserReward(arg0 common.Address) (*big.Int, error) {
	return _NewbieTask.Contract.UserReward(&_NewbieTask.CallOpts, arg0)
}

// UserReward is a free data retrieval call binding the contract method 0x000a74be.
//
// Solidity: function userReward(address ) view returns(uint256)
func (_NewbieTask *NewbieTaskCallerSession) UserReward(arg0 common.Address) (*big.Int, error) {
	return _NewbieTask.Contract.UserReward(&_NewbieTask.CallOpts, arg0)
}

// WithdrawLimit is a free data retrieval call binding the contract method 0xf848d541.
//
// Solidity: function withdrawLimit() view returns(address)
func (_NewbieTask *NewbieTaskCaller) WithdrawLimit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NewbieTask.contract.Call(opts, &out, "withdrawLimit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawLimit is a free data retrieval call binding the contract method 0xf848d541.
//
// Solidity: function withdrawLimit() view returns(address)
func (_NewbieTask *NewbieTaskSession) WithdrawLimit() (common.Address, error) {
	return _NewbieTask.Contract.WithdrawLimit(&_NewbieTask.CallOpts)
}

// WithdrawLimit is a free data retrieval call binding the contract method 0xf848d541.
//
// Solidity: function withdrawLimit() view returns(address)
func (_NewbieTask *NewbieTaskCallerSession) WithdrawLimit() (common.Address, error) {
	return _NewbieTask.Contract.WithdrawLimit(&_NewbieTask.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_NewbieTask *NewbieTaskTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_NewbieTask *NewbieTaskSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AddAdmin(&_NewbieTask.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AddAdmin(&_NewbieTask.TransactOpts, account)
}

// AdminSetClaimType is a paid mutator transaction binding the contract method 0xba1c6a10.
//
// Solidity: function adminSetClaimType(uint256 claimType_, uint256 amount_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetClaimType(opts *bind.TransactOpts, claimType_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetClaimType", claimType_, amount_)
}

// AdminSetClaimType is a paid mutator transaction binding the contract method 0xba1c6a10.
//
// Solidity: function adminSetClaimType(uint256 claimType_, uint256 amount_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetClaimType(claimType_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetClaimType(&_NewbieTask.TransactOpts, claimType_, amount_)
}

// AdminSetClaimType is a paid mutator transaction binding the contract method 0xba1c6a10.
//
// Solidity: function adminSetClaimType(uint256 claimType_, uint256 amount_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetClaimType(claimType_ *big.Int, amount_ *big.Int) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetClaimType(&_NewbieTask.TransactOpts, claimType_, amount_)
}

// AdminSetFaceAuth is a paid mutator transaction binding the contract method 0xf5223a31.
//
// Solidity: function adminSetFaceAuth(address faceAuthAddress_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetFaceAuth(opts *bind.TransactOpts, faceAuthAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetFaceAuth", faceAuthAddress_)
}

// AdminSetFaceAuth is a paid mutator transaction binding the contract method 0xf5223a31.
//
// Solidity: function adminSetFaceAuth(address faceAuthAddress_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetFaceAuth(faceAuthAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetFaceAuth(&_NewbieTask.TransactOpts, faceAuthAddress_)
}

// AdminSetFaceAuth is a paid mutator transaction binding the contract method 0xf5223a31.
//
// Solidity: function adminSetFaceAuth(address faceAuthAddress_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetFaceAuth(faceAuthAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetFaceAuth(&_NewbieTask.TransactOpts, faceAuthAddress_)
}

// AdminSetIntoBinding is a paid mutator transaction binding the contract method 0x2550cd71.
//
// Solidity: function adminSetIntoBinding(address bindingAddress_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetIntoBinding(opts *bind.TransactOpts, bindingAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetIntoBinding", bindingAddress_)
}

// AdminSetIntoBinding is a paid mutator transaction binding the contract method 0x2550cd71.
//
// Solidity: function adminSetIntoBinding(address bindingAddress_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetIntoBinding(bindingAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetIntoBinding(&_NewbieTask.TransactOpts, bindingAddress_)
}

// AdminSetIntoBinding is a paid mutator transaction binding the contract method 0x2550cd71.
//
// Solidity: function adminSetIntoBinding(address bindingAddress_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetIntoBinding(bindingAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetIntoBinding(&_NewbieTask.TransactOpts, bindingAddress_)
}

// AdminSetIntoEarlyBird is a paid mutator transaction binding the contract method 0x130ddba9.
//
// Solidity: function adminSetIntoEarlyBird(address intoEarlyBirdAddress_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetIntoEarlyBird(opts *bind.TransactOpts, intoEarlyBirdAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetIntoEarlyBird", intoEarlyBirdAddress_)
}

// AdminSetIntoEarlyBird is a paid mutator transaction binding the contract method 0x130ddba9.
//
// Solidity: function adminSetIntoEarlyBird(address intoEarlyBirdAddress_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetIntoEarlyBird(intoEarlyBirdAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetIntoEarlyBird(&_NewbieTask.TransactOpts, intoEarlyBirdAddress_)
}

// AdminSetIntoEarlyBird is a paid mutator transaction binding the contract method 0x130ddba9.
//
// Solidity: function adminSetIntoEarlyBird(address intoEarlyBirdAddress_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetIntoEarlyBird(intoEarlyBirdAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetIntoEarlyBird(&_NewbieTask.TransactOpts, intoEarlyBirdAddress_)
}

// AdminSetIntoSocialMining is a paid mutator transaction binding the contract method 0x304fd019.
//
// Solidity: function adminSetIntoSocialMining(address socialMiningAddress_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetIntoSocialMining(opts *bind.TransactOpts, socialMiningAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetIntoSocialMining", socialMiningAddress_)
}

// AdminSetIntoSocialMining is a paid mutator transaction binding the contract method 0x304fd019.
//
// Solidity: function adminSetIntoSocialMining(address socialMiningAddress_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetIntoSocialMining(socialMiningAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetIntoSocialMining(&_NewbieTask.TransactOpts, socialMiningAddress_)
}

// AdminSetIntoSocialMining is a paid mutator transaction binding the contract method 0x304fd019.
//
// Solidity: function adminSetIntoSocialMining(address socialMiningAddress_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetIntoSocialMining(socialMiningAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetIntoSocialMining(&_NewbieTask.TransactOpts, socialMiningAddress_)
}

// AdminSetMaxReward is a paid mutator transaction binding the contract method 0x3816a4bd.
//
// Solidity: function adminSetMaxReward(uint256 maxReward_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetMaxReward(opts *bind.TransactOpts, maxReward_ *big.Int) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetMaxReward", maxReward_)
}

// AdminSetMaxReward is a paid mutator transaction binding the contract method 0x3816a4bd.
//
// Solidity: function adminSetMaxReward(uint256 maxReward_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetMaxReward(maxReward_ *big.Int) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetMaxReward(&_NewbieTask.TransactOpts, maxReward_)
}

// AdminSetMaxReward is a paid mutator transaction binding the contract method 0x3816a4bd.
//
// Solidity: function adminSetMaxReward(uint256 maxReward_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetMaxReward(maxReward_ *big.Int) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetMaxReward(&_NewbieTask.TransactOpts, maxReward_)
}

// AdminSetPledgePool is a paid mutator transaction binding the contract method 0xaf13e937.
//
// Solidity: function adminSetPledgePool(address pledgePoolAddress_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetPledgePool(opts *bind.TransactOpts, pledgePoolAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetPledgePool", pledgePoolAddress_)
}

// AdminSetPledgePool is a paid mutator transaction binding the contract method 0xaf13e937.
//
// Solidity: function adminSetPledgePool(address pledgePoolAddress_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetPledgePool(pledgePoolAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetPledgePool(&_NewbieTask.TransactOpts, pledgePoolAddress_)
}

// AdminSetPledgePool is a paid mutator transaction binding the contract method 0xaf13e937.
//
// Solidity: function adminSetPledgePool(address pledgePoolAddress_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetPledgePool(pledgePoolAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetPledgePool(&_NewbieTask.TransactOpts, pledgePoolAddress_)
}

// AdminSetServerAddress is a paid mutator transaction binding the contract method 0x3f662e56.
//
// Solidity: function adminSetServerAddress(address serverAddress_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetServerAddress(opts *bind.TransactOpts, serverAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetServerAddress", serverAddress_)
}

// AdminSetServerAddress is a paid mutator transaction binding the contract method 0x3f662e56.
//
// Solidity: function adminSetServerAddress(address serverAddress_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetServerAddress(serverAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetServerAddress(&_NewbieTask.TransactOpts, serverAddress_)
}

// AdminSetServerAddress is a paid mutator transaction binding the contract method 0x3f662e56.
//
// Solidity: function adminSetServerAddress(address serverAddress_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetServerAddress(serverAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetServerAddress(&_NewbieTask.TransactOpts, serverAddress_)
}

// AdminSetTaskTotalRewardThreshold is a paid mutator transaction binding the contract method 0x1aa8068c.
//
// Solidity: function adminSetTaskTotalRewardThreshold(uint256 taskTotalRewardThreshold_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetTaskTotalRewardThreshold(opts *bind.TransactOpts, taskTotalRewardThreshold_ *big.Int) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetTaskTotalRewardThreshold", taskTotalRewardThreshold_)
}

// AdminSetTaskTotalRewardThreshold is a paid mutator transaction binding the contract method 0x1aa8068c.
//
// Solidity: function adminSetTaskTotalRewardThreshold(uint256 taskTotalRewardThreshold_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetTaskTotalRewardThreshold(taskTotalRewardThreshold_ *big.Int) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetTaskTotalRewardThreshold(&_NewbieTask.TransactOpts, taskTotalRewardThreshold_)
}

// AdminSetTaskTotalRewardThreshold is a paid mutator transaction binding the contract method 0x1aa8068c.
//
// Solidity: function adminSetTaskTotalRewardThreshold(uint256 taskTotalRewardThreshold_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetTaskTotalRewardThreshold(taskTotalRewardThreshold_ *big.Int) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetTaskTotalRewardThreshold(&_NewbieTask.TransactOpts, taskTotalRewardThreshold_)
}

// AdminSetTox is a paid mutator transaction binding the contract method 0xb2fc941c.
//
// Solidity: function adminSetTox(address toxAddress_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetTox(opts *bind.TransactOpts, toxAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetTox", toxAddress_)
}

// AdminSetTox is a paid mutator transaction binding the contract method 0xb2fc941c.
//
// Solidity: function adminSetTox(address toxAddress_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetTox(toxAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetTox(&_NewbieTask.TransactOpts, toxAddress_)
}

// AdminSetTox is a paid mutator transaction binding the contract method 0xb2fc941c.
//
// Solidity: function adminSetTox(address toxAddress_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetTox(toxAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetTox(&_NewbieTask.TransactOpts, toxAddress_)
}

// AdminSetWithdrawLimit is a paid mutator transaction binding the contract method 0xee391e9d.
//
// Solidity: function adminSetWithdrawLimit(address withdrawLimitAddress_) returns()
func (_NewbieTask *NewbieTaskTransactor) AdminSetWithdrawLimit(opts *bind.TransactOpts, withdrawLimitAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "adminSetWithdrawLimit", withdrawLimitAddress_)
}

// AdminSetWithdrawLimit is a paid mutator transaction binding the contract method 0xee391e9d.
//
// Solidity: function adminSetWithdrawLimit(address withdrawLimitAddress_) returns()
func (_NewbieTask *NewbieTaskSession) AdminSetWithdrawLimit(withdrawLimitAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetWithdrawLimit(&_NewbieTask.TransactOpts, withdrawLimitAddress_)
}

// AdminSetWithdrawLimit is a paid mutator transaction binding the contract method 0xee391e9d.
//
// Solidity: function adminSetWithdrawLimit(address withdrawLimitAddress_) returns()
func (_NewbieTask *NewbieTaskTransactorSession) AdminSetWithdrawLimit(withdrawLimitAddress_ common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.AdminSetWithdrawLimit(&_NewbieTask.TransactOpts, withdrawLimitAddress_)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_NewbieTask *NewbieTaskTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_NewbieTask *NewbieTaskSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.BatchAddAdmin(&_NewbieTask.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_NewbieTask *NewbieTaskTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.BatchAddAdmin(&_NewbieTask.TransactOpts, amounts)
}

// Claim is a paid mutator transaction binding the contract method 0x7b91a85a.
//
// Solidity: function claim(bytes signature, address sender, uint256 chainId, uint256 claimType, bool taskStatus) returns()
func (_NewbieTask *NewbieTaskTransactor) Claim(opts *bind.TransactOpts, signature []byte, sender common.Address, chainId *big.Int, claimType *big.Int, taskStatus bool) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "claim", signature, sender, chainId, claimType, taskStatus)
}

// Claim is a paid mutator transaction binding the contract method 0x7b91a85a.
//
// Solidity: function claim(bytes signature, address sender, uint256 chainId, uint256 claimType, bool taskStatus) returns()
func (_NewbieTask *NewbieTaskSession) Claim(signature []byte, sender common.Address, chainId *big.Int, claimType *big.Int, taskStatus bool) (*types.Transaction, error) {
	return _NewbieTask.Contract.Claim(&_NewbieTask.TransactOpts, signature, sender, chainId, claimType, taskStatus)
}

// Claim is a paid mutator transaction binding the contract method 0x7b91a85a.
//
// Solidity: function claim(bytes signature, address sender, uint256 chainId, uint256 claimType, bool taskStatus) returns()
func (_NewbieTask *NewbieTaskTransactorSession) Claim(signature []byte, sender common.Address, chainId *big.Int, claimType *big.Int, taskStatus bool) (*types.Transaction, error) {
	return _NewbieTask.Contract.Claim(&_NewbieTask.TransactOpts, signature, sender, chainId, claimType, taskStatus)
}

// ClaimV2 is a paid mutator transaction binding the contract method 0x561e09f6.
//
// Solidity: function claimV2(bytes signature, address sender, uint256 chainId, uint256 claimType, bool taskStatus) returns()
func (_NewbieTask *NewbieTaskTransactor) ClaimV2(opts *bind.TransactOpts, signature []byte, sender common.Address, chainId *big.Int, claimType *big.Int, taskStatus bool) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "claimV2", signature, sender, chainId, claimType, taskStatus)
}

// ClaimV2 is a paid mutator transaction binding the contract method 0x561e09f6.
//
// Solidity: function claimV2(bytes signature, address sender, uint256 chainId, uint256 claimType, bool taskStatus) returns()
func (_NewbieTask *NewbieTaskSession) ClaimV2(signature []byte, sender common.Address, chainId *big.Int, claimType *big.Int, taskStatus bool) (*types.Transaction, error) {
	return _NewbieTask.Contract.ClaimV2(&_NewbieTask.TransactOpts, signature, sender, chainId, claimType, taskStatus)
}

// ClaimV2 is a paid mutator transaction binding the contract method 0x561e09f6.
//
// Solidity: function claimV2(bytes signature, address sender, uint256 chainId, uint256 claimType, bool taskStatus) returns()
func (_NewbieTask *NewbieTaskTransactorSession) ClaimV2(signature []byte, sender common.Address, chainId *big.Int, claimType *big.Int, taskStatus bool) (*types.Transaction, error) {
	return _NewbieTask.Contract.ClaimV2(&_NewbieTask.TransactOpts, signature, sender, chainId, claimType, taskStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_NewbieTask *NewbieTaskTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_NewbieTask *NewbieTaskSession) Initialize() (*types.Transaction, error) {
	return _NewbieTask.Contract.Initialize(&_NewbieTask.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_NewbieTask *NewbieTaskTransactorSession) Initialize() (*types.Transaction, error) {
	return _NewbieTask.Contract.Initialize(&_NewbieTask.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_NewbieTask *NewbieTaskTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_NewbieTask *NewbieTaskSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.RemoveAdmin(&_NewbieTask.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_NewbieTask *NewbieTaskTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _NewbieTask.Contract.RemoveAdmin(&_NewbieTask.TransactOpts, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_NewbieTask *NewbieTaskTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NewbieTask.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_NewbieTask *NewbieTaskSession) RenounceAdmin() (*types.Transaction, error) {
	return _NewbieTask.Contract.RenounceAdmin(&_NewbieTask.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_NewbieTask *NewbieTaskTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _NewbieTask.Contract.RenounceAdmin(&_NewbieTask.TransactOpts)
}

// NewbieTaskAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the NewbieTask contract.
type NewbieTaskAdminAddedIterator struct {
	Event *NewbieTaskAdminAdded // Event containing the contract specifics and raw log

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
func (it *NewbieTaskAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewbieTaskAdminAdded)
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
		it.Event = new(NewbieTaskAdminAdded)
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
func (it *NewbieTaskAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewbieTaskAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewbieTaskAdminAdded represents a AdminAdded event raised by the NewbieTask contract.
type NewbieTaskAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address account)
func (_NewbieTask *NewbieTaskFilterer) FilterAdminAdded(opts *bind.FilterOpts) (*NewbieTaskAdminAddedIterator, error) {

	logs, sub, err := _NewbieTask.contract.FilterLogs(opts, "AdminAdded")
	if err != nil {
		return nil, err
	}
	return &NewbieTaskAdminAddedIterator{contract: _NewbieTask.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address account)
func (_NewbieTask *NewbieTaskFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *NewbieTaskAdminAdded) (event.Subscription, error) {

	logs, sub, err := _NewbieTask.contract.WatchLogs(opts, "AdminAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewbieTaskAdminAdded)
				if err := _NewbieTask.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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
func (_NewbieTask *NewbieTaskFilterer) ParseAdminAdded(log types.Log) (*NewbieTaskAdminAdded, error) {
	event := new(NewbieTaskAdminAdded)
	if err := _NewbieTask.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NewbieTaskAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the NewbieTask contract.
type NewbieTaskAdminRemovedIterator struct {
	Event *NewbieTaskAdminRemoved // Event containing the contract specifics and raw log

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
func (it *NewbieTaskAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewbieTaskAdminRemoved)
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
		it.Event = new(NewbieTaskAdminRemoved)
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
func (it *NewbieTaskAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewbieTaskAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewbieTaskAdminRemoved represents a AdminRemoved event raised by the NewbieTask contract.
type NewbieTaskAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address account)
func (_NewbieTask *NewbieTaskFilterer) FilterAdminRemoved(opts *bind.FilterOpts) (*NewbieTaskAdminRemovedIterator, error) {

	logs, sub, err := _NewbieTask.contract.FilterLogs(opts, "AdminRemoved")
	if err != nil {
		return nil, err
	}
	return &NewbieTaskAdminRemovedIterator{contract: _NewbieTask.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address account)
func (_NewbieTask *NewbieTaskFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *NewbieTaskAdminRemoved) (event.Subscription, error) {

	logs, sub, err := _NewbieTask.contract.WatchLogs(opts, "AdminRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewbieTaskAdminRemoved)
				if err := _NewbieTask.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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
func (_NewbieTask *NewbieTaskFilterer) ParseAdminRemoved(log types.Log) (*NewbieTaskAdminRemoved, error) {
	event := new(NewbieTaskAdminRemoved)
	if err := _NewbieTask.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NewbieTaskClaimLogIterator is returned from FilterClaimLog and is used to iterate over the raw logs and unpacked data for ClaimLog events raised by the NewbieTask contract.
type NewbieTaskClaimLogIterator struct {
	Event *NewbieTaskClaimLog // Event containing the contract specifics and raw log

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
func (it *NewbieTaskClaimLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewbieTaskClaimLog)
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
		it.Event = new(NewbieTaskClaimLog)
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
func (it *NewbieTaskClaimLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewbieTaskClaimLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewbieTaskClaimLog represents a ClaimLog event raised by the NewbieTask contract.
type NewbieTaskClaimLog struct {
	User      common.Address
	Amount    *big.Int
	ClaimType *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimLog is a free log retrieval operation binding the contract event 0x33269a4e92a2210f809a73cf4d7a0bcb90717755a6b93cf03c75a0ed1080d95e.
//
// Solidity: event ClaimLog(address user, uint256 amount, uint256 claimType)
func (_NewbieTask *NewbieTaskFilterer) FilterClaimLog(opts *bind.FilterOpts) (*NewbieTaskClaimLogIterator, error) {

	logs, sub, err := _NewbieTask.contract.FilterLogs(opts, "ClaimLog")
	if err != nil {
		return nil, err
	}
	return &NewbieTaskClaimLogIterator{contract: _NewbieTask.contract, event: "ClaimLog", logs: logs, sub: sub}, nil
}

// WatchClaimLog is a free log subscription operation binding the contract event 0x33269a4e92a2210f809a73cf4d7a0bcb90717755a6b93cf03c75a0ed1080d95e.
//
// Solidity: event ClaimLog(address user, uint256 amount, uint256 claimType)
func (_NewbieTask *NewbieTaskFilterer) WatchClaimLog(opts *bind.WatchOpts, sink chan<- *NewbieTaskClaimLog) (event.Subscription, error) {

	logs, sub, err := _NewbieTask.contract.WatchLogs(opts, "ClaimLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewbieTaskClaimLog)
				if err := _NewbieTask.contract.UnpackLog(event, "ClaimLog", log); err != nil {
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

// ParseClaimLog is a log parse operation binding the contract event 0x33269a4e92a2210f809a73cf4d7a0bcb90717755a6b93cf03c75a0ed1080d95e.
//
// Solidity: event ClaimLog(address user, uint256 amount, uint256 claimType)
func (_NewbieTask *NewbieTaskFilterer) ParseClaimLog(log types.Log) (*NewbieTaskClaimLog, error) {
	event := new(NewbieTaskClaimLog)
	if err := _NewbieTask.contract.UnpackLog(event, "ClaimLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NewbieTaskInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NewbieTask contract.
type NewbieTaskInitializedIterator struct {
	Event *NewbieTaskInitialized // Event containing the contract specifics and raw log

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
func (it *NewbieTaskInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewbieTaskInitialized)
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
		it.Event = new(NewbieTaskInitialized)
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
func (it *NewbieTaskInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewbieTaskInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewbieTaskInitialized represents a Initialized event raised by the NewbieTask contract.
type NewbieTaskInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NewbieTask *NewbieTaskFilterer) FilterInitialized(opts *bind.FilterOpts) (*NewbieTaskInitializedIterator, error) {

	logs, sub, err := _NewbieTask.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NewbieTaskInitializedIterator{contract: _NewbieTask.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NewbieTask *NewbieTaskFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NewbieTaskInitialized) (event.Subscription, error) {

	logs, sub, err := _NewbieTask.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewbieTaskInitialized)
				if err := _NewbieTask.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NewbieTask *NewbieTaskFilterer) ParseInitialized(log types.Log) (*NewbieTaskInitialized, error) {
	event := new(NewbieTaskInitialized)
	if err := _NewbieTask.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
