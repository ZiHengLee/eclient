// 本代码由脚本根据stock_info.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package eastmoney

import (
	"context"
)

type StockInfoReply struct {
	reply *Reply
}

func (r *StockInfoReply) Err() error {
	return r.reply.Err()
}

func (r *StockInfoReply) Get() (stock *StockInfo, err error) {
	resp := &StockInfoResp{}
	stock = &StockInfo{}
	err = r.reply.Get(&resp.BaseResp, resp)
	if err == nil {
		for k, v := range resp.Data {
			switch k {
			case "f57":
				stock.Code = v.(string)
			case "f58":
				stock.StockShortName, _ = v.(string)
			case "f84":
				stock.TotalQuity, _ = v.(float64)
			case "f85":
				stock.FloatQuity, _ = v.(float64)
			case "f127":
				stock.Industry, _ = v.(string)
			case "f116":
				stock.TotalMarketValue, _ = v.(float64)
			case "f117":
				stock.FloatMarketValue, _ = v.(float64)
			case "f189":
				stock.ListingDate, _ = v.(float64)
			}
		}
	}
	return
}

func (c Client) StockInfo(ctx context.Context, param *StockInfoParam) *StockInfoReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, StockInfoMethod, StockInfoUrl, param)
	}
	return &StockInfoReply{
		reply: reply,
	}
}
