package usd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/ZiHengLee/eclient/utils/httpclient"
	"github.com/ZiHengLee/eclient/utils/logger"
)

var (
	ErrInvalidJsonResp = errors.New("invalid json response")
)

type IResp interface {
	GetData() interface{}
}

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`

	UsedWeight map[string]int64 `json:"-"`
}

func (r *BaseResp) GetData() interface{} {
	return nil
}

type Reply struct {
	err   error
	reply *httpclient.Reply
}

func (r Reply) Err() error {
	if r.err != nil {
		return r.err
	}
	return r.reply.Err()
}

func (r Reply) Get(base *BaseResp, resp IResp) (err error) {
	if r.err != nil {
		return r.err
	}
	res, err := r.reply.Get()
	if err != nil {
		return
	}

	logger.Info("binance resp:%#v", *res)
	if base != nil {
		for k, v := range res.Header {
			if strings.HasPrefix(k, "X-Mbx-Used-Weight") && len(v) > 0 {
				w, err := strconv.ParseInt(v[0], 10, 64)
				if err != nil {
					logger.Warn("parse response header:%v err:%v", k, err)
				} else {
					if base.UsedWeight == nil {
						base.UsedWeight = make(map[string]int64)
					}
					base.UsedWeight[k] = w
				}
			}
		}
	}
	body := res.Body
	defer body.Close()
	dat, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	logger.Info("binance resp body:%v", string(dat))
	if len(dat) > 0 {
		err = ErrInvalidJsonResp
		if resp != nil {
			if dat[0] == '{' {
				err = json.Unmarshal(dat, resp)
			} else if d := resp.GetData(); d != nil {
				err = json.Unmarshal(dat, d)
			}
		}
	} else if res.StatusCode != 200 {
		err = fmt.Errorf("http response code:%v", res.StatusCode)
	}
	return
}
