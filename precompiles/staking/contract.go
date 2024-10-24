// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorSrcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorDstAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"beginRedelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationHeight\",\"type\":\"uint256\"}],\"name\":\"cancelUnbondingDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commission\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"pubkey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"createValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"delegatorAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorAddr\",\"type\":\"string\"}],\"name\":\"delegation\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"internalType\":\"structDelegation\",\"name\":\"delegation\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"delegatorAddr\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"delegatorDelegations\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"internalType\":\"structDelegation\",\"name\":\"delegation\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structDelegationResponse[]\",\"name\":\"delegationResponses\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"paginationResult\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"delegatorAddr\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"delegatorUnbondingDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"unbondingId\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"}],\"internalType\":\"structUnbondingDelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structUnbondingDelegation[]\",\"name\":\"unbondingResponses\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"paginationResult\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"delegatorAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorAddr\",\"type\":\"string\"}],\"name\":\"delegatorValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commissionRates\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"updateTime\",\"type\":\"uint256\"}],\"internalType\":\"structCommission\",\"name\":\"commission\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"},{\"internalType\":\"uint64[]\",\"name\":\"unbondingIds\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidator\",\"name\":\"validator\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"delegatorAddr\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"delegatorValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commissionRates\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"updateTime\",\"type\":\"uint256\"}],\"internalType\":\"structCommission\",\"name\":\"commission\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"},{\"internalType\":\"uint64[]\",\"name\":\"unbondingIds\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"paginationResult\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isNull\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structNullableUint\",\"name\":\"commissionRate\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isNull\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structNullableUint\",\"name\":\"minSelfDelegation\",\"type\":\"tuple\"}],\"name\":\"editValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"maxValidators\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxEntries\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"historicalEntries\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"bondDenom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minCommissionRate\",\"type\":\"uint256\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"notBondedTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondedTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"srcValidatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dstValidatorAddress\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"redelegations\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorSrcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorDstAddress\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesDst\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"unbondingId\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"}],\"internalType\":\"structRedelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structRedelegation\",\"name\":\"redelegation\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesDst\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"unbondingId\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"}],\"internalType\":\"structRedelegationEntry\",\"name\":\"redelegationEntry\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structRedelegationEntryResponse[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structRedelegationResponse[]\",\"name\":\"redelegationResponses\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"paginationResult\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"delegatorAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorAddr\",\"type\":\"string\"}],\"name\":\"unbondingDelegation\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"unbondingId\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"}],\"internalType\":\"structUnbondingDelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structUnbondingDelegation\",\"name\":\"unbond\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"validator\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commissionRates\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"updateTime\",\"type\":\"uint256\"}],\"internalType\":\"structCommission\",\"name\":\"commission\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"},{\"internalType\":\"uint64[]\",\"name\":\"unbondingIds\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidator\",\"name\":\"validator\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddr\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"validatorDelegations\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"internalType\":\"structDelegation\",\"name\":\"delegation\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structDelegationResponse[]\",\"name\":\"delegationResponses\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"paginationResult\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddr\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"validatorUnbondingDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"delegatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"unbondingId\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"}],\"internalType\":\"structUnbondingDelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structUnbondingDelegation[]\",\"name\":\"unbondingResponses\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"paginationResult\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"validators\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"operatorAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commissionRates\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"updateTime\",\"type\":\"uint256\"}],\"internalType\":\"structCommission\",\"name\":\"commission\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"},{\"internalType\":\"uint64[]\",\"name\":\"unbondingIds\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"paginationResult\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// Delegation is a free data retrieval call binding the contract method 0x25b8eda2.
//
// Solidity: function delegation(string delegatorAddr, string validatorAddr) view returns((string,string,uint256) delegation, uint256 balance)
func (_Staking *StakingCaller) Delegation(opts *bind.CallOpts, delegatorAddr string, validatorAddr string) (struct {
	Delegation Delegation
	Balance    *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegation", delegatorAddr, validatorAddr)

	outstruct := new(struct {
		Delegation Delegation
		Balance    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Delegation = *abi.ConvertType(out[0], new(Delegation)).(*Delegation)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Delegation is a free data retrieval call binding the contract method 0x25b8eda2.
//
// Solidity: function delegation(string delegatorAddr, string validatorAddr) view returns((string,string,uint256) delegation, uint256 balance)
func (_Staking *StakingSession) Delegation(delegatorAddr string, validatorAddr string) (struct {
	Delegation Delegation
	Balance    *big.Int
}, error) {
	return _Staking.Contract.Delegation(&_Staking.CallOpts, delegatorAddr, validatorAddr)
}

// Delegation is a free data retrieval call binding the contract method 0x25b8eda2.
//
// Solidity: function delegation(string delegatorAddr, string validatorAddr) view returns((string,string,uint256) delegation, uint256 balance)
func (_Staking *StakingCallerSession) Delegation(delegatorAddr string, validatorAddr string) (struct {
	Delegation Delegation
	Balance    *big.Int
}, error) {
	return _Staking.Contract.Delegation(&_Staking.CallOpts, delegatorAddr, validatorAddr)
}

// DelegatorDelegations is a free data retrieval call binding the contract method 0x062cac23.
//
// Solidity: function delegatorDelegations(string delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((string,string,uint256),uint256)[] delegationResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingCaller) DelegatorDelegations(opts *bind.CallOpts, delegatorAddr string, pagination PageRequest) (struct {
	DelegationResponses []DelegationResponse
	PaginationResult    PageResponse
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegatorDelegations", delegatorAddr, pagination)

	outstruct := new(struct {
		DelegationResponses []DelegationResponse
		PaginationResult    PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DelegationResponses = *abi.ConvertType(out[0], new([]DelegationResponse)).(*[]DelegationResponse)
	outstruct.PaginationResult = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// DelegatorDelegations is a free data retrieval call binding the contract method 0x062cac23.
//
// Solidity: function delegatorDelegations(string delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((string,string,uint256),uint256)[] delegationResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingSession) DelegatorDelegations(delegatorAddr string, pagination PageRequest) (struct {
	DelegationResponses []DelegationResponse
	PaginationResult    PageResponse
}, error) {
	return _Staking.Contract.DelegatorDelegations(&_Staking.CallOpts, delegatorAddr, pagination)
}

// DelegatorDelegations is a free data retrieval call binding the contract method 0x062cac23.
//
// Solidity: function delegatorDelegations(string delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((string,string,uint256),uint256)[] delegationResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingCallerSession) DelegatorDelegations(delegatorAddr string, pagination PageRequest) (struct {
	DelegationResponses []DelegationResponse
	PaginationResult    PageResponse
}, error) {
	return _Staking.Contract.DelegatorDelegations(&_Staking.CallOpts, delegatorAddr, pagination)
}

// DelegatorUnbondingDelegations is a free data retrieval call binding the contract method 0xba157bc5.
//
// Solidity: function delegatorUnbondingDelegations(string delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[])[] unbondingResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingCaller) DelegatorUnbondingDelegations(opts *bind.CallOpts, delegatorAddr string, pagination PageRequest) (struct {
	UnbondingResponses []UnbondingDelegation
	PaginationResult   PageResponse
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegatorUnbondingDelegations", delegatorAddr, pagination)

	outstruct := new(struct {
		UnbondingResponses []UnbondingDelegation
		PaginationResult   PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UnbondingResponses = *abi.ConvertType(out[0], new([]UnbondingDelegation)).(*[]UnbondingDelegation)
	outstruct.PaginationResult = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// DelegatorUnbondingDelegations is a free data retrieval call binding the contract method 0xba157bc5.
//
// Solidity: function delegatorUnbondingDelegations(string delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[])[] unbondingResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingSession) DelegatorUnbondingDelegations(delegatorAddr string, pagination PageRequest) (struct {
	UnbondingResponses []UnbondingDelegation
	PaginationResult   PageResponse
}, error) {
	return _Staking.Contract.DelegatorUnbondingDelegations(&_Staking.CallOpts, delegatorAddr, pagination)
}

// DelegatorUnbondingDelegations is a free data retrieval call binding the contract method 0xba157bc5.
//
// Solidity: function delegatorUnbondingDelegations(string delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[])[] unbondingResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingCallerSession) DelegatorUnbondingDelegations(delegatorAddr string, pagination PageRequest) (struct {
	UnbondingResponses []UnbondingDelegation
	PaginationResult   PageResponse
}, error) {
	return _Staking.Contract.DelegatorUnbondingDelegations(&_Staking.CallOpts, delegatorAddr, pagination)
}

