package httputils

import "net/http"

type Session interface {
	Set(r *http.Request, w http.ResponseWriter, key string, value string) error
	Get(r *http.Request, key string) (string, error)
	Remove(r *http.Request, w http.ResponseWriter,  key string) error
	Destroy(r *http.Request, w http.ResponseWriter) error
}
