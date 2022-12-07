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

// ERC20BridgeMetaData contains all meta data concerning the ERC20Bridge contract.
var ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"}],\"name\":\"DepositedNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"FeeManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_liquidityPools\",\"type\":\"address\"}],\"name\":\"LiquidityPoolsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayBridge\",\"type\":\"address\"}],\"name\":\"RelayBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Reverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"TokenManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"ValidatorAddressUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAppAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"contractFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_liquidityPools\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeAppAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_relayBridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"isExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"isReverted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityPools\",\"outputs\":[{\"internalType\":\"contractLiquidityPools\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayBridge\",\"outputs\":[{\"internalType\":\"contractIRelayBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"revertSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"reverted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_liquidityPools\",\"type\":\"address\"}],\"name\":\"setLiquidityPools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayBridge\",\"type\":\"address\"}],\"name\":\"setRelayBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"setTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerStorage\",\"outputs\":[{\"internalType\":\"contractISignerStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenManager\",\"outputs\":[{\"internalType\":\"contractTokenManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BridgeMetaData.ABI instead.
var ERC20BridgeABI = ERC20BridgeMetaData.ABI

// ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type ERC20Bridge struct {
	ERC20BridgeCaller     // Read-only binding to the contract
	ERC20BridgeTransactor // Write-only binding to the contract
	ERC20BridgeFilterer   // Log filterer for contract events
}

// ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BridgeSession struct {
	Contract     *ERC20Bridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BridgeCallerSession struct {
	Contract *ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BridgeTransactorSession struct {
	Contract     *ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BridgeRaw struct {
	Contract *ERC20Bridge // Generic contract binding to access the raw methods on
}

// ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BridgeCallerRaw struct {
	Contract *ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BridgeTransactorRaw struct {
	Contract *ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Bridge creates a new instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20Bridge(address common.Address, backend bind.ContractBackend) (*ERC20Bridge, error) {
	contract, err := bindERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Bridge{ERC20BridgeCaller: ERC20BridgeCaller{contract: contract}, ERC20BridgeTransactor: ERC20BridgeTransactor{contract: contract}, ERC20BridgeFilterer: ERC20BridgeFilterer{contract: contract}}, nil
}

// NewERC20BridgeCaller creates a new read-only instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*ERC20BridgeCaller, error) {
	contract, err := bindERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeCaller{contract: contract}, nil
}

// NewERC20BridgeTransactor creates a new write-only instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BridgeTransactor, error) {
	contract, err := bindERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeTransactor{contract: contract}, nil
}

// NewERC20BridgeFilterer creates a new log filterer instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BridgeFilterer, error) {
	contract, err := bindERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeFilterer{contract: contract}, nil
}

// bindERC20Bridge binds a generic wrapper to an already deployed contract.
func bindERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Bridge *ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Bridge.Contract.ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Bridge *ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Bridge *ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Bridge *ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Bridge *ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Bridge *ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// BridgeAppAddress is a free data retrieval call binding the contract method 0xd7ac8777.
//
// Solidity: function bridgeAppAddress() view returns(address)
func (_ERC20Bridge *ERC20BridgeCaller) BridgeAppAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "bridgeAppAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAppAddress is a free data retrieval call binding the contract method 0xd7ac8777.
//
// Solidity: function bridgeAppAddress() view returns(address)
func (_ERC20Bridge *ERC20BridgeSession) BridgeAppAddress() (common.Address, error) {
	return _ERC20Bridge.Contract.BridgeAppAddress(&_ERC20Bridge.CallOpts)
}

