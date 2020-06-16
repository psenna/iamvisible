package handlers

import "github.com/gin-gonic/gin"

// GetHome the api index
func GetHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Check a system visibility",
			"docs":    "todo",
		})
	}
}
