package json

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

type Response struct {
	Status int   `json:"status"`
	Error  Error `json:"error,omitempty"`
	Data   any   `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(Response{
		Status: status,
		Data:   data,
	})
}

func WriteError(w http.ResponseWriter, status int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(Response{
		Status: status,
		Error: Error{
			Message: message,
		},
	})
}