// BridgeAppAddress is a free data retrieval call binding the contract method 0xd7ac8777.
//
// Solidity: function bridgeAppAddress() view returns(address)
func (_ERC20Bridge *ERC20BridgeCallerSession) BridgeAppAddress() (common.Address, error) {
	return _ERC20Bridge.Contract.BridgeAppAddress(&_ERC20Bridge.CallOpts)
}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCaller) Executed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "executed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeSession) Executed(arg0 [32]byte) (bool, error) {
	return _ERC20Bridge.Contract.Executed(&_ERC20Bridge.CallOpts, arg0)
}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCallerSession) Executed(arg0 [32]byte) (bool, error) {
	return _ERC20Bridge.Contract.Executed(&_ERC20Bridge.CallOpts, arg0)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_ERC20Bridge *ERC20BridgeCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_ERC20Bridge *ERC20BridgeSession) FeeManager() (common.Address, error) {
	return _ERC20Bridge.Contract.FeeManager(&_ERC20Bridge.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_ERC20Bridge *ERC20BridgeCallerSession) FeeManager() (common.Address, error) {
	return _ERC20Bridge.Contract.FeeManager(&_ERC20Bridge.CallOpts)
}

// IsExecuted is a free data retrieval call binding the contract method 0xe210a1f9.
//
// Solidity: function isExecuted(address _sender, bytes _data, address _token, address _receiver, uint256 _amount, uint256 _fee) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCaller) IsExecuted(opts *bind.CallOpts, _sender common.Address, _data []byte, _token common.Address, _receiver common.Address, _amount *big.Int, _fee *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "isExecuted", _sender, _data, _token, _receiver, _amount, _fee)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecuted is a free data retrieval call binding the contract method 0xe210a1f9.
//
// Solidity: function isExecuted(address _sender, bytes _data, address _token, address _receiver, uint256 _amount, uint256 _fee) view returns(bool)
func (_ERC20Bridge *ERC20BridgeSession) IsExecuted(_sender common.Address, _data []byte, _token common.Address, _receiver common.Address, _amount *big.Int, _fee *big.Int) (bool, error) {
	return _ERC20Bridge.Contract.IsExecuted(&_ERC20Bridge.CallOpts, _sender, _data, _token, _receiver, _amount, _fee)
}

// IsExecuted is a free data retrieval call binding the contract method 0xe210a1f9.
//
// Solidity: function isExecuted(address _sender, bytes _data, address _token, address _receiver, uint256 _amount, uint256 _fee) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCallerSession) IsExecuted(_sender common.Address, _data []byte, _token common.Address, _receiver common.Address, _amount *big.Int, _fee *big.Int) (bool, error) {
	return _ERC20Bridge.Contract.IsExecuted(&_ERC20Bridge.CallOpts, _sender, _data, _token, _receiver, _amount, _fee)
}

// IsReverted is a free data retrieval call binding the contract method 0x79a83810.
//
// Solidity: function isReverted(address _sender, bytes _data, address _token, address _receiver, uint256 _amount, uint256 _fee) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCaller) IsReverted(opts *bind.CallOpts, _sender common.Address, _data []byte, _token common.Address, _receiver common.Address, _amount *big.Int, _fee *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "isReverted", _sender, _data, _token, _receiver, _amount, _fee)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReverted is a free data retrieval call binding the contract method 0x79a83810.
//
// Solidity: function isReverted(address _sender, bytes _data, address _token, address _receiver, uint256 _amount, uint256 _fee) view returns(bool)
func (_ERC20Bridge *ERC20BridgeSession) IsReverted(_sender common.Address, _data []byte, _token common.Address, _receiver common.Address, _amount *big.Int, _fee *big.Int) (bool, error) {
	return _ERC20Bridge.Contract.IsReverted(&_ERC20Bridge.CallOpts, _sender, _data, _token, _receiver, _amount, _fee)
}

// IsReverted is a free data retrieval call binding the contract method 0x79a83810.
//
// Solidity: function isReverted(address _sender, bytes _data, address _token, address _receiver, uint256 _amount, uint256 _fee) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCallerSession) IsReverted(_sender common.Address, _data []byte, _token common.Address, _receiver common.Address, _amount *big.Int, _fee *big.Int) (bool, error) {
	return _ERC20Bridge.Contract.IsReverted(&_ERC20Bridge.CallOpts, _sender, _data, _token, _receiver, _amount, _fee)
}

// LiquidityPools is a free data retrieval call binding the contract method 0xd6efd7c3.
//
// Solidity: function liquidityPools() view returns(address)
func (_ERC20Bridge *ERC20BridgeCaller) LiquidityPools(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "liquidityPools")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityPools is a free data retrieval call binding the contract method 0xd6efd7c3.
//
// Solidity: function liquidityPools() view returns(address)
func (_ERC20Bridge *ERC20BridgeSession) LiquidityPools() (common.Address, error) {
	return _ERC20Bridge.Contract.LiquidityPools(&_ERC20Bridge.CallOpts)
}

// LiquidityPools is a free data retrieval call binding the contract method 0xd6efd7c3.
//
// Solidity: function liquidityPools() view returns(address)
func (_ERC20Bridge *ERC20BridgeCallerSession) LiquidityPools() (common.Address, error) {
	return _ERC20Bridge.Contract.LiquidityPools(&_ERC20Bridge.CallOpts)
}

