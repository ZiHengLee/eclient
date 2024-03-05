// 获取历史k线数据
package binance

import (
	"context"
)

const UsdKlineMethod = "GET"

const UsdKlineUrl = "https://fapi.binance.com/fapi/v1/klines"

type UsdKlineParam struct {
	BaseParam
	Symbol    string `url:"symbol"`
	Limit     int64  `url:"limit,omitempty"` //默认 500; 最大 1000.
	Interval  string `url:"interval,omitempty"`
	StartTime int64  `url:"startTime,omitempty"`
	EndTime   int64  `url:"endTime,omitempty"`
}

func (k UsdKlineParam) RequireAuth() bool {
	return false
}

func (k *UsdKlineParam) Prepare(ctx context.Context) (err error) {
	return
}

type UsdKlineResp struct {
	BaseResp
	Klines [][]interface{} `json:"Klines"`
}

func (k *UsdKlineResp) GetData() interface{} {
	return &k.Klines
}

// [
//   [
//     1499040000000,      // k线开盘时间
//     "0.01634790",       // 开盘价
//     "0.80000000",       // 最高价
//     "0.01575800",       // 最低价
//     "0.01577100",       // 收盘价(当前K线未结束的即为最新价)
//     "148976.11427815",  // 成交量
//     1499644799999,      // k线收盘时间
//     "2434.19055334",    // 成交额
//     308,                // 成交笔数
//     "1756.87402397",    // 主动买入成交量
//     "28.46694368",      // 主动买入成交额
//     "0" // 请忽略该参数
//   ]
// ]
