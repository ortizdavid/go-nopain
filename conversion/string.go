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

// string -> int or nil
func StringToIntOrNil(strNumber string) *int {
	if strNumber == "" {
		return nil
	}
	numInt, err := strconv.Atoi(strNumber)
	if err != nil {
		return nil
	}
	return &numInt
}

// string -> int32 or nil
func StringToInt32OrNil(strNumber string) *int32 {
	if strNumber == "" {
		return nil
	}
	numInt, err := strconv.ParseInt(strNumber, 10, 32)
	if err != nil {
		return nil
	}
	result := int32(numInt)
	return &result
}

// string -> int64 or nil
func StringToInt64OrNil(strNumber string) *int64 {
	if strNumber == "" {
		return nil
	}
	numInt, err := strconv.ParseInt(strNumber, 10, 64)
	if err != nil {
		return nil
	}
	result := int64(numInt)
	return &result
}

// string -> float32 or nil
func StringToFloat32OrNil(strNumber string) *float32 {
	if strNumber == "" {
		return nil
	}
	numFloat, err := strconv.ParseFloat(strNumber, 32)
	if err != nil {
		return nil
	}
	result := float32(numFloat)
	return &result
}

// string -> float64 or nil
func StringToFloat64OrNil(strNumber string) *float64 {
	if strNumber == "" {
		return nil
	}
	numFloat, err := strconv.ParseFloat(strNumber, 64)
	if err != nil {
		return nil
	}
	result := float64(numFloat)
	return &result
}

// string -> String or Nil
func StringOrNil(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}
