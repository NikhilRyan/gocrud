package db

import (
	"fmt"

	"gorm.io/gorm"
)

func InsertData(db *gorm.DB, tableName string, data map[string]interface{}) error {
	insertSQL := fmt.Sprintf("INSERT INTO %s (", tableName)
	valuesSQL := "VALUES ("
	params := []interface{}{}

	for key, value := range data {
		insertSQL += fmt.Sprintf("%s, ", key)
		valuesSQL += "?, "
		params = append(params, value)
	}
	insertSQL = insertSQL[:len(insertSQL)-2] + ") "
	valuesSQL = valuesSQL[:len(valuesSQL)-2] + ")"

	finalSQL := insertSQL + valuesSQL

	return db.Exec(finalSQL, params...).Error
}
