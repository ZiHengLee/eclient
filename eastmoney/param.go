package eastmoney

import (
	"context"
)

type IParam interface {
}

type BaseParam struct {
}

func (p *BaseParam) Prepare(ctx context.Context) (err error) {
	return
}
