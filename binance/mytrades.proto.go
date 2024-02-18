package binance

const MyTradesMethod = "GET"

const MyTradesUrl = "https://api.binance.com/api/v3/myTrades"

type MyTradesParam struct {
	BaseParam
	Symbol    string `url:"symbol,omitempty"` //交易对
	OrderId   int64  `url:"orderId,omitempty"`
	StartTime int64  `url:"startTime,omitempty"`
	EndTime   int64  `url:"endTime,omitempty"`
	FromId    int64  `url:"fromId,omitempty"` //起始Trade id。 默认获取最新交易。
	Limit     int64  `url:"limit,omitempty"`  //默认 500; 最大 1000.
}

type TradeInfo struct {
	Symbol          string `json:"symbol"`          //交易对
	Id              int64  `json:"id"`              //28457, // trade ID
	OrderId         int64  `json:"orderId"`         // 100234, // 订单ID
	OrderListId     int64  `json:"orderListId"`     // -1, // OCO订单的ID，不然就是-1
	Price           string `json:"price"`           // "4.00000100", // 成交价格
	Qty             string `json:"qty"`             // "12.00000000", // 成交量
	QuoteQty        string `json:"quoteQty"`        // "48.000012", // 成交金额
	Commission      string `json:"commission"`      // "10.10000000", // 交易费金额
	CommissionAsset string `json:"commissionAsset"` // "BNB", // 交易费资产类型
	Time            int64  `json:"time"`            // 1499865549590, // 交易时间
	IsBuyer         bool   `json:"isBuyer"`         // true, // 是否是买家
	IsMaker         bool   `json:"isMaker"`         // false, // 是否是挂单方
	IsBestMatch     bool   `json:"isBestMatch"`     // true
}

type MyTradesResp struct {
	BaseResp
	Data []*TradeInfo `json:"data"`
}

func (r *MyTradesResp) GetData() interface{} {
	return &r.Data
}
