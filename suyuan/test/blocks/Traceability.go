// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blocks

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

// GoodsTraceData is an auto generated low-level Go binding around an user-defined struct.
type GoodsTraceData struct {
	Addr      string
	Status    uint16
	Timestamp *big.Int
	Data      string
}

// TraceabilityMetaData contains all meta data concerning the Traceability contract.
var TraceabilityMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"goodsTp\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_addr\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"goodsid\",\"type\":\"string\"}],\"name\":\"newGoodsEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_passwd\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"goodsStatus\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"goodsID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newaddr\",\"type\":\"string\"}],\"name\":\"changeGoodsStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_passwd\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"goodsID\",\"type\":\"string\"}],\"name\":\"createGoods\",\"outputs\":[{\"internalType\":\"contractGoods\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"goodsId\",\"type\":\"string\"}],\"name\":\"getGoods\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"status\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structGoods.TraceData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"goodsId\",\"type\":\"string\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraceabilityABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceabilityMetaData.ABI instead.
var TraceabilityABI = TraceabilityMetaData.ABI

// Traceability is an auto generated Go binding around an Ethereum contract.
type Traceability struct {
	TraceabilityCaller     // Read-only binding to the contract
	TraceabilityTransactor // Write-only binding to the contract
	TraceabilityFilterer   // Log filterer for contract events
}

// TraceabilityCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraceabilityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceabilityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraceabilityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceabilityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraceabilityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceabilitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraceabilitySession struct {
	Contract     *Traceability     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceabilityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraceabilityCallerSession struct {
	Contract *TraceabilityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TraceabilityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraceabilityTransactorSession struct {
	Contract     *TraceabilityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TraceabilityRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraceabilityRaw struct {
	Contract *Traceability // Generic contract binding to access the raw methods on
}

// TraceabilityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraceabilityCallerRaw struct {
	Contract *TraceabilityCaller // Generic read-only contract binding to access the raw methods on
}

// TraceabilityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraceabilityTransactorRaw struct {
	Contract *TraceabilityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraceability creates a new instance of Traceability, bound to a specific deployed contract.
func NewTraceability(address common.Address, backend bind.ContractBackend) (*Traceability, error) {
	contract, err := bindTraceability(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Traceability{TraceabilityCaller: TraceabilityCaller{contract: contract}, TraceabilityTransactor: TraceabilityTransactor{contract: contract}, TraceabilityFilterer: TraceabilityFilterer{contract: contract}}, nil
}

// NewTraceabilityCaller creates a new read-only instance of Traceability, bound to a specific deployed contract.
func NewTraceabilityCaller(address common.Address, caller bind.ContractCaller) (*TraceabilityCaller, error) {
	contract, err := bindTraceability(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraceabilityCaller{contract: contract}, nil
}

// NewTraceabilityTransactor creates a new write-only instance of Traceability, bound to a specific deployed contract.
func NewTraceabilityTransactor(address common.Address, transactor bind.ContractTransactor) (*TraceabilityTransactor, error) {
	contract, err := bindTraceability(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraceabilityTransactor{contract: contract}, nil
}

// NewTraceabilityFilterer creates a new log filterer instance of Traceability, bound to a specific deployed contract.
func NewTraceabilityFilterer(address common.Address, filterer bind.ContractFilterer) (*TraceabilityFilterer, error) {
	contract, err := bindTraceability(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraceabilityFilterer{contract: contract}, nil
}

// bindTraceability binds a generic wrapper to an already deployed contract.
func bindTraceability(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TraceabilityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Traceability *TraceabilityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Traceability.Contract.TraceabilityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Traceability *TraceabilityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Traceability.Contract.TraceabilityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Traceability *TraceabilityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Traceability.Contract.TraceabilityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Traceability *TraceabilityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Traceability.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Traceability *TraceabilityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Traceability.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Traceability *TraceabilityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Traceability.Contract.contract.Transact(opts, method, params...)
}

// GetGoods is a free data retrieval call binding the contract method 0xaf64c789.
//
// Solidity: function getGoods(string goodsId) view returns((string,uint16,uint256,string)[])
func (_Traceability *TraceabilityCaller) GetGoods(opts *bind.CallOpts, goodsId string) ([]GoodsTraceData, error) {
	var out []interface{}
	err := _Traceability.contract.Call(opts, &out, "getGoods", goodsId)

	if err != nil {
		return *new([]GoodsTraceData), err
	}

	out0 := *abi.ConvertType(out[0], new([]GoodsTraceData)).(*[]GoodsTraceData)

	return out0, err

}

// GetGoods is a free data retrieval call binding the contract method 0xaf64c789.
//
// Solidity: function getGoods(string goodsId) view returns((string,uint16,uint256,string)[])
func (_Traceability *TraceabilitySession) GetGoods(goodsId string) ([]GoodsTraceData, error) {
	return _Traceability.Contract.GetGoods(&_Traceability.CallOpts, goodsId)
}

// GetGoods is a free data retrieval call binding the contract method 0xaf64c789.
//
// Solidity: function getGoods(string goodsId) view returns((string,uint16,uint256,string)[])
func (_Traceability *TraceabilityCallerSession) GetGoods(goodsId string) ([]GoodsTraceData, error) {
	return _Traceability.Contract.GetGoods(&_Traceability.CallOpts, goodsId)
}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string goodsId) view returns(uint16)
func (_Traceability *TraceabilityCaller) GetStatus(opts *bind.CallOpts, goodsId string) (uint16, error) {
	var out []interface{}
	err := _Traceability.contract.Call(opts, &out, "getStatus", goodsId)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string goodsId) view returns(uint16)
func (_Traceability *TraceabilitySession) GetStatus(goodsId string) (uint16, error) {
	return _Traceability.Contract.GetStatus(&_Traceability.CallOpts, goodsId)
}

// GetStatus is a free data retrieval call binding the contract method 0x22b05ed2.
//
// Solidity: function getStatus(string goodsId) view returns(uint16)
func (_Traceability *TraceabilityCallerSession) GetStatus(goodsId string) (uint16, error) {
	return _Traceability.Contract.GetStatus(&_Traceability.CallOpts, goodsId)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_Traceability *TraceabilityCaller) Login(opts *bind.CallOpts, username string, passwd string) (bool, error) {
	var out []interface{}
	err := _Traceability.contract.Call(opts, &out, "login", username, passwd)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_Traceability *TraceabilitySession) Login(username string, passwd string) (bool, error) {
	return _Traceability.Contract.Login(&_Traceability.CallOpts, username, passwd)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string username, string passwd) view returns(bool)
func (_Traceability *TraceabilityCallerSession) Login(username string, passwd string) (bool, error) {
	return _Traceability.Contract.Login(&_Traceability.CallOpts, username, passwd)
}

// ChangeGoodsStatus is a paid mutator transaction binding the contract method 0x42279385.
//
// Solidity: function changeGoodsStatus(string _username, string _passwd, uint16 goodsStatus, string data, string goodsID, string newaddr) returns()
func (_Traceability *TraceabilityTransactor) ChangeGoodsStatus(opts *bind.TransactOpts, _username string, _passwd string, goodsStatus uint16, data string, goodsID string, newaddr string) (*types.Transaction, error) {
	return _Traceability.contract.Transact(opts, "changeGoodsStatus", _username, _passwd, goodsStatus, data, goodsID, newaddr)
}

// ChangeGoodsStatus is a paid mutator transaction binding the contract method 0x42279385.
//
// Solidity: function changeGoodsStatus(string _username, string _passwd, uint16 goodsStatus, string data, string goodsID, string newaddr) returns()
func (_Traceability *TraceabilitySession) ChangeGoodsStatus(_username string, _passwd string, goodsStatus uint16, data string, goodsID string, newaddr string) (*types.Transaction, error) {
	return _Traceability.Contract.ChangeGoodsStatus(&_Traceability.TransactOpts, _username, _passwd, goodsStatus, data, goodsID, newaddr)
}

// ChangeGoodsStatus is a paid mutator transaction binding the contract method 0x42279385.
//
// Solidity: function changeGoodsStatus(string _username, string _passwd, uint16 goodsStatus, string data, string goodsID, string newaddr) returns()
func (_Traceability *TraceabilityTransactorSession) ChangeGoodsStatus(_username string, _passwd string, goodsStatus uint16, data string, goodsID string, newaddr string) (*types.Transaction, error) {
	return _Traceability.Contract.ChangeGoodsStatus(&_Traceability.TransactOpts, _username, _passwd, goodsStatus, data, goodsID, newaddr)
}

// CreateGoods is a paid mutator transaction binding the contract method 0x7ea69e9f.
//
// Solidity: function createGoods(string _username, string _passwd, string data, string goodsID) returns(address)
func (_Traceability *TraceabilityTransactor) CreateGoods(opts *bind.TransactOpts, _username string, _passwd string, data string, goodsID string) (*types.Transaction, error) {
	return _Traceability.contract.Transact(opts, "createGoods", _username, _passwd, data, goodsID)
}

// CreateGoods is a paid mutator transaction binding the contract method 0x7ea69e9f.
//
// Solidity: function createGoods(string _username, string _passwd, string data, string goodsID) returns(address)
func (_Traceability *TraceabilitySession) CreateGoods(_username string, _passwd string, data string, goodsID string) (*types.Transaction, error) {
	return _Traceability.Contract.CreateGoods(&_Traceability.TransactOpts, _username, _passwd, data, goodsID)
}

// CreateGoods is a paid mutator transaction binding the contract method 0x7ea69e9f.
//
// Solidity: function createGoods(string _username, string _passwd, string data, string goodsID) returns(address)
func (_Traceability *TraceabilityTransactorSession) CreateGoods(_username string, _passwd string, data string, goodsID string) (*types.Transaction, error) {
	return _Traceability.Contract.CreateGoods(&_Traceability.TransactOpts, _username, _passwd, data, goodsID)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Traceability *TraceabilityTransactor) Register(opts *bind.TransactOpts, username string, passwd string) (*types.Transaction, error) {
	return _Traceability.contract.Transact(opts, "register", username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Traceability *TraceabilitySession) Register(username string, passwd string) (*types.Transaction, error) {
	return _Traceability.Contract.Register(&_Traceability.TransactOpts, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string username, string passwd) returns()
func (_Traceability *TraceabilityTransactorSession) Register(username string, passwd string) (*types.Transaction, error) {
	return _Traceability.Contract.Register(&_Traceability.TransactOpts, username, passwd)
}

// TraceabilityNewGoodsEventIterator is returned from FilterNewGoodsEvent and is used to iterate over the raw logs and unpacked data for NewGoodsEvent events raised by the Traceability contract.
type TraceabilityNewGoodsEventIterator struct {
	Event *TraceabilityNewGoodsEvent // Event containing the contract specifics and raw log

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
func (it *TraceabilityNewGoodsEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraceabilityNewGoodsEvent)
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
		it.Event = new(TraceabilityNewGoodsEvent)
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
func (it *TraceabilityNewGoodsEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraceabilityNewGoodsEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraceabilityNewGoodsEvent represents a NewGoodsEvent event raised by the Traceability contract.
type TraceabilityNewGoodsEvent struct {
	Addr    string
	Data    string
	Goodsid string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewGoodsEvent is a free log retrieval operation binding the contract event 0x30d18c7c4624d4b6e90e21ae02bab5166256618bc6876873dd5ba8c5ed3db86f.
//
// Solidity: event newGoodsEvent(string _addr, string _data, string goodsid)
func (_Traceability *TraceabilityFilterer) FilterNewGoodsEvent(opts *bind.FilterOpts) (*TraceabilityNewGoodsEventIterator, error) {

	logs, sub, err := _Traceability.contract.FilterLogs(opts, "newGoodsEvent")
	if err != nil {
		return nil, err
	}
	return &TraceabilityNewGoodsEventIterator{contract: _Traceability.contract, event: "newGoodsEvent", logs: logs, sub: sub}, nil
}

// WatchNewGoodsEvent is a free log subscription operation binding the contract event 0x30d18c7c4624d4b6e90e21ae02bab5166256618bc6876873dd5ba8c5ed3db86f.
//
// Solidity: event newGoodsEvent(string _addr, string _data, string goodsid)
func (_Traceability *TraceabilityFilterer) WatchNewGoodsEvent(opts *bind.WatchOpts, sink chan<- *TraceabilityNewGoodsEvent) (event.Subscription, error) {

	logs, sub, err := _Traceability.contract.WatchLogs(opts, "newGoodsEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraceabilityNewGoodsEvent)
				if err := _Traceability.contract.UnpackLog(event, "newGoodsEvent", log); err != nil {
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

// ParseNewGoodsEvent is a log parse operation binding the contract event 0x30d18c7c4624d4b6e90e21ae02bab5166256618bc6876873dd5ba8c5ed3db86f.
//
// Solidity: event newGoodsEvent(string _addr, string _data, string goodsid)
func (_Traceability *TraceabilityFilterer) ParseNewGoodsEvent(log types.Log) (*TraceabilityNewGoodsEvent, error) {
	event := new(TraceabilityNewGoodsEvent)
	if err := _Traceability.contract.UnpackLog(event, "newGoodsEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
