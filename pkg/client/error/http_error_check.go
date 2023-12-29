package error

import (
	"net/http"

	"vultr-stat/pkg/lib/json"
)

func CheckErrorResponse(res *http.Response) error {
	if res.StatusCode >= 400 {
		body, err := json.ReadReader(res.Body, new(HttpErrorResponse))
		if err != nil {
			return err
		}
		return body.ToError()
	}
	if res.StatusCode >= 300 {
		err := NewRedirectError(res)
		return err
	}
	return nil
}
