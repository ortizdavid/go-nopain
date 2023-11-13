package conversion

import (
	"testing"
	"fmt"
)

func TestFloatToInt(t *testing.T) {
	fmt.Println(Float32ToInt(10.9))
	fmt.Println(Float32ToInt32(120.9))
	fmt.Println(Float32ToInt64(10.9))	
}

func TestFloatToString(t *testing.T) {
	fmt.Println(Float32ToString(10.9))
	fmt.Println(Float64ToString(10.9))
}