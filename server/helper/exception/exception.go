package exception

type ValidationError struct {
	Err error
}

func NewValidationError(err error) *ValidationError {
	return &ValidationError{err}
}

func (e *ValidationError) Error() string {
	return e.Err.Error()
}