// RelayBridge is a free data retrieval call binding the contract method 0xa76aa0e2.
//
// Solidity: function relayBridge() view returns(address)
func (_ERC20Bridge *ERC20BridgeCaller) RelayBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "relayBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayBridge is a free data retrieval call binding the contract method 0xa76aa0e2.
//
// Solidity: function relayBridge() view returns(address)
func (_ERC20Bridge *ERC20BridgeSession) RelayBridge() (common.Address, error) {
	return _ERC20Bridge.Contract.RelayBridge(&_ERC20Bridge.CallOpts)
}

// RelayBridge is a free data retrieval call binding the contract method 0xa76aa0e2.
//
// Solidity: function relayBridge() view returns(address)
func (_ERC20Bridge *ERC20BridgeCallerSession) RelayBridge() (common.Address, error) {
	return _ERC20Bridge.Contract.RelayBridge(&_ERC20Bridge.CallOpts)
}

// Reverted is a free data retrieval call binding the contract method 0xabe61ec4.
//
// Solidity: function reverted(bytes32 ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCaller) Reverted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "reverted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Reverted is a free data retrieval call binding the contract method 0xabe61ec4.
//
// Solidity: function reverted(bytes32 ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeSession) Reverted(arg0 [32]byte) (bool, error) {
	return _ERC20Bridge.Contract.Reverted(&_ERC20Bridge.CallOpts, arg0)
}

// Reverted is a free data retrieval call binding the contract method 0xabe61ec4.
//
// Solidity: function reverted(bytes32 ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCallerSession) Reverted(arg0 [32]byte) (bool, error) {
	return _ERC20Bridge.Contract.Reverted(&_ERC20Bridge.CallOpts, arg0)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_ERC20Bridge *ERC20BridgeCaller) SignerStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "signerStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_ERC20Bridge *ERC20BridgeSession) SignerStorage() (common.Address, error) {
	return _ERC20Bridge.Contract.SignerStorage(&_ERC20Bridge.CallOpts)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_ERC20Bridge *ERC20BridgeCallerSession) SignerStorage() (common.Address, error) {
	return _ERC20Bridge.Contract.SignerStorage(&_ERC20Bridge.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_ERC20Bridge *ERC20BridgeCaller) TokenManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "tokenManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_ERC20Bridge *ERC20BridgeSession) TokenManager() (common.Address, error) {
	return _ERC20Bridge.Contract.TokenManager(&_ERC20Bridge.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_ERC20Bridge *ERC20BridgeCallerSession) TokenManager() (common.Address, error) {
	return _ERC20Bridge.Contract.TokenManager(&_ERC20Bridge.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address _token, uint256 _chainId, address _receiver, uint256 _amount) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _chainId *big.Int, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "deposit", _token, _chainId, _receiver, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address _token, uint256 _chainId, address _receiver, uint256 _amount) returns()
func (_ERC20Bridge *ERC20BridgeSession) Deposit(_token common.Address, _chainId *big.Int, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Deposit(&_ERC20Bridge.TransactOpts, _token, _chainId, _receiver, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address _token, uint256 _chainId, address _receiver, uint256 _amount) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) Deposit(_token common.Address, _chainId *big.Int, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Deposit(&_ERC20Bridge.TransactOpts, _token, _chainId, _receiver, _amount)
}

// DepositNative is a paid mutator transaction binding the contract method 0xf0194945.
//
// Solidity: function depositNative(uint256 _chainId, address _receiver) payable returns()
func (_ERC20Bridge *ERC20BridgeTransactor) DepositNative(opts *bind.TransactOpts, _chainId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "depositNative", _chainId, _receiver)
}

// DepositNative is a paid mutator transaction binding the contract method 0xf0194945.
//
// Solidity: function depositNative(uint256 _chainId, address _receiver) payable returns()
func (_ERC20Bridge *ERC20BridgeSession) DepositNative(_chainId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.DepositNative(&_ERC20Bridge.TransactOpts, _chainId, _receiver)
}

// DepositNative is a paid mutator transaction binding the contract method 0xf0194945.
//
// Solidity: function depositNative(uint256 _chainId, address _receiver) payable returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) DepositNative(_chainId *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.DepositNative(&_ERC20Bridge.TransactOpts, _chainId, _receiver)
}

// Execute is a paid mutator transaction binding the contract method 0x59efcb15.
//
// Solidity: function execute(uint256 , bytes data) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) Execute(opts *bind.TransactOpts, arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "execute", arg0, data)
}

// Execute is a paid mutator transaction binding the contract method 0x59efcb15.
//
// Solidity: function execute(uint256 , bytes data) returns()
func (_ERC20Bridge *ERC20BridgeSession) Execute(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Execute(&_ERC20Bridge.TransactOpts, arg0, data)
}

