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

// BridgeValidatorFeePoolMetaData contains all meta data concerning the BridgeValidatorFeePool contract.
var BridgeValidatorFeePoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Collected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Bridge\",\"type\":\"address\"}],\"name\":\"ERC20BridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"LimitPerTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validatorFeeReceiver\",\"type\":\"address\"}],\"name\":\"ValidatorFeeReceiverUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20Bridge\",\"outputs\":[{\"internalType\":\"contractERC20Bridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"homeChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc20Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_homeChainId\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"limitPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Bridge\",\"type\":\"address\"}],\"name\":\"setERC20Bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"setLimitPerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorFeeReceiver\",\"type\":\"address\"}],\"name\":\"setValidatorFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerStorage\",\"outputs\":[{\"internalType\":\"contractISignerStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BridgeValidatorFeePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeValidatorFeePoolMetaData.ABI instead.
var BridgeValidatorFeePoolABI = BridgeValidatorFeePoolMetaData.ABI

// BridgeValidatorFeePool is an auto generated Go binding around an Ethereum contract.
type BridgeValidatorFeePool struct {
	BridgeValidatorFeePoolCaller     // Read-only binding to the contract
	BridgeValidatorFeePoolTransactor // Write-only binding to the contract
	BridgeValidatorFeePoolFilterer   // Log filterer for contract events
}

// BridgeValidatorFeePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeValidatorFeePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeValidatorFeePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeValidatorFeePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeValidatorFeePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeValidatorFeePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeValidatorFeePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeValidatorFeePoolSession struct {
	Contract     *BridgeValidatorFeePool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeValidatorFeePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeValidatorFeePoolCallerSession struct {
	Contract *BridgeValidatorFeePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BridgeValidatorFeePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeValidatorFeePoolTransactorSession struct {
	Contract     *BridgeValidatorFeePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BridgeValidatorFeePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeValidatorFeePoolRaw struct {
	Contract *BridgeValidatorFeePool // Generic contract binding to access the raw methods on
}

// BridgeValidatorFeePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeValidatorFeePoolCallerRaw struct {
	Contract *BridgeValidatorFeePoolCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeValidatorFeePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeValidatorFeePoolTransactorRaw struct {
	Contract *BridgeValidatorFeePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeValidatorFeePool creates a new instance of BridgeValidatorFeePool, bound to a specific deployed contract.
func NewBridgeValidatorFeePool(address common.Address, backend bind.ContractBackend) (*BridgeValidatorFeePool, error) {
	contract, err := bindBridgeValidatorFeePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorFeePool{BridgeValidatorFeePoolCaller: BridgeValidatorFeePoolCaller{contract: contract}, BridgeValidatorFeePoolTransactor: BridgeValidatorFeePoolTransactor{contract: contract}, BridgeValidatorFeePoolFilterer: BridgeValidatorFeePoolFilterer{contract: contract}}, nil
}

// NewBridgeValidatorFeePoolCaller creates a new read-only instance of BridgeValidatorFeePool, bound to a specific deployed contract.
func NewBridgeValidatorFeePoolCaller(address common.Address, caller bind.ContractCaller) (*BridgeValidatorFeePoolCaller, error) {
	contract, err := bindBridgeValidatorFeePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorFeePoolCaller{contract: contract}, nil
}

// NewBridgeValidatorFeePoolTransactor creates a new write-only instance of BridgeValidatorFeePool, bound to a specific deployed contract.
func NewBridgeValidatorFeePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeValidatorFeePoolTransactor, error) {
	contract, err := bindBridgeValidatorFeePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorFeePoolTransactor{contract: contract}, nil
}

// NewBridgeValidatorFeePoolFilterer creates a new log filterer instance of BridgeValidatorFeePool, bound to a specific deployed contract.
func NewBridgeValidatorFeePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeValidatorFeePoolFilterer, error) {
	contract, err := bindBridgeValidatorFeePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorFeePoolFilterer{contract: contract}, nil
}

// bindBridgeValidatorFeePool binds a generic wrapper to an already deployed contract.
func bindBridgeValidatorFeePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeValidatorFeePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeValidatorFeePool.Contract.BridgeValidatorFeePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.BridgeValidatorFeePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.BridgeValidatorFeePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeValidatorFeePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.contract.Transact(opts, method, params...)
}

