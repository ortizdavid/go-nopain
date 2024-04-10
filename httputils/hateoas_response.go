package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type hateoasJson struct {
	Message *string `json:"message,omitempty"`
	Status int `json:"status"`
	Count *int `json:"count,omitempty"`
	Data any `json:"data,omitempty"`
	Links link `json:"links"`
}


type link struct {
	Path string `json:"path"`
	Self string `json:"self"`
	Rel string `json:"rel,omitempty"`
}


func WriteHateoasJson(w http.ResponseWriter, r *http.Request, statusCode int, data any, count int) {
	writeJsonHeader(w, statusCode)
	
	basePath := fmt.Sprintf("%s://%s", r.URL.Scheme, r.URL.Host)
	selfLink := basePath + r.URL.Path
	
	response := hateoasJson{
		Status:  statusCode,
		Count:   &count,
		Data:    data,
		Links:   link{
			Path: basePath,
			Self: selfLink,
			Rel:  "",
		},
	}
	encodeHateoas(w, response)
}


func encodeHateoas(w http.ResponseWriter, response hateoasJson) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

