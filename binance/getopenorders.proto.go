package binance

const GetOpenOrdersMethod = "GET"

const GetOpenOrdersUrl = "https://api.binance.com/api/v3/openOrders"

type GetOpenOrdersParam struct {
	BaseParam
	Symbol string `url:"symbol,omitempty"` //交易对
}

type GetOpenOrdersResp struct {
	BaseResp
	Data []*OrderInfo `json:"data"`
}

func (r *GetOpenOrdersResp) GetData() interface{} {
	return &r.Data
}
