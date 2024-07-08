package main

import (
	"github.com/gin-gonic/gin"
	"gocrud/pkg/cache"
	"gocrud/pkg/crud"
	"gocrud/pkg/handlers"
	"gocrud/pkg/repositories"
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

	// Initialize caches
	crud.InitCache(10*time.Minute, 10*time.Minute)
	cache.InitMemCache(10*time.Minute, 10*time.Minute)
	cache.InitRedis("localhost:6379", "", 0)

	userRepository := repositories.NewUserRepository(db)

	router := gin.Default()
	router.POST("/create", handlers.CreateHandler(db))
	router.GET("/read", handlers.ReadHandler(db))
	router.POST("/read-with-joins", handlers.ReadWithJoinsHandler(db))
	router.PUT("/update", handlers.UpdateHandler(db))
	router.DELETE("/delete", handlers.DeleteHandler(db))
	router.GET("/columns", handlers.GetColumnInfoHandler(db))
	router.GET("/user", handlers.GetUserHandler(userRepository))

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
