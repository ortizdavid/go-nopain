package conversion

import "strconv"

// int -> string
func IntToString(number int) string {
	return strconv.Itoa(number)
}

// int -> float32
func IntToFloat32(number int) float32 {
	return float32(number)
}

// int -> float64
func IntToFloat64(number int) float64 {
	return float64(number)
}

// int32 -> string
func Int32ToString(number int32) string {
	return strconv.FormatInt(int64(number), 10)
}

// int32 -> float32
func Int32ToFloat32(number int32) float32 {
	return float32(number)
}

// int32 -> float64
func Int32ToFloat64(number int32) float64 {
	return float64(number)
}

// int64 -> string
func Int64ToString(number int64) string {
	return strconv.FormatInt(number, 10)
}

// int64 -> float32
func Int64ToFloat32(number int64) float32 {
	return float32(number)
}

// int64 -> float64
func Int64ToFloat64(number int64) float64 {
	return float64(number)
}
