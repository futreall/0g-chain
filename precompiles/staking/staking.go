package staking

import (
	"fmt"
	"strings"

	precopmiles_common "github.com/0glabs/0g-chain/precompiles/common"
	"github.com/cosmos/cosmos-sdk/store/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/x/evm/statedb"
)

const (
	PrecompileAddress = "0x0000000000000000000000000000000000001001"

	// txs
	StakingFunctionCreateValidator           = "createValidator"
	StakingFunctionEditValidator             = "editValidator"
	StakingFunctionDelegate                  = "delegate"
	StakingFunctionBeginRedelegate           = "beginRedelegate"
	StakingFunctionUndelegate                = "undelegate"
	StakingFunctionCancelUnbondingDelegation = "cancelUnbondingDelegation"
	// queries
	StakingFunctionValidators                    = "validators"
	StakingFunctionValidator                     = "validator"
	StakingFunctionValidatorDelegations          = "validatorDelegations"
	StakingFunctionValidatorUnbondingDelegations = "validatorUnbondingDelegations"
	StakingFunctionDelegation                    = "delegation"
	StakingFunctionUnbondingDelegation           = "unbondingDelegation"
	StakingFunctionDelegatorDelegations          = "delegatorDelegations"
	StakingFunctionDelegatorUnbondingDelegations = "delegatorUnbondingDelegations"
	StakingFunctionRedelegations                 = "redelegations"
	StakingFunctionDelegatorValidators           = "delegatorValidators"
	StakingFunctionDelegatorValidator            = "delegatorValidator"
	StakingFunctionPool                          = "pool"
	StakingFunctionParams                        = "params"
)

var _ vm.PrecompiledContract = &StakingPrecompile{}

type StakingPrecompile struct {
	abi           abi.ABI
	stakingKeeper *stakingkeeper.Keeper
}

func NewStakingPrecompile(stakingKeeper *stakingkeeper.Keeper) (*StakingPrecompile, error) {
	abi, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return &StakingPrecompile{
		abi:           abi,
		stakingKeeper: stakingKeeper,
	}, nil
}

// Address implements vm.PrecompiledContract.
func (s *StakingPrecompile) Address() common.Address {
	return common.HexToAddress(PrecompileAddress)
}

// RequiredGas implements vm.PrecompiledContract.
func (s *StakingPrecompile) RequiredGas(input []byte) uint64 {
	return 0
}

// Run implements vm.PrecompiledContract.
func (s *StakingPrecompile) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	// parse input
	if len(contract.Input) < 4 {
		return nil, vm.ErrExecutionReverted
	}
	method, err := s.abi.MethodById(contract.Input[:4])
	if err != nil {
		return nil, vm.ErrExecutionReverted
	}
	args, err := method.Inputs.Unpack(contract.Input[4:])
	if err != nil {
		return nil, err
	}
	// get state db and context
	stateDB, ok := evm.StateDB.(*statedb.StateDB)
	if !ok {
		return nil, fmt.Errorf(precopmiles_common.ErrGetStateDB)
	}
	ctx := stateDB.GetContext()
	// reset gas config
	ctx = ctx.WithKVGasConfig(types.KVGasConfig())
	initialGas := ctx.GasMeter().GasConsumed()

	var bz []byte
	switch method.Name {
	// queries
	case StakingFunctionValidators:
		bz, err = s.Validators(ctx, evm, method, args)
	case StakingFunctionValidator:
		bz, err = s.Validator(ctx, evm, method, args)
	case StakingFunctionValidatorDelegations:
		bz, err = s.ValidatorDelegations(ctx, evm, method, args)
	case StakingFunctionValidatorUnbondingDelegations:
		bz, err = s.ValidatorUnbondingDelegations(ctx, evm, method, args)
	case StakingFunctionDelegation:
		bz, err = s.Delegation(ctx, evm, method, args)
	case StakingFunctionUnbondingDelegation:
		bz, err = s.UnbondingDelegation(ctx, evm, method, args)
	case StakingFunctionDelegatorDelegations:
		bz, err = s.DelegatorDelegations(ctx, evm, method, args)
	case StakingFunctionDelegatorUnbondingDelegations:
		bz, err = s.DelegatorUnbondingDelegations(ctx, evm, method, args)
	case StakingFunctionRedelegations:
		bz, err = s.Redelegations(ctx, evm, method, args)
	case StakingFunctionDelegatorValidators:
		bz, err = s.DelegatorValidators(ctx, evm, method, args)
	case StakingFunctionDelegatorValidator:
		bz, err = s.DelegatorValidator(ctx, evm, method, args)
	case StakingFunctionPool:
		bz, err = s.Pool(ctx, evm, method, args)
	case StakingFunctionParams:
		bz, err = s.Params(ctx, evm, method, args)
	// txs
	case StakingFunctionCreateValidator:
		bz, err = s.CreateValidator(ctx, evm, stateDB, contract, method, args)
	case StakingFunctionEditValidator:
		bz, err = s.EditValidator(ctx, evm, stateDB, contract, method, args)
	case StakingFunctionDelegate:
		bz, err = s.Delegate(ctx, evm, stateDB, contract, method, args)
	case StakingFunctionBeginRedelegate:
		bz, err = s.BeginRedelegate(ctx, evm, stateDB, contract, method, args)
	case StakingFunctionUndelegate:
		bz, err = s.Undelegate(ctx, evm, stateDB, contract, method, args)
	case StakingFunctionCancelUnbondingDelegation:
		bz, err = s.CancelUnbondingDelegation(ctx, evm, stateDB, contract, method, args)
	}

	if err != nil {
		return nil, err
	}

	cost := ctx.GasMeter().GasConsumed() - initialGas

	if !contract.UseGas(cost) {
		return nil, vm.ErrOutOfGas
	}
	return bz, nil
}
