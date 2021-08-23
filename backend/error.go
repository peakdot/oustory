package main

import (
	"net/http"

	"github.com/peakdot/oustory/common/oapi"
)

var (
	ErrInvalidAmount = oapi.NewError(1001, http.StatusBadRequest, "Invalid amount. Must be greater than 0.")
)
