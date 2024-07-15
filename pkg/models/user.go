package models

type User struct {
	ID    int    `gorm:"column:id" redis:"id"`
	Name  string `gorm:"column:name" redis:"name"`
	Age   int    `gorm:"column:age" redis:"age"`
	Email string `gorm:"column:email" redis:"email"`
}
