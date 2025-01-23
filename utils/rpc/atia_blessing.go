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
	_ = abi.ConvertType
)

// IAtiaShrinePrayerCountInfo is an auto generated low-level Go binding around an user-defined struct.
type IAtiaShrinePrayerCountInfo struct {
	SyncedDayIndex uint64
	CountPerDay    uint64
}

// IAtiaShrineRestorationCost is an auto generated low-level Go binding around an user-defined struct.
type IAtiaShrineRestorationCost struct {
	LowerMilestone uint32
	CostInUsd      *big.Int
}

// PythConverter is an auto generated low-level Go binding around an user-defined struct.
type PythConverter struct {
	TokenInDecimal   uint8
	TokenOutDecimal  uint8
	Pyth             common.Address
	PriceId          [32]byte
	MaxAcceptableAge *big.Int
}

// AtiaBlessingMetaData contains all meta data concerning the AtiaBlessing contract.
var AtiaBlessingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrActivateBeforeStartingBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAlreadyActivatedStreak\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerNotDelegated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"expo1\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"expo2\",\"type\":\"int32\"},{\"internalType\":\"int64\",\"name\":\"price1\",\"type\":\"int64\"}],\"name\":\"ErrComputedPriceTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrExceedDailyPrayThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"}],\"name\":\"ErrExponentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrFreezeTimeEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInsufficientPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidActiveStreakLifeTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidOrderOfRestorationCosts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidResetDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoLostStreak\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrStartNewStreakByDelegatee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroLostStreakCount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"streakLifeTime\",\"type\":\"uint256\"}],\"name\":\"ActiveStreakLifeTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dailyPrayerThreshold\",\"type\":\"uint256\"}],\"name\":\"DailyPrayerThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFreezePeriod\",\"type\":\"uint256\"}],\"name\":\"FreezePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPaymentToken\",\"type\":\"address\"}],\"name\":\"PaymentTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prayer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"syncedDayIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"countPerDay\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIAtiaShrine.PrayerCountInfo\",\"name\":\"prayerCountInfo\",\"type\":\"tuple\"}],\"name\":\"PrayerCountSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tokenInDecimal\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenOutDecimal\",\"type\":\"uint8\"},{\"internalType\":\"contractIPyth\",\"name\":\"pyth\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"priceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxAcceptableAge\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPythConverter\",\"name\":\"priceFeedData\",\"type\":\"tuple\"}],\"name\":\"PriceFeedDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newResetDuration\",\"type\":\"uint256\"}],\"name\":\"ResetDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"lowerMilestone\",\"type\":\"uint32\"},{\"internalType\":\"uint80\",\"name\":\"costInUsd\",\"type\":\"uint80\"}],\"indexed\":false,\"internalType\":\"structIAtiaShrine.RestorationCost[]\",\"name\":\"newRestorationCosts\",\"type\":\"tuple[]\"}],\"name\":\"RestorationCostUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"restoredStreakCount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costInUsd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costInPaymentToken\",\"type\":\"uint256\"}],\"name\":\"StreakRestored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newLastUpdated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"longestStreakCount\",\"type\":\"uint256\"}],\"name\":\"StreakUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasuryUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"activateStreak\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"streaks\",\"type\":\"uint256[]\"}],\"name\":\"forceRestoreStreak\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getActivationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isLostStreak\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasPrayedToday\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveStreakLifeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRestorationCosts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"lowerMilestone\",\"type\":\"uint32\"},{\"internalType\":\"uint80\",\"name\":\"costInUsd\",\"type\":\"uint80\"}],\"internalType\":\"structIAtiaShrine.RestorationCost[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prayer\",\"type\":\"address\"}],\"name\":\"getDailyPrayerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDailyPrayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFreezePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaymentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceFeedData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tokenInDecimal\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenOutDecimal\",\"type\":\"uint8\"},{\"internalType\":\"contractIPyth\",\"name\":\"pyth\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"priceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxAcceptableAge\",\"type\":\"uint256\"}],\"internalType\":\"structPythConverter\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResetDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lostStreakCount\",\"type\":\"uint256\"}],\"name\":\"getRestorationCosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"costInUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costInPaymentToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStartedAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getStreak\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentStreakCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastActivated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longestStreakCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lostStreakCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streakLifeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startedAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streakLifeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resetTime\",\"type\":\"uint256\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"freezePeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"lowerMilestone\",\"type\":\"uint32\"},{\"internalType\":\"uint80\",\"name\":\"costInUsd\",\"type\":\"uint80\"}],\"internalType\":\"structIAtiaShrine.RestorationCost[]\",\"name\":\"restorationCosts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tokenInDecimal\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenOutDecimal\",\"type\":\"uint8\"},{\"internalType\":\"contractIPyth\",\"name\":\"pyth\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"priceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxAcceptableAge\",\"type\":\"uint256\"}],\"internalType\":\"structPythConverter\",\"name\":\"priceFeedData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"dailyPrayerThreshold\",\"type\":\"uint256\"}],\"name\":\"initializeV3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isRestorable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"restoreStreak\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dailyPrayerThreshold\",\"type\":\"uint256\"}],\"name\":\"setDailyPrayerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFreezePeriod\",\"type\":\"uint256\"}],\"name\":\"setFreezePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"setPaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tokenInDecimal\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenOutDecimal\",\"type\":\"uint8\"},{\"internalType\":\"contractIPyth\",\"name\":\"pyth\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"priceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxAcceptableAge\",\"type\":\"uint256\"}],\"internalType\":\"structPythConverter\",\"name\":\"priceFeedData\",\"type\":\"tuple\"}],\"name\":\"setPriceFeedData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newResetDuration\",\"type\":\"uint256\"}],\"name\":\"setResetDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"lowerMilestone\",\"type\":\"uint32\"},{\"internalType\":\"uint80\",\"name\":\"costInUsd\",\"type\":\"uint80\"}],\"internalType\":\"structIAtiaShrine.RestorationCost[]\",\"name\":\"restorationCosts\",\"type\":\"tuple[]\"}],\"name\":\"setRestorationCosts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AtiaBlessingABI is the input ABI used to generate the binding from.
// Deprecated: Use AtiaBlessingMetaData.ABI instead.
var AtiaBlessingABI = AtiaBlessingMetaData.ABI

// AtiaBlessing is an auto generated Go binding around an Ethereum contract.
type AtiaBlessing struct {
	AtiaBlessingCaller     // Read-only binding to the contract
	AtiaBlessingTransactor // Write-only binding to the contract
	AtiaBlessingFilterer   // Log filterer for contract events
}

// AtiaBlessingCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtiaBlessingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtiaBlessingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtiaBlessingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtiaBlessingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtiaBlessingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtiaBlessingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtiaBlessingSession struct {
	Contract     *AtiaBlessing     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtiaBlessingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtiaBlessingCallerSession struct {
	Contract *AtiaBlessingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AtiaBlessingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtiaBlessingTransactorSession struct {
	Contract     *AtiaBlessingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AtiaBlessingRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtiaBlessingRaw struct {
	Contract *AtiaBlessing // Generic contract binding to access the raw methods on
}

// AtiaBlessingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtiaBlessingCallerRaw struct {
	Contract *AtiaBlessingCaller // Generic read-only contract binding to access the raw methods on
}

// AtiaBlessingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtiaBlessingTransactorRaw struct {
	Contract *AtiaBlessingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtiaBlessing creates a new instance of AtiaBlessing, bound to a specific deployed contract.
func NewAtiaBlessing(address common.Address, backend bind.ContractBackend) (*AtiaBlessing, error) {
	contract, err := bindAtiaBlessing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AtiaBlessing{AtiaBlessingCaller: AtiaBlessingCaller{contract: contract}, AtiaBlessingTransactor: AtiaBlessingTransactor{contract: contract}, AtiaBlessingFilterer: AtiaBlessingFilterer{contract: contract}}, nil
}

// NewAtiaBlessingCaller creates a new read-only instance of AtiaBlessing, bound to a specific deployed contract.
func NewAtiaBlessingCaller(address common.Address, caller bind.ContractCaller) (*AtiaBlessingCaller, error) {
	contract, err := bindAtiaBlessing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingCaller{contract: contract}, nil
}

// NewAtiaBlessingTransactor creates a new write-only instance of AtiaBlessing, bound to a specific deployed contract.
func NewAtiaBlessingTransactor(address common.Address, transactor bind.ContractTransactor) (*AtiaBlessingTransactor, error) {
	contract, err := bindAtiaBlessing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingTransactor{contract: contract}, nil
}

// NewAtiaBlessingFilterer creates a new log filterer instance of AtiaBlessing, bound to a specific deployed contract.
func NewAtiaBlessingFilterer(address common.Address, filterer bind.ContractFilterer) (*AtiaBlessingFilterer, error) {
	contract, err := bindAtiaBlessing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingFilterer{contract: contract}, nil
}

// bindAtiaBlessing binds a generic wrapper to an already deployed contract.
func bindAtiaBlessing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AtiaBlessingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtiaBlessing *AtiaBlessingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtiaBlessing.Contract.AtiaBlessingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtiaBlessing *AtiaBlessingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.AtiaBlessingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtiaBlessing *AtiaBlessingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.AtiaBlessingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtiaBlessing *AtiaBlessingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtiaBlessing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtiaBlessing *AtiaBlessingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtiaBlessing *AtiaBlessingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AtiaBlessing *AtiaBlessingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AtiaBlessing *AtiaBlessingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AtiaBlessing.Contract.DEFAULTADMINROLE(&_AtiaBlessing.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AtiaBlessing *AtiaBlessingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AtiaBlessing.Contract.DEFAULTADMINROLE(&_AtiaBlessing.CallOpts)
}

// GetActivationStatus is a free data retrieval call binding the contract method 0x6d16c15b.
//
// Solidity: function getActivationStatus(address user) view returns(bool isLostStreak, bool hasPrayedToday)
func (_AtiaBlessing *AtiaBlessingCaller) GetActivationStatus(opts *bind.CallOpts, user common.Address) (struct {
	IsLostStreak   bool
	HasPrayedToday bool
}, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getActivationStatus", user)

	outstruct := new(struct {
		IsLostStreak   bool
		HasPrayedToday bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsLostStreak = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.HasPrayedToday = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetActivationStatus is a free data retrieval call binding the contract method 0x6d16c15b.
//
// Solidity: function getActivationStatus(address user) view returns(bool isLostStreak, bool hasPrayedToday)
func (_AtiaBlessing *AtiaBlessingSession) GetActivationStatus(user common.Address) (struct {
	IsLostStreak   bool
	HasPrayedToday bool
}, error) {
	return _AtiaBlessing.Contract.GetActivationStatus(&_AtiaBlessing.CallOpts, user)
}

// GetActivationStatus is a free data retrieval call binding the contract method 0x6d16c15b.
//
// Solidity: function getActivationStatus(address user) view returns(bool isLostStreak, bool hasPrayedToday)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetActivationStatus(user common.Address) (struct {
	IsLostStreak   bool
	HasPrayedToday bool
}, error) {
	return _AtiaBlessing.Contract.GetActivationStatus(&_AtiaBlessing.CallOpts, user)
}

// GetActiveStreakLifeTime is a free data retrieval call binding the contract method 0x01a1476c.
//
// Solidity: function getActiveStreakLifeTime() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCaller) GetActiveStreakLifeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getActiveStreakLifeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveStreakLifeTime is a free data retrieval call binding the contract method 0x01a1476c.
//
// Solidity: function getActiveStreakLifeTime() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingSession) GetActiveStreakLifeTime() (*big.Int, error) {
	return _AtiaBlessing.Contract.GetActiveStreakLifeTime(&_AtiaBlessing.CallOpts)
}

// GetActiveStreakLifeTime is a free data retrieval call binding the contract method 0x01a1476c.
//
// Solidity: function getActiveStreakLifeTime() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetActiveStreakLifeTime() (*big.Int, error) {
	return _AtiaBlessing.Contract.GetActiveStreakLifeTime(&_AtiaBlessing.CallOpts)
}

// GetAllRestorationCosts is a free data retrieval call binding the contract method 0x36cee514.
//
// Solidity: function getAllRestorationCosts() view returns((uint32,uint80)[])
func (_AtiaBlessing *AtiaBlessingCaller) GetAllRestorationCosts(opts *bind.CallOpts) ([]IAtiaShrineRestorationCost, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getAllRestorationCosts")

	if err != nil {
		return *new([]IAtiaShrineRestorationCost), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAtiaShrineRestorationCost)).(*[]IAtiaShrineRestorationCost)

	return out0, err

}

// GetAllRestorationCosts is a free data retrieval call binding the contract method 0x36cee514.
//
// Solidity: function getAllRestorationCosts() view returns((uint32,uint80)[])
func (_AtiaBlessing *AtiaBlessingSession) GetAllRestorationCosts() ([]IAtiaShrineRestorationCost, error) {
	return _AtiaBlessing.Contract.GetAllRestorationCosts(&_AtiaBlessing.CallOpts)
}

// GetAllRestorationCosts is a free data retrieval call binding the contract method 0x36cee514.
//
// Solidity: function getAllRestorationCosts() view returns((uint32,uint80)[])
func (_AtiaBlessing *AtiaBlessingCallerSession) GetAllRestorationCosts() ([]IAtiaShrineRestorationCost, error) {
	return _AtiaBlessing.Contract.GetAllRestorationCosts(&_AtiaBlessing.CallOpts)
}

// GetDailyPrayerCount is a free data retrieval call binding the contract method 0x89464619.
//
// Solidity: function getDailyPrayerCount(address prayer) view returns(uint256, uint256)
func (_AtiaBlessing *AtiaBlessingCaller) GetDailyPrayerCount(opts *bind.CallOpts, prayer common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getDailyPrayerCount", prayer)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDailyPrayerCount is a free data retrieval call binding the contract method 0x89464619.
//
// Solidity: function getDailyPrayerCount(address prayer) view returns(uint256, uint256)
func (_AtiaBlessing *AtiaBlessingSession) GetDailyPrayerCount(prayer common.Address) (*big.Int, *big.Int, error) {
	return _AtiaBlessing.Contract.GetDailyPrayerCount(&_AtiaBlessing.CallOpts, prayer)
}

// GetDailyPrayerCount is a free data retrieval call binding the contract method 0x89464619.
//
// Solidity: function getDailyPrayerCount(address prayer) view returns(uint256, uint256)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetDailyPrayerCount(prayer common.Address) (*big.Int, *big.Int, error) {
	return _AtiaBlessing.Contract.GetDailyPrayerCount(&_AtiaBlessing.CallOpts, prayer)
}

// GetDailyPrayerThreshold is a free data retrieval call binding the contract method 0xeb7113be.
//
// Solidity: function getDailyPrayerThreshold() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCaller) GetDailyPrayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getDailyPrayerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDailyPrayerThreshold is a free data retrieval call binding the contract method 0xeb7113be.
//
// Solidity: function getDailyPrayerThreshold() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingSession) GetDailyPrayerThreshold() (*big.Int, error) {
	return _AtiaBlessing.Contract.GetDailyPrayerThreshold(&_AtiaBlessing.CallOpts)
}

// GetDailyPrayerThreshold is a free data retrieval call binding the contract method 0xeb7113be.
//
// Solidity: function getDailyPrayerThreshold() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetDailyPrayerThreshold() (*big.Int, error) {
	return _AtiaBlessing.Contract.GetDailyPrayerThreshold(&_AtiaBlessing.CallOpts)
}

// GetFreezePeriod is a free data retrieval call binding the contract method 0x9fe4f9d6.
//
// Solidity: function getFreezePeriod() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCaller) GetFreezePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getFreezePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFreezePeriod is a free data retrieval call binding the contract method 0x9fe4f9d6.
//
// Solidity: function getFreezePeriod() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingSession) GetFreezePeriod() (*big.Int, error) {
	return _AtiaBlessing.Contract.GetFreezePeriod(&_AtiaBlessing.CallOpts)
}

// GetFreezePeriod is a free data retrieval call binding the contract method 0x9fe4f9d6.
//
// Solidity: function getFreezePeriod() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetFreezePeriod() (*big.Int, error) {
	return _AtiaBlessing.Contract.GetFreezePeriod(&_AtiaBlessing.CallOpts)
}

// GetPaymentToken is a free data retrieval call binding the contract method 0xd41c3a65.
//
// Solidity: function getPaymentToken() view returns(address)
func (_AtiaBlessing *AtiaBlessingCaller) GetPaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getPaymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPaymentToken is a free data retrieval call binding the contract method 0xd41c3a65.
//
// Solidity: function getPaymentToken() view returns(address)
func (_AtiaBlessing *AtiaBlessingSession) GetPaymentToken() (common.Address, error) {
	return _AtiaBlessing.Contract.GetPaymentToken(&_AtiaBlessing.CallOpts)
}

// GetPaymentToken is a free data retrieval call binding the contract method 0xd41c3a65.
//
// Solidity: function getPaymentToken() view returns(address)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetPaymentToken() (common.Address, error) {
	return _AtiaBlessing.Contract.GetPaymentToken(&_AtiaBlessing.CallOpts)
}

// GetPriceFeedData is a free data retrieval call binding the contract method 0x0fe50f43.
//
// Solidity: function getPriceFeedData() view returns((uint8,uint8,address,bytes32,uint256))
func (_AtiaBlessing *AtiaBlessingCaller) GetPriceFeedData(opts *bind.CallOpts) (PythConverter, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getPriceFeedData")

	if err != nil {
		return *new(PythConverter), err
	}

	out0 := *abi.ConvertType(out[0], new(PythConverter)).(*PythConverter)

	return out0, err

}

// GetPriceFeedData is a free data retrieval call binding the contract method 0x0fe50f43.
//
// Solidity: function getPriceFeedData() view returns((uint8,uint8,address,bytes32,uint256))
func (_AtiaBlessing *AtiaBlessingSession) GetPriceFeedData() (PythConverter, error) {
	return _AtiaBlessing.Contract.GetPriceFeedData(&_AtiaBlessing.CallOpts)
}

