package binance

const ApiRestrictionsMethod = "GET"

const ApiRestrictionsUrl = "https://api.binance.com/sapi/v1/account/apiRestrictions"

type ApiRestrictionsParam struct {
	BaseParam
}

type ApiRestrictionsResp struct {
	BaseResp
	IpRestrict                     bool  `json:"ipRestrict"`                     //false,  // 是否限制ip访问
	CreateTime                     int64 `json:"createTime"`                     //1623840271000,   // 创建时间
	EnableWithdrawals              bool  `json:"enableWithdrawals"`              //false,   // 此选项允许通过此api提现。开启提现选项必须添加IP访问限制过滤器
	EnableInternalTransfer         bool  `json:"enableInternalTransfer"`         //true,  // 此选项授权此密钥在您的母账户和子账户之间划转资金
	PermitsUniversalTransfer       bool  `json:"permitsUniversalTransfer"`       //true,  // 授权该密钥可用于专用的万向划转接口，用以操作其支持的多种类型资金划转。各业务自身的划转接口使用权限，不受本授权影响
	EnableVanillaOptions           bool  `json:"enableVanillaOptions"`           //false,  // 欧式期权交易权限
	EnableReading                  bool  `json:"enableReading"`                  //true,
	EnableFutures                  bool  `json:"enableFutures"`                  //false,  // 合约交易权限，需注意开通合约账户之前创建的API Key不支持合约API功能
	EnableMargin                   bool  `json:"enableMargin"`                   //false,   // 此选项在全仓账户完成划转后可编辑
	EnableSpotAndMarginTrading     bool  `json:"enableSpotAndMarginTrading"`     //false, // 现货和杠杆交易权限
	TradingAuthorityExpirationTime int64 `json:"tradingAuthorityExpirationTime"` //1628985600000  // 现货和杠杆交易权限到期时间，如果没有则不返回该字段
}
