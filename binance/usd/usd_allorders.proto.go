// 获取历史所有订单
package usd

import (
	"context"

	"github.com/ZiHengLee/eclient/binance"
)

const UsdAllOrdersMethod = "GET"

const UsdAllOrdersUrl = "https://fapi.binance.com/fapi/v1/allOrders"

/*
1.查询时间范围最大不得超过7天
2.默认查询最近7天内的数据
*/
type UsdAllOrdersParam struct {
	binance.BaseParam
	Symbol    string `url:"symbol"`
	Limit     int64  `url:"orderId,omitempty"` //默认 500; 最大 1000.
	Interval  string `url:"interval,omitempty"`
	StartTime int64  `url:"startTime,omitempty"` //起始时间
	EndTime   int64  `url:"endTime,omitempty"`   //结束时间
	OrderId   int64  `url:"orderId,omitempty"`   //只返回此orderID及之后的订单，缺省返回最近的订单
}

func (k UsdAllOrdersParam) RequireAuth() bool {
	return true
}

func (k *UsdAllOrdersParam) Prepare(ctx context.Context) (err error) {
	return
}

type UsdAllOrdersResp struct {
	BaseResp
	Orders []Order `json:"Orders"`
}

type Order struct {
	AvgPrice                string `json:"avgPrice"`
	ClientOrderId           string `json:"clientOrderId"`
	CumQuote                string `json:"cumQuote"`
	ExecutedQty             string `json:"executedQty"`
	OrderId                 int64  `json:"orderId"`
	OrigQty                 string `json:"origQty"`
	OrigType                string `json:"origType"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	Status                  string `json:"status"`
	StopPrice               string `json:"stopPrice"`
	ClosePosition           bool   `json:"closePosition"`
	Symbol                  string `json:"symbol"`
	Time                    int64  `json:"time"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	UpdateTime              int64  `json:"updateTime"`
	WorkingType             string `json:"workingType"`
	PriceProtect            bool   `json:"priceProtect"`
	PriceMatch              string `json:"priceMatch"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int64  `json:"goodTillDate"`
}

func (k *UsdAllOrdersResp) GetData() interface{} {
	return &k.Orders
}

// [
//   {
//     "avgPrice": "0.00000",              // 平均成交价
//     "clientOrderId": "abc",             // 用户自定义的订单号
//     "cumQuote": "0",                    // 成交金额
//     "executedQty": "0",                 // 成交量
//     "orderId": 1917641,                 // 系统订单号
//     "origQty": "0.40",                  // 原始委托数量
//     "origType": "TRAILING_STOP_MARKET", // 触发前订单类型
//     "price": "0",                       // 委托价格
//     "reduceOnly": false,                // 是否仅减仓
//     "side": "BUY",                      // 买卖方向
//     "positionSide": "SHORT",            // 持仓方向
//     "status": "NEW",                    // 订单状态
//     "stopPrice": "9300",                // 触发价，对`TRAILING_STOP_MARKET`无效
//     "closePosition": false,             // 是否条件全平仓
//     "symbol": "BTCUSDT",                // 交易对
//     "time": 1579276756075,              // 订单时间
//     "timeInForce": "GTC",               // 有效方法
//     "type": "TRAILING_STOP_MARKET",     // 订单类型
//     "activatePrice": "9020", // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "priceRate": "0.3", // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "updateTime": 1579276756075,       // 更新时间
//     "workingType": "CONTRACT_PRICE",   // 条件价格触发类型
//     "priceProtect": false,             // 是否开启条件单触发保护
//     "priceMatch": "NONE",              //盘口价格下单模式
//     "selfTradePreventionMode": "NONE", //订单自成交保护模式
//     "goodTillDate": 0      //订单TIF为GTD时的自动取消时间
//   }
// ]