// Erc20Bridge is a free data retrieval call binding the contract method 0xf8ffea7e.
//
// Solidity: function erc20Bridge() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCaller) Erc20Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeValidatorFeePool.contract.Call(opts, &out, "erc20Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20Bridge is a free data retrieval call binding the contract method 0xf8ffea7e.
//
// Solidity: function erc20Bridge() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) Erc20Bridge() (common.Address, error) {
	return _BridgeValidatorFeePool.Contract.Erc20Bridge(&_BridgeValidatorFeePool.CallOpts)
}

// Erc20Bridge is a free data retrieval call binding the contract method 0xf8ffea7e.
//
// Solidity: function erc20Bridge() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCallerSession) Erc20Bridge() (common.Address, error) {
	return _BridgeValidatorFeePool.Contract.Erc20Bridge(&_BridgeValidatorFeePool.CallOpts)
}

// HomeChainId is a free data retrieval call binding the contract method 0xef126967.
//
// Solidity: function homeChainId() view returns(uint256)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCaller) HomeChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeValidatorFeePool.contract.Call(opts, &out, "homeChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HomeChainId is a free data retrieval call binding the contract method 0xef126967.
//
// Solidity: function homeChainId() view returns(uint256)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) HomeChainId() (*big.Int, error) {
	return _BridgeValidatorFeePool.Contract.HomeChainId(&_BridgeValidatorFeePool.CallOpts)
}

// HomeChainId is a free data retrieval call binding the contract method 0xef126967.
//
// Solidity: function homeChainId() view returns(uint256)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCallerSession) HomeChainId() (*big.Int, error) {
	return _BridgeValidatorFeePool.Contract.HomeChainId(&_BridgeValidatorFeePool.CallOpts)
}

// LimitPerToken is a free data retrieval call binding the contract method 0x0bb51dd5.
//
// Solidity: function limitPerToken(address ) view returns(uint256)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCaller) LimitPerToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeValidatorFeePool.contract.Call(opts, &out, "limitPerToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitPerToken is a free data retrieval call binding the contract method 0x0bb51dd5.
//
// Solidity: function limitPerToken(address ) view returns(uint256)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) LimitPerToken(arg0 common.Address) (*big.Int, error) {
	return _BridgeValidatorFeePool.Contract.LimitPerToken(&_BridgeValidatorFeePool.CallOpts, arg0)
}

// LimitPerToken is a free data retrieval call binding the contract method 0x0bb51dd5.
//
// Solidity: function limitPerToken(address ) view returns(uint256)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCallerSession) LimitPerToken(arg0 common.Address) (*big.Int, error) {
	return _BridgeValidatorFeePool.Contract.LimitPerToken(&_BridgeValidatorFeePool.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeValidatorFeePool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) Owner() (common.Address, error) {
	return _BridgeValidatorFeePool.Contract.Owner(&_BridgeValidatorFeePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCallerSession) Owner() (common.Address, error) {
	return _BridgeValidatorFeePool.Contract.Owner(&_BridgeValidatorFeePool.CallOpts)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCaller) SignerStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeValidatorFeePool.contract.Call(opts, &out, "signerStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) SignerStorage() (common.Address, error) {
	return _BridgeValidatorFeePool.Contract.SignerStorage(&_BridgeValidatorFeePool.CallOpts)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCallerSession) SignerStorage() (common.Address, error) {
	return _BridgeValidatorFeePool.Contract.SignerStorage(&_BridgeValidatorFeePool.CallOpts)
}

// ValidatorFeeReceiver is a free data retrieval call binding the contract method 0x22816a4c.
//
// Solidity: function validatorFeeReceiver() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCaller) ValidatorFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeValidatorFeePool.contract.Call(opts, &out, "validatorFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorFeeReceiver is a free data retrieval call binding the contract method 0x22816a4c.
//
// Solidity: function validatorFeeReceiver() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) ValidatorFeeReceiver() (common.Address, error) {
	return _BridgeValidatorFeePool.Contract.ValidatorFeeReceiver(&_BridgeValidatorFeePool.CallOpts)
}

