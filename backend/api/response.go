package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorData represents the standardized error response details.
type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Response represents the standardized API response shape.
type Response struct {
	Data  interface{} `json:"data"`
	Error *ErrorData  `json:"error"`
}

// SendSuccess writes a successful JSON response matching the API contract.
func SendSuccess(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response{
		Data:  data,
		Error: nil,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding success response: %v", err)
	}
}

// SendError writes a failed JSON response matching the API contract.
func SendError(w http.ResponseWriter, status int, code string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response{
		Data: nil,
		Error: &ErrorData{
			Code:    code,
			Message: message,
		},
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding error response: %v", err)
	}
}
