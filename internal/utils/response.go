package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Sucess  bool   `json:"success"`
	Message string `json:"message"`
	Date    any    `json:"data,omitempty"`
}

func SendResponse(w http.ResponseWriter, sucess bool, message string, data any) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{
		Sucess:  sucess,
		Message: message,
		Date:    data,
	})
}
