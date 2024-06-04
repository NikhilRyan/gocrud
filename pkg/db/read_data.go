package db

import (
	"fmt"

	"gorm.io/gorm"
)

func ReadData(db *gorm.DB, tableName, key string, data map[string]interface{}) (interface{}, error) {
	var result []map[string]interface{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", tableName, key)
	value, exists := data[key]
	if !exists {
		return nil, fmt.Errorf("key %s not found in data", key)
	}

	if err := db.Raw(query, value).Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
