package client

import (
	"fmt"
	"io"
	"net/http"

	"vultr-stat/pkg/config/conf"
	"vultr-stat/pkg/config/consts"
	"vultr-stat/pkg/lib/json"
	"vultr-stat/pkg/lib/web/request"
	urlutil "vultr-stat/pkg/lib/web/url"
)

type Client struct {
}

func NewVultrClient() *Client {
	return &Client{}
}

func (ac *Client) Os() (*OsResponse, error) {
	res, err := RequestHttp(nil, nil, GetHeader(), http.MethodGet, url("/os"))
	if err != nil {
		return nil, err
	}
	return json.ReadReader(res.Body, new(OsResponse))
}

func (ac *Client) Account() (*AccountResponse, error) {
	res, err := RequestHttp(nil, nil, GetHeader(), http.MethodGet, url("/account"))
	if err != nil {
		return nil, err
	}
	return json.ReadReader(res.Body, new(AccountResponse))
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

type QueryString interface {
	ToQueryString() string
}

func urlQs(path string, qs QueryString) string {
	return urlutil.GetUrlQs(consts.Endpoint, path, qs.ToQueryString())
}

func GetHeader() map[string]string {
	headers := make(map[string]string)
	cf := conf.GetConf()
	headers["Authorization"] = fmt.Sprintf("Bearer %s", cf.VultrApiKey)

	return headers
}

func RequestHttp(body io.Reader, err error, headers map[string]string, method string, url string) (*http.Response, error) {
	res, err := request.Request(method, url, headers, body, err)
	if err != nil {
		return nil, err
	}
	if err := CheckErrorResponse(res); err != nil {
		return nil, err
	}
	return res, err
}
