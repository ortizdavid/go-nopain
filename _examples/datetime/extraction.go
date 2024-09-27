package main

import (
	"fmt"
	"time"

	"github.com/ortizdavid/go-nopain/datetime"
)

func main() {

	date := time.Now()
	fmt.Println("Date and Time Extraction")
	fmt.Println("Original: ", date)
	fmt.Println("Date: ", datetime.ExtractDate(date))
	fmt.Println("Date Time: ", datetime.ExtractDateTime(date))
	fmt.Println("Year: ", datetime.ExtractYear(date))
	fmt.Println("Week: ", datetime.ExtractWeek(date))
	fmt.Println("Day: ", datetime.ExtractDay(date))
	fmt.Println("Month: ", datetime.ExtractMonth(date))
	fmt.Println("Hour: ", datetime.ExtractHour(date))
	fmt.Println("Seconds: ", datetime.ExtractSeconds(date))
	fmt.Println("Milliseconds: ", datetime.ExtractMilliseconds(date))
}