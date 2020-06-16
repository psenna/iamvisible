package handlers

import "github.com/gin-gonic/gin"

// PostCheckVisibility Check is a http request is visible
func PostCheckVisibility() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Check a system visibility",
			"docs":    "todo",
		})
	}
}
