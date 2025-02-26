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
	// Example usage of FromJsonFile
	/*var jsonData []Record
	err := serialization.FromJsonFile("files/data.json", &jsonData)
	if err != nil {
		log.Fatalf("Error converting form JSON file: %v", err)
	}
	fmt.Println("JSON Data:", jsonData)*/

	// Example usage of FromXmlFile
	var xmlData []Record
	err := serialization.FromXmlFile("files/data.xml", &xmlData)
	if err != nil {
		log.Fatalf("Error converting from XML file: %v", err)
	}
	fmt.Println("XML Data:", xmlData)
}
