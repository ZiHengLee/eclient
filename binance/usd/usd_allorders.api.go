// 本代码由脚本根据usd_allorders.proto.go和usd_tmpl_api.go文件自动生成
// 请不要直接修改本文件

package usd

import (
	"context"
)

type UsdAllOrdersReply struct {
	reply *Reply
}

func (r *UsdAllOrdersReply) Err() error {
	return r.reply.Err()
}

func (r *UsdAllOrdersReply) Get() (resp *UsdAllOrdersResp, err error) {
	resp = &UsdAllOrdersResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) UsdAllOrders(ctx context.Context, param *UsdAllOrdersParam) *UsdAllOrdersReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, UsdAllOrdersMethod, UsdAllOrdersUrl, param)
	}
	return &UsdAllOrdersReply{
		reply: reply,
	}
}
