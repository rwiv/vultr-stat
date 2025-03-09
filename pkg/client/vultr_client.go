package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rwiv/vultr-stat/pkg/common"
	"github.com/rwiv/vultr-stat/pkg/lib/json"
	"github.com/rwiv/vultr-stat/pkg/lib/web/request"
	urlutil "github.com/rwiv/vultr-stat/pkg/lib/web/url"
)

type Client struct {
	ApiKey string
}

func NewVultrClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
	}
}

func (ac *Client) Os() (*OsResponse, error) {
	res, err := RequestHttp(nil, nil, ac.GetHeader(), http.MethodGet, url("/os"))
	if err != nil {
		return nil, err
	}
	return json.ReadReader(res.Body, new(OsResponse))
}

func (ac *Client) Account() (*AccountResponse, error) {
	res, err := RequestHttp(nil, nil, ac.GetHeader(), http.MethodGet, url("/account"))
	if err != nil {
		return nil, err
	}
	return json.ReadReader(res.Body, new(AccountResponse))
}

func (ac *Client) Instances() (*InstanceResponse, error) {
	res, err := RequestHttp(nil, nil, ac.GetHeader(), http.MethodGet, url("/instances"))
	if err != nil {
		return nil, err
	}
	return json.ReadReader(res.Body, new(InstanceResponse))
}

//func (ac *Client) Oses() ([]*OsResponse, error) {
//	res, err := cutil.RequestHttp(nil, nil, http.MethodGet, url("/os"))
//	if err != nil {
//		return nil, err
//	}
//	return json.ReadReaderSlice(res.Body, make([]*OsResponse, 0))
//}

func url(path string) string {
	return urlutil.GetUrl(common.Endpoint, path)
}

type QueryString interface {
	ToQueryString() string
}

func urlQs(path string, qs QueryString) string {
	return urlutil.GetUrlQs(common.Endpoint, path, qs.ToQueryString())
}

func (ac *Client) GetHeader() map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", ac.ApiKey)

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
