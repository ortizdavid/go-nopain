package main

import (
	"fmt"
	"log"

	"github.com/ortizdavid/go-nopain/httputils"
)

func main() {
	client := httputils.NewHttpClient()

	url := "https://dog.ceo/dog-api/"
	
	resp, err := client.Head(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp.Body))
}