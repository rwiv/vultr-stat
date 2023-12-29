package client

import (
	"fmt"
	"net/http"

	"vultr-stat/pkg/lib/json"
)

type HttpError struct {
	Message string `json:"error"`
	Status  int32  `json:"status"`
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Status, e.Message)
}

func CheckErrorResponse(res *http.Response) error {
	if res.StatusCode >= 400 {
		body, err := json.ReadReader(res.Body, new(HttpError))
		if err != nil {
			return err
		}
		return body
	}
	return nil
}
