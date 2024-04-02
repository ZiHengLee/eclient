package binance

import (
	"context"

	"github.com/slowly-richer/eclient/utils/number"
)

const DepthMethod = "GET"

const DepthUrl = "https://api.binance.com/api/v3/depth"

type DepthParam struct {
	BaseParam
	Symbol string `url:"symbol"`
	Limit  int64  `url:"limit,omitempty"`
}

func (p *DepthParam) RequireAuth() bool {
	return false
}

func (p *DepthParam) Prepare(ctx context.Context) (err error) {
	return
}

type DepthInfo struct {
	LastUpdateId int64              `json:"lastUpdateId"`
	Bids         [][2]number.Number `json:"bids"`
	Asks         [][2]number.Number `json:"asks"`
}

type DepthResp struct {
	BaseResp

	DepthInfo
}
