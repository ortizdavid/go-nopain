package messages

import (
	"fmt"
	"log"
)


func PrintOnError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}


func LogFailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
