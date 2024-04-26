package httputils

import (
	"errors"
	"net/http"
	"sync"
)

// ApiKeyMiddleware represents middleware for API key authentication.
// It allows setting a default API key, which can be overridden by the X-API-Key header in the request.
type ApiKeyMiddleware struct {
	defaultKey string
	mu sync.RWMutex
}

// Return a ApiKeyMiddleware object, and initialize default key with the apiKey passed
func NewApiKeyMiddleWare(apiKey string) ApiKeyMiddleware {
	return ApiKeyMiddleware{
		defaultKey: apiKey,
	}
}

// Apply applies the API key middleware to a handler. It wraps the handler function.
// This method should be used with mux.Handle instead of mux.HandleFunc.
// Example usage: mux.Handle("GET /protected", middleware.Apply(protectedHandler)).
func (apiMid *ApiKeyMiddleware) Apply(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return apiMid.applyMiddleware(http.HandlerFunc(handler))
}

// applyMiddleware applies the API key middleware to a handler function.
// It retrieves the API key from the request header and validates it against the default key.
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

// SetDefaultKey sets the default API key.
// It returns an error if the provided key is empty.
func (apiMid *ApiKeyMiddleware) SetDefaultKey(value string) error {
	if value == "" {
		return errors.New("API key cannot be empty")
	}
	apiMid.mu.Lock()
	apiMid.defaultKey = value
	apiMid.mu.Unlock()
	return nil
}

// getDefaultKey returns the default API key.
func (apiMid *ApiKeyMiddleware) getDefaultKey() string {
	apiMid.mu.Lock()
	defer apiMid.mu.Unlock()
	return apiMid.defaultKey
}
