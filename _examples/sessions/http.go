package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ortizdavid/go-nopain/httputils"
)

func main() {
	// Create a new SessionHttp instance with a custom cookie name and expiration
	session := httputils.NewSessionHttp("mySessionCookie", 30*time.Second)

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		// Set a session value
		err := session.Set(w, "username", "john_doe")
		if err != nil {
			http.Error(w, "Unable to set session value", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Session value set")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		// Get a session value
		value, err := session.Get(r, "username")
		if err != nil {
			http.Error(w, "Unable to get session value", http.StatusInternalServerError)
			return
		}
		if value == "" {
			fmt.Fprintln(w, "No session value found")
		} else {
			fmt.Fprintf(w, "Session value: %s\n", value)
		}
	})

	http.HandleFunc("/set-expiration", func(w http.ResponseWriter, r *http.Request) {
		// Change session expiration
		session.SetExpiration(30 * time.Minute)
		fmt.Fprintln(w, "Session expiration set to 30 minutes")
	})

	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		// Remove a session value
		err := session.Remove(w, "username")
		if err != nil {
			http.Error(w, "Unable to remove session value", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Session value removed")
	})

	http.HandleFunc("/destroy", func(w http.ResponseWriter, r *http.Request) {
		// Destroy the session
		err := session.Destroy(r, w)
		if err != nil {
			http.Error(w, "Unable to destroy session", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Session destroyed")
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
