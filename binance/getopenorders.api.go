// 本代码由脚本根据getopenorders.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type GetOpenOrdersReply struct {
	reply *Reply
}

func (r *GetOpenOrdersReply) Err() error {
	return r.reply.Err()
}

func (r *GetOpenOrdersReply) Get() (resp *GetOpenOrdersResp, err error) {
	resp = &GetOpenOrdersResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) GetOpenOrders(ctx context.Context, param *GetOpenOrdersParam) *GetOpenOrdersReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, GetOpenOrdersMethod, GetOpenOrdersUrl, param)
	}
	return &GetOpenOrdersReply{
		reply: reply,
	}
}
