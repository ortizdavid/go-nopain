package conversion

import (
	"testing"
	"fmt"
)

func TestFloat32ToInt(t *testing.T) {
	fmt.Println(Float32ToInt(10.9))
	fmt.Println(Float32ToInt32(120.9))
	fmt.Println(Float32ToInt64(10.9))	
}

func TestFloat64ToInt32(t *testing.T) {
	fmt.Println(Float64ToInt32(10.9))
	fmt.Println(Float64ToInt32(120.9))
	fmt.Println(Float64ToInt32(10.9))	
}

func TestFloat64ToInt64(t *testing.T) {
	fmt.Println(Float64ToInt64(10.9))
	fmt.Println(Float64ToInt64(120.9))
	fmt.Println(Float64ToInt64(10.9))	
}

func TestFloatToString(t *testing.T) {
	fmt.Println(Float32ToString(10.9))
	fmt.Println(Float64ToString(10.9))
}