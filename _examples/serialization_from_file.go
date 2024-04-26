package main

import (
	"fmt"
	"log"

	"github.com/ortizdavid/go-nopain/serialization"
)

// Define a struct to represent a CSV record
type Record struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// Example usage of FromCsvFileToJson
	var jsonData []Record
	err := serialization.FromCsvFileToJson("data.csv", &jsonData)
	if err != nil {
		log.Fatalf("Error converting CSV to JSON: %v", err)
	}
	fmt.Println("JSON Data:", jsonData)

	// Example usage of FromCsvFileToXml
	var xmlData []Record
	err = serialization.FromCsvFileToXml("data.csv", &xmlData)
	if err != nil {
		log.Fatalf("Error converting CSV to XML: %v", err)
	}
	fmt.Println("XML Data:", xmlData)
}
