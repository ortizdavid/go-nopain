package httputils

import (
	"log"
	"net/http"
)

func LogHttpRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}