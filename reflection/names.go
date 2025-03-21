package reflection

import (
	"reflect"
	"runtime"
	"strings"
)

// Name of Function -> includes package
func NameOfFunc(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

// Short Name of Function -> only name
func NameOfFuncShort(fn interface{}) string {
	parts := strings.Split(NameOfFunc(fn), ".")
	return parts[len(parts) - 1]
}

// Name of Interface or Struct
func NameOfType(v interface{}) string {
	return reflect.TypeOf(v).Name()
}
