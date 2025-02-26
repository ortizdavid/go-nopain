package main

import (
	"fmt"

	"github.com/ortizdavid/go-nopain/serialization"
)

func main() {

	xml := `<person>
		<name>Andre</name>
		<age>23</age>
	</person>`


	jsonData, _ := serialization.XmlToJson([]byte(xml))
	fmt.Println(string(jsonData))
}