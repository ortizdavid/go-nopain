package serialization

import (
	"bytes"
	"reflect"
	"testing"
)

type testJsonStruct struct {
	Name  string
	Age   int
	Email string
}

func Test_SerializeJson(t *testing.T) {
	obj := testJsonStruct{Name: "John", Age: 30, Email: "john@example.com"}

	jsonData, err := SerializeJson(obj)
	if err != nil {
		t.Errorf("SerializeJson() error = %v", err)
		return
	}

	var newObj testJsonStruct
	err = UnserializeJson(jsonData, &newObj)
	if err != nil {
		t.Errorf("UnserializeJson() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Deserialized object is not equal to the original object")
	}
}

func Test_UnserializeJson(t *testing.T) {

	obj := testJsonStruct{Name: "John", Age: 30, Email: "john@example.com"}
	jsonData, err := SerializeJson(obj)
	if err != nil {
		t.Errorf("SerializeJson() error = %v", err)
		return
	}

	var newObj testJsonStruct
	err = UnserializeJson(jsonData, &newObj)
	if err != nil {
		t.Errorf("UnserializeJson() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Deserialized object is not equal to the original object")
	}
}

func Test_EncodeJson(t *testing.T) {
	obj := testJsonStruct{Name: "John", Age: 30, Email: "john@example.com"}

	var buf bytes.Buffer
	err := EncodeJson(&buf, obj)
	if err != nil {
		t.Errorf("EncodeJson() error = %v", err)
		return
	}

	var newObj testJsonStruct
	err = DecodeJson(&buf, &newObj)
	if err != nil {
		t.Errorf("DecodeJson() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Decoded object is not equal to the original object")
	}
}

func Test_DecodeJson(t *testing.T) {

	obj := testJsonStruct{Name: "John", Age: 30, Email: "john@example.com"}
	var buf bytes.Buffer
	err := EncodeJson(&buf, obj)
	if err != nil {
		t.Errorf("EncodeJson() error = %v", err)
		return
	}

	var newObj testJsonStruct
	err = DecodeJson(&buf, &newObj)
	if err != nil {
		t.Errorf("DecodeJson() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Decoded object is not equal to the original object")
	}
}
