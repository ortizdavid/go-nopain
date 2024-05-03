package random

import (
	"testing"
)

func TestInt(t *testing.T) {
	min, max := 0, 10
	result := Int(min, max)
	if result < min || result >= max {
		t.Errorf("Int(%d, %d) returned %d, which is out of range", min, max, result)
	}
}

func TestInt32(t *testing.T) {
	min, max := int32(0), int32(10)
	result := Int32(min, max)
	if result < min || result >= max {
		t.Errorf("Int32(%d, %d) returned %d, which is out of range", min, max, result)
	}
}

func TestInt64(t *testing.T) {
	min, max := int64(0), int64(10)
	result := Int64(min, max)
	if result < min || result >= max {
		t.Errorf("Int64(%d, %d) returned %d, which is out of range", min, max, result)
	}
}

func TestFloat32(t *testing.T) {
	min, max := float32(0), float32(10)
	result := Float32(min, max)
	if result < min || result >= max {
		t.Errorf("Float32(%f, %f) returned %f, which is out of range", min, max, result)
	}
}

func TestFloat64(t *testing.T) {
	min, max := float64(0), float64(10)
	result := Float64(min, max)
	if result < min || result >= max {
		t.Errorf("Float64(%f, %f) returned %f, which is out of range", min, max, result)
	}
}
