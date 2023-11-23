package reflection

import (
	"fmt"
	"reflect"
)


func VarDump(variable interface{}) {
	t := reflect.TypeOf(variable)
	v := reflect.ValueOf(variable)
	fmt.Printf("Type: %t, Value: %v\n\n", t, v)
}