package binance

import (
	"context"
)

const PostUserDataStreamMethod = "POST"

const PostUserDataStreamUrl = "https://api.binance.com/api/v3/userDataStream"

type PostUserDataStreamParam struct {
	BaseParam
}

func (p *PostUserDataStreamParam) Prepare(ctx context.Context) (err error) {
	return
}

func (p *PostUserDataStreamParam) RequireAuth() bool {
	return false
}

type PostUserDataStreamResp struct {
	BaseResp
	ListenKey string `json:"listenKey"`
}

const PutUserDataStreamMethod = "PUT"

const PutUserDataStreamUrl = "https://api.binance.com/api/v3/userDataStream"

type PutUserDataStreamParam struct {
	BaseParam
}

type PutUserDataStreamResp struct {
	BaseResp
}
