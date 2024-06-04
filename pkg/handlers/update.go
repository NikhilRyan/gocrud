package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(req.Table) == 0 || len(req.Key) == 0 || len(req.Data) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Table name, key, and data are required"})
			return
		}

		if err := db.UpdateData(req.Table, req.Key, req.Data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
	}
}
