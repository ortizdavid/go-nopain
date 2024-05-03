package conversion

import "testing"

func TestConvertAnyToString(t *testing.T) {
    // Test case for int
    var intValue interface{} = 123
    expectedInt := "123"
    resultInt := ConvertAnyToString(intValue)
    if resultInt != expectedInt {
        t.Errorf("ConvertAnyToString(%v) returned %s, expected %s", intValue, resultInt, expectedInt)
    }

    // Test case for float
    var floatValue interface{} = 123.45
    expectedFloat := "123.45"
    resultFloat := ConvertAnyToString(floatValue)
    if resultFloat != expectedFloat {
        t.Errorf("ConvertAnyToString(%v) returned %s, expected %s", floatValue, resultFloat, expectedFloat)
    }

    // Test case for string
    var strValue interface{} = "hello"
    expectedStr := "hello"
    resultStr := ConvertAnyToString(strValue)
    if resultStr != expectedStr {
        t.Errorf("ConvertAnyToString(%v) returned %s, expected %s", strValue, resultStr, expectedStr)
    }

    // Test case for boolean
    var boolValue interface{} = true
    expectedBool := "true"
    resultBool := ConvertAnyToString(boolValue)
    if resultBool != expectedBool {
        t.Errorf("ConvertAnyToString(%v) returned %s, expected %s", boolValue, resultBool, expectedBool)
    }

    // Test case for slice
    var sliceValue interface{} = []int{1, 2, 3}
    expectedSlice := "[1 2 3]"
    resultSlice := ConvertAnyToString(sliceValue)
    if resultSlice != expectedSlice {
        t.Errorf("ConvertAnyToString(%v) returned %s, expected %s", sliceValue, resultSlice, expectedSlice)
    }

    // Test case for struct
    type testStruct struct {
        Name  string
        Age   int
    }
    var structValue interface{} = testStruct{Name: "John", Age: 30}
    expectedStruct := "{John 30}"
    resultStruct := ConvertAnyToString(structValue)
    if resultStruct != expectedStruct {
        t.Errorf("ConvertAnyToString(%v) returned %s, expected %s", structValue, resultStruct, expectedStruct)
    }
}
