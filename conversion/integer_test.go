package conversion

import (
	"fmt"
	"testing"
)

func TestIntToFloat(t *testing.T) {
	fmt.Println(IntToFloat32(9))
	fmt.Println(IntToFloat64(9))
}

func TestIntToString(t *testing.T) {
	fmt.Println(IntToString(13))
	fmt.Println(Int32ToString(33))
	fmt.Println(Int64ToString(343))
}