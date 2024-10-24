package common

import (
	"math/big"
	"strings"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
)

func ToLowerHexWithoutPrefix(addr common.Address) string {
	return strings.ToLower(addr.Hex()[2:])
}

// BigIntToLegacyDec converts a uint number (18 decimals) to math.LegacyDec (18 decimals)
func BigIntToLegacyDec(x *big.Int) math.LegacyDec {
	return math.LegacyNewDecFromBigIntWithPrec(x, math.LegacyPrecision)
}
