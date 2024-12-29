package main

import (
	"github.com/cholid97/go-kredit/config"
	"github.com/cholid97/go-kredit/controllers"
	"github.com/cholid97/go-kredit/repositories"
	"github.com/cholid97/go-kredit/routes"
	"github.com/cholid97/go-kredit/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Setup Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, userController)

	// Start the server
	router.Run(":8080")
}
