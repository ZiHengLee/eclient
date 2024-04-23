package eastmoney

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ApiKey struct {
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
	fmt.Println(url)
	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	return
}
