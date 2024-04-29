package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// ConvertToCamelCase recursively converts keys in a struct to camel case.
func ConvertToCamelCase(data interface{}) interface{} {
	fmt.Println(data)
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Map {
		return convertMapKeysToCamelCase(data.(map[string]interface{}))
	}

	// Convert struct fields to camel case
	return convertStructFieldsToCamelCase(data)
}

// convertMapKeysToCamelCase converts keys in a map to camel case.
func convertMapKeysToCamelCase(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range data {
		// Convert key to camel case
		camelCaseKey := toCamelCase(key)

		// Check if value is another map
		if childMap, ok := value.(map[string]interface{}); ok {
			// Recursively convert child map keys to camel case
			result[camelCaseKey] = convertMapKeysToCamelCase(childMap)
		} else {
			result[camelCaseKey] = value
		}
	}

	return result
}

// convertStructFieldsToCamelCase converts keys in a struct to camel case.
func convertStructFieldsToCamelCase(data interface{}) interface{} {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Create a new instance of the same type as the input
	newData := reflect.New(v.Type()).Interface()

	// Iterate over struct fields and convert keys to camel case
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		camelCaseKey := toCamelCase(field.Name)
		reflect.ValueOf(newData).Elem().FieldByName(camelCaseKey).Set(v.Field(i))
	}
	return newData
}

// toCamelCase converts a string to camel case.
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if i > 0 {
			parts[i] = strings.Title(part)
		}
	}
	return strings.Join(parts, "")
}
