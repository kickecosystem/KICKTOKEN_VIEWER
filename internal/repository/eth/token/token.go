// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minteradmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseTrigger\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"meltTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnallow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"destroyFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimMinterAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimMelterAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromMelters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"melteradmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"meltBatchTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"transferPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMelter\",\"type\":\"address\"}],\"name\":\"transferMelterAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mintBatchToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mintBatchFrozenTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"whitelistadmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFromMinters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingMelterAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBurner\",\"type\":\"address\"}],\"name\":\"transferWhitelistAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingPauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToMinters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"freezeTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintFrozenTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingMinterAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addToMelters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimBurner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"transferMinterAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWhiteListAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrozenToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"burnTrigger\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingBurner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Melt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FrozenTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousMelter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMelter\",\"type\":\"address\"}],\"name\":\"MelterTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousMinter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"MinterTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousPauser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"PauserTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousBurner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBurner\",\"type\":\"address\"}],\"name\":\"BurnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"WhitelistAdminTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address spender) constant returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address spender) constant returns(uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address spender) constant returns(uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// Burnallow is a free data retrieval call binding the contract method 0x0dc6f69c.
//
// Solidity: function burnallow() constant returns(bool)
func (_Token *TokenCaller) Burnallow(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "burnallow")
	return *ret0, err
}

// Burnallow is a free data retrieval call binding the contract method 0x0dc6f69c.
//
// Solidity: function burnallow() constant returns(bool)
func (_Token *TokenSession) Burnallow() (bool, error) {
	return _Token.Contract.Burnallow(&_Token.CallOpts)
}

// Burnallow is a free data retrieval call binding the contract method 0x0dc6f69c.
//
// Solidity: function burnallow() constant returns(bool)
func (_Token *TokenCallerSession) Burnallow() (bool, error) {
	return _Token.Contract.Burnallow(&_Token.CallOpts)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() constant returns(address)
func (_Token *TokenCaller) Burner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "burner")
	return *ret0, err
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() constant returns(address)
func (_Token *TokenSession) Burner() (common.Address, error) {
	return _Token.Contract.Burner(&_Token.CallOpts)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() constant returns(address)
func (_Token *TokenCallerSession) Burner() (common.Address, error) {
	return _Token.Contract.Burner(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) constant returns(bool)
func (_Token *TokenCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isWhitelisted", account)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) constant returns(bool)
func (_Token *TokenSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Token.Contract.IsWhitelisted(&_Token.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) constant returns(bool)
func (_Token *TokenCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Token.Contract.IsWhitelisted(&_Token.CallOpts, account)
}

// Melteradmin is a free data retrieval call binding the contract method 0x35cde1ce.
//
// Solidity: function melteradmin() constant returns(address)
func (_Token *TokenCaller) Melteradmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "melteradmin")
	return *ret0, err
}

// Melteradmin is a free data retrieval call binding the contract method 0x35cde1ce.
//
// Solidity: function melteradmin() constant returns(address)
func (_Token *TokenSession) Melteradmin() (common.Address, error) {
	return _Token.Contract.Melteradmin(&_Token.CallOpts)
}

// Melteradmin is a free data retrieval call binding the contract method 0x35cde1ce.
//
// Solidity: function melteradmin() constant returns(address)
func (_Token *TokenCallerSession) Melteradmin() (common.Address, error) {
	return _Token.Contract.Melteradmin(&_Token.CallOpts)
}

// Minteradmin is a free data retrieval call binding the contract method 0x08c985bf.
//
// Solidity: function minteradmin() constant returns(address)
func (_Token *TokenCaller) Minteradmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "minteradmin")
	return *ret0, err
}

// Minteradmin is a free data retrieval call binding the contract method 0x08c985bf.
//
// Solidity: function minteradmin() constant returns(address)
func (_Token *TokenSession) Minteradmin() (common.Address, error) {
	return _Token.Contract.Minteradmin(&_Token.CallOpts)
}

// Minteradmin is a free data retrieval call binding the contract method 0x08c985bf.
//
// Solidity: function minteradmin() constant returns(address)
func (_Token *TokenCallerSession) Minteradmin() (common.Address, error) {
	return _Token.Contract.Minteradmin(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Token *TokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Token *TokenSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Token *TokenCallerSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() constant returns(address)
func (_Token *TokenCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "pauser")
	return *ret0, err
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() constant returns(address)
func (_Token *TokenSession) Pauser() (common.Address, error) {
	return _Token.Contract.Pauser(&_Token.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() constant returns(address)
func (_Token *TokenCallerSession) Pauser() (common.Address, error) {
	return _Token.Contract.Pauser(&_Token.CallOpts)
}

// PendingBurner is a free data retrieval call binding the contract method 0xf2fbce99.
//
// Solidity: function pendingBurner() constant returns(address)
func (_Token *TokenCaller) PendingBurner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "pendingBurner")
	return *ret0, err
}

// PendingBurner is a free data retrieval call binding the contract method 0xf2fbce99.
//
// Solidity: function pendingBurner() constant returns(address)
func (_Token *TokenSession) PendingBurner() (common.Address, error) {
	return _Token.Contract.PendingBurner(&_Token.CallOpts)
}

// PendingBurner is a free data retrieval call binding the contract method 0xf2fbce99.
//
// Solidity: function pendingBurner() constant returns(address)
func (_Token *TokenCallerSession) PendingBurner() (common.Address, error) {
	return _Token.Contract.PendingBurner(&_Token.CallOpts)
}

// PendingMelterAdmin is a free data retrieval call binding the contract method 0x70766280.
//
// Solidity: function pendingMelterAdmin() constant returns(address)
func (_Token *TokenCaller) PendingMelterAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "pendingMelterAdmin")
	return *ret0, err
}

// PendingMelterAdmin is a free data retrieval call binding the contract method 0x70766280.
//
// Solidity: function pendingMelterAdmin() constant returns(address)
func (_Token *TokenSession) PendingMelterAdmin() (common.Address, error) {
	return _Token.Contract.PendingMelterAdmin(&_Token.CallOpts)
}

// PendingMelterAdmin is a free data retrieval call binding the contract method 0x70766280.
//
// Solidity: function pendingMelterAdmin() constant returns(address)
func (_Token *TokenCallerSession) PendingMelterAdmin() (common.Address, error) {
	return _Token.Contract.PendingMelterAdmin(&_Token.CallOpts)
}

