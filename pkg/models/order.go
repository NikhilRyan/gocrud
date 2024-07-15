package models

type Order struct {
	ID        int     `gorm:"column:id" redis:"id"`
	UserID    int     `gorm:"column:user_id" redis:"user_id"`
	ProductID int     `gorm:"column:product_id" redis:"product_id"`
	Total     float64 `gorm:"column:total" redis:"total"`
	Status    string  `gorm:"column:status" redis:"status"`
	Date      string  `gorm:"column:date" redis:"date"`
}
