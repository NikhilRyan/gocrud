package crud

import (
	"gorm.io/gorm"
)

type ColumnInfo struct {
	ColumnName string `gorm:"column:column_name"`
	DataType   string `gorm:"column:data_type"`
}

func GetColumnInfo(db *gorm.DB, tableName string) ([]ColumnInfo, error) {
	var columns []ColumnInfo
	query := `
        SELECT column_name, data_type
        FROM information_schema.columns
        WHERE table_name = ?
    `
	if err := db.Raw(query, tableName).Scan(&columns).Error; err != nil {
		return nil, err
	}
	return columns, nil
}
