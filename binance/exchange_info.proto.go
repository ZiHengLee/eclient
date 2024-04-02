package binance

import (
	"context"
	"encoding/json"
	"fmt"

	//"git.coinracing.com/coint/strade/utils/logger"
	"github.com/slowly-richer/eclient/utils/number"
)

const ExchangeInfoMethod = "GET"

const ExchangeInfoUrl = "https://api.binance.com/api/v3/exchangeInfo"

type ExchangeInfoParam struct {
	BaseParam
	Symbol  string `url:"symbol,omitempty"`
	Symbols string `url:"symbols,omitempty"`
}

func (p *ExchangeInfoParam) RequireAuth() bool {
	return false
}

func (p *ExchangeInfoParam) Prepare(ctx context.Context) (err error) {
	return
}

type ExchangeFilter = map[string]interface{}

/*
价格过滤器 用于检测订单中 price 参数的合法性。包含以下三个部分:

minPrice 定义了 price/stopPrice 允许的最小值。
maxPrice 定义了 price/stopPrice 允许的最大值。
tickSize 定义了 price/stopPrice 的步进间隔，即price必须等于minPrice+(tickSize的整数倍)
以上每一项均可为0，为0时代表这一项不再做限制。

逻辑伪代码如下:

price >= minPrice
price <= maxPrice
(price-minPrice) % tickSize == 0
*/
type ExchangeFilterPriceFilter struct {
	FilterType string         `json:"filterType"` //"PRICE_FILTER",
	MinPrice   *number.Number `json:"minPrice"`   //"0.00100000",
	MaxPrice   *number.Number `json:"maxPrice"`   //"100000.00000000",
	TickSize   *number.Number `json:"tickSize"`   //"0.00100000"
}

/*
Lots是拍卖术语，LOT_SIZE 过滤器对订单中的 quantity 也就是数量参数进行合法性检查。包含三个部分:

minQty 表示 quantity/icebergQty 允许的最小值。
maxQty 表示 quantity/icebergQty 允许的最大值。
stepSize 表示 quantity/icebergQty 允许的步进值。
逻辑伪代码如下:

quantity >= minQty
quantity <= maxQty
(quantity-minQty) % stepSize == 0
*/
type ExchangeFilterLotSize struct {
	FilterType string         `json:"filterType"` //"LOT_SIZE",
	MinQty     *number.Number `json:"minQty"`     //"0.00100000",
	MaxQty     *number.Number `json:"maxQty"`     //"100000.00000000",
	StepSize   *number.Number `json:"stepSize"`   //"0.00100000"
}

/*
MARKET_LOT_SIZE过滤器为交易对上的MARKET订单定义了数量(即拍卖中的"手数")规则。 共有3部分：

minQty定义了允许的最小quantity。
maxQty定义了允许的最大数量。
stepSize定义了可以增加/减少数量的间隔。
为了通过market lot size，quantity必须满足以下条件：

quantity >= minQty
quantity <= maxQty
(quantity-minQty) % stepSize == 0
*/
type ExchangeFilterMarketLotSize struct {
	FilterType string         `json:"filterType"` //"MARKET_LOT_SIZE",
	MinQty     *number.Number `json:"minQty"`     //"0.00100000",
	MaxQty     *number.Number `json:"maxQty"`     //"100000.00000000",
	StepSize   *number.Number `json:"stepSize"`   //"0.00100000"
}

/*
MIN_NOTIONAL过滤器定义了交易对订单所允许的最小名义价值(成交额)。 订单的名义价值是价格*数量。 如果是高级订单(比如止盈止损订单STOP_LOSS_LIMIT)，名义价值会按照stopPrice * quantity来计算。 如果是冰山订单，名义价值会按照price * icebergQty来计算。 applyToMarket确定 MIN_NOTIONAL过滤器是否也将应用于MARKET订单。
由于MARKET订单没有价格，因此会在最后avgPriceMins分钟内使用平均价格。
avgPriceMins是计算平均价格的分钟数。 0表示使用最后的价格。
*/
type ExchangeFilterMinNotional struct {
	FilterType    string         `json:"filterType"`    //"MIN_NOTIONAL",
	MinNotional   *number.Number `json:"minNotional"`   //"0.00100000",
	ApplyToMarket bool           `json:"applyToMarket"` //true,
	AvgPriceMins  int            `json:"avgPriceMins"`  //5
}

