package cutil

import (
	"io"
	"net/http"

	ce "vultr-stat/pkg/client/error"
	"vultr-stat/pkg/lib/web/request"
)

func RequestHttp(body io.Reader, err error, method string, url string) (*http.Response, error) {
	return RequestHttpHeaders(body, err, GetHeaders(), method, url)
}

func RequestHttpHeaders(body io.Reader, err error, headers map[string]string, method string, url string) (*http.Response, error) {
	res, err := request.Request(method, url, headers, body, err)
	if err != nil {
		return nil, err
	}
	if err := ce.CheckErrorResponse(res); err != nil {
		return nil, err
	}
	return res, err
}
