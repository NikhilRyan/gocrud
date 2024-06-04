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

func TestCreateHandler(t *testing.T) {
	// Setup
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	_ = db.AutoMigrate(&db.Request{})

	router := gin.Default()
	router.POST("/create", CreateHandler(db))

	// Test case
	reqBody := `{"table": "users", "key": "id", "data": {"name": "John Doe", "age": 30}}`
	req, _ := http.NewRequest("POST", "/create", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Record created successfully")
}
