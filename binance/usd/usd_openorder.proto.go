// 查询当前挂单 (USER_DATA)
package usd

import (
	"context"

	"github.com/slowly-richer/eclient/binance"
)

const UsdOpenOrderMethod = "GET"

const UsdOpenOrderUrl = "https://fapi.binance.com/fapi/v1/openOrder"

/*
1.orderId 与 origClientOrderId 中的一个为必填参数
2.查询的订单如果已经成交或取消，将返回报错 "Order does not exist."
*/
type UsdOpenOrderParam struct {
	binance.BaseParam
	Symbol            string `url:"symbol"`
	OrigClientOrderId string `url:"origClientOrderId,omitempty"` //用户自定义的订单号
	OrderId           int64  `url:"orderId,omitempty"`           //系统订单号
}

func (k UsdOpenOrderParam) RequireAuth() bool {
	return true
}

func (k *UsdOpenOrderParam) Prepare(ctx context.Context) (err error) {
	return
}

type UsdOpenOrderResp struct {
	BaseResp
	Order
}

func (k *UsdOpenOrderResp) GetData() interface{} {
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
