// 下单
package usd

import (
	"context"

	"github.com/ZiHengLee/eclient/binance"
)

const UsdOrderMethod = "POST"

const UsdOrderUrl = "https://fapi.binance.com/fapi/v1/order"

/*
条件单的触发必须:

如果订单参数priceProtect为true:
达到触发价时，MARK_PRICE(标记价格)与CONTRACT_PRICE(合约最新价)之间的价差不能超过改symbol触发保护阈值
触发保护阈值请参考接口GET /fapi/v1/exchangeInfo 返回内容相应symbol中"triggerProtect"字段
STOP, STOP_MARKET 止损单:
买入: 最新合约价格/标记价格高于等于触发价stopPrice
卖出: 最新合约价格/标记价格低于等于触发价stopPrice
TAKE_PROFIT, TAKE_PROFIT_MARKET 止盈单:
买入: 最新合约价格/标记价格低于等于触发价stopPrice
卖出: 最新合约价格/标记价格高于等于触发价stopPrice
TRAILING_STOP_MARKET 跟踪止损单:
买入: 当合约价格/标记价格区间最低价格低于激活价格activationPrice,且最新合约价格/标记价高于等于最低价设定回调幅度。
卖出: 当合约价格/标记价格区间最高价格高于激活价格activationPrice,且最新合约价格/标记价低于等于最高价设定回调幅度。
TRAILING_STOP_MARKET 跟踪止损单如果遇到报错 {"code": -2021, "msg": "Order would immediately trigger."}
表示订单不满足以下条件:

买入: 指定的activationPrice 必须小于 latest price
卖出: 指定的activationPrice 必须大于 latest price
newOrderRespType 如果传 RESULT:

MARKET 订单将直接返回成交结果；
配合使用特殊 timeInForce 的 LIMIT 订单将直接返回成交或过期拒绝结果。
STOP_MARKET, TAKE_PROFIT_MARKET 配合 closePosition=true:

条件单触发依照上述条件单触发逻辑
条件触发后，平掉当时持有所有多头仓位(若为卖单)或当时持有所有空头仓位(若为买单)
不支持 quantity 参数
自带只平仓属性，不支持reduceOnly参数
双开模式下,LONG方向上不支持BUY; SHORT 方向上不支持SELL
selfTradePreventionMode 仅在 timeInForce为IOC或GTC或GTD时生效.

极端行情时，timeInForce为GTD的订单自动取消可能有一定延迟
*/
type UsdOrderParam struct {
	binance.BaseParam
	Symbol              string `json:"symbol"`                        // 交易对
	Side                string `json:"side"`                          // 买卖方向 SELL, BUY
	PositionSide        string `json:"positionSide,omitempty"`        // 持仓方向
	Type                string `json:"type"`                          // 订单类型 LIMIT, MARKET, STOP, TAKE_PROFIT, STOP_MARKET, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
	ReduceOnly          bool   `json:"reduceOnly,omitempty"`          // 是否仅减仓
	Quantity            string `json:"quantity,omitempty"`            // 下单数量
	Price               string `json:"price,omitempty"`               // 委托价格
	NewClientOrderId    string `json:"newClientOrderId,omitempty"`    // 用户自定义的订单号
	StopPrice           string `json:"stopPrice,omitempty"`           // 触发价
	ClosePosition       bool   `json:"closePosition,omitempty"`       // 触发后全部平仓
	ActivationPrice     string `json:"activationPrice,omitempty"`     // 追踪止损激活价格
	CallbackRate        string `json:"callbackRate,omitempty"`        // 追踪止损回调比例
	TimeInForce         string `json:"timeInForce,omitempty"`         // 有效方法
	WorkingType         string `json:"workingType,omitempty"`         // stopPrice 触发类型
	PriceProtect        string `json:"priceProtect,omitempty"`        // 条件单触发保护
	NewOrderRespType    string `json:"newOrderRespType,omitempty"`    // 返回类型
	PriceMatch          string `json:"priceMatch,omitempty"`          // 价格匹配
	SelfTradePrevention string `json:"selfTradePrevention,omitempty"` // 自成交保护模式
	GoodTillDate        int64  `json:"goodTillDate,omitempty"`        // GTD时订单的自动取消时间
}

func (k UsdOrderParam) RequireAuth() bool {
	return true
}

func (k *UsdOrderParam) Prepare(ctx context.Context) (err error) {
	return
}

type UsdOrderResp struct {
	BaseResp
	Order
}

func (k *UsdOrderResp) GetData() interface{} {
	return &k.Order
}

// {
//     "avgPrice": "0.00000",              // 平均成交价
//     "clientOrderId": "abc",             // 用户自定义的订单号
//     "cumQuote": "0",                        // 成交金额
//     "executedQty": "0",                 // 成交量
//     "orderId": 1917641,                 // 系统订单号
//     "origQty": "0.40",                  // 原始委托数量
//     "origType": "TRAILING_STOP_MARKET", // 触发前订单类型
//     "price": "0",                   // 委托价格
//     "reduceOnly": false,                // 是否仅减仓
//     "side": "BUY",                      // 买卖方向
//     "status": "NEW",                    // 订单状态
//     "positionSide": "SHORT", // 持仓方向
//     "stopPrice": "9300",                    // 触发价，对`TRAILING_STOP_MARKET`无效
//     "closePosition": false,   // 是否条件全平仓
//     "symbol": "BTCUSDT",                // 交易对
//     "time": 1579276756075,              // 订单时间
//     "timeInForce": "GTC",               // 有效方法
//     "type": "TRAILING_STOP_MARKET",     // 订单类型
//     "activatePrice": "9020", // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "priceRate": "0.3", // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "updateTime": 1579276756075,        // 更新时间
//     "workingType": "CONTRACT_PRICE", // 条件价格触发类型
//     "priceProtect": false,            // 是否开启条件单触发保护
//     "priceMatch": "NONE",             //盘口价格下单模式
//     "selfTradePreventionMode": "NONE", //订单自成交保护模式
//     "goodTillDate": 0      //订单TIF为GTD时的自动取消时间
// }
