package httputils

import (
	"fmt"
	"math"
	"net/http"

	"github.com/ortizdavid/go-nopain/conversion"
)

type Pagination[T any] struct {
	Items           []T `json:"items"`
	PageIndex       int `json:"page_index"`
	TotalItems      int `json:"total_items"`
	TotalPages      int `json:"total_pages"`
	PreviousPageUrl string `json:"previous_page_url"`
	NextPageUrl     string `json:"next_page_url"`
}

func NewPagination[T any](r *http.Request, items []T, count int, pageIndex int, pageSize int) (*Pagination[T], error) {
	if pageIndex < 0 {
		return nil, fmt.Errorf("page index must be >= 0")
	}
	if pageSize < 1 {
		return nil, fmt.Errorf("page size must be >= 1")
	}
	pagination := Pagination[T]{
		Items:           items,
		PageIndex:       pageIndex,
		TotalItems:      count,
		TotalPages:      int(math.Ceil(float64(count) / float64(pageSize))),
		PreviousPageUrl: "",
		NextPageUrl:     "",
	}
	pagination.calculateUrls(r, pageIndex, pageSize)
	return &pagination, nil
}

func (p *Pagination[T]) HasNextPage() bool {
	return p.PageIndex < p.TotalPages - 1
}

func (p *Pagination[T]) HasPreviousPage() bool {
	return p.PageIndex > 0
}

func (p *Pagination[T]) calculateUrls(r *http.Request, pageIndex, pageSize int) {
	baseUrl := getRequestBaseUrl(r)

	if pageIndex < p.TotalPages {
		p.NextPageUrl = getPageUrl(baseUrl, pageIndex+1, pageSize)
	}
	if pageIndex > 1 {
		p.PreviousPageUrl = getPageUrl(baseUrl, pageIndex-1, pageSize)
	}
}

func getRequestBaseUrl(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host + r.URL.Path
}

func getPageUrl(baseUrl string, pageNumber, pageSize int) string {
	return baseUrl + "?pageIndex=" + conversion.IntToString(pageNumber) + "&pageSize=" + conversion.IntToString(pageSize)
}