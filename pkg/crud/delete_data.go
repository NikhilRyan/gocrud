package crud

import (
	"fmt"

	"gorm.io/gorm"
)

func DeleteData(db *gorm.DB, tableName, key string) error {
	deleteSQL := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", tableName, key)
	return db.Exec(deleteSQL, key).Error
}
