package httputils

import (
	"fmt"
	"net/http"
	"github.com/ortizdavid/go-nopain/serialization"
)

// jsonResponse represents the structure of a JSON response.
type jsonResponse struct {
	Message *string `json:"message,omitempty"` // Message field for the response (optional)
	Status  int     `json:"status"`             // Status code of the response
	Count   *int    `json:"count,omitempty"`   // Count field for the response (optional)
	Data    any     `json:"data,omitempty"`    // Data field for the response (optional)
}

// WriteJson writes a simple JSON response with the provided status code and data.
func WriteJson(w http.ResponseWriter, statusCode int, data any) {
	writeJsonHeader(w, statusCode)
	response := jsonResponse{
		Status: statusCode,
		Data:   data,
	}
	encodeJson(w, response)
}

// WriteJsonSimple writes a simple JSON response with the provided status code
func WriteJsonSimple(w http.ResponseWriter, statusCode int, data any) {
	writeJsonHeader(w, statusCode)
	encodeJson(w, data)
}


// WriteJsonPaginated writes a paginated JSON response to the provided http.ResponseWriter.
// It includes the paginated items, pagination metadata,
// and handles potential errors during pagination or JSON encoding.
func WriteJsonPaginated[T any](w http.ResponseWriter, r *http.Request, items []T, count int, currentPage int, limit int) {
	writeJsonHeader(w, 200)
	pagination, err := NewPagination(r, items, count, currentPage, limit)
	if err != nil {
		WriteJsonError(w, err.Error(), http.StatusInternalServerError)
	}
	if err := serialization.EncodeJson(w, pagination); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// WriteJsonError writes a JSON error response with the provided message and status code.
func WriteJsonError(w http.ResponseWriter, message string, statusCode int) {
	writeJsonHeader(w, statusCode)
	response := jsonResponse{
		Message: &message,
		Status:  statusCode,
	}
	encodeJson(w, response)
}

// writeJsonHeader writes the JSON content type header and sets the HTTP status code.
func writeJsonHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
}

// encodeJson encodes the jsonResponse struct to JSON format and writes it to the response writer.
func encodeJson(w http.ResponseWriter, response interface{}) {
	err := serialization.EncodeJson(w, response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
