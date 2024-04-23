package eastmoney

import (
	"context"
)

////=====header====

type TmplNameReply struct {
	reply *Reply
}

func (r *TmplNameReply) Err() error {
	return r.reply.Err()
}

func (r *TmplNameReply) Get() (resp *TmplNameResp, err error) {
	resp = &TmplNameResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) TmplName(ctx context.Context, param *TmplNameParam) *TmplNameReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, TmplNameMethod, TmplNameUrl, param)
	}
	return &TmplNameReply{
		reply: reply,
	}
}
