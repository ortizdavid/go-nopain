package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonResponse struct {
	Message string `json:"message,omitempty"`
	Status int `json:"status"`
	Count *int `json:"count,omitempty"`
	Data any `json:"data,omitempty"`
}

func WriteJson(w http.ResponseWriter, statusCode int, data any, count int) {
	writeJsonHeader(w, statusCode)
	response := jsonResponse{
		Status:  statusCode,
		Count:   &count,
		Data:    data,
	}
	encodeJson(w, response)
}

func WriteJsonSimple(w http.ResponseWriter, statusCode int, data any) {
	writeJsonHeader(w, statusCode)
	response := jsonResponse{
		Status:  statusCode,
		Data:    data,
	}
	encodeJson(w, response)
}

func WriteJsonError(w http.ResponseWriter, message string, statusCode int) {
	writeJsonHeader(w, statusCode)
	response := jsonResponse{
		Message: message,
		Status:  statusCode,
	}
	encodeJson(w, response)
}

func writeJsonHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
}

func encodeJson(w http.ResponseWriter, response jsonResponse) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}