package crud

import (
	"fmt"

	"gorm.io/gorm"
)

func UpdateData(db *gorm.DB, tableName, key string, data map[string]interface{}) error {
	updateSQL := fmt.Sprintf("UPDATE %s SET ", tableName)
	params := []interface{}{}

	for col, value := range data {
		if col != key {
			updateSQL += fmt.Sprintf("%s = ?, ", col)
			params = append(params, value)
		}
	}
	updateSQL = updateSQL[:len(updateSQL)-2] + fmt.Sprintf(" WHERE %s = ?", key)
	params = append(params, data[key])

	return db.Exec(updateSQL, params...).Error
}
