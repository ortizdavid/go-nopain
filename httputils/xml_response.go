package httputils

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ortizdavid/go-nopain/serialization"
)

// xmlResponse represents the structure of an XML response.
type xmlResponse struct {
	XMLName xml.Name `xml:"response"`            // XML root element name
	Message *string  `xml:"message,omitempty"`   // Message field for the response (optional)
	Status  int      `xml:"status"`               // Status code of the response
	Count   *int     `xml:"count,omitempty"`     // Count field for the response (optional)
	Data    any      `data:"data,omitempty"`     // Data field for the response (optional)
}

// WriteXmlPaginated writes a paginated XML response to the provided http.ResponseWriter.
// It includes the provided status code, paginated items, pagination metadata,
// and handles potential errors during pagination or XML encoding.
func WriteXmlPaginated[T any](w http.ResponseWriter, r *http.Request, items []T, count int64, currentPage int, limit int) {
	writeXmlHeader(w, 200)
	paginationXML, err := NewPaginationXML(r, items, count, currentPage, limit)
	if err != nil {
		WriteXmlError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := serialization.EncodeXml(w, paginationXML); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// WriteXml writes a simple XML response with the provided status code and data.
func WriteXml(w http.ResponseWriter, statusCode int, data any) {
	writeXmlHeader(w, statusCode)
	response := xmlResponse{
		Status: statusCode,
		Data:   data,
	}
	encodeXml(w, response)
}

// WriteXmlError writes an XML error response with the provided message and status code.
func WriteXmlError(w http.ResponseWriter, message string, statusCode int) {
	writeXmlHeader(w, statusCode)
	response := xmlResponse{
		Message: &message,
		Status:  statusCode,
	}
	encodeXml(w, response)
}

// writeXmlHeader writes the XML content type header and sets the HTTP status code.
func writeXmlHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-type", "application/xml")
	w.WriteHeader(statusCode)
}

// encodeXml encodes the xmlResponse struct to XML format and writes it to the response writer.
func encodeXml(w http.ResponseWriter, response xmlResponse) {
	err := serialization.EncodeXml(w, response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
