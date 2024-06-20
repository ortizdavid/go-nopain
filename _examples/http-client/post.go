package main

import (
	"fmt"
	"github.com/ortizdavid/go-nopain/httputils"
)

func main() {
	client := httputils.NewHttpClient()
	//client.SetHeader("X-API-KEY", "key123")
	
	data :=  `
	{
		"UserId": 123,
		"TaskName": "Other",
		"StartDate": "2024-01-01",
		"EndDate": "20024-01-01"
	}`
	  
	// Example GET request
	response, err := client.Post("http://localhost:8000/tasks", data)
	if err != nil {
	  fmt.Println("Error:", err)
	  return
	}
  
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Body:", string(response.Body))
}