// ValidatorFeeReceiver is a free data retrieval call binding the contract method 0x22816a4c.
//
// Solidity: function validatorFeeReceiver() view returns(address)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolCallerSession) ValidatorFeeReceiver() (common.Address, error) {
	return _BridgeValidatorFeePool.Contract.ValidatorFeeReceiver(&_BridgeValidatorFeePool.CallOpts)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address _token) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactor) Collect(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.contract.Transact(opts, "collect", _token)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address _token) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) Collect(_token common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.Collect(&_BridgeValidatorFeePool.TransactOpts, _token)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address _token) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactorSession) Collect(_token common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.Collect(&_BridgeValidatorFeePool.TransactOpts, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _signerStorage, address _erc20Bridge, address _validatorFeeReceiver, uint256 _homeChainId) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactor) Initialize(opts *bind.TransactOpts, _signerStorage common.Address, _erc20Bridge common.Address, _validatorFeeReceiver common.Address, _homeChainId *big.Int) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.contract.Transact(opts, "initialize", _signerStorage, _erc20Bridge, _validatorFeeReceiver, _homeChainId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _signerStorage, address _erc20Bridge, address _validatorFeeReceiver, uint256 _homeChainId) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) Initialize(_signerStorage common.Address, _erc20Bridge common.Address, _validatorFeeReceiver common.Address, _homeChainId *big.Int) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.Initialize(&_BridgeValidatorFeePool.TransactOpts, _signerStorage, _erc20Bridge, _validatorFeeReceiver, _homeChainId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _signerStorage, address _erc20Bridge, address _validatorFeeReceiver, uint256 _homeChainId) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactorSession) Initialize(_signerStorage common.Address, _erc20Bridge common.Address, _validatorFeeReceiver common.Address, _homeChainId *big.Int) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.Initialize(&_BridgeValidatorFeePool.TransactOpts, _signerStorage, _erc20Bridge, _validatorFeeReceiver, _homeChainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.RenounceOwnership(&_BridgeValidatorFeePool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.RenounceOwnership(&_BridgeValidatorFeePool.TransactOpts)
}

// SetERC20Bridge is a paid mutator transaction binding the contract method 0xe88866a5.
//
// Solidity: function setERC20Bridge(address _erc20Bridge) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactor) SetERC20Bridge(opts *bind.TransactOpts, _erc20Bridge common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.contract.Transact(opts, "setERC20Bridge", _erc20Bridge)
}

// SetERC20Bridge is a paid mutator transaction binding the contract method 0xe88866a5.
//
// Solidity: function setERC20Bridge(address _erc20Bridge) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) SetERC20Bridge(_erc20Bridge common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.SetERC20Bridge(&_BridgeValidatorFeePool.TransactOpts, _erc20Bridge)
}

// SetERC20Bridge is a paid mutator transaction binding the contract method 0xe88866a5.
//
// Solidity: function setERC20Bridge(address _erc20Bridge) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactorSession) SetERC20Bridge(_erc20Bridge common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.SetERC20Bridge(&_BridgeValidatorFeePool.TransactOpts, _erc20Bridge)
}

// SetLimitPerToken is a paid mutator transaction binding the contract method 0x84d1a52a.
//
// Solidity: function setLimitPerToken(address _token, uint256 _limit) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactor) SetLimitPerToken(opts *bind.TransactOpts, _token common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.contract.Transact(opts, "setLimitPerToken", _token, _limit)
}

// SetLimitPerToken is a paid mutator transaction binding the contract method 0x84d1a52a.
//
// Solidity: function setLimitPerToken(address _token, uint256 _limit) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) SetLimitPerToken(_token common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.SetLimitPerToken(&_BridgeValidatorFeePool.TransactOpts, _token, _limit)
}

// SetLimitPerToken is a paid mutator transaction binding the contract method 0x84d1a52a.
//
// Solidity: function setLimitPerToken(address _token, uint256 _limit) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactorSession) SetLimitPerToken(_token common.Address, _limit *big.Int) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.SetLimitPerToken(&_BridgeValidatorFeePool.TransactOpts, _token, _limit)
}

// SetValidatorFeeReceiver is a paid mutator transaction binding the contract method 0x8a4b4e5a.
//
// Solidity: function setValidatorFeeReceiver(address _validatorFeeReceiver) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactor) SetValidatorFeeReceiver(opts *bind.TransactOpts, _validatorFeeReceiver common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.contract.Transact(opts, "setValidatorFeeReceiver", _validatorFeeReceiver)
}

