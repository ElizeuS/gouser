package routes

import (
	"github.com/ElizeuS/gouser/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("cliente")
		{
			users.POST("/", controllers.CreateUser)
			users.GET("/:uuid", controllers.ShowUser)
			users.PUT("/:uuid", controllers.UpdateUser)
			users.DELETE("/:uuid", controllers.DeleteUser)
			users.GET("/", controllers.ShowUsers)
		}
	}
	return router
}
