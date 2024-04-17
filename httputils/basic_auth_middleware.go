package httputils

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// BasicAuthMiddleware provides middleware for basic authentication
type BasicAuthMiddleware struct {
	users []UserBasicAuth
}

// UserBasicAuth represents a user with username and password
type UserBasicAuth struct {
	Username string
	Password string
}

// NewBasicAuthMiddleware returns a new BasicAuthMiddleware object with the provided username and password.
func NewBasicAuthMiddleware(allowedUsers []UserBasicAuth) *BasicAuthMiddleware {
    return &BasicAuthMiddleware{
    	users: allowedUsers,
    }
}

// NewBasicAuthMiddleware returns a new BasicAuthMiddleware object with the provided username and password.
func (ba *BasicAuthMiddleware) Apply(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return ba.applyMiddleware(http.HandlerFunc(handler))
}

// applyMiddleware applies the basic authentication middleware.
func (ba *BasicAuthMiddleware) applyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
        authHeader := r.Header.Get("Authorization")
        
        // Check if the Authorization header is set
        if authHeader == "" {
            // No Authorization header provided, request authentication
            w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, "Unauthorized access\n")
            return
        }
        
        // Check if the Authorization header starts with "Basic"
        if !strings.HasPrefix(authHeader, "Basic ") {
            // Invalid Authorization header
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprint(w, "Invalid Authorization header\n")
            return
        }
        
        // Decode the base64-encoded credentials
        credentials, err := base64.StdEncoding.DecodeString(authHeader[len("Basic "):])
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, "Error decoding credentials: %v\n", err)
            return
        }
        
        // Split the credentials into username and password
        parts := strings.SplitN(string(credentials), ":", 2)
        if len(parts) != 2 {
            // Malformed credentials
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprint(w, "Malformed credentials\n")
            return
        }

        // Check if the provided username and password are valid
        username, password := parts[0], parts[1]
        if !ba.isValidUser(username, password) {
            // Incorrect username or password
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, "Unauthorized access\n")
            return
        }
        
		next.ServeHTTP(w, r)
	})
}

// isValidUser checks if the provided username and password are valid
func (ba *BasicAuthMiddleware) isValidUser(username, password string) bool {
    for _, user := range ba.users {
        if user.Username == username && user.Password == password {
            return true
        }
    }
    return false
}