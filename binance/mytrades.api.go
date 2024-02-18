// 本代码由脚本根据mytrades.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type MyTradesReply struct {
	reply *Reply
}

func (r *MyTradesReply) Err() error {
	return r.reply.Err()
}

func (r *MyTradesReply) Get() (resp *MyTradesResp, err error) {
	resp = &MyTradesResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) MyTrades(ctx context.Context, param *MyTradesParam) *MyTradesReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, MyTradesMethod, MyTradesUrl, param)
	}
	return &MyTradesReply{
		reply: reply,
	}
}
