package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthToken Retun a middleware that authorize the acess based on Authorization header
// Temporary just for tests
func AuthToken(authToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		if !strings.Contains(bearerToken, "Bearer") || !strings.Contains(bearerToken, authToken) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
