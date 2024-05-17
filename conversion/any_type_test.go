package conversion

import "testing"

func Test_AnyToString(t *testing.T) {
    // Test case for int
    var intValue interface{} = 123
    expectedInt := "123"
    resultInt := AnyToString(intValue)
    if resultInt != expectedInt {
        t.Errorf("AnyToString(%v) returned %s, expected %s", intValue, resultInt, expectedInt)
    }

    // Test case for float
    var floatValue interface{} = 123.45
    expectedFloat := "123.45"
    resultFloat := AnyToString(floatValue)
    if resultFloat != expectedFloat {
        t.Errorf("AnyToString(%v) returned %s, expected %s", floatValue, resultFloat, expectedFloat)
    }

    // Test case for string
    var strValue interface{} = "hello"
    expectedStr := "hello"
    resultStr := AnyToString(strValue)
    if resultStr != expectedStr {
        t.Errorf("AnyToString(%v) returned %s, expected %s", strValue, resultStr, expectedStr)
    }

    // Test case for boolean
    var boolValue interface{} = true
    expectedBool := "true"
    resultBool := AnyToString(boolValue)
    if resultBool != expectedBool {
        t.Errorf("AnyToString(%v) returned %s, expected %s", boolValue, resultBool, expectedBool)
    }

    // Test case for slice
    var sliceValue interface{} = []int{1, 2, 3}
    expectedSlice := "[1 2 3]"
    resultSlice := AnyToString(sliceValue)
    if resultSlice != expectedSlice {
        t.Errorf("AnyToString(%v) returned %s, expected %s", sliceValue, resultSlice, expectedSlice)
    }

    // Test case for struct
    type testStruct struct {
        Name  string
        Age   int
    }
    var structValue interface{} = testStruct{Name: "John", Age: 30}
    expectedStruct := "{John 30}"
    resultStruct := AnyToString(structValue)
    if resultStruct != expectedStruct {
        t.Errorf("AnyToString(%v) returned %s, expected %s", structValue, resultStruct, expectedStruct)
    }
}
