package httputils

import (
	"encoding/xml"
	"net/http"
)

type PaginationXML struct {
	XMLName  xml.Name    `xml:"pagination"`
	Items    interface{} `xml:"items"`
	MetaData MetaData    `xml:"metadata"`
}

// NewPaginationXML creates a Pagination object for XML response.
func NewPaginationXML[T any](r *http.Request, items []T, count int64, currentPage int, limit int) (*PaginationXML, error) {
	pagination, err := NewPagination(r, items, count, currentPage, limit)
	if err != nil {
		return nil, err
	}
	return pagination.ToXML(), nil
}

func (p *Pagination[T]) ToXML() *PaginationXML {
	xmlData := PaginationXML{
		Items:    p.Items,
		MetaData: p.MetaData,
	}
	return &xmlData
}