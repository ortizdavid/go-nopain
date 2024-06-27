package main

import (
	"fmt"
	"github.com/ortizdavid/go-nopain/httputils"
)

func main() {
	client := httputils.NewHttpClient()

	data :=  `
	{
		"customerType": 1,
		"customerName": "Kelson Firmino",
		"identificationNumber": "Q3D95FS672HOA",
		"gender": "Female",
		"birthDate": "1997-08-01",
		"email": "eliandra@gmail.com",
		"phone": "+2493698765",
		"address": "Luanda, Viana, Zango"
	}`
	  
	// Example Post request
	response, err := client.Post("http://localhost:5062/api/customers", data)
	if err != nil {
	  fmt.Println("Error:", err)
	  return
	}
  
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Body:", string(response.Body))
}