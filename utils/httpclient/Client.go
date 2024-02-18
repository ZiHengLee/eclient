package httpclient

import (
	"net/http"
)

type HttpClient struct {
	opt    *Option
	client *http.Client
}

func NewHttpClient(opt *Option) (c *HttpClient, err error) {
	cli := &http.Client{}
	c = &HttpClient{
		opt:    opt,
		client: cli,
	}
	return
}

func (c HttpClient) Send(req *http.Request) (reply *Reply) {
	reply = &Reply{
		ch: make(chan *replyBody, 1),
	}
	go func() {
		res, err := c.client.Do(req)
		body := &replyBody{
			err: err,
			res: res,
		}
		tgt := c.opt.Host
		if len(tgt) == 0 {
			tgt = req.Host
		}
		if len(tgt) == 0 {
			tgt = req.URL.Host
		}

		reply.ch <- body
		close(reply.ch)
	}()
	return
}
