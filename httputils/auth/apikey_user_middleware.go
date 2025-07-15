package httputils

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// ApiKeyUserMiddleware represents middleware for API key authentication.
// It maintains a list of user IDs and their corresponding API keys.
type ApiKeyUserMiddleware struct {
	userApikeys	map[string]string // List of user IDs and API keys
	mu			sync.RWMutex // Mutex for thread safety
}


// UserApiKey represents a user ID and its associated API key.
// This struct is used to store user credentials.
type UserApiKey struct {
	UserId	string
	ApiKey	string
}


// NewApiKeyUserMiddleware creates a new instance of ApiKeyUserMiddleware with the provided list of user keys.
func NewApiKeyUserMiddleware(userKeyList []UserApiKey) *ApiKeyUserMiddleware {
	apiKeyMap := make(map[string]string)
	for _, uk:= range userKeyList {
		apiKeyMap[uk.UserId] = uk.ApiKey
	}
	return &ApiKeyUserMiddleware{
		userApikeys: apiKeyMap,
		mu:          sync.RWMutex{},
	}
}

// Apply applies the API key middleware to a handler. It wraps the handler function.
// This method should be used with mux.Handle instead of mux.HandleFunc.
// Example usage: mux.Handle("GET /protected", middleware.Apply(protectedHandler)).
func (apiMid *ApiKeyUserMiddleware) Apply(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return apiMid.applyMiddleware(http.HandlerFunc(handler))
}


// applyMiddleware applies the API key middleware to a handler function.
// It retrieves the API key from the request header and validates it.
// The X-API-Key header must be valid and non-empty for authentication to succeed.
func (apiMid *ApiKeyUserMiddleware) applyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("X-User-Id")
		if userId == "" {
			log.Printf("Unauthorized request. User ID missing. URL: %s", r.URL.Path)
			http.Error(w, "Unauthorized. User Id missing", http.StatusUnauthorized)
			return
		}
		requestedAPIKey := r.Header.Get("X-API-KEY")
		if requestedAPIKey == "" {
			log.Printf("Unauthorized request. API Key missing. User ID: %s, URL: %s", userId, r.URL.Path)
			http.Error(w, "Unauthorized. API Key missing", http.StatusUnauthorized)
			return
		}
		apiKey, err := apiMid.GetApiKey(userId)
		if err != nil {
			log.Printf("Unauthorized request. User ID: %s, Error: %s, URL: %s", userId, err.Error(), r.URL.Path)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if requestedAPIKey != apiKey {
			log.Printf("Unauthorized request. Invalid API Key. User ID: %s, URL: %s", userId, r.URL.Path)
			http.Error(w, "Unauthorized. Invalid API Key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}


// GetApiKey retrieves the API key for a specific user ID.
// It searches the list of user API keys maintained by the middleware.
func (apiMid *ApiKeyUserMiddleware) GetApiKey(userId string) (string, error) {
	apiMid.mu.RLock()
	defer apiMid.mu.RUnlock()
	apiKey, exists := apiMid.userApikeys[userId]
	if !exists {
		return "", fmt.Errorf("key not found for user %s", userId)
	}
	return apiKey, nil
}


