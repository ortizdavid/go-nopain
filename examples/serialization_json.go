package main

import (
	"fmt"

	"github.com/ortizdavid/go-nopain/serialization"
)

type message struct {
	Text   string 
	Number int
	Boolean bool
}


func main() {
	//jsonSerialization()
	jsonUnserialization()
}

func jsonSerialization() {
	data := message{
		Text:    "hello",
		Number:  123,
		Boolean: true,
	}

	jsonData, _ := serialization.SerializeJson(data)
	fmt.Println(string(jsonData))
}

func jsonUnserialization() {
	j := `{
		"Text": "hello",
		"Number": 123,
		"Boolean": true
	}
	`

	var msg message

	err := serialization.UnserializeJson([]byte(j), &msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)
}
