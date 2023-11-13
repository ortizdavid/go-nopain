package conversion

import (
	"fmt"
	"testing"
)

func TestIntToFloat(t *testing.T) {
	fmt.Println(IntToFloat32(9))
	fmt.Println(IntToFloat64(9))
}

func TestInt32ToFloat(t *testing.T) {
	fmt.Println(Int32ToFloat32(92))
	fmt.Println(Int32ToFloat64(39))
}

func TestInt64ToFloat(t *testing.T) {
	fmt.Println(Int64ToFloat32(922))
	fmt.Println(Int64ToFloat64(922334))
}

func TestIntToString(t *testing.T) {
	fmt.Println(IntToString(13))
	fmt.Println(Int32ToString(3))
	fmt.Println(Int64ToString(12345678))
}