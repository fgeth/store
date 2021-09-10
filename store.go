// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Store

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getItem\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x608060405234801561001057600080fd5b50604080518082019091526005808252640312e302e360dc1b602090920191825261003d91600091610043565b50610117565b82805461004f906100dc565b90600052602060002090601f01602090048101928261007157600085556100b7565b82601f1061008a57805160ff19168380011785556100b7565b828001600101855582156100b7579182015b828111156100b757825182559160200191906001019061009c565b506100c39291506100c7565b5090565b5b808211156100c357600081556001016100c8565b600181811c908216806100f057607f821691505b6020821081141561011157634e487b7160e01b600052602260045260246000fd5b50919050565b6105b2806101266000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630d8e6e2c1461005c5780634909b29b1461008f57806354fd4d50146100a25780637a4c982e146100aa578063b6010fcd146100bf575b600080fd5b6040805180820190915260058152640312e302e360dc1b60208201525b60405161008691906104b3565b60405180910390f35b61007961009d3660046103ca565b6100d2565b610079610177565b6100bd6100b8366004610407565b610184565b005b6100796100cd3660046103ca565b6101f4565b8051602081830181018051600182529282019190930120915280546100f69061052b565b80601f01602080910402602001604051908101604052809291908181526020018280546101229061052b565b801561016f5780601f106101445761010080835404028352916020019161016f565b820191906000526020600020905b81548152906001019060200180831161015257829003601f168201915b505050505081565b600080546100f69061052b565b806001836040516101959190610497565b908152602001604051809103902090805190602001906101b69291906102a4565b507f523139e7aa4a7267d8d1b859cba73a73993683d42bc00a7a06770e7b6ca57ef382826040516101e89291906104cd565b60405180910390a15050565b60606001826040516102069190610497565b9081526020016040518091039020805461021f9061052b565b80601f016020809104026020016040519081016040528092919081815260200182805461024b9061052b565b80156102985780601f1061026d57610100808354040283529160200191610298565b820191906000526020600020905b81548152906001019060200180831161027b57829003601f168201915b50505050509050919050565b8280546102b09061052b565b90600052602060002090601f0160209004810192826102d25760008555610318565b82601f106102eb57805160ff1916838001178555610318565b82800160010185558215610318579182015b828111156103185782518255916020019190600101906102fd565b50610324929150610328565b5090565b5b808211156103245760008155600101610329565b600082601f83011261034e57600080fd5b813567ffffffffffffffff8082111561036957610369610566565b604051601f8301601f19908116603f0116810190828211818310171561039157610391610566565b816040528381528660208588010111156103aa57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000602082840312156103dc57600080fd5b813567ffffffffffffffff8111156103f357600080fd5b6103ff8482850161033d565b949350505050565b6000806040838503121561041a57600080fd5b823567ffffffffffffffff8082111561043257600080fd5b61043e8683870161033d565b9350602085013591508082111561045457600080fd5b506104618582860161033d565b9150509250929050565b600081518084526104838160208601602086016104fb565b601f01601f19169290920160200192915050565b600082516104a98184602087016104fb565b9190910192915050565b6020815260006104c6602083018461046b565b9392505050565b6040815260006104e0604083018561046b565b82810360208401526104f2818561046b565b95945050505050565b60005b838110156105165781810151838201526020016104fe565b83811115610525576000848401525b50505050565b600181811c9082168061053f57607f821691505b6020821081141561056057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220742a6f8db360cecbe4a7a6bcb7c9be133b4848dcc17d5d0464a75d051659918764736f6c63430008060033"

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// GetItem is a free data retrieval call binding the contract method 0xb6010fcd.
//
// Solidity: function getItem(string key) view returns(string)
func (_Store *StoreCaller) GetItem(opts *bind.CallOpts, key string) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getItem", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetItem is a free data retrieval call binding the contract method 0xb6010fcd.
//
// Solidity: function getItem(string key) view returns(string)
func (_Store *StoreSession) GetItem(key string) (string, error) {
	return _Store.Contract.GetItem(&_Store.CallOpts, key)
}

// GetItem is a free data retrieval call binding the contract method 0xb6010fcd.
//
// Solidity: function getItem(string key) view returns(string)
func (_Store *StoreCallerSession) GetItem(key string) (string, error) {
	return _Store.Contract.GetItem(&_Store.CallOpts, key)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Store *StoreCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Store *StoreSession) GetVersion() (string, error) {
	return _Store.Contract.GetVersion(&_Store.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_Store *StoreCallerSession) GetVersion() (string, error) {
	return _Store.Contract.GetVersion(&_Store.CallOpts)
}

// Items is a free data retrieval call binding the contract method 0x4909b29b.
//
// Solidity: function items(string ) view returns(string)
func (_Store *StoreCaller) Items(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "items", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x4909b29b.
//
// Solidity: function items(string ) view returns(string)
func (_Store *StoreSession) Items(arg0 string) (string, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x4909b29b.
//
// Solidity: function items(string ) view returns(string)
func (_Store *StoreCallerSession) Items(arg0 string) (string, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreCallerSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0x7a4c982e.
//
// Solidity: function setItem(string key, string value) returns()
func (_Store *StoreTransactor) SetItem(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0x7a4c982e.
//
// Solidity: function setItem(string key, string value) returns()
func (_Store *StoreSession) SetItem(key string, value string) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0x7a4c982e.
//
// Solidity: function setItem(string key, string value) returns()
func (_Store *StoreTransactorSession) SetItem(key string, value string) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}

// StoreItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Store contract.
type StoreItemSetIterator struct {
	Event *StoreItemSet // Event containing the contract specifics and raw log

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
func (it *StoreItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreItemSet)
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
		it.Event = new(StoreItemSet)
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
func (it *StoreItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreItemSet represents a ItemSet event raised by the Store contract.
type StoreItemSet struct {
	Key   string
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0x523139e7aa4a7267d8d1b859cba73a73993683d42bc00a7a06770e7b6ca57ef3.
//
// Solidity: event ItemSet(string key, string value)
func (_Store *StoreFilterer) FilterItemSet(opts *bind.FilterOpts) (*StoreItemSetIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &StoreItemSetIterator{contract: _Store.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0x523139e7aa4a7267d8d1b859cba73a73993683d42bc00a7a06770e7b6ca57ef3.
//
// Solidity: event ItemSet(string key, string value)
func (_Store *StoreFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *StoreItemSet) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreItemSet)
				if err := _Store.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

// ParseItemSet is a log parse operation binding the contract event 0x523139e7aa4a7267d8d1b859cba73a73993683d42bc00a7a06770e7b6ca57ef3.
//
// Solidity: event ItemSet(string key, string value)
func (_Store *StoreFilterer) ParseItemSet(log types.Log) (*StoreItemSet, error) {
	event := new(StoreItemSet)
	if err := _Store.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
