package main

import (
	"fmt"
	"github.com/ortizdavid/go-nopain/httputils/client"
)

func main() {
	client := httputils.NewHttpClient()
	//client.SetHeader("X-API-KEY", "key123")
	
	data := "files/patterns.png"
	  
	// Example Post request
	response, err := client.Post("http://localhost:8080/upload", data)
	if err != nil {
	  fmt.Println("Error:", err)
	  return
	}
  
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Body:", string(response.Body))
}