package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	ID   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func TestReadHandler(t *testing.T) {
	// Setup
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	router := gin.Default()
	router.GET("/read", ReadHandler(db))

	// Test case
	reqBody := `{
        "table": "users",
        "conditions": {"age": 30},
        "order_by": ["name ASC", "age DESC"],
        "limit": 10,
        "offset": 5,
        "struct": User{}
    }`
	req, _ := http.NewRequest("GET", "/read", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
