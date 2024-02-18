package binance

import (
	"github.com/ZiHengLee/eclient/utils/number"
)

type BaseStreamMsg struct {
	Event     string `json:"e"` // 事件类型
	Symbol    string `json:"s"` // 交易对
	EventTime int64  `json:"E"` // 事件时间
}

type DepthUpdateMsg struct {
	BaseStreamMsg
	FirstUpdateId int64              `json:"U"` // 从上次推送至今新增的第一个 update Id
	LastUpdateId  int64              `json:"u"` // 从上次推送至今新增的最后一个 update Id
	Buyer         [][2]number.Number `json:"b"` // 变动的买单深度 [变动的价格档位, 数量]
	Asker         [][2]number.Number `json:"a"` // 变动的卖单深度 [变动的价格档位, 数量]
}

type ExecutionReportMsg struct {
	BaseStreamMsg
	ClientOrderId       string `json:"c"` //"mUvoqJxFIILMdfAW5iGSOW", // clientOrderId,本次操作的订单ID
	OrigClientOrderId   string `json:"C"` // "",                       // 原始订单自定义ID(原始订单，指撤单操作的对象。撤单本身被视为另一个订单)
	OrderId             int64  `json:"i"` // 4293153,                  // orderId
	Ignore              int64  `json:"I"` // 8641984,                  // 请忽略
	Side                string `json:"S"` //"BUY",                    // 订单方向
	Type                string `json:"o"` //"LIMIT",                  // 订单类型
	OrderTime           int64  `json:"O"` // 1499405658657,            // 订单创建时间
	TimeInForce         string `json:"f"` //"GTC",                    // 有效方式
	IceQty              string `json:"F"` // "0.00000000",             // 冰山订单数量
	Quantity            string `json:"q"` //"1.00000000",             // 订单原始数量
	QuoteQty            string `json:"Q"` // "0.00000000"              // Quote Order Qty
	Price               string `json:"p"` // "0.10264410",             // 订单原始价格
	ProfitLossPrice     string `json:"P"` // "0.00000000",             // 止盈止损单触发价格
	Action              string `json:"x"` // "NEW",                    // 本次事件的具体执行类型
	Status              string `json:"X"` // "NEW",                    // 订单的当前状态
	ExecutedQty         string `json:"z"` // "0.00000000",             // 订单累计已成交量
	CummulativeQuoteQty string `json:"Z"` // "0.00000000",             // 订单累计已成交金额
	/*
	  "g": -1,                       // OCO订单 OrderListId
	  "r": "NONE",                   // 订单被拒绝的原因
	  "l": "0.00000000",             // 订单末次成交量
	  "L": "0.00000000",             // 订单末次成交价格
	  "n": "0",                      // 手续费数量
	  "N": null,                     // 手续费资产类别
	  "T": 1499405658657,            // 成交时间
	  "t": -1,                       // 成交ID
	  "w": true,                     // 订单是否在订单簿上？
	  "m": false,                    // 该成交是作为挂单成交吗？
	  "M": false,                    // 请忽略
	  "Y": "0.00000000",              // 订单末次成交金额
	*/
}