// PendingMinterAdmin is a free data retrieval call binding the contract method 0xce1d82f5.
//
// Solidity: function pendingMinterAdmin() constant returns(address)
func (_Token *TokenCaller) PendingMinterAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "pendingMinterAdmin")
	return *ret0, err
}

// PendingMinterAdmin is a free data retrieval call binding the contract method 0xce1d82f5.
//
// Solidity: function pendingMinterAdmin() constant returns(address)
func (_Token *TokenSession) PendingMinterAdmin() (common.Address, error) {
	return _Token.Contract.PendingMinterAdmin(&_Token.CallOpts)
}

// PendingMinterAdmin is a free data retrieval call binding the contract method 0xce1d82f5.
//
// Solidity: function pendingMinterAdmin() constant returns(address)
func (_Token *TokenCallerSession) PendingMinterAdmin() (common.Address, error) {
	return _Token.Contract.PendingMinterAdmin(&_Token.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_Token *TokenCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_Token *TokenSession) PendingOwner() (common.Address, error) {
	return _Token.Contract.PendingOwner(&_Token.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_Token *TokenCallerSession) PendingOwner() (common.Address, error) {
	return _Token.Contract.PendingOwner(&_Token.CallOpts)
}

// PendingPauser is a free data retrieval call binding the contract method 0x9a7165e4.
//
// Solidity: function pendingPauser() constant returns(address)
func (_Token *TokenCaller) PendingPauser(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "pendingPauser")
	return *ret0, err
}

// PendingPauser is a free data retrieval call binding the contract method 0x9a7165e4.
//
// Solidity: function pendingPauser() constant returns(address)
func (_Token *TokenSession) PendingPauser() (common.Address, error) {
	return _Token.Contract.PendingPauser(&_Token.CallOpts)
}

// PendingPauser is a free data retrieval call binding the contract method 0x9a7165e4.
//
// Solidity: function pendingPauser() constant returns(address)
func (_Token *TokenCallerSession) PendingPauser() (common.Address, error) {
	return _Token.Contract.PendingPauser(&_Token.CallOpts)
}

// PendingWhiteListAdmin is a free data retrieval call binding the contract method 0xe020455c.
//
// Solidity: function pendingWhiteListAdmin() constant returns(address)
func (_Token *TokenCaller) PendingWhiteListAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "pendingWhiteListAdmin")
	return *ret0, err
}

// PendingWhiteListAdmin is a free data retrieval call binding the contract method 0xe020455c.
//
// Solidity: function pendingWhiteListAdmin() constant returns(address)
func (_Token *TokenSession) PendingWhiteListAdmin() (common.Address, error) {
	return _Token.Contract.PendingWhiteListAdmin(&_Token.CallOpts)
}

// PendingWhiteListAdmin is a free data retrieval call binding the contract method 0xe020455c.
//
// Solidity: function pendingWhiteListAdmin() constant returns(address)
func (_Token *TokenCallerSession) PendingWhiteListAdmin() (common.Address, error) {
	return _Token.Contract.PendingWhiteListAdmin(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Whitelistadmin is a free data retrieval call binding the contract method 0x55e6ef34.
//
// Solidity: function whitelistadmin() constant returns(address)
func (_Token *TokenCaller) Whitelistadmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "whitelistadmin")
	return *ret0, err
}

// Whitelistadmin is a free data retrieval call binding the contract method 0x55e6ef34.
//
// Solidity: function whitelistadmin() constant returns(address)
func (_Token *TokenSession) Whitelistadmin() (common.Address, error) {
	return _Token.Contract.Whitelistadmin(&_Token.CallOpts)
}

// Whitelistadmin is a free data retrieval call binding the contract method 0x55e6ef34.
//
// Solidity: function whitelistadmin() constant returns(address)
func (_Token *TokenCallerSession) Whitelistadmin() (common.Address, error) {
	return _Token.Contract.Whitelistadmin(&_Token.CallOpts)
}

// AddToMelters is a paid mutator transaction binding the contract method 0xd8da5ce9.
//
// Solidity: function addToMelters(address account) returns()
func (_Token *TokenTransactor) AddToMelters(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addToMelters", account)
}

// AddToMelters is a paid mutator transaction binding the contract method 0xd8da5ce9.
//
// Solidity: function addToMelters(address account) returns()
func (_Token *TokenSession) AddToMelters(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddToMelters(&_Token.TransactOpts, account)
}

// AddToMelters is a paid mutator transaction binding the contract method 0xd8da5ce9.
//
// Solidity: function addToMelters(address account) returns()
func (_Token *TokenTransactorSession) AddToMelters(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddToMelters(&_Token.TransactOpts, account)
}

// AddToMinters is a paid mutator transaction binding the contract method 0xa15c15af.
//
// Solidity: function addToMinters(address account) returns()
func (_Token *TokenTransactor) AddToMinters(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addToMinters", account)
}

// AddToMinters is a paid mutator transaction binding the contract method 0xa15c15af.
//
// Solidity: function addToMinters(address account) returns()
func (_Token *TokenSession) AddToMinters(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddToMinters(&_Token.TransactOpts, account)
}

