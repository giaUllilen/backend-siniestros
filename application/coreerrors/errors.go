package coreerrors

import (
	"errors"
	"strings"
)

type CvcError struct {
	Code string
	Err  error
}

func NewCvcError(code string, message ...string) *CvcError {

	errorMsg := strings.Join(message,",")

	return &CvcError{
		Code: code,
		Err:  errors.New(errorMsg),
	}
}

func (e *CvcError) Error() string {
	return e.Err.Error()
}

func (e *CvcError) Unwrap() error { return e.Err }
