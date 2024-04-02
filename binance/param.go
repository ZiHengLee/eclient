package binance

import (
	"context"

	"github.com/slowly-richer/eclient/utils/time"
)

type IParam interface {
	RequireAuth() bool
}

type BaseParam struct {
	RecvWindow int64 `url:"recvWindow,omitempty"`
	Timestamp  int64 `url:"timestamp,omitempty"`
}

func (p BaseParam) RequireAuth() bool {
	return true
}

func (p *BaseParam) Prepare(ctx context.Context) (err error) {
	if p.Timestamp <= 0 {
		p.Timestamp = time.UnixMilli()
	}
	return
}