// GetPriceFeedData is a free data retrieval call binding the contract method 0x0fe50f43.
//
// Solidity: function getPriceFeedData() view returns((uint8,uint8,address,bytes32,uint256))
func (_AtiaBlessing *AtiaBlessingCallerSession) GetPriceFeedData() (PythConverter, error) {
	return _AtiaBlessing.Contract.GetPriceFeedData(&_AtiaBlessing.CallOpts)
}

// GetResetDuration is a free data retrieval call binding the contract method 0x1614252f.
//
// Solidity: function getResetDuration() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCaller) GetResetDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getResetDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResetDuration is a free data retrieval call binding the contract method 0x1614252f.
//
// Solidity: function getResetDuration() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingSession) GetResetDuration() (*big.Int, error) {
	return _AtiaBlessing.Contract.GetResetDuration(&_AtiaBlessing.CallOpts)
}

// GetResetDuration is a free data retrieval call binding the contract method 0x1614252f.
//
// Solidity: function getResetDuration() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetResetDuration() (*big.Int, error) {
	return _AtiaBlessing.Contract.GetResetDuration(&_AtiaBlessing.CallOpts)
}

// GetRestorationCosts is a free data retrieval call binding the contract method 0xfc44e00b.
//
// Solidity: function getRestorationCosts(uint256 lostStreakCount) view returns(uint256 costInUsd, uint256 costInPaymentToken)
func (_AtiaBlessing *AtiaBlessingCaller) GetRestorationCosts(opts *bind.CallOpts, lostStreakCount *big.Int) (struct {
	CostInUsd          *big.Int
	CostInPaymentToken *big.Int
}, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getRestorationCosts", lostStreakCount)

	outstruct := new(struct {
		CostInUsd          *big.Int
		CostInPaymentToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CostInUsd = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CostInPaymentToken = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRestorationCosts is a free data retrieval call binding the contract method 0xfc44e00b.
//
// Solidity: function getRestorationCosts(uint256 lostStreakCount) view returns(uint256 costInUsd, uint256 costInPaymentToken)
func (_AtiaBlessing *AtiaBlessingSession) GetRestorationCosts(lostStreakCount *big.Int) (struct {
	CostInUsd          *big.Int
	CostInPaymentToken *big.Int
}, error) {
	return _AtiaBlessing.Contract.GetRestorationCosts(&_AtiaBlessing.CallOpts, lostStreakCount)
}

// GetRestorationCosts is a free data retrieval call binding the contract method 0xfc44e00b.
//
// Solidity: function getRestorationCosts(uint256 lostStreakCount) view returns(uint256 costInUsd, uint256 costInPaymentToken)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetRestorationCosts(lostStreakCount *big.Int) (struct {
	CostInUsd          *big.Int
	CostInPaymentToken *big.Int
}, error) {
	return _AtiaBlessing.Contract.GetRestorationCosts(&_AtiaBlessing.CallOpts, lostStreakCount)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AtiaBlessing *AtiaBlessingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AtiaBlessing *AtiaBlessingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AtiaBlessing.Contract.GetRoleAdmin(&_AtiaBlessing.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AtiaBlessing.Contract.GetRoleAdmin(&_AtiaBlessing.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AtiaBlessing *AtiaBlessingCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AtiaBlessing *AtiaBlessingSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AtiaBlessing.Contract.GetRoleMember(&_AtiaBlessing.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AtiaBlessing.Contract.GetRoleMember(&_AtiaBlessing.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AtiaBlessing *AtiaBlessingSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AtiaBlessing.Contract.GetRoleMemberCount(&_AtiaBlessing.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AtiaBlessing.Contract.GetRoleMemberCount(&_AtiaBlessing.CallOpts, role)
}

// GetStartedAtBlock is a free data retrieval call binding the contract method 0xdcedeb64.
//
// Solidity: function getStartedAtBlock() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCaller) GetStartedAtBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getStartedAtBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStartedAtBlock is a free data retrieval call binding the contract method 0xdcedeb64.
//
// Solidity: function getStartedAtBlock() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingSession) GetStartedAtBlock() (*big.Int, error) {
	return _AtiaBlessing.Contract.GetStartedAtBlock(&_AtiaBlessing.CallOpts)
}

// GetStartedAtBlock is a free data retrieval call binding the contract method 0xdcedeb64.
//
// Solidity: function getStartedAtBlock() view returns(uint256)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetStartedAtBlock() (*big.Int, error) {
	return _AtiaBlessing.Contract.GetStartedAtBlock(&_AtiaBlessing.CallOpts)
}

// GetStreak is a free data retrieval call binding the contract method 0x5eeadb0d.
//
// Solidity: function getStreak(address user) view returns(uint256 currentStreakCount, uint256 lastActivated, uint256 longestStreakCount, uint256 lostStreakCount)
func (_AtiaBlessing *AtiaBlessingCaller) GetStreak(opts *bind.CallOpts, user common.Address) (struct {
	CurrentStreakCount *big.Int
	LastActivated      *big.Int
	LongestStreakCount *big.Int
	LostStreakCount    *big.Int
}, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getStreak", user)

	outstruct := new(struct {
		CurrentStreakCount *big.Int
		LastActivated      *big.Int
		LongestStreakCount *big.Int
		LostStreakCount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentStreakCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastActivated = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LongestStreakCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LostStreakCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStreak is a free data retrieval call binding the contract method 0x5eeadb0d.
//
// Solidity: function getStreak(address user) view returns(uint256 currentStreakCount, uint256 lastActivated, uint256 longestStreakCount, uint256 lostStreakCount)
func (_AtiaBlessing *AtiaBlessingSession) GetStreak(user common.Address) (struct {
	CurrentStreakCount *big.Int
	LastActivated      *big.Int
	LongestStreakCount *big.Int
	LostStreakCount    *big.Int
}, error) {
	return _AtiaBlessing.Contract.GetStreak(&_AtiaBlessing.CallOpts, user)
}

// GetStreak is a free data retrieval call binding the contract method 0x5eeadb0d.
//
// Solidity: function getStreak(address user) view returns(uint256 currentStreakCount, uint256 lastActivated, uint256 longestStreakCount, uint256 lostStreakCount)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetStreak(user common.Address) (struct {
	CurrentStreakCount *big.Int
	LastActivated      *big.Int
	LongestStreakCount *big.Int
	LostStreakCount    *big.Int
}, error) {
	return _AtiaBlessing.Contract.GetStreak(&_AtiaBlessing.CallOpts, user)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_AtiaBlessing *AtiaBlessingCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_AtiaBlessing *AtiaBlessingSession) GetTreasury() (common.Address, error) {
	return _AtiaBlessing.Contract.GetTreasury(&_AtiaBlessing.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_AtiaBlessing *AtiaBlessingCallerSession) GetTreasury() (common.Address, error) {
	return _AtiaBlessing.Contract.GetTreasury(&_AtiaBlessing.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AtiaBlessing *AtiaBlessingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AtiaBlessing *AtiaBlessingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AtiaBlessing.Contract.HasRole(&_AtiaBlessing.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AtiaBlessing *AtiaBlessingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AtiaBlessing.Contract.HasRole(&_AtiaBlessing.CallOpts, role, account)
}

// IsRestorable is a free data retrieval call binding the contract method 0x333ebe12.
//
// Solidity: function isRestorable(address user) view returns(bool)
func (_AtiaBlessing *AtiaBlessingCaller) IsRestorable(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "isRestorable", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRestorable is a free data retrieval call binding the contract method 0x333ebe12.
//
// Solidity: function isRestorable(address user) view returns(bool)
func (_AtiaBlessing *AtiaBlessingSession) IsRestorable(user common.Address) (bool, error) {
	return _AtiaBlessing.Contract.IsRestorable(&_AtiaBlessing.CallOpts, user)
}

// IsRestorable is a free data retrieval call binding the contract method 0x333ebe12.
//
// Solidity: function isRestorable(address user) view returns(bool)
func (_AtiaBlessing *AtiaBlessingCallerSession) IsRestorable(user common.Address) (bool, error) {
	return _AtiaBlessing.Contract.IsRestorable(&_AtiaBlessing.CallOpts, user)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AtiaBlessing *AtiaBlessingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AtiaBlessing.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AtiaBlessing *AtiaBlessingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AtiaBlessing.Contract.SupportsInterface(&_AtiaBlessing.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AtiaBlessing *AtiaBlessingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AtiaBlessing.Contract.SupportsInterface(&_AtiaBlessing.CallOpts, interfaceId)
}

// ActivateStreak is a paid mutator transaction binding the contract method 0x31711309.
//
// Solidity: function activateStreak(address to) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) ActivateStreak(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "activateStreak", to)
}

// ActivateStreak is a paid mutator transaction binding the contract method 0x31711309.
//
// Solidity: function activateStreak(address to) returns()
func (_AtiaBlessing *AtiaBlessingSession) ActivateStreak(to common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.ActivateStreak(&_AtiaBlessing.TransactOpts, to)
}

// ActivateStreak is a paid mutator transaction binding the contract method 0x31711309.
//
// Solidity: function activateStreak(address to) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) ActivateStreak(to common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.ActivateStreak(&_AtiaBlessing.TransactOpts, to)
}

// ForceRestoreStreak is a paid mutator transaction binding the contract method 0x199fef65.
//
// Solidity: function forceRestoreStreak(address[] addresses, uint256[] streaks) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) ForceRestoreStreak(opts *bind.TransactOpts, addresses []common.Address, streaks []*big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "forceRestoreStreak", addresses, streaks)
}

// ForceRestoreStreak is a paid mutator transaction binding the contract method 0x199fef65.
//
// Solidity: function forceRestoreStreak(address[] addresses, uint256[] streaks) returns()
func (_AtiaBlessing *AtiaBlessingSession) ForceRestoreStreak(addresses []common.Address, streaks []*big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.ForceRestoreStreak(&_AtiaBlessing.TransactOpts, addresses, streaks)
}

// ForceRestoreStreak is a paid mutator transaction binding the contract method 0x199fef65.
//
// Solidity: function forceRestoreStreak(address[] addresses, uint256[] streaks) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) ForceRestoreStreak(addresses []common.Address, streaks []*big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.ForceRestoreStreak(&_AtiaBlessing.TransactOpts, addresses, streaks)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AtiaBlessing *AtiaBlessingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.GrantRole(&_AtiaBlessing.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.GrantRole(&_AtiaBlessing.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6ab36f2.
//
// Solidity: function initialize(uint256 streakLifeTime, uint256 startedAtBlock, address admin) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) Initialize(opts *bind.TransactOpts, streakLifeTime *big.Int, startedAtBlock *big.Int, admin common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "initialize", streakLifeTime, startedAtBlock, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6ab36f2.
//
// Solidity: function initialize(uint256 streakLifeTime, uint256 startedAtBlock, address admin) returns()
func (_AtiaBlessing *AtiaBlessingSession) Initialize(streakLifeTime *big.Int, startedAtBlock *big.Int, admin common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.Initialize(&_AtiaBlessing.TransactOpts, streakLifeTime, startedAtBlock, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6ab36f2.
//
// Solidity: function initialize(uint256 streakLifeTime, uint256 startedAtBlock, address admin) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) Initialize(streakLifeTime *big.Int, startedAtBlock *big.Int, admin common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.Initialize(&_AtiaBlessing.TransactOpts, streakLifeTime, startedAtBlock, admin)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xf1056766.
//
// Solidity: function initializeV2(uint256 streakLifeTime, uint256 resetTime) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) InitializeV2(opts *bind.TransactOpts, streakLifeTime *big.Int, resetTime *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "initializeV2", streakLifeTime, resetTime)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xf1056766.
//
// Solidity: function initializeV2(uint256 streakLifeTime, uint256 resetTime) returns()
func (_AtiaBlessing *AtiaBlessingSession) InitializeV2(streakLifeTime *big.Int, resetTime *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.InitializeV2(&_AtiaBlessing.TransactOpts, streakLifeTime, resetTime)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0xf1056766.
//
// Solidity: function initializeV2(uint256 streakLifeTime, uint256 resetTime) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) InitializeV2(streakLifeTime *big.Int, resetTime *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.InitializeV2(&_AtiaBlessing.TransactOpts, streakLifeTime, resetTime)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x25a853fc.
//
// Solidity: function initializeV3(uint256 freezePeriod, address paymentToken, address treasury, (uint32,uint80)[] restorationCosts, (uint8,uint8,address,bytes32,uint256) priceFeedData, uint256 dailyPrayerThreshold) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) InitializeV3(opts *bind.TransactOpts, freezePeriod *big.Int, paymentToken common.Address, treasury common.Address, restorationCosts []IAtiaShrineRestorationCost, priceFeedData PythConverter, dailyPrayerThreshold *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "initializeV3", freezePeriod, paymentToken, treasury, restorationCosts, priceFeedData, dailyPrayerThreshold)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x25a853fc.
//
// Solidity: function initializeV3(uint256 freezePeriod, address paymentToken, address treasury, (uint32,uint80)[] restorationCosts, (uint8,uint8,address,bytes32,uint256) priceFeedData, uint256 dailyPrayerThreshold) returns()
func (_AtiaBlessing *AtiaBlessingSession) InitializeV3(freezePeriod *big.Int, paymentToken common.Address, treasury common.Address, restorationCosts []IAtiaShrineRestorationCost, priceFeedData PythConverter, dailyPrayerThreshold *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.InitializeV3(&_AtiaBlessing.TransactOpts, freezePeriod, paymentToken, treasury, restorationCosts, priceFeedData, dailyPrayerThreshold)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x25a853fc.
//
// Solidity: function initializeV3(uint256 freezePeriod, address paymentToken, address treasury, (uint32,uint80)[] restorationCosts, (uint8,uint8,address,bytes32,uint256) priceFeedData, uint256 dailyPrayerThreshold) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) InitializeV3(freezePeriod *big.Int, paymentToken common.Address, treasury common.Address, restorationCosts []IAtiaShrineRestorationCost, priceFeedData PythConverter, dailyPrayerThreshold *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.InitializeV3(&_AtiaBlessing.TransactOpts, freezePeriod, paymentToken, treasury, restorationCosts, priceFeedData, dailyPrayerThreshold)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AtiaBlessing *AtiaBlessingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.RenounceRole(&_AtiaBlessing.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.RenounceRole(&_AtiaBlessing.TransactOpts, role, account)
}

// RestoreStreak is a paid mutator transaction binding the contract method 0x43afef80.
//
// Solidity: function restoreStreak(address to, uint256 maxAmountIn) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) RestoreStreak(opts *bind.TransactOpts, to common.Address, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "restoreStreak", to, maxAmountIn)
}

// RestoreStreak is a paid mutator transaction binding the contract method 0x43afef80.
//
// Solidity: function restoreStreak(address to, uint256 maxAmountIn) returns()
func (_AtiaBlessing *AtiaBlessingSession) RestoreStreak(to common.Address, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.RestoreStreak(&_AtiaBlessing.TransactOpts, to, maxAmountIn)
}

// RestoreStreak is a paid mutator transaction binding the contract method 0x43afef80.
//
// Solidity: function restoreStreak(address to, uint256 maxAmountIn) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) RestoreStreak(to common.Address, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.RestoreStreak(&_AtiaBlessing.TransactOpts, to, maxAmountIn)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AtiaBlessing *AtiaBlessingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.RevokeRole(&_AtiaBlessing.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.RevokeRole(&_AtiaBlessing.TransactOpts, role, account)
}

// SetDailyPrayerThreshold is a paid mutator transaction binding the contract method 0x4d77dbc5.
//
// Solidity: function setDailyPrayerThreshold(uint256 dailyPrayerThreshold) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) SetDailyPrayerThreshold(opts *bind.TransactOpts, dailyPrayerThreshold *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "setDailyPrayerThreshold", dailyPrayerThreshold)
}

// SetDailyPrayerThreshold is a paid mutator transaction binding the contract method 0x4d77dbc5.
//
// Solidity: function setDailyPrayerThreshold(uint256 dailyPrayerThreshold) returns()
func (_AtiaBlessing *AtiaBlessingSession) SetDailyPrayerThreshold(dailyPrayerThreshold *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetDailyPrayerThreshold(&_AtiaBlessing.TransactOpts, dailyPrayerThreshold)
}

// SetDailyPrayerThreshold is a paid mutator transaction binding the contract method 0x4d77dbc5.
//
// Solidity: function setDailyPrayerThreshold(uint256 dailyPrayerThreshold) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) SetDailyPrayerThreshold(dailyPrayerThreshold *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetDailyPrayerThreshold(&_AtiaBlessing.TransactOpts, dailyPrayerThreshold)
}

// SetFreezePeriod is a paid mutator transaction binding the contract method 0x57120165.
//
// Solidity: function setFreezePeriod(uint256 newFreezePeriod) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) SetFreezePeriod(opts *bind.TransactOpts, newFreezePeriod *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "setFreezePeriod", newFreezePeriod)
}

// SetFreezePeriod is a paid mutator transaction binding the contract method 0x57120165.
//
// Solidity: function setFreezePeriod(uint256 newFreezePeriod) returns()
func (_AtiaBlessing *AtiaBlessingSession) SetFreezePeriod(newFreezePeriod *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetFreezePeriod(&_AtiaBlessing.TransactOpts, newFreezePeriod)
}

// SetFreezePeriod is a paid mutator transaction binding the contract method 0x57120165.
//
// Solidity: function setFreezePeriod(uint256 newFreezePeriod) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) SetFreezePeriod(newFreezePeriod *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetFreezePeriod(&_AtiaBlessing.TransactOpts, newFreezePeriod)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address paymentToken) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) SetPaymentToken(opts *bind.TransactOpts, paymentToken common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "setPaymentToken", paymentToken)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address paymentToken) returns()
func (_AtiaBlessing *AtiaBlessingSession) SetPaymentToken(paymentToken common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetPaymentToken(&_AtiaBlessing.TransactOpts, paymentToken)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address paymentToken) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) SetPaymentToken(paymentToken common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetPaymentToken(&_AtiaBlessing.TransactOpts, paymentToken)
}

// SetPriceFeedData is a paid mutator transaction binding the contract method 0x1b308773.
//
// Solidity: function setPriceFeedData((uint8,uint8,address,bytes32,uint256) priceFeedData) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) SetPriceFeedData(opts *bind.TransactOpts, priceFeedData PythConverter) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "setPriceFeedData", priceFeedData)
}

// SetPriceFeedData is a paid mutator transaction binding the contract method 0x1b308773.
//
// Solidity: function setPriceFeedData((uint8,uint8,address,bytes32,uint256) priceFeedData) returns()
func (_AtiaBlessing *AtiaBlessingSession) SetPriceFeedData(priceFeedData PythConverter) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetPriceFeedData(&_AtiaBlessing.TransactOpts, priceFeedData)
}

// SetPriceFeedData is a paid mutator transaction binding the contract method 0x1b308773.
//
// Solidity: function setPriceFeedData((uint8,uint8,address,bytes32,uint256) priceFeedData) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) SetPriceFeedData(priceFeedData PythConverter) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetPriceFeedData(&_AtiaBlessing.TransactOpts, priceFeedData)
}

// SetResetDuration is a paid mutator transaction binding the contract method 0x886ce65c.
//
// Solidity: function setResetDuration(uint256 newResetDuration) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) SetResetDuration(opts *bind.TransactOpts, newResetDuration *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "setResetDuration", newResetDuration)
}

