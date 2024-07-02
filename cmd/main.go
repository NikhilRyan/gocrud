package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"gocrud/pkg/crud"
	"gocrud/pkg/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Initialize cache
	crud.InitCache(cache.DefaultExpiration, 10*time.Minute)

	router := gin.Default()
	router.POST("/create", handlers.CreateHandler(db))
	router.GET("/read", handlers.ReadHandler(db))
	router.POST("/read-with-joins", handlers.ReadWithJoinsHandler(db))
	router.PUT("/update", handlers.UpdateHandler(db))
	router.DELETE("/delete", handlers.DeleteHandler(db))
	router.GET("/columns", handlers.GetColumnInfoHandler(db))

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
