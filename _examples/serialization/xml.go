package main

import (
	"fmt"
	"encoding/xml"
	"github.com/ortizdavid/go-nopain/serialization"
)


type person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {

	//jsonSerialization()
	//xmlSerialization()
	xmlUnserialization()
}

func xmlSerialization() {
	data2 := person{
		XMLName: xml.Name{},
		Name:    "Andre",
		Age:     23,
	}

	xmlData, _ := serialization.SerializeXml(data2)
	fmt.Println(string(xmlData))
}

func xmlUnserialization()  {
	x := `<person>
		<name>Andre</name>
		<age>23</age>
	</person>
	`
	var p person
	serialization.UnserializeXml([]byte(x), &p)
	fmt.Println(p)
}