// 本代码由脚本根据usd_order.proto.go和usd_tmpl_api.go文件自动生成
// 请不要直接修改本文件

package usd

import (
	"context"
)

type UsdOrderReply struct {
	reply *Reply
}

func (r *UsdOrderReply) Err() error {
	return r.reply.Err()
}

func (r *UsdOrderReply) Get() (resp *UsdOrderResp, err error) {
	resp = &UsdOrderResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) UsdOrder(ctx context.Context, param *UsdOrderParam) *UsdOrderReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, UsdOrderMethod, UsdOrderUrl, param)
	}
	return &UsdOrderReply{
		reply: reply,
	}
}
