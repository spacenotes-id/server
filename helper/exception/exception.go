package exception

type HttpError struct {
	Error string `json:"error" example:"Something went wrong"`
}

type ValErrors struct {
	Errors []string `json:"errors" example:"Invalid email,Username is required"`
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
