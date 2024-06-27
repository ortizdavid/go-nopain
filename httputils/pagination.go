package httputils

import (
	"encoding/xml"
	"fmt"
	"math"
	"net/http"

	"github.com/ortizdavid/go-nopain/conversion"
)

type Pagination[T any] struct {
	Items           []T `json:"items"`
	MetaData			MetaData `json:"pagination"`
}

type MetaData struct {
	XMLName         xml.Name `json:"-" xml:"metadata"`
	CurrentPage     int      `json:"current_page" xml:"currentPage"`
	TotalItems      int      `json:"total_items" xml:"totalItems"`
	TotalPages      int      `json:"total_pages" xml:"totalPages"`
	FirstPageUrl    string   `json:"first_page_url" xml:"firstPageUrl"`
	PreviousPageUrl string   `json:"previous_page_url" xml:"previousPageUrl"`
	NextPageUrl     string   `json:"next_page_url" xml:"nextPageUrl"`
	LastPageUrl     string   `json:"last_page_url" xml:"lastPageUrl"`
}

func NewPagination[T any](r *http.Request, items []T, count int, currentPage int, limit int) (*Pagination[T], error) {
	if currentPage < 0 {
		return nil, fmt.Errorf("current page must be >= 0")
	}
	if limit < 1 {
		return nil, fmt.Errorf("page size must be >= 1")
	}
	pagination := Pagination[T]{
		Items:      items,
		MetaData:	MetaData{
			CurrentPage:     currentPage,
			TotalItems:      count,
			TotalPages:      int(math.Ceil(float64(count) / float64(limit))),
			FirstPageUrl: 	 "",
			PreviousPageUrl: "",
			NextPageUrl:     "",
			LastPageUrl:	 "",
		},
	}
	pagination.calculateUrls(r, currentPage, limit)
	return &pagination, nil
}

func (p *Pagination[T]) HasNextPage() bool {
	return p.MetaData.CurrentPage < p.MetaData.TotalPages - 1
}

func (p *Pagination[T]) HasPreviousPage() bool {
	return p.MetaData.CurrentPage > 0
}

func (p *Pagination[T]) calculateUrls(r *http.Request, currentPage, limit int) {
	baseUrl := getRequestBaseUrl(r)
	
	p.MetaData.FirstPageUrl = getPageUrl(baseUrl, 1, limit)
	if currentPage < p.MetaData.TotalPages {
		p.MetaData.NextPageUrl = getPageUrl(baseUrl, currentPage+1, limit)
	}
	if currentPage > 1 {
		p.MetaData.PreviousPageUrl = getPageUrl(baseUrl, currentPage-1, limit)
	}
	p.MetaData.LastPageUrl = getPageUrl(baseUrl, p.MetaData.TotalPages, limit)
}

func getRequestBaseUrl(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host + r.URL.Path
}

func getPageUrl(baseUrl string, pageNumber, limit int) string {
	return baseUrl + "?current_page=" + conversion.IntToString(pageNumber) + "&limit=" + conversion.IntToString(limit)
}