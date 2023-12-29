package error

type HttpError struct {
	Status    int
	Message   string
	Uuid      string
	Timestamp string
	Exception string
}

func (e *HttpError) Error() string {
	return e.Message
}
