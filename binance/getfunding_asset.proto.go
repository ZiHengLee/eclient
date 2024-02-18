package binance

const GetFundingAssetMethod = "POST"

const GetFundingAssetUrl = "https://api.binance.com/sapi/v1/asset/get-funding-asset"

type GetFundingAssetParam struct {
	BaseParam
	Asset            string `url:"asset,omitempty"` //交易对
	NeedBtcValuation bool   `url:"needBtcValuation,omitempty"`
}

type FundingAssetItem struct {
	Asset        string `json:"asset"`        //"USDT",
	Free         string `json:"free"`         //"1",    // 可用余额
	Locked       string `json:"locked"`       //"0",  // 锁定资金
	Freeze       string `json:"freeze"`       //"0",  //冻结资金
	WithDrawing  string `json:"withdrawing"`  //"0",  // 提币
	BtcValuation string `json:"btcValuation"` //"0.00000091"  // btc估值
}

type GetFundingAssetResp struct {
	BaseResp
	Data []*FundingAssetItem `json:"data"`
}

func (r *GetFundingAssetResp) GetData() interface{} {
	return &r.Data
}
