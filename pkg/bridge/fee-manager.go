// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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
)

// FeeManagerMetaData contains all meta data concerning the FeeManager contract.
var FeeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_foundationAddress\",\"type\":\"address\"}],\"name\":\"FoundationAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_liquidityPools\",\"type\":\"address\"}],\"name\":\"LiquidityPoolsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorFee\",\"type\":\"address\"}],\"name\":\"ValidatorFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorRefundFee\",\"type\":\"uint256\"}],\"name\":\"ValidatorRefundFeeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"foundationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerStorage\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_liquidityPools\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_foundationAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_validatorFeePool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validatorRefundFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityPools\",\"outputs\":[{\"internalType\":\"contractLiquidityPools\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidityRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_foundationAddress\",\"type\":\"address\"}],\"name\":\"setFoundationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_liquidityPools\",\"type\":\"address\"}],\"name\":\"setLiquidityPools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityReward\",\"type\":\"uint256\"}],\"name\":\"setTokenFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_validatorFee\",\"type\":\"address\"}],\"name\":\"setValidatorFeePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorRefundFee\",\"type\":\"uint256\"}],\"name\":\"setValidatorRefundFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerStorage\",\"outputs\":[{\"internalType\":\"contractSignerStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorFeePool\",\"outputs\":[{\"internalType\":\"contractBridgeValidatorFeePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorRefundFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FeeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeManagerMetaData.ABI instead.
var FeeManagerABI = FeeManagerMetaData.ABI

// FeeManager is an auto generated Go binding around an Ethereum contract.
type FeeManager struct {
	FeeManagerCaller     // Read-only binding to the contract
	FeeManagerTransactor // Write-only binding to the contract
	FeeManagerFilterer   // Log filterer for contract events
}

// FeeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeManagerSession struct {
	Contract     *FeeManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeManagerCallerSession struct {
	Contract *FeeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FeeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeManagerTransactorSession struct {
	Contract     *FeeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FeeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeManagerRaw struct {
	Contract *FeeManager // Generic contract binding to access the raw methods on
}

// FeeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeManagerCallerRaw struct {
	Contract *FeeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FeeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeManagerTransactorRaw struct {
	Contract *FeeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeManager creates a new instance of FeeManager, bound to a specific deployed contract.
func NewFeeManager(address common.Address, backend bind.ContractBackend) (*FeeManager, error) {
	contract, err := bindFeeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeManager{FeeManagerCaller: FeeManagerCaller{contract: contract}, FeeManagerTransactor: FeeManagerTransactor{contract: contract}, FeeManagerFilterer: FeeManagerFilterer{contract: contract}}, nil
}

// NewFeeManagerCaller creates a new read-only instance of FeeManager, bound to a specific deployed contract.
func NewFeeManagerCaller(address common.Address, caller bind.ContractCaller) (*FeeManagerCaller, error) {
	contract, err := bindFeeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeManagerCaller{contract: contract}, nil
}

// NewFeeManagerTransactor creates a new write-only instance of FeeManager, bound to a specific deployed contract.
func NewFeeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeManagerTransactor, error) {
	contract, err := bindFeeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeManagerTransactor{contract: contract}, nil
}

// NewFeeManagerFilterer creates a new log filterer instance of FeeManager, bound to a specific deployed contract.
func NewFeeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeManagerFilterer, error) {
	contract, err := bindFeeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeManagerFilterer{contract: contract}, nil
}

// bindFeeManager binds a generic wrapper to an already deployed contract.
func bindFeeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeManager *FeeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeManager.Contract.FeeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeManager *FeeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.Contract.FeeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeManager *FeeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeManager.Contract.FeeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeManager *FeeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeManager *FeeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeManager *FeeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeManager.Contract.contract.Transact(opts, method, params...)
}

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 amount) view returns(uint256 fee)
func (_FeeManager *FeeManagerCaller) CalculateFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "calculateFee", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 amount) view returns(uint256 fee)
func (_FeeManager *FeeManagerSession) CalculateFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _FeeManager.Contract.CalculateFee(&_FeeManager.CallOpts, token, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x8b28ab1e.
//
// Solidity: function calculateFee(address token, uint256 amount) view returns(uint256 fee)
func (_FeeManager *FeeManagerCallerSession) CalculateFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _FeeManager.Contract.CalculateFee(&_FeeManager.CallOpts, token, amount)
}

