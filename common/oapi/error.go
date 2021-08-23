package oapi

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
)

type APIError struct {
	Code    int
	Message string
	Status  int
}

func NewError(code int, status int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

func NewValidationError(message string) *APIError {
	return &APIError{
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func SendErr(w http.ResponseWriter, apierr *APIError) error {
	w.Header().Set("Content-Type", "text/plain")
	if apierr.Code != 0 {
		w.Header().Set("Error-Code", strconv.Itoa(apierr.Code))
	}
	w.WriteHeader(apierr.Status)
	_, err := w.Write([]byte(apierr.Message))
	return err
}

// Shortcut that sends error response immediately
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	log.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}

func Forbidden(w http.ResponseWriter) {
	ClientError(w, http.StatusForbidden)
}
