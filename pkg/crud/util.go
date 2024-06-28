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

func getStructFields(obj interface{}) string {
	t := reflect.TypeOf(obj)
	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("gorm")
		if tag != "" && strings.HasPrefix(tag, "column:") {
			columns := strings.Split(tag, ":")
			fields = append(fields, columns[1])
		} else {
			fields = append(fields, field.Name)
		}
	}
	return strings.Join(fields, ", ")
}
