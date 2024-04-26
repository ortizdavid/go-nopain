package main

import (
	"fmt"
	"net/http"
	"github.com/ortizdavid/go-nopain/httputils"
)


// Define a list of users
var users = []httputils.UserBasicAuth{
    {"user1", "pass1"},
    {"user2", "pass2"},
}


func main() {

	mux := http.NewServeMux()
    // Define your handler function 

	middleware := httputils.NewBasicAuthMiddleware(users)
    
    // Wrap your handler function with BasicAuth middleware
    mux.Handle("GET /", middleware.Apply(helloHandler))
    
    // Start the HTTP server
    http.ListenAndServe(":7000", mux)
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, authenticated user!")
}