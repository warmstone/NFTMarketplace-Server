// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PLATFORM_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_BID_INCREMENT_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctionCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deprecatedTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startPriceUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"highestBidTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"highestBidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAuction\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startPriceUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"durationHours\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disablePriceFeed\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startPriceUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"highestBidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialFeeRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paymentTokenAllowed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingReturns\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"placeBid\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"platformFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceFeeds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"feed\",\"type\":\"address\",\"internalType\":\"contractAggregatorV3Interface\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxStaleness\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteTokenAmount\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteUsdAmount\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setERC20PriceFeed\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxStaleness\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentTokenAllowed\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPlatformFee\",\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxStaleness\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFeeRecipient\",\"inputs\":[{\"name\":\"newRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawBid\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuctionCreated\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startPriceUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionEnded\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"usdPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidPlaced\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bidUsdAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidWithdrawn\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRecipientUpdated\",\"inputs\":[{\"name\":\"oldRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentTokenAllowedUpdated\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlatformFeeUpdated\",\"inputs\":[{\"name\":\"oldFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFeedConfigured\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feed\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"maxStaleness\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFeedDisabled\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AuctionEndedAlready\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionNotEnded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BidTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectPayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDuration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOraclePrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoyalty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketplaceNotApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoPendingReturn\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotFeeRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PaymentTokenNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriceFeedNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SellerCannotBid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StaleOraclePrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Contracts *ContractsCaller) BASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Contracts *ContractsSession) BASISPOINTS() (*big.Int, error) {
	return _Contracts.Contract.BASISPOINTS(&_Contracts.CallOpts)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Contracts *ContractsCallerSession) BASISPOINTS() (*big.Int, error) {
	return _Contracts.Contract.BASISPOINTS(&_Contracts.CallOpts)
}

// MAXPLATFORMFEE is a free data retrieval call binding the contract method 0x3998a681.
//
// Solidity: function MAX_PLATFORM_FEE() view returns(uint256)
func (_Contracts *ContractsCaller) MAXPLATFORMFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "MAX_PLATFORM_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPLATFORMFEE is a free data retrieval call binding the contract method 0x3998a681.
//
// Solidity: function MAX_PLATFORM_FEE() view returns(uint256)
func (_Contracts *ContractsSession) MAXPLATFORMFEE() (*big.Int, error) {
	return _Contracts.Contract.MAXPLATFORMFEE(&_Contracts.CallOpts)
}

// MAXPLATFORMFEE is a free data retrieval call binding the contract method 0x3998a681.
//
// Solidity: function MAX_PLATFORM_FEE() view returns(uint256)
func (_Contracts *ContractsCallerSession) MAXPLATFORMFEE() (*big.Int, error) {
	return _Contracts.Contract.MAXPLATFORMFEE(&_Contracts.CallOpts)
}

// MINBIDINCREMENTBPS is a free data retrieval call binding the contract method 0x14414cc0.
//
// Solidity: function MIN_BID_INCREMENT_BPS() view returns(uint256)
func (_Contracts *ContractsCaller) MINBIDINCREMENTBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "MIN_BID_INCREMENT_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBIDINCREMENTBPS is a free data retrieval call binding the contract method 0x14414cc0.
//
// Solidity: function MIN_BID_INCREMENT_BPS() view returns(uint256)
func (_Contracts *ContractsSession) MINBIDINCREMENTBPS() (*big.Int, error) {
	return _Contracts.Contract.MINBIDINCREMENTBPS(&_Contracts.CallOpts)
}

