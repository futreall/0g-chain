package staking_test

import (
	"encoding/base64"
	"math/big"

	"cosmossdk.io/math"
	stakingprecompile "github.com/0glabs/0g-chain/precompiles/staking"
	"github.com/ethereum/go-ethereum/common"
)

func (s *StakingTestSuite) TestCreateValidator() {
	method := stakingprecompile.StakingFunctionCreateValidator
	description := stakingprecompile.Description{
		Moniker:         "test node",
		Identity:        "test node identity",
		Website:         "http://test.node.com",
		SecurityContact: "test node security contract",
		Details:         "test node details",
	}
	commission := stakingprecompile.CommissionRates{
		Rate:          math.LegacyOneDec().BigInt(),
		MaxRate:       math.LegacyOneDec().BigInt(),
		MaxChangeRate: math.LegacyOneDec().BigInt(),
	}
	minSelfDelegation := big.NewInt(1)
	pubkey := "eh/aR8BGUBIYI/Ust0NVBxZafLDAm7344F9dKzZU+7g="
	value := big.NewInt(100000000)

	testCases := []struct {
		name          string
		malleate      func() []byte
		gas           uint64
		callerAddress *common.Address
		postCheck     func(data []byte)
		expError      bool
		errContains   string
	}{
		{
			"fail - ErrPubKeyInvalidLength",
			func() []byte {
				input, err := s.abi.Pack(
					method,
					description,
					commission,
					minSelfDelegation,
					s.signerOne.HexAddr,
					value,
				)
				s.Assert().NoError(err)
				return input
			},
			200000,
			nil,
			func([]byte) {},
			true,
			stakingprecompile.ErrPubKeyInvalidLength,
		},
		{
			"success",
			func() []byte {
				input, err := s.abi.Pack(
					method,
					description,
					commission,
					minSelfDelegation,
					pubkey,
					value,
				)
				s.Assert().NoError(err)
				return input
			},
			200000,
			nil,
			func(data []byte) {},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()

			bz, err := s.runTx(tc.malleate(), s.signerOne, 10000000)

			if tc.expError {
				s.Require().ErrorContains(err, tc.errContains)
				s.Require().Empty(bz)
			} else {
				s.Require().NoError(err)
				// query the validator in the staking keeper
				validator := s.StakingKeeper.Validator(s.Ctx, s.signerOne.ValAddr)
				s.Require().NoError(err)

				s.Require().NotNil(validator, "expected validator not to be nil")
				tc.postCheck(bz)

				isBonded := validator.IsBonded()
				s.Require().Equal(false, isBonded, "expected validator bonded to be %t; got %t", false, isBonded)

				consPubKey, err := validator.ConsPubKey()
				s.Require().NoError(err)
				consPubKeyBase64 := base64.StdEncoding.EncodeToString(consPubKey.Bytes())
				s.Require().Equal(pubkey, consPubKeyBase64, "expected validator pubkey to be %s; got %s", pubkey, consPubKeyBase64)

				operator := validator.GetOperator()
				s.Require().Equal(s.signerOne.ValAddr, operator, "expected validator operator to be %s; got %s", s.signerOne.ValAddr, operator)

				commissionRate := validator.GetCommission()
				s.Require().Equal(commission.Rate.String(), commissionRate.BigInt().String(), "expected validator commission rate to be %s; got %s", commission.Rate.String(), commissionRate.String())

				valMinSelfDelegation := validator.GetMinSelfDelegation()
				s.Require().Equal(minSelfDelegation.String(), valMinSelfDelegation.String(), "expected validator min self delegation to be %s; got %s", minSelfDelegation.String(), valMinSelfDelegation.String())

				moniker := validator.GetMoniker()
				s.Require().Equal(description.Moniker, moniker, "expected validator moniker to be %s; got %s", description.Moniker, moniker)

				jailed := validator.IsJailed()
				s.Require().Equal(false, jailed, "expected validator jailed to be %t; got %t", false, jailed)
			}
		})
	}
}
