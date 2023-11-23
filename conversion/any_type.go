package conversion

import "fmt"

func ConvertAnyToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}