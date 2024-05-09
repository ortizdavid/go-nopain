package main

import (
	"fmt"
)

// FormatCurrency formats the given amount as currency with the specified symbol.
// It takes a float64 amount and a string symbol as input and returns the formatted currency string.
func FormatCurrency(amount float64, symbol string) string {
	// Format the amount with two decimal places and append the currency symbol.
	return fmt.Sprintf("%.2f %s", amount, symbol)
}

func main() {
	// Define the amount and symbol
	amount := 1234.56
	symbol := "$"

	// Format the currency using the FormatCurrency function
	formatted := FormatCurrency(amount, symbol)

	// Print the formatted currency
	fmt.Println("Formatted currency:", formatted)
}
