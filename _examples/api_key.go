package main

import (
	"fmt"
	"net/http"

	"github.com/ortizdavid/go-nopain/httputils"
)

func main() {

	mux := http.NewServeMux()
	middleware := httputils.NewApiKeyMiddleWare("key123")

	mux.HandleFunc("GET /", indexHandler)
	mux.HandleFunc("GET /public", publicHandler)
	mux.Handle("GET /protected-get", middleware.Apply(protectedHandler))
	mux.Handle("DELETE /protected-delete", middleware.Apply(protectedHandler2))

	fmt.Println("Listen at http://127.0.0.1:7000")
	http.ListenAndServe(":7000", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index")
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Public Content")
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Protected GET Content")
}

func protectedHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Protected DELETE Content")
}


