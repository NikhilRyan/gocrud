package crud

import (
	"encoding/json"
	"reflect"
	"strings"
)

func detectSQLType(value interface{}) string {
	switch value.(type) {
	case int, int32, int64:
		return "INTEGER"
	case float32, float64:
		return "REAL"
	case bool:
		return "BOOLEAN"
	case string:
		return "TEXT"
	default:
		// Check if the value can be marshalled into JSON
		if _, err := json.Marshal(value); err == nil {
			return "JSON"
		}
		return "BLOB"
	}
}

func getStructFields(dataStruct interface{}) string {
	val := reflect.ValueOf(dataStruct)

	// Check if it's a pointer to a slice
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Check if it's a slice
	if val.Kind() == reflect.Slice {
		val = reflect.New(val.Type().Elem()).Elem()
	}

	if val.Kind() != reflect.Struct {
		return "*"
	}

	var fields []string
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("gorm")
		if tag != "" {
			// Extract the column name from the tag
			tagParts := strings.Split(tag, ":")
			if len(tagParts) == 2 && tagParts[0] == "column" {
				fields = append(fields, tagParts[1])
			}
		}
	}

	return strings.Join(fields, ", ")
}
