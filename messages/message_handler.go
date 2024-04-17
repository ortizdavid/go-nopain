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

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", err, msg)
	}
}
