package wombat

import "math/big"

var (
	DefaultGas = Gas{Swap: 125000}
	WAD        = big.NewInt(1e18)
	WADI       = big.NewInt(1e18)
)
