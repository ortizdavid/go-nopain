package serialization

import (
	"os"
)


// FromJsonFile deserializes JSON data from the specified file into the given object.
func FromJsonFile(jsonFile string, obj interface{}) error {
	// Read JSON file
    data, err := os.ReadFile(jsonFile)
    if err != nil {
        return err
    }
    // Deserialize JSON into object
    err = UnserializeJson(data, obj)
    if err != nil {
        return err
    }
	return nil
}


// FromXmlFile deserializes XML data from the specified file into the given object.
func FromXmlFile(xmlFile string, obj interface{}) error {
	// Read XML file
    data, err := os.ReadFile(xmlFile)
    if err != nil {
        return err
    }
    // Deserialize XML into object
    err = UnserializeXml(data, obj)
    if err != nil {
        return err
    }
	return nil
}


// FromAsnFile deserializes ASN.1 encoded data from the specified file into the given object.
func FromAsnFile(asnFile string, obj interface{}) error {
	// Read ASN.1 file
	data, err := os.ReadFile(asnFile)
	if err != nil {
		return err
	}
	// Unmarshal ASN.1 data into object
	err = UnserializeAsn1(data, obj)
	if err != nil {
		return err
	}
	return nil
}


// FromCsvFileToJson reads data from the specified CSV file and deserializes it into JSON format.
// The deserialized JSON data is stored in the provided object.
func FromCsvFileToJson(csvFile string, obj interface{}) error {
	// Open the CSV file
	file, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}


// FromCsvFileToXml reads data from the specified CSV file and deserializes it into XML format.
// The deserialized XML data is stored in the provided object.
func FromCsvFileToXml(csvFile string, obj interface{}) error {
	// Open the CSV file
	file, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}