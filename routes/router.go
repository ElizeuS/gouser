package routes

import (
	"github.com/ElizeuS/gouser/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.ShowUser)
		}

	}
	return router
}
