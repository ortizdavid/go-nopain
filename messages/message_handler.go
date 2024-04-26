package messages

import (
	"fmt"
	"log"
)

// PrintOnError prints the error message if it's not nil.
func PrintOnError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// LogFailOnError logs the error message and terminates the program if it's not nil.
func LogFailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// PanicOnError panics with the error message if it's not nil.
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// FailOnError logs the error message along with a custom message and terminates the program if the error is not nil.
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", err, msg)
	}
}
