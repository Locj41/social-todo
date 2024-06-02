package common

import (
	"net/http"

)

type appError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewFullErrResponse(statusCode int, root error, msg string, log string, key string) *appError {
	return &appError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(root error, msg string, log string, key string) *appError {
	return &appError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnthorized(root error, msg string, log string, key string) *appError {
	return &appError{
		StatusCode: http.StatusUnauthorized,
		Message: msg,
		Log: log,
		Key: key,
	}
}

func (e *appError) RootError() error {
	if err, ok := e.RootErr.(*appError); ok {
		return err.RootError()
	}

	return e.RootError()
}

func (e *appError) Error() string {
	return e.RootErr.Error()
}