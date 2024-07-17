package httputils

import (
	"log"
	"net/http"
)

func logRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}

func UseLoggingMiddleware(mux *http.ServeMux) {
	logRequests(mux)
}