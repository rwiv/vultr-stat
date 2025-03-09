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

func (ac *Client) Instances(withUsage bool) ([]*InstanceInfo, error) {
	res, err := RequestHttp(nil, nil, ac.GetHeader(), http.MethodGet, url("/instances"))
	if err != nil {
		return nil, err
	}
	raw, err := json.ReadReader(res.Body, new(InstanceResponse))
	if err != nil {
		return nil, err
	}
	instances := make([]*InstanceInfo, 0)
	for _, instance := range raw.Instances {
		var usedBandwidth *int64
		if withUsage {
			bres, err := ac.Bandwidth(instance.Id)
			if err != nil {
				return nil, err
			}
			tmp := int64(bres.Sum().ToGb().OutgoingBytes)
			usedBandwidth = &tmp
		}
		instances = append(instances, instance.ToInfo(usedBandwidth))
	}
	return instances, nil
}

func (ac *Client) Bandwidth(instanceId string) (*InstanceBandwidthResponse, error) {
	paths := fmt.Sprintf("/instances/%s/bandwidth", instanceId)
	res, err := RequestHttp(nil, nil, ac.GetHeader(), http.MethodGet, url(paths))
	if err != nil {
		return nil, err
	}
	return json.ReadReader(res.Body, new(InstanceBandwidthResponse))
}

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
