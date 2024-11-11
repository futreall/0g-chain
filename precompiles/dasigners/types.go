package dasigners

import (
	"fmt"
	"math/big"

	precopmiles_common "github.com/0glabs/0g-chain/precompiles/common"
	dasignerstypes "github.com/0glabs/0g-chain/x/dasigners/v1/types"
	"github.com/ethereum/go-ethereum/common"
)

type BN254G1Point = struct {
	X *big.Int "json:\"X\""
	Y *big.Int "json:\"Y\""
}

type BN254G2Point = struct {
	X [2]*big.Int "json:\"X\""
	Y [2]*big.Int "json:\"Y\""
}

type IDASignersSignerDetail = struct {
	Signer common.Address "json:\"signer\""
	Socket string         "json:\"socket\""
	PkG1   BN254G1Point   "json:\"pkG1\""
	PkG2   BN254G2Point   "json:\"pkG2\""
}

type IDASignersParams = struct {
	TokensPerVote     *big.Int "json:\"tokensPerVote\""
	MaxVotesPerSigner *big.Int "json:\"maxVotesPerSigner\""
	MaxQuorums        *big.Int "json:\"maxQuorums\""
	EpochBlocks       *big.Int "json:\"epochBlocks\""
	EncodedSlices     *big.Int "json:\"encodedSlices\""
}

func NewBN254G1Point(b []byte) BN254G1Point {
	return BN254G1Point{
		X: new(big.Int).SetBytes(b[:32]),
		Y: new(big.Int).SetBytes(b[32:64]),
	}
}

func SerializeG1(p BN254G1Point) []byte {
	b := make([]byte, 0)
	b = append(b, common.LeftPadBytes(p.X.Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(p.Y.Bytes(), 32)...)
	return b
}

func NewBN254G2Point(b []byte) BN254G2Point {
	return BN254G2Point{
		X: [2]*big.Int{
			new(big.Int).SetBytes(b[:32]),
			new(big.Int).SetBytes(b[32:64]),
		},
		Y: [2]*big.Int{
			new(big.Int).SetBytes(b[64:96]),
			new(big.Int).SetBytes(b[96:128]),
		},
	}
}

func SerializeG2(p BN254G2Point) []byte {
	b := make([]byte, 0)
	b = append(b, common.LeftPadBytes(p.X[0].Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(p.X[1].Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(p.Y[0].Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(p.Y[1].Bytes(), 32)...)
	return b
}

func NewQueryQuorumCountRequest(args []interface{}) (*dasignerstypes.QueryQuorumCountRequest, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precopmiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}

	return &dasignerstypes.QueryQuorumCountRequest{
		EpochNumber: args[0].(*big.Int).Uint64(),
	}, nil
}

func NewQuerySignerRequest(args []interface{}) (*dasignerstypes.QuerySignerRequest, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precopmiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}
	accounts := args[0].([]common.Address)
	req := dasignerstypes.QuerySignerRequest{
		Accounts: make([]string, len(accounts)),
	}
	for i, account := range accounts {
		req.Accounts[i] = precopmiles_common.ToLowerHexWithoutPrefix(account)
	}
	return &req, nil
}

func NewQueryEpochQuorumRequest(args []interface{}) (*dasignerstypes.QueryEpochQuorumRequest, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf(precopmiles_common.ErrInvalidNumberOfArgs, 2, len(args))
	}

	return &dasignerstypes.QueryEpochQuorumRequest{
		EpochNumber: args[0].(*big.Int).Uint64(),
		QuorumId:    args[1].(*big.Int).Uint64(),
	}, nil
}

func NewQueryEpochQuorumRowRequest(args []interface{}) (*dasignerstypes.QueryEpochQuorumRowRequest, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf(precopmiles_common.ErrInvalidNumberOfArgs, 3, len(args))
	}

	return &dasignerstypes.QueryEpochQuorumRowRequest{
		EpochNumber: args[0].(*big.Int).Uint64(),
		QuorumId:    args[1].(*big.Int).Uint64(),
		RowIndex:    args[2].(uint32),
	}, nil
}

func NewQueryAggregatePubkeyG1Request(args []interface{}) (*dasignerstypes.QueryAggregatePubkeyG1Request, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf(precopmiles_common.ErrInvalidNumberOfArgs, 3, len(args))
	}

	return &dasignerstypes.QueryAggregatePubkeyG1Request{
		EpochNumber:  args[0].(*big.Int).Uint64(),
		QuorumId:     args[1].(*big.Int).Uint64(),
		QuorumBitmap: args[2].([]byte),
	}, nil
}

func NewIDASignersSignerDetail(signer *dasignerstypes.Signer) IDASignersSignerDetail {
	return IDASignersSignerDetail{
		Signer: common.HexToAddress(signer.Account),
		Socket: signer.Socket,
		PkG1:   NewBN254G1Point(signer.PubkeyG1),
		PkG2:   NewBN254G2Point(signer.PubkeyG2),
	}
}

func NewMsgRegisterSigner(args []interface{}) (*dasignerstypes.MsgRegisterSigner, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf(precopmiles_common.ErrInvalidNumberOfArgs, 2, len(args))
	}

	signer := args[0].(IDASignersSignerDetail)
	return &dasignerstypes.MsgRegisterSigner{
		Signer: &dasignerstypes.Signer{
			Account:  precopmiles_common.ToLowerHexWithoutPrefix(signer.Signer),
			Socket:   signer.Socket,
			PubkeyG1: SerializeG1(signer.PkG1),
			PubkeyG2: SerializeG2(signer.PkG2),
		},
		Signature: SerializeG1(args[1].(BN254G1Point)),
	}, nil
}

func NewMsgRegisterNextEpoch(args []interface{}, account string) (*dasignerstypes.MsgRegisterNextEpoch, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precopmiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}

	return &dasignerstypes.MsgRegisterNextEpoch{
		Account:   account,
		Signature: SerializeG1(args[0].(BN254G1Point)),
	}, nil
}

func NewMsgUpdateSocket(args []interface{}, account string) (*dasignerstypes.MsgUpdateSocket, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(precopmiles_common.ErrInvalidNumberOfArgs, 1, len(args))
	}

	return &dasignerstypes.MsgUpdateSocket{
		Account: account,
		Socket:  args[0].(string),
	}, nil
}
