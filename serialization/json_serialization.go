package serialization

import (
	"encoding/json"
	"io"
)


// SerializeJson serializes the given object to JSON
func SerializeJson(obj interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// UnserializeJson deserializes JSON data into the given data structure
func UnserializeJson(jsonData []byte, obj interface{}) error {
	err := json.Unmarshal(jsonData, obj)
	if err != nil {
		return err
	}
	return nil
}

// EncodeJson encodes the given object as JSON and writes it to the writer
func EncodeJson(obj interface{}, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(obj)
	if err != nil {
		return err
	}
	return nil
}

// DecodeJson decodes JSON from the reader into the given data structure
func DecodeJson(reader io.Reader, v interface{}) error {
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(v)
	if err != nil {
		return err
	}
	return nil
}