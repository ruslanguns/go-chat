package errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrInternal      = errors.New("internal error")
	ErrAlreadyExists = errors.New("already exists")
)

type AppError struct {
	Err error
	Msg string
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Msg, e.Err)
}

func (e AppError) ErrorType() error {
	return e.Err
}

func NewAppError(err error, msg string) AppError {
	return AppError{Err: err, Msg: msg}
}