type SymbolExchangeInfo struct {
	Symbol              string   `json:"symbol"`              //"ETHBTC",
	Status              string   `json:"status"`              //"TRADING",
	BaseAsset           string   `json:"baseAsset"`           //"ETH",
	BaseAssetPrecision  int      `json:"baseAssetPrecision"`  //8,
	QuoteAsset          string   `json:"quoteAsset"`          // "BTC",
	QuotePrecision      int      `json:"quotePrecision"`      // 8,
	QuoteAssetPrecision int      `json:"quoteAssetPrecision"` // 8,
	OrderTypes          []string `json:"orderTypes"`          /*[
	    "LIMIT",
	    "LIMIT_MAKER",
	    "MARKET",
	    "STOP_LOSS",
	    "STOP_LOSS_LIMIT",
	    "TAKE_PROFIT",
	    "TAKE_PROFIT_LIMIT"
	]*/
	IcebergAllowed         bool             `json:"icebergAllowed"`         // true,
	OcoAllowed             bool             `json:"ocoAllowed"`             // true,
	IsSpotTradingAllowed   bool             `json:"isSpotTradingAllowed"`   // true,
	IsMarginTradingAllowed bool             `json:"isMarginTradingAllowed"` // true,
	Filters                []ExchangeFilter `json:"filters"`                /* [
	    //这些在"过滤器"部分中定义
	    //所有限制都是可选的
	]*/
	Permissions []string `json:"permissions"` /* [
	   "SPOT",
	   "MARGIN"
	 ]*/

	PriceFilter   ExchangeFilterPriceFilter   `json:"-"`
	LotSize       ExchangeFilterLotSize       `json:"-"`
	MinNotional   ExchangeFilterMinNotional   `json:"-"`
	MarketLotSize ExchangeFilterMarketLotSize `json:"-"`
}

func (s *SymbolExchangeInfo) Build() (err error) {
	m := map[string]interface{}{
		"PRICE_FILTER":    &s.PriceFilter,
		"LOT_SIZE":        &s.LotSize,
		"MIN_NOTIONAL":    &s.MinNotional,
		"MARKET_LOT_SIZE": &s.MarketLotSize,
	}
	for _, f := range s.Filters {
		if k, ok := f["filterType"]; ok {
			sk := k.(string)
			if obj, ok := m[sk]; ok {
				dat, err := json.Marshal(f)
				if err != nil {
					return err
				}
				err = json.Unmarshal(dat, obj)
				if err != nil {
					return err
				}
			}
		}
	}
	return
}

func (s *SymbolExchangeInfo) CheckPrice(p *number.Number, adjust int) (err error) {
	f := s.PriceFilter
	if f.TickSize != nil && f.TickSize.Sign() != 0 {
		if adjust != number.AdjustNone {
			p.Sub(f.MinPrice).Div(f.TickSize).Adjust(0, adjust)
			p.Mul(f.TickSize).Add(f.MinPrice)
		} else {
			sp := p.Copy()
			sp.Sub(f.MinPrice)
			sp.Mod(f.TickSize)
			if sp.Sign() != 0 {
				err = fmt.Errorf("price unmatch TickSize:%v", f.TickSize.ToString())
				return
			}
		}
	}
	if f.MinPrice != nil && f.MinPrice.Sign() != 0 && p.Cmp(f.MinPrice) < 0 {
		err = fmt.Errorf("price less MinPrice:%v", f.MinPrice.ToString())
		return
	}
	if f.MaxPrice != nil && f.MaxPrice.Sign() != 0 && p.Cmp(f.MaxPrice) > 0 {
		err = fmt.Errorf("price great MaxPrice:%v", f.MaxPrice.ToString())
		return
	}
	return
}

func (s *SymbolExchangeInfo) CheckQty(q *number.Number, adjust int) (err error) {
	f := s.LotSize
	if adjust != number.AdjustNone && f.StepSize != nil && f.StepSize.Sign() > 0 {
		q.Sub(f.MinQty).Div(f.StepSize).Adjust(0, adjust)
		q.Mul(f.StepSize).Add(f.MinQty)
	}
	if f.MinQty != nil && f.MinQty.Sign() > 0 && q.Cmp(f.MinQty) < 0 {
		err = fmt.Errorf("qty less MinQty:%v", f.MinQty.ToString())
	} else if f.MaxQty != nil && f.MaxQty.Sign() > 0 && q.Cmp(f.MaxQty) > 0 {
		err = fmt.Errorf("qty great MaxQty:%v", f.MaxQty.ToString())
	}
	return
}

