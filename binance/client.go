package binance

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/ZiHengLee/eclient/utils/crypto"
	"github.com/ZiHengLee/eclient/utils/httpclient"
	"github.com/ZiHengLee/eclient/utils/logger"
)

type Client struct {
	auth *ApiKey
	cli  *httpclient.HttpClient
}

func NewClient(key, secret string) (c *Client) {
	cli, err := httpclient.NewHttpClient(&httpclient.Option{})
	if err != nil {
		logger.Error("new http client err:%v", err)
		return
	}
	asecret, err := decrypt(secret)
	if err != nil {
		logger.Error("binance decrypt secret:%v err:%v", key, err)
		return
	}
	return &Client{
		auth: &ApiKey{
			Key:    key,
			secret: asecret,
		},
		cli: cli,
	}
}

func decrypt(s string) (ret []byte, err error) {
	idx := strings.Index(s, ":")
	if idx > 0 {
		var kval string
		if gOpt != nil {
			kval = gOpt.CryptoKeys[s[:idx]]
		}
		if len(kval) == 0 {
			logger.Error("can't find binance cryptokey for %v", s)
			return
		}
		ckey := sha256.Sum256([]byte(kval))
		c := crypto.NewAESCoder(ckey[:16])
		var key []byte
		key, err = base64.StdEncoding.DecodeString(s[idx+1:])
		if err != nil {
			return
		}
		buf, err1 := c.AES128CBCDecrypt(key)
		if err1 != nil {
			err = err1
			return
		}
		ret = buf
	} else {
		ret = []byte(s)
	}
	return
}

func (c *Client) send(ctx context.Context, req *http.Request) (reply *httpclient.Reply) {
	logger.Info("binance client send:%v %v", req.Method, req.URL.String())
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
