// 本代码由脚本根据kline.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type KlineReply struct {
	reply *Reply
}

func (r *KlineReply) Err() error {
	return r.reply.Err()
}

func (r *KlineReply) Get() (resp *KlineResp, err error) {
	resp = &KlineResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) Kline(ctx context.Context, param *KlineParam) *KlineReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, KlineMethod, KlineUrl, param)
	}
	return &KlineReply{
		reply: reply,
	}
}
