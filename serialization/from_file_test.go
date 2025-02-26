package serialization

import (
	"os"
	"reflect"
	"testing"
)

func Test_FromJsonFile(t *testing.T) {
	type TestStruct struct {
		Name  string
		Age   int
		Email string
	}

	// Create a temporary JSON file for testing
	testData := []byte(`{"Name": "John", "Age": 30, "Email": "john@example.com"}`)
	tmpfile := createTempFile(t, "test.json", testData)
	defer os.Remove(tmpfile.Name())

	var obj TestStruct
	err := FromJsonFile(tmpfile.Name(), &obj)
	if err != nil {
		t.Errorf("FromJsonFile() error = %v", err)
		return
	}

	expected := TestStruct{Name: "John", Age: 30, Email: "john@example.com"}
	if !reflect.DeepEqual(obj, expected) {
		t.Errorf("Deserialized object is not equal to the expected object")
	}
}

func Test_FromXmlFile(t *testing.T) {
	type TestStruct struct {
		Name  string
		Age   int
		Email string
	}

	// Create a temporary XML file for testing
	testData := []byte(`<TestStruct><Name>John</Name><Age>30</Age><Email>john@example.com</Email></TestStruct>`)
	tmpfile := createTempFile(t, "test.xml", testData)
	defer os.Remove(tmpfile.Name())

	var obj TestStruct
	err := FromXmlFile(tmpfile.Name(), &obj)
	if err != nil {
		t.Errorf("FromXmlFile() error = %v", err)
		return
	}

	expected := TestStruct{Name: "John", Age: 30, Email: "john@example.com"}
	if !reflect.DeepEqual(obj, expected) {
		t.Errorf("Deserialized object is not equal to the expected object")
	}
}

func Test_FromAsnFile(t *testing.T) {
	type TestStruct struct {
		Name  string
		Age   int
		Email string
	}

	// Create a temporary ASN.1 file for testing
	testData, err := SerializeAsn1(TestStruct{Name: "John", Age: 30, Email: "john@example.com"})
	if err != nil {
		t.Errorf("SerializeAsn1() error = %v", err)
		return
	}
	tmpfile := createTempFile(t, "test.asn", testData)
	defer os.Remove(tmpfile.Name())

	var obj TestStruct
	err = FromAsn1File(tmpfile.Name(), &obj)
	if err != nil {
		t.Errorf("FromAsnFile() error = %v", err)
		return
	}

	expected := TestStruct{Name: "John", Age: 30, Email: "john@example.com"}
	if !reflect.DeepEqual(obj, expected) {
		t.Errorf("Deserialized object is not equal to the expected object")
	}
}

func createTempFile(t *testing.T, filename string, data []byte) *os.File {
	tmpfile, err := os.CreateTemp("", filename)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	if _, err := tmpfile.Write(data); err != nil {
		t.Fatalf("Failed to write data to temporary file: %v", err)
	}
	return tmpfile
}
