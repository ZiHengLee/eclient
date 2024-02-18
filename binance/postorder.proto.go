package binance

const PostOrderMethod = "POST"

const PostOrderUrl = "https://api.binance.com/api/v3/order"

type PostOrderParam struct {
	BaseParam
	Symbol           string `url:"symbol"`
	Side             string `url:"side"`
	Type             string `url:"type"`
	TimeInForce      string `url:"timeInForce,omitempty"`
	Quantity         string `url:"quantity,omitempty"`
	QuoteOrderQty    string `url:"quoteOrderQty,omitempty"`
	Price            string `url:"price,omitempty"`
	NewClientOrderId string `url:"newClientOrderId,omitempty"`
	StopPrice        string `url:"stopPrice,omitempty"`
	IcebergQty       string `url:"icebergQty,omitempty"`
	NewOrderRespType string `url:"newOrderRespType,omitempty"`
}

type OrderFill struct {
	TradeId         int64  `json:"tradeId"`
	Price           string `json:"price"`           //"4000.00000000", // 交易的价格
	Qty             string `json:"qty"`             //"1.00000000", // 交易的数量
	Commission      string `json:"commission"`      //"4.00000000", // 手续费金额
	CommissionAsset string `json:"commissionAsset"` //"USDT" // 手续费的币种
}

type PostOrderResp struct {
	BaseResp
	BasicOrderInfo

	TransactTime int64        `json:"transactTime"` //1507725176595, // 交易的时间戳
	Fills        []*OrderFill `json:"fills"`        // 订单中交易的信息
}
