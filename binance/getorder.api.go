// 本代码由脚本根据getorder.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type GetOrderReply struct {
	reply *Reply
}

func (r *GetOrderReply) Err() error {
	return r.reply.Err()
}

func (r *GetOrderReply) Get() (resp *GetOrderResp, err error) {
	resp = &GetOrderResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) GetOrder(ctx context.Context, param *GetOrderParam) *GetOrderReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, GetOrderMethod, GetOrderUrl, param)
	}
	return &GetOrderReply{
		reply: reply,
	}
}