// Execute is a paid mutator transaction binding the contract method 0x59efcb15.
//
// Solidity: function execute(uint256 , bytes data) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) Execute(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Execute(&_ERC20Bridge.TransactOpts, arg0, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _signerStorage, address _tokenManager, address _liquidityPools, address _feeManager, address _bridgeAppAddress, address _relayBridge) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) Initialize(opts *bind.TransactOpts, _signerStorage common.Address, _tokenManager common.Address, _liquidityPools common.Address, _feeManager common.Address, _bridgeAppAddress common.Address, _relayBridge common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "initialize", _signerStorage, _tokenManager, _liquidityPools, _feeManager, _bridgeAppAddress, _relayBridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _signerStorage, address _tokenManager, address _liquidityPools, address _feeManager, address _bridgeAppAddress, address _relayBridge) returns()
func (_ERC20Bridge *ERC20BridgeSession) Initialize(_signerStorage common.Address, _tokenManager common.Address, _liquidityPools common.Address, _feeManager common.Address, _bridgeAppAddress common.Address, _relayBridge common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Initialize(&_ERC20Bridge.TransactOpts, _signerStorage, _tokenManager, _liquidityPools, _feeManager, _bridgeAppAddress, _relayBridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _signerStorage, address _tokenManager, address _liquidityPools, address _feeManager, address _bridgeAppAddress, address _relayBridge) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) Initialize(_signerStorage common.Address, _tokenManager common.Address, _liquidityPools common.Address, _feeManager common.Address, _bridgeAppAddress common.Address, _relayBridge common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.Initialize(&_ERC20Bridge.TransactOpts, _signerStorage, _tokenManager, _liquidityPools, _feeManager, _bridgeAppAddress, _relayBridge)
}

// RevertSend is a paid mutator transaction binding the contract method 0x0d788db0.
//
// Solidity: function revertSend(uint256 , bytes data) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) RevertSend(opts *bind.TransactOpts, arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "revertSend", arg0, data)
}

// RevertSend is a paid mutator transaction binding the contract method 0x0d788db0.
//
// Solidity: function revertSend(uint256 , bytes data) returns()
func (_ERC20Bridge *ERC20BridgeSession) RevertSend(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.RevertSend(&_ERC20Bridge.TransactOpts, arg0, data)
}

// RevertSend is a paid mutator transaction binding the contract method 0x0d788db0.
//
// Solidity: function revertSend(uint256 , bytes data) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) RevertSend(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.RevertSend(&_ERC20Bridge.TransactOpts, arg0, data)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) SetFeeManager(opts *bind.TransactOpts, _feeManager common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "setFeeManager", _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_ERC20Bridge *ERC20BridgeSession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SetFeeManager(&_ERC20Bridge.TransactOpts, _feeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) SetFeeManager(_feeManager common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SetFeeManager(&_ERC20Bridge.TransactOpts, _feeManager)
}

// SetLiquidityPools is a paid mutator transaction binding the contract method 0xc5bc0753.
//
// Solidity: function setLiquidityPools(address _liquidityPools) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) SetLiquidityPools(opts *bind.TransactOpts, _liquidityPools common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "setLiquidityPools", _liquidityPools)
}

// SetLiquidityPools is a paid mutator transaction binding the contract method 0xc5bc0753.
//
// Solidity: function setLiquidityPools(address _liquidityPools) returns()
func (_ERC20Bridge *ERC20BridgeSession) SetLiquidityPools(_liquidityPools common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SetLiquidityPools(&_ERC20Bridge.TransactOpts, _liquidityPools)
}

// SetLiquidityPools is a paid mutator transaction binding the contract method 0xc5bc0753.
//
// Solidity: function setLiquidityPools(address _liquidityPools) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) SetLiquidityPools(_liquidityPools common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SetLiquidityPools(&_ERC20Bridge.TransactOpts, _liquidityPools)
}

// SetRelayBridge is a paid mutator transaction binding the contract method 0x65011549.
//
// Solidity: function setRelayBridge(address _relayBridge) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) SetRelayBridge(opts *bind.TransactOpts, _relayBridge common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "setRelayBridge", _relayBridge)
}

// SetRelayBridge is a paid mutator transaction binding the contract method 0x65011549.
//
// Solidity: function setRelayBridge(address _relayBridge) returns()
func (_ERC20Bridge *ERC20BridgeSession) SetRelayBridge(_relayBridge common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SetRelayBridge(&_ERC20Bridge.TransactOpts, _relayBridge)
}

