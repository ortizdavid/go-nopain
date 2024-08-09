package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/ortizdavid/go-nopain/httputils"
)

// Create a new global SessionHttp instance with a custom cookie name and expiration
var session = httputils.NewSessionHttp("mySessionCookie", 30*time.Second)

func main() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	http.HandleFunc("/set-expiration", setExpiration)
	http.HandleFunc("/remove", remove)
	http.HandleFunc("/destroy", destroy)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	// Set a session value
	err := session.Set(w, "user_name", "john_doe")
	if err != nil {
		http.Error(w, "Unable to set session value", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Session value set")
}

func get(w http.ResponseWriter, r *http.Request) {
	// Get a session value
	value, err := session.Get(r, "user_name")
	if err != nil {
		http.Error(w, "Unable to get session value", http.StatusInternalServerError)
		return
	}
	if value == "" {
		fmt.Fprintln(w, "No session value found")
	} else {
		fmt.Fprintf(w, "Session value: %s\n", value)
	}
}

func setExpiration(w http.ResponseWriter, r *http.Request) {
	// Change session expiration
	session.SetExpiration(30 * time.Minute)
	fmt.Fprintln(w, "Session expiration set to 30 minutes")
}

func remove(w http.ResponseWriter, r *http.Request) {
	// Remove a session value
	err := session.Remove(w, "user_name")
	if err != nil {
		http.Error(w, "Unable to remove session value", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Session value removed")
}

func destroy(w http.ResponseWriter, r *http.Request) {
	// Destroy the session
	err := session.Destroy(r, w)
	if err != nil {
		http.Error(w, "Unable to destroy session", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Session destroyed")
}