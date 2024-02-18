// 本代码由脚本根据allorders.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type AllOrdersReply struct {
	reply *Reply
}

func (r *AllOrdersReply) Err() error {
	return r.reply.Err()
}

func (r *AllOrdersReply) Get() (resp *AllOrdersResp, err error) {
	resp = &AllOrdersResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) AllOrders(ctx context.Context, param *AllOrdersParam) *AllOrdersReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, AllOrdersMethod, AllOrdersUrl, param)
	}
	return &AllOrdersReply{
		reply: reply,
	}
}