func (s *SymbolExchangeInfo) CheckMarketQty(q *number.Number, adjust int) (err error) {
	f := s.MarketLotSize
	if adjust != number.AdjustNone && f.StepSize != nil && f.StepSize.Sign() > 0 {
		q.Sub(f.MinQty).Div(f.StepSize).Floor(0)
		q.Mul(f.StepSize).Add(f.MinQty)
	}
	if f.MinQty != nil && f.MinQty.Sign() > 0 && q.Cmp(f.MinQty) < 0 {
		err = fmt.Errorf("qty less MinQty:%v", f.MinQty.ToString())
	} else if f.MaxQty != nil && f.MaxQty.Sign() > 0 && q.Cmp(f.MaxQty) > 0 {
		err = fmt.Errorf("qty great MaxQty:%v", f.MaxQty.ToString())
	}
	return
}

func (s *SymbolExchangeInfo) CheckQuoteQty(q *number.Number, adjust int) (err error) {
	var step *number.Number
	if s.LotSize.StepSize != nil && s.PriceFilter.TickSize != nil {
		step = s.LotSize.StepSize.Copy().Mul(s.PriceFilter.TickSize)
	}
	if step != nil && step.Sign() > 0 && adjust != number.AdjustNone {
		n := q.Copy().Div(step).Adjust(0, adjust)
		q = n.Mul(step)
	}

	f := s.MinNotional
	if f.MinNotional != nil && f.MinNotional.Cmp(q) > 0 {
		err = fmt.Errorf("quote qty less MinNotional:%v", f.MinNotional.ToString())
	}
	return
}

func (s *SymbolExchangeInfo) adjust(q, min, max, step *number.Number, roundMethod int) (v *number.Number) {
	if step == nil || step.Sign() == 0 {
		v = q.Copy()
	} else {
		if min == nil {
			min = &number.Number{}
		}
		n := q.Copy().Sub(min).Div(step)
		if roundMethod == number.AdjustFloor {
			n.Floor(0)
		} else if roundMethod == number.AdjustCeil {
			n.Ceil(0)
		} else if roundMethod == number.AdjustRound {
			n.Round(0)
		}
		v = min.Copy().Add(n.Mul(step))
	}

	if min != nil && min.Sign() > 0 && v.Cmp(min) < 0 {
		v = min.Copy()
	} else if max != nil && max.Sign() > 0 && v.Cmp(max) > 0 {
		v = max.Copy()
	}
	return
}

func (s *SymbolExchangeInfo) AdjustPrice(q *number.Number, roundMethod int) (v *number.Number) {
	f := s.PriceFilter
	return s.adjust(q, f.MinPrice, f.MaxPrice, f.TickSize, roundMethod)
}

func (s *SymbolExchangeInfo) AdjustQty(q *number.Number, roundMethod int) (v *number.Number) {
	f := s.LotSize
	return s.adjust(q, f.MinQty, f.MaxQty, f.StepSize, roundMethod)
}

func (s *SymbolExchangeInfo) AdjustMarketQty(q *number.Number, roundMethod int) (v *number.Number) {
	f := s.MarketLotSize
	t := s.adjust(q, f.MinQty, f.MaxQty, f.StepSize, roundMethod)
	return s.AdjustQty(t, roundMethod)
}

func (s *SymbolExchangeInfo) AdjustQuoteQty(q *number.Number, roundMethod int) (v *number.Number) {
	f := s.MinNotional
	if f.MinNotional != nil && f.MinNotional.Cmp(q) > 0 {
		v = f.MinNotional.Copy()
	} else {
		pf := s.PriceFilter
		qf := s.LotSize
		v = q.Copy()
		if (pf.TickSize != nil && pf.TickSize.Sign() > 0) && (qf.StepSize != nil && qf.StepSize.Sign() > 0) {
			step := pf.TickSize.Copy().Mul(qf.StepSize)
			if step.Sign() > 0 {
				n := q.Copy().Div(step).Adjust(0, roundMethod)
				v = n.Mul(step)
			}
		}
		precise := s.QuotePrecision
		v.Adjust(precise, roundMethod)
	}
	return
}

type ExchangeInfoResp struct {
	BaseResp
	Timezone   string               `json:"timezone"`
	ServerTime int64                `json:"serverTime"`
	Symbols    []SymbolExchangeInfo `json:"symbols"`
}

func (r *ExchangeInfoResp) Build() (err error) {
	for i := range r.Symbols {
		err = r.Symbols[i].Build()
		if err != nil {
			return
		}
	}
	return
}
