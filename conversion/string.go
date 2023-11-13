package conversion

import (
	"strconv"
)

// string -> int
func StringToInt(strNumber string) (int, error) {
	result, err := strconv.Atoi(strNumber)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// string -> int32
func StringToInt32(strNumber string) (int32, error) {
	result, err := strconv.ParseInt(strNumber, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(result), nil
}

// string -> int64
func StringToInt64(strNumber string) (int64, error) {
	result, err := strconv.ParseInt(strNumber, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// string -> float32
func StringToFloat32(strNumber string) (float32, error) {
	result, err := strconv.ParseFloat(strNumber, 32)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}

// string -> float64
func StringToFloat64(strNumber string) (float64, error) {
	result, err := strconv.ParseFloat(strNumber, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
