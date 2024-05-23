package httputils

import (
	"net/http"
	"time"
)

type CsrfMiddleware struct {
	tokenLength int
	cookieName  string
	headerName  string
	expiration  time.Duration
}

func NewCsrfMiddleware() CsrfMiddleware {
	return CsrfMiddleware{}
}

func (csrf CsrfMiddleware) Apply(next http.Handler) http.Handler {
	return nil
}

func (csrf CsrfMiddleware) CreateToken(w http.ResponseWriter) (string, error) {
	return "", nil
}

func (csrf CsrfMiddleware) ValidateToken(w http.ResponseWriter) error {
	return nil
}