package services

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	data, _ := json.Marshal(v)
	if len(data) > 0 {
		w.Write(data)
	}
}

func WriteErrorJSON(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, ErrorResponse{
		Status:  "error",
		Message: message,
	})
}