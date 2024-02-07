package errs

import "net/http"

type AppError struct {
	Code    int `json:",omitempty"`
	Message string
}

func (a AppError) Asmessage() *AppError {
	return &AppError{
		Message: a.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func Newunexpectederror(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
