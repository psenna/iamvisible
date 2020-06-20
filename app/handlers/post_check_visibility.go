package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/psenna/iamvisible/app/models"
)

// PostCheckVisibility Check is a http request is visible
func PostCheckVisibility() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := models.GetCheckRequest()

		c.ShouldBindJSON(&request)

		status, message := request.RunCheckRequest()

		c.JSON(status, message)
	}
}
