// 本代码由脚本根据getfunding_asset.proto.go和tmpl_api.go文件自动生成
// 请不要直接修改本文件

package binance

import (
	"context"
)

type GetFundingAssetReply struct {
	reply *Reply
}

func (r *GetFundingAssetReply) Err() error {
	return r.reply.Err()
}

func (r *GetFundingAssetReply) Get() (resp *GetFundingAssetResp, err error) {
	resp = &GetFundingAssetResp{}
	err = r.reply.Get(&resp.BaseResp, resp)
	return
}

func (c Client) GetFundingAsset(ctx context.Context, param *GetFundingAssetParam) *GetFundingAssetReply {
	reply := &Reply{}
	if param != nil {
		reply.err = param.Prepare(ctx)
	}
	if reply.err == nil {
		reply.reply, reply.err = c.sendParam(ctx, GetFundingAssetMethod, GetFundingAssetUrl, param)
	}
	return &GetFundingAssetReply{
		reply: reply,
	}
}
