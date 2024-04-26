package main

import (
	"net/http"
	"github.com/ortizdavid/go-nopain/httputils"
)

func greet(w http.ResponseWriter, r *http.Request) {
	httputils.WriteXml(w, 200, 
	struct{
		Text string `json:"text"`
		Number int `json:"number"`
		Boolean bool `json:"boolean"`
	}{
		Text: "hello",
		Number: 123,
		Boolean: true,
	}, 1)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}