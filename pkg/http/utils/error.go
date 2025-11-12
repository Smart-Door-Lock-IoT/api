package httputils

import "net/http"

type Error struct {
	Code    int
	Message string
	Err     error
}

func NewError(code int, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewInternalError(err error) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: "Terjadi kesalahan tak terduga pada server!",
		Err:     err,
	}
}
