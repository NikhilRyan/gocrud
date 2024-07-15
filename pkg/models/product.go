package models

type Product struct {
	ID    int     `json:"id" gorm:"column:id"`
	Name  string  `json:"name" gorm:"column:name"`
	Price float64 `json:"price" gorm:"column:price"`
}
