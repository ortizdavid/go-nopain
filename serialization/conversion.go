package serialization

import (
	"encoding/xml"
	"fmt"
	"reflect"
)


// JsonToXml converts JSON data to XML format.
func JsonToXml(jsonData []byte) ([]byte, error) {
	// Unmarshal JSON data into a generic interface{}
	var v interface{}
	if err := UnserializeJson(jsonData, &v); err != nil {
		return nil, err
	}
	// Marshal the generic interface{} into XML format
	xmlData, err := SerializeXml(v)
	if err != nil {
		return nil, err
	}
	return xmlData, nil
}


// XmlToJson converts XML data to JSON format.
func XmlToJson(xmlData []byte) ([]byte, error) {
	// Unmarshal XML data into a generic interface{}
	var v interface{}
	if err := UnserializeXml(xmlData, &v); err != nil {
		return nil, err
	}
	// Marshal the generic interface{} into JSON format
	jsonData, err := SerializeJson(v)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}


// JsonToAsn1 converts JSON data to ASN.1 format.
func JsonToAsn1(jsonData []byte) ([]byte, error) {
	// Unmarshal JSON data into a generic interface{}
	var v interface{}
	if err := UnserializeJson(jsonData, &v); err != nil {
		return nil, err
	}
	// Marshal the generic interface{} into ASN.1 format
	asn1Data, err := SerializeAsn1(v)
	if err != nil {
		return nil, err
	}
	return asn1Data, nil
}


// XmlToAsn1 converts XML data to ASN.1 format.
func XmlToAsn1(xmlData []byte) ([]byte, error) {
	// Unmarshal XML data into a generic interface{}
	var v interface{}
	if err := UnserializeXml(xmlData, &v); err != nil {
		return nil, err
	}
	// Marshal the generic interface{} into ASN.1 format
	asn1Data, err := SerializeAsn1(v)
	if err != nil {
		return nil, err
	}
	return asn1Data, nil
}


// Map
// MapToJson converts a map[string]interface{} to JSON format.
func MapToJson(data map[string]interface{}) ([]byte, error) {
	jsonData, err := SerializeJson(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}


// MapToXml converts a map[string]interface{} to XML format.
func MapToXml(data map[string]interface{}) ([]byte, error) {
	// Create a new struct type with XML tags based on the keys in the map
	type dynamicStruct struct {
		XMLName xml.Name
		Fields  []xml.Name `xml:",any"`
	}
	// Create an instance of the dynamic struct
	instance := dynamicStruct{
		XMLName: xml.Name{Local: "Data"},
	}
	// Iterate over the map and create XML elements dynamically
	for key := range data {
		instance.Fields = append(instance.Fields, xml.Name{Local: key})
	}
	// Marshal the dynamic struct to XML format
	xmlData, err := SerializeXml(instance)
	if err != nil {
		return nil, err
	}
	return xmlData, nil
}


// Struct
// MapToStruct converts a map[string]interface{} to a struct.
func MapToStruct(data map[string]interface{}, obj interface{}) error {
	// Get the type of the provided struct
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Ptr || objType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("obj must be a pointer to a struct")
	}
	// Get the value of the provided struct
	objValue := reflect.ValueOf(obj).Elem()
	// Iterate over the fields of the struct
	for i := 0; i < objType.Elem().NumField(); i++ {
		fieldType := objType.Elem().Field(i)
		fieldName := fieldType.Name
		fieldValue := objValue.FieldByName(fieldName)
		// Check if the field exists in the map
		value, ok := data[fieldName]
		if !ok {
			continue // Field not found in map, skip
		}
		// Convert the value from the map to the field's type
		mapValue := reflect.ValueOf(value)
		if !mapValue.Type().AssignableTo(fieldType.Type) {
			return fmt.Errorf("value type does not match field type for field %s", fieldName)
		}
		// Set the field value
		fieldValue.Set(mapValue)
	}
	return nil
}


// StructToMap converts a struct to a map[string]interface{}.
func StructToMap(obj interface{}, data map[string]interface{}) error {
	// Get the type of the provided struct
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Struct {
		return fmt.Errorf("obj must be a struct")
	}
	// Get the value of the provided struct
	objValue := reflect.ValueOf(obj)
	// Iterate over the fields of the struct
	for i := 0; i < objType.NumField(); i++ {
		fieldType := objType.Field(i)
		fieldName := fieldType.Name
		fieldValue := objValue.Field(i).Interface()
		// Add the field name and value to the map
		data[fieldName] = fieldValue
	}
	return nil
}