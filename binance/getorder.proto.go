package binance

const GetOrderMethod = "GET"

const GetOrderUrl = "https://api.binance.com/api/v3/order"

type GetOrderParam struct {
	BaseParam
	Symbol            string `url:"symbol"`
	OrderId           int64  `url:"orderId,omitempty"`
	OrigClientOrderId string `url:"origClientOrderId,omitempty"`
	NewClientOrderId  string `url:"newClientOrderId,omitempty"`
}

type BasicOrderInfo struct {
	Symbol              string `json:"symbol"`              //交易对
	OrderId             int64  `json:"orderId"`             //28, // 系统的订单ID
	OrderListId         int64  `json:"orderListId"`         //-1, // OCO订单ID，否则为 -1
	ClientOrderId       string `json:"clientOrderId"`       //"6gCrw2kRUAF9CvJDGP16IP", // 客户自己设置的ID
	Price               string `json:"price"`               //"0.00000000", // 订单价格
	OrigQty             string `json:"origQty"`             //"10.00000000", // 用户设置的原始订单数量
	ExecutedQty         string `json:"executedQty"`         // "10.00000000", // 交易的订单数量
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"` // "10.00000000", // 累计交易的金额
	Status              string `json:"status"`              //"FILLED", // 订单状态
	TimeInForce         string `json:"timeInForce"`         //"GTC", // 订单的时效方式
	Type                string `json:"type"`                //"MARKET", // 订单类型， 比如市价单，现价单等
	Side                string `json:"side"`                //"SELL", // 订单方向，买还是卖

}

type OrderInfo struct {
	BasicOrderInfo

	StopPrice         string `json:"stopPrice"`         //"0.0", // 止损价格
	IcebergQty        string `json:"icebergQty"`        //"0.0", // 冰山数量
	Time              int64  `json:"time"`              //1499827319559, // 订单时间
	UpdateTime        int64  `json:"updateTime"`        //1499827319559, // 最后更新时间
	IsWorking         bool   `json:"isWorking"`         //true, // 订单是否出现在orderbook中
	OrigQuoteOrderQty string `json:"origQuoteOrderQty"` //"0.000000" // 原始的交易金额
}

type GetOrderResp struct {
	BaseResp
	OrderInfo
}
