package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/BurntSushi/toml"
)

type Key struct {
	Key    string `json:"key" yaml:"key" toml:"key"`
	Secret string `json:"secret" yaml:"secret" toml:"secret"`
}

type BinanceOption struct {
	ApiKey             Key                      `toml:"apikey"`
	Account            *AccountInfoParam        `toml:"Account"`
	GetOrder           *GetOrderParam           `toml:"GetOrder"`
	MyTrades           *MyTradesParam           `toml:"MyTrades"`
	ExchangeInfo       *ExchangeInfoParam       `toml:"ExchangeInfo"`
	Depth              *DepthParam              `toml:"Depth"`
	GetFundingAsset    *GetFundingAssetParam    `toml:"GetFundingAsset"`
	ApiRestrictions    *ApiRestrictionsParam    `toml:"ApiRestrictions"`
	TickerPrice        *TickerPriceParam        `toml:"TickerPrice"`
	Ticker24hr         *Ticker24hrParam         `toml:"Ticker24hr"`
	PostUserDataStream *PostUserDataStreamParam `toml:"PostUserDataStream"`
	Klines             *KlineParam              `toml:"Klines"`
}

func TestBinanceApi(t *testing.T) {
	conf := "client.toml"
	opt := BinanceOption{}
	_, err := toml.DecodeFile(conf, &opt)
	if err != nil {
		log.Fatalf("parse option file:%v err:%v", conf, err)
		return
	}

	cli := NewClient(opt.ApiKey.Key, string(opt.ApiKey.Secret))

	ctx := context.Background()

	tc := []string{
		"account", "exchangeinfo", "getorder", "mytrades", "depth", "getfundingasset", "apirestrictions", "tickerprice", "ticker24hr", "postuserdatastream",
	}
	for _, cmd := range tc {
		switch cmd {
		case "account":
			resp, err := cli.AccountInfo(ctx, opt.Account).Get()
			output(resp, err)
		case "exchangeinfo":
			resp, err := cli.ExchangeInfo(ctx, opt.ExchangeInfo).Get()
			output(resp, err)
		case "getorder":
			resp, err := cli.GetOrder(ctx, opt.GetOrder).Get()
			output(resp, err)
		case "mytrades":
			resp, err := cli.MyTrades(ctx, opt.MyTrades).Get()
			output(resp, err)
		case "depth":
			resp, err := cli.Depth(ctx, opt.Depth).Get()
			output(resp, err)
		case "getfundingasset":
			resp, err := cli.GetFundingAsset(ctx, opt.GetFundingAsset).Get()
			output(resp, err)
		case "apirestrictions":
			resp, err := cli.ApiRestrictions(ctx, opt.ApiRestrictions).Get()
			output(resp, err)
		case "tickerprice":
			resp, err := cli.TickerPrice(ctx, opt.TickerPrice).Get()
			output(resp, err)
		case "ticker24hr":
			resp, err := cli.Ticker24hr(ctx, opt.Ticker24hr).Get()
			output(resp, err)
		case "postuserdatastream":
			resp, err := cli.PostUserDataStream(ctx, opt.PostUserDataStream).Get()
			output(resp, err)
		default:
			fmt.Printf("unknown cmd:%v\n", cmd)
		}
	}
}

func TestKlines(t *testing.T) {
	conf := "client.toml"
	opt := BinanceOption{}
	_, err := toml.DecodeFile(conf, &opt)
	if err != nil {
		log.Fatalf("parse option file:%v err:%v", conf, err)
		return
	}

	cli := NewClient(opt.ApiKey.Key, string(opt.ApiKey.Secret))

	ctx := context.Background()
	resp, err := cli.Kline(ctx, opt.Klines).Get()
	output(resp, err)
}

func output(resp interface{}, err error) {
	if err != nil {
		log.Fatalf("request err:%v", err)
		return
	}
	dat, err := json.MarshalIndent(resp, "", "  ")
	if err == nil {
		fmt.Printf("resp:\n%v\n", string(dat))
	} else {
		fmt.Printf("resp:%#v\n", resp)
	}
}
