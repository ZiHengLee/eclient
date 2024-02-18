// 本代码由脚本根据ticker24hr.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type Ticker24hrReply struct {
	reply *Reply
}

func (r *Ticker24hrReply) Err() error {
	return r.reply.Err()
}

func (r *Ticker24hrReply) Get() (resp *Ticker24hrResp, err error) {
	resp = &Ticker24hrResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) Ticker24hr(ctx context.Context, param *Ticker24hrParam) *Ticker24hrReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, Ticker24hrMethod, Ticker24hrUrl, param)
	}
	return &Ticker24hrReply{
		reply: reply,
	}
}