// FoundationAddress is a free data retrieval call binding the contract method 0xfcf07c6b.
//
// Solidity: function foundationAddress() view returns(address)
func (_FeeManager *FeeManagerCaller) FoundationAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "foundationAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FoundationAddress is a free data retrieval call binding the contract method 0xfcf07c6b.
//
// Solidity: function foundationAddress() view returns(address)
func (_FeeManager *FeeManagerSession) FoundationAddress() (common.Address, error) {
	return _FeeManager.Contract.FoundationAddress(&_FeeManager.CallOpts)
}

// FoundationAddress is a free data retrieval call binding the contract method 0xfcf07c6b.
//
// Solidity: function foundationAddress() view returns(address)
func (_FeeManager *FeeManagerCallerSession) FoundationAddress() (common.Address, error) {
	return _FeeManager.Contract.FoundationAddress(&_FeeManager.CallOpts)
}

// LiquidityPools is a free data retrieval call binding the contract method 0xd6efd7c3.
//
// Solidity: function liquidityPools() view returns(address)
func (_FeeManager *FeeManagerCaller) LiquidityPools(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "liquidityPools")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityPools is a free data retrieval call binding the contract method 0xd6efd7c3.
//
// Solidity: function liquidityPools() view returns(address)
func (_FeeManager *FeeManagerSession) LiquidityPools() (common.Address, error) {
	return _FeeManager.Contract.LiquidityPools(&_FeeManager.CallOpts)
}

// LiquidityPools is a free data retrieval call binding the contract method 0xd6efd7c3.
//
// Solidity: function liquidityPools() view returns(address)
func (_FeeManager *FeeManagerCallerSession) LiquidityPools() (common.Address, error) {
	return _FeeManager.Contract.LiquidityPools(&_FeeManager.CallOpts)
}

// LiquidityRewardPercentage is a free data retrieval call binding the contract method 0x1f71680b.
//
// Solidity: function liquidityRewardPercentage(address ) view returns(uint256)
func (_FeeManager *FeeManagerCaller) LiquidityRewardPercentage(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "liquidityRewardPercentage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityRewardPercentage is a free data retrieval call binding the contract method 0x1f71680b.
//
// Solidity: function liquidityRewardPercentage(address ) view returns(uint256)
func (_FeeManager *FeeManagerSession) LiquidityRewardPercentage(arg0 common.Address) (*big.Int, error) {
	return _FeeManager.Contract.LiquidityRewardPercentage(&_FeeManager.CallOpts, arg0)
}

// LiquidityRewardPercentage is a free data retrieval call binding the contract method 0x1f71680b.
//
// Solidity: function liquidityRewardPercentage(address ) view returns(uint256)
func (_FeeManager *FeeManagerCallerSession) LiquidityRewardPercentage(arg0 common.Address) (*big.Int, error) {
	return _FeeManager.Contract.LiquidityRewardPercentage(&_FeeManager.CallOpts, arg0)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_FeeManager *FeeManagerCaller) SignerStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "signerStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_FeeManager *FeeManagerSession) SignerStorage() (common.Address, error) {
	return _FeeManager.Contract.SignerStorage(&_FeeManager.CallOpts)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_FeeManager *FeeManagerCallerSession) SignerStorage() (common.Address, error) {
	return _FeeManager.Contract.SignerStorage(&_FeeManager.CallOpts)
}

