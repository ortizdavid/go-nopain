package httputils

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type xmlResponse struct {
	XMLName xml.Name `xml:"response"`
	Message string `xml:"message"`
	Status  int `xml:"status"`
	Count   int `xml:"count"`
	Data    any `data:"data,omitempty"`
}

func WriteXml(w http.ResponseWriter, statusCode int, data any, count int) {
	writeXmlHeader(w, statusCode)
	response := xmlResponse{
		Status:  statusCode,
		Count:   count,
		Data:    data,
	}
	encodeXml(w, response)
}

func WriteXmlSimple(w http.ResponseWriter, statusCode int, data any) {
	writeXmlHeader(w, statusCode)
	response := xmlResponse{
		Status:  statusCode,
		Data:    data,
	}
	encodeXml(w, response)
}

func WriteXmlError(w http.ResponseWriter, message string, statusCode int) {
	writeXmlHeader(w, statusCode)
	response := xmlResponse{
		Message: message,
		Status:  statusCode,
	}
	encodeXml(w, response)
}

func writeXmlHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-type", "application/xml")
	w.WriteHeader(statusCode)
}

func encodeXml(w http.ResponseWriter, response xmlResponse) {
	err := xml.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
