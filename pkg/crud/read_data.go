package crud

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func ReadData(db *gorm.DB, tableName string, conditions map[string]interface{}) (interface{}, error) {
	var result []map[string]interface{}
	var conditionStrings []string
	var params []interface{}

	for key, value := range conditions {
		conditionStrings = append(conditionStrings, fmt.Sprintf("%s = ?", key))
		params = append(params, value)
	}

	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	if len(conditionStrings) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(conditionStrings, " AND "))
	}

	if err := db.Raw(query, params...).Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
