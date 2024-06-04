package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Request struct {
	Table string                 `json:"table" binding:"required"`
	Key   string                 `json:"key" binding:"required"`
	Data  map[string]interface{} `json:"data"`
}

func CreateHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(req.Table) == 0 || len(req.Key) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Table name and key are required"})
			return
		}

		if err := db.CreateTable(req.Table, req.Data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.InsertData(req.Table, req.Data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Record created successfully"})
	}
}
