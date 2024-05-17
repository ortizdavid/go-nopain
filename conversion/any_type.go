package conversion

import "fmt"

// interface{} -> string.
func AnyToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