// MINBIDINCREMENTBPS is a free data retrieval call binding the contract method 0x14414cc0.
//
// Solidity: function MIN_BID_INCREMENT_BPS() view returns(uint256)
func (_Contracts *ContractsCallerSession) MINBIDINCREMENTBPS() (*big.Int, error) {
	return _Contracts.Contract.MINBIDINCREMENTBPS(&_Contracts.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contracts *ContractsCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contracts *ContractsSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Contracts.Contract.UPGRADEINTERFACEVERSION(&_Contracts.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contracts *ContractsCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Contracts.Contract.UPGRADEINTERFACEVERSION(&_Contracts.CallOpts)
}

// AuctionCounter is a free data retrieval call binding the contract method 0xa7e76644.
//
// Solidity: function auctionCounter() view returns(uint256)
func (_Contracts *ContractsCaller) AuctionCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "auctionCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionCounter is a free data retrieval call binding the contract method 0xa7e76644.
//
// Solidity: function auctionCounter() view returns(uint256)
func (_Contracts *ContractsSession) AuctionCounter() (*big.Int, error) {
	return _Contracts.Contract.AuctionCounter(&_Contracts.CallOpts)
}

// AuctionCounter is a free data retrieval call binding the contract method 0xa7e76644.
//
// Solidity: function auctionCounter() view returns(uint256)
func (_Contracts *ContractsCallerSession) AuctionCounter() (*big.Int, error) {
	return _Contracts.Contract.AuctionCounter(&_Contracts.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, address deprecatedTokenAddress, uint256 startPriceUsd, uint256 highestBidUsd, address highestBidder, uint256 endTime, bool active, address highestBidTokenAddress, uint256 highestBidAmount)
func (_Contracts *ContractsCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller                 common.Address
	NftContract            common.Address
	TokenId                *big.Int
	DeprecatedTokenAddress common.Address
	StartPriceUsd          *big.Int
	HighestBidUsd          *big.Int
	HighestBidder          common.Address
	EndTime                *big.Int
	Active                 bool
	HighestBidTokenAddress common.Address
	HighestBidAmount       *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		Seller                 common.Address
		NftContract            common.Address
		TokenId                *big.Int
		DeprecatedTokenAddress common.Address
		StartPriceUsd          *big.Int
		HighestBidUsd          *big.Int
		HighestBidder          common.Address
		EndTime                *big.Int
		Active                 bool
		HighestBidTokenAddress common.Address
		HighestBidAmount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeprecatedTokenAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.StartPriceUsd = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HighestBidUsd = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.EndTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.HighestBidTokenAddress = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.HighestBidAmount = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, address deprecatedTokenAddress, uint256 startPriceUsd, uint256 highestBidUsd, address highestBidder, uint256 endTime, bool active, address highestBidTokenAddress, uint256 highestBidAmount)
func (_Contracts *ContractsSession) Auctions(arg0 *big.Int) (struct {
	Seller                 common.Address
	NftContract            common.Address
	TokenId                *big.Int
	DeprecatedTokenAddress common.Address
	StartPriceUsd          *big.Int
	HighestBidUsd          *big.Int
	HighestBidder          common.Address
	EndTime                *big.Int
	Active                 bool
	HighestBidTokenAddress common.Address
	HighestBidAmount       *big.Int
}, error) {
	return _Contracts.Contract.Auctions(&_Contracts.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, address deprecatedTokenAddress, uint256 startPriceUsd, uint256 highestBidUsd, address highestBidder, uint256 endTime, bool active, address highestBidTokenAddress, uint256 highestBidAmount)
func (_Contracts *ContractsCallerSession) Auctions(arg0 *big.Int) (struct {
	Seller                 common.Address
	NftContract            common.Address
	TokenId                *big.Int
	DeprecatedTokenAddress common.Address
	StartPriceUsd          *big.Int
	HighestBidUsd          *big.Int
	HighestBidder          common.Address
	EndTime                *big.Int
	Active                 bool
	HighestBidTokenAddress common.Address
	HighestBidAmount       *big.Int
}, error) {
	return _Contracts.Contract.Auctions(&_Contracts.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Contracts *ContractsCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Contracts *ContractsSession) FeeRecipient() (common.Address, error) {
	return _Contracts.Contract.FeeRecipient(&_Contracts.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Contracts *ContractsCallerSession) FeeRecipient() (common.Address, error) {
	return _Contracts.Contract.FeeRecipient(&_Contracts.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 auctionId) view returns(address seller, address nftContract, uint256 tokenId, uint256 startPriceUsd, uint256 highestBidUsd, address highestBidTokenAddress, uint256 highestBidAmount, address highestBidder, uint256 endTime, bool active)
func (_Contracts *ContractsCaller) GetAuction(opts *bind.CallOpts, auctionId *big.Int) (struct {
	Seller                 common.Address
	NftContract            common.Address
	TokenId                *big.Int
	StartPriceUsd          *big.Int
	HighestBidUsd          *big.Int
	HighestBidTokenAddress common.Address
	HighestBidAmount       *big.Int
	HighestBidder          common.Address
	EndTime                *big.Int
	Active                 bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAuction", auctionId)

	outstruct := new(struct {
		Seller                 common.Address
		NftContract            common.Address
		TokenId                *big.Int
		StartPriceUsd          *big.Int
		HighestBidUsd          *big.Int
		HighestBidTokenAddress common.Address
		HighestBidAmount       *big.Int
		HighestBidder          common.Address
		EndTime                *big.Int
		Active                 bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartPriceUsd = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.HighestBidUsd = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HighestBidTokenAddress = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.HighestBidAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.EndTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 auctionId) view returns(address seller, address nftContract, uint256 tokenId, uint256 startPriceUsd, uint256 highestBidUsd, address highestBidTokenAddress, uint256 highestBidAmount, address highestBidder, uint256 endTime, bool active)
func (_Contracts *ContractsSession) GetAuction(auctionId *big.Int) (struct {
	Seller                 common.Address
	NftContract            common.Address
	TokenId                *big.Int
	StartPriceUsd          *big.Int
	HighestBidUsd          *big.Int
	HighestBidTokenAddress common.Address
	HighestBidAmount       *big.Int
	HighestBidder          common.Address
	EndTime                *big.Int
	Active                 bool
}, error) {
	return _Contracts.Contract.GetAuction(&_Contracts.CallOpts, auctionId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 auctionId) view returns(address seller, address nftContract, uint256 tokenId, uint256 startPriceUsd, uint256 highestBidUsd, address highestBidTokenAddress, uint256 highestBidAmount, address highestBidder, uint256 endTime, bool active)
func (_Contracts *ContractsCallerSession) GetAuction(auctionId *big.Int) (struct {
	Seller                 common.Address
	NftContract            common.Address
	TokenId                *big.Int
	StartPriceUsd          *big.Int
	HighestBidUsd          *big.Int
	HighestBidTokenAddress common.Address
	HighestBidAmount       *big.Int
	HighestBidder          common.Address
	EndTime                *big.Int
	Active                 bool
}, error) {
	return _Contracts.Contract.GetAuction(&_Contracts.CallOpts, auctionId)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Contracts *ContractsCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Contracts *ContractsSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Contracts.Contract.OnERC721Received(&_Contracts.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Contracts *ContractsCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Contracts.Contract.OnERC721Received(&_Contracts.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// PaymentTokenAllowed is a free data retrieval call binding the contract method 0x06d49f92.
//
// Solidity: function paymentTokenAllowed(address ) view returns(bool)
func (_Contracts *ContractsCaller) PaymentTokenAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "paymentTokenAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PaymentTokenAllowed is a free data retrieval call binding the contract method 0x06d49f92.
//
// Solidity: function paymentTokenAllowed(address ) view returns(bool)
func (_Contracts *ContractsSession) PaymentTokenAllowed(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.PaymentTokenAllowed(&_Contracts.CallOpts, arg0)
}

// PaymentTokenAllowed is a free data retrieval call binding the contract method 0x06d49f92.
//
// Solidity: function paymentTokenAllowed(address ) view returns(bool)
func (_Contracts *ContractsCallerSession) PaymentTokenAllowed(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.PaymentTokenAllowed(&_Contracts.CallOpts, arg0)
}

// PendingReturns is a free data retrieval call binding the contract method 0x3911bdf6.
//
// Solidity: function pendingReturns(uint256 , address , address ) view returns(uint256)
func (_Contracts *ContractsCaller) PendingReturns(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "pendingReturns", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReturns is a free data retrieval call binding the contract method 0x3911bdf6.
//
// Solidity: function pendingReturns(uint256 , address , address ) view returns(uint256)
func (_Contracts *ContractsSession) PendingReturns(arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _Contracts.Contract.PendingReturns(&_Contracts.CallOpts, arg0, arg1, arg2)
}

// PendingReturns is a free data retrieval call binding the contract method 0x3911bdf6.
//
// Solidity: function pendingReturns(uint256 , address , address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) PendingReturns(arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _Contracts.Contract.PendingReturns(&_Contracts.CallOpts, arg0, arg1, arg2)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_Contracts *ContractsCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_Contracts *ContractsSession) PlatformFee() (*big.Int, error) {
	return _Contracts.Contract.PlatformFee(&_Contracts.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_Contracts *ContractsCallerSession) PlatformFee() (*big.Int, error) {
	return _Contracts.Contract.PlatformFee(&_Contracts.CallOpts)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address feed, uint8 tokenDecimals, uint256 maxStaleness, bool active)
func (_Contracts *ContractsCaller) PriceFeeds(opts *bind.CallOpts, arg0 common.Address) (struct {
	Feed          common.Address
	TokenDecimals uint8
	MaxStaleness  *big.Int
	Active        bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "priceFeeds", arg0)

	outstruct := new(struct {
		Feed          common.Address
		TokenDecimals uint8
		MaxStaleness  *big.Int
		Active        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Feed = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenDecimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.MaxStaleness = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address feed, uint8 tokenDecimals, uint256 maxStaleness, bool active)
func (_Contracts *ContractsSession) PriceFeeds(arg0 common.Address) (struct {
	Feed          common.Address
	TokenDecimals uint8
	MaxStaleness  *big.Int
	Active        bool
}, error) {
	return _Contracts.Contract.PriceFeeds(&_Contracts.CallOpts, arg0)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address feed, uint8 tokenDecimals, uint256 maxStaleness, bool active)
func (_Contracts *ContractsCallerSession) PriceFeeds(arg0 common.Address) (struct {
	Feed          common.Address
	TokenDecimals uint8
	MaxStaleness  *big.Int
	Active        bool
}, error) {
	return _Contracts.Contract.PriceFeeds(&_Contracts.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contracts *ContractsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contracts *ContractsSession) ProxiableUUID() ([32]byte, error) {
	return _Contracts.Contract.ProxiableUUID(&_Contracts.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contracts *ContractsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Contracts.Contract.ProxiableUUID(&_Contracts.CallOpts)
}

// QuoteTokenAmount is a free data retrieval call binding the contract method 0xf359b59d.
//
// Solidity: function quoteTokenAmount(address tokenAddress, uint256 usdAmount) view returns(uint256 tokenAmount)
func (_Contracts *ContractsCaller) QuoteTokenAmount(opts *bind.CallOpts, tokenAddress common.Address, usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quoteTokenAmount", tokenAddress, usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteTokenAmount is a free data retrieval call binding the contract method 0xf359b59d.
//
// Solidity: function quoteTokenAmount(address tokenAddress, uint256 usdAmount) view returns(uint256 tokenAmount)
func (_Contracts *ContractsSession) QuoteTokenAmount(tokenAddress common.Address, usdAmount *big.Int) (*big.Int, error) {
	return _Contracts.Contract.QuoteTokenAmount(&_Contracts.CallOpts, tokenAddress, usdAmount)
}

// QuoteTokenAmount is a free data retrieval call binding the contract method 0xf359b59d.
//
// Solidity: function quoteTokenAmount(address tokenAddress, uint256 usdAmount) view returns(uint256 tokenAmount)
func (_Contracts *ContractsCallerSession) QuoteTokenAmount(tokenAddress common.Address, usdAmount *big.Int) (*big.Int, error) {
	return _Contracts.Contract.QuoteTokenAmount(&_Contracts.CallOpts, tokenAddress, usdAmount)
}

// QuoteUsdAmount is a free data retrieval call binding the contract method 0x80d93fb8.
//
// Solidity: function quoteUsdAmount(address tokenAddress, uint256 tokenAmount) view returns(uint256 usdAmount)
func (_Contracts *ContractsCaller) QuoteUsdAmount(opts *bind.CallOpts, tokenAddress common.Address, tokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quoteUsdAmount", tokenAddress, tokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteUsdAmount is a free data retrieval call binding the contract method 0x80d93fb8.
//
// Solidity: function quoteUsdAmount(address tokenAddress, uint256 tokenAmount) view returns(uint256 usdAmount)
func (_Contracts *ContractsSession) QuoteUsdAmount(tokenAddress common.Address, tokenAmount *big.Int) (*big.Int, error) {
	return _Contracts.Contract.QuoteUsdAmount(&_Contracts.CallOpts, tokenAddress, tokenAmount)
}

// QuoteUsdAmount is a free data retrieval call binding the contract method 0x80d93fb8.
//
// Solidity: function quoteUsdAmount(address tokenAddress, uint256 tokenAmount) view returns(uint256 usdAmount)
func (_Contracts *ContractsCallerSession) QuoteUsdAmount(tokenAddress common.Address, tokenAmount *big.Int) (*big.Int, error) {
	return _Contracts.Contract.QuoteUsdAmount(&_Contracts.CallOpts, tokenAddress, tokenAmount)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Contracts *ContractsCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Contracts *ContractsSession) Version() (string, error) {
	return _Contracts.Contract.Version(&_Contracts.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Contracts *ContractsCallerSession) Version() (string, error) {
	return _Contracts.Contract.Version(&_Contracts.CallOpts)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x61beb1d7.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startPriceUsd, uint256 durationHours) returns(uint256)
func (_Contracts *ContractsTransactor) CreateAuction(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, startPriceUsd *big.Int, durationHours *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createAuction", nftContract, tokenId, startPriceUsd, durationHours)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x61beb1d7.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startPriceUsd, uint256 durationHours) returns(uint256)
func (_Contracts *ContractsSession) CreateAuction(nftContract common.Address, tokenId *big.Int, startPriceUsd *big.Int, durationHours *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateAuction(&_Contracts.TransactOpts, nftContract, tokenId, startPriceUsd, durationHours)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x61beb1d7.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startPriceUsd, uint256 durationHours) returns(uint256)
func (_Contracts *ContractsTransactorSession) CreateAuction(nftContract common.Address, tokenId *big.Int, startPriceUsd *big.Int, durationHours *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateAuction(&_Contracts.TransactOpts, nftContract, tokenId, startPriceUsd, durationHours)
}

// DisablePriceFeed is a paid mutator transaction binding the contract method 0xdd15e73a.
//
// Solidity: function disablePriceFeed(address tokenAddress) returns()
func (_Contracts *ContractsTransactor) DisablePriceFeed(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "disablePriceFeed", tokenAddress)
}

// DisablePriceFeed is a paid mutator transaction binding the contract method 0xdd15e73a.
//
// Solidity: function disablePriceFeed(address tokenAddress) returns()
func (_Contracts *ContractsSession) DisablePriceFeed(tokenAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.DisablePriceFeed(&_Contracts.TransactOpts, tokenAddress)
}

// DisablePriceFeed is a paid mutator transaction binding the contract method 0xdd15e73a.
//
// Solidity: function disablePriceFeed(address tokenAddress) returns()
func (_Contracts *ContractsTransactorSession) DisablePriceFeed(tokenAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.DisablePriceFeed(&_Contracts.TransactOpts, tokenAddress)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_Contracts *ContractsTransactor) EndAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "endAuction", auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_Contracts *ContractsSession) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.EndAuction(&_Contracts.TransactOpts, auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_Contracts *ContractsTransactorSession) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.EndAuction(&_Contracts.TransactOpts, auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address initialFeeRecipient) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialFeeRecipient common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize", initialOwner, initialFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address initialFeeRecipient) returns()
func (_Contracts *ContractsSession) Initialize(initialOwner common.Address, initialFeeRecipient common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, initialOwner, initialFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address initialFeeRecipient) returns()
func (_Contracts *ContractsTransactorSession) Initialize(initialOwner common.Address, initialFeeRecipient common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, initialOwner, initialFeeRecipient)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xf1e2f348.
//
// Solidity: function placeBid(uint256 auctionId, address tokenAddress, uint256 bidAmount) payable returns()
func (_Contracts *ContractsTransactor) PlaceBid(opts *bind.TransactOpts, auctionId *big.Int, tokenAddress common.Address, bidAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "placeBid", auctionId, tokenAddress, bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xf1e2f348.
//
// Solidity: function placeBid(uint256 auctionId, address tokenAddress, uint256 bidAmount) payable returns()
func (_Contracts *ContractsSession) PlaceBid(auctionId *big.Int, tokenAddress common.Address, bidAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PlaceBid(&_Contracts.TransactOpts, auctionId, tokenAddress, bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xf1e2f348.
//
// Solidity: function placeBid(uint256 auctionId, address tokenAddress, uint256 bidAmount) payable returns()
func (_Contracts *ContractsTransactorSession) PlaceBid(auctionId *big.Int, tokenAddress common.Address, bidAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PlaceBid(&_Contracts.TransactOpts, auctionId, tokenAddress, bidAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// SetERC20PriceFeed is a paid mutator transaction binding the contract method 0x3902bd89.
//
// Solidity: function setERC20PriceFeed(address tokenAddress, address feed, uint256 maxStaleness) returns()
func (_Contracts *ContractsTransactor) SetERC20PriceFeed(opts *bind.TransactOpts, tokenAddress common.Address, feed common.Address, maxStaleness *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setERC20PriceFeed", tokenAddress, feed, maxStaleness)
}

// SetERC20PriceFeed is a paid mutator transaction binding the contract method 0x3902bd89.
//
// Solidity: function setERC20PriceFeed(address tokenAddress, address feed, uint256 maxStaleness) returns()
func (_Contracts *ContractsSession) SetERC20PriceFeed(tokenAddress common.Address, feed common.Address, maxStaleness *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetERC20PriceFeed(&_Contracts.TransactOpts, tokenAddress, feed, maxStaleness)
}

// SetERC20PriceFeed is a paid mutator transaction binding the contract method 0x3902bd89.
//
// Solidity: function setERC20PriceFeed(address tokenAddress, address feed, uint256 maxStaleness) returns()
func (_Contracts *ContractsTransactorSession) SetERC20PriceFeed(tokenAddress common.Address, feed common.Address, maxStaleness *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetERC20PriceFeed(&_Contracts.TransactOpts, tokenAddress, feed, maxStaleness)
}

// SetPaymentTokenAllowed is a paid mutator transaction binding the contract method 0x28336098.
//
// Solidity: function setPaymentTokenAllowed(address tokenAddress, bool allowed) returns()
func (_Contracts *ContractsTransactor) SetPaymentTokenAllowed(opts *bind.TransactOpts, tokenAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPaymentTokenAllowed", tokenAddress, allowed)
}

// SetPaymentTokenAllowed is a paid mutator transaction binding the contract method 0x28336098.
//
// Solidity: function setPaymentTokenAllowed(address tokenAddress, bool allowed) returns()
func (_Contracts *ContractsSession) SetPaymentTokenAllowed(tokenAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetPaymentTokenAllowed(&_Contracts.TransactOpts, tokenAddress, allowed)
}

// SetPaymentTokenAllowed is a paid mutator transaction binding the contract method 0x28336098.
//
// Solidity: function setPaymentTokenAllowed(address tokenAddress, bool allowed) returns()
func (_Contracts *ContractsTransactorSession) SetPaymentTokenAllowed(tokenAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetPaymentTokenAllowed(&_Contracts.TransactOpts, tokenAddress, allowed)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 newFee) returns()
func (_Contracts *ContractsTransactor) SetPlatformFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPlatformFee", newFee)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 newFee) returns()
func (_Contracts *ContractsSession) SetPlatformFee(newFee *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetPlatformFee(&_Contracts.TransactOpts, newFee)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 newFee) returns()
func (_Contracts *ContractsTransactorSession) SetPlatformFee(newFee *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetPlatformFee(&_Contracts.TransactOpts, newFee)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x2bcd09ee.
//
// Solidity: function setPriceFeed(address tokenAddress, address feed, uint8 tokenDecimals, uint256 maxStaleness) returns()
func (_Contracts *ContractsTransactor) SetPriceFeed(opts *bind.TransactOpts, tokenAddress common.Address, feed common.Address, tokenDecimals uint8, maxStaleness *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPriceFeed", tokenAddress, feed, tokenDecimals, maxStaleness)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x2bcd09ee.
//
// Solidity: function setPriceFeed(address tokenAddress, address feed, uint8 tokenDecimals, uint256 maxStaleness) returns()
func (_Contracts *ContractsSession) SetPriceFeed(tokenAddress common.Address, feed common.Address, tokenDecimals uint8, maxStaleness *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetPriceFeed(&_Contracts.TransactOpts, tokenAddress, feed, tokenDecimals, maxStaleness)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x2bcd09ee.
//
// Solidity: function setPriceFeed(address tokenAddress, address feed, uint8 tokenDecimals, uint256 maxStaleness) returns()
func (_Contracts *ContractsTransactorSession) SetPriceFeed(tokenAddress common.Address, feed common.Address, tokenDecimals uint8, maxStaleness *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetPriceFeed(&_Contracts.TransactOpts, tokenAddress, feed, tokenDecimals, maxStaleness)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newRecipient) returns()
func (_Contracts *ContractsTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, newRecipient common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateFeeRecipient", newRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newRecipient) returns()
func (_Contracts *ContractsSession) UpdateFeeRecipient(newRecipient common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateFeeRecipient(&_Contracts.TransactOpts, newRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newRecipient) returns()
func (_Contracts *ContractsTransactorSession) UpdateFeeRecipient(newRecipient common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateFeeRecipient(&_Contracts.TransactOpts, newRecipient)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contracts *ContractsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contracts *ContractsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpgradeToAndCall(&_Contracts.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contracts *ContractsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.UpgradeToAndCall(&_Contracts.TransactOpts, newImplementation, data)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x801f597f.
//
// Solidity: function withdrawBid(uint256 auctionId, address tokenAddress) returns()
func (_Contracts *ContractsTransactor) WithdrawBid(opts *bind.TransactOpts, auctionId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawBid", auctionId, tokenAddress)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x801f597f.
//
// Solidity: function withdrawBid(uint256 auctionId, address tokenAddress) returns()
func (_Contracts *ContractsSession) WithdrawBid(auctionId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawBid(&_Contracts.TransactOpts, auctionId, tokenAddress)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x801f597f.
//
// Solidity: function withdrawBid(uint256 auctionId, address tokenAddress) returns()
func (_Contracts *ContractsTransactorSession) WithdrawBid(auctionId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawBid(&_Contracts.TransactOpts, auctionId, tokenAddress)
}

// ContractsAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the Contracts contract.
type ContractsAuctionCreatedIterator struct {
	Event *ContractsAuctionCreated // Event containing the contract specifics and raw log

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
func (it *ContractsAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAuctionCreated)
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
		it.Event = new(ContractsAuctionCreated)
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
func (it *ContractsAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAuctionCreated represents a AuctionCreated event raised by the Contracts contract.
type ContractsAuctionCreated struct {
	AuctionId     *big.Int
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartPriceUsd *big.Int
	EndTime       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0x06b9e486c68303eb64052e0493f906f3d93a1b7149b6b8dcff221aebd16c3513.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPriceUsd, uint256 endTime)
func (_Contracts *ContractsFilterer) FilterAuctionCreated(opts *bind.FilterOpts, auctionId []*big.Int, seller []common.Address, nftContract []common.Address) (*ContractsAuctionCreatedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AuctionCreated", auctionIdRule, sellerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAuctionCreatedIterator{contract: _Contracts.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0x06b9e486c68303eb64052e0493f906f3d93a1b7149b6b8dcff221aebd16c3513.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPriceUsd, uint256 endTime)
func (_Contracts *ContractsFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *ContractsAuctionCreated, auctionId []*big.Int, seller []common.Address, nftContract []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AuctionCreated", auctionIdRule, sellerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAuctionCreated)
				if err := _Contracts.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0x06b9e486c68303eb64052e0493f906f3d93a1b7149b6b8dcff221aebd16c3513.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPriceUsd, uint256 endTime)
func (_Contracts *ContractsFilterer) ParseAuctionCreated(log types.Log) (*ContractsAuctionCreated, error) {
	event := new(ContractsAuctionCreated)
	if err := _Contracts.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the Contracts contract.
type ContractsAuctionEndedIterator struct {
	Event *ContractsAuctionEnded // Event containing the contract specifics and raw log

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
func (it *ContractsAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAuctionEnded)
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
		it.Event = new(ContractsAuctionEnded)
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
func (it *ContractsAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAuctionEnded represents a AuctionEnded event raised by the Contracts contract.
type ContractsAuctionEnded struct {
	AuctionId    *big.Int
	Buyer        common.Address
	TokenAddress common.Address
	Price        *big.Int
	UsdPrice     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0x596165d0521c3cb4157fad2621686f086daed4663acb3d03441a92b9277f5683.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed buyer, address indexed tokenAddress, uint256 price, uint256 usdPrice)
func (_Contracts *ContractsFilterer) FilterAuctionEnded(opts *bind.FilterOpts, auctionId []*big.Int, buyer []common.Address, tokenAddress []common.Address) (*ContractsAuctionEndedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AuctionEnded", auctionIdRule, buyerRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAuctionEndedIterator{contract: _Contracts.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0x596165d0521c3cb4157fad2621686f086daed4663acb3d03441a92b9277f5683.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed buyer, address indexed tokenAddress, uint256 price, uint256 usdPrice)
func (_Contracts *ContractsFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *ContractsAuctionEnded, auctionId []*big.Int, buyer []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AuctionEnded", auctionIdRule, buyerRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAuctionEnded)
				if err := _Contracts.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0x596165d0521c3cb4157fad2621686f086daed4663acb3d03441a92b9277f5683.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed buyer, address indexed tokenAddress, uint256 price, uint256 usdPrice)
func (_Contracts *ContractsFilterer) ParseAuctionEnded(log types.Log) (*ContractsAuctionEnded, error) {
	event := new(ContractsAuctionEnded)
	if err := _Contracts.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the Contracts contract.
type ContractsBidPlacedIterator struct {
	Event *ContractsBidPlaced // Event containing the contract specifics and raw log

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
func (it *ContractsBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBidPlaced)
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
		it.Event = new(ContractsBidPlaced)
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
func (it *ContractsBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBidPlaced represents a BidPlaced event raised by the Contracts contract.
type ContractsBidPlaced struct {
	AuctionId    *big.Int
	Bidder       common.Address
	TokenAddress common.Address
	BidAmount    *big.Int
	BidUsdAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x2808decb743a25d04efe1bd3dc192acde3be644e2f6ad1dce5d3c46643e1c602.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, address indexed tokenAddress, uint256 bidAmount, uint256 bidUsdAmount)
func (_Contracts *ContractsFilterer) FilterBidPlaced(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address, tokenAddress []common.Address) (*ContractsBidPlacedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "BidPlaced", auctionIdRule, bidderRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractsBidPlacedIterator{contract: _Contracts.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x2808decb743a25d04efe1bd3dc192acde3be644e2f6ad1dce5d3c46643e1c602.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, address indexed tokenAddress, uint256 bidAmount, uint256 bidUsdAmount)
func (_Contracts *ContractsFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *ContractsBidPlaced, auctionId []*big.Int, bidder []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "BidPlaced", auctionIdRule, bidderRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBidPlaced)
				if err := _Contracts.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x2808decb743a25d04efe1bd3dc192acde3be644e2f6ad1dce5d3c46643e1c602.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, address indexed tokenAddress, uint256 bidAmount, uint256 bidUsdAmount)
func (_Contracts *ContractsFilterer) ParseBidPlaced(log types.Log) (*ContractsBidPlaced, error) {
	event := new(ContractsBidPlaced)
	if err := _Contracts.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsBidWithdrawnIterator is returned from FilterBidWithdrawn and is used to iterate over the raw logs and unpacked data for BidWithdrawn events raised by the Contracts contract.
type ContractsBidWithdrawnIterator struct {
	Event *ContractsBidWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractsBidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBidWithdrawn)
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
		it.Event = new(ContractsBidWithdrawn)
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
func (it *ContractsBidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBidWithdrawn represents a BidWithdrawn event raised by the Contracts contract.
type ContractsBidWithdrawn struct {
	AuctionId    *big.Int
	Bidder       common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBidWithdrawn is a free log retrieval operation binding the contract event 0x718753587878e0812c16c95d4bec2f56bc77661c1242d74a9e747463e23583f4.
//
// Solidity: event BidWithdrawn(uint256 indexed auctionId, address indexed bidder, address indexed tokenAddress, uint256 amount)
func (_Contracts *ContractsFilterer) FilterBidWithdrawn(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address, tokenAddress []common.Address) (*ContractsBidWithdrawnIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "BidWithdrawn", auctionIdRule, bidderRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractsBidWithdrawnIterator{contract: _Contracts.contract, event: "BidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBidWithdrawn is a free log subscription operation binding the contract event 0x718753587878e0812c16c95d4bec2f56bc77661c1242d74a9e747463e23583f4.
//
// Solidity: event BidWithdrawn(uint256 indexed auctionId, address indexed bidder, address indexed tokenAddress, uint256 amount)
func (_Contracts *ContractsFilterer) WatchBidWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractsBidWithdrawn, auctionId []*big.Int, bidder []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "BidWithdrawn", auctionIdRule, bidderRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBidWithdrawn)
				if err := _Contracts.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
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

// ParseBidWithdrawn is a log parse operation binding the contract event 0x718753587878e0812c16c95d4bec2f56bc77661c1242d74a9e747463e23583f4.
//
// Solidity: event BidWithdrawn(uint256 indexed auctionId, address indexed bidder, address indexed tokenAddress, uint256 amount)
func (_Contracts *ContractsFilterer) ParseBidWithdrawn(log types.Log) (*ContractsBidWithdrawn, error) {
	event := new(ContractsBidWithdrawn)
	if err := _Contracts.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsFeeRecipientUpdatedIterator is returned from FilterFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientUpdated events raised by the Contracts contract.
type ContractsFeeRecipientUpdatedIterator struct {
	Event *ContractsFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFeeRecipientUpdated)
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
		it.Event = new(ContractsFeeRecipientUpdated)
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
func (it *ContractsFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the Contracts contract.
type ContractsFeeRecipientUpdated struct {
	OldRecipient common.Address
	NewRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientUpdated is a free log retrieval operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed oldRecipient, address indexed newRecipient)
func (_Contracts *ContractsFilterer) FilterFeeRecipientUpdated(opts *bind.FilterOpts, oldRecipient []common.Address, newRecipient []common.Address) (*ContractsFeeRecipientUpdatedIterator, error) {

	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FeeRecipientUpdated", oldRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractsFeeRecipientUpdatedIterator{contract: _Contracts.contract, event: "FeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientUpdated is a free log subscription operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed oldRecipient, address indexed newRecipient)
func (_Contracts *ContractsFilterer) WatchFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *ContractsFeeRecipientUpdated, oldRecipient []common.Address, newRecipient []common.Address) (event.Subscription, error) {

	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FeeRecipientUpdated", oldRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFeeRecipientUpdated)
				if err := _Contracts.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
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

// ParseFeeRecipientUpdated is a log parse operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed oldRecipient, address indexed newRecipient)
func (_Contracts *ContractsFilterer) ParseFeeRecipientUpdated(log types.Log) (*ContractsFeeRecipientUpdated, error) {
	event := new(ContractsFeeRecipientUpdated)
	if err := _Contracts.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contracts contract.
type ContractsInitializedIterator struct {
	Event *ContractsInitialized // Event containing the contract specifics and raw log

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
func (it *ContractsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsInitialized)
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
		it.Event = new(ContractsInitialized)
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
func (it *ContractsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsInitialized represents a Initialized event raised by the Contracts contract.
type ContractsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contracts *ContractsFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractsInitializedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractsInitializedIterator{contract: _Contracts.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contracts *ContractsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractsInitialized) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsInitialized)
				if err := _Contracts.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contracts *ContractsFilterer) ParseInitialized(log types.Log) (*ContractsInitialized, error) {
	event := new(ContractsInitialized)
	if err := _Contracts.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPaymentTokenAllowedUpdatedIterator is returned from FilterPaymentTokenAllowedUpdated and is used to iterate over the raw logs and unpacked data for PaymentTokenAllowedUpdated events raised by the Contracts contract.
type ContractsPaymentTokenAllowedUpdatedIterator struct {
	Event *ContractsPaymentTokenAllowedUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsPaymentTokenAllowedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPaymentTokenAllowedUpdated)
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
		it.Event = new(ContractsPaymentTokenAllowedUpdated)
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
func (it *ContractsPaymentTokenAllowedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPaymentTokenAllowedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPaymentTokenAllowedUpdated represents a PaymentTokenAllowedUpdated event raised by the Contracts contract.
type ContractsPaymentTokenAllowedUpdated struct {
	TokenAddress common.Address
	Allowed      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenAllowedUpdated is a free log retrieval operation binding the contract event 0x359d3f0085554124a1c9fc8f53602b98df473a1d955d73500399fb2ef8b8cf18.
//
// Solidity: event PaymentTokenAllowedUpdated(address indexed tokenAddress, bool allowed)
func (_Contracts *ContractsFilterer) FilterPaymentTokenAllowedUpdated(opts *bind.FilterOpts, tokenAddress []common.Address) (*ContractsPaymentTokenAllowedUpdatedIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PaymentTokenAllowedUpdated", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPaymentTokenAllowedUpdatedIterator{contract: _Contracts.contract, event: "PaymentTokenAllowedUpdated", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenAllowedUpdated is a free log subscription operation binding the contract event 0x359d3f0085554124a1c9fc8f53602b98df473a1d955d73500399fb2ef8b8cf18.
//
// Solidity: event PaymentTokenAllowedUpdated(address indexed tokenAddress, bool allowed)
func (_Contracts *ContractsFilterer) WatchPaymentTokenAllowedUpdated(opts *bind.WatchOpts, sink chan<- *ContractsPaymentTokenAllowedUpdated, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PaymentTokenAllowedUpdated", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPaymentTokenAllowedUpdated)
				if err := _Contracts.contract.UnpackLog(event, "PaymentTokenAllowedUpdated", log); err != nil {
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

// ParsePaymentTokenAllowedUpdated is a log parse operation binding the contract event 0x359d3f0085554124a1c9fc8f53602b98df473a1d955d73500399fb2ef8b8cf18.
//
// Solidity: event PaymentTokenAllowedUpdated(address indexed tokenAddress, bool allowed)
func (_Contracts *ContractsFilterer) ParsePaymentTokenAllowedUpdated(log types.Log) (*ContractsPaymentTokenAllowedUpdated, error) {
	event := new(ContractsPaymentTokenAllowedUpdated)
	if err := _Contracts.contract.UnpackLog(event, "PaymentTokenAllowedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPlatformFeeUpdatedIterator is returned from FilterPlatformFeeUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeUpdated events raised by the Contracts contract.
type ContractsPlatformFeeUpdatedIterator struct {
	Event *ContractsPlatformFeeUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsPlatformFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPlatformFeeUpdated)
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
		it.Event = new(ContractsPlatformFeeUpdated)
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
func (it *ContractsPlatformFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPlatformFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPlatformFeeUpdated represents a PlatformFeeUpdated event raised by the Contracts contract.
type ContractsPlatformFeeUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeUpdated is a free log retrieval operation binding the contract event 0xd347e206f25a89b917fc9482f1a2d294d749baa4dc9bde7fb495ee11fe491643.
//
// Solidity: event PlatformFeeUpdated(uint256 oldFee, uint256 newFee)
func (_Contracts *ContractsFilterer) FilterPlatformFeeUpdated(opts *bind.FilterOpts) (*ContractsPlatformFeeUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsPlatformFeeUpdatedIterator{contract: _Contracts.contract, event: "PlatformFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeUpdated is a free log subscription operation binding the contract event 0xd347e206f25a89b917fc9482f1a2d294d749baa4dc9bde7fb495ee11fe491643.
//
// Solidity: event PlatformFeeUpdated(uint256 oldFee, uint256 newFee)
func (_Contracts *ContractsFilterer) WatchPlatformFeeUpdated(opts *bind.WatchOpts, sink chan<- *ContractsPlatformFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPlatformFeeUpdated)
				if err := _Contracts.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
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

// ParsePlatformFeeUpdated is a log parse operation binding the contract event 0xd347e206f25a89b917fc9482f1a2d294d749baa4dc9bde7fb495ee11fe491643.
//
// Solidity: event PlatformFeeUpdated(uint256 oldFee, uint256 newFee)
func (_Contracts *ContractsFilterer) ParsePlatformFeeUpdated(log types.Log) (*ContractsPlatformFeeUpdated, error) {
	event := new(ContractsPlatformFeeUpdated)
	if err := _Contracts.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPriceFeedConfiguredIterator is returned from FilterPriceFeedConfigured and is used to iterate over the raw logs and unpacked data for PriceFeedConfigured events raised by the Contracts contract.
type ContractsPriceFeedConfiguredIterator struct {
	Event *ContractsPriceFeedConfigured // Event containing the contract specifics and raw log

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
func (it *ContractsPriceFeedConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPriceFeedConfigured)
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
		it.Event = new(ContractsPriceFeedConfigured)
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
func (it *ContractsPriceFeedConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPriceFeedConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPriceFeedConfigured represents a PriceFeedConfigured event raised by the Contracts contract.
type ContractsPriceFeedConfigured struct {
	TokenAddress  common.Address
	Feed          common.Address
	TokenDecimals uint8
	MaxStaleness  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedConfigured is a free log retrieval operation binding the contract event 0x90f4c9edb23e8bfa67a001b49edb54575ab0758b9b09260445238c592615dc80.
//
// Solidity: event PriceFeedConfigured(address indexed tokenAddress, address indexed feed, uint8 tokenDecimals, uint256 maxStaleness)
func (_Contracts *ContractsFilterer) FilterPriceFeedConfigured(opts *bind.FilterOpts, tokenAddress []common.Address, feed []common.Address) (*ContractsPriceFeedConfiguredIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var feedRule []interface{}
	for _, feedItem := range feed {
		feedRule = append(feedRule, feedItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PriceFeedConfigured", tokenAddressRule, feedRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPriceFeedConfiguredIterator{contract: _Contracts.contract, event: "PriceFeedConfigured", logs: logs, sub: sub}, nil
}

// WatchPriceFeedConfigured is a free log subscription operation binding the contract event 0x90f4c9edb23e8bfa67a001b49edb54575ab0758b9b09260445238c592615dc80.
//
// Solidity: event PriceFeedConfigured(address indexed tokenAddress, address indexed feed, uint8 tokenDecimals, uint256 maxStaleness)
func (_Contracts *ContractsFilterer) WatchPriceFeedConfigured(opts *bind.WatchOpts, sink chan<- *ContractsPriceFeedConfigured, tokenAddress []common.Address, feed []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var feedRule []interface{}
	for _, feedItem := range feed {
		feedRule = append(feedRule, feedItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PriceFeedConfigured", tokenAddressRule, feedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPriceFeedConfigured)
				if err := _Contracts.contract.UnpackLog(event, "PriceFeedConfigured", log); err != nil {
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

// ParsePriceFeedConfigured is a log parse operation binding the contract event 0x90f4c9edb23e8bfa67a001b49edb54575ab0758b9b09260445238c592615dc80.
//
// Solidity: event PriceFeedConfigured(address indexed tokenAddress, address indexed feed, uint8 tokenDecimals, uint256 maxStaleness)
func (_Contracts *ContractsFilterer) ParsePriceFeedConfigured(log types.Log) (*ContractsPriceFeedConfigured, error) {
	event := new(ContractsPriceFeedConfigured)
	if err := _Contracts.contract.UnpackLog(event, "PriceFeedConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPriceFeedDisabledIterator is returned from FilterPriceFeedDisabled and is used to iterate over the raw logs and unpacked data for PriceFeedDisabled events raised by the Contracts contract.
type ContractsPriceFeedDisabledIterator struct {
	Event *ContractsPriceFeedDisabled // Event containing the contract specifics and raw log

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
func (it *ContractsPriceFeedDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPriceFeedDisabled)
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
		it.Event = new(ContractsPriceFeedDisabled)
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
func (it *ContractsPriceFeedDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPriceFeedDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPriceFeedDisabled represents a PriceFeedDisabled event raised by the Contracts contract.
type ContractsPriceFeedDisabled struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedDisabled is a free log retrieval operation binding the contract event 0x8a16c5098f5027c3179beb0ba478798ae5f6f1ea2253eb59e972567d36fbcc4f.
//
// Solidity: event PriceFeedDisabled(address indexed tokenAddress)
func (_Contracts *ContractsFilterer) FilterPriceFeedDisabled(opts *bind.FilterOpts, tokenAddress []common.Address) (*ContractsPriceFeedDisabledIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PriceFeedDisabled", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPriceFeedDisabledIterator{contract: _Contracts.contract, event: "PriceFeedDisabled", logs: logs, sub: sub}, nil
}

// WatchPriceFeedDisabled is a free log subscription operation binding the contract event 0x8a16c5098f5027c3179beb0ba478798ae5f6f1ea2253eb59e972567d36fbcc4f.
//
// Solidity: event PriceFeedDisabled(address indexed tokenAddress)
func (_Contracts *ContractsFilterer) WatchPriceFeedDisabled(opts *bind.WatchOpts, sink chan<- *ContractsPriceFeedDisabled, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PriceFeedDisabled", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPriceFeedDisabled)
				if err := _Contracts.contract.UnpackLog(event, "PriceFeedDisabled", log); err != nil {
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

// ParsePriceFeedDisabled is a log parse operation binding the contract event 0x8a16c5098f5027c3179beb0ba478798ae5f6f1ea2253eb59e972567d36fbcc4f.
//
// Solidity: event PriceFeedDisabled(address indexed tokenAddress)
func (_Contracts *ContractsFilterer) ParsePriceFeedDisabled(log types.Log) (*ContractsPriceFeedDisabled, error) {
	event := new(ContractsPriceFeedDisabled)
	if err := _Contracts.contract.UnpackLog(event, "PriceFeedDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Contracts contract.
type ContractsUpgradedIterator struct {
	Event *ContractsUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUpgraded)
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
		it.Event = new(ContractsUpgraded)
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
func (it *ContractsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUpgraded represents a Upgraded event raised by the Contracts contract.
type ContractsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contracts *ContractsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractsUpgradedIterator{contract: _Contracts.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contracts *ContractsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUpgraded)
				if err := _Contracts.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contracts *ContractsFilterer) ParseUpgraded(log types.Log) (*ContractsUpgraded, error) {
	event := new(ContractsUpgraded)
	if err := _Contracts.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