// SetValidatorFeeReceiver is a paid mutator transaction binding the contract method 0x8a4b4e5a.
//
// Solidity: function setValidatorFeeReceiver(address _validatorFeeReceiver) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) SetValidatorFeeReceiver(_validatorFeeReceiver common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.SetValidatorFeeReceiver(&_BridgeValidatorFeePool.TransactOpts, _validatorFeeReceiver)
}

// SetValidatorFeeReceiver is a paid mutator transaction binding the contract method 0x8a4b4e5a.
//
// Solidity: function setValidatorFeeReceiver(address _validatorFeeReceiver) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactorSession) SetValidatorFeeReceiver(_validatorFeeReceiver common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.SetValidatorFeeReceiver(&_BridgeValidatorFeePool.TransactOpts, _validatorFeeReceiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.TransferOwnership(&_BridgeValidatorFeePool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.TransferOwnership(&_BridgeValidatorFeePool.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeValidatorFeePool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolSession) Receive() (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.Receive(&_BridgeValidatorFeePool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeValidatorFeePool.Contract.Receive(&_BridgeValidatorFeePool.TransactOpts)
}

// BridgeValidatorFeePoolCollectedIterator is returned from FilterCollected and is used to iterate over the raw logs and unpacked data for Collected events raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolCollectedIterator struct {
	Event *BridgeValidatorFeePoolCollected // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorFeePoolCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorFeePoolCollected)
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
		it.Event = new(BridgeValidatorFeePoolCollected)
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
func (it *BridgeValidatorFeePoolCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorFeePoolCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorFeePoolCollected represents a Collected event raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolCollected struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollected is a free log retrieval operation binding the contract event 0x8c22f554c81b54107cd95e734e4ef45767214974a540f34ea4a4f8c3fc7b13c3.
//
// Solidity: event Collected(address token, uint256 amount)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) FilterCollected(opts *bind.FilterOpts) (*BridgeValidatorFeePoolCollectedIterator, error) {

	logs, sub, err := _BridgeValidatorFeePool.contract.FilterLogs(opts, "Collected")
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorFeePoolCollectedIterator{contract: _BridgeValidatorFeePool.contract, event: "Collected", logs: logs, sub: sub}, nil
}