// AddToMinters is a paid mutator transaction binding the contract method 0xa15c15af.
//
// Solidity: function addToMinters(address account) returns()
func (_Token *TokenTransactorSession) AddToMinters(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddToMinters(&_Token.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Token *TokenTransactor) AddWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addWhitelisted", account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Token *TokenSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddWhitelisted(&_Token.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Token *TokenTransactorSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddWhitelisted(&_Token.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Token *TokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Token *TokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns(bool)
func (_Token *TokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns(bool)
func (_Token *TokenSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.Contract.ApproveAndCall(&_Token.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns(bool)
func (_Token *TokenTransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Token.Contract.ApproveAndCall(&_Token.TransactOpts, _spender, _value, _extraData)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Token *TokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Token *TokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Token *TokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Token *TokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Token *TokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.BurnFrom(&_Token.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Token *TokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.BurnFrom(&_Token.TransactOpts, account, amount)
}

// BurnTrigger is a paid mutator transaction binding the contract method 0xeeada434.
//
// Solidity: function burnTrigger() returns()
func (_Token *TokenTransactor) BurnTrigger(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burnTrigger")
}

// BurnTrigger is a paid mutator transaction binding the contract method 0xeeada434.
//
// Solidity: function burnTrigger() returns()
func (_Token *TokenSession) BurnTrigger() (*types.Transaction, error) {
	return _Token.Contract.BurnTrigger(&_Token.TransactOpts)
}

// BurnTrigger is a paid mutator transaction binding the contract method 0xeeada434.
//
// Solidity: function burnTrigger() returns()
func (_Token *TokenTransactorSession) BurnTrigger() (*types.Transaction, error) {
	return _Token.Contract.BurnTrigger(&_Token.TransactOpts)
}

// ClaimBurner is a paid mutator transaction binding the contract method 0xd9489aa2.
//
// Solidity: function claimBurner() returns()
func (_Token *TokenTransactor) ClaimBurner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimBurner")
}

// ClaimBurner is a paid mutator transaction binding the contract method 0xd9489aa2.
//
// Solidity: function claimBurner() returns()
func (_Token *TokenSession) ClaimBurner() (*types.Transaction, error) {
	return _Token.Contract.ClaimBurner(&_Token.TransactOpts)
}

// ClaimBurner is a paid mutator transaction binding the contract method 0xd9489aa2.
//
// Solidity: function claimBurner() returns()
func (_Token *TokenTransactorSession) ClaimBurner() (*types.Transaction, error) {
	return _Token.Contract.ClaimBurner(&_Token.TransactOpts)
}

// ClaimMelterAdmin is a paid mutator transaction binding the contract method 0x1e57a2b2.
//
// Solidity: function claimMelterAdmin() returns()
func (_Token *TokenTransactor) ClaimMelterAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimMelterAdmin")
}

// ClaimMelterAdmin is a paid mutator transaction binding the contract method 0x1e57a2b2.
//
// Solidity: function claimMelterAdmin() returns()
func (_Token *TokenSession) ClaimMelterAdmin() (*types.Transaction, error) {
	return _Token.Contract.ClaimMelterAdmin(&_Token.TransactOpts)
}

// ClaimMelterAdmin is a paid mutator transaction binding the contract method 0x1e57a2b2.
//
// Solidity: function claimMelterAdmin() returns()
func (_Token *TokenTransactorSession) ClaimMelterAdmin() (*types.Transaction, error) {
	return _Token.Contract.ClaimMelterAdmin(&_Token.TransactOpts)
}

// ClaimMinterAdmin is a paid mutator transaction binding the contract method 0x1d2c0b3d.
//
// Solidity: function claimMinterAdmin() returns()
func (_Token *TokenTransactor) ClaimMinterAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimMinterAdmin")
}

// ClaimMinterAdmin is a paid mutator transaction binding the contract method 0x1d2c0b3d.
//
// Solidity: function claimMinterAdmin() returns()
func (_Token *TokenSession) ClaimMinterAdmin() (*types.Transaction, error) {
	return _Token.Contract.ClaimMinterAdmin(&_Token.TransactOpts)
}

// ClaimMinterAdmin is a paid mutator transaction binding the contract method 0x1d2c0b3d.
//
// Solidity: function claimMinterAdmin() returns()
func (_Token *TokenTransactorSession) ClaimMinterAdmin() (*types.Transaction, error) {
	return _Token.Contract.ClaimMinterAdmin(&_Token.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Token *TokenTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Token *TokenSession) ClaimOwnership() (*types.Transaction, error) {
	return _Token.Contract.ClaimOwnership(&_Token.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Token *TokenTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Token.Contract.ClaimOwnership(&_Token.TransactOpts)
}

// ClaimPauser is a paid mutator transaction binding the contract method 0x13d71f6e.
//
// Solidity: function claimPauser() returns()
func (_Token *TokenTransactor) ClaimPauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimPauser")
}

// ClaimPauser is a paid mutator transaction binding the contract method 0x13d71f6e.
//
// Solidity: function claimPauser() returns()
func (_Token *TokenSession) ClaimPauser() (*types.Transaction, error) {
	return _Token.Contract.ClaimPauser(&_Token.TransactOpts)
}

// ClaimPauser is a paid mutator transaction binding the contract method 0x13d71f6e.
//
// Solidity: function claimPauser() returns()
func (_Token *TokenTransactorSession) ClaimPauser() (*types.Transaction, error) {
	return _Token.Contract.ClaimPauser(&_Token.TransactOpts)
}

// ClaimWhitelistAdmin is a paid mutator transaction binding the contract method 0x434450e8.
//
// Solidity: function claimWhitelistAdmin() returns()
func (_Token *TokenTransactor) ClaimWhitelistAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimWhitelistAdmin")
}

// ClaimWhitelistAdmin is a paid mutator transaction binding the contract method 0x434450e8.
//
// Solidity: function claimWhitelistAdmin() returns()
func (_Token *TokenSession) ClaimWhitelistAdmin() (*types.Transaction, error) {
	return _Token.Contract.ClaimWhitelistAdmin(&_Token.TransactOpts)
}

// ClaimWhitelistAdmin is a paid mutator transaction binding the contract method 0x434450e8.
//
// Solidity: function claimWhitelistAdmin() returns()
func (_Token *TokenTransactorSession) ClaimWhitelistAdmin() (*types.Transaction, error) {
	return _Token.Contract.ClaimWhitelistAdmin(&_Token.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, spender, subtractedValue)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(address account, uint256 amount) returns()
func (_Token *TokenTransactor) Destroy(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "destroy", account, amount)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(address account, uint256 amount) returns()
func (_Token *TokenSession) Destroy(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Destroy(&_Token.TransactOpts, account, amount)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(address account, uint256 amount) returns()
func (_Token *TokenTransactorSession) Destroy(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Destroy(&_Token.TransactOpts, account, amount)
}

// DestroyFrozen is a paid mutator transaction binding the contract method 0x171177b3.
//
// Solidity: function destroyFrozen(address account, uint256 amount) returns()
func (_Token *TokenTransactor) DestroyFrozen(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "destroyFrozen", account, amount)
}

// DestroyFrozen is a paid mutator transaction binding the contract method 0x171177b3.
//
// Solidity: function destroyFrozen(address account, uint256 amount) returns()
func (_Token *TokenSession) DestroyFrozen(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DestroyFrozen(&_Token.TransactOpts, account, amount)
}

// DestroyFrozen is a paid mutator transaction binding the contract method 0x171177b3.
//
// Solidity: function destroyFrozen(address account, uint256 amount) returns()
func (_Token *TokenTransactorSession) DestroyFrozen(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DestroyFrozen(&_Token.TransactOpts, account, amount)
}

// FreezeTokens is a paid mutator transaction binding the contract method 0xa4df6c6a.
//
// Solidity: function freezeTokens(address account, uint256 amount) returns(bool)
func (_Token *TokenTransactor) FreezeTokens(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "freezeTokens", account, amount)
}

// FreezeTokens is a paid mutator transaction binding the contract method 0xa4df6c6a.
//
// Solidity: function freezeTokens(address account, uint256 amount) returns(bool)
func (_Token *TokenSession) FreezeTokens(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.FreezeTokens(&_Token.TransactOpts, account, amount)
}

// FreezeTokens is a paid mutator transaction binding the contract method 0xa4df6c6a.
//
// Solidity: function freezeTokens(address account, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) FreezeTokens(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.FreezeTokens(&_Token.TransactOpts, account, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, spender, addedValue)
}

// MeltBatchTokens is a paid mutator transaction binding the contract method 0x3d7cd488.
//
// Solidity: function meltBatchTokens(address[] accounts, uint256[] amounts) returns(bool)
func (_Token *TokenTransactor) MeltBatchTokens(opts *bind.TransactOpts, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "meltBatchTokens", accounts, amounts)
}

