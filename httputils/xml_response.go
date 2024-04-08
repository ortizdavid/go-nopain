package httputils

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type XmlResponse struct {
	XMLName xml.Name `xml:"response"`
	Message string `xml:"message"`
	Status  int `xml:"status"`
	Count   int `xml:"count"`
	Data    any `data:"data"`
}

func WriteXml(w http.ResponseWriter, r *http.Request, message string, statusCode int, data any, count int) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(statusCode)
	response := XmlResponse{
		Message: message,
		Status:  statusCode,
		Count:   count,
		Data:    data,
	}
	err := xml.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func WriteXmlError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-type", "application/xml")
	w.WriteHeader(statusCode)
	response := XmlResponse{
		Message: message,
		Status:  statusCode,
	}
	err := xml.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}