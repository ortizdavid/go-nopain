package main

import (
	"fmt"
	"log"
	"github.com/ortizdavid/go-nopain/httputils/client"
)

func main() {
	client := httputils.NewHttpClient()

	url := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp.Body))
}