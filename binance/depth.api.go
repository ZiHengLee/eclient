// 本代码由脚本根据depth.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type DepthReply struct {
	reply *Reply
}

func (r *DepthReply) Err() error {
	return r.reply.Err()
}

func (r *DepthReply) Get() (resp *DepthResp, err error) {
	resp = &DepthResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) Depth(ctx context.Context, param *DepthParam) *DepthReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, DepthMethod, DepthUrl, param)
	}
	return &DepthReply{
		reply: reply,
	}
}
