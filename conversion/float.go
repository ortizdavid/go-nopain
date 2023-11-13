package conversion

import (
	"strconv"
)

//  float32 -> string
func Float32ToString(number float32) string {
	return strconv.FormatFloat(float64(number), 'f', -1, 32)
}

//  float32 -> int
func Float32ToInt(number float32) int {
	return int(number)
}

//  float32 -> int32
func Float32ToInt32(number float32) int32 {
	return int32(number)
}

//  float32 -> int64
func Float32ToInt64(number float32) int64 {
	return int64(number)
}

//  float64 -> string
func Float64ToString(number float64) string {
	return strconv.FormatFloat(number, 'f', -1, 64)
}

//  float64 -> int
func Float64ToInt(number float64) int {
	return int(number)
}

//  float64 -> int32
func Float64ToInt32(number float64) int32 {
	return int32(number)
}

//  float32 -> int64
func Float64ToInt64(number float64) int64 {
	return int64(number)
}
