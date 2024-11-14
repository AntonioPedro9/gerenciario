package routes

import (
	handlers "server/cmd/api/handlers"
	"server/cmd/api/middlewares"
	"server/internals/repositories"
	services "server/internals/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userRepository *repositories.UserRepository) {
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	api := r.Group("/api")
	auth := api.Group("", middlewares.AuthMiddleware())

	api.POST("/users", userHandler.CreateUser)
	auth.GET("/users/me", userHandler.GetUser)
	auth.PATCH("/users/me", userHandler.UpdateUserData)
	auth.PATCH("/users/me/password", userHandler.UpdateUserPassword)
	auth.DELETE("/users/me", userHandler.DeleteUser)
}
