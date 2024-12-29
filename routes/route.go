package routes

import (
	"github.com/cholid97/go-kredit/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/contract", userController.GetContracts)
		userRoutes.GET("/:id", userController.GetUser)
		userRoutes.POST("", userController.CreateUser)
		userRoutes.POST("/ctr", userController.CreateContract)
	}
}
