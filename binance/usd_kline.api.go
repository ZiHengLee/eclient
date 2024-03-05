// 本代码由脚本根据usd_kline.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type UsdKlineReply struct {
	reply *Reply
}

func (r *UsdKlineReply) Err() error {
	return r.reply.Err()
}

func (r *UsdKlineReply) Get() (resp *UsdKlineResp, err error) {
	resp = &UsdKlineResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) UsdKline(ctx context.Context, param *UsdKlineParam) *UsdKlineReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, UsdKlineMethod, UsdKlineUrl, param)
	}
	return &UsdKlineReply{
		reply: reply,
	}
}
