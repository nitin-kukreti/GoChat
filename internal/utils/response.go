package utils

import (
	"encoding/json"
	"net/http"
)

type CustomResponse[T any] struct {
	Data       T      `json:"data"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func WriteJSON[T any](w http.ResponseWriter, status int, message string, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := CustomResponse[T]{
		Data:       data,
		Message:    message,
		StatusCode: status,
	}
	json.NewEncoder(w).Encode(resp)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, message, struct{}{})
}