// SetResetDuration is a paid mutator transaction binding the contract method 0x886ce65c.
//
// Solidity: function setResetDuration(uint256 newResetDuration) returns()
func (_AtiaBlessing *AtiaBlessingSession) SetResetDuration(newResetDuration *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetResetDuration(&_AtiaBlessing.TransactOpts, newResetDuration)
}

// SetResetDuration is a paid mutator transaction binding the contract method 0x886ce65c.
//
// Solidity: function setResetDuration(uint256 newResetDuration) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) SetResetDuration(newResetDuration *big.Int) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetResetDuration(&_AtiaBlessing.TransactOpts, newResetDuration)
}

// SetRestorationCosts is a paid mutator transaction binding the contract method 0x07a4f002.
//
// Solidity: function setRestorationCosts((uint32,uint80)[] restorationCosts) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) SetRestorationCosts(opts *bind.TransactOpts, restorationCosts []IAtiaShrineRestorationCost) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "setRestorationCosts", restorationCosts)
}

// SetRestorationCosts is a paid mutator transaction binding the contract method 0x07a4f002.
//
// Solidity: function setRestorationCosts((uint32,uint80)[] restorationCosts) returns()
func (_AtiaBlessing *AtiaBlessingSession) SetRestorationCosts(restorationCosts []IAtiaShrineRestorationCost) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetRestorationCosts(&_AtiaBlessing.TransactOpts, restorationCosts)
}

