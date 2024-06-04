package db

import (
	"encoding/json"
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
