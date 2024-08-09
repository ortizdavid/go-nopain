package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-nopain/random"
)

// Create a new gorilla session store
var store = httputils.NewGorillaStore("secret-key")
var session = httputils.NewSessionGorilla(store, "session-name", 10 * time.Second)

func main() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	http.HandleFunc("/remove", remove)
	http.HandleFunc("/destroy", destroy)
	fmt.Println("Starting server on :8081")
	http.ListenAndServe(":8081", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	err := session.Set(r, w, "username", "Anna")
	session.Set(r, w, "password", "Anna123$")
	session.Set(r, w, "company", "company .Corp")
	session.Set(r, w, "session_id", random.String(20))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Session set successfully")
}

func get(w http.ResponseWriter, r *http.Request) {
	username, _ := session.Get(r, "user_name")
	password, _ := session.Get(r, "password")
	empresa, _ := session.Get(r, "company")
	sessionId, _ := session.Get(r, "session_id")
	fmt.Fprintln(w, "Username:", username)
	fmt.Fprintln(w, "Password:", password)
	fmt.Fprintln(w, "Company:", empresa)
	fmt.Fprintln(w, "Session Id:", sessionId)
}

func remove(w http.ResponseWriter, r *http.Request) {
	err := session.Remove(r, w, "user_name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Username removed from session")
}

func destroy(w http.ResponseWriter, r *http.Request) {
	err := session.Destroy(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Session destroyed")
}