package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ErrInternal = &Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}

	ErrBadRequest = &Error{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
	}

	ErrGameNotFoundRequest = &Error{
		Code:    http.StatusNotFound,
		Message: "Not Found",
	}

	ErrObjectIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Request object should be provided",
	}
)

type Error struct {
	Code    int
	Message string
}

func (err *Error) Error() string {
	return err.String()
}

func (err *Error) String() string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("Error %d: %s", err.Code, err.Message)
}

func (err *Error) JSON() []byte {
	if err == nil {
		return []byte("{}")
	}

	res, _ := json.Marshal(err)
	return res
}

func (err *Error) StatusCode() int {
	if err == nil {
		return http.StatusOK
	}

	return err.Code
}
