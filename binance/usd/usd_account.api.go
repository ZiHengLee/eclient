// 本代码由脚本根据usd_account.proto.go和usd_tmpl_api.go文件自动生成
// 请不要直接修改本文件

package usd

import (
	"context"
)

type UsdAccountReply struct {
	reply *Reply
}

func (r *UsdAccountReply) Err() error {
	return r.reply.Err()
}

func (r *UsdAccountReply) Get() (resp *UsdAccountResp, err error) {
	resp = &UsdAccountResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) UsdAccount(ctx context.Context, param *UsdAccountParam) *UsdAccountReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, UsdAccountMethod, UsdAccountUrl, param)
	}
	return &UsdAccountReply{
		reply: reply,
	}
}
