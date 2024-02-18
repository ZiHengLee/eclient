package binance

import (
	"context"
)

const Ticker24hrMethod = "GET"

const Ticker24hrUrl = "https://api.binance.com/api/v3/ticker/24hr"

type Ticker24hrParam struct {
	BaseParam
	Symbol string `url:"symbol,omitempty"` //交易对
}

func (p *Ticker24hrParam) RequireAuth() bool {
	return false
}

func (p *Ticker24hrParam) Prepare(ctx context.Context) (err error) {
	return
}

type Ticker24hrInfo struct {
	Symbol             string `json:"symbol"`             //交易对
	PriceChange        string `json:"priceChange"`        //"-94.99999800",
	PriceChangePercent string `json:"priceChangePercent"` //"-95.960",
	WeightedAvgPrice   string `json:"weightedAvgPrice"`   //"0.29628482",
	PrevClosePrice     string `json:"prevClosePrice"`     //"0.10002000",
	LastPrice          string `json:"lastPrice"`          //"4.00000200",
	LastQty            string `json:"lastQty"`            //"200.00000000",
	BidPrice           string `json:"bidPrice"`           //"4.00000000",
	BidQty             string `json:"bidQty"`             //"100.00000000",
	AskPrice           string `json:"askPrice"`           //"4.00000200",
	AskQty             string `json:"askQty"`             //"100.00000000",
	OpenPrice          string `json:"openPrice"`          //"99.00000000",
	HighPrice          string `json:"highPrice"`          //"100.00000000",
	LowPrice           string `json:"lowPrice"`           //"0.10000000",
	Volume             string `json:"volume"`             //"8913.30000000",
	QuoteVolume        string `json:"quoteVolume"`        //"15.30000000",
	OpenTime           int64  `json:"openTime"`           //1499783499040,
	CloseTime          int64  `json:"closeTime"`          //1499869899040,
	FirstId            int64  `json:"firstId"`            //28385,   // 首笔成交id
	LastId             int64  `json:"lastId"`             //28460,    // 末笔成交id
	Count              int64  `json:"count"`              //76         // 成交笔数
}

type Ticker24hrResp struct {
	BaseResp
	Ticker24hrInfo
	Data []*Ticker24hrInfo
}

func (r *Ticker24hrResp) GetData() interface{} {
	return &r.Data
}
