// 本代码由脚本根据postorder.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type PostOrderReply struct {
	reply *Reply
}

func (r *PostOrderReply) Err() error {
	return r.reply.Err()
}

func (r *PostOrderReply) Get() (resp *PostOrderResp, err error) {
	resp = &PostOrderResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) PostOrder(ctx context.Context, param *PostOrderParam) *PostOrderReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, PostOrderMethod, PostOrderUrl, param)
	}
	return &PostOrderReply{
		reply: reply,
	}
}
