package main

import (
	"fmt"
	"github.com/ortizdavid/go-nopain/datetime"
)

func main() {
	fmt.Println("Some Operations\n")
	fmt.Println("--------------------------")
	fmt.Println("Add Days: ", datetime.AddDaysStr("1994-01-01", 12))
	fmt.Println("Add Weeks: ", datetime.AddWeeksStr("1994-01-01", 3))
	fmt.Println("Add Years: ", datetime.AddYearsStr("1994-01-01", 20))
	fmt.Println("Subtract Years: ", datetime.SubtractYearsStr("1994-01-01", 5))
	fmt.Println("--------------------------")
	fmt.Println("Age: ", datetime.GetAgeStr("1994-01-01"))
	fmt.Println("--------------------------")
	fmt.Println("Time Ago: ", datetime.TimeAgoStr("2014-01-19 16:20:16"))
	fmt.Println("Time Ago Betwen: ", datetime.TimeAgoBetweenStr("2019-01-19 16:20:16", "2023-12-31 23:59:00"))
	fmt.Println()
}