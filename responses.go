package helper

import (
	"encoding/json"
	"net/http"
)

const (
	ContentTypeHeader     = "Content-Type"
	ApplicationJSON       = "application/json"
	ErrorLabel            = "error"
	MessageLabel          = "message"
	ErrorEncodingResponse = "error encoding response"
)

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set(ContentTypeHeader, ApplicationJSON)
	w.WriteHeader(statusCode)
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{ErrorLabel: message}); encodeErr != nil {
		LogError(encodeErr, ErrorEncodingResponse)
	}
}

func RespondWithMessage(w http.ResponseWriter, message string) {
	w.Header().Set(ContentTypeHeader, ApplicationJSON)
	if encodeErr := json.NewEncoder(w).Encode(map[string]string{MessageLabel: message}); encodeErr != nil {
		LogError(encodeErr, ErrorEncodingResponse)
	}
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set(ContentTypeHeader, ApplicationJSON)
	w.WriteHeader(statusCode)
	if encodeErr := json.NewEncoder(w).Encode(data); encodeErr != nil {
		LogError(encodeErr, ErrorEncodingResponse)
	}
}
