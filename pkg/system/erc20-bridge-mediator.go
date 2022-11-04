// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package system

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

// ERC20BridgeMediatorMetaData contains all meta data concerning the ERC20BridgeMediator contract.
var ERC20BridgeMediatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mediatorAddress\",\"type\":\"address\"}],\"name\":\"MediatorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sourceData\",\"type\":\"bytes\"}],\"name\":\"mediate\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"symbolToToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenToSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20BridgeMediatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BridgeMediatorMetaData.ABI instead.
var ERC20BridgeMediatorABI = ERC20BridgeMediatorMetaData.ABI

// ERC20BridgeMediator is an auto generated Go binding around an Ethereum contract.
type ERC20BridgeMediator struct {
	ERC20BridgeMediatorCaller     // Read-only binding to the contract
	ERC20BridgeMediatorTransactor // Write-only binding to the contract
	ERC20BridgeMediatorFilterer   // Log filterer for contract events
}

// ERC20BridgeMediatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BridgeMediatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeMediatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BridgeMediatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeMediatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BridgeMediatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeMediatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BridgeMediatorSession struct {
	Contract     *ERC20BridgeMediator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20BridgeMediatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BridgeMediatorCallerSession struct {
	Contract *ERC20BridgeMediatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ERC20BridgeMediatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BridgeMediatorTransactorSession struct {
	Contract     *ERC20BridgeMediatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ERC20BridgeMediatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BridgeMediatorRaw struct {
	Contract *ERC20BridgeMediator // Generic contract binding to access the raw methods on
}

// ERC20BridgeMediatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BridgeMediatorCallerRaw struct {
	Contract *ERC20BridgeMediatorCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BridgeMediatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BridgeMediatorTransactorRaw struct {
	Contract *ERC20BridgeMediatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20BridgeMediator creates a new instance of ERC20BridgeMediator, bound to a specific deployed contract.
func NewERC20BridgeMediator(address common.Address, backend bind.ContractBackend) (*ERC20BridgeMediator, error) {
	contract, err := bindERC20BridgeMediator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeMediator{ERC20BridgeMediatorCaller: ERC20BridgeMediatorCaller{contract: contract}, ERC20BridgeMediatorTransactor: ERC20BridgeMediatorTransactor{contract: contract}, ERC20BridgeMediatorFilterer: ERC20BridgeMediatorFilterer{contract: contract}}, nil
}

// NewERC20BridgeMediatorCaller creates a new read-only instance of ERC20BridgeMediator, bound to a specific deployed contract.
func NewERC20BridgeMediatorCaller(address common.Address, caller bind.ContractCaller) (*ERC20BridgeMediatorCaller, error) {
	contract, err := bindERC20BridgeMediator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeMediatorCaller{contract: contract}, nil
}

// NewERC20BridgeMediatorTransactor creates a new write-only instance of ERC20BridgeMediator, bound to a specific deployed contract.
func NewERC20BridgeMediatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BridgeMediatorTransactor, error) {
	contract, err := bindERC20BridgeMediator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeMediatorTransactor{contract: contract}, nil
}

// NewERC20BridgeMediatorFilterer creates a new log filterer instance of ERC20BridgeMediator, bound to a specific deployed contract.
func NewERC20BridgeMediatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BridgeMediatorFilterer, error) {
	contract, err := bindERC20BridgeMediator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeMediatorFilterer{contract: contract}, nil
}

// bindERC20BridgeMediator binds a generic wrapper to an already deployed contract.
func bindERC20BridgeMediator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BridgeMediatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20BridgeMediator *ERC20BridgeMediatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20BridgeMediator.Contract.ERC20BridgeMediatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20BridgeMediator *ERC20BridgeMediatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.ERC20BridgeMediatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20BridgeMediator *ERC20BridgeMediatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.ERC20BridgeMediatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20BridgeMediator *ERC20BridgeMediatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20BridgeMediator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20BridgeMediator *ERC20BridgeMediatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20BridgeMediator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20BridgeMediator *ERC20BridgeMediatorSession) Owner() (common.Address, error) {
	return _ERC20BridgeMediator.Contract.Owner(&_ERC20BridgeMediator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20BridgeMediator *ERC20BridgeMediatorCallerSession) Owner() (common.Address, error) {
	return _ERC20BridgeMediator.Contract.Owner(&_ERC20BridgeMediator.CallOpts)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa20126ff.
//
// Solidity: function symbolToToken(uint256 , string ) view returns(address)
func (_ERC20BridgeMediator *ERC20BridgeMediatorCaller) SymbolToToken(opts *bind.CallOpts, arg0 *big.Int, arg1 string) (common.Address, error) {
	var out []interface{}
	err := _ERC20BridgeMediator.contract.Call(opts, &out, "symbolToToken", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SymbolToToken is a free data retrieval call binding the contract method 0xa20126ff.
//
// Solidity: function symbolToToken(uint256 , string ) view returns(address)
func (_ERC20BridgeMediator *ERC20BridgeMediatorSession) SymbolToToken(arg0 *big.Int, arg1 string) (common.Address, error) {
	return _ERC20BridgeMediator.Contract.SymbolToToken(&_ERC20BridgeMediator.CallOpts, arg0, arg1)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa20126ff.
//
// Solidity: function symbolToToken(uint256 , string ) view returns(address)
func (_ERC20BridgeMediator *ERC20BridgeMediatorCallerSession) SymbolToToken(arg0 *big.Int, arg1 string) (common.Address, error) {
	return _ERC20BridgeMediator.Contract.SymbolToToken(&_ERC20BridgeMediator.CallOpts, arg0, arg1)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0xeff7bb86.
//
// Solidity: function tokenToSymbol(uint256 , address ) view returns(string)
func (_ERC20BridgeMediator *ERC20BridgeMediatorCaller) TokenToSymbol(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (string, error) {
	var out []interface{}
	err := _ERC20BridgeMediator.contract.Call(opts, &out, "tokenToSymbol", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenToSymbol is a free data retrieval call binding the contract method 0xeff7bb86.
//
// Solidity: function tokenToSymbol(uint256 , address ) view returns(string)
func (_ERC20BridgeMediator *ERC20BridgeMediatorSession) TokenToSymbol(arg0 *big.Int, arg1 common.Address) (string, error) {
	return _ERC20BridgeMediator.Contract.TokenToSymbol(&_ERC20BridgeMediator.CallOpts, arg0, arg1)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0xeff7bb86.
//
// Solidity: function tokenToSymbol(uint256 , address ) view returns(string)
func (_ERC20BridgeMediator *ERC20BridgeMediatorCallerSession) TokenToSymbol(arg0 *big.Int, arg1 common.Address) (string, error) {
	return _ERC20BridgeMediator.Contract.TokenToSymbol(&_ERC20BridgeMediator.CallOpts, arg0, arg1)
}

// AddToken is a paid mutator transaction binding the contract method 0x11457604.
//
// Solidity: function addToken(string symbol, uint256 chainId, address token) returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactor) AddToken(opts *bind.TransactOpts, symbol string, chainId *big.Int, token common.Address) (*types.Transaction, error) {
	return _ERC20BridgeMediator.contract.Transact(opts, "addToken", symbol, chainId, token)
}

// AddToken is a paid mutator transaction binding the contract method 0x11457604.
//
// Solidity: function addToken(string symbol, uint256 chainId, address token) returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorSession) AddToken(symbol string, chainId *big.Int, token common.Address) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.AddToken(&_ERC20BridgeMediator.TransactOpts, symbol, chainId, token)
}

// AddToken is a paid mutator transaction binding the contract method 0x11457604.
//
// Solidity: function addToken(string symbol, uint256 chainId, address token) returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactorSession) AddToken(symbol string, chainId *big.Int, token common.Address) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.AddToken(&_ERC20BridgeMediator.TransactOpts, symbol, chainId, token)
}

// Mediate is a paid mutator transaction binding the contract method 0x8b68e0f7.
//
// Solidity: function mediate(uint256 sourceChain, uint256 destinationChain, bytes sourceData) returns(bytes)
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactor) Mediate(opts *bind.TransactOpts, sourceChain *big.Int, destinationChain *big.Int, sourceData []byte) (*types.Transaction, error) {
	return _ERC20BridgeMediator.contract.Transact(opts, "mediate", sourceChain, destinationChain, sourceData)
}

// Mediate is a paid mutator transaction binding the contract method 0x8b68e0f7.
//
// Solidity: function mediate(uint256 sourceChain, uint256 destinationChain, bytes sourceData) returns(bytes)
func (_ERC20BridgeMediator *ERC20BridgeMediatorSession) Mediate(sourceChain *big.Int, destinationChain *big.Int, sourceData []byte) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.Mediate(&_ERC20BridgeMediator.TransactOpts, sourceChain, destinationChain, sourceData)
}

// Mediate is a paid mutator transaction binding the contract method 0x8b68e0f7.
//
// Solidity: function mediate(uint256 sourceChain, uint256 destinationChain, bytes sourceData) returns(bytes)
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactorSession) Mediate(sourceChain *big.Int, destinationChain *big.Int, sourceData []byte) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.Mediate(&_ERC20BridgeMediator.TransactOpts, sourceChain, destinationChain, sourceData)
}

// RemoveToken is a paid mutator transaction binding the contract method 0xe409ed98.
//
// Solidity: function removeToken(string symbol, uint256 chainId) returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactor) RemoveToken(opts *bind.TransactOpts, symbol string, chainId *big.Int) (*types.Transaction, error) {
	return _ERC20BridgeMediator.contract.Transact(opts, "removeToken", symbol, chainId)
}

// RemoveToken is a paid mutator transaction binding the contract method 0xe409ed98.
//
// Solidity: function removeToken(string symbol, uint256 chainId) returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorSession) RemoveToken(symbol string, chainId *big.Int) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.RemoveToken(&_ERC20BridgeMediator.TransactOpts, symbol, chainId)
}

// RemoveToken is a paid mutator transaction binding the contract method 0xe409ed98.
//
// Solidity: function removeToken(string symbol, uint256 chainId) returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactorSession) RemoveToken(symbol string, chainId *big.Int) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.RemoveToken(&_ERC20BridgeMediator.TransactOpts, symbol, chainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20BridgeMediator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.RenounceOwnership(&_ERC20BridgeMediator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.RenounceOwnership(&_ERC20BridgeMediator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20BridgeMediator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.TransferOwnership(&_ERC20BridgeMediator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20BridgeMediator *ERC20BridgeMediatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20BridgeMediator.Contract.TransferOwnership(&_ERC20BridgeMediator.TransactOpts, newOwner)
}

