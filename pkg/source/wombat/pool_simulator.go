package wombat

import (
	"encoding/json"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
)

type Pool struct {
	pool.Pool
	assets map[string]Asset
	gas    Gas
}

func NewPoolSimulator(entityPool entity.Pool) (*Pool, error) {
	var extra Extra
	if err := json.Unmarshal([]byte(entityPool.Extra), &extra); err != nil {
		return nil, err
	}

	var tokens = make([]string, len(entityPool.Tokens))
	for _, token := range entityPool.Tokens {
		tokens = append(tokens, token.Address)
	}

	return &Pool{
		Pool: pool.Pool{
			Info: pool.PoolInfo{
				Address:  entityPool.Address,
				Type:     entityPool.Type,
				Exchange: entityPool.Exchange,
				Tokens:   tokens,
				Checked:  false,
			},
		},
		assets: extra.AssetMap,
		gas:    DefaultGas,
	}, nil
}

func (p *Pool) CalcAmountOut(
	tokenAmountIn pool.TokenAmount,
	tokenOut string,
) (*pool.CalcAmountOutResult, error) {
	return nil, nil
}
