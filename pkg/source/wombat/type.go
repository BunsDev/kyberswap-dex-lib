package wombat

import "math/big"

type Extra struct {
	HaircutRate *big.Int
	AssetMap    map[string]Asset
}

type Asset struct {
	Cash                    *big.Int
	Liability               *big.Int
	UnderlyingTokenDecimals uint8
}

type Gas struct {
	Swap int64
}
