package error

type HttpErrorResponse struct {
	Status    int
	Message   string
	Uuid      string
	Timestamp string
	Exception string
}

func (res *HttpErrorResponse) ToError() *HttpError {
	return &HttpError{
		Status:    res.Status,
		Message:   res.Message,
		Uuid:      res.Uuid,
		Timestamp: res.Timestamp,
		Exception: res.Exception,
	}
}
