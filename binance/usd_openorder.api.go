// 本代码由脚本根据usd_openorder.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type UsdOpenOrderReply struct {
	reply *Reply
}

func (r *UsdOpenOrderReply) Err() error {
	return r.reply.Err()
}

func (r *UsdOpenOrderReply) Get() (resp *UsdOpenOrderResp, err error) {
	resp = &UsdOpenOrderResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) UsdOpenOrder(ctx context.Context, param *UsdOpenOrderParam) *UsdOpenOrderReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, UsdOpenOrderMethod, UsdOpenOrderUrl, param)
	}
	return &UsdOpenOrderReply{
		reply: reply,
	}
}
