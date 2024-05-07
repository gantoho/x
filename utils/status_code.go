package utils

import "net/http"

const (
	OK                  = http.StatusOK
	NotFound            = http.StatusNotFound
	BadRequest          = http.StatusBadRequest
	Unauthorized        = http.StatusUnauthorized
	InternalServerError = http.StatusInternalServerError
	MovedPermanently    = http.StatusMovedPermanently
)
