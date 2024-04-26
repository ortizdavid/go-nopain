package serialization

import (
    "encoding/asn1"
    "io"
)


// SerializeAsn1 serializes the given object to ASN.1 format.
func SerializeAsn1(obj interface{}) ([]byte, error) {
    asn1Data, err := asn1.Marshal(obj)
    if err != nil {
        return nil, err
    }
    return asn1Data, nil
}


// UnserializeAsn1 deserializes ASN.1 data into the given data structure.
func UnserializeAsn1(asn1Data []byte, obj interface{}) error {
    _, err := asn1.Unmarshal(asn1Data, obj)
    if err != nil {
        return err
    }
    return nil
}

// EncodeAsn1 encodes the given object as ASN.1 and writes it to the writer.
func EncodeAsn1(writer io.Writer, obj interface{}) error {
    asn1Data, err := asn1.Marshal(obj)
    if err != nil {
        return err
    }
    _, err = writer.Write(asn1Data)
    return err
}

// DecodeAsn1 decodes ASN.1 from the reader into the given data structure.
func DecodeAsn1(reader io.Reader, v interface{}) error {
    asn1Data, err := io.ReadAll(reader)
    if err != nil {
        return err
    }
    _, err = asn1.Unmarshal(asn1Data, v)
    return err
}
