package app

import (
	"github.com/gin-gonic/gin"
	"github.com/psenna/iamvisible/app/handlers"
)

func GetServer() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	ginInstance := gin.Default()

	ginInstance.GET("/", handlers.GetHome())

	ginInstance.POST("/check")

	return ginInstance
}
