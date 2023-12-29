package client

import (
	"net/http"
	"vultr-stat/pkg/client/client/types"

	"vultr-stat/pkg/client/consts"
	"vultr-stat/pkg/lib/cutil"
	"vultr-stat/pkg/lib/json"
	urlutil "vultr-stat/pkg/lib/web/url"
)

type Client struct {
}

func NewVultrClient() *Client {
	return &Client{}
}

func (ac *Client) Os() (*types.OsResponse, error) {
	res, err := cutil.RequestHttp(nil, nil, http.MethodGet, url("/os"))
	if err != nil {
		return nil, err
	}
	return json.ReadReader(res.Body, new(types.OsResponse))
}

//func (ac *Client) Oses() ([]*OsResponse, error) {
//	res, err := cutil.RequestHttp(nil, nil, http.MethodGet, url("/os"))
//	if err != nil {
//		return nil, err
//	}
//	return json.ReadReaderSlice(res.Body, make([]*OsResponse, 0))
//}

func url(path string) string {
	return urlutil.GetUrl(consts.Endpoint, path)
}

func urlQs(path string, qs cutil.QueryString) string {
	return urlutil.GetUrlQs(consts.Endpoint, path, qs.ToQueryString())
}
