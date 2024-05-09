package conversion

import "testing"

func Test_StringToInt(t *testing.T) {
    // Test case for a valid integer string
    str := "123"
    expected := 123
    result := StringToInt(str)
    if result != expected {
        t.Errorf("StringToInt(%s) returned %d, expected %d", str, result, expected)
    }

    // Test case for an invalid integer string
    str = "abc"
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("StringToInt(%s) did not panic as expected", str)
        }
    }()
    StringToInt(str)
}

func Test_StringToInt32(t *testing.T) {
    // Test case for a valid int32 string
    str := "123"
    expected := int32(123)
    result := StringToInt32(str)
    if result != expected {
        t.Errorf("StringToInt32(%s) returned %d, expected %d", str, result, expected)
    }

    // Test case for an invalid int32 string
    str = "abc"
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("StringToInt32(%s) did not panic as expected", str)
        }
    }()
    StringToInt32(str)
}

func Test_StringToInt64(t *testing.T) {
    // Test case for a valid int64 string
    str := "123"
    expected := int64(123)
    result := StringToInt64(str)
    if result != expected {
        t.Errorf("StringToInt64(%s) returned %d, expected %d", str, result, expected)
    }

    // Test case for an invalid int64 string
    str = "abc"
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("StringToInt64(%s) did not panic as expected", str)
        }
    }()
    StringToInt64(str)
}

func Test_StringToFloat32(t *testing.T) {
    // Test case for a valid float32 string
    str := "123.45"
    expected := float32(123.45)
    result := StringToFloat32(str)
    if result != expected {
        t.Errorf("StringToFloat32(%s) returned %f, expected %f", str, result, expected)
    }

    // Test case for an invalid float32 string
    str = "abc"
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("StringToFloat32(%s) did not panic as expected", str)
        }
    }()
    StringToFloat32(str)
}

func Test_StringToFloat64(t *testing.T) {
    // Test case for a valid float64 string
    str := "123.45"
    expected := 123.45
    result := StringToFloat64(str)
    if result != expected {
        t.Errorf("StringToFloat64(%s) returned %f, expected %f", str, result, expected)
    }

    // Test case for an invalid float64 string
    str = "abc"
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("StringToFloat64(%s) did not panic as expected", str)
        }
    }()
    StringToFloat64(str)
}
