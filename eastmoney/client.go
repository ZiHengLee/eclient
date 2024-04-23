package eastmoney

import (
	"context"
	"net/http"

	"github.com/slowly-richer/eclient/utils/httpclient"
	"github.com/slowly-richer/eclient/utils/logger"
)

type Client struct {
	auth *ApiKey
	cli  *httpclient.HttpClient
}

func NewClient() (c *Client) {
	cli, err := httpclient.NewHttpClient(&httpclient.Option{})
	if err != nil {
		logger.Error("new http client err:%v", err)
		return
	}
	return &Client{
		auth: &ApiKey{},
		cli:  cli,
	}
}

func (c *Client) send(ctx context.Context, req *http.Request) (reply *httpclient.Reply) {
	// logger.Info("binance client send:%v %v", req.Method, req.URL.String())
	reply = c.cli.Send(req)
	return
}

func (c *Client) sendParam(ctx context.Context, method, url string, param IParam) (reply *httpclient.Reply, err error) {
	req, err := c.auth.NewRequest(method, url, param)
	if err != nil {
		logger.Warn("create request err:%v", err)
		return nil, err
	}
	return c.send(ctx, req), nil
}
