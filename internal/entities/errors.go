package entities

type RequestError struct {
	Msg        string
	StatusCode int
}

func (e *RequestError) Error() string {
	return e.Msg
}
