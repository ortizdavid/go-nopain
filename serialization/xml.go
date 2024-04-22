package serialization

import (
	"encoding/xml"
	"io"
)


// SerializeXml serializes the given object to JSON
func SerializeXml(obj interface{}) ([]byte, error) {
	xmlData, err := xml.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return xmlData, nil
}

// UnserializeXml deserializes JSON data into the given data structure
func UnserializeXml(xmlData []byte, obj interface{}) error {
	err := xml.Unmarshal(xmlData, obj)
	if err != nil {
		return err
	}
	return nil
}

// EncodeXml encodes the given object as JSON and writes it to the writer
func EncodeXml(obj interface{}, writer io.Writer) error {
	encoder := xml.NewEncoder(writer)
	err := encoder.Encode(obj)
	if err != nil {
		return err
	}
	return nil
}

// DecodeXml decodes JSON from the reader into the given data structure
func DecodeXml(reader io.Reader, v interface{}) error {
	decoder := xml.NewDecoder(reader)
	err := decoder.Decode(v)
	if err != nil {
		return err
	}
	return nil
}