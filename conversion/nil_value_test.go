package conversion

import (
	"testing"
)

func TestStringToIntOrNil(t *testing.T) {
	// Test cases where string is convertible to int
	validStr := "123"
	if val := StringToIntOrNil(validStr); val == nil || *val != 123 {
		t.Errorf("StringToIntOrNil(%s) returned unexpected result: %v", validStr, val)
	}

	// Test case where string is empty
	emptyStr := ""
	if val := StringToIntOrNil(emptyStr); val != nil {
		t.Errorf("StringToIntOrNil(%s) returned unexpected result: %v", emptyStr, val)
	}

	// Test case where string is not convertible to int
	invalidStr := "abc"
	if val := StringToIntOrNil(invalidStr); val != nil {
		t.Errorf("StringToIntOrNil(%s) returned unexpected result: %v", invalidStr, val)
	}
}

func TestStringToInt32OrNil(t *testing.T) {
	// Test cases where string is convertible to int32
	validStr := "123"
	if val := StringToInt32OrNil(validStr); val == nil || *val != 123 {
		t.Errorf("StringToInt32OrNil(%s) returned unexpected result: %v", validStr, val)
	}

	// Test case where string is empty
	emptyStr := ""
	if val := StringToInt32OrNil(emptyStr); val != nil {
		t.Errorf("StringToInt32OrNil(%s) returned unexpected result: %v", emptyStr, val)
	}

	// Test case where string is not convertible to int32
	invalidStr := "abc"
	if val := StringToInt32OrNil(invalidStr); val != nil {
		t.Errorf("StringToInt32OrNil(%s) returned unexpected result: %v", invalidStr, val)
	}
}

func TestStringToInt64OrNil(t *testing.T) {
	// Test cases where string is convertible to int64
	validStr := "123"
	if val := StringToInt64OrNil(validStr); val == nil || *val != 123 {
		t.Errorf("StringToInt64OrNil(%s) returned unexpected result: %v", validStr, val)
	}

	// Test case where string is empty
	emptyStr := ""
	if val := StringToInt64OrNil(emptyStr); val != nil {
		t.Errorf("StringToInt64OrNil(%s) returned unexpected result: %v", emptyStr, val)
	}

	// Test case where string is not convertible to int64
	invalidStr := "abc"
	if val := StringToInt64OrNil(invalidStr); val != nil {
		t.Errorf("StringToInt64OrNil(%s) returned unexpected result: %v", invalidStr, val)
	}
}

func TestStringToFloat32OrNil(t *testing.T) {
	// Test cases where string is convertible to float32
	validStr := "123.45"
	if val := StringToFloat32OrNil(validStr); val == nil || *val != 123.45 {
		t.Errorf("StringToFloat32OrNil(%s) returned unexpected result: %v", validStr, val)
	}

	// Test case where string is empty
	emptyStr := ""
	if val := StringToFloat32OrNil(emptyStr); val != nil {
		t.Errorf("StringToFloat32OrNil(%s) returned unexpected result: %v", emptyStr, val)
	}

	// Test case where string is not convertible to float32
	invalidStr := "abc"
	if val := StringToFloat32OrNil(invalidStr); val != nil {
		t.Errorf("StringToFloat32OrNil(%s) returned unexpected result: %v", invalidStr, val)
	}
}

func TestStringToFloat64OrNil(t *testing.T) {
	// Test cases where string is convertible to float64
	validStr := "123.45"
	if val := StringToFloat64OrNil(validStr); val == nil || *val != 123.45 {
		t.Errorf("StringToFloat64OrNil(%s) returned unexpected result: %v", validStr, val)
	}

	// Test case where string is empty
	emptyStr := ""
	if val := StringToFloat64OrNil(emptyStr); val != nil {
		t.Errorf("StringToFloat64OrNil(%s) returned unexpected result: %v", emptyStr, val)
	}

	// Test case where string is not convertible to float64
	invalidStr := "abc"
	if val := StringToFloat64OrNil(invalidStr); val != nil {
		t.Errorf("StringToFloat64OrNil(%s) returned unexpected result: %v", invalidStr, val)
	}
}

func TestStringOrNil(t *testing.T) {
	// Test cases where string is not empty
	str := "test"
	if val := StringOrNil(str); val == nil || *val != str {
		t.Errorf("StringOrNil(%s) returned unexpected result: %v", str, val)
	}

	// Test case where string is empty
	emptyStr := ""
	if val := StringOrNil(emptyStr); val != nil {
		t.Errorf("StringOrNil(%s) returned unexpected result: %v", emptyStr, val)
	}
}

func TestIntOrNilToString(t *testing.T) {
	// Test cases where int is not nil
	num := 123
	if val := IntOrNilToString(&num); val != "123" {
		t.Errorf("IntOrNilToString(%v) returned unexpected result: %s", num, val)
	}

	// Test case where int is nil
	var nilNum *int
	if val := IntOrNilToString(nilNum); val != "" {
		t.Errorf("IntOrNilToString(%v) returned unexpected result: %s", nilNum, val)
	}
}

func TestInt32OrNilToString(t *testing.T) {
	// Test cases where int32 is not nil
	num := int32(123)
	if val := Int32OrNilToString(&num); val != "123" {
		t.Errorf("Int32OrNilToString(%v) returned unexpected result: %s", num, val)
	}

	// Test case where int32 is nil
	var nilNum *int32
	if val := Int32OrNilToString(nilNum); val != "" {
		t.Errorf("Int32OrNilToString(%v) returned unexpected result: %s", nilNum, val)
	}
}

func TestInt64OrNilToString(t *testing.T) {
	// Test cases where int64 is not nil
	num := int64(123)
	if val := Int64OrNilToString(&num); val != "123" {
		t.Errorf("Int64OrNilToString(%v) returned unexpected result: %s", num, val)
	}

	// Test case where int64 is nil
	var nilNum *int64
	if val := Int64OrNilToString(nilNum); val != "" {
		t.Errorf("Int64OrNilToString(%v) returned unexpected result: %s", nilNum, val)
	}
}

func TestFloat32OrNilToString(t *testing.T) {
	// Test cases where float32 is not nil
	num := float32(123.45)
	if val := Float32OrNilToString(&num); val != "123.45" {
		t.Errorf("Float32OrNilToString(%v) returned unexpected result: %s", num, val)
	}

	// Test case where float32 is nil
	var nilNum *float32
	if val := Float32OrNilToString(nilNum); val != "" {
		t.Errorf("Float32OrNilToString(%v) returned unexpected result: %s", nilNum, val)
	}
}

func TestFloat64OrNilToString(t *testing.T) {
	// Test cases where float64 is not nil
	num := float64(123.45)
	if val := Float64OrNilToString(&num); val != "123.45" {
		t.Errorf("Float64OrNilToString(%v) returned unexpected result: %s", num, val)
	}

	// Test case where float64 is nil
	var nilNum *float64
	if val := Float64OrNilToString(nilNum); val != "" {
		t.Errorf("Float64OrNilToString(%v) returned unexpected result: %s", nilNum, val)
	}
}
