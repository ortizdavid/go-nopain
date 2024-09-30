package conversion

import (
	"fmt"
	"strconv"
)

// StringToInt converts a string to an int and returns an error if conversion fails.
func StringToInt(strNumber string) (int, error) {
	result, err := strconv.Atoi(strNumber)
	if err != nil {
		return 0, fmt.Errorf("error converting string to int '%s': %w", strNumber, err)
	}
	return result, nil
}

// StringToInt32 converts a string to an int32 and returns an error if conversion fails.
func StringToInt32(strNumber string) (int32, error) {
	result, err := strconv.ParseInt(strNumber, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("error converting string to int32 '%s': %w", strNumber, err)
	}
	return int32(result), nil
}

// StringToInt64 converts a string to an int64 and returns an error if conversion fails.
func StringToInt64(strNumber string) (int64, error) {
	result, err := strconv.ParseInt(strNumber, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting string to int64 '%s': %w", strNumber, err)
	}
	return result, nil
}

// StringToInt8 converts a string to an int8 and returns an error if conversion fails.
func StringToInt8(strNumber string) (int8, error) {
	result, err := strconv.ParseInt(strNumber, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("error converting string to int8 '%s': %w", strNumber, err)
	}
	// Check if the result is within the valid range for int8
	if result < -128 || result > 127 {
		return 0, fmt.Errorf("value out of range for int8: %d", result)
	}
	return int8(result), nil
}

// StringToFloat32 converts a string to a float32 and returns an error if conversion fails.
func StringToFloat32(strNumber string) (float32, error) {
	result, err := strconv.ParseFloat(strNumber, 32)
	if err != nil {
		return 0, fmt.Errorf("error converting string to float32 '%s': %w", strNumber, err)
	}
	return float32(result), nil
}

// StringToFloat64 converts a string to a float64 and returns an error if conversion fails.
func StringToFloat64(strNumber string) (float64, error) {
	result, err := strconv.ParseFloat(strNumber, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting string to float64 '%s': %w", strNumber, err)
	}
	return result, nil
}
