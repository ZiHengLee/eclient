// 本代码由脚本根据api_restrictions.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type ApiRestrictionsReply struct {
	reply *Reply
}

func (r *ApiRestrictionsReply) Err() error {
	return r.reply.Err()
}

func (r *ApiRestrictionsReply) Get() (resp *ApiRestrictionsResp, err error) {
	resp = &ApiRestrictionsResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) ApiRestrictions(ctx context.Context, param *ApiRestrictionsParam) *ApiRestrictionsReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, ApiRestrictionsMethod, ApiRestrictionsUrl, param)
	}
	return &ApiRestrictionsReply{
		reply: reply,
	}
}
