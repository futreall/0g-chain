package dasigners

import (
	"fmt"
	"strings"

	precopmiles_common "github.com/0glabs/0g-chain/precompiles/common"
	dasignerskeeper "github.com/0glabs/0g-chain/x/dasigners/v1/keeper"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/x/evm/statedb"
)

const (
	PrecompileAddress = "0x0000000000000000000000000000000000001000"

	RequiredGasMax uint64 = 1000_000_000

	DASignersFunctionParams            = "params"
	DASignersFunctionEpochNumber       = "epochNumber"
	DASignersFunctionQuorumCount       = "quorumCount"
	DASignersFunctionGetSigner         = "getSigner"
	DASignersFunctionGetQuorum         = "getQuorum"
	DASignersFunctionGetQuorumRow      = "getQuorumRow"
	DASignersFunctionRegisterSigner    = "registerSigner"
	DASignersFunctionUpdateSocket      = "updateSocket"
	DASignersFunctionRegisterNextEpoch = "registerNextEpoch"
	DASignersFunctionGetAggPkG1        = "getAggPkG1"
	DASignersFunctionIsSigner          = "isSigner"
	DASignersFunctionRegisteredEpoch   = "registeredEpoch"
)

var RequiredGasBasic = map[string]uint64{
	DASignersFunctionParams:            1000,
	DASignersFunctionEpochNumber:       1000,
	DASignersFunctionQuorumCount:       1000,
	DASignersFunctionGetSigner:         100000,
	DASignersFunctionGetQuorum:         100000,
	DASignersFunctionGetQuorumRow:      10000,
	DASignersFunctionRegisterSigner:    100000,
	DASignersFunctionUpdateSocket:      50000,
	DASignersFunctionRegisterNextEpoch: 100000,
	DASignersFunctionGetAggPkG1:        1000000,
	DASignersFunctionIsSigner:          10000,
	DASignersFunctionRegisteredEpoch:   10000,
}

var KVGasConfig storetypes.GasConfig = storetypes.GasConfig{
	HasCost:          0,
	DeleteCost:       0,
	ReadCostFlat:     0,
	ReadCostPerByte:  0,
	WriteCostFlat:    0,
	WriteCostPerByte: 0,
	IterNextCostFlat: 0,
}

var _ vm.PrecompiledContract = &DASignersPrecompile{}

type DASignersPrecompile struct {
	abi             abi.ABI
	dasignersKeeper dasignerskeeper.Keeper
}

func NewDASignersPrecompile(dasignersKeeper dasignerskeeper.Keeper) (*DASignersPrecompile, error) {
	abi, err := abi.JSON(strings.NewReader(DASignersABI))
	if err != nil {
		return nil, err
	}
	return &DASignersPrecompile{
		abi:             abi,
		dasignersKeeper: dasignersKeeper,
	}, nil
}

// Address implements vm.PrecompiledContract.
func (d *DASignersPrecompile) Address() common.Address {
	return common.HexToAddress(PrecompileAddress)
}

// RequiredGas implements vm.PrecompiledContract.
func (d *DASignersPrecompile) RequiredGas(input []byte) uint64 {
	method, err := d.abi.MethodById(input[:4])
	if err != nil {
		return RequiredGasMax
	}
	if gas, ok := RequiredGasBasic[method.Name]; ok {
		return gas
	}
	return RequiredGasMax
}

// Run implements vm.PrecompiledContract.
func (d *DASignersPrecompile) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	// parse input
	if len(contract.Input) < 4 {
		return nil, vm.ErrExecutionReverted
	}
	method, err := d.abi.MethodById(contract.Input[:4])
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
	ctx = ctx.WithKVGasConfig(KVGasConfig)
	initialGas := ctx.GasMeter().GasConsumed()

	var bz []byte
	switch method.Name {
	// queries
	case DASignersFunctionParams:
		bz, err = d.Params(ctx, evm, method, args)
	case DASignersFunctionEpochNumber:
		bz, err = d.EpochNumber(ctx, evm, method, args)
	case DASignersFunctionQuorumCount:
		bz, err = d.QuorumCount(ctx, evm, method, args)
	case DASignersFunctionGetSigner:
		bz, err = d.GetSigner(ctx, evm, method, args)
	case DASignersFunctionGetQuorum:
		bz, err = d.GetQuorum(ctx, evm, method, args)
	case DASignersFunctionGetQuorumRow:
		bz, err = d.GetQuorumRow(ctx, evm, method, args)
	case DASignersFunctionGetAggPkG1:
		bz, err = d.GetAggPkG1(ctx, evm, method, args)
	case DASignersFunctionIsSigner:
		bz, err = d.IsSigner(ctx, evm, method, args)
	case DASignersFunctionRegisteredEpoch:
		bz, err = d.RegisteredEpoch(ctx, evm, method, args)
	// txs
	case DASignersFunctionRegisterSigner:
		bz, err = d.RegisterSigner(ctx, evm, stateDB, contract, method, args)
	case DASignersFunctionRegisterNextEpoch:
		bz, err = d.RegisterNextEpoch(ctx, evm, stateDB, contract, method, args)
	case DASignersFunctionUpdateSocket:
		bz, err = d.UpdateSocket(ctx, evm, stateDB, contract, method, args)
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
