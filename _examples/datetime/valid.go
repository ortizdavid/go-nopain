package main

import (
	"fmt"

	"github.com/ortizdavid/go-nopain/datetime"
)

func main() {
	valid1 := datetime.IsValidDate("0000-00-00")
	valid2 := datetime.IsValidDate("2024-01-10")
	valid3 := datetime.IsValidDate("20000-9870-0")
	valid4 := datetime.IsValidDate("2034-06-30")
	
	fmt.Println("Validate 1 == ", valid1)
	fmt.Println("Validate 2 == ", valid2)
	fmt.Println("Validate 3 == ", valid3)
	fmt.Println("Validate 4 == ", valid4)
}