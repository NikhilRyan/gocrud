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

type OrderWithUser struct {
	OrderID     int    `gorm:"column:order_id"`
	Total       int    `gorm:"column:total"`
	Name        string `gorm:"column:name"`
	Email       string `gorm:"column:email"`
	ProductName string `gorm:"column:product_name"`
}

func TestReadWithJoinsHandler(t *testing.T) {
	// Setup
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	router := gin.Default()
	router.POST("/read-with-joins", ReadWithJoinsHandler(db))

	// Test case
	reqBody := `{
        "table": "orders",
        "joins": [
            {
                "join_type": "INNER",
                "table": "users",
                "on_condition": "orders.user_id = users.id",
                "selects": ["users.name", "users.email"],
                "conditions": {
                    "users.age": 30
                }
            },
            {
                "join_type": "LEFT",
                "table": "products",
                "on_condition": "orders.product_id = products.id",
                "selects": ["products.name as product_name"],
                "conditions": {
                    "products.price >": 100
                }
            }
        ],
        "columns": ["orders.id", "orders.total"],
        "conditions": {
            "orders.status": "completed",
            "orders.created_at >": "2022-01-01",
            "orders.note LIKE": "%urgent%"
        },
        "order_by": ["orders.date DESC"],
        "limit": 10,
        "offset": 5,
        "struct": OrderWithUser{}
    }`
	req, _ := http.NewRequest("POST", "/read-with-joins", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
