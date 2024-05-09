package conversion

import (
    "testing"
)

func Test_Float32ToString(t *testing.T) {
    // Test case for positive float32
    var num1 float32 = 123.45
    expected1 := "123.45"
    if result := Float32ToString(num1); result != expected1 {
        t.Errorf("Float32ToString(%v) returned %s, expected %s", num1, result, expected1)
    }

    // Test case for negative float32
    var num2 float32 = -67.89
    expected2 := "-67.89"
    if result := Float32ToString(num2); result != expected2 {
        t.Errorf("Float32ToString(%v) returned %s, expected %s", num2, result, expected2)
    }
}

func Test_Float32ToInt(t *testing.T) {
    // Test case for positive float32
    var num1 float32 = 123.45
    expected1 := 123
    if result := Float32ToInt(num1); result != expected1 {
        t.Errorf("Float32ToInt(%v) returned %d, expected %d", num1, result, expected1)
    }

    // Test case for negative float32
    var num2 float32 = -67.89
    expected2 := -67
    if result := Float32ToInt(num2); result != expected2 {
        t.Errorf("Float32ToInt(%v) returned %d, expected %d", num2, result, expected2)
    }
}

func Test_Float32ToInt32(t *testing.T) {
    // Test case for positive float32
    var num1 float32 = 123.45
    expected1 := int32(123)
    if result := Float32ToInt32(num1); result != expected1 {
        t.Errorf("Float32ToInt32(%v) returned %d, expected %d", num1, result, expected1)
    }

    // Test case for negative float32
    var num2 float32 = -67.89
    expected2 := int32(-67)
    if result := Float32ToInt32(num2); result != expected2 {
        t.Errorf("Float32ToInt32(%v) returned %d, expected %d", num2, result, expected2)
    }
}

func Test_Float32ToInt64(t *testing.T) {
    // Test case for positive float32
    var num1 float32 = 123.45
    expected1 := int64(123)
    if result := Float32ToInt64(num1); result != expected1 {
        t.Errorf("Float32ToInt64(%v) returned %d, expected %d", num1, result, expected1)
    }

    // Test case for negative float32
    var num2 float32 = -67.89
    expected2 := int64(-67)
    if result := Float32ToInt64(num2); result != expected2 {
        t.Errorf("Float32ToInt64(%v) returned %d, expected %d", num2, result, expected2)
    }
}

func Test_Float64ToString(t *testing.T) {
    // Test case for positive float64
    var num1 float64 = 123.45
    expected1 := "123.45"
    if result := Float64ToString(num1); result != expected1 {
        t.Errorf("Float64ToString(%v) returned %s, expected %s", num1, result, expected1)
    }

    // Test case for negative float64
    var num2 float64 = -67.89
    expected2 := "-67.89"
    if result := Float64ToString(num2); result != expected2 {
        t.Errorf("Float64ToString(%v) returned %s, expected %s", num2, result, expected2)
    }
}

func Test_Float64ToInt(t *testing.T) {
    // Test case for positive float64
    var num1 float64 = 123.45
    expected1 := 123
    if result := Float64ToInt(num1); result != expected1 {
        t.Errorf("Float64ToInt(%v) returned %d, expected %d", num1, result, expected1)
    }

    // Test case for negative float64
    var num2 float64 = -67.89
    expected2 := -67
    if result := Float64ToInt(num2); result != expected2 {
        t.Errorf("Float64ToInt(%v) returned %d, expected %d", num2, result, expected2)
    }
}

func Test_Float64ToInt32(t *testing.T) {
    // Test case for positive float64
    var num1 float64 = 123.45
    expected1 := int32(123)
    if result := Float64ToInt32(num1); result != expected1 {
        t.Errorf("Float64ToInt32(%v) returned %d, expected %d", num1, result, expected1)
    }

    // Test case for negative float64
    var num2 float64 = -67.89
    expected2 := int32(-67)
    if result := Float64ToInt32(num2); result != expected2 {
        t.Errorf("Float64ToInt32(%v) returned %d, expected %d", num2, result, expected2)
    }
}

func Test_Float64ToInt64(t *testing.T) {
    // Test case for positive float64
    var num1 float64 = 123.45
    expected1 := int64(123)
    if result := Float64ToInt64(num1); result != expected1 {
        t.Errorf("Float64ToInt64(%v) returned %d, expected %d", num1, result, expected1)
    }

    // Test case for negative float64
    var num2 float64 = -67.89
    expected2 := int64(-67)
    if result := Float64ToInt64(num2); result != expected2 {
        t.Errorf("Float64ToInt64(%v) returned %d, expected %d", num2, result, expected2)
    }
}
