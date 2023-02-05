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

// FarmAccountingInfo is an auto generated low-level Go binding around an user-defined struct.
type FarmAccountingInfo struct {
	Finished *big.Int
	Duration uint32
	Reward   *big.Int
}

// RewardsMetaData contains all meta data concerning the Rewards contract.
var RewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20Pods\",\"name\":\"farmableToken_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardsTokensLimit_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DurationTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutputArrayTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardsTokenAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardsTokenNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardsTokensLimitReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"RewardsTokensLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameDistributor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFarmableTokenAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRewardsTokenAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldDistributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDistributor\",\"type\":\"address\"}],\"name\":\"DistributorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"FarmCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"}],\"name\":\"addRewardsToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"rewardsToken\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"rewardsToken\",\"type\":\"address\"}],\"name\":\"farmInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"finished\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"internalType\":\"uint184\",\"name\":\"reward\",\"type\":\"uint184\"}],\"internalType\":\"structFarmAccounting.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"farmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsTokensLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor_\",\"type\":\"address\"}],\"name\":\"setDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"startFarming\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Pods\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateBalances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardsMetaData.ABI instead.
var RewardsABI = RewardsMetaData.ABI

// Rewards is an auto generated Go binding around an Ethereum contract.
type Rewards struct {
	RewardsCaller     // Read-only binding to the contract
	RewardsTransactor // Write-only binding to the contract
	RewardsFilterer   // Log filterer for contract events
}

// RewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardsSession struct {
	Contract     *Rewards          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardsCallerSession struct {
	Contract *RewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardsTransactorSession struct {
	Contract     *RewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardsRaw struct {
	Contract *Rewards // Generic contract binding to access the raw methods on
}

// RewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardsCallerRaw struct {
	Contract *RewardsCaller // Generic read-only contract binding to access the raw methods on
}

// RewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardsTransactorRaw struct {
	Contract *RewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewards creates a new instance of Rewards, bound to a specific deployed contract.
func NewRewards(address common.Address, backend bind.ContractBackend) (*Rewards, error) {
	contract, err := bindRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewards{RewardsCaller: RewardsCaller{contract: contract}, RewardsTransactor: RewardsTransactor{contract: contract}, RewardsFilterer: RewardsFilterer{contract: contract}}, nil
}

// NewRewardsCaller creates a new read-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsCaller(address common.Address, caller bind.ContractCaller) (*RewardsCaller, error) {
	contract, err := bindRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsCaller{contract: contract}, nil
}

// NewRewardsTransactor creates a new write-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardsTransactor, error) {
	contract, err := bindRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsTransactor{contract: contract}, nil
}

// NewRewardsFilterer creates a new log filterer instance of Rewards, bound to a specific deployed contract.
func NewRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardsFilterer, error) {
	contract, err := bindRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardsFilterer{contract: contract}, nil
}

