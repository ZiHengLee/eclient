// 本代码由脚本根据tickerprice.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type TickerPriceReply struct {
	reply *Reply
}

func (r *TickerPriceReply) Err() error {
	return r.reply.Err()
}

func (r *TickerPriceReply) Get() (resp *TickerPriceResp, err error) {
	resp = &TickerPriceResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) TickerPrice(ctx context.Context, param *TickerPriceParam) *TickerPriceReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, TickerPriceMethod, TickerPriceUrl, param)
	}
	return &TickerPriceReply{
		reply: reply,
	}
}
