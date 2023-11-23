package conversion

import (
	"log"
	"strconv"
)

// string -> int
func StringToInt(strNumber string) int {
	result, err := strconv.Atoi(strNumber)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}

// string -> int32
func StringToInt32(strNumber string) int32 {
	result, err := strconv.ParseInt(strNumber, 10, 32)
	if err != nil {
		log.Fatal(err.Error())
	}
	return int32(result)
}

// string -> int64
func StringToInt64(strNumber string) int64 {
	result, err := strconv.ParseInt(strNumber, 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}

// string -> float32
func StringToFloat32(strNumber string) float32 {
	result, err := strconv.ParseFloat(strNumber, 32)
	if err != nil {
		log.Fatal(err.Error())
	}
	return float32(result)
}

// string -> float64
func StringToFloat64(strNumber string) float64 {
	result, err := strconv.ParseFloat(strNumber, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