// SetRelayBridge is a paid mutator transaction binding the contract method 0x65011549.
//
// Solidity: function setRelayBridge(address _relayBridge) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) SetRelayBridge(_relayBridge common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SetRelayBridge(&_ERC20Bridge.TransactOpts, _relayBridge)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) SetTokenManager(opts *bind.TransactOpts, _tokenManager common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "setTokenManager", _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_ERC20Bridge *ERC20BridgeSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SetTokenManager(&_ERC20Bridge.TransactOpts, _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SetTokenManager(&_ERC20Bridge.TransactOpts, _tokenManager)
}

// ERC20BridgeDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ERC20Bridge contract.
type ERC20BridgeDepositedIterator struct {
	Event *ERC20BridgeDeposited // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeDeposited)
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
		it.Event = new(ERC20BridgeDeposited)
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
func (it *ERC20BridgeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeDeposited represents a Deposited event raised by the ERC20Bridge contract.
type ERC20BridgeDeposited struct {
	Sender             common.Address
	Token              common.Address
	DestinationChainId *big.Int
	Receiver           common.Address
	Fee                *big.Int
	TransferAmount     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xb64b8d2a723c97c9a2594ccb3778465edeb4e8b73151aee791183eeda118efff.
//
// Solidity: event Deposited(address sender, address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterDeposited(opts *bind.FilterOpts) (*ERC20BridgeDepositedIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeDepositedIterator{contract: _ERC20Bridge.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xb64b8d2a723c97c9a2594ccb3778465edeb4e8b73151aee791183eeda118efff.
//
// Solidity: event Deposited(address sender, address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ERC20BridgeDeposited) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeDeposited)
				if err := _ERC20Bridge.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xb64b8d2a723c97c9a2594ccb3778465edeb4e8b73151aee791183eeda118efff.
//
// Solidity: event Deposited(address sender, address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseDeposited(log types.Log) (*ERC20BridgeDeposited, error) {
	event := new(ERC20BridgeDeposited)
	if err := _ERC20Bridge.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeDepositedNativeIterator is returned from FilterDepositedNative and is used to iterate over the raw logs and unpacked data for DepositedNative events raised by the ERC20Bridge contract.
type ERC20BridgeDepositedNativeIterator struct {
	Event *ERC20BridgeDepositedNative // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeDepositedNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeDepositedNative)
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
		it.Event = new(ERC20BridgeDepositedNative)
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
func (it *ERC20BridgeDepositedNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeDepositedNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeDepositedNative represents a DepositedNative event raised by the ERC20Bridge contract.
type ERC20BridgeDepositedNative struct {
	Sender             common.Address
	Token              common.Address
	DestinationChainId *big.Int
	Receiver           common.Address
	Fee                *big.Int
	TransferAmount     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositedNative is a free log retrieval operation binding the contract event 0x594aceaea2948a8384c265adbb963c46e8d8b5c405d4d3b98622384fbbce00ac.
//
// Solidity: event DepositedNative(address sender, address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterDepositedNative(opts *bind.FilterOpts) (*ERC20BridgeDepositedNativeIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeDepositedNativeIterator{contract: _ERC20Bridge.contract, event: "DepositedNative", logs: logs, sub: sub}, nil
}

// WatchDepositedNative is a free log subscription operation binding the contract event 0x594aceaea2948a8384c265adbb963c46e8d8b5c405d4d3b98622384fbbce00ac.
//
// Solidity: event DepositedNative(address sender, address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchDepositedNative(opts *bind.WatchOpts, sink chan<- *ERC20BridgeDepositedNative) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeDepositedNative)
				if err := _ERC20Bridge.contract.UnpackLog(event, "DepositedNative", log); err != nil {
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

// ParseDepositedNative is a log parse operation binding the contract event 0x594aceaea2948a8384c265adbb963c46e8d8b5c405d4d3b98622384fbbce00ac.
//
// Solidity: event DepositedNative(address sender, address token, uint256 destinationChainId, address receiver, uint256 fee, uint256 transferAmount)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseDepositedNative(log types.Log) (*ERC20BridgeDepositedNative, error) {
	event := new(ERC20BridgeDepositedNative)
	if err := _ERC20Bridge.contract.UnpackLog(event, "DepositedNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeFeeManagerUpdatedIterator is returned from FilterFeeManagerUpdated and is used to iterate over the raw logs and unpacked data for FeeManagerUpdated events raised by the ERC20Bridge contract.
type ERC20BridgeFeeManagerUpdatedIterator struct {
	Event *ERC20BridgeFeeManagerUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeFeeManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeFeeManagerUpdated)
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
		it.Event = new(ERC20BridgeFeeManagerUpdated)
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
func (it *ERC20BridgeFeeManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeFeeManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeFeeManagerUpdated represents a FeeManagerUpdated event raised by the ERC20Bridge contract.
type ERC20BridgeFeeManagerUpdated struct {
	FeeManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeManagerUpdated is a free log retrieval operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address _feeManager)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterFeeManagerUpdated(opts *bind.FilterOpts) (*ERC20BridgeFeeManagerUpdatedIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "FeeManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeFeeManagerUpdatedIterator{contract: _ERC20Bridge.contract, event: "FeeManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeManagerUpdated is a free log subscription operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address _feeManager)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchFeeManagerUpdated(opts *bind.WatchOpts, sink chan<- *ERC20BridgeFeeManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "FeeManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeFeeManagerUpdated)
				if err := _ERC20Bridge.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
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

// ParseFeeManagerUpdated is a log parse operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address _feeManager)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseFeeManagerUpdated(log types.Log) (*ERC20BridgeFeeManagerUpdated, error) {
	event := new(ERC20BridgeFeeManagerUpdated)
	if err := _ERC20Bridge.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20Bridge contract.
type ERC20BridgeInitializedIterator struct {
	Event *ERC20BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeInitialized)
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
		it.Event = new(ERC20BridgeInitialized)
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
func (it *ERC20BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeInitialized represents a Initialized event raised by the ERC20Bridge contract.
type ERC20BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20BridgeInitializedIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeInitializedIterator{contract: _ERC20Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeInitialized)
				if err := _ERC20Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC20Bridge *ERC20BridgeFilterer) ParseInitialized(log types.Log) (*ERC20BridgeInitialized, error) {
	event := new(ERC20BridgeInitialized)
	if err := _ERC20Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeLiquidityPoolsUpdatedIterator is returned from FilterLiquidityPoolsUpdated and is used to iterate over the raw logs and unpacked data for LiquidityPoolsUpdated events raised by the ERC20Bridge contract.
type ERC20BridgeLiquidityPoolsUpdatedIterator struct {
	Event *ERC20BridgeLiquidityPoolsUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeLiquidityPoolsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeLiquidityPoolsUpdated)
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
		it.Event = new(ERC20BridgeLiquidityPoolsUpdated)
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
func (it *ERC20BridgeLiquidityPoolsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeLiquidityPoolsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeLiquidityPoolsUpdated represents a LiquidityPoolsUpdated event raised by the ERC20Bridge contract.
type ERC20BridgeLiquidityPoolsUpdated struct {
	LiquidityPools common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidityPoolsUpdated is a free log retrieval operation binding the contract event 0xa267dfaba1c8ec87bb2497438b406981ff010a22f9def8ac81f7d79f06d85db6.
//
// Solidity: event LiquidityPoolsUpdated(address _liquidityPools)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterLiquidityPoolsUpdated(opts *bind.FilterOpts) (*ERC20BridgeLiquidityPoolsUpdatedIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "LiquidityPoolsUpdated")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeLiquidityPoolsUpdatedIterator{contract: _ERC20Bridge.contract, event: "LiquidityPoolsUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityPoolsUpdated is a free log subscription operation binding the contract event 0xa267dfaba1c8ec87bb2497438b406981ff010a22f9def8ac81f7d79f06d85db6.
//
// Solidity: event LiquidityPoolsUpdated(address _liquidityPools)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchLiquidityPoolsUpdated(opts *bind.WatchOpts, sink chan<- *ERC20BridgeLiquidityPoolsUpdated) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "LiquidityPoolsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeLiquidityPoolsUpdated)
				if err := _ERC20Bridge.contract.UnpackLog(event, "LiquidityPoolsUpdated", log); err != nil {
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
func (_ERC20Bridge *ERC20BridgeFilterer) ParseLiquidityPoolsUpdated(log types.Log) (*ERC20BridgeLiquidityPoolsUpdated, error) {
	event := new(ERC20BridgeLiquidityPoolsUpdated)
	if err := _ERC20Bridge.contract.UnpackLog(event, "LiquidityPoolsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeRelayBridgeUpdatedIterator is returned from FilterRelayBridgeUpdated and is used to iterate over the raw logs and unpacked data for RelayBridgeUpdated events raised by the ERC20Bridge contract.
type ERC20BridgeRelayBridgeUpdatedIterator struct {
	Event *ERC20BridgeRelayBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeRelayBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeRelayBridgeUpdated)
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
		it.Event = new(ERC20BridgeRelayBridgeUpdated)
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
func (it *ERC20BridgeRelayBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeRelayBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeRelayBridgeUpdated represents a RelayBridgeUpdated event raised by the ERC20Bridge contract.
type ERC20BridgeRelayBridgeUpdated struct {
	RelayBridge common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayBridgeUpdated is a free log retrieval operation binding the contract event 0xcabfbcf520a6680a53751431f8f30b90fe2a0c9d0fb95f4c3b9dc418f45730a4.
//
// Solidity: event RelayBridgeUpdated(address _relayBridge)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterRelayBridgeUpdated(opts *bind.FilterOpts) (*ERC20BridgeRelayBridgeUpdatedIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "RelayBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeRelayBridgeUpdatedIterator{contract: _ERC20Bridge.contract, event: "RelayBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchRelayBridgeUpdated is a free log subscription operation binding the contract event 0xcabfbcf520a6680a53751431f8f30b90fe2a0c9d0fb95f4c3b9dc418f45730a4.
//
// Solidity: event RelayBridgeUpdated(address _relayBridge)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchRelayBridgeUpdated(opts *bind.WatchOpts, sink chan<- *ERC20BridgeRelayBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "RelayBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeRelayBridgeUpdated)
				if err := _ERC20Bridge.contract.UnpackLog(event, "RelayBridgeUpdated", log); err != nil {
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

// ParseRelayBridgeUpdated is a log parse operation binding the contract event 0xcabfbcf520a6680a53751431f8f30b90fe2a0c9d0fb95f4c3b9dc418f45730a4.
//
// Solidity: event RelayBridgeUpdated(address _relayBridge)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseRelayBridgeUpdated(log types.Log) (*ERC20BridgeRelayBridgeUpdated, error) {
	event := new(ERC20BridgeRelayBridgeUpdated)
	if err := _ERC20Bridge.contract.UnpackLog(event, "RelayBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the ERC20Bridge contract.
type ERC20BridgeRevertedIterator struct {
	Event *ERC20BridgeReverted // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeReverted)
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
		it.Event = new(ERC20BridgeReverted)
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
func (it *ERC20BridgeRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeReverted represents a Reverted event raised by the ERC20Bridge contract.
type ERC20BridgeReverted struct {
	Sender        common.Address
	Token         common.Address
	SourceChainId *big.Int
	Receiver      common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0x6b79def8f04467814352cf318743c5f80de8e5481a67b938b1ea3f9887683501.
//
// Solidity: event Reverted(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterReverted(opts *bind.FilterOpts) (*ERC20BridgeRevertedIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "Reverted")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeRevertedIterator{contract: _ERC20Bridge.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0x6b79def8f04467814352cf318743c5f80de8e5481a67b938b1ea3f9887683501.
//
// Solidity: event Reverted(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *ERC20BridgeReverted) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "Reverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeReverted)
				if err := _ERC20Bridge.contract.UnpackLog(event, "Reverted", log); err != nil {
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

// ParseReverted is a log parse operation binding the contract event 0x6b79def8f04467814352cf318743c5f80de8e5481a67b938b1ea3f9887683501.
//
// Solidity: event Reverted(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseReverted(log types.Log) (*ERC20BridgeReverted, error) {
	event := new(ERC20BridgeReverted)
	if err := _ERC20Bridge.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeTokenManagerUpdatedIterator is returned from FilterTokenManagerUpdated and is used to iterate over the raw logs and unpacked data for TokenManagerUpdated events raised by the ERC20Bridge contract.
type ERC20BridgeTokenManagerUpdatedIterator struct {
	Event *ERC20BridgeTokenManagerUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeTokenManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeTokenManagerUpdated)
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
		it.Event = new(ERC20BridgeTokenManagerUpdated)
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
func (it *ERC20BridgeTokenManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeTokenManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeTokenManagerUpdated represents a TokenManagerUpdated event raised by the ERC20Bridge contract.
type ERC20BridgeTokenManagerUpdated struct {
	TokenManager common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenManagerUpdated is a free log retrieval operation binding the contract event 0x160ba3f04f73a1d23914cbe64026f708d7bea7296d144a3a2f759b97349a63fd.
//
// Solidity: event TokenManagerUpdated(address _tokenManager)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterTokenManagerUpdated(opts *bind.FilterOpts) (*ERC20BridgeTokenManagerUpdatedIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "TokenManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeTokenManagerUpdatedIterator{contract: _ERC20Bridge.contract, event: "TokenManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenManagerUpdated is a free log subscription operation binding the contract event 0x160ba3f04f73a1d23914cbe64026f708d7bea7296d144a3a2f759b97349a63fd.
//
// Solidity: event TokenManagerUpdated(address _tokenManager)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchTokenManagerUpdated(opts *bind.WatchOpts, sink chan<- *ERC20BridgeTokenManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "TokenManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeTokenManagerUpdated)
				if err := _ERC20Bridge.contract.UnpackLog(event, "TokenManagerUpdated", log); err != nil {
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

// ParseTokenManagerUpdated is a log parse operation binding the contract event 0x160ba3f04f73a1d23914cbe64026f708d7bea7296d144a3a2f759b97349a63fd.
//
// Solidity: event TokenManagerUpdated(address _tokenManager)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseTokenManagerUpdated(log types.Log) (*ERC20BridgeTokenManagerUpdated, error) {
	event := new(ERC20BridgeTokenManagerUpdated)
	if err := _ERC20Bridge.contract.UnpackLog(event, "TokenManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the ERC20Bridge contract.
type ERC20BridgeTransferredIterator struct {
	Event *ERC20BridgeTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeTransferred)
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
		it.Event = new(ERC20BridgeTransferred)
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
func (it *ERC20BridgeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeTransferred represents a Transferred event raised by the ERC20Bridge contract.
type ERC20BridgeTransferred struct {
	Sender        common.Address
	Token         common.Address
	SourceChainId *big.Int
	Receiver      common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0x621c87a009a8535d1a52333309173ce7117c61f0d228bb8737404185ea6764d2.
//
// Solidity: event Transferred(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterTransferred(opts *bind.FilterOpts) (*ERC20BridgeTransferredIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeTransferredIterator{contract: _ERC20Bridge.contract, event: "Transferred", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0x621c87a009a8535d1a52333309173ce7117c61f0d228bb8737404185ea6764d2.
//
// Solidity: event Transferred(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *ERC20BridgeTransferred) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeTransferred)
				if err := _ERC20Bridge.contract.UnpackLog(event, "Transferred", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0x621c87a009a8535d1a52333309173ce7117c61f0d228bb8737404185ea6764d2.
//
// Solidity: event Transferred(address sender, address token, uint256 sourceChainId, address receiver, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseTransferred(log types.Log) (*ERC20BridgeTransferred, error) {
	event := new(ERC20BridgeTransferred)
	if err := _ERC20Bridge.contract.UnpackLog(event, "Transferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeValidatorAddressUpdatedIterator is returned from FilterValidatorAddressUpdated and is used to iterate over the raw logs and unpacked data for ValidatorAddressUpdated events raised by the ERC20Bridge contract.
type ERC20BridgeValidatorAddressUpdatedIterator struct {
	Event *ERC20BridgeValidatorAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeValidatorAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeValidatorAddressUpdated)
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
		it.Event = new(ERC20BridgeValidatorAddressUpdated)
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
func (it *ERC20BridgeValidatorAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeValidatorAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeValidatorAddressUpdated represents a ValidatorAddressUpdated event raised by the ERC20Bridge contract.
type ERC20BridgeValidatorAddressUpdated struct {
	ValidatorAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorAddressUpdated is a free log retrieval operation binding the contract event 0xdf570ab3f5307e9adf3ff792b3dbfa0cd0d9f82e667552def683ae20be981f72.
//
// Solidity: event ValidatorAddressUpdated(address _validatorAddress)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterValidatorAddressUpdated(opts *bind.FilterOpts) (*ERC20BridgeValidatorAddressUpdatedIterator, error) {

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "ValidatorAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeValidatorAddressUpdatedIterator{contract: _ERC20Bridge.contract, event: "ValidatorAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorAddressUpdated is a free log subscription operation binding the contract event 0xdf570ab3f5307e9adf3ff792b3dbfa0cd0d9f82e667552def683ae20be981f72.
//
// Solidity: event ValidatorAddressUpdated(address _validatorAddress)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchValidatorAddressUpdated(opts *bind.WatchOpts, sink chan<- *ERC20BridgeValidatorAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "ValidatorAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeValidatorAddressUpdated)
				if err := _ERC20Bridge.contract.UnpackLog(event, "ValidatorAddressUpdated", log); err != nil {
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

// ParseValidatorAddressUpdated is a log parse operation binding the contract event 0xdf570ab3f5307e9adf3ff792b3dbfa0cd0d9f82e667552def683ae20be981f72.
//
// Solidity: event ValidatorAddressUpdated(address _validatorAddress)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseValidatorAddressUpdated(log types.Log) (*ERC20BridgeValidatorAddressUpdated, error) {
	event := new(ERC20BridgeValidatorAddressUpdated)
	if err := _ERC20Bridge.contract.UnpackLog(event, "ValidatorAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
