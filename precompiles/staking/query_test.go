package staking_test

import (
	stakingprecompile "github.com/0glabs/0g-chain/precompiles/staking"
	sdk "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
)

func (s *StakingTestSuite) TestValidators() {
	method := stakingprecompile.StakingFunctionValidators

	testCases := []struct {
		name        string
		malleate    func() []byte
		postCheck   func(bz []byte)
		gas         uint64
		expErr      bool
		errContains string
	}{
		{
			"success",
			func() []byte {
				input, err := s.abi.Pack(
					method,
					"",
					query.PageRequest{
						Limit:      10,
						CountTotal: true,
					},
				)
				s.Assert().NoError(err)
				return input
			},
			func(data []byte) {
				out, err := s.abi.Methods[method].Outputs.Unpack(data)
				s.Require().NoError(err, "failed to unpack output")
				validators := out[0].([]stakingprecompile.Validator)
				paginationResult := out[1].(stakingprecompile.PageResponse)
				s.Assert().EqualValues(3, len(validators))
				s.Assert().EqualValues(3, paginationResult.Total)
			},
			100000,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupTest()
			s.AddDelegation(s.signerOne.HexAddr, s.signerTwo.HexAddr, sdk.NewIntFromUint64(1000000))

			bz, err := s.runTx(tc.malleate(), s.signerOne, 10000000)

			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.errContains)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(bz)
				tc.postCheck(bz)
			}
		})
	}
}
