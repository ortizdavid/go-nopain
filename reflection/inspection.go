package reflection

import (
	"fmt"
	"reflect"
)

// InspectStruct prints information about the given object, including its kind, type, name, attributes, and public methods.
func InspectStruct(obj interface{}) {
	fmt.Println("Name: ", NameOfType(obj))
	fmt.Println("Kind: ", reflect.TypeOf(obj).Kind())
	fmt.Println("Type: ", reflect.TypeOf(obj))
	fmt.Println("Type Name: ", reflect.TypeOf(obj).Name())
	fmt.Println("Attributes: ", GetFields(obj))
	fmt.Println("Public Methods: ", GetPublicMethods(obj))
	fmt.Println()
}

// InspectStructFields prints information about the fields of the given object, including their names, values, and kinds.
func InspectStructFields(obj interface{}) {
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	numFields := objType.NumField()
	fmt.Println("Number of fields: ", numFields)
	fmt.Println("Field\tValue\tKind")
	for i := 0; i < numFields; i++ {
		fmt.Printf("%s\t%#v\t%v\n", objType.Field(i).Name, objVal.Field(i), objVal.Field(i).Kind())
	}
}

// CountTypes counts the occurrences of different types among the provided elements.
func CountTypes(elements ...interface{}) {
	var (
		countFloats   int
		countInts     int
		countStrings  int
		countBools    int
		countMaps     int
		countStructs  int
	)

	for _, elem := range elements {
		switch elem.(type) {
		case float32:
		case float64:
			countFloats++
		case int:
		case int16:
		case int32:
		case int64:
			countInts++
		case string:
			countStrings++
		case bool:
			countBools++
		case map[string]interface{}:
			countMaps++
		case struct{}:
			countStructs++
		}
	}
	fmt.Println("Integers: ", countInts)
	fmt.Println("Floats: ", countFloats)
	fmt.Println("Strings: ", countStrings)
	fmt.Println("Booleans: ", countBools)
	fmt.Println("Maps: ", countMaps)
	fmt.Println("Structs: ", countStructs)
}