// bindRewards binds a generic wrapper to an already deployed contract.
func bindRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.RewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transact(opts, method, params...)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Rewards *RewardsCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Rewards *RewardsSession) Distributor() (common.Address, error) {
	return _Rewards.Contract.Distributor(&_Rewards.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Rewards *RewardsCallerSession) Distributor() (common.Address, error) {
	return _Rewards.Contract.Distributor(&_Rewards.CallOpts)
}

// FarmInfo is a free data retrieval call binding the contract method 0x2915cbec.
//
// Solidity: function farmInfo(address rewardsToken) view returns((uint40,uint32,uint184))
func (_Rewards *RewardsCaller) FarmInfo(opts *bind.CallOpts, rewardsToken common.Address) (FarmAccountingInfo, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "farmInfo", rewardsToken)

	if err != nil {
		return *new(FarmAccountingInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(FarmAccountingInfo)).(*FarmAccountingInfo)

	return out0, err

}

// FarmInfo is a free data retrieval call binding the contract method 0x2915cbec.
//
// Solidity: function farmInfo(address rewardsToken) view returns((uint40,uint32,uint184))
func (_Rewards *RewardsSession) FarmInfo(rewardsToken common.Address) (FarmAccountingInfo, error) {
	return _Rewards.Contract.FarmInfo(&_Rewards.CallOpts, rewardsToken)
}

// FarmInfo is a free data retrieval call binding the contract method 0x2915cbec.
//
// Solidity: function farmInfo(address rewardsToken) view returns((uint40,uint32,uint184))
func (_Rewards *RewardsCallerSession) FarmInfo(rewardsToken common.Address) (FarmAccountingInfo, error) {
	return _Rewards.Contract.FarmInfo(&_Rewards.CallOpts, rewardsToken)
}

// Farmed is a free data retrieval call binding the contract method 0xb1bd3517.
//
// Solidity: function farmed(address rewardsToken, address account) view returns(uint256)
func (_Rewards *RewardsCaller) Farmed(opts *bind.CallOpts, rewardsToken common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "farmed", rewardsToken, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Farmed is a free data retrieval call binding the contract method 0xb1bd3517.
//
// Solidity: function farmed(address rewardsToken, address account) view returns(uint256)
func (_Rewards *RewardsSession) Farmed(rewardsToken common.Address, account common.Address) (*big.Int, error) {
	return _Rewards.Contract.Farmed(&_Rewards.CallOpts, rewardsToken, account)
}

// Farmed is a free data retrieval call binding the contract method 0xb1bd3517.
//
// Solidity: function farmed(address rewardsToken, address account) view returns(uint256)
func (_Rewards *RewardsCallerSession) Farmed(rewardsToken common.Address, account common.Address) (*big.Int, error) {
	return _Rewards.Contract.Farmed(&_Rewards.CallOpts, rewardsToken, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsSession) Owner() (common.Address, error) {
	return _Rewards.Contract.Owner(&_Rewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsCallerSession) Owner() (common.Address, error) {
	return _Rewards.Contract.Owner(&_Rewards.CallOpts)
}

// RewardsTokens is a free data retrieval call binding the contract method 0xcd00671b.
//
// Solidity: function rewardsTokens() view returns(address[])
func (_Rewards *RewardsCaller) RewardsTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "rewardsTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RewardsTokens is a free data retrieval call binding the contract method 0xcd00671b.
//
// Solidity: function rewardsTokens() view returns(address[])
func (_Rewards *RewardsSession) RewardsTokens() ([]common.Address, error) {
	return _Rewards.Contract.RewardsTokens(&_Rewards.CallOpts)
}

// RewardsTokens is a free data retrieval call binding the contract method 0xcd00671b.
//
// Solidity: function rewardsTokens() view returns(address[])
func (_Rewards *RewardsCallerSession) RewardsTokens() ([]common.Address, error) {
	return _Rewards.Contract.RewardsTokens(&_Rewards.CallOpts)
}

// RewardsTokensLimit is a free data retrieval call binding the contract method 0x3717a689.
//
// Solidity: function rewardsTokensLimit() view returns(uint256)
func (_Rewards *RewardsCaller) RewardsTokensLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "rewardsTokensLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsTokensLimit is a free data retrieval call binding the contract method 0x3717a689.
//
// Solidity: function rewardsTokensLimit() view returns(uint256)
func (_Rewards *RewardsSession) RewardsTokensLimit() (*big.Int, error) {
	return _Rewards.Contract.RewardsTokensLimit(&_Rewards.CallOpts)
}

// RewardsTokensLimit is a free data retrieval call binding the contract method 0x3717a689.
//
// Solidity: function rewardsTokensLimit() view returns(uint256)
func (_Rewards *RewardsCallerSession) RewardsTokensLimit() (*big.Int, error) {
	return _Rewards.Contract.RewardsTokensLimit(&_Rewards.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Rewards *RewardsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Rewards *RewardsSession) Token() (common.Address, error) {
	return _Rewards.Contract.Token(&_Rewards.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Rewards *RewardsCallerSession) Token() (common.Address, error) {
	return _Rewards.Contract.Token(&_Rewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Rewards *RewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Rewards *RewardsSession) TotalSupply() (*big.Int, error) {
	return _Rewards.Contract.TotalSupply(&_Rewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Rewards *RewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _Rewards.Contract.TotalSupply(&_Rewards.CallOpts)
}

// AddRewardsToken is a paid mutator transaction binding the contract method 0x66fa3e1e.
//
// Solidity: function addRewardsToken(address rewardsToken) returns()
func (_Rewards *RewardsTransactor) AddRewardsToken(opts *bind.TransactOpts, rewardsToken common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "addRewardsToken", rewardsToken)
}

// AddRewardsToken is a paid mutator transaction binding the contract method 0x66fa3e1e.
//
// Solidity: function addRewardsToken(address rewardsToken) returns()
func (_Rewards *RewardsSession) AddRewardsToken(rewardsToken common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.AddRewardsToken(&_Rewards.TransactOpts, rewardsToken)
}

// AddRewardsToken is a paid mutator transaction binding the contract method 0x66fa3e1e.
//
// Solidity: function addRewardsToken(address rewardsToken) returns()
func (_Rewards *RewardsTransactorSession) AddRewardsToken(rewardsToken common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.AddRewardsToken(&_Rewards.TransactOpts, rewardsToken)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address rewardsToken) returns()
func (_Rewards *RewardsTransactor) Claim(opts *bind.TransactOpts, rewardsToken common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "claim", rewardsToken)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address rewardsToken) returns()
func (_Rewards *RewardsSession) Claim(rewardsToken common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.Claim(&_Rewards.TransactOpts, rewardsToken)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address rewardsToken) returns()
func (_Rewards *RewardsTransactorSession) Claim(rewardsToken common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.Claim(&_Rewards.TransactOpts, rewardsToken)
}

// Claim0 is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Rewards *RewardsTransactor) Claim0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "claim0")
}

// Claim0 is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Rewards *RewardsSession) Claim0() (*types.Transaction, error) {
	return _Rewards.Contract.Claim0(&_Rewards.TransactOpts)
}

// Claim0 is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Rewards *RewardsTransactorSession) Claim0() (*types.Transaction, error) {
	return _Rewards.Contract.Claim0(&_Rewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewards.Contract.RenounceOwnership(&_Rewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewards.Contract.RenounceOwnership(&_Rewards.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Rewards *RewardsTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Rewards *RewardsSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.RescueFunds(&_Rewards.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Rewards *RewardsTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.RescueFunds(&_Rewards.TransactOpts, token, amount)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address distributor_) returns()
func (_Rewards *RewardsTransactor) SetDistributor(opts *bind.TransactOpts, distributor_ common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setDistributor", distributor_)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address distributor_) returns()
func (_Rewards *RewardsSession) SetDistributor(distributor_ common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetDistributor(&_Rewards.TransactOpts, distributor_)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address distributor_) returns()
func (_Rewards *RewardsTransactorSession) SetDistributor(distributor_ common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetDistributor(&_Rewards.TransactOpts, distributor_)
}

// StartFarming is a paid mutator transaction binding the contract method 0xe6235ba7.
//
// Solidity: function startFarming(address rewardsToken, uint256 amount, uint256 period) returns()
func (_Rewards *RewardsTransactor) StartFarming(opts *bind.TransactOpts, rewardsToken common.Address, amount *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "startFarming", rewardsToken, amount, period)
}

// StartFarming is a paid mutator transaction binding the contract method 0xe6235ba7.
//
// Solidity: function startFarming(address rewardsToken, uint256 amount, uint256 period) returns()
func (_Rewards *RewardsSession) StartFarming(rewardsToken common.Address, amount *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.StartFarming(&_Rewards.TransactOpts, rewardsToken, amount, period)
}

// StartFarming is a paid mutator transaction binding the contract method 0xe6235ba7.
//
// Solidity: function startFarming(address rewardsToken, uint256 amount, uint256 period) returns()
func (_Rewards *RewardsTransactorSession) StartFarming(rewardsToken common.Address, amount *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.StartFarming(&_Rewards.TransactOpts, rewardsToken, amount, period)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.TransferOwnership(&_Rewards.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.TransferOwnership(&_Rewards.TransactOpts, newOwner)
}

// UpdateBalances is a paid mutator transaction binding the contract method 0x7a1f1aa9.
//
// Solidity: function updateBalances(address from, address to, uint256 amount) returns()
func (_Rewards *RewardsTransactor) UpdateBalances(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "updateBalances", from, to, amount)
}

// UpdateBalances is a paid mutator transaction binding the contract method 0x7a1f1aa9.
//
// Solidity: function updateBalances(address from, address to, uint256 amount) returns()
func (_Rewards *RewardsSession) UpdateBalances(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.UpdateBalances(&_Rewards.TransactOpts, from, to, amount)
}

// UpdateBalances is a paid mutator transaction binding the contract method 0x7a1f1aa9.
//
// Solidity: function updateBalances(address from, address to, uint256 amount) returns()
func (_Rewards *RewardsTransactorSession) UpdateBalances(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.UpdateBalances(&_Rewards.TransactOpts, from, to, amount)
}

// RewardsDistributorChangedIterator is returned from FilterDistributorChanged and is used to iterate over the raw logs and unpacked data for DistributorChanged events raised by the Rewards contract.
type RewardsDistributorChangedIterator struct {
	Event *RewardsDistributorChanged // Event containing the contract specifics and raw log

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
func (it *RewardsDistributorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsDistributorChanged)
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
		it.Event = new(RewardsDistributorChanged)
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
func (it *RewardsDistributorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsDistributorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsDistributorChanged represents a DistributorChanged event raised by the Rewards contract.
type RewardsDistributorChanged struct {
	OldDistributor common.Address
	NewDistributor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributorChanged is a free log retrieval operation binding the contract event 0xa9f739537fc57540bed0a44e33e27baa63290d865cc15f0f16cf17d38c998a4d.
//
// Solidity: event DistributorChanged(address oldDistributor, address newDistributor)
func (_Rewards *RewardsFilterer) FilterDistributorChanged(opts *bind.FilterOpts) (*RewardsDistributorChangedIterator, error) {

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "DistributorChanged")
	if err != nil {
		return nil, err
	}
	return &RewardsDistributorChangedIterator{contract: _Rewards.contract, event: "DistributorChanged", logs: logs, sub: sub}, nil
}

// WatchDistributorChanged is a free log subscription operation binding the contract event 0xa9f739537fc57540bed0a44e33e27baa63290d865cc15f0f16cf17d38c998a4d.
//
// Solidity: event DistributorChanged(address oldDistributor, address newDistributor)
func (_Rewards *RewardsFilterer) WatchDistributorChanged(opts *bind.WatchOpts, sink chan<- *RewardsDistributorChanged) (event.Subscription, error) {

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "DistributorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsDistributorChanged)
				if err := _Rewards.contract.UnpackLog(event, "DistributorChanged", log); err != nil {
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

// ParseDistributorChanged is a log parse operation binding the contract event 0xa9f739537fc57540bed0a44e33e27baa63290d865cc15f0f16cf17d38c998a4d.
//
// Solidity: event DistributorChanged(address oldDistributor, address newDistributor)
func (_Rewards *RewardsFilterer) ParseDistributorChanged(log types.Log) (*RewardsDistributorChanged, error) {
	event := new(RewardsDistributorChanged)
	if err := _Rewards.contract.UnpackLog(event, "DistributorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsFarmCreatedIterator is returned from FilterFarmCreated and is used to iterate over the raw logs and unpacked data for FarmCreated events raised by the Rewards contract.
type RewardsFarmCreatedIterator struct {
	Event *RewardsFarmCreated // Event containing the contract specifics and raw log

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
func (it *RewardsFarmCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsFarmCreated)
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
		it.Event = new(RewardsFarmCreated)
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
func (it *RewardsFarmCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsFarmCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsFarmCreated represents a FarmCreated event raised by the Rewards contract.
type RewardsFarmCreated struct {
	Token  common.Address
	Reward common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFarmCreated is a free log retrieval operation binding the contract event 0x6bff9ddd187ef283e9c7726f406ab27bcc3719a41b6bee3585c7447183cffcec.
//
// Solidity: event FarmCreated(address token, address reward)
func (_Rewards *RewardsFilterer) FilterFarmCreated(opts *bind.FilterOpts) (*RewardsFarmCreatedIterator, error) {

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "FarmCreated")
	if err != nil {
		return nil, err
	}
	return &RewardsFarmCreatedIterator{contract: _Rewards.contract, event: "FarmCreated", logs: logs, sub: sub}, nil
}

// WatchFarmCreated is a free log subscription operation binding the contract event 0x6bff9ddd187ef283e9c7726f406ab27bcc3719a41b6bee3585c7447183cffcec.
//
// Solidity: event FarmCreated(address token, address reward)
func (_Rewards *RewardsFilterer) WatchFarmCreated(opts *bind.WatchOpts, sink chan<- *RewardsFarmCreated) (event.Subscription, error) {

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "FarmCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsFarmCreated)
				if err := _Rewards.contract.UnpackLog(event, "FarmCreated", log); err != nil {
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

// ParseFarmCreated is a log parse operation binding the contract event 0x6bff9ddd187ef283e9c7726f406ab27bcc3719a41b6bee3585c7447183cffcec.
//
// Solidity: event FarmCreated(address token, address reward)
func (_Rewards *RewardsFilterer) ParseFarmCreated(log types.Log) (*RewardsFarmCreated, error) {
	event := new(RewardsFarmCreated)
	if err := _Rewards.contract.UnpackLog(event, "FarmCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rewards contract.
type RewardsOwnershipTransferredIterator struct {
	Event *RewardsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewardsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsOwnershipTransferred)
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
		it.Event = new(RewardsOwnershipTransferred)
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
func (it *RewardsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsOwnershipTransferred represents a OwnershipTransferred event raised by the Rewards contract.
type RewardsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewardsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewardsOwnershipTransferredIterator{contract: _Rewards.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewardsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsOwnershipTransferred)
				if err := _Rewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) ParseOwnershipTransferred(log types.Log) (*RewardsOwnershipTransferred, error) {
	event := new(RewardsOwnershipTransferred)
	if err := _Rewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the Rewards contract.
type RewardsRewardAddedIterator struct {
	Event *RewardsRewardAdded // Event containing the contract specifics and raw log

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
func (it *RewardsRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsRewardAdded)
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
		it.Event = new(RewardsRewardAdded)
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
func (it *RewardsRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsRewardAdded represents a RewardAdded event raised by the Rewards contract.
type RewardsRewardAdded struct {
	Token    common.Address
	Reward   *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0x6a6f77044107a33658235d41bedbbaf2fe9ccdceb313143c947a5e76e1ec8474.
//
// Solidity: event RewardAdded(address token, uint256 reward, uint256 duration)
func (_Rewards *RewardsFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*RewardsRewardAddedIterator, error) {

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &RewardsRewardAddedIterator{contract: _Rewards.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0x6a6f77044107a33658235d41bedbbaf2fe9ccdceb313143c947a5e76e1ec8474.
//
// Solidity: event RewardAdded(address token, uint256 reward, uint256 duration)
func (_Rewards *RewardsFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *RewardsRewardAdded) (event.Subscription, error) {

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsRewardAdded)
				if err := _Rewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0x6a6f77044107a33658235d41bedbbaf2fe9ccdceb313143c947a5e76e1ec8474.
//
// Solidity: event RewardAdded(address token, uint256 reward, uint256 duration)
func (_Rewards *RewardsFilterer) ParseRewardAdded(log types.Log) (*RewardsRewardAdded, error) {
	event := new(RewardsRewardAdded)
	if err := _Rewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
