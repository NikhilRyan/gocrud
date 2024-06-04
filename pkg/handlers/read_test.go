package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestReadHandler(t *testing.T) {
	// Setup
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	_ = db.AutoMigrate(&db.Request{})

	router := gin.Default()
	router.GET("/read", ReadHandler(db))

	// Test case
	reqBody := `{"table": "users", "key": "id", "data": {"id": 1}}`
	req, _ := http.NewRequest("GET", "/read", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
