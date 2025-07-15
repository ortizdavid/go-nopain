package reflection

import (
	"fmt"
	"reflect"
)

func InstanceOf(obj interface{}, kind interface{}) bool {
	return reflect.TypeOf(obj) == reflect.TypeOf(kind)
}

// check if is atype xx or yy
func CheckType(v interface{}) {
	fmt.Printf("Type is %v", reflect.TypeOf(v))
}