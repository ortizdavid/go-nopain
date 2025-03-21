package reflection

import (
	"fmt"
	"reflect"
)

// VarDump prints the type and value of multiple variables.
func VarDump(variables ...interface{}) {
	for _, variable := range variables {

		if variable == nil {
			fmt.Printf("Type: <nil>\nValue: <nil>\n\n")
			continue
		}

		v := reflect.ValueOf(variable)
		t := reflect.TypeOf(variable)

		fmt.Printf("Type: %s\n", t)

		switch v.Kind() {
		case reflect.Struct:
			fmt.Println("Value:")
			for i := 0; i < v.NumField(); i++ {
				fmt.Printf("  %s: %v\n", t.Field(i).Name, v.Field(i).Interface())
			}

		case reflect.Slice, reflect.Array:
			fmt.Println("Value: [")
			for i := 0; i < v.Len(); i++ {
				fmt.Printf("  %v\n", v.Index(i).Interface())
			}
			fmt.Println("]")

		case reflect.Map:
			fmt.Println("Value: {")
			for _, key := range v.MapKeys() {
				fmt.Printf("  %v: %v\n", key.Interface(), v.MapIndex(key).Interface())
			}
			fmt.Println("}")

		default:
			fmt.Printf("Value: %v\n", v.Interface())
		}

		fmt.Println()
	}
}


// VarDumpBasic basic prints the type and value of multiple variables.
func VarDumpBasic(variables ...interface{}) {
	for _, variable := range variables {
		if variable == nil {
			fmt.Printf("Type: <nil>\nValue: <nil>\n")
			continue
		}

		t := reflect.TypeOf(variable)
		v := reflect.ValueOf(variable)

		fmt.Printf("Type: %s\nValue: %#v\n\n", t, v)
	}
}
