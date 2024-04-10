package httputils

import (
	"errors"
	"net/http"
	"sync"
)

type ApiKeyMiddleware struct {
	defaultKey string
	mu sync.RWMutex
}

// Return a ApiKeyMiddleware object with the apiKey passed
func NewApiKeyMiddleWare(apiKey string) ApiKeyMiddleware {
	return ApiKeyMiddleware{
		defaultKey: apiKey,
	}
}

// Apply Api Key middleware to an handler. Wraps a handler.
// must be used for mux.Handle intead of mux.HandleFunc.
// Example: mux.Handle("GET /protected", protectedHandler).
func (apiMid *ApiKeyMiddleware) Apply(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return apiMid.applyMiddleware(http.HandlerFunc(handler))
}

// Apply Api Key middleware to a handler.
// Get value from header X-API-Key and pass to next request.
// X-API-Key must be valid and non empty.
func (apiMid *ApiKeyMiddleware) applyMiddleware(next http.Handler) http.Handler {
	validApiKey := apiMid.getDefaultKey()

	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "Unauthorized. API Key missing", http.StatusUnauthorized)
			return
		}
		if apiKey != validApiKey {
			http.Error(w, "Unauthorized. Invalid API Key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}


func (apiMid *ApiKeyMiddleware) SetDefaultKey(value string) error {
	if value == "" {
		return errors.New("API key cannot be empty")
	}
	apiMid.mu.Lock()
	apiMid.defaultKey = value
	apiMid.mu.Unlock()
	return nil
}


func (apiMid *ApiKeyMiddleware) getDefaultKey() string {
	apiMid.mu.Lock()
	defer apiMid.mu.Unlock()
	return apiMid.defaultKey
}
