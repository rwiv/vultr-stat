package client

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/rwiv/vultr-stat/pkg/common"
	"github.com/rwiv/vultr-stat/pkg/lib/json"
	"github.com/rwiv/vultr-stat/pkg/lib/web/request"
	urlutil "github.com/rwiv/vultr-stat/pkg/lib/web/url"
)

type VultrClient struct {
	ApiKey string
}

func NewVultrClient(apiKey string) *VultrClient {
	return &VultrClient{
		ApiKey: apiKey,
	}
}

func (vc *VultrClient) Os() (*OsResponse, error) {
	res, err := RequestHttp(nil, nil, vc.GetHeader(), http.MethodGet, url("/os"))
	if err != nil {
		return nil, err
	}
	return json.ReadReader(res.Body, new(OsResponse))
}

func (vc *VultrClient) Account() (*AccountResponse, error) {
	res, err := RequestHttp(nil, nil, vc.GetHeader(), http.MethodGet, url("/account"))
	if err != nil {
		return nil, err
	}
	return json.ReadReader(res.Body, new(AccountResponse))
}

func (vc *VultrClient) Instances(withUsedBandwidth bool) ([]*InstanceInfo, error) {
	res, err := RequestHttp(nil, nil, vc.GetHeader(), http.MethodGet, url("/instances"))
	if err != nil {
		return nil, err
	}
	raw, err := json.ReadReader(res.Body, new(InstanceResponse))
	if err != nil {
		return nil, err
	}

	instances := make([]*InstanceInfo, 0)

	var wg sync.WaitGroup
	results := make(chan *InstanceInfo, len(raw.Instances))
	for _, instance := range raw.Instances {
		wg.Add(1)
		go vc.solveInstanceInfo(instance, withUsedBandwidth, &wg, results)
	}
	wg.Wait()
	close(results)
	for result := range results {
		instances = append(instances, result)
	}

	return instances, nil
}

func (vc *VultrClient) solveInstanceInfo(
	instance *Instance,
	withUsedBandwidth bool,
	wg *sync.WaitGroup,
	results chan<- *InstanceInfo,
) {
	defer wg.Done()
	var usedBandwidth *int64
	if withUsedBandwidth {
		bres, err := vc.Bandwidth(instance.Id)
		if err != nil {
			fmt.Println(err)
		}
		tmp := int64(bres.Sum().ToGb().OutgoingBytes)
		usedBandwidth = &tmp
	}
	results <- instance.ToInfo(usedBandwidth)
}

func (vc *VultrClient) Bandwidth(instanceId string) (*InstanceBandwidthResponse, error) {
	paths := fmt.Sprintf("/instances/%s/bandwidth", instanceId)
	res, err := RequestHttp(nil, nil, vc.GetHeader(), http.MethodGet, url(paths))
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

func (vc *VultrClient) GetHeader() map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", vc.ApiKey)

	return headers
}

func RequestHttp(
	body io.Reader,
	err error,
	headers map[string]string,
	method string,
	url string,
) (*http.Response, error) {
	res, err := request.Request(method, url, headers, body, err)
	if err != nil {
		return nil, err
	}
	if err := CheckErrorResponse(res); err != nil {
		return nil, err
	}
	return res, err
}
