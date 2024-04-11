package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// hateoasJson represents the structure of a HATEOAS JSON response.
type hateoasJson struct {
	Message *string `json:"message,omitempty"` // Message field for the response (optional)
	Status  int     `json:"status"`             // Status code of the response
	Count   *int    `json:"count,omitempty"`   // Count field for the response (optional)
	Data    any     `json:"data,omitempty"`    // Data field for the response (optional)
	Links   link    `json:"links"`             // Links field for HATEOAS
}

// link represents the structure of a link in a HATEOAS JSON response.
type link struct {
	Path string `json:"path"` // Path of the link
	Self string `json:"self"` // Self link
	Rel  string `json:"rel,omitempty"` // Relationship of the link (optional)
}

// WriteHateoasJson writes a HATEOAS JSON response with the provided status code, data, and count.
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

// encodeHateoas encodes the hateoasJson struct to JSON format and writes it to the response writer.
func encodeHateoas(w http.ResponseWriter, response hateoasJson) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
