# 币安u本位合约所有接口
```
1.测试服务器连通性 PING GET /fapi/v1/ping
2.获取服务器时间 GET /fapi/v1/time
3.获取交易规则和交易对 GET /fapi/v1/exchangeInfo
4.深度信息 GET /fapi/v1/depth
5.近期成交 GET /fapi/v1/trades
6.查询历史成交(MARKET_DATA) GET /fapi/v1/historicalTrades
7.近期成交(归集) GET /fapi/v1/aggTrades
8.K线数据 GET /fapi/v1/klines
9.连续合约K线数据 GET /fapi/v1/continuousKlines 
10.价格指数K线数据 GET /fapi/v1/indexPriceKlines
11.标记价格K线数据 GET /fapi/v1/markPriceKlines
12.溢价指数K线数据 GET /fapi/v1/premiumIndexKlines
13.最新标记价格和资金费率 GET /fapi/v1/premiumIndex
14.查询资金费率历史 GET /fapi/v1/fundingRate
15.查询资金费率信息 GET /fapi/v1/fundingInfo
16.24hr价格变动情况 GET /fapi/v1/ticker/24hr
17.最新价格 GET /fapi/v1/ticker/price
18.最新价格V2 GET /fapi/v2/ticker/price
19.当前最优挂单 GET /fapi/v1/ticker/bookTicker
20.获取未平仓合约数 GET /fapi/v1/openInterest
21.季度合约历史结算价 GET /futures/data/delivery-price
22.合约持仓量 GET /futures/data/openInterestHist
23.大户账户数多空比 GET /futures/data/topLongShortAccountRatio
24.大户持仓量多空比 GET /futures/data/topLongShortPositionRatio
25.多空持仓人数比 GET /futures/data/globalLongShortAccountRatio
26.合约主动买卖量 GET /futures/data/takerlongshortRatio
27.杠杆代币历史净值K线 GET /fapi/v1/lvtKlines
28.综合指数交易对信息 GET /fapi/v1/indexInfo
29.多资产模式资产汇率指数 GET /fapi/v1/assetIndex
30.查询指数价格成分 GET /fapi/v1/constituents
```

# 文件生成
```
make api
``````