// MeltBatchTokens is a paid mutator transaction binding the contract method 0x3d7cd488.
//
// Solidity: function meltBatchTokens(address[] accounts, uint256[] amounts) returns(bool)
func (_Token *TokenSession) MeltBatchTokens(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MeltBatchTokens(&_Token.TransactOpts, accounts, amounts)
}

// MeltBatchTokens is a paid mutator transaction binding the contract method 0x3d7cd488.
//
// Solidity: function meltBatchTokens(address[] accounts, uint256[] amounts) returns(bool)
func (_Token *TokenTransactorSession) MeltBatchTokens(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MeltBatchTokens(&_Token.TransactOpts, accounts, amounts)
}

// MeltTokens is a paid mutator transaction binding the contract method 0x0c15f601.
//
// Solidity: function meltTokens(address account, uint256 amount) returns(bool)
func (_Token *TokenTransactor) MeltTokens(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "meltTokens", account, amount)
}

// MeltTokens is a paid mutator transaction binding the contract method 0x0c15f601.
//
// Solidity: function meltTokens(address account, uint256 amount) returns(bool)
func (_Token *TokenSession) MeltTokens(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MeltTokens(&_Token.TransactOpts, account, amount)
}

// MeltTokens is a paid mutator transaction binding the contract method 0x0c15f601.
//
// Solidity: function meltTokens(address account, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) MeltTokens(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MeltTokens(&_Token.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_Token *TokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, account, amount)
}

// MintBatchFrozenTokens is a paid mutator transaction binding the contract method 0x54b2b176.
//
// Solidity: function mintBatchFrozenTokens(address[] accounts, uint256[] amounts) returns(bool)
func (_Token *TokenTransactor) MintBatchFrozenTokens(opts *bind.TransactOpts, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintBatchFrozenTokens", accounts, amounts)
}

// MintBatchFrozenTokens is a paid mutator transaction binding the contract method 0x54b2b176.
//
// Solidity: function mintBatchFrozenTokens(address[] accounts, uint256[] amounts) returns(bool)
func (_Token *TokenSession) MintBatchFrozenTokens(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintBatchFrozenTokens(&_Token.TransactOpts, accounts, amounts)
}

// MintBatchFrozenTokens is a paid mutator transaction binding the contract method 0x54b2b176.
//
// Solidity: function mintBatchFrozenTokens(address[] accounts, uint256[] amounts) returns(bool)
func (_Token *TokenTransactorSession) MintBatchFrozenTokens(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintBatchFrozenTokens(&_Token.TransactOpts, accounts, amounts)
}

// MintBatchToken is a paid mutator transaction binding the contract method 0x4c808904.
//
// Solidity: function mintBatchToken(address[] accounts, uint256[] amounts) returns(bool)
func (_Token *TokenTransactor) MintBatchToken(opts *bind.TransactOpts, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintBatchToken", accounts, amounts)
}

// MintBatchToken is a paid mutator transaction binding the contract method 0x4c808904.
//
// Solidity: function mintBatchToken(address[] accounts, uint256[] amounts) returns(bool)
func (_Token *TokenSession) MintBatchToken(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintBatchToken(&_Token.TransactOpts, accounts, amounts)
}

// MintBatchToken is a paid mutator transaction binding the contract method 0x4c808904.
//
// Solidity: function mintBatchToken(address[] accounts, uint256[] amounts) returns(bool)
func (_Token *TokenTransactorSession) MintBatchToken(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintBatchToken(&_Token.TransactOpts, accounts, amounts)
}

// MintFrozenTokens is a paid mutator transaction binding the contract method 0xc4157b29.
//
// Solidity: function mintFrozenTokens(address account, uint256 amount) returns(bool)
func (_Token *TokenTransactor) MintFrozenTokens(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mintFrozenTokens", account, amount)
}

// MintFrozenTokens is a paid mutator transaction binding the contract method 0xc4157b29.
//
// Solidity: function mintFrozenTokens(address account, uint256 amount) returns(bool)
func (_Token *TokenSession) MintFrozenTokens(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintFrozenTokens(&_Token.TransactOpts, account, amount)
}

// MintFrozenTokens is a paid mutator transaction binding the contract method 0xc4157b29.
//
// Solidity: function mintFrozenTokens(address account, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) MintFrozenTokens(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MintFrozenTokens(&_Token.TransactOpts, account, amount)
}

// PauseTrigger is a paid mutator transaction binding the contract method 0x096a2e32.
//
// Solidity: function pauseTrigger() returns()
func (_Token *TokenTransactor) PauseTrigger(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pauseTrigger")
}

// PauseTrigger is a paid mutator transaction binding the contract method 0x096a2e32.
//
// Solidity: function pauseTrigger() returns()
func (_Token *TokenSession) PauseTrigger() (*types.Transaction, error) {
	return _Token.Contract.PauseTrigger(&_Token.TransactOpts)
}

// PauseTrigger is a paid mutator transaction binding the contract method 0x096a2e32.
//
// Solidity: function pauseTrigger() returns()
func (_Token *TokenTransactorSession) PauseTrigger() (*types.Transaction, error) {
	return _Token.Contract.PauseTrigger(&_Token.TransactOpts)
}

// RemoveFromMelters is a paid mutator transaction binding the contract method 0x238fd731.
//
// Solidity: function removeFromMelters(address account) returns()
func (_Token *TokenTransactor) RemoveFromMelters(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeFromMelters", account)
}

// RemoveFromMelters is a paid mutator transaction binding the contract method 0x238fd731.
//
// Solidity: function removeFromMelters(address account) returns()
func (_Token *TokenSession) RemoveFromMelters(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveFromMelters(&_Token.TransactOpts, account)
}

