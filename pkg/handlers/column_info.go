package handlers

import (
	"github.com/gin-gonic/gin"
	"gocrud/pkg/crud"
	"gorm.io/gorm"
	"net/http"
)

func GetColumnInfoHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableName := c.Query("table")
		if tableName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Table name is required"})
			return
		}

		columns, err := crud.GetColumnInfo(db, tableName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"columns": columns})
	}
}
