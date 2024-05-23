package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/ortizdavid/go-nopain/httputils"
	"github.com/ortizdavid/go-nopain/random"
)



func main() {
	// Create a new gorilla session store
	store := httputils.NewGorillaStore("secret-key")
	session := httputils.NewSessionGorilla(store, "session-name", 10 * time.Second)

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		err := session.Set(r, w, "username", "victor")
		session.Set(r, w, "password", "victor")
		session.Set(r, w, "company", "clinica 12")
		session.Set(r, w, "session_id", random.String(20))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Session set successfully")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		username, _ := session.Get(r, "username")
		password, _ := session.Get(r, "password")
		empresa, _ := session.Get(r, "company")
		sessionId, _ := session.Get(r, "session_id")
		/*if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}*/
		fmt.Fprintln(w, "Username:", username)
		fmt.Fprintln(w, "Password:", password)
		fmt.Fprintln(w, "Empresa:", empresa)
		fmt.Fprintln(w, "Session Id:", sessionId)
	})

	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		err := session.Remove(r, w, "username")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Username removed from session")
	})

	http.HandleFunc("/destroy", func(w http.ResponseWriter, r *http.Request) {
		err := session.Destroy(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Session destroyed")
	})

	http.ListenAndServe(":8080", nil)
}