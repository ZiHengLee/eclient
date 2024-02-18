// 本代码由脚本根据deleteorder.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type DeleteOrderReply struct {
	reply *Reply
}

func (r *DeleteOrderReply) Err() error {
	return r.reply.Err()
}

func (r *DeleteOrderReply) Get() (resp *DeleteOrderResp, err error) {
	resp = &DeleteOrderResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) DeleteOrder(ctx context.Context, param *DeleteOrderParam) *DeleteOrderReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, DeleteOrderMethod, DeleteOrderUrl, param)
	}
	return &DeleteOrderReply{
		reply: reply,
	}
}
