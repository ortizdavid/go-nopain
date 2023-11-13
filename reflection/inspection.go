package reflection

import (
	"fmt"
	"reflect"
)

func VarDump(variable interface{}) {
	fmt.Printf("Type: %s, Value: %v\n", reflect.TypeOf(variable), variable)
}