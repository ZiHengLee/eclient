package binance

const AllOrdersMethod = "GET"

const AllOrdersUrl = "https://api.binance.com/api/v3/allOrders"

type AllOrdersParam struct {
	BaseParam
	Symbol    string `url:"symbol"`
	OrderId   int64  `url:"orderId,omitempty"`
	StartTime int64  `url:"startTime,omitempty"`
	EndTime   int64  `url:"endTime,omitempty"`
	Limit     int64  `url:"limit,omitempty"`
}

type AllOrdersResp struct {
	BaseResp
	Data []*OrderInfo `json:"data"`
}

func (r *AllOrdersResp) GetData() interface{} {
	return &r.Data
}
