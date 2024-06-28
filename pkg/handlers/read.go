package handlers

import (
	"github.com/gin-gonic/gin"
	"gocrud/pkg/crud"
	"gorm.io/gorm"
	"net/http"
)

func ReadHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req crud.QueryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(req.Table) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Table name is required"})
			return
		}

		result, err := crud.ReadData(db, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
