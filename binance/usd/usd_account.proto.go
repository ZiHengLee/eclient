// 获取历史k线数据
package usd

import (
	"context"

	"github.com/slowly-richer/eclient/binance"
)

const UsdAccountMethod = "GET"

const UsdAccountUrl = "https://fapi.binance.com/fapi/v2/account"

type UsdAccountParam struct {
	binance.BaseParam
}

func (k UsdAccountParam) RequireAuth() bool {
	return true
}

func (k *UsdAccountParam) Prepare(ctx context.Context) (err error) {
	return
}

type UsdAccountResp struct {
	BaseResp
	FeeTier                     int        `json:"feeTier"`
	CanTrade                    bool       `json:"canTrade"`
	CanDeposit                  bool       `json:"canDeposit"`
	CanWithdraw                 bool       `json:"canWithdraw"`
	UpdateTime                  int64      `json:"updateTime"`
	MultiAssetsMargin           bool       `json:"multiAssetsMargin"`
	TradeGroupId                int        `json:"tradeGroupId"`
	TotalInitialMargin          string     `json:"totalInitialMargin"`
	TotalMaintMargin            string     `json:"totalMaintMargin"`
	TotalWalletBalance          string     `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string     `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string     `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string     `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string     `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string     `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string     `json:"totalCrossUnPnl"`
	AvailableBalance            string     `json:"availableBalance"`
	MaxWithdrawAmount           string     `json:"maxWithdrawAmount"`
	Assets                      []Asset    `json:"assets"`
	Positions                   []Position `json:"positions"`
}
type Asset struct {
	Asset                  string `json:"asset"`
	WalletBalance          string `json:"walletBalance"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	MarginBalance          string `json:"marginBalance"`
	MaintMargin            string `json:"maintMargin"`
	InitialMargin          string `json:"initialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int64  `json:"updateTime"`
}

type Position struct {
	Symbol                 string `json:"symbol"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Leverage               string `json:"leverage"`
	Isolated               bool   `json:"isolated"`
	EntryPrice             string `json:"entryPrice"`
	MaxNotional            string `json:"maxNotional"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	UpdateTime             int64  `json:"updateTime"`
}

// {
//     "feeTier": 0,  // 手续费等级
//     "canTrade": true,  // 是否可以交易
//     "canDeposit": true,  // 是否可以入金
//     "canWithdraw": true, // 是否可以出金
//     "updateTime": 0,     // 保留字段，请忽略
//     "multiAssetsMargin": false,
//     "tradeGroupId": -1,
//     "totalInitialMargin": "0.00000000",  // 当前所需起始保证金总额(存在逐仓请忽略), 仅计算usdt资产
//     "totalMaintMargin": "0.00000000",  // 维持保证金总额, 仅计算usdt资产
//     "totalWalletBalance": "23.72469206",   // 账户总余额, 仅计算usdt资产
//     "totalUnrealizedProfit": "0.00000000",  // 持仓未实现盈亏总额, 仅计算usdt资产
//     "totalMarginBalance": "23.72469206",  // 保证金总余额, 仅计算usdt资产
//     "totalPositionInitialMargin": "0.00000000",  // 持仓所需起始保证金(基于最新标记价格), 仅计算usdt资产
//     "totalOpenOrderInitialMargin": "0.00000000",  // 当前挂单所需起始保证金(基于最新标记价格), 仅计算usdt资产
//     "totalCrossWalletBalance": "23.72469206",  // 全仓账户余额, 仅计算usdt资产
//     "totalCrossUnPnl": "0.00000000",    // 全仓持仓未实现盈亏总额, 仅计算usdt资产
//     "availableBalance": "23.72469206",       // 可用余额, 仅计算usdt资产
//     "maxWithdrawAmount": "23.72469206"     // 最大可转出余额, 仅计算usdt资产
//     "assets": [
//         {
//             "asset": "USDT",        //资产
//             "walletBalance": "23.72469206",  //余额
//             "unrealizedProfit": "0.00000000",  // 未实现盈亏
//             "marginBalance": "23.72469206",  // 保证金余额
//             "maintMargin": "0.00000000",    // 维持保证金
//             "initialMargin": "0.00000000",  // 当前所需起始保证金
//             "positionInitialMargin": "0.00000000",  // 持仓所需起始保证金(基于最新标记价格)
//             "openOrderInitialMargin": "0.00000000", // 当前挂单所需起始保证金(基于最新标记价格)
//             "crossWalletBalance": "23.72469206",  //全仓账户余额
//             "crossUnPnl": "0.00000000" // 全仓持仓未实现盈亏
//             "availableBalance": "126.72469206",       // 可用余额
//             "maxWithdrawAmount": "23.72469206",     // 最大可转出余额
//             "marginAvailable": true,   // 是否可用作联合保证金
//             "updateTime": 1625474304765  //更新时间
//         }
//     ],
//     "positions": [  // 头寸，将返回所有市场symbol。
//         //根据用户持仓模式展示持仓方向，即单向模式下只返回BOTH持仓情况，双向模式下只返回 LONG 和 SHORT 持仓情况
//         {
//             "symbol": "BTCUSDT",  // 交易对
//             "initialMargin": "0",   // 当前所需起始保证金(基于最新标记价格)
//             "maintMargin": "0", //维持保证金
//             "unrealizedProfit": "0.00000000",  // 持仓未实现盈亏
//             "positionInitialMargin": "0",  // 持仓所需起始保证金(基于最新标记价格)
//             "openOrderInitialMargin": "0",  // 当前挂单所需起始保证金(基于最新标记价格)
//             "leverage": "100",  // 杠杆倍率
//             "isolated": true,  // 是否是逐仓模式
//             "entryPrice": "0.00000",  // 持仓成本价
//             "maxNotional": "250000",  // 当前杠杆下用户可用的最大名义价值
//             "bidNotional": "0",  // 买单净值，忽略
//             "askNotional": "0",  // 卖单净值，忽略
//             "positionSide": "BOTH",  // 持仓方向
//             "positionAmt": "0",      // 持仓数量
//             "updateTime": 0         // 更新时间
//         }
//     ]
// }
