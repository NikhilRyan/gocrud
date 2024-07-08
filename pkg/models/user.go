package models

type User struct {
	UserID int    `redis:"user_id" gorm:"column:user_id"`
	Age    int    `redis:"age" gorm:"column:age"`
	Name   string `gorm:"column:name"`
}
