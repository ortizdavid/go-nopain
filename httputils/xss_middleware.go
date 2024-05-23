package httputils

import (
	"net/http"
	"time"
)

type XssMiddleware struct {
	tokenLength int
	expiration  time.Duration
	cookieName  string
	headerName  string
}

func NewXssMiddleware() *XssMiddleware {
	return &XssMiddleware{}
}

func (csrf *XssMiddleware) Apply(next http.Handler) http.Handler {
	return nil
}