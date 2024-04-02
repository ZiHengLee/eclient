package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ApiKey struct {
	Key string `json:"key" yaml:"key" toml:"key"`

	Secret []byte
}

func (k ApiKey) Sign(body string) string {
	mac := hmac.New(sha256.New, k.Secret)
	mac.Write([]byte(body))
	sum := mac.Sum(nil)
	return hex.EncodeToString(sum)
}

func (k ApiKey) BuildQuery(param IParam) (q string, err error) {
	v, err := query.Values(param)
	if err != nil {
		return
	}
	q = v.Encode()
	if len(q) == 0 {
		return
	}
	if param.RequireAuth() {
		s := k.Sign(q)
		q += "&signature=" + s
	}
	return
}

func (k ApiKey) NewRequest(method, uri string, param IParam) (req *http.Request, err error) {
	url := uri
	if param != nil {
		q, e := k.BuildQuery(param)
		if e != nil {
			return nil, e
		}
		if len(q) > 0 {
			url += "?" + q
		}
	}
	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	req.Header.Add("X-MBX-APIKEY", k.Key)
	return
}