// ERC20BridgeMediatorAddedTokenIterator is returned from FilterAddedToken and is used to iterate over the raw logs and unpacked data for AddedToken events raised by the ERC20BridgeMediator contract.
type ERC20BridgeMediatorAddedTokenIterator struct {
	Event *ERC20BridgeMediatorAddedToken // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeMediatorAddedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeMediatorAddedToken)
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
		it.Event = new(ERC20BridgeMediatorAddedToken)
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
func (it *ERC20BridgeMediatorAddedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeMediatorAddedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeMediatorAddedToken represents a AddedToken event raised by the ERC20BridgeMediator contract.
type ERC20BridgeMediatorAddedToken struct {
	Symbol  string
	ChainId *big.Int
	Token   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddedToken is a free log retrieval operation binding the contract event 0x1eea762d9514744d01160b932a7a6c1cbee2788c286607b3cc00999f6a530ea9.
//
// Solidity: event AddedToken(string symbol, uint256 chainId, address token)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) FilterAddedToken(opts *bind.FilterOpts) (*ERC20BridgeMediatorAddedTokenIterator, error) {

	logs, sub, err := _ERC20BridgeMediator.contract.FilterLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeMediatorAddedTokenIterator{contract: _ERC20BridgeMediator.contract, event: "AddedToken", logs: logs, sub: sub}, nil
}

// WatchAddedToken is a free log subscription operation binding the contract event 0x1eea762d9514744d01160b932a7a6c1cbee2788c286607b3cc00999f6a530ea9.
//
// Solidity: event AddedToken(string symbol, uint256 chainId, address token)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) WatchAddedToken(opts *bind.WatchOpts, sink chan<- *ERC20BridgeMediatorAddedToken) (event.Subscription, error) {

	logs, sub, err := _ERC20BridgeMediator.contract.WatchLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeMediatorAddedToken)
				if err := _ERC20BridgeMediator.contract.UnpackLog(event, "AddedToken", log); err != nil {
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

// ParseAddedToken is a log parse operation binding the contract event 0x1eea762d9514744d01160b932a7a6c1cbee2788c286607b3cc00999f6a530ea9.
//
// Solidity: event AddedToken(string symbol, uint256 chainId, address token)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) ParseAddedToken(log types.Log) (*ERC20BridgeMediatorAddedToken, error) {
	event := new(ERC20BridgeMediatorAddedToken)
	if err := _ERC20BridgeMediator.contract.UnpackLog(event, "AddedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeMediatorMediatorAddressIterator is returned from FilterMediatorAddress and is used to iterate over the raw logs and unpacked data for MediatorAddress events raised by the ERC20BridgeMediator contract.
type ERC20BridgeMediatorMediatorAddressIterator struct {
	Event *ERC20BridgeMediatorMediatorAddress // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeMediatorMediatorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeMediatorMediatorAddress)
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
		it.Event = new(ERC20BridgeMediatorMediatorAddress)
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
func (it *ERC20BridgeMediatorMediatorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeMediatorMediatorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeMediatorMediatorAddress represents a MediatorAddress event raised by the ERC20BridgeMediator contract.
type ERC20BridgeMediatorMediatorAddress struct {
	MediatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMediatorAddress is a free log retrieval operation binding the contract event 0xbb7882191c73e87a43b5ce5731ce01aceb048e6cafc123d2122cc3d365ad7c31.
//
// Solidity: event MediatorAddress(address mediatorAddress)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) FilterMediatorAddress(opts *bind.FilterOpts) (*ERC20BridgeMediatorMediatorAddressIterator, error) {

	logs, sub, err := _ERC20BridgeMediator.contract.FilterLogs(opts, "MediatorAddress")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeMediatorMediatorAddressIterator{contract: _ERC20BridgeMediator.contract, event: "MediatorAddress", logs: logs, sub: sub}, nil
}

// WatchMediatorAddress is a free log subscription operation binding the contract event 0xbb7882191c73e87a43b5ce5731ce01aceb048e6cafc123d2122cc3d365ad7c31.
//
// Solidity: event MediatorAddress(address mediatorAddress)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) WatchMediatorAddress(opts *bind.WatchOpts, sink chan<- *ERC20BridgeMediatorMediatorAddress) (event.Subscription, error) {

	logs, sub, err := _ERC20BridgeMediator.contract.WatchLogs(opts, "MediatorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeMediatorMediatorAddress)
				if err := _ERC20BridgeMediator.contract.UnpackLog(event, "MediatorAddress", log); err != nil {
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

// ParseMediatorAddress is a log parse operation binding the contract event 0xbb7882191c73e87a43b5ce5731ce01aceb048e6cafc123d2122cc3d365ad7c31.
//
// Solidity: event MediatorAddress(address mediatorAddress)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) ParseMediatorAddress(log types.Log) (*ERC20BridgeMediatorMediatorAddress, error) {
	event := new(ERC20BridgeMediatorMediatorAddress)
	if err := _ERC20BridgeMediator.contract.UnpackLog(event, "MediatorAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeMediatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20BridgeMediator contract.
type ERC20BridgeMediatorOwnershipTransferredIterator struct {
	Event *ERC20BridgeMediatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeMediatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeMediatorOwnershipTransferred)
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
		it.Event = new(ERC20BridgeMediatorOwnershipTransferred)
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
func (it *ERC20BridgeMediatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeMediatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeMediatorOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20BridgeMediator contract.
type ERC20BridgeMediatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20BridgeMediatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20BridgeMediator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeMediatorOwnershipTransferredIterator{contract: _ERC20BridgeMediator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20BridgeMediatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20BridgeMediator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeMediatorOwnershipTransferred)
				if err := _ERC20BridgeMediator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20BridgeMediatorOwnershipTransferred, error) {
	event := new(ERC20BridgeMediatorOwnershipTransferred)
	if err := _ERC20BridgeMediator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeMediatorRemovedTokenIterator is returned from FilterRemovedToken and is used to iterate over the raw logs and unpacked data for RemovedToken events raised by the ERC20BridgeMediator contract.
type ERC20BridgeMediatorRemovedTokenIterator struct {
	Event *ERC20BridgeMediatorRemovedToken // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeMediatorRemovedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeMediatorRemovedToken)
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
		it.Event = new(ERC20BridgeMediatorRemovedToken)
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
func (it *ERC20BridgeMediatorRemovedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeMediatorRemovedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeMediatorRemovedToken represents a RemovedToken event raised by the ERC20BridgeMediator contract.
type ERC20BridgeMediatorRemovedToken struct {
	Symbol  string
	ChainId *big.Int
	Token   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedToken is a free log retrieval operation binding the contract event 0x747510f5d47eee15c8194e6944a1339805da89ef768309e849b2aacde53f3b44.
//
// Solidity: event RemovedToken(string symbol, uint256 chainId, address token)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) FilterRemovedToken(opts *bind.FilterOpts) (*ERC20BridgeMediatorRemovedTokenIterator, error) {

	logs, sub, err := _ERC20BridgeMediator.contract.FilterLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeMediatorRemovedTokenIterator{contract: _ERC20BridgeMediator.contract, event: "RemovedToken", logs: logs, sub: sub}, nil
}

// WatchRemovedToken is a free log subscription operation binding the contract event 0x747510f5d47eee15c8194e6944a1339805da89ef768309e849b2aacde53f3b44.
//
// Solidity: event RemovedToken(string symbol, uint256 chainId, address token)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) WatchRemovedToken(opts *bind.WatchOpts, sink chan<- *ERC20BridgeMediatorRemovedToken) (event.Subscription, error) {

	logs, sub, err := _ERC20BridgeMediator.contract.WatchLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeMediatorRemovedToken)
				if err := _ERC20BridgeMediator.contract.UnpackLog(event, "RemovedToken", log); err != nil {
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

// ParseRemovedToken is a log parse operation binding the contract event 0x747510f5d47eee15c8194e6944a1339805da89ef768309e849b2aacde53f3b44.
//
// Solidity: event RemovedToken(string symbol, uint256 chainId, address token)
func (_ERC20BridgeMediator *ERC20BridgeMediatorFilterer) ParseRemovedToken(log types.Log) (*ERC20BridgeMediatorRemovedToken, error) {
	event := new(ERC20BridgeMediatorRemovedToken)
	if err := _ERC20BridgeMediator.contract.UnpackLog(event, "RemovedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