// RemoveFromMelters is a paid mutator transaction binding the contract method 0x238fd731.
//
// Solidity: function removeFromMelters(address account) returns()
func (_Token *TokenTransactorSession) RemoveFromMelters(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveFromMelters(&_Token.TransactOpts, account)
}

// RemoveFromMinters is a paid mutator transaction binding the contract method 0x6e722920.
//
// Solidity: function removeFromMinters(address account) returns()
func (_Token *TokenTransactor) RemoveFromMinters(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeFromMinters", account)
}

// RemoveFromMinters is a paid mutator transaction binding the contract method 0x6e722920.
//
// Solidity: function removeFromMinters(address account) returns()
func (_Token *TokenSession) RemoveFromMinters(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveFromMinters(&_Token.TransactOpts, account)
}

// RemoveFromMinters is a paid mutator transaction binding the contract method 0x6e722920.
//
// Solidity: function removeFromMinters(address account) returns()
func (_Token *TokenTransactorSession) RemoveFromMinters(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveFromMinters(&_Token.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Token *TokenTransactor) RemoveWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeWhitelisted", account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Token *TokenSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveWhitelisted(&_Token.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Token *TokenTransactorSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveWhitelisted(&_Token.TransactOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, sender, recipient, amount)
}

// TransferFrozenToken is a paid mutator transaction binding the contract method 0xe50d3180.
//
// Solidity: function transferFrozenToken(address from, address to, uint256 amount) returns(bool)
func (_Token *TokenTransactor) TransferFrozenToken(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrozenToken", from, to, amount)
}

// TransferFrozenToken is a paid mutator transaction binding the contract method 0xe50d3180.
//
// Solidity: function transferFrozenToken(address from, address to, uint256 amount) returns(bool)
func (_Token *TokenSession) TransferFrozenToken(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrozenToken(&_Token.TransactOpts, from, to, amount)
}

// TransferFrozenToken is a paid mutator transaction binding the contract method 0xe50d3180.
//
// Solidity: function transferFrozenToken(address from, address to, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) TransferFrozenToken(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrozenToken(&_Token.TransactOpts, from, to, amount)
}

// TransferMelterAdmin is a paid mutator transaction binding the contract method 0x483a8df9.
//
// Solidity: function transferMelterAdmin(address newMelter) returns()
func (_Token *TokenTransactor) TransferMelterAdmin(opts *bind.TransactOpts, newMelter common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferMelterAdmin", newMelter)
}

// TransferMelterAdmin is a paid mutator transaction binding the contract method 0x483a8df9.
//
// Solidity: function transferMelterAdmin(address newMelter) returns()
func (_Token *TokenSession) TransferMelterAdmin(newMelter common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferMelterAdmin(&_Token.TransactOpts, newMelter)
}

// TransferMelterAdmin is a paid mutator transaction binding the contract method 0x483a8df9.
//
// Solidity: function transferMelterAdmin(address newMelter) returns()
func (_Token *TokenTransactorSession) TransferMelterAdmin(newMelter common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferMelterAdmin(&_Token.TransactOpts, newMelter)
}

// TransferMinterAdmin is a paid mutator transaction binding the contract method 0xdba03d81.
//
// Solidity: function transferMinterAdmin(address newMinter) returns()
func (_Token *TokenTransactor) TransferMinterAdmin(opts *bind.TransactOpts, newMinter common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferMinterAdmin", newMinter)
}

// TransferMinterAdmin is a paid mutator transaction binding the contract method 0xdba03d81.
//
// Solidity: function transferMinterAdmin(address newMinter) returns()
func (_Token *TokenSession) TransferMinterAdmin(newMinter common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferMinterAdmin(&_Token.TransactOpts, newMinter)
}

