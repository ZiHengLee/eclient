// 本代码由脚本根据exchange_info.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type ExchangeInfoReply struct {
	reply *Reply
}

func (r *ExchangeInfoReply) Err() error {
	return r.reply.Err()
}

func (r *ExchangeInfoReply) Get() (resp *ExchangeInfoResp, err error) {
	resp = &ExchangeInfoResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) ExchangeInfo(ctx context.Context, param *ExchangeInfoParam) *ExchangeInfoReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, ExchangeInfoMethod, ExchangeInfoUrl, param)
	}
	return &ExchangeInfoReply{
		reply: reply,
	}
}
