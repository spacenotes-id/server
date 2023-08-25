package exception

import "fmt"

type HttpError struct {
	message    string
	statusCode uint
}

func (h *HttpError) Error() string {
	return h.message
}

func (h *HttpError) StatusCode() uint {
	return h.statusCode
}

func NewHTTPError(statusCode uint, message string, args ...any) *HttpError {
	msg := fmt.Sprintf(message, args...)
	return &HttpError{msg, statusCode}
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