// TokenFeePercentage is a free data retrieval call binding the contract method 0x417e7465.
//
// Solidity: function tokenFeePercentage(address ) view returns(uint256)
func (_FeeManager *FeeManagerCaller) TokenFeePercentage(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "tokenFeePercentage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenFeePercentage is a free data retrieval call binding the contract method 0x417e7465.
//
// Solidity: function tokenFeePercentage(address ) view returns(uint256)
func (_FeeManager *FeeManagerSession) TokenFeePercentage(arg0 common.Address) (*big.Int, error) {
	return _FeeManager.Contract.TokenFeePercentage(&_FeeManager.CallOpts, arg0)
}

// TokenFeePercentage is a free data retrieval call binding the contract method 0x417e7465.
//
// Solidity: function tokenFeePercentage(address ) view returns(uint256)
func (_FeeManager *FeeManagerCallerSession) TokenFeePercentage(arg0 common.Address) (*big.Int, error) {
	return _FeeManager.Contract.TokenFeePercentage(&_FeeManager.CallOpts, arg0)
}

// ValidatorFeePool is a free data retrieval call binding the contract method 0x665bbce0.
//
// Solidity: function validatorFeePool() view returns(address)
func (_FeeManager *FeeManagerCaller) ValidatorFeePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "validatorFeePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorFeePool is a free data retrieval call binding the contract method 0x665bbce0.
//
// Solidity: function validatorFeePool() view returns(address)
func (_FeeManager *FeeManagerSession) ValidatorFeePool() (common.Address, error) {
	return _FeeManager.Contract.ValidatorFeePool(&_FeeManager.CallOpts)
}

// ValidatorFeePool is a free data retrieval call binding the contract method 0x665bbce0.
//
// Solidity: function validatorFeePool() view returns(address)
func (_FeeManager *FeeManagerCallerSession) ValidatorFeePool() (common.Address, error) {
	return _FeeManager.Contract.ValidatorFeePool(&_FeeManager.CallOpts)
}

// ValidatorRefundFee is a free data retrieval call binding the contract method 0x5c57c36c.
//
// Solidity: function validatorRefundFee() view returns(uint256)
func (_FeeManager *FeeManagerCaller) ValidatorRefundFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "validatorRefundFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorRefundFee is a free data retrieval call binding the contract method 0x5c57c36c.
//
// Solidity: function validatorRefundFee() view returns(uint256)
func (_FeeManager *FeeManagerSession) ValidatorRefundFee() (*big.Int, error) {
	return _FeeManager.Contract.ValidatorRefundFee(&_FeeManager.CallOpts)
}

// ValidatorRefundFee is a free data retrieval call binding the contract method 0x5c57c36c.
//
// Solidity: function validatorRefundFee() view returns(uint256)
func (_FeeManager *FeeManagerCallerSession) ValidatorRefundFee() (*big.Int, error) {
	return _FeeManager.Contract.ValidatorRefundFee(&_FeeManager.CallOpts)
}

// ValidatorRewardPercentage is a free data retrieval call binding the contract method 0x2f732967.
//
// Solidity: function validatorRewardPercentage(address ) view returns(uint256)
func (_FeeManager *FeeManagerCaller) ValidatorRewardPercentage(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeManager.contract.Call(opts, &out, "validatorRewardPercentage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorRewardPercentage is a free data retrieval call binding the contract method 0x2f732967.
//
// Solidity: function validatorRewardPercentage(address ) view returns(uint256)
func (_FeeManager *FeeManagerSession) ValidatorRewardPercentage(arg0 common.Address) (*big.Int, error) {
	return _FeeManager.Contract.ValidatorRewardPercentage(&_FeeManager.CallOpts, arg0)
}

// ValidatorRewardPercentage is a free data retrieval call binding the contract method 0x2f732967.
//
// Solidity: function validatorRewardPercentage(address ) view returns(uint256)
func (_FeeManager *FeeManagerCallerSession) ValidatorRewardPercentage(arg0 common.Address) (*big.Int, error) {
	return _FeeManager.Contract.ValidatorRewardPercentage(&_FeeManager.CallOpts, arg0)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x1de9d9b6.
//
// Solidity: function distributeRewards(address token) returns()
func (_FeeManager *FeeManagerTransactor) DistributeRewards(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "distributeRewards", token)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x1de9d9b6.
//
// Solidity: function distributeRewards(address token) returns()
func (_FeeManager *FeeManagerSession) DistributeRewards(token common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.DistributeRewards(&_FeeManager.TransactOpts, token)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x1de9d9b6.
//
// Solidity: function distributeRewards(address token) returns()
func (_FeeManager *FeeManagerTransactorSession) DistributeRewards(token common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.DistributeRewards(&_FeeManager.TransactOpts, token)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _signerStorage, address _liquidityPools, address _foundationAddress, address _validatorFeePool, uint256 _validatorRefundFee) returns()
func (_FeeManager *FeeManagerTransactor) Initialize(opts *bind.TransactOpts, _signerStorage common.Address, _liquidityPools common.Address, _foundationAddress common.Address, _validatorFeePool common.Address, _validatorRefundFee *big.Int) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "initialize", _signerStorage, _liquidityPools, _foundationAddress, _validatorFeePool, _validatorRefundFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _signerStorage, address _liquidityPools, address _foundationAddress, address _validatorFeePool, uint256 _validatorRefundFee) returns()
func (_FeeManager *FeeManagerSession) Initialize(_signerStorage common.Address, _liquidityPools common.Address, _foundationAddress common.Address, _validatorFeePool common.Address, _validatorRefundFee *big.Int) (*types.Transaction, error) {
	return _FeeManager.Contract.Initialize(&_FeeManager.TransactOpts, _signerStorage, _liquidityPools, _foundationAddress, _validatorFeePool, _validatorRefundFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _signerStorage, address _liquidityPools, address _foundationAddress, address _validatorFeePool, uint256 _validatorRefundFee) returns()
func (_FeeManager *FeeManagerTransactorSession) Initialize(_signerStorage common.Address, _liquidityPools common.Address, _foundationAddress common.Address, _validatorFeePool common.Address, _validatorRefundFee *big.Int) (*types.Transaction, error) {
	return _FeeManager.Contract.Initialize(&_FeeManager.TransactOpts, _signerStorage, _liquidityPools, _foundationAddress, _validatorFeePool, _validatorRefundFee)
}

// SetFoundationAddress is a paid mutator transaction binding the contract method 0xf41377ca.
//
// Solidity: function setFoundationAddress(address _foundationAddress) returns()
func (_FeeManager *FeeManagerTransactor) SetFoundationAddress(opts *bind.TransactOpts, _foundationAddress common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "setFoundationAddress", _foundationAddress)
}

// SetFoundationAddress is a paid mutator transaction binding the contract method 0xf41377ca.
//
// Solidity: function setFoundationAddress(address _foundationAddress) returns()
func (_FeeManager *FeeManagerSession) SetFoundationAddress(_foundationAddress common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.SetFoundationAddress(&_FeeManager.TransactOpts, _foundationAddress)
}

// SetFoundationAddress is a paid mutator transaction binding the contract method 0xf41377ca.
//
// Solidity: function setFoundationAddress(address _foundationAddress) returns()
func (_FeeManager *FeeManagerTransactorSession) SetFoundationAddress(_foundationAddress common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.SetFoundationAddress(&_FeeManager.TransactOpts, _foundationAddress)
}

// SetLiquidityPools is a paid mutator transaction binding the contract method 0xc5bc0753.
//
// Solidity: function setLiquidityPools(address _liquidityPools) returns()
func (_FeeManager *FeeManagerTransactor) SetLiquidityPools(opts *bind.TransactOpts, _liquidityPools common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "setLiquidityPools", _liquidityPools)
}

// SetLiquidityPools is a paid mutator transaction binding the contract method 0xc5bc0753.
//
// Solidity: function setLiquidityPools(address _liquidityPools) returns()
func (_FeeManager *FeeManagerSession) SetLiquidityPools(_liquidityPools common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.SetLiquidityPools(&_FeeManager.TransactOpts, _liquidityPools)
}

// SetLiquidityPools is a paid mutator transaction binding the contract method 0xc5bc0753.
//
// Solidity: function setLiquidityPools(address _liquidityPools) returns()
func (_FeeManager *FeeManagerTransactorSession) SetLiquidityPools(_liquidityPools common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.SetLiquidityPools(&_FeeManager.TransactOpts, _liquidityPools)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x302fee96.
//
// Solidity: function setTokenFee(address token, uint256 tokenFee, uint256 validatorReward, uint256 liquidityReward) returns()
func (_FeeManager *FeeManagerTransactor) SetTokenFee(opts *bind.TransactOpts, token common.Address, tokenFee *big.Int, validatorReward *big.Int, liquidityReward *big.Int) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "setTokenFee", token, tokenFee, validatorReward, liquidityReward)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x302fee96.
//
// Solidity: function setTokenFee(address token, uint256 tokenFee, uint256 validatorReward, uint256 liquidityReward) returns()
func (_FeeManager *FeeManagerSession) SetTokenFee(token common.Address, tokenFee *big.Int, validatorReward *big.Int, liquidityReward *big.Int) (*types.Transaction, error) {
	return _FeeManager.Contract.SetTokenFee(&_FeeManager.TransactOpts, token, tokenFee, validatorReward, liquidityReward)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x302fee96.
//
// Solidity: function setTokenFee(address token, uint256 tokenFee, uint256 validatorReward, uint256 liquidityReward) returns()
func (_FeeManager *FeeManagerTransactorSession) SetTokenFee(token common.Address, tokenFee *big.Int, validatorReward *big.Int, liquidityReward *big.Int) (*types.Transaction, error) {
	return _FeeManager.Contract.SetTokenFee(&_FeeManager.TransactOpts, token, tokenFee, validatorReward, liquidityReward)
}

// SetValidatorFeePool is a paid mutator transaction binding the contract method 0x756fdbe8.
//
// Solidity: function setValidatorFeePool(address _validatorFee) returns()
func (_FeeManager *FeeManagerTransactor) SetValidatorFeePool(opts *bind.TransactOpts, _validatorFee common.Address) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "setValidatorFeePool", _validatorFee)
}

// SetValidatorFeePool is a paid mutator transaction binding the contract method 0x756fdbe8.
//
// Solidity: function setValidatorFeePool(address _validatorFee) returns()
func (_FeeManager *FeeManagerSession) SetValidatorFeePool(_validatorFee common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.SetValidatorFeePool(&_FeeManager.TransactOpts, _validatorFee)
}

// SetValidatorFeePool is a paid mutator transaction binding the contract method 0x756fdbe8.
//
// Solidity: function setValidatorFeePool(address _validatorFee) returns()
func (_FeeManager *FeeManagerTransactorSession) SetValidatorFeePool(_validatorFee common.Address) (*types.Transaction, error) {
	return _FeeManager.Contract.SetValidatorFeePool(&_FeeManager.TransactOpts, _validatorFee)
}

// SetValidatorRefundFee is a paid mutator transaction binding the contract method 0xd28a1766.
//
// Solidity: function setValidatorRefundFee(uint256 _validatorRefundFee) returns()
func (_FeeManager *FeeManagerTransactor) SetValidatorRefundFee(opts *bind.TransactOpts, _validatorRefundFee *big.Int) (*types.Transaction, error) {
	return _FeeManager.contract.Transact(opts, "setValidatorRefundFee", _validatorRefundFee)
}

// SetValidatorRefundFee is a paid mutator transaction binding the contract method 0xd28a1766.
//
// Solidity: function setValidatorRefundFee(uint256 _validatorRefundFee) returns()
func (_FeeManager *FeeManagerSession) SetValidatorRefundFee(_validatorRefundFee *big.Int) (*types.Transaction, error) {
	return _FeeManager.Contract.SetValidatorRefundFee(&_FeeManager.TransactOpts, _validatorRefundFee)
}

// SetValidatorRefundFee is a paid mutator transaction binding the contract method 0xd28a1766.
//
// Solidity: function setValidatorRefundFee(uint256 _validatorRefundFee) returns()
func (_FeeManager *FeeManagerTransactorSession) SetValidatorRefundFee(_validatorRefundFee *big.Int) (*types.Transaction, error) {
	return _FeeManager.Contract.SetValidatorRefundFee(&_FeeManager.TransactOpts, _validatorRefundFee)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeManager *FeeManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeManager *FeeManagerSession) Receive() (*types.Transaction, error) {
	return _FeeManager.Contract.Receive(&_FeeManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FeeManager *FeeManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _FeeManager.Contract.Receive(&_FeeManager.TransactOpts)
}

// FeeManagerFoundationAddressUpdatedIterator is returned from FilterFoundationAddressUpdated and is used to iterate over the raw logs and unpacked data for FoundationAddressUpdated events raised by the FeeManager contract.
type FeeManagerFoundationAddressUpdatedIterator struct {
	Event *FeeManagerFoundationAddressUpdated // Event containing the contract specifics and raw log

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
func (it *FeeManagerFoundationAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerFoundationAddressUpdated)
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
		it.Event = new(FeeManagerFoundationAddressUpdated)
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
func (it *FeeManagerFoundationAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerFoundationAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerFoundationAddressUpdated represents a FoundationAddressUpdated event raised by the FeeManager contract.
type FeeManagerFoundationAddressUpdated struct {
	FoundationAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFoundationAddressUpdated is a free log retrieval operation binding the contract event 0xd451575e0d555b0fcd35add8eff1239d7486e3f4f256239a9af21ea316250e4b.
//
// Solidity: event FoundationAddressUpdated(address _foundationAddress)
func (_FeeManager *FeeManagerFilterer) FilterFoundationAddressUpdated(opts *bind.FilterOpts) (*FeeManagerFoundationAddressUpdatedIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "FoundationAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeManagerFoundationAddressUpdatedIterator{contract: _FeeManager.contract, event: "FoundationAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchFoundationAddressUpdated is a free log subscription operation binding the contract event 0xd451575e0d555b0fcd35add8eff1239d7486e3f4f256239a9af21ea316250e4b.
//
// Solidity: event FoundationAddressUpdated(address _foundationAddress)
func (_FeeManager *FeeManagerFilterer) WatchFoundationAddressUpdated(opts *bind.WatchOpts, sink chan<- *FeeManagerFoundationAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "FoundationAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerFoundationAddressUpdated)
				if err := _FeeManager.contract.UnpackLog(event, "FoundationAddressUpdated", log); err != nil {
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

// ParseFoundationAddressUpdated is a log parse operation binding the contract event 0xd451575e0d555b0fcd35add8eff1239d7486e3f4f256239a9af21ea316250e4b.
//
// Solidity: event FoundationAddressUpdated(address _foundationAddress)
func (_FeeManager *FeeManagerFilterer) ParseFoundationAddressUpdated(log types.Log) (*FeeManagerFoundationAddressUpdated, error) {
	event := new(FeeManagerFoundationAddressUpdated)
	if err := _FeeManager.contract.UnpackLog(event, "FoundationAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FeeManager contract.
type FeeManagerInitializedIterator struct {
	Event *FeeManagerInitialized // Event containing the contract specifics and raw log

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
func (it *FeeManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerInitialized)
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
		it.Event = new(FeeManagerInitialized)
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
func (it *FeeManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerInitialized represents a Initialized event raised by the FeeManager contract.
type FeeManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FeeManager *FeeManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*FeeManagerInitializedIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FeeManagerInitializedIterator{contract: _FeeManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_FeeManager *FeeManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FeeManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerInitialized)
				if err := _FeeManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_FeeManager *FeeManagerFilterer) ParseInitialized(log types.Log) (*FeeManagerInitialized, error) {
	event := new(FeeManagerInitialized)
	if err := _FeeManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerLiquidityPoolsUpdatedIterator is returned from FilterLiquidityPoolsUpdated and is used to iterate over the raw logs and unpacked data for LiquidityPoolsUpdated events raised by the FeeManager contract.
type FeeManagerLiquidityPoolsUpdatedIterator struct {
	Event *FeeManagerLiquidityPoolsUpdated // Event containing the contract specifics and raw log

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
func (it *FeeManagerLiquidityPoolsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerLiquidityPoolsUpdated)
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
		it.Event = new(FeeManagerLiquidityPoolsUpdated)
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
func (it *FeeManagerLiquidityPoolsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerLiquidityPoolsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerLiquidityPoolsUpdated represents a LiquidityPoolsUpdated event raised by the FeeManager contract.
type FeeManagerLiquidityPoolsUpdated struct {
	LiquidityPools common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidityPoolsUpdated is a free log retrieval operation binding the contract event 0xa267dfaba1c8ec87bb2497438b406981ff010a22f9def8ac81f7d79f06d85db6.
//
// Solidity: event LiquidityPoolsUpdated(address _liquidityPools)
func (_FeeManager *FeeManagerFilterer) FilterLiquidityPoolsUpdated(opts *bind.FilterOpts) (*FeeManagerLiquidityPoolsUpdatedIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "LiquidityPoolsUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeManagerLiquidityPoolsUpdatedIterator{contract: _FeeManager.contract, event: "LiquidityPoolsUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityPoolsUpdated is a free log subscription operation binding the contract event 0xa267dfaba1c8ec87bb2497438b406981ff010a22f9def8ac81f7d79f06d85db6.
//
// Solidity: event LiquidityPoolsUpdated(address _liquidityPools)
func (_FeeManager *FeeManagerFilterer) WatchLiquidityPoolsUpdated(opts *bind.WatchOpts, sink chan<- *FeeManagerLiquidityPoolsUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "LiquidityPoolsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerLiquidityPoolsUpdated)
				if err := _FeeManager.contract.UnpackLog(event, "LiquidityPoolsUpdated", log); err != nil {
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

// ParseLiquidityPoolsUpdated is a log parse operation binding the contract event 0xa267dfaba1c8ec87bb2497438b406981ff010a22f9def8ac81f7d79f06d85db6.
//
// Solidity: event LiquidityPoolsUpdated(address _liquidityPools)
func (_FeeManager *FeeManagerFilterer) ParseLiquidityPoolsUpdated(log types.Log) (*FeeManagerLiquidityPoolsUpdated, error) {
	event := new(FeeManagerLiquidityPoolsUpdated)
	if err := _FeeManager.contract.UnpackLog(event, "LiquidityPoolsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerValidatorFeeUpdatedIterator is returned from FilterValidatorFeeUpdated and is used to iterate over the raw logs and unpacked data for ValidatorFeeUpdated events raised by the FeeManager contract.
type FeeManagerValidatorFeeUpdatedIterator struct {
	Event *FeeManagerValidatorFeeUpdated // Event containing the contract specifics and raw log

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
func (it *FeeManagerValidatorFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerValidatorFeeUpdated)
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
		it.Event = new(FeeManagerValidatorFeeUpdated)
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
func (it *FeeManagerValidatorFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerValidatorFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerValidatorFeeUpdated represents a ValidatorFeeUpdated event raised by the FeeManager contract.
type FeeManagerValidatorFeeUpdated struct {
	ValidatorFee common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorFeeUpdated is a free log retrieval operation binding the contract event 0xfea9cde7862d6dd80ef0836638fe70c3cc79378595b0e136ebae4b6d62cca88f.
//
// Solidity: event ValidatorFeeUpdated(address _validatorFee)
func (_FeeManager *FeeManagerFilterer) FilterValidatorFeeUpdated(opts *bind.FilterOpts) (*FeeManagerValidatorFeeUpdatedIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "ValidatorFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeManagerValidatorFeeUpdatedIterator{contract: _FeeManager.contract, event: "ValidatorFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorFeeUpdated is a free log subscription operation binding the contract event 0xfea9cde7862d6dd80ef0836638fe70c3cc79378595b0e136ebae4b6d62cca88f.
//
// Solidity: event ValidatorFeeUpdated(address _validatorFee)
func (_FeeManager *FeeManagerFilterer) WatchValidatorFeeUpdated(opts *bind.WatchOpts, sink chan<- *FeeManagerValidatorFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "ValidatorFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerValidatorFeeUpdated)
				if err := _FeeManager.contract.UnpackLog(event, "ValidatorFeeUpdated", log); err != nil {
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

// ParseValidatorFeeUpdated is a log parse operation binding the contract event 0xfea9cde7862d6dd80ef0836638fe70c3cc79378595b0e136ebae4b6d62cca88f.
//
// Solidity: event ValidatorFeeUpdated(address _validatorFee)
func (_FeeManager *FeeManagerFilterer) ParseValidatorFeeUpdated(log types.Log) (*FeeManagerValidatorFeeUpdated, error) {
	event := new(FeeManagerValidatorFeeUpdated)
	if err := _FeeManager.contract.UnpackLog(event, "ValidatorFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeManagerValidatorRefundFeeUpdatedIterator is returned from FilterValidatorRefundFeeUpdated and is used to iterate over the raw logs and unpacked data for ValidatorRefundFeeUpdated events raised by the FeeManager contract.
type FeeManagerValidatorRefundFeeUpdatedIterator struct {
	Event *FeeManagerValidatorRefundFeeUpdated // Event containing the contract specifics and raw log

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
func (it *FeeManagerValidatorRefundFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeManagerValidatorRefundFeeUpdated)
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
		it.Event = new(FeeManagerValidatorRefundFeeUpdated)
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
func (it *FeeManagerValidatorRefundFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeManagerValidatorRefundFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeManagerValidatorRefundFeeUpdated represents a ValidatorRefundFeeUpdated event raised by the FeeManager contract.
type FeeManagerValidatorRefundFeeUpdated struct {
	ValidatorRefundFee *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRefundFeeUpdated is a free log retrieval operation binding the contract event 0x3783e624a91710496aafab2b6a1ce02bc5d2125fcf72e85dddda3aceb675dfc7.
//
// Solidity: event ValidatorRefundFeeUpdated(uint256 _validatorRefundFee)
func (_FeeManager *FeeManagerFilterer) FilterValidatorRefundFeeUpdated(opts *bind.FilterOpts) (*FeeManagerValidatorRefundFeeUpdatedIterator, error) {

	logs, sub, err := _FeeManager.contract.FilterLogs(opts, "ValidatorRefundFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeManagerValidatorRefundFeeUpdatedIterator{contract: _FeeManager.contract, event: "ValidatorRefundFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorRefundFeeUpdated is a free log subscription operation binding the contract event 0x3783e624a91710496aafab2b6a1ce02bc5d2125fcf72e85dddda3aceb675dfc7.
//
// Solidity: event ValidatorRefundFeeUpdated(uint256 _validatorRefundFee)
func (_FeeManager *FeeManagerFilterer) WatchValidatorRefundFeeUpdated(opts *bind.WatchOpts, sink chan<- *FeeManagerValidatorRefundFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeManager.contract.WatchLogs(opts, "ValidatorRefundFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeManagerValidatorRefundFeeUpdated)
				if err := _FeeManager.contract.UnpackLog(event, "ValidatorRefundFeeUpdated", log); err != nil {
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

// ParseValidatorRefundFeeUpdated is a log parse operation binding the contract event 0x3783e624a91710496aafab2b6a1ce02bc5d2125fcf72e85dddda3aceb675dfc7.
//
// Solidity: event ValidatorRefundFeeUpdated(uint256 _validatorRefundFee)
func (_FeeManager *FeeManagerFilterer) ParseValidatorRefundFeeUpdated(log types.Log) (*FeeManagerValidatorRefundFeeUpdated, error) {
	event := new(FeeManagerValidatorRefundFeeUpdated)
	if err := _FeeManager.contract.UnpackLog(event, "ValidatorRefundFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
