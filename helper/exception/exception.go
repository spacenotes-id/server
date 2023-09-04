package exception

type HttpError struct {
	Error string `json:"error"`
}

type ValErrors struct {
	Errors []string `json:"errors"`
}

type ValidationError struct {
	Err error
}

func NewValidationError(err error) *ValidationError {
	return &ValidationError{err}
}

func (e *ValidationError) Error() string {
	return e.Err.Error()
}
