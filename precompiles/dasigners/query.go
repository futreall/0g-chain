package dasigners

import (
	"fmt"
	"math/big"

	precopmiles_common "github.com/0glabs/0g-chain/precompiles/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

func (d *DASignersPrecompile) Params(ctx sdk.Context, _ *vm.EVM, method *abi.Method, _ []interface{}) ([]byte, error) {
	params := d.dasignersKeeper.GetParams(ctx)
	return method.Outputs.Pack(IDASignersParams{
		TokensPerVote:     big.NewInt(int64(params.TokensPerVote)),
		MaxVotesPerSigner: big.NewInt(int64(params.MaxVotesPerSigner)),
		MaxQuorums:        big.NewInt(int64(params.MaxQuorums)),
		EpochBlocks:       big.NewInt(int64(params.EpochBlocks)),
		EncodedSlices:     big.NewInt(int64(params.EncodedSlices)),
	})
}

func (d *DASignersPrecompile) EpochNumber(ctx sdk.Context, _ *vm.EVM, method *abi.Method, _ []interface{}) ([]byte, error) {
	epochNumber, err := d.dasignersKeeper.GetEpochNumber(ctx)
	if err != nil {
		return nil, err
	}
	return method.Outputs.Pack(big.NewInt(int64(epochNumber)))
}

func (d *DASignersPrecompile) QuorumCount(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryQuorumCountRequest(args)
	if err != nil {
		return nil, err
	}
	response, err := d.dasignersKeeper.QuorumCount(ctx, req)
	if err != nil {
		return nil, err
	}
	return method.Outputs.Pack(big.NewInt(int64(response.QuorumCount)))
}

func (d *DASignersPrecompile) GetSigner(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQuerySignerRequest(args)
	if err != nil {
		return nil, err
	}
	response, err := d.dasignersKeeper.Signer(sdk.WrapSDKContext(ctx), req)
	if err != nil {
		return nil, err
	}
	signers := make([]IDASignersSignerDetail, len(response.Signer))
	for i, signer := range response.Signer {
		signers[i] = NewIDASignersSignerDetail(signer)
	}
	return method.Outputs.Pack(signers)
}

func (d *DASignersPrecompile) IsSigner(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precopmiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}
	account := precopmiles_common.ToLowerHexWithoutPrefix(args[0].(common.Address))
	_, found, err := d.dasignersKeeper.GetSigner(ctx, account)
	if err != nil {
		return nil, err
	}
	return method.Outputs.Pack(found)
}

func (d *DASignersPrecompile) RegisteredEpoch(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf(precopmiles_common.ErrInvalidNumberOfArgs, 2, len(args))
	}
	account := precopmiles_common.ToLowerHexWithoutPrefix(args[0].(common.Address))
	epoch := args[1].(*big.Int).Uint64()
	_, found, err := d.dasignersKeeper.GetRegistration(ctx, epoch, account)
	if err != nil {
		return nil, err
	}
	return method.Outputs.Pack(found)
}

func (d *DASignersPrecompile) GetQuorum(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryEpochQuorumRequest(args)
	if err != nil {
		return nil, err
	}
	response, err := d.dasignersKeeper.EpochQuorum(sdk.WrapSDKContext(ctx), req)
	if err != nil {
		return nil, err
	}
	signers := make([]common.Address, len(response.Quorum.Signers))
	for i, signer := range response.Quorum.Signers {
		signers[i] = common.HexToAddress(signer)
	}
	return method.Outputs.Pack(signers)
}

func (d *DASignersPrecompile) GetQuorumRow(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryEpochQuorumRowRequest(args)
	if err != nil {
		return nil, err
	}
	response, err := d.dasignersKeeper.EpochQuorumRow(sdk.WrapSDKContext(ctx), req)
	if err != nil {
		return nil, err
	}
	return method.Outputs.Pack(common.HexToAddress(response.Signer))
}

func (d *DASignersPrecompile) GetAggPkG1(ctx sdk.Context, _ *vm.EVM, method *abi.Method, args []interface{}) ([]byte, error) {
	req, err := NewQueryAggregatePubkeyG1Request(args)
	if err != nil {
		return nil, err
	}
	response, err := d.dasignersKeeper.AggregatePubkeyG1(sdk.WrapSDKContext(ctx), req)
	if err != nil {
		return nil, err
	}
	return method.Outputs.Pack(NewBN254G1Point(response.AggregatePubkeyG1), big.NewInt(int64(response.Total)), big.NewInt(int64(response.Hit)))
}
