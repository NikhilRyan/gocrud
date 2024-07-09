package handlers

import (
	"github.com/gin-gonic/gin"
	"gocrud/pkg/crud"
	"gocrud/pkg/models"
	"net/http"
)

func GenerateReadQueryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.QueryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query, params := crud.GenerateReadQuery(&req)
		c.JSON(http.StatusOK, gin.H{
			"query":  query,
			"params": params,
		})
	}
}

func GenerateReadWithJoinsQueryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.JoinRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query, params := crud.GenerateReadWithJoinsQuery(&req)
		c.JSON(http.StatusOK, gin.H{
			"query":  query,
			"params": params,
		})
	}
}
