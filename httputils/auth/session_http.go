package httputils

import (
	"net/http"
	"time"
)

// SessionHttp stores session cookie details.
type SessionHttp struct {
	cookieName string        // Name of the session cookie.
	expiration time.Duration // Expiration time of the session.
}

// NewSessionHttp creates a new session with a custom cookie name and expiration.
func NewSessionHttp(cookieName string, expiration time.Duration) *SessionHttp {
	return &SessionHttp{
		cookieName: cookieName,
		expiration: expiration,
	}
}

// NewSessionHttpDefault creates a new session with default settings.
func NewSessionHttpDefault() *SessionHttp {
	return &SessionHttp{
		cookieName: "default-session",
		expiration: 15 * time.Minute,
	}
}

// SetExpiration updates the session expiration time.
func (s *SessionHttp) SetExpiration(expiration time.Duration) {
	s.expiration = expiration
}

// Set sets a key-value pair in the session.
func (s *SessionHttp) Set(w http.ResponseWriter, key string, value string) error {
	cookie := &http.Cookie{
		Name: key,
		Value: value,
		Expires: time.Now().Add(s.expiration),
		Path: "/",
	}
	http.SetCookie(w, cookie)
	return nil
}

// Get retrieves a value from the session by key.
func (s *SessionHttp) Get(r *http.Request, key string) (string, error) {
	cookie, err := r.Cookie(key)
	if err != nil {
		if err != http.ErrNoCookie {
			return "", nil // No cookie found, return empty string without error
		}
		return "", err // Other errors, return the error
	}
	return cookie.Value, nil
}

// Remove Session key
func (s *SessionHttp) Remove(w http.ResponseWriter,  key string) error {
	cookie := &http.Cookie{
		Name: key,
		Value: "",
		Expires: time.Unix(0, 0),
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	return nil
}

// Destroy current session
func (s *SessionHttp) Destroy(r *http.Request, w http.ResponseWriter) error {
	for _, cookie := range r.Cookies() {
		c := &http.Cookie{
			Name: cookie.Name,
			Value: "",
			Expires: time.Unix(0, 0),
			Path: "/",
			MaxAge: -1,
		} 
		http.SetCookie(w, c)
	}
	return nil
}
