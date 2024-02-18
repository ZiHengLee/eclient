package binance

const DeleteOrderMethod = "DELETE"

const DeleteOrderUrl = "https://api.binance.com/api/v3/order"

type DeleteOrderParam struct {
	BaseParam
	Symbol            string `url:"symbol"`
	OrderId           int64  `url:"orderId,omitempty"`
	OrigClientOrderId string `url:"origClientOrderId,omitempty"`
	NewClientOrderId  string `url:"newClientOrderId,omitempty"`
}

type DeleteOrderResp struct {
	BaseResp
	Symbol              string `json:"symbol"`              //"LTCBTC"
	OrigClientOrderId   string `json:"origClientOrderId"`   //"myOrder1",
	OrderId             int64  `json:"orderId"`             //4,
	OrderListId         int64  `json:"orderListId"`         //-1, // OCO订单ID，否则为 -1
	ClientOrderId       string `json:"clientOrderId"`       //"cancelMyOrder1",
	Price               string `json:"price"`               //"2.00000000",
	OrigQty             string `json:"origQty"`             //"1.00000000",
	ExecutedQty         string `json:"executedQty"`         //"0.00000000",
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"` //"0.00000000",
	Status              string `json:"status"`              //"CANCELED",
	TimeInForce         string `json:"timeInForce"`         //"GTC",
	Type                string `json:"type"`                //"LIMIT",
	Side                string `json:"side"`                //"BUY"
}
