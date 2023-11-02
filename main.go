package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Binding from json
type Register struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		var json Register
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"register": json,
		})
	})
	r.Run("localhost:8080")
}