// DelegatorValidator is a free data retrieval call binding the contract method 0x4118fcc0.
//
// Solidity: function delegatorValidator(string delegatorAddr, string validatorAddr) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[]) validator)
func (_Staking *StakingCaller) DelegatorValidator(opts *bind.CallOpts, delegatorAddr string, validatorAddr string) (Validator, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegatorValidator", delegatorAddr, validatorAddr)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// DelegatorValidator is a free data retrieval call binding the contract method 0x4118fcc0.
//
// Solidity: function delegatorValidator(string delegatorAddr, string validatorAddr) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[]) validator)
func (_Staking *StakingSession) DelegatorValidator(delegatorAddr string, validatorAddr string) (Validator, error) {
	return _Staking.Contract.DelegatorValidator(&_Staking.CallOpts, delegatorAddr, validatorAddr)
}

// DelegatorValidator is a free data retrieval call binding the contract method 0x4118fcc0.
//
// Solidity: function delegatorValidator(string delegatorAddr, string validatorAddr) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[]) validator)
func (_Staking *StakingCallerSession) DelegatorValidator(delegatorAddr string, validatorAddr string) (Validator, error) {
	return _Staking.Contract.DelegatorValidator(&_Staking.CallOpts, delegatorAddr, validatorAddr)
}

// DelegatorValidators is a free data retrieval call binding the contract method 0x0d4d2384.
//
// Solidity: function delegatorValidators(string delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[])[] validators, (bytes,uint64) paginationResult)
func (_Staking *StakingCaller) DelegatorValidators(opts *bind.CallOpts, delegatorAddr string, pagination PageRequest) (struct {
	Validators       []Validator
	PaginationResult PageResponse
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "delegatorValidators", delegatorAddr, pagination)

	outstruct := new(struct {
		Validators       []Validator
		PaginationResult PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]Validator)).(*[]Validator)
	outstruct.PaginationResult = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// DelegatorValidators is a free data retrieval call binding the contract method 0x0d4d2384.
//
// Solidity: function delegatorValidators(string delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[])[] validators, (bytes,uint64) paginationResult)
func (_Staking *StakingSession) DelegatorValidators(delegatorAddr string, pagination PageRequest) (struct {
	Validators       []Validator
	PaginationResult PageResponse
}, error) {
	return _Staking.Contract.DelegatorValidators(&_Staking.CallOpts, delegatorAddr, pagination)
}

// DelegatorValidators is a free data retrieval call binding the contract method 0x0d4d2384.
//
// Solidity: function delegatorValidators(string delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[])[] validators, (bytes,uint64) paginationResult)
func (_Staking *StakingCallerSession) DelegatorValidators(delegatorAddr string, pagination PageRequest) (struct {
	Validators       []Validator
	PaginationResult PageResponse
}, error) {
	return _Staking.Contract.DelegatorValidators(&_Staking.CallOpts, delegatorAddr, pagination)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((int64,uint32,uint32,uint32,string,uint256) params)
func (_Staking *StakingCaller) Params(opts *bind.CallOpts) (Params, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "params")

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((int64,uint32,uint32,uint32,string,uint256) params)
func (_Staking *StakingSession) Params() (Params, error) {
	return _Staking.Contract.Params(&_Staking.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((int64,uint32,uint32,uint32,string,uint256) params)
func (_Staking *StakingCallerSession) Params() (Params, error) {
	return _Staking.Contract.Params(&_Staking.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(uint256 notBondedTokens, uint256 bondedTokens)
func (_Staking *StakingCaller) Pool(opts *bind.CallOpts) (struct {
	NotBondedTokens *big.Int
	BondedTokens    *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "pool")

	outstruct := new(struct {
		NotBondedTokens *big.Int
		BondedTokens    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NotBondedTokens = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BondedTokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(uint256 notBondedTokens, uint256 bondedTokens)
func (_Staking *StakingSession) Pool() (struct {
	NotBondedTokens *big.Int
	BondedTokens    *big.Int
}, error) {
	return _Staking.Contract.Pool(&_Staking.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(uint256 notBondedTokens, uint256 bondedTokens)
func (_Staking *StakingCallerSession) Pool() (struct {
	NotBondedTokens *big.Int
	BondedTokens    *big.Int
}, error) {
	return _Staking.Contract.Pool(&_Staking.CallOpts)
}

// Redelegations is a free data retrieval call binding the contract method 0xe4227092.
//
// Solidity: function redelegations(string delegatorAddress, string srcValidatorAddress, string dstValidatorAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns(((string,string,string,(int64,int64,uint256,uint256,uint64,int64)[]),((int64,int64,uint256,uint256,uint64,int64),uint256)[])[] redelegationResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingCaller) Redelegations(opts *bind.CallOpts, delegatorAddress string, srcValidatorAddress string, dstValidatorAddress string, pageRequest PageRequest) (struct {
	RedelegationResponses []RedelegationResponse
	PaginationResult      PageResponse
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "redelegations", delegatorAddress, srcValidatorAddress, dstValidatorAddress, pageRequest)

	outstruct := new(struct {
		RedelegationResponses []RedelegationResponse
		PaginationResult      PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RedelegationResponses = *abi.ConvertType(out[0], new([]RedelegationResponse)).(*[]RedelegationResponse)
	outstruct.PaginationResult = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Redelegations is a free data retrieval call binding the contract method 0xe4227092.
//
// Solidity: function redelegations(string delegatorAddress, string srcValidatorAddress, string dstValidatorAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns(((string,string,string,(int64,int64,uint256,uint256,uint64,int64)[]),((int64,int64,uint256,uint256,uint64,int64),uint256)[])[] redelegationResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingSession) Redelegations(delegatorAddress string, srcValidatorAddress string, dstValidatorAddress string, pageRequest PageRequest) (struct {
	RedelegationResponses []RedelegationResponse
	PaginationResult      PageResponse
}, error) {
	return _Staking.Contract.Redelegations(&_Staking.CallOpts, delegatorAddress, srcValidatorAddress, dstValidatorAddress, pageRequest)
}

// Redelegations is a free data retrieval call binding the contract method 0xe4227092.
//
// Solidity: function redelegations(string delegatorAddress, string srcValidatorAddress, string dstValidatorAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns(((string,string,string,(int64,int64,uint256,uint256,uint64,int64)[]),((int64,int64,uint256,uint256,uint64,int64),uint256)[])[] redelegationResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingCallerSession) Redelegations(delegatorAddress string, srcValidatorAddress string, dstValidatorAddress string, pageRequest PageRequest) (struct {
	RedelegationResponses []RedelegationResponse
	PaginationResult      PageResponse
}, error) {
	return _Staking.Contract.Redelegations(&_Staking.CallOpts, delegatorAddress, srcValidatorAddress, dstValidatorAddress, pageRequest)
}

// UnbondingDelegation is a free data retrieval call binding the contract method 0x20bc02be.
//
// Solidity: function unbondingDelegation(string delegatorAddr, string validatorAddr) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[]) unbond)
func (_Staking *StakingCaller) UnbondingDelegation(opts *bind.CallOpts, delegatorAddr string, validatorAddr string) (UnbondingDelegation, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "unbondingDelegation", delegatorAddr, validatorAddr)

	if err != nil {
		return *new(UnbondingDelegation), err
	}

	out0 := *abi.ConvertType(out[0], new(UnbondingDelegation)).(*UnbondingDelegation)

	return out0, err

}

// UnbondingDelegation is a free data retrieval call binding the contract method 0x20bc02be.
//
// Solidity: function unbondingDelegation(string delegatorAddr, string validatorAddr) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[]) unbond)
func (_Staking *StakingSession) UnbondingDelegation(delegatorAddr string, validatorAddr string) (UnbondingDelegation, error) {
	return _Staking.Contract.UnbondingDelegation(&_Staking.CallOpts, delegatorAddr, validatorAddr)
}

// UnbondingDelegation is a free data retrieval call binding the contract method 0x20bc02be.
//
// Solidity: function unbondingDelegation(string delegatorAddr, string validatorAddr) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[]) unbond)
func (_Staking *StakingCallerSession) UnbondingDelegation(delegatorAddr string, validatorAddr string) (UnbondingDelegation, error) {
	return _Staking.Contract.UnbondingDelegation(&_Staking.CallOpts, delegatorAddr, validatorAddr)
}

// Validator is a free data retrieval call binding the contract method 0x0bc82a17.
//
// Solidity: function validator(string validatorAddress) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[]) validator)
func (_Staking *StakingCaller) Validator(opts *bind.CallOpts, validatorAddress string) (Validator, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validator", validatorAddress)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x0bc82a17.
//
// Solidity: function validator(string validatorAddress) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[]) validator)
func (_Staking *StakingSession) Validator(validatorAddress string) (Validator, error) {
	return _Staking.Contract.Validator(&_Staking.CallOpts, validatorAddress)
}

// Validator is a free data retrieval call binding the contract method 0x0bc82a17.
//
// Solidity: function validator(string validatorAddress) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[]) validator)
func (_Staking *StakingCallerSession) Validator(validatorAddress string) (Validator, error) {
	return _Staking.Contract.Validator(&_Staking.CallOpts, validatorAddress)
}

// ValidatorDelegations is a free data retrieval call binding the contract method 0x85cadbda.
//
// Solidity: function validatorDelegations(string validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((string,string,uint256),uint256)[] delegationResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingCaller) ValidatorDelegations(opts *bind.CallOpts, validatorAddr string, pagination PageRequest) (struct {
	DelegationResponses []DelegationResponse
	PaginationResult    PageResponse
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validatorDelegations", validatorAddr, pagination)

	outstruct := new(struct {
		DelegationResponses []DelegationResponse
		PaginationResult    PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DelegationResponses = *abi.ConvertType(out[0], new([]DelegationResponse)).(*[]DelegationResponse)
	outstruct.PaginationResult = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// ValidatorDelegations is a free data retrieval call binding the contract method 0x85cadbda.
//
// Solidity: function validatorDelegations(string validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((string,string,uint256),uint256)[] delegationResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingSession) ValidatorDelegations(validatorAddr string, pagination PageRequest) (struct {
	DelegationResponses []DelegationResponse
	PaginationResult    PageResponse
}, error) {
	return _Staking.Contract.ValidatorDelegations(&_Staking.CallOpts, validatorAddr, pagination)
}

// ValidatorDelegations is a free data retrieval call binding the contract method 0x85cadbda.
//
// Solidity: function validatorDelegations(string validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((string,string,uint256),uint256)[] delegationResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingCallerSession) ValidatorDelegations(validatorAddr string, pagination PageRequest) (struct {
	DelegationResponses []DelegationResponse
	PaginationResult    PageResponse
}, error) {
	return _Staking.Contract.ValidatorDelegations(&_Staking.CallOpts, validatorAddr, pagination)
}

// ValidatorUnbondingDelegations is a free data retrieval call binding the contract method 0x1394d67f.
//
// Solidity: function validatorUnbondingDelegations(string validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[])[] unbondingResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingCaller) ValidatorUnbondingDelegations(opts *bind.CallOpts, validatorAddr string, pagination PageRequest) (struct {
	UnbondingResponses []UnbondingDelegation
	PaginationResult   PageResponse
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validatorUnbondingDelegations", validatorAddr, pagination)

	outstruct := new(struct {
		UnbondingResponses []UnbondingDelegation
		PaginationResult   PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UnbondingResponses = *abi.ConvertType(out[0], new([]UnbondingDelegation)).(*[]UnbondingDelegation)
	outstruct.PaginationResult = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// ValidatorUnbondingDelegations is a free data retrieval call binding the contract method 0x1394d67f.
//
// Solidity: function validatorUnbondingDelegations(string validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[])[] unbondingResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingSession) ValidatorUnbondingDelegations(validatorAddr string, pagination PageRequest) (struct {
	UnbondingResponses []UnbondingDelegation
	PaginationResult   PageResponse
}, error) {
	return _Staking.Contract.ValidatorUnbondingDelegations(&_Staking.CallOpts, validatorAddr, pagination)
}

// ValidatorUnbondingDelegations is a free data retrieval call binding the contract method 0x1394d67f.
//
// Solidity: function validatorUnbondingDelegations(string validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,(int64,int64,uint256,uint256,uint64,int64)[])[] unbondingResponses, (bytes,uint64) paginationResult)
func (_Staking *StakingCallerSession) ValidatorUnbondingDelegations(validatorAddr string, pagination PageRequest) (struct {
	UnbondingResponses []UnbondingDelegation
	PaginationResult   PageResponse
}, error) {
	return _Staking.Contract.ValidatorUnbondingDelegations(&_Staking.CallOpts, validatorAddr, pagination)
}

// Validators is a free data retrieval call binding the contract method 0x186b2167.
//
// Solidity: function validators(string status, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[])[] validators, (bytes,uint64) paginationResult)
func (_Staking *StakingCaller) Validators(opts *bind.CallOpts, status string, pagination PageRequest) (struct {
	Validators       []Validator
	PaginationResult PageResponse
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validators", status, pagination)

	outstruct := new(struct {
		Validators       []Validator
		PaginationResult PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]Validator)).(*[]Validator)
	outstruct.PaginationResult = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0x186b2167.
//
// Solidity: function validators(string status, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[])[] validators, (bytes,uint64) paginationResult)
func (_Staking *StakingSession) Validators(status string, pagination PageRequest) (struct {
	Validators       []Validator
	PaginationResult PageResponse
}, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, status, pagination)
}

// Validators is a free data retrieval call binding the contract method 0x186b2167.
//
// Solidity: function validators(string status, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),uint256),uint256,int64,uint64[])[] validators, (bytes,uint64) paginationResult)
func (_Staking *StakingCallerSession) Validators(status string, pagination PageRequest) (struct {
	Validators       []Validator
	PaginationResult PageResponse
}, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, status, pagination)
}

// BeginRedelegate is a paid mutator transaction binding the contract method 0x2e436cf2.
//
// Solidity: function beginRedelegate(string validatorSrcAddress, string validatorDstAddress, uint256 amount) returns(uint256 completionTime)
func (_Staking *StakingTransactor) BeginRedelegate(opts *bind.TransactOpts, validatorSrcAddress string, validatorDstAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "beginRedelegate", validatorSrcAddress, validatorDstAddress, amount)
}

// BeginRedelegate is a paid mutator transaction binding the contract method 0x2e436cf2.
//
// Solidity: function beginRedelegate(string validatorSrcAddress, string validatorDstAddress, uint256 amount) returns(uint256 completionTime)
func (_Staking *StakingSession) BeginRedelegate(validatorSrcAddress string, validatorDstAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.BeginRedelegate(&_Staking.TransactOpts, validatorSrcAddress, validatorDstAddress, amount)
}

// BeginRedelegate is a paid mutator transaction binding the contract method 0x2e436cf2.
//
// Solidity: function beginRedelegate(string validatorSrcAddress, string validatorDstAddress, uint256 amount) returns(uint256 completionTime)
func (_Staking *StakingTransactorSession) BeginRedelegate(validatorSrcAddress string, validatorDstAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.BeginRedelegate(&_Staking.TransactOpts, validatorSrcAddress, validatorDstAddress, amount)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x43970281.
//
// Solidity: function cancelUnbondingDelegation(string validatorAddress, uint256 amount, uint256 creationHeight) returns()
func (_Staking *StakingTransactor) CancelUnbondingDelegation(opts *bind.TransactOpts, validatorAddress string, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "cancelUnbondingDelegation", validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x43970281.
//
// Solidity: function cancelUnbondingDelegation(string validatorAddress, uint256 amount, uint256 creationHeight) returns()
func (_Staking *StakingSession) CancelUnbondingDelegation(validatorAddress string, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CancelUnbondingDelegation(&_Staking.TransactOpts, validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x43970281.
//
// Solidity: function cancelUnbondingDelegation(string validatorAddress, uint256 amount, uint256 creationHeight) returns()
func (_Staking *StakingTransactorSession) CancelUnbondingDelegation(validatorAddress string, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CancelUnbondingDelegation(&_Staking.TransactOpts, validatorAddress, amount, creationHeight)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x3adeba1e.
//
// Solidity: function createValidator((string,string,string,string,string) description, (uint256,uint256,uint256) commission, uint256 minSelfDelegation, string pubkey, uint256 value) returns()
func (_Staking *StakingTransactor) CreateValidator(opts *bind.TransactOpts, description Description, commission CommissionRates, minSelfDelegation *big.Int, pubkey string, value *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "createValidator", description, commission, minSelfDelegation, pubkey, value)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x3adeba1e.
//
// Solidity: function createValidator((string,string,string,string,string) description, (uint256,uint256,uint256) commission, uint256 minSelfDelegation, string pubkey, uint256 value) returns()
func (_Staking *StakingSession) CreateValidator(description Description, commission CommissionRates, minSelfDelegation *big.Int, pubkey string, value *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator(&_Staking.TransactOpts, description, commission, minSelfDelegation, pubkey, value)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x3adeba1e.
//
// Solidity: function createValidator((string,string,string,string,string) description, (uint256,uint256,uint256) commission, uint256 minSelfDelegation, string pubkey, uint256 value) returns()
func (_Staking *StakingTransactorSession) CreateValidator(description Description, commission CommissionRates, minSelfDelegation *big.Int, pubkey string, value *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator(&_Staking.TransactOpts, description, commission, minSelfDelegation, pubkey, value)
}

// Delegate is a paid mutator transaction binding the contract method 0x03f24de1.
//
// Solidity: function delegate(string validatorAddress, uint256 amount) returns(bool success)
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x03f24de1.
//
// Solidity: function delegate(string validatorAddress, uint256 amount) returns(bool success)
func (_Staking *StakingSession) Delegate(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x03f24de1.
//
// Solidity: function delegate(string validatorAddress, uint256 amount) returns(bool success)
func (_Staking *StakingTransactorSession) Delegate(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, validatorAddress, amount)
}

// EditValidator is a paid mutator transaction binding the contract method 0x34397fbb.
//
// Solidity: function editValidator((string,string,string,string,string) description, (bool,uint256) commissionRate, (bool,uint256) minSelfDelegation) returns()
func (_Staking *StakingTransactor) EditValidator(opts *bind.TransactOpts, description Description, commissionRate NullableUint, minSelfDelegation NullableUint) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "editValidator", description, commissionRate, minSelfDelegation)
}

// EditValidator is a paid mutator transaction binding the contract method 0x34397fbb.
//
// Solidity: function editValidator((string,string,string,string,string) description, (bool,uint256) commissionRate, (bool,uint256) minSelfDelegation) returns()
func (_Staking *StakingSession) EditValidator(description Description, commissionRate NullableUint, minSelfDelegation NullableUint) (*types.Transaction, error) {
	return _Staking.Contract.EditValidator(&_Staking.TransactOpts, description, commissionRate, minSelfDelegation)
}

// EditValidator is a paid mutator transaction binding the contract method 0x34397fbb.
//
// Solidity: function editValidator((string,string,string,string,string) description, (bool,uint256) commissionRate, (bool,uint256) minSelfDelegation) returns()
func (_Staking *StakingTransactorSession) EditValidator(description Description, commissionRate NullableUint, minSelfDelegation NullableUint) (*types.Transaction, error) {
	return _Staking.Contract.EditValidator(&_Staking.TransactOpts, description, commissionRate, minSelfDelegation)
}

// Undelegate is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string validatorAddress, uint256 amount) returns(uint256 completionTime)
func (_Staking *StakingTransactor) Undelegate(opts *bind.TransactOpts, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "undelegate", validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string validatorAddress, uint256 amount) returns(uint256 completionTime)
func (_Staking *StakingSession) Undelegate(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string validatorAddress, uint256 amount) returns(uint256 completionTime)
func (_Staking *StakingTransactorSession) Undelegate(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, validatorAddress, amount)
}
