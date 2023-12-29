package error

import (
	"fmt"
	"net/http"
)

type RedirectError struct {
	Status      int
	RedirectUri string
}

func NewRedirectError(res *http.Response) *RedirectError {
	return &RedirectError{
		Status:      res.StatusCode,
		RedirectUri: res.Header.Get("Location"),
	}
}

func (e *RedirectError) Error() string {
	return fmt.Sprintf("please go to %s", e.RedirectUri)
}
