package serialization

import (
	"bytes"
	"reflect"
	"testing"
)

type testXmlStruct struct {
	Name  string
	Age   int
	Email string
}


func TestSerializeXml(t *testing.T) {
	obj := testXmlStruct{Name: "John", Age: 30, Email: "john@example.com"}

	xmlData, err := SerializeXml(obj)
	if err != nil {
		t.Errorf("SerializeXml() error = %v", err)
		return
	}

	var newObj testXmlStruct
	err = UnserializeXml(xmlData, &newObj)
	if err != nil {
		t.Errorf("UnserializeXml() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Deserialized object is not equal to the original object")
	}
}

func TestUnserializeXml(t *testing.T) {

	obj := testXmlStruct{Name: "John", Age: 30, Email: "john@example.com"}
	xmlData, err := SerializeXml(obj)
	if err != nil {
		t.Errorf("SerializeXml() error = %v", err)
		return
	}

	var newObj testXmlStruct
	err = UnserializeXml(xmlData, &newObj)
	if err != nil {
		t.Errorf("UnserializeXml() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Deserialized object is not equal to the original object")
	}
}

func TestEncodeXml(t *testing.T) {

	obj := testXmlStruct{Name: "John", Age: 30, Email: "john@example.com"}

	var buf bytes.Buffer
	err := EncodeXml(&buf, obj)
	if err != nil {
		t.Errorf("EncodeXml() error = %v", err)
		return
	}

	var newObj testXmlStruct
	err = DecodeXml(&buf, &newObj)
	if err != nil {
		t.Errorf("DecodeXml() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Decoded object is not equal to the original object")
	}
}

func TestDecodeXml(t *testing.T) {
	
	obj := testXmlStruct{Name: "John", Age: 30, Email: "john@example.com"}
	var buf bytes.Buffer
	err := EncodeXml(&buf, obj)
	if err != nil {
		t.Errorf("EncodeXml() error = %v", err)
		return
	}

	var newObj testXmlStruct
	err = DecodeXml(&buf, &newObj)
	if err != nil {
		t.Errorf("DecodeXml() error = %v", err)
		return
	}

	if !reflect.DeepEqual(obj, newObj) {
		t.Errorf("Decoded object is not equal to the original object")
	}
}
