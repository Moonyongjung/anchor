// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rpc

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

// AnchoringMetaData contains all meta data concerning the Anchoring contract.
var AnchoringMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"callBlockList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"merkleRoot\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"startBlock\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endBlock\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"merkleroot\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"startBlock\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endBlock\",\"type\":\"string\"}],\"name\":\"invokeBlockList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610690806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631b7167611461003b578063591916d214610057575b600080fd5b6100556004803603810190610050919061047f565b610077565b005b61005f6100c9565b60405161006e939291906105ae565b60405180910390f35b8260008001908051906020019061008f929190610282565b5081600060010190805190602001906100a9929190610282565b5080600060020190805190602001906100c3929190610282565b50505050565b60608060606000800180546100dd90610629565b80601f016020809104026020016040519081016040528092919081815260200182805461010990610629565b80156101565780601f1061012b57610100808354040283529160200191610156565b820191906000526020600020905b81548152906001019060200180831161013957829003601f168201915b505050505092506000600101805461016d90610629565b80601f016020809104026020016040519081016040528092919081815260200182805461019990610629565b80156101e65780601f106101bb576101008083540402835291602001916101e6565b820191906000526020600020905b8154815290600101906020018083116101c957829003601f168201915b50505050509150600060020180546101fd90610629565b80601f016020809104026020016040519081016040528092919081815260200182805461022990610629565b80156102765780601f1061024b57610100808354040283529160200191610276565b820191906000526020600020905b81548152906001019060200180831161025957829003601f168201915b50505050509050909192565b82805461028e90610629565b90600052602060002090601f0160209004810192826102b057600085556102f7565b82601f106102c957805160ff19168380011785556102f7565b828001600101855582156102f7579182015b828111156102f65782518255916020019190600101906102db565b5b5090506103049190610308565b5090565b5b80821115610321576000816000905550600101610309565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61038c82610343565b810181811067ffffffffffffffff821117156103ab576103aa610354565b5b80604052505050565b60006103be610325565b90506103ca8282610383565b919050565b600067ffffffffffffffff8211156103ea576103e9610354565b5b6103f382610343565b9050602081019050919050565b82818337600083830152505050565b600061042261041d846103cf565b6103b4565b90508281526020810184848401111561043e5761043d61033e565b5b610449848285610400565b509392505050565b600082601f83011261046657610465610339565b5b813561047684826020860161040f565b91505092915050565b6000806000606084860312156104985761049761032f565b5b600084013567ffffffffffffffff8111156104b6576104b5610334565b5b6104c286828701610451565b935050602084013567ffffffffffffffff8111156104e3576104e2610334565b5b6104ef86828701610451565b925050604084013567ffffffffffffffff8111156105105761050f610334565b5b61051c86828701610451565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b60005b83811015610560578082015181840152602081019050610545565b8381111561056f576000848401525b50505050565b600061058082610526565b61058a8185610531565b935061059a818560208601610542565b6105a381610343565b840191505092915050565b600060608201905081810360008301526105c88186610575565b905081810360208301526105dc8185610575565b905081810360408301526105f08184610575565b9050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061064157607f821691505b602082108103610654576106536105fa565b5b5091905056fea264697066735822122040d537a99b6ddfa76ecc39abbd6e3024b4545aeea02e983a1038a698e600afd464736f6c634300080e0033",
}

// AnchoringABI is the input ABI used to generate the binding from.
// Deprecated: Use AnchoringMetaData.ABI instead.
var AnchoringABI = AnchoringMetaData.ABI

// AnchoringBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AnchoringMetaData.Bin instead.
var AnchoringBin = AnchoringMetaData.Bin

// DeployAnchoring deploys a new Ethereum contract, binding an instance of Anchoring to it.
func DeployAnchoring(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Anchoring, error) {
	parsed, err := AnchoringMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AnchoringBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Anchoring{AnchoringCaller: AnchoringCaller{contract: contract}, AnchoringTransactor: AnchoringTransactor{contract: contract}, AnchoringFilterer: AnchoringFilterer{contract: contract}}, nil
}

// Anchoring is an auto generated Go binding around an Ethereum contract.
type Anchoring struct {
	AnchoringCaller     // Read-only binding to the contract
	AnchoringTransactor // Write-only binding to the contract
	AnchoringFilterer   // Log filterer for contract events
}

// AnchoringCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnchoringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchoringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnchoringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchoringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnchoringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchoringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnchoringSession struct {
	Contract     *Anchoring        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnchoringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnchoringCallerSession struct {
	Contract *AnchoringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AnchoringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnchoringTransactorSession struct {
	Contract     *AnchoringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AnchoringRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnchoringRaw struct {
	Contract *Anchoring // Generic contract binding to access the raw methods on
}

// AnchoringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnchoringCallerRaw struct {
	Contract *AnchoringCaller // Generic read-only contract binding to access the raw methods on
}

// AnchoringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnchoringTransactorRaw struct {
	Contract *AnchoringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnchoring creates a new instance of Anchoring, bound to a specific deployed contract.
func NewAnchoring(address common.Address, backend bind.ContractBackend) (*Anchoring, error) {
	contract, err := bindAnchoring(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Anchoring{AnchoringCaller: AnchoringCaller{contract: contract}, AnchoringTransactor: AnchoringTransactor{contract: contract}, AnchoringFilterer: AnchoringFilterer{contract: contract}}, nil
}

// NewAnchoringCaller creates a new read-only instance of Anchoring, bound to a specific deployed contract.
func NewAnchoringCaller(address common.Address, caller bind.ContractCaller) (*AnchoringCaller, error) {
	contract, err := bindAnchoring(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnchoringCaller{contract: contract}, nil
}

// NewAnchoringTransactor creates a new write-only instance of Anchoring, bound to a specific deployed contract.
func NewAnchoringTransactor(address common.Address, transactor bind.ContractTransactor) (*AnchoringTransactor, error) {
	contract, err := bindAnchoring(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnchoringTransactor{contract: contract}, nil
}

// NewAnchoringFilterer creates a new log filterer instance of Anchoring, bound to a specific deployed contract.
func NewAnchoringFilterer(address common.Address, filterer bind.ContractFilterer) (*AnchoringFilterer, error) {
	contract, err := bindAnchoring(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnchoringFilterer{contract: contract}, nil
}

// bindAnchoring binds a generic wrapper to an already deployed contract.
func bindAnchoring(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AnchoringABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Anchoring *AnchoringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Anchoring.Contract.AnchoringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Anchoring *AnchoringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anchoring.Contract.AnchoringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Anchoring *AnchoringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Anchoring.Contract.AnchoringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Anchoring *AnchoringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Anchoring.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Anchoring *AnchoringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anchoring.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Anchoring *AnchoringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Anchoring.Contract.contract.Transact(opts, method, params...)
}

// CallBlockList is a free data retrieval call binding the contract method 0x591916d2.
//
// Solidity: function callBlockList() view returns(string merkleRoot, string startBlock, string endBlock)
func (_Anchoring *AnchoringCaller) CallBlockList(opts *bind.CallOpts) (struct {
	MerkleRoot string
	StartBlock string
	EndBlock   string
}, error) {
	var out []interface{}
	err := _Anchoring.contract.Call(opts, &out, "callBlockList")

	outstruct := new(struct {
		MerkleRoot string
		StartBlock string
		EndBlock   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MerkleRoot = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.EndBlock = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// CallBlockList is a free data retrieval call binding the contract method 0x591916d2.
//
// Solidity: function callBlockList() view returns(string merkleRoot, string startBlock, string endBlock)
func (_Anchoring *AnchoringSession) CallBlockList() (struct {
	MerkleRoot string
	StartBlock string
	EndBlock   string
}, error) {
	return _Anchoring.Contract.CallBlockList(&_Anchoring.CallOpts)
}

// CallBlockList is a free data retrieval call binding the contract method 0x591916d2.
//
// Solidity: function callBlockList() view returns(string merkleRoot, string startBlock, string endBlock)
func (_Anchoring *AnchoringCallerSession) CallBlockList() (struct {
	MerkleRoot string
	StartBlock string
	EndBlock   string
}, error) {
	return _Anchoring.Contract.CallBlockList(&_Anchoring.CallOpts)
}

// InvokeBlockList is a paid mutator transaction binding the contract method 0x1b716761.
//
// Solidity: function invokeBlockList(string merkleroot, string startBlock, string endBlock) returns()
func (_Anchoring *AnchoringTransactor) InvokeBlockList(opts *bind.TransactOpts, merkleroot string, startBlock string, endBlock string) (*types.Transaction, error) {
	return _Anchoring.contract.Transact(opts, "invokeBlockList", merkleroot, startBlock, endBlock)
}

// InvokeBlockList is a paid mutator transaction binding the contract method 0x1b716761.
//
// Solidity: function invokeBlockList(string merkleroot, string startBlock, string endBlock) returns()
func (_Anchoring *AnchoringSession) InvokeBlockList(merkleroot string, startBlock string, endBlock string) (*types.Transaction, error) {
	return _Anchoring.Contract.InvokeBlockList(&_Anchoring.TransactOpts, merkleroot, startBlock, endBlock)
}

// InvokeBlockList is a paid mutator transaction binding the contract method 0x1b716761.
//
// Solidity: function invokeBlockList(string merkleroot, string startBlock, string endBlock) returns()
func (_Anchoring *AnchoringTransactorSession) InvokeBlockList(merkleroot string, startBlock string, endBlock string) (*types.Transaction, error) {
	return _Anchoring.Contract.InvokeBlockList(&_Anchoring.TransactOpts, merkleroot, startBlock, endBlock)
}