// TransferMinterAdmin is a paid mutator transaction binding the contract method 0xdba03d81.
//
// Solidity: function transferMinterAdmin(address newMinter) returns()
func (_Token *TokenTransactorSession) TransferMinterAdmin(newMinter common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferMinterAdmin(&_Token.TransactOpts, newMinter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferPauser is a paid mutator transaction binding the contract method 0x4421ea21.
//
// Solidity: function transferPauser(address newPauser) returns()
func (_Token *TokenTransactor) TransferPauser(opts *bind.TransactOpts, newPauser common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferPauser", newPauser)
}

// TransferPauser is a paid mutator transaction binding the contract method 0x4421ea21.
//
// Solidity: function transferPauser(address newPauser) returns()
func (_Token *TokenSession) TransferPauser(newPauser common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferPauser(&_Token.TransactOpts, newPauser)
}

// TransferPauser is a paid mutator transaction binding the contract method 0x4421ea21.
//
// Solidity: function transferPauser(address newPauser) returns()
func (_Token *TokenTransactorSession) TransferPauser(newPauser common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferPauser(&_Token.TransactOpts, newPauser)
}

// TransferWhitelistAdmin is a paid mutator transaction binding the contract method 0x73cfffa0.
//
// Solidity: function transferWhitelistAdmin(address newBurner) returns()
func (_Token *TokenTransactor) TransferWhitelistAdmin(opts *bind.TransactOpts, newBurner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferWhitelistAdmin", newBurner)
}

// TransferWhitelistAdmin is a paid mutator transaction binding the contract method 0x73cfffa0.
//
// Solidity: function transferWhitelistAdmin(address newBurner) returns()
func (_Token *TokenSession) TransferWhitelistAdmin(newBurner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferWhitelistAdmin(&_Token.TransactOpts, newBurner)
}

// TransferWhitelistAdmin is a paid mutator transaction binding the contract method 0x73cfffa0.
//
// Solidity: function transferWhitelistAdmin(address newBurner) returns()
func (_Token *TokenTransactorSession) TransferWhitelistAdmin(newBurner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferWhitelistAdmin(&_Token.TransactOpts, newBurner)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenBurnerTransferredIterator is returned from FilterBurnerTransferred and is used to iterate over the raw logs and unpacked data for BurnerTransferred events raised by the Token contract.
type TokenBurnerTransferredIterator struct {
	Event *TokenBurnerTransferred // Event containing the contract specifics and raw log

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
func (it *TokenBurnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBurnerTransferred)
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
		it.Event = new(TokenBurnerTransferred)
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
func (it *TokenBurnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBurnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBurnerTransferred represents a BurnerTransferred event raised by the Token contract.
type TokenBurnerTransferred struct {
	PreviousBurner common.Address
	NewBurner      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBurnerTransferred is a free log retrieval operation binding the contract event 0xeee90df09f0bc4f2408d904f2b5c22873e54185001290d225b5b2ced52128149.
//
// Solidity: event BurnerTransferred(address indexed previousBurner, address indexed newBurner)
func (_Token *TokenFilterer) FilterBurnerTransferred(opts *bind.FilterOpts, previousBurner []common.Address, newBurner []common.Address) (*TokenBurnerTransferredIterator, error) {

	var previousBurnerRule []interface{}
	for _, previousBurnerItem := range previousBurner {
		previousBurnerRule = append(previousBurnerRule, previousBurnerItem)
	}
	var newBurnerRule []interface{}
	for _, newBurnerItem := range newBurner {
		newBurnerRule = append(newBurnerRule, newBurnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "BurnerTransferred", previousBurnerRule, newBurnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenBurnerTransferredIterator{contract: _Token.contract, event: "BurnerTransferred", logs: logs, sub: sub}, nil
}

// WatchBurnerTransferred is a free log subscription operation binding the contract event 0xeee90df09f0bc4f2408d904f2b5c22873e54185001290d225b5b2ced52128149.
//
// Solidity: event BurnerTransferred(address indexed previousBurner, address indexed newBurner)
func (_Token *TokenFilterer) WatchBurnerTransferred(opts *bind.WatchOpts, sink chan<- *TokenBurnerTransferred, previousBurner []common.Address, newBurner []common.Address) (event.Subscription, error) {

	var previousBurnerRule []interface{}
	for _, previousBurnerItem := range previousBurner {
		previousBurnerRule = append(previousBurnerRule, previousBurnerItem)
	}
	var newBurnerRule []interface{}
	for _, newBurnerItem := range newBurner {
		newBurnerRule = append(newBurnerRule, newBurnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "BurnerTransferred", previousBurnerRule, newBurnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBurnerTransferred)
				if err := _Token.contract.UnpackLog(event, "BurnerTransferred", log); err != nil {
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

// ParseBurnerTransferred is a log parse operation binding the contract event 0xeee90df09f0bc4f2408d904f2b5c22873e54185001290d225b5b2ced52128149.
//
// Solidity: event BurnerTransferred(address indexed previousBurner, address indexed newBurner)
func (_Token *TokenFilterer) ParseBurnerTransferred(log types.Log) (*TokenBurnerTransferred, error) {
	event := new(TokenBurnerTransferred)
	if err := _Token.contract.UnpackLog(event, "BurnerTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenFreezeIterator is returned from FilterFreeze and is used to iterate over the raw logs and unpacked data for Freeze events raised by the Token contract.
type TokenFreezeIterator struct {
	Event *TokenFreeze // Event containing the contract specifics and raw log

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
func (it *TokenFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFreeze)
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
		it.Event = new(TokenFreeze)
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
func (it *TokenFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFreeze represents a Freeze event raised by the Token contract.
type TokenFreeze struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFreeze is a free log retrieval operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 amount)
func (_Token *TokenFilterer) FilterFreeze(opts *bind.FilterOpts, from []common.Address) (*TokenFreezeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return &TokenFreezeIterator{contract: _Token.contract, event: "Freeze", logs: logs, sub: sub}, nil
}

// WatchFreeze is a free log subscription operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 amount)
func (_Token *TokenFilterer) WatchFreeze(opts *bind.WatchOpts, sink chan<- *TokenFreeze, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Freeze", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFreeze)
				if err := _Token.contract.UnpackLog(event, "Freeze", log); err != nil {
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

// ParseFreeze is a log parse operation binding the contract event 0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0.
//
// Solidity: event Freeze(address indexed from, uint256 amount)
func (_Token *TokenFilterer) ParseFreeze(log types.Log) (*TokenFreeze, error) {
	event := new(TokenFreeze)
	if err := _Token.contract.UnpackLog(event, "Freeze", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenFrozenTransferIterator is returned from FilterFrozenTransfer and is used to iterate over the raw logs and unpacked data for FrozenTransfer events raised by the Token contract.
type TokenFrozenTransferIterator struct {
	Event *TokenFrozenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenFrozenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFrozenTransfer)
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
		it.Event = new(TokenFrozenTransfer)
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
func (it *TokenFrozenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFrozenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFrozenTransfer represents a FrozenTransfer event raised by the Token contract.
type TokenFrozenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFrozenTransfer is a free log retrieval operation binding the contract event 0x9958976925baa403abf86afa44976a9effff3e01931eafd0ea2689f40852fd59.
//
// Solidity: event FrozenTransfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) FilterFrozenTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenFrozenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "FrozenTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenFrozenTransferIterator{contract: _Token.contract, event: "FrozenTransfer", logs: logs, sub: sub}, nil
}

// WatchFrozenTransfer is a free log subscription operation binding the contract event 0x9958976925baa403abf86afa44976a9effff3e01931eafd0ea2689f40852fd59.
//
// Solidity: event FrozenTransfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) WatchFrozenTransfer(opts *bind.WatchOpts, sink chan<- *TokenFrozenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "FrozenTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFrozenTransfer)
				if err := _Token.contract.UnpackLog(event, "FrozenTransfer", log); err != nil {
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

// ParseFrozenTransfer is a log parse operation binding the contract event 0x9958976925baa403abf86afa44976a9effff3e01931eafd0ea2689f40852fd59.
//
// Solidity: event FrozenTransfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) ParseFrozenTransfer(log types.Log) (*TokenFrozenTransfer, error) {
	event := new(TokenFrozenTransfer)
	if err := _Token.contract.UnpackLog(event, "FrozenTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenMeltIterator is returned from FilterMelt and is used to iterate over the raw logs and unpacked data for Melt events raised by the Token contract.
type TokenMeltIterator struct {
	Event *TokenMelt // Event containing the contract specifics and raw log

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
func (it *TokenMeltIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMelt)
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
		it.Event = new(TokenMelt)
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
func (it *TokenMeltIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMeltIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMelt represents a Melt event raised by the Token contract.
type TokenMelt struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMelt is a free log retrieval operation binding the contract event 0x0bb0b52882b12b41cdf6b733954f1133183ca85efebcda11b4506bc6926d326b.
//
// Solidity: event Melt(address indexed from, uint256 amount)
func (_Token *TokenFilterer) FilterMelt(opts *bind.FilterOpts, from []common.Address) (*TokenMeltIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Melt", fromRule)
	if err != nil {
		return nil, err
	}
	return &TokenMeltIterator{contract: _Token.contract, event: "Melt", logs: logs, sub: sub}, nil
}

// WatchMelt is a free log subscription operation binding the contract event 0x0bb0b52882b12b41cdf6b733954f1133183ca85efebcda11b4506bc6926d326b.
//
// Solidity: event Melt(address indexed from, uint256 amount)
func (_Token *TokenFilterer) WatchMelt(opts *bind.WatchOpts, sink chan<- *TokenMelt, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Melt", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMelt)
				if err := _Token.contract.UnpackLog(event, "Melt", log); err != nil {
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

// ParseMelt is a log parse operation binding the contract event 0x0bb0b52882b12b41cdf6b733954f1133183ca85efebcda11b4506bc6926d326b.
//
// Solidity: event Melt(address indexed from, uint256 amount)
func (_Token *TokenFilterer) ParseMelt(log types.Log) (*TokenMelt, error) {
	event := new(TokenMelt)
	if err := _Token.contract.UnpackLog(event, "Melt", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenMelterTransferredIterator is returned from FilterMelterTransferred and is used to iterate over the raw logs and unpacked data for MelterTransferred events raised by the Token contract.
type TokenMelterTransferredIterator struct {
	Event *TokenMelterTransferred // Event containing the contract specifics and raw log

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
func (it *TokenMelterTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMelterTransferred)
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
		it.Event = new(TokenMelterTransferred)
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
func (it *TokenMelterTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMelterTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMelterTransferred represents a MelterTransferred event raised by the Token contract.
type TokenMelterTransferred struct {
	PreviousMelter common.Address
	NewMelter      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMelterTransferred is a free log retrieval operation binding the contract event 0x2e564b778e0bd386914daff857c4db4480f555c393032a2f0ab2d62de42c2977.
//
// Solidity: event MelterTransferred(address indexed previousMelter, address indexed newMelter)
func (_Token *TokenFilterer) FilterMelterTransferred(opts *bind.FilterOpts, previousMelter []common.Address, newMelter []common.Address) (*TokenMelterTransferredIterator, error) {

	var previousMelterRule []interface{}
	for _, previousMelterItem := range previousMelter {
		previousMelterRule = append(previousMelterRule, previousMelterItem)
	}
	var newMelterRule []interface{}
	for _, newMelterItem := range newMelter {
		newMelterRule = append(newMelterRule, newMelterItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "MelterTransferred", previousMelterRule, newMelterRule)
	if err != nil {
		return nil, err
	}
	return &TokenMelterTransferredIterator{contract: _Token.contract, event: "MelterTransferred", logs: logs, sub: sub}, nil
}

// WatchMelterTransferred is a free log subscription operation binding the contract event 0x2e564b778e0bd386914daff857c4db4480f555c393032a2f0ab2d62de42c2977.
//
// Solidity: event MelterTransferred(address indexed previousMelter, address indexed newMelter)
func (_Token *TokenFilterer) WatchMelterTransferred(opts *bind.WatchOpts, sink chan<- *TokenMelterTransferred, previousMelter []common.Address, newMelter []common.Address) (event.Subscription, error) {

	var previousMelterRule []interface{}
	for _, previousMelterItem := range previousMelter {
		previousMelterRule = append(previousMelterRule, previousMelterItem)
	}
	var newMelterRule []interface{}
	for _, newMelterItem := range newMelter {
		newMelterRule = append(newMelterRule, newMelterItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "MelterTransferred", previousMelterRule, newMelterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMelterTransferred)
				if err := _Token.contract.UnpackLog(event, "MelterTransferred", log); err != nil {
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

// ParseMelterTransferred is a log parse operation binding the contract event 0x2e564b778e0bd386914daff857c4db4480f555c393032a2f0ab2d62de42c2977.
//
// Solidity: event MelterTransferred(address indexed previousMelter, address indexed newMelter)
func (_Token *TokenFilterer) ParseMelterTransferred(log types.Log) (*TokenMelterTransferred, error) {
	event := new(TokenMelterTransferred)
	if err := _Token.contract.UnpackLog(event, "MelterTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenMintFrozenIterator is returned from FilterMintFrozen and is used to iterate over the raw logs and unpacked data for MintFrozen events raised by the Token contract.
type TokenMintFrozenIterator struct {
	Event *TokenMintFrozen // Event containing the contract specifics and raw log

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
func (it *TokenMintFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMintFrozen)
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
		it.Event = new(TokenMintFrozen)
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
func (it *TokenMintFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMintFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMintFrozen represents a MintFrozen event raised by the Token contract.
type TokenMintFrozen struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintFrozen is a free log retrieval operation binding the contract event 0xba89ad6709373f454c31524e8c39cef3cdd4b0e8cfde0ccddbd419a2e488be6b.
//
// Solidity: event MintFrozen(address indexed to, uint256 amount)
func (_Token *TokenFilterer) FilterMintFrozen(opts *bind.FilterOpts, to []common.Address) (*TokenMintFrozenIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "MintFrozen", toRule)
	if err != nil {
		return nil, err
	}
	return &TokenMintFrozenIterator{contract: _Token.contract, event: "MintFrozen", logs: logs, sub: sub}, nil
}

// WatchMintFrozen is a free log subscription operation binding the contract event 0xba89ad6709373f454c31524e8c39cef3cdd4b0e8cfde0ccddbd419a2e488be6b.
//
// Solidity: event MintFrozen(address indexed to, uint256 amount)
func (_Token *TokenFilterer) WatchMintFrozen(opts *bind.WatchOpts, sink chan<- *TokenMintFrozen, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "MintFrozen", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMintFrozen)
				if err := _Token.contract.UnpackLog(event, "MintFrozen", log); err != nil {
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

// ParseMintFrozen is a log parse operation binding the contract event 0xba89ad6709373f454c31524e8c39cef3cdd4b0e8cfde0ccddbd419a2e488be6b.
//
// Solidity: event MintFrozen(address indexed to, uint256 amount)
func (_Token *TokenFilterer) ParseMintFrozen(log types.Log) (*TokenMintFrozen, error) {
	event := new(TokenMintFrozen)
	if err := _Token.contract.UnpackLog(event, "MintFrozen", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenMinterTransferredIterator is returned from FilterMinterTransferred and is used to iterate over the raw logs and unpacked data for MinterTransferred events raised by the Token contract.
type TokenMinterTransferredIterator struct {
	Event *TokenMinterTransferred // Event containing the contract specifics and raw log

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
func (it *TokenMinterTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMinterTransferred)
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
		it.Event = new(TokenMinterTransferred)
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
func (it *TokenMinterTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMinterTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMinterTransferred represents a MinterTransferred event raised by the Token contract.
type TokenMinterTransferred struct {
	PreviousMinter common.Address
	NewMinter      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinterTransferred is a free log retrieval operation binding the contract event 0x02ad39e5173f89bdd5497202bd74024b5da045106c3163ddb078d2e89ff6d6de.
//
// Solidity: event MinterTransferred(address indexed previousMinter, address indexed newMinter)
func (_Token *TokenFilterer) FilterMinterTransferred(opts *bind.FilterOpts, previousMinter []common.Address, newMinter []common.Address) (*TokenMinterTransferredIterator, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "MinterTransferred", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &TokenMinterTransferredIterator{contract: _Token.contract, event: "MinterTransferred", logs: logs, sub: sub}, nil
}

// WatchMinterTransferred is a free log subscription operation binding the contract event 0x02ad39e5173f89bdd5497202bd74024b5da045106c3163ddb078d2e89ff6d6de.
//
// Solidity: event MinterTransferred(address indexed previousMinter, address indexed newMinter)
func (_Token *TokenFilterer) WatchMinterTransferred(opts *bind.WatchOpts, sink chan<- *TokenMinterTransferred, previousMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "MinterTransferred", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMinterTransferred)
				if err := _Token.contract.UnpackLog(event, "MinterTransferred", log); err != nil {
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

// ParseMinterTransferred is a log parse operation binding the contract event 0x02ad39e5173f89bdd5497202bd74024b5da045106c3163ddb078d2e89ff6d6de.
//
// Solidity: event MinterTransferred(address indexed previousMinter, address indexed newMinter)
func (_Token *TokenFilterer) ParseMinterTransferred(log types.Log) (*TokenMinterTransferred, error) {
	event := new(TokenMinterTransferred)
	if err := _Token.contract.UnpackLog(event, "MinterTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Token *TokenFilterer) ParseOwnershipTransferred(log types.Log) (*TokenOwnershipTransferred, error) {
	event := new(TokenOwnershipTransferred)
	if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenPauserTransferredIterator is returned from FilterPauserTransferred and is used to iterate over the raw logs and unpacked data for PauserTransferred events raised by the Token contract.
type TokenPauserTransferredIterator struct {
	Event *TokenPauserTransferred // Event containing the contract specifics and raw log

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
func (it *TokenPauserTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPauserTransferred)
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
		it.Event = new(TokenPauserTransferred)
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
func (it *TokenPauserTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPauserTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPauserTransferred represents a PauserTransferred event raised by the Token contract.
type TokenPauserTransferred struct {
	PreviousPauser common.Address
	NewPauser      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPauserTransferred is a free log retrieval operation binding the contract event 0x51c4874e0f23f262e04a38c51751336dde72126d67f53eb672aaff02996b3ef6.
//
// Solidity: event PauserTransferred(address indexed previousPauser, address indexed newPauser)
func (_Token *TokenFilterer) FilterPauserTransferred(opts *bind.FilterOpts, previousPauser []common.Address, newPauser []common.Address) (*TokenPauserTransferredIterator, error) {

	var previousPauserRule []interface{}
	for _, previousPauserItem := range previousPauser {
		previousPauserRule = append(previousPauserRule, previousPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "PauserTransferred", previousPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return &TokenPauserTransferredIterator{contract: _Token.contract, event: "PauserTransferred", logs: logs, sub: sub}, nil
}

// WatchPauserTransferred is a free log subscription operation binding the contract event 0x51c4874e0f23f262e04a38c51751336dde72126d67f53eb672aaff02996b3ef6.
//
// Solidity: event PauserTransferred(address indexed previousPauser, address indexed newPauser)
func (_Token *TokenFilterer) WatchPauserTransferred(opts *bind.WatchOpts, sink chan<- *TokenPauserTransferred, previousPauser []common.Address, newPauser []common.Address) (event.Subscription, error) {

	var previousPauserRule []interface{}
	for _, previousPauserItem := range previousPauser {
		previousPauserRule = append(previousPauserRule, previousPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "PauserTransferred", previousPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPauserTransferred)
				if err := _Token.contract.UnpackLog(event, "PauserTransferred", log); err != nil {
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

// ParsePauserTransferred is a log parse operation binding the contract event 0x51c4874e0f23f262e04a38c51751336dde72126d67f53eb672aaff02996b3ef6.
//
// Solidity: event PauserTransferred(address indexed previousPauser, address indexed newPauser)
func (_Token *TokenFilterer) ParsePauserTransferred(log types.Log) (*TokenPauserTransferred, error) {
	event := new(TokenPauserTransferred)
	if err := _Token.contract.UnpackLog(event, "PauserTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenWhitelistAdminTransferredIterator is returned from FilterWhitelistAdminTransferred and is used to iterate over the raw logs and unpacked data for WhitelistAdminTransferred events raised by the Token contract.
type TokenWhitelistAdminTransferredIterator struct {
	Event *TokenWhitelistAdminTransferred // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistAdminTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistAdminTransferred)
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
		it.Event = new(TokenWhitelistAdminTransferred)
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
func (it *TokenWhitelistAdminTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistAdminTransferred represents a WhitelistAdminTransferred event raised by the Token contract.
type TokenWhitelistAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdminTransferred is a free log retrieval operation binding the contract event 0x780008ae49c9be745883f9a2ca6d2e68e9063350aaa0782ee3f00316965d2515.
//
// Solidity: event WhitelistAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_Token *TokenFilterer) FilterWhitelistAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*TokenWhitelistAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "WhitelistAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistAdminTransferredIterator{contract: _Token.contract, event: "WhitelistAdminTransferred", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdminTransferred is a free log subscription operation binding the contract event 0x780008ae49c9be745883f9a2ca6d2e68e9063350aaa0782ee3f00316965d2515.
//
// Solidity: event WhitelistAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_Token *TokenFilterer) WatchWhitelistAdminTransferred(opts *bind.WatchOpts, sink chan<- *TokenWhitelistAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "WhitelistAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistAdminTransferred)
				if err := _Token.contract.UnpackLog(event, "WhitelistAdminTransferred", log); err != nil {
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

// ParseWhitelistAdminTransferred is a log parse operation binding the contract event 0x780008ae49c9be745883f9a2ca6d2e68e9063350aaa0782ee3f00316965d2515.
//
// Solidity: event WhitelistAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_Token *TokenFilterer) ParseWhitelistAdminTransferred(log types.Log) (*TokenWhitelistAdminTransferred, error) {
	event := new(TokenWhitelistAdminTransferred)
	if err := _Token.contract.UnpackLog(event, "WhitelistAdminTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