// WatchCollected is a free log subscription operation binding the contract event 0x8c22f554c81b54107cd95e734e4ef45767214974a540f34ea4a4f8c3fc7b13c3.
//
// Solidity: event Collected(address token, uint256 amount)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) WatchCollected(opts *bind.WatchOpts, sink chan<- *BridgeValidatorFeePoolCollected) (event.Subscription, error) {

	logs, sub, err := _BridgeValidatorFeePool.contract.WatchLogs(opts, "Collected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorFeePoolCollected)
				if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "Collected", log); err != nil {
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

// ParseCollected is a log parse operation binding the contract event 0x8c22f554c81b54107cd95e734e4ef45767214974a540f34ea4a4f8c3fc7b13c3.
//
// Solidity: event Collected(address token, uint256 amount)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) ParseCollected(log types.Log) (*BridgeValidatorFeePoolCollected, error) {
	event := new(BridgeValidatorFeePoolCollected)
	if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "Collected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeValidatorFeePoolERC20BridgeUpdatedIterator is returned from FilterERC20BridgeUpdated and is used to iterate over the raw logs and unpacked data for ERC20BridgeUpdated events raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolERC20BridgeUpdatedIterator struct {
	Event *BridgeValidatorFeePoolERC20BridgeUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorFeePoolERC20BridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorFeePoolERC20BridgeUpdated)
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
		it.Event = new(BridgeValidatorFeePoolERC20BridgeUpdated)
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
func (it *BridgeValidatorFeePoolERC20BridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorFeePoolERC20BridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorFeePoolERC20BridgeUpdated represents a ERC20BridgeUpdated event raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolERC20BridgeUpdated struct {
	Erc20Bridge common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC20BridgeUpdated is a free log retrieval operation binding the contract event 0x4dcff21652c7d0d1d66be34beea4eeea84364a02f61ed804cfd084772324e63b.
//
// Solidity: event ERC20BridgeUpdated(address erc20Bridge)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) FilterERC20BridgeUpdated(opts *bind.FilterOpts) (*BridgeValidatorFeePoolERC20BridgeUpdatedIterator, error) {

	logs, sub, err := _BridgeValidatorFeePool.contract.FilterLogs(opts, "ERC20BridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorFeePoolERC20BridgeUpdatedIterator{contract: _BridgeValidatorFeePool.contract, event: "ERC20BridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchERC20BridgeUpdated is a free log subscription operation binding the contract event 0x4dcff21652c7d0d1d66be34beea4eeea84364a02f61ed804cfd084772324e63b.
//
// Solidity: event ERC20BridgeUpdated(address erc20Bridge)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) WatchERC20BridgeUpdated(opts *bind.WatchOpts, sink chan<- *BridgeValidatorFeePoolERC20BridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeValidatorFeePool.contract.WatchLogs(opts, "ERC20BridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorFeePoolERC20BridgeUpdated)
				if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "ERC20BridgeUpdated", log); err != nil {
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

// ParseERC20BridgeUpdated is a log parse operation binding the contract event 0x4dcff21652c7d0d1d66be34beea4eeea84364a02f61ed804cfd084772324e63b.
//
// Solidity: event ERC20BridgeUpdated(address erc20Bridge)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) ParseERC20BridgeUpdated(log types.Log) (*BridgeValidatorFeePoolERC20BridgeUpdated, error) {
	event := new(BridgeValidatorFeePoolERC20BridgeUpdated)
	if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "ERC20BridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeValidatorFeePoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolInitializedIterator struct {
	Event *BridgeValidatorFeePoolInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorFeePoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorFeePoolInitialized)
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
		it.Event = new(BridgeValidatorFeePoolInitialized)
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
func (it *BridgeValidatorFeePoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorFeePoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorFeePoolInitialized represents a Initialized event raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeValidatorFeePoolInitializedIterator, error) {

	logs, sub, err := _BridgeValidatorFeePool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorFeePoolInitializedIterator{contract: _BridgeValidatorFeePool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeValidatorFeePoolInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeValidatorFeePool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorFeePoolInitialized)
				if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) ParseInitialized(log types.Log) (*BridgeValidatorFeePoolInitialized, error) {
	event := new(BridgeValidatorFeePoolInitialized)
	if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeValidatorFeePoolLimitPerTokenUpdatedIterator is returned from FilterLimitPerTokenUpdated and is used to iterate over the raw logs and unpacked data for LimitPerTokenUpdated events raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolLimitPerTokenUpdatedIterator struct {
	Event *BridgeValidatorFeePoolLimitPerTokenUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorFeePoolLimitPerTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorFeePoolLimitPerTokenUpdated)
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
		it.Event = new(BridgeValidatorFeePoolLimitPerTokenUpdated)
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
func (it *BridgeValidatorFeePoolLimitPerTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorFeePoolLimitPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorFeePoolLimitPerTokenUpdated represents a LimitPerTokenUpdated event raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolLimitPerTokenUpdated struct {
	Token common.Address
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLimitPerTokenUpdated is a free log retrieval operation binding the contract event 0x335be255e3120c1ab7d33be1ca6907a55bc81129991ef594795b796932d431d9.
//
// Solidity: event LimitPerTokenUpdated(address token, uint256 limit)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) FilterLimitPerTokenUpdated(opts *bind.FilterOpts) (*BridgeValidatorFeePoolLimitPerTokenUpdatedIterator, error) {

	logs, sub, err := _BridgeValidatorFeePool.contract.FilterLogs(opts, "LimitPerTokenUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorFeePoolLimitPerTokenUpdatedIterator{contract: _BridgeValidatorFeePool.contract, event: "LimitPerTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchLimitPerTokenUpdated is a free log subscription operation binding the contract event 0x335be255e3120c1ab7d33be1ca6907a55bc81129991ef594795b796932d431d9.
//
// Solidity: event LimitPerTokenUpdated(address token, uint256 limit)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) WatchLimitPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *BridgeValidatorFeePoolLimitPerTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeValidatorFeePool.contract.WatchLogs(opts, "LimitPerTokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorFeePoolLimitPerTokenUpdated)
				if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "LimitPerTokenUpdated", log); err != nil {
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

// ParseLimitPerTokenUpdated is a log parse operation binding the contract event 0x335be255e3120c1ab7d33be1ca6907a55bc81129991ef594795b796932d431d9.
//
// Solidity: event LimitPerTokenUpdated(address token, uint256 limit)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) ParseLimitPerTokenUpdated(log types.Log) (*BridgeValidatorFeePoolLimitPerTokenUpdated, error) {
	event := new(BridgeValidatorFeePoolLimitPerTokenUpdated)
	if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "LimitPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeValidatorFeePoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolOwnershipTransferredIterator struct {
	Event *BridgeValidatorFeePoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorFeePoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorFeePoolOwnershipTransferred)
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
		it.Event = new(BridgeValidatorFeePoolOwnershipTransferred)
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
func (it *BridgeValidatorFeePoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorFeePoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorFeePoolOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeValidatorFeePoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeValidatorFeePool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorFeePoolOwnershipTransferredIterator{contract: _BridgeValidatorFeePool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeValidatorFeePoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeValidatorFeePool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorFeePoolOwnershipTransferred)
				if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeValidatorFeePoolOwnershipTransferred, error) {
	event := new(BridgeValidatorFeePoolOwnershipTransferred)
	if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeValidatorFeePoolValidatorFeeReceiverUpdatedIterator is returned from FilterValidatorFeeReceiverUpdated and is used to iterate over the raw logs and unpacked data for ValidatorFeeReceiverUpdated events raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolValidatorFeeReceiverUpdatedIterator struct {
	Event *BridgeValidatorFeePoolValidatorFeeReceiverUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeValidatorFeePoolValidatorFeeReceiverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeValidatorFeePoolValidatorFeeReceiverUpdated)
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
		it.Event = new(BridgeValidatorFeePoolValidatorFeeReceiverUpdated)
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
func (it *BridgeValidatorFeePoolValidatorFeeReceiverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeValidatorFeePoolValidatorFeeReceiverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeValidatorFeePoolValidatorFeeReceiverUpdated represents a ValidatorFeeReceiverUpdated event raised by the BridgeValidatorFeePool contract.
type BridgeValidatorFeePoolValidatorFeeReceiverUpdated struct {
	ValidatorFeeReceiver common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterValidatorFeeReceiverUpdated is a free log retrieval operation binding the contract event 0x94a216cae37990b5cdf93069f4dc84904e2448bd81b87b8ea6510086f623d201.
//
// Solidity: event ValidatorFeeReceiverUpdated(address validatorFeeReceiver)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) FilterValidatorFeeReceiverUpdated(opts *bind.FilterOpts) (*BridgeValidatorFeePoolValidatorFeeReceiverUpdatedIterator, error) {

	logs, sub, err := _BridgeValidatorFeePool.contract.FilterLogs(opts, "ValidatorFeeReceiverUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeValidatorFeePoolValidatorFeeReceiverUpdatedIterator{contract: _BridgeValidatorFeePool.contract, event: "ValidatorFeeReceiverUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorFeeReceiverUpdated is a free log subscription operation binding the contract event 0x94a216cae37990b5cdf93069f4dc84904e2448bd81b87b8ea6510086f623d201.
//
// Solidity: event ValidatorFeeReceiverUpdated(address validatorFeeReceiver)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) WatchValidatorFeeReceiverUpdated(opts *bind.WatchOpts, sink chan<- *BridgeValidatorFeePoolValidatorFeeReceiverUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeValidatorFeePool.contract.WatchLogs(opts, "ValidatorFeeReceiverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeValidatorFeePoolValidatorFeeReceiverUpdated)
				if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "ValidatorFeeReceiverUpdated", log); err != nil {
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

// ParseValidatorFeeReceiverUpdated is a log parse operation binding the contract event 0x94a216cae37990b5cdf93069f4dc84904e2448bd81b87b8ea6510086f623d201.
//
// Solidity: event ValidatorFeeReceiverUpdated(address validatorFeeReceiver)
func (_BridgeValidatorFeePool *BridgeValidatorFeePoolFilterer) ParseValidatorFeeReceiverUpdated(log types.Log) (*BridgeValidatorFeePoolValidatorFeeReceiverUpdated, error) {
	event := new(BridgeValidatorFeePoolValidatorFeeReceiverUpdated)
	if err := _BridgeValidatorFeePool.contract.UnpackLog(event, "ValidatorFeeReceiverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
