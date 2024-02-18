// 本代码由脚本根据accountinfo.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type AccountInfoReply struct {
	reply *Reply
}

func (r *AccountInfoReply) Err() error {
	return r.reply.Err()
}

func (r *AccountInfoReply) Get() (resp *AccountInfoResp, err error) {
	resp = &AccountInfoResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) AccountInfo(ctx context.Context, param *AccountInfoParam) *AccountInfoReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, AccountInfoMethod, AccountInfoUrl, param)
	}
	return &AccountInfoReply{
		reply: reply,
	}
}