// SetRestorationCosts is a paid mutator transaction binding the contract method 0x07a4f002.
//
// Solidity: function setRestorationCosts((uint32,uint80)[] restorationCosts) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) SetRestorationCosts(restorationCosts []IAtiaShrineRestorationCost) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetRestorationCosts(&_AtiaBlessing.TransactOpts, restorationCosts)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury) returns()
func (_AtiaBlessing *AtiaBlessingTransactor) SetTreasury(opts *bind.TransactOpts, treasury common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.contract.Transact(opts, "setTreasury", treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury) returns()
func (_AtiaBlessing *AtiaBlessingSession) SetTreasury(treasury common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetTreasury(&_AtiaBlessing.TransactOpts, treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury) returns()
func (_AtiaBlessing *AtiaBlessingTransactorSession) SetTreasury(treasury common.Address) (*types.Transaction, error) {
	return _AtiaBlessing.Contract.SetTreasury(&_AtiaBlessing.TransactOpts, treasury)
}

// AtiaBlessingActiveStreakLifeTimeChangedIterator is returned from FilterActiveStreakLifeTimeChanged and is used to iterate over the raw logs and unpacked data for ActiveStreakLifeTimeChanged events raised by the AtiaBlessing contract.
type AtiaBlessingActiveStreakLifeTimeChangedIterator struct {
	Event *AtiaBlessingActiveStreakLifeTimeChanged // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingActiveStreakLifeTimeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingActiveStreakLifeTimeChanged)
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
		it.Event = new(AtiaBlessingActiveStreakLifeTimeChanged)
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
func (it *AtiaBlessingActiveStreakLifeTimeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingActiveStreakLifeTimeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingActiveStreakLifeTimeChanged represents a ActiveStreakLifeTimeChanged event raised by the AtiaBlessing contract.
type AtiaBlessingActiveStreakLifeTimeChanged struct {
	StreakLifeTime *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterActiveStreakLifeTimeChanged is a free log retrieval operation binding the contract event 0x90fa25b1bc7d8c44b57f07b2856b862d1b3f56e8dc385055b6b76b1977eea976.
//
// Solidity: event ActiveStreakLifeTimeChanged(uint256 streakLifeTime)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterActiveStreakLifeTimeChanged(opts *bind.FilterOpts) (*AtiaBlessingActiveStreakLifeTimeChangedIterator, error) {

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "ActiveStreakLifeTimeChanged")
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingActiveStreakLifeTimeChangedIterator{contract: _AtiaBlessing.contract, event: "ActiveStreakLifeTimeChanged", logs: logs, sub: sub}, nil
}

// WatchActiveStreakLifeTimeChanged is a free log subscription operation binding the contract event 0x90fa25b1bc7d8c44b57f07b2856b862d1b3f56e8dc385055b6b76b1977eea976.
//
// Solidity: event ActiveStreakLifeTimeChanged(uint256 streakLifeTime)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchActiveStreakLifeTimeChanged(opts *bind.WatchOpts, sink chan<- *AtiaBlessingActiveStreakLifeTimeChanged) (event.Subscription, error) {

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "ActiveStreakLifeTimeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingActiveStreakLifeTimeChanged)
				if err := _AtiaBlessing.contract.UnpackLog(event, "ActiveStreakLifeTimeChanged", log); err != nil {
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

// ParseActiveStreakLifeTimeChanged is a log parse operation binding the contract event 0x90fa25b1bc7d8c44b57f07b2856b862d1b3f56e8dc385055b6b76b1977eea976.
//
// Solidity: event ActiveStreakLifeTimeChanged(uint256 streakLifeTime)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseActiveStreakLifeTimeChanged(log types.Log) (*AtiaBlessingActiveStreakLifeTimeChanged, error) {
	event := new(AtiaBlessingActiveStreakLifeTimeChanged)
	if err := _AtiaBlessing.contract.UnpackLog(event, "ActiveStreakLifeTimeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingDailyPrayerThresholdUpdatedIterator is returned from FilterDailyPrayerThresholdUpdated and is used to iterate over the raw logs and unpacked data for DailyPrayerThresholdUpdated events raised by the AtiaBlessing contract.
type AtiaBlessingDailyPrayerThresholdUpdatedIterator struct {
	Event *AtiaBlessingDailyPrayerThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingDailyPrayerThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingDailyPrayerThresholdUpdated)
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
		it.Event = new(AtiaBlessingDailyPrayerThresholdUpdated)
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
func (it *AtiaBlessingDailyPrayerThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingDailyPrayerThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingDailyPrayerThresholdUpdated represents a DailyPrayerThresholdUpdated event raised by the AtiaBlessing contract.
type AtiaBlessingDailyPrayerThresholdUpdated struct {
	DailyPrayerThreshold *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDailyPrayerThresholdUpdated is a free log retrieval operation binding the contract event 0x29dc16cc311bf8e4de48a39e8fa751ddd5083c329187056f59857917df8a312e.
//
// Solidity: event DailyPrayerThresholdUpdated(uint256 dailyPrayerThreshold)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterDailyPrayerThresholdUpdated(opts *bind.FilterOpts) (*AtiaBlessingDailyPrayerThresholdUpdatedIterator, error) {

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "DailyPrayerThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingDailyPrayerThresholdUpdatedIterator{contract: _AtiaBlessing.contract, event: "DailyPrayerThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyPrayerThresholdUpdated is a free log subscription operation binding the contract event 0x29dc16cc311bf8e4de48a39e8fa751ddd5083c329187056f59857917df8a312e.
//
// Solidity: event DailyPrayerThresholdUpdated(uint256 dailyPrayerThreshold)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchDailyPrayerThresholdUpdated(opts *bind.WatchOpts, sink chan<- *AtiaBlessingDailyPrayerThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "DailyPrayerThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingDailyPrayerThresholdUpdated)
				if err := _AtiaBlessing.contract.UnpackLog(event, "DailyPrayerThresholdUpdated", log); err != nil {
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

// ParseDailyPrayerThresholdUpdated is a log parse operation binding the contract event 0x29dc16cc311bf8e4de48a39e8fa751ddd5083c329187056f59857917df8a312e.
//
// Solidity: event DailyPrayerThresholdUpdated(uint256 dailyPrayerThreshold)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseDailyPrayerThresholdUpdated(log types.Log) (*AtiaBlessingDailyPrayerThresholdUpdated, error) {
	event := new(AtiaBlessingDailyPrayerThresholdUpdated)
	if err := _AtiaBlessing.contract.UnpackLog(event, "DailyPrayerThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingFreezePeriodUpdatedIterator is returned from FilterFreezePeriodUpdated and is used to iterate over the raw logs and unpacked data for FreezePeriodUpdated events raised by the AtiaBlessing contract.
type AtiaBlessingFreezePeriodUpdatedIterator struct {
	Event *AtiaBlessingFreezePeriodUpdated // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingFreezePeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingFreezePeriodUpdated)
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
		it.Event = new(AtiaBlessingFreezePeriodUpdated)
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
func (it *AtiaBlessingFreezePeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingFreezePeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingFreezePeriodUpdated represents a FreezePeriodUpdated event raised by the AtiaBlessing contract.
type AtiaBlessingFreezePeriodUpdated struct {
	NewFreezePeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFreezePeriodUpdated is a free log retrieval operation binding the contract event 0xe5cfb067fd3553c2c8f52bd13990a943e4eedc525c139207ebf6ae790f6d190b.
//
// Solidity: event FreezePeriodUpdated(uint256 newFreezePeriod)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterFreezePeriodUpdated(opts *bind.FilterOpts) (*AtiaBlessingFreezePeriodUpdatedIterator, error) {

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "FreezePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingFreezePeriodUpdatedIterator{contract: _AtiaBlessing.contract, event: "FreezePeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchFreezePeriodUpdated is a free log subscription operation binding the contract event 0xe5cfb067fd3553c2c8f52bd13990a943e4eedc525c139207ebf6ae790f6d190b.
//
// Solidity: event FreezePeriodUpdated(uint256 newFreezePeriod)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchFreezePeriodUpdated(opts *bind.WatchOpts, sink chan<- *AtiaBlessingFreezePeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "FreezePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingFreezePeriodUpdated)
				if err := _AtiaBlessing.contract.UnpackLog(event, "FreezePeriodUpdated", log); err != nil {
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

// ParseFreezePeriodUpdated is a log parse operation binding the contract event 0xe5cfb067fd3553c2c8f52bd13990a943e4eedc525c139207ebf6ae790f6d190b.
//
// Solidity: event FreezePeriodUpdated(uint256 newFreezePeriod)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseFreezePeriodUpdated(log types.Log) (*AtiaBlessingFreezePeriodUpdated, error) {
	event := new(AtiaBlessingFreezePeriodUpdated)
	if err := _AtiaBlessing.contract.UnpackLog(event, "FreezePeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AtiaBlessing contract.
type AtiaBlessingInitializedIterator struct {
	Event *AtiaBlessingInitialized // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingInitialized)
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
		it.Event = new(AtiaBlessingInitialized)
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
func (it *AtiaBlessingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingInitialized represents a Initialized event raised by the AtiaBlessing contract.
type AtiaBlessingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterInitialized(opts *bind.FilterOpts) (*AtiaBlessingInitializedIterator, error) {

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingInitializedIterator{contract: _AtiaBlessing.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AtiaBlessingInitialized) (event.Subscription, error) {

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingInitialized)
				if err := _AtiaBlessing.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AtiaBlessing *AtiaBlessingFilterer) ParseInitialized(log types.Log) (*AtiaBlessingInitialized, error) {
	event := new(AtiaBlessingInitialized)
	if err := _AtiaBlessing.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingPaymentTokenUpdatedIterator is returned from FilterPaymentTokenUpdated and is used to iterate over the raw logs and unpacked data for PaymentTokenUpdated events raised by the AtiaBlessing contract.
type AtiaBlessingPaymentTokenUpdatedIterator struct {
	Event *AtiaBlessingPaymentTokenUpdated // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingPaymentTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingPaymentTokenUpdated)
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
		it.Event = new(AtiaBlessingPaymentTokenUpdated)
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
func (it *AtiaBlessingPaymentTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingPaymentTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingPaymentTokenUpdated represents a PaymentTokenUpdated event raised by the AtiaBlessing contract.
type AtiaBlessingPaymentTokenUpdated struct {
	NewPaymentToken common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenUpdated is a free log retrieval operation binding the contract event 0xbd4032c1c91da2791730ea1bbc82c6b6f857da7c0a8318143d19ef74e62cd913.
//
// Solidity: event PaymentTokenUpdated(address newPaymentToken)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterPaymentTokenUpdated(opts *bind.FilterOpts) (*AtiaBlessingPaymentTokenUpdatedIterator, error) {

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "PaymentTokenUpdated")
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingPaymentTokenUpdatedIterator{contract: _AtiaBlessing.contract, event: "PaymentTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenUpdated is a free log subscription operation binding the contract event 0xbd4032c1c91da2791730ea1bbc82c6b6f857da7c0a8318143d19ef74e62cd913.
//
// Solidity: event PaymentTokenUpdated(address newPaymentToken)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchPaymentTokenUpdated(opts *bind.WatchOpts, sink chan<- *AtiaBlessingPaymentTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "PaymentTokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingPaymentTokenUpdated)
				if err := _AtiaBlessing.contract.UnpackLog(event, "PaymentTokenUpdated", log); err != nil {
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

// ParsePaymentTokenUpdated is a log parse operation binding the contract event 0xbd4032c1c91da2791730ea1bbc82c6b6f857da7c0a8318143d19ef74e62cd913.
//
// Solidity: event PaymentTokenUpdated(address newPaymentToken)
func (_AtiaBlessing *AtiaBlessingFilterer) ParsePaymentTokenUpdated(log types.Log) (*AtiaBlessingPaymentTokenUpdated, error) {
	event := new(AtiaBlessingPaymentTokenUpdated)
	if err := _AtiaBlessing.contract.UnpackLog(event, "PaymentTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingPrayerCountSyncedIterator is returned from FilterPrayerCountSynced and is used to iterate over the raw logs and unpacked data for PrayerCountSynced events raised by the AtiaBlessing contract.
type AtiaBlessingPrayerCountSyncedIterator struct {
	Event *AtiaBlessingPrayerCountSynced // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingPrayerCountSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingPrayerCountSynced)
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
		it.Event = new(AtiaBlessingPrayerCountSynced)
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
func (it *AtiaBlessingPrayerCountSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingPrayerCountSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingPrayerCountSynced represents a PrayerCountSynced event raised by the AtiaBlessing contract.
type AtiaBlessingPrayerCountSynced struct {
	Prayer          common.Address
	PrayerCountInfo IAtiaShrinePrayerCountInfo
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPrayerCountSynced is a free log retrieval operation binding the contract event 0xce6788d4098848c2ef2ca76ed97a6572db9785a6bbf18303399728e383f2eb00.
//
// Solidity: event PrayerCountSynced(address indexed prayer, (uint64,uint64) prayerCountInfo)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterPrayerCountSynced(opts *bind.FilterOpts, prayer []common.Address) (*AtiaBlessingPrayerCountSyncedIterator, error) {

	var prayerRule []interface{}
	for _, prayerItem := range prayer {
		prayerRule = append(prayerRule, prayerItem)
	}

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "PrayerCountSynced", prayerRule)
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingPrayerCountSyncedIterator{contract: _AtiaBlessing.contract, event: "PrayerCountSynced", logs: logs, sub: sub}, nil
}

// WatchPrayerCountSynced is a free log subscription operation binding the contract event 0xce6788d4098848c2ef2ca76ed97a6572db9785a6bbf18303399728e383f2eb00.
//
// Solidity: event PrayerCountSynced(address indexed prayer, (uint64,uint64) prayerCountInfo)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchPrayerCountSynced(opts *bind.WatchOpts, sink chan<- *AtiaBlessingPrayerCountSynced, prayer []common.Address) (event.Subscription, error) {

	var prayerRule []interface{}
	for _, prayerItem := range prayer {
		prayerRule = append(prayerRule, prayerItem)
	}

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "PrayerCountSynced", prayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingPrayerCountSynced)
				if err := _AtiaBlessing.contract.UnpackLog(event, "PrayerCountSynced", log); err != nil {
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

// ParsePrayerCountSynced is a log parse operation binding the contract event 0xce6788d4098848c2ef2ca76ed97a6572db9785a6bbf18303399728e383f2eb00.
//
// Solidity: event PrayerCountSynced(address indexed prayer, (uint64,uint64) prayerCountInfo)
func (_AtiaBlessing *AtiaBlessingFilterer) ParsePrayerCountSynced(log types.Log) (*AtiaBlessingPrayerCountSynced, error) {
	event := new(AtiaBlessingPrayerCountSynced)
	if err := _AtiaBlessing.contract.UnpackLog(event, "PrayerCountSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingPriceFeedDataUpdatedIterator is returned from FilterPriceFeedDataUpdated and is used to iterate over the raw logs and unpacked data for PriceFeedDataUpdated events raised by the AtiaBlessing contract.
type AtiaBlessingPriceFeedDataUpdatedIterator struct {
	Event *AtiaBlessingPriceFeedDataUpdated // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingPriceFeedDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingPriceFeedDataUpdated)
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
		it.Event = new(AtiaBlessingPriceFeedDataUpdated)
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
func (it *AtiaBlessingPriceFeedDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingPriceFeedDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingPriceFeedDataUpdated represents a PriceFeedDataUpdated event raised by the AtiaBlessing contract.
type AtiaBlessingPriceFeedDataUpdated struct {
	PriceFeedData PythConverter
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedDataUpdated is a free log retrieval operation binding the contract event 0x62789006c233d6af569da152e52aed4ecbf9a487701604d3456e4419e93121a2.
//
// Solidity: event PriceFeedDataUpdated((uint8,uint8,address,bytes32,uint256) priceFeedData)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterPriceFeedDataUpdated(opts *bind.FilterOpts) (*AtiaBlessingPriceFeedDataUpdatedIterator, error) {

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "PriceFeedDataUpdated")
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingPriceFeedDataUpdatedIterator{contract: _AtiaBlessing.contract, event: "PriceFeedDataUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceFeedDataUpdated is a free log subscription operation binding the contract event 0x62789006c233d6af569da152e52aed4ecbf9a487701604d3456e4419e93121a2.
//
// Solidity: event PriceFeedDataUpdated((uint8,uint8,address,bytes32,uint256) priceFeedData)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchPriceFeedDataUpdated(opts *bind.WatchOpts, sink chan<- *AtiaBlessingPriceFeedDataUpdated) (event.Subscription, error) {

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "PriceFeedDataUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingPriceFeedDataUpdated)
				if err := _AtiaBlessing.contract.UnpackLog(event, "PriceFeedDataUpdated", log); err != nil {
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

// ParsePriceFeedDataUpdated is a log parse operation binding the contract event 0x62789006c233d6af569da152e52aed4ecbf9a487701604d3456e4419e93121a2.
//
// Solidity: event PriceFeedDataUpdated((uint8,uint8,address,bytes32,uint256) priceFeedData)
func (_AtiaBlessing *AtiaBlessingFilterer) ParsePriceFeedDataUpdated(log types.Log) (*AtiaBlessingPriceFeedDataUpdated, error) {
	event := new(AtiaBlessingPriceFeedDataUpdated)
	if err := _AtiaBlessing.contract.UnpackLog(event, "PriceFeedDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingResetDurationUpdatedIterator is returned from FilterResetDurationUpdated and is used to iterate over the raw logs and unpacked data for ResetDurationUpdated events raised by the AtiaBlessing contract.
type AtiaBlessingResetDurationUpdatedIterator struct {
	Event *AtiaBlessingResetDurationUpdated // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingResetDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingResetDurationUpdated)
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
		it.Event = new(AtiaBlessingResetDurationUpdated)
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
func (it *AtiaBlessingResetDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingResetDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingResetDurationUpdated represents a ResetDurationUpdated event raised by the AtiaBlessing contract.
type AtiaBlessingResetDurationUpdated struct {
	NewResetDuration *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterResetDurationUpdated is a free log retrieval operation binding the contract event 0x4534d7ef285882c47146e4b0a0f0192a2e3e8dc978949664ddf2c82c5baef367.
//
// Solidity: event ResetDurationUpdated(uint256 newResetDuration)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterResetDurationUpdated(opts *bind.FilterOpts) (*AtiaBlessingResetDurationUpdatedIterator, error) {

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "ResetDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingResetDurationUpdatedIterator{contract: _AtiaBlessing.contract, event: "ResetDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchResetDurationUpdated is a free log subscription operation binding the contract event 0x4534d7ef285882c47146e4b0a0f0192a2e3e8dc978949664ddf2c82c5baef367.
//
// Solidity: event ResetDurationUpdated(uint256 newResetDuration)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchResetDurationUpdated(opts *bind.WatchOpts, sink chan<- *AtiaBlessingResetDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "ResetDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingResetDurationUpdated)
				if err := _AtiaBlessing.contract.UnpackLog(event, "ResetDurationUpdated", log); err != nil {
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

// ParseResetDurationUpdated is a log parse operation binding the contract event 0x4534d7ef285882c47146e4b0a0f0192a2e3e8dc978949664ddf2c82c5baef367.
//
// Solidity: event ResetDurationUpdated(uint256 newResetDuration)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseResetDurationUpdated(log types.Log) (*AtiaBlessingResetDurationUpdated, error) {
	event := new(AtiaBlessingResetDurationUpdated)
	if err := _AtiaBlessing.contract.UnpackLog(event, "ResetDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingRestorationCostUpdatedIterator is returned from FilterRestorationCostUpdated and is used to iterate over the raw logs and unpacked data for RestorationCostUpdated events raised by the AtiaBlessing contract.
type AtiaBlessingRestorationCostUpdatedIterator struct {
	Event *AtiaBlessingRestorationCostUpdated // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingRestorationCostUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingRestorationCostUpdated)
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
		it.Event = new(AtiaBlessingRestorationCostUpdated)
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
func (it *AtiaBlessingRestorationCostUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingRestorationCostUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingRestorationCostUpdated represents a RestorationCostUpdated event raised by the AtiaBlessing contract.
type AtiaBlessingRestorationCostUpdated struct {
	NewRestorationCosts []IAtiaShrineRestorationCost
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRestorationCostUpdated is a free log retrieval operation binding the contract event 0xdb16718cd968791fb87ccb9de4392526cf98a50acf9f2ce4a3e66e95e6a51380.
//
// Solidity: event RestorationCostUpdated((uint32,uint80)[] newRestorationCosts)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterRestorationCostUpdated(opts *bind.FilterOpts) (*AtiaBlessingRestorationCostUpdatedIterator, error) {

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "RestorationCostUpdated")
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingRestorationCostUpdatedIterator{contract: _AtiaBlessing.contract, event: "RestorationCostUpdated", logs: logs, sub: sub}, nil
}

// WatchRestorationCostUpdated is a free log subscription operation binding the contract event 0xdb16718cd968791fb87ccb9de4392526cf98a50acf9f2ce4a3e66e95e6a51380.
//
// Solidity: event RestorationCostUpdated((uint32,uint80)[] newRestorationCosts)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchRestorationCostUpdated(opts *bind.WatchOpts, sink chan<- *AtiaBlessingRestorationCostUpdated) (event.Subscription, error) {

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "RestorationCostUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingRestorationCostUpdated)
				if err := _AtiaBlessing.contract.UnpackLog(event, "RestorationCostUpdated", log); err != nil {
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

// ParseRestorationCostUpdated is a log parse operation binding the contract event 0xdb16718cd968791fb87ccb9de4392526cf98a50acf9f2ce4a3e66e95e6a51380.
//
// Solidity: event RestorationCostUpdated((uint32,uint80)[] newRestorationCosts)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseRestorationCostUpdated(log types.Log) (*AtiaBlessingRestorationCostUpdated, error) {
	event := new(AtiaBlessingRestorationCostUpdated)
	if err := _AtiaBlessing.contract.UnpackLog(event, "RestorationCostUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AtiaBlessing contract.
type AtiaBlessingRoleAdminChangedIterator struct {
	Event *AtiaBlessingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingRoleAdminChanged)
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
		it.Event = new(AtiaBlessingRoleAdminChanged)
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
func (it *AtiaBlessingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingRoleAdminChanged represents a RoleAdminChanged event raised by the AtiaBlessing contract.
type AtiaBlessingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AtiaBlessingRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingRoleAdminChangedIterator{contract: _AtiaBlessing.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AtiaBlessingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingRoleAdminChanged)
				if err := _AtiaBlessing.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseRoleAdminChanged(log types.Log) (*AtiaBlessingRoleAdminChanged, error) {
	event := new(AtiaBlessingRoleAdminChanged)
	if err := _AtiaBlessing.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AtiaBlessing contract.
type AtiaBlessingRoleGrantedIterator struct {
	Event *AtiaBlessingRoleGranted // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingRoleGranted)
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
		it.Event = new(AtiaBlessingRoleGranted)
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
func (it *AtiaBlessingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingRoleGranted represents a RoleGranted event raised by the AtiaBlessing contract.
type AtiaBlessingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AtiaBlessingRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingRoleGrantedIterator{contract: _AtiaBlessing.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AtiaBlessingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingRoleGranted)
				if err := _AtiaBlessing.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseRoleGranted(log types.Log) (*AtiaBlessingRoleGranted, error) {
	event := new(AtiaBlessingRoleGranted)
	if err := _AtiaBlessing.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AtiaBlessing contract.
type AtiaBlessingRoleRevokedIterator struct {
	Event *AtiaBlessingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingRoleRevoked)
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
		it.Event = new(AtiaBlessingRoleRevoked)
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
func (it *AtiaBlessingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingRoleRevoked represents a RoleRevoked event raised by the AtiaBlessing contract.
type AtiaBlessingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AtiaBlessingRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingRoleRevokedIterator{contract: _AtiaBlessing.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AtiaBlessingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingRoleRevoked)
				if err := _AtiaBlessing.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseRoleRevoked(log types.Log) (*AtiaBlessingRoleRevoked, error) {
	event := new(AtiaBlessingRoleRevoked)
	if err := _AtiaBlessing.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingStreakRestoredIterator is returned from FilterStreakRestored and is used to iterate over the raw logs and unpacked data for StreakRestored events raised by the AtiaBlessing contract.
type AtiaBlessingStreakRestoredIterator struct {
	Event *AtiaBlessingStreakRestored // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingStreakRestoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingStreakRestored)
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
		it.Event = new(AtiaBlessingStreakRestored)
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
func (it *AtiaBlessingStreakRestoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingStreakRestoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingStreakRestored represents a StreakRestored event raised by the AtiaBlessing contract.
type AtiaBlessingStreakRestored struct {
	User                common.Address
	RestoredStreakCount *big.Int
	LastUpdated         *big.Int
	CostInUsd           *big.Int
	CostInPaymentToken  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterStreakRestored is a free log retrieval operation binding the contract event 0x2362cec5a8600f629df24ff52b88e2acf8e748d8726be66a93a3f4db7ceb8b94.
//
// Solidity: event StreakRestored(address indexed user, uint256 indexed restoredStreakCount, uint256 indexed lastUpdated, uint256 costInUsd, uint256 costInPaymentToken)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterStreakRestored(opts *bind.FilterOpts, user []common.Address, restoredStreakCount []*big.Int, lastUpdated []*big.Int) (*AtiaBlessingStreakRestoredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var restoredStreakCountRule []interface{}
	for _, restoredStreakCountItem := range restoredStreakCount {
		restoredStreakCountRule = append(restoredStreakCountRule, restoredStreakCountItem)
	}
	var lastUpdatedRule []interface{}
	for _, lastUpdatedItem := range lastUpdated {
		lastUpdatedRule = append(lastUpdatedRule, lastUpdatedItem)
	}

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "StreakRestored", userRule, restoredStreakCountRule, lastUpdatedRule)
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingStreakRestoredIterator{contract: _AtiaBlessing.contract, event: "StreakRestored", logs: logs, sub: sub}, nil
}

// WatchStreakRestored is a free log subscription operation binding the contract event 0x2362cec5a8600f629df24ff52b88e2acf8e748d8726be66a93a3f4db7ceb8b94.
//
// Solidity: event StreakRestored(address indexed user, uint256 indexed restoredStreakCount, uint256 indexed lastUpdated, uint256 costInUsd, uint256 costInPaymentToken)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchStreakRestored(opts *bind.WatchOpts, sink chan<- *AtiaBlessingStreakRestored, user []common.Address, restoredStreakCount []*big.Int, lastUpdated []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var restoredStreakCountRule []interface{}
	for _, restoredStreakCountItem := range restoredStreakCount {
		restoredStreakCountRule = append(restoredStreakCountRule, restoredStreakCountItem)
	}
	var lastUpdatedRule []interface{}
	for _, lastUpdatedItem := range lastUpdated {
		lastUpdatedRule = append(lastUpdatedRule, lastUpdatedItem)
	}

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "StreakRestored", userRule, restoredStreakCountRule, lastUpdatedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingStreakRestored)
				if err := _AtiaBlessing.contract.UnpackLog(event, "StreakRestored", log); err != nil {
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

// ParseStreakRestored is a log parse operation binding the contract event 0x2362cec5a8600f629df24ff52b88e2acf8e748d8726be66a93a3f4db7ceb8b94.
//
// Solidity: event StreakRestored(address indexed user, uint256 indexed restoredStreakCount, uint256 indexed lastUpdated, uint256 costInUsd, uint256 costInPaymentToken)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseStreakRestored(log types.Log) (*AtiaBlessingStreakRestored, error) {
	event := new(AtiaBlessingStreakRestored)
	if err := _AtiaBlessing.contract.UnpackLog(event, "StreakRestored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingStreakUpdatedIterator is returned from FilterStreakUpdated and is used to iterate over the raw logs and unpacked data for StreakUpdated events raised by the AtiaBlessing contract.
type AtiaBlessingStreakUpdatedIterator struct {
	Event *AtiaBlessingStreakUpdated // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingStreakUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingStreakUpdated)
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
		it.Event = new(AtiaBlessingStreakUpdated)
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
func (it *AtiaBlessingStreakUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingStreakUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingStreakUpdated represents a StreakUpdated event raised by the AtiaBlessing contract.
type AtiaBlessingStreakUpdated struct {
	User               common.Address
	NewAmount          *big.Int
	NewLastUpdated     *big.Int
	LongestStreakCount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterStreakUpdated is a free log retrieval operation binding the contract event 0x62fbe39954276c01e2d7a17f048cec9d25cc2099b61f7858a7accd2b4216382e.
//
// Solidity: event StreakUpdated(address indexed user, uint256 indexed newAmount, uint256 indexed newLastUpdated, uint256 longestStreakCount)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterStreakUpdated(opts *bind.FilterOpts, user []common.Address, newAmount []*big.Int, newLastUpdated []*big.Int) (*AtiaBlessingStreakUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}
	var newLastUpdatedRule []interface{}
	for _, newLastUpdatedItem := range newLastUpdated {
		newLastUpdatedRule = append(newLastUpdatedRule, newLastUpdatedItem)
	}

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "StreakUpdated", userRule, newAmountRule, newLastUpdatedRule)
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingStreakUpdatedIterator{contract: _AtiaBlessing.contract, event: "StreakUpdated", logs: logs, sub: sub}, nil
}

// WatchStreakUpdated is a free log subscription operation binding the contract event 0x62fbe39954276c01e2d7a17f048cec9d25cc2099b61f7858a7accd2b4216382e.
//
// Solidity: event StreakUpdated(address indexed user, uint256 indexed newAmount, uint256 indexed newLastUpdated, uint256 longestStreakCount)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchStreakUpdated(opts *bind.WatchOpts, sink chan<- *AtiaBlessingStreakUpdated, user []common.Address, newAmount []*big.Int, newLastUpdated []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}
	var newLastUpdatedRule []interface{}
	for _, newLastUpdatedItem := range newLastUpdated {
		newLastUpdatedRule = append(newLastUpdatedRule, newLastUpdatedItem)
	}

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "StreakUpdated", userRule, newAmountRule, newLastUpdatedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingStreakUpdated)
				if err := _AtiaBlessing.contract.UnpackLog(event, "StreakUpdated", log); err != nil {
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

// ParseStreakUpdated is a log parse operation binding the contract event 0x62fbe39954276c01e2d7a17f048cec9d25cc2099b61f7858a7accd2b4216382e.
//
// Solidity: event StreakUpdated(address indexed user, uint256 indexed newAmount, uint256 indexed newLastUpdated, uint256 longestStreakCount)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseStreakUpdated(log types.Log) (*AtiaBlessingStreakUpdated, error) {
	event := new(AtiaBlessingStreakUpdated)
	if err := _AtiaBlessing.contract.UnpackLog(event, "StreakUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtiaBlessingTreasuryUpdatedIterator is returned from FilterTreasuryUpdated and is used to iterate over the raw logs and unpacked data for TreasuryUpdated events raised by the AtiaBlessing contract.
type AtiaBlessingTreasuryUpdatedIterator struct {
	Event *AtiaBlessingTreasuryUpdated // Event containing the contract specifics and raw log

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
func (it *AtiaBlessingTreasuryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtiaBlessingTreasuryUpdated)
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
		it.Event = new(AtiaBlessingTreasuryUpdated)
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
func (it *AtiaBlessingTreasuryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtiaBlessingTreasuryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtiaBlessingTreasuryUpdated represents a TreasuryUpdated event raised by the AtiaBlessing contract.
type AtiaBlessingTreasuryUpdated struct {
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdated is a free log retrieval operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_AtiaBlessing *AtiaBlessingFilterer) FilterTreasuryUpdated(opts *bind.FilterOpts) (*AtiaBlessingTreasuryUpdatedIterator, error) {

	logs, sub, err := _AtiaBlessing.contract.FilterLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return &AtiaBlessingTreasuryUpdatedIterator{contract: _AtiaBlessing.contract, event: "TreasuryUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdated is a free log subscription operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_AtiaBlessing *AtiaBlessingFilterer) WatchTreasuryUpdated(opts *bind.WatchOpts, sink chan<- *AtiaBlessingTreasuryUpdated) (event.Subscription, error) {

	logs, sub, err := _AtiaBlessing.contract.WatchLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtiaBlessingTreasuryUpdated)
				if err := _AtiaBlessing.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
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

// ParseTreasuryUpdated is a log parse operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_AtiaBlessing *AtiaBlessingFilterer) ParseTreasuryUpdated(log types.Log) (*AtiaBlessingTreasuryUpdated, error) {
	event := new(AtiaBlessingTreasuryUpdated)
	if err := _AtiaBlessing.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
