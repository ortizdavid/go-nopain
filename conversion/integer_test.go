package conversion

import (
	"testing"
)

func Test_IntToString(t *testing.T) {
	// Test case for positive int
	num1 := 123
	expected1 := "123"
	if result := IntToString(num1); result != expected1 {
		t.Errorf("IntToString(%d) returned %s, expected %s", num1, result, expected1)
	}

	// Test case for negative int
	num2 := -456
	expected2 := "-456"
	if result := IntToString(num2); result != expected2 {
		t.Errorf("IntToString(%d) returned %s, expected %s", num2, result, expected2)
	}
}

func Test_IntToFloat32(t *testing.T) {
	// Test case for positive int
	num1 := 123
	expected1 := float32(123)
	if result := IntToFloat32(num1); result != expected1 {
		t.Errorf("IntToFloat32(%d) returned %f, expected %f", num1, result, expected1)
	}

	// Test case for negative int
	num2 := -456
	expected2 := float32(-456)
	if result := IntToFloat32(num2); result != expected2 {
		t.Errorf("IntToFloat32(%d) returned %f, expected %f", num2, result, expected2)
	}
}

func Test_IntToFloat64(t *testing.T) {
	// Test case for positive int
	num1 := 123
	expected1 := float64(123)
	if result := IntToFloat64(num1); result != expected1 {
		t.Errorf("IntToFloat64(%d) returned %f, expected %f", num1, result, expected1)
	}

	// Test case for negative int
	num2 := -456
	expected2 := float64(-456)
	if result := IntToFloat64(num2); result != expected2 {
		t.Errorf("IntToFloat64(%d) returned %f, expected %f", num2, result, expected2)
	}
}

func Test_Int32ToString(t *testing.T) {
	// Test case for positive int32
	num1 := int32(123)
	expected1 := "123"
	if result := Int32ToString(num1); result != expected1 {
		t.Errorf("Int32ToString(%d) returned %s, expected %s", num1, result, expected1)
	}

	// Test case for negative int32
	num2 := int32(-456)
	expected2 := "-456"
	if result := Int32ToString(num2); result != expected2 {
		t.Errorf("Int32ToString(%d) returned %s, expected %s", num2, result, expected2)
	}
}

func Test_Int32ToFloat32(t *testing.T) {
	// Test case for positive int32
	num1 := int32(123)
	expected1 := float32(123)
	if result := Int32ToFloat32(num1); result != expected1 {
		t.Errorf("Int32ToFloat32(%d) returned %f, expected %f", num1, result, expected1)
	}

	// Test case for negative int32
	num2 := int32(-456)
	expected2 := float32(-456)
	if result := Int32ToFloat32(num2); result != expected2 {
		t.Errorf("Int32ToFloat32(%d) returned %f, expected %f", num2, result, expected2)
	}
}

func Test_Int32ToFloat64(t *testing.T) {
	// Test case for positive int32
	num1 := int32(123)
	expected1 := float64(123)
	if result := Int32ToFloat64(num1); result != expected1 {
		t.Errorf("Int32ToFloat64(%d) returned %f, expected %f", num1, result, expected1)
	}

	// Test case for negative int32
	num2 := int32(-456)
	expected2 := float64(-456)
	if result := Int32ToFloat64(num2); result != expected2 {
		t.Errorf("Int32ToFloat64(%d) returned %f, expected %f", num2, result, expected2)
	}
}

func Test_Int64ToString(t *testing.T) {
	// Test case for positive int64
	num1 := int64(123)
	expected1 := "123"
	if result := Int64ToString(num1); result != expected1 {
		t.Errorf("Int64ToString(%d) returned %s, expected %s", num1, result, expected1)
	}

	// Test case for negative int64
	num2 := int64(-456)
	expected2 := "-456"
	if result := Int64ToString(num2); result != expected2 {
		t.Errorf("Int64ToString(%d) returned %s, expected %s", num2, result, expected2)
	}
}

func Test_Int64ToFloat32(t *testing.T) {
	// Test case for positive int64
	num1 := int64(123)
	expected1 := float32(123)
	if result := Int64ToFloat32(num1); result != expected1 {
		t.Errorf("Int64ToFloat32(%d) returned %f, expected %f", num1, result, expected1)
	}

	// Test case for negative int64
	num2 := int64(-456)
	expected2 := float32(-456)
	if result := Int64ToFloat32(num2); result != expected2 {
		t.Errorf("Int64ToFloat32(%d) returned %f, expected %f", num2, result, expected2)
	}
}

func Test_Int64ToFloat64(t *testing.T) {
	// Test case for positive int64
	num1 := int64(123)
	expected1 := float64(123)
	if result := Int64ToFloat64(num1); result != expected1 {
		t.Errorf("Int64ToFloat64(%d) returned %f, expected %f", num1, result, expected1)
	}

	// Test case for negative int64
	num2 := int64(-456)
	expected2 := float64(-456)
	if result := Int64ToFloat64(num2); result != expected2 {
		t.Errorf("Int64ToFloat64(%d) returned %f, expected %f", num2, result, expected2)
	}
}
