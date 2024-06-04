package db

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB, tableName string, data map[string]interface{}) error {
	createTableSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT", tableName)
	for key, value := range data {
		createTableSQL += fmt.Sprintf(", %s %s", key, detectSQLType(value))
	}
	createTableSQL += ")"

	return db.Exec(createTableSQL).Error
}
