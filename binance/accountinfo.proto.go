package binance

const AccountInfoMethod = "GET"

const AccountInfoUrl = "https://api.binance.com/api/v3/account"

type AccountInfoParam struct {
	BaseParam
}

type BalanceInfo struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type AccountInfoResp struct {
	BaseResp

	AccountType string         `json:"accountType"`
	Balances    []*BalanceInfo `json:"balances"`
	UpdateTime  int64          `json:"updateTime"`
	/*
	  "makerCommission": 15,
	  "takerCommission": 15,
	  "buyerCommission": 0,
	  "sellerCommission": 0,
	  "canTrade": true,
	  "canWithdraw": true,
	  "canDeposit": true,
	  "updateTime": 123456789,
	  "accountType": "SPOT",
	  "balances": [
	    {
	      "asset": "BTC",
	      "free": "4723846.89208129",
	      "locked": "0.00000000"
	    },
	    {
	      "asset": "LTC",
	      "free": "4763368.68006011",
	      "locked": "0.00000000"
	    }
	  ],
	  "permissions": [
	    "SPOT"
	  ]
	*/
}
