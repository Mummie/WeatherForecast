package main

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func sendErrorResponse(w http.ResponseWriter, err error, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var errorResponse ErrorResponse
	switch true {
	case errors.Is(err, ErrInvalidGeoPoints):
		errorResponse = ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    ErrInvalidGeoPoints.Error(),
		}
	case errors.Is(err, ErrInvalidResponse):
		errorResponse = ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    ErrInvalidResponse.Error(),
		}
	case errors.Is(err, ErrInvalidCoordinates):
		errorResponse = ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    ErrInvalidCoordinates.Error(),
		}
	default:
		errorResponse = ErrorResponse{
			StatusCode: statusCode,
			Message:    message,
		}
	}
	w.WriteHeader(errorResponse.StatusCode)
	json.NewEncoder(w).Encode(errorResponse)
}
