package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/psenna/iamvisible/app/handlers"
	"github.com/psenna/iamvisible/app/middlewares"
)

func GetServer() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	ginInstance := gin.Default()

	ginInstance.GET("/", handlers.GetHome())

	protected := ginInstance.Group("")

	if token := os.Getenv("IAMVISIBLE_TOKEN"); token != "" {
		fmt.Println("Check protected with token")
		protected.Use(middlewares.AuthToken(token))
	}

	protected.POST("/check", handlers.PostCheckVisibility())

	return ginInstance
}
