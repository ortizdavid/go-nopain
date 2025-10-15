package conversion

import (
	"strconv"
	"time"
)

// string -> int or nil
func StringToIntOrNil(strNumber string) *int {
	if strNumber == "" {
		return nil
	}
	numInt, err := strconv.Atoi(strNumber)
	if err != nil {
		return nil
	}
	return &numInt
}

// string -> int32 or nil
func StringToInt32OrNil(strNumber string) *int32 {
	if strNumber == "" {
		return nil
	}
	numInt, err := strconv.ParseInt(strNumber, 10, 32)
	if err != nil {
		return nil
	}
	result := int32(numInt)
	return &result
}

// string -> int64 or nil
func StringToInt64OrNil(strNumber string) *int64 {
	if strNumber == "" {
		return nil
	}
	numInt, err := strconv.ParseInt(strNumber, 10, 64)
	if err != nil {
		return nil
	}
	result := int64(numInt)
	return &result
}

// string -> float32 or nil
func StringToFloat32OrNil(strNumber string) *float32 {
	if strNumber == "" {
		return nil
	}
	numFloat, err := strconv.ParseFloat(strNumber, 32)
	if err != nil {
		return nil
	}
	result := float32(numFloat)
	return &result
}

// string -> float64 or nil
func StringToFloat64OrNil(strNumber string) *float64 {
	if strNumber == "" {
		return nil
	}
	numFloat, err := strconv.ParseFloat(strNumber, 64)
	if err != nil {
		return nil
	}
	result := float64(numFloat)
	return &result
}

// string -> String or Nil
func StringOrNil(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

// Int or Nil -> string
func IntOrNilToString(number *int) string {
	if number == nil {
		return ""
	}
	return strconv.Itoa(*number)
}

// Int32 or Nil -> string
func Int32OrNilToString(number *int32) string {
	if number == nil {
		return ""
	}
	return strconv.Itoa(int(*number))
}

// Int642 or Nil -> string
func Int64OrNilToString(number *int64) string {
	if number == nil {
		return ""
	}
	return strconv.Itoa(int(*number))
}

// Int32 or Nil -> string
func Float32OrNilToString(number *float32) string {
	if number == nil {
		return ""
	}
	return strconv.FormatFloat(float64(*number), 'f', -1, 32)
}

// Int642 or Nil -> string
func Float64OrNilToString(number *float64) string {
	if number == nil {
		return ""
	}
	return strconv.FormatFloat(float64(*number), 'f', -1, 64)
}

// string -> *string
func StringToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// *string -> string
func PtrToString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

// formatTimePtr converts *time.Time to *string in RFC3339 format
func FormatTimePtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}
