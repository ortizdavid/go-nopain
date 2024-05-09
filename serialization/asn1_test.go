package serialization

import (
	"bytes"
	"reflect"
	"testing"
)

type testAsn1Struct struct {
	Name  string
	Age   int
	Email string
}

func Test_SerializeAsn1(t *testing.T) {
	
	obj := testAsn1Struct{Name: "John", Age: 30, Email: "john@example.com"}

	asn1Data, err := SerializeAsn1(obj)
	if err != nil {
		t.Errorf("SerializeAsn1() error = %v", err)
		return
	}

	var newObj testAsn1Struct
	err = UnserializeAsn1(asn1Data, &newObj)
	if err != nil {
		t.Errorf("UnserializeAsn1() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Deserialized object is not equal to the original object")
	}
}

func Test_UnserializeAsn1(t *testing.T) {
	
	obj := testAsn1Struct{Name: "John", Age: 30, Email: "john@example.com"}
	asn1Data, err := SerializeAsn1(obj)
	if err != nil {
		t.Errorf("SerializeAsn1() error = %v", err)
		return
	}

	var newObj testAsn1Struct
	err = UnserializeAsn1(asn1Data, &newObj)
	if err != nil {
		t.Errorf("UnserializeAsn1() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Deserialized object is not equal to the original object")
	}
}

func Test_EncodeAsn1(t *testing.T) {
	
	obj := testAsn1Struct{Name: "John", Age: 30, Email: "john@example.com"}

	var buf bytes.Buffer
	err := EncodeAsn1(&buf, obj)
	if err != nil {
		t.Errorf("EncodeAsn1() error = %v", err)
		return
	}

	var newObj testAsn1Struct
	err = DecodeAsn1(&buf, &newObj)
	if err != nil {
		t.Errorf("DecodeAsn1() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Decoded object is not equal to the original object")
	}
}

func Test_DecodeAsn1(t *testing.T) {

	obj := testAsn1Struct{Name: "John", Age: 30, Email: "john@example.com"}
	var buf bytes.Buffer
	err := EncodeAsn1(&buf, obj)
	if err != nil {
		t.Errorf("EncodeAsn1() error = %v", err)
		return
	}

	var newObj testAsn1Struct
	err = DecodeAsn1(&buf, &newObj)
	if err != nil {
		t.Errorf("DecodeAsn1() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Decoded object is not equal to the original object")
	}
}
