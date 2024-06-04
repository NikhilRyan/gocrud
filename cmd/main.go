package main

import (
	"gocrud/pkg/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()
	router.POST("/create", handlers.CreateHandler(db))
	router.GET("/read", handlers.ReadHandler(db))
	router.PUT("/update", handlers.UpdateHandler(db))
	router.DELETE("/delete", handlers.DeleteHandler(db))

	router.Run(":8080")
}
