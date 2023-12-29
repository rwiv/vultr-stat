package client

import (
	"io"

	"vultr-stat/pkg/client/consts"
	"vultr-stat/pkg/lib/cutil"
	"vultr-stat/pkg/lib/json"
	urlutil "vultr-stat/pkg/lib/web/url"
)

const (
	prefix = "/api/access"
)

type Client struct {
}

func NewAccessClient() *Client {
	return &Client{}
}

func url(path string) string {
	return urlutil.GetUrl(consts.Endpoint, prefix+path)
}

func urlQs(path string, qs cutil.QueryString) string {
	return urlutil.GetUrlQs(consts.Endpoint, prefix+path, qs.ToQueryString())
}

func requestOne(body io.Reader, err error, method string, url string) (*FileInfo, error) {
	res, err := cutil.RequestHttp(body, err, method, url)
	if err != nil {
		return nil, err
	}
	return json.ReadReader(res.Body, new(FileInfo))
}

func requestList(body io.Reader, err error, method string, url string) ([]*FileInfo, error) {
	res, err := cutil.RequestHttp(body, err, method, url)
	if err != nil {
		return nil, err
	}
	return json.ReadReaderSlice(res.Body, make([]*FileInfo, 0))
}
