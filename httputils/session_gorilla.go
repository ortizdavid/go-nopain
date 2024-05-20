package httputils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/ortizdavid/go-nopain/encryption"
	"github.com/ortizdavid/go-nopain/random"
)

// SessionGorilla implements the Session interface using gorilla/sessions
type SessionGorilla struct {
	store sessions.Store
	name string
	secretKey string
	expiration time.Duration //expiration time
}

// NewGorillaStore creates a new gorilla session store with a secret key
func NewGorillaStore(secretKey string) sessions.Store {
	return sessions.NewCookieStore([]byte(secretKey))
}

// NewSessionGorilla creates a new SessionGorilla
func NewSessionGorilla(store sessions.Store, name string, expiration time.Duration) *SessionGorilla {
	return &SessionGorilla{
		store: store,
		name:  name,
		expiration: expiration,
	}
}

func NewSessionGorillaDefault() *SessionGorilla {
	secretKey := encryption.GenerateUUID()
	store := sessions.NewCookieStore([]byte(secretKey))
	return &SessionGorilla{
		store: store,
		name:  random.String(15),
		secretKey: secretKey,
		expiration: 15 * time.Minute,
	}
}

// Set sets a key-value pair in the session
func (s *SessionGorilla) Set(r *http.Request, w http.ResponseWriter, key string, value string) error {
	session, err := s.store.Get(r, s.name)
	if err != nil {
		return err
	}
	session.Values[key] = value
	session.Options.MaxAge = int(s.expiration.Seconds())
	return session.Save(r, w)
}

// Get retrieves a value from the session by key
func (s *SessionGorilla) Get(r *http.Request, key string) (string, error) {
	session, err := s.store.Get(r, s.name)
	if err != nil {
		return "", err
	}
	value, ok := session.Values[key].(string)
	if !ok {
		return "", fmt.Errorf("key not found or not a string")
	}
	return value, nil
}

// Remove removes a key-value pair from the session
func (s *SessionGorilla) Remove(r *http.Request, w http.ResponseWriter, key string) error {
	session, err := s.store.Get(r, s.name)
	if err != nil {
		return err
	}
	delete(session.Values, key)
	return session.Save(r, w)
}

// Destroy invalidates the entire session
func (s *SessionGorilla) Destroy(r *http.Request, w http.ResponseWriter) error {
	session, err := s.store.Get(r, s.name)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1 // Setting MaxAge to -1 deletes the cookie
	return session.Save(r, w)
}

func (s * SessionGorilla) GetSecretKey() string {
	return s.secretKey
}

// SetExpiration sets the expiration duration for sessions.
func (s *SessionGorilla) SetExpiration(expiration time.Duration)  {
	s.expiration = expiration
}