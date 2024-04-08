package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Count int `json:"count"`
	Data any `json:"data"`
}

func WriteJson(w http.ResponseWriter, r *http.Request, message string, statusCode int, data any, count int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := JsonResponse{
		Message: message,
		Status:  statusCode,
		Count:   count,
		Data:    data,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func WriteJsonError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	response := JsonResponse{
		Message: message,
		Status:  statusCode,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
