package reflection

import "reflect"

// GetTags returns a map with field names and associated tags
func GetTags(obj interface{}, tagKey string) map[string]string {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	tags := make(map[string]string)
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		tags[field.Name] = field.Tag.Get(tagKey)
	}
	return tags
}

// HasTag verify if a field has a tag
func HasTag(obj interface{}, fieldName string, tagKey string) bool {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	field, found := objType.FieldByName(fieldName)
	if !found {
		return false
	}

	tag := field.Tag.Get(tagKey)
	return tag != ""
}

// GetTag retrieves the value of a specific tag key for a given field
func GetTag(obj interface{}, fieldName string, tagKey string) string {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if objType.Kind() != reflect.Struct {
		return ""
	}

	field, found := objType.FieldByName(fieldName)
	if !found {
		return ""
	}

	return field.Tag.Get(tagKey)
}


// GetFieldsTag returns names of fields that contains a given tag
func GetFieldsWithTag(obj interface{}, tagKey string, tagValue string) []string {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	var fields []string
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if field.Tag.Get(tagKey) == tagValue {
			fields = append(fields, field.Name)
		}
	}
	return fields
}