package conversion

import "fmt"

// interface{} -> string.
func ConvertAnyToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
