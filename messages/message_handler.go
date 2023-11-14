package messages

import (
	"fmt"
	"log"
)


func PrintFailOrError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func LogFailOrError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


