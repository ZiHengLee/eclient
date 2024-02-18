// 本代码由脚本根据userdatastream.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type PostUserDataStreamReply struct {
	reply *Reply
}

func (r *PostUserDataStreamReply) Err() error {
	return r.reply.Err()
}

func (r *PostUserDataStreamReply) Get() (resp *PostUserDataStreamResp, err error) {
	resp = &PostUserDataStreamResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) PostUserDataStream(ctx context.Context, param *PostUserDataStreamParam) *PostUserDataStreamReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, PostUserDataStreamMethod, PostUserDataStreamUrl, param)
	}
	return &PostUserDataStreamReply{
		reply: reply,
	}
}

type PutUserDataStreamReply struct {
	reply *Reply
}

func (r *PutUserDataStreamReply) Err() error {
	return r.reply.Err()
}

func (r *PutUserDataStreamReply) Get() (resp *PutUserDataStreamResp, err error) {
	resp = &PutUserDataStreamResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) PutUserDataStream(ctx context.Context, param *PutUserDataStreamParam) *PutUserDataStreamReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, PutUserDataStreamMethod, PutUserDataStreamUrl, param)
	}
	return &PutUserDataStreamReply{
		reply: reply,
	}
}
