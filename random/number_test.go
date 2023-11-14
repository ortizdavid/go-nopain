package random

import (
	"fmt"
	"testing"
)

func TestRandomInt(t *testing.T) {
	fmt.Println(Int(1, 100))
}

func TestRandomInt32(t *testing.T) {
	fmt.Println(Int32(1, 100))
}

func TestRandomInt64(t *testing.T) {
	fmt.Println(Int64(1, 100))
}

func TestRandomFloat32(t *testing.T) {
	fmt.Println(Float32(1, 100))
}

func TestRandomFloat64(t *testing.T) {
	fmt.Println(Float64(1, 100))
}


