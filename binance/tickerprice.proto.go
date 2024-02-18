package binance

import (
	"context"
)

const TickerPriceMethod = "GET"

const TickerPriceUrl = "https://api.binance.com/api/v3/ticker/price"

type TickerPriceParam struct {
	BaseParam
	Symbol string `url:"symbol,omitempty"` //交易对
}

func (p *TickerPriceParam) RequireAuth() bool {
	return false
}

func (p *TickerPriceParam) Prepare(ctx context.Context) (err error) {
	return
}

type TickerPriceInfo struct {
	Symbol string `json:"symbol"` //交易对
	Price  string `json:"price"`  // "4.00000100", // 成交价格
}

type TickerPriceResp struct {
	BaseResp
	Data []*TickerPriceInfo `json:"data"`
}

func (r *TickerPriceResp) GetData() interface{} {
	return &r.Data
}
