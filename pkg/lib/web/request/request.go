package request

import (
	"io"
	"net/http"
)

func Request(method string, url string, headers map[string]string, body io.Reader, err error) (*http.Response, error) {
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return c.Do(req)
}
