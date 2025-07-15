package reflection

import (
	"reflect"
)

func GetPublicMethods(obj interface{}) []string {
	objType := reflect.TypeOf(obj)
	var methods []string
	for i := 0; i < objType.NumMethod(); i++ {
		methods = append(methods, objType.Method(i).Name)
	}
	return methods
}

func GetFields(obj interface{}) []string {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem() 
	}
	var fields []string
	for i := 0; i < objType.NumField(); i++ {
		fields = append(fields, objType.Field(i).Name)
	}
	return fields
}

func ExistsField(obj interface{}, fieldName string) bool {
	field := reflect.ValueOf(obj).FieldByName(fieldName)
	return field.IsValid() 
}

func ExistsMethod(obj interface{}, methodName string) bool {
	method := reflect.ValueOf(obj).MethodByName(methodName)
	return method.IsValid() 
}
	
func CallMethod(any interface{}, methodName string, args... interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(methodName).Call(inputs)
}
