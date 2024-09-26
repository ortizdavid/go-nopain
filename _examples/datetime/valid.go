package main

import (
	"fmt"
	"github.com/ortizdavid/go-nopain/datetime"
)

func main() {
	// Validations
	// date
	fmt.Println("Validating date")
	fmt.Println("1 == ", datetime.IsValidDate("0000-00-00"))
	fmt.Println("2 == ", datetime.IsValidDate("2024-09-270"))
	
	// date time 
	fmt.Println("Validating date time")
	fmt.Println("1 == ", datetime.IsValidDateTime("20000-008-10 10:00:2"))
	fmt.Println("2 == ", datetime.IsValidDateTime("2024-09-27 12:54:09"))

	// time 
	fmt.Println("Validating time")
	fmt.Println("1 == ", datetime.IsValidTime("1:00:802"))
	fmt.Println("2 == ", datetime.IsValidTime("12:54:09"))
}