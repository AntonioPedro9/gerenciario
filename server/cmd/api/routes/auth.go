package routes

import (
	handlers "server/cmd/api/handlers"
	"server/internals/repositories"
	services "server/internals/services"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, userRepository *repositories.UserRepository) {
	authService := services.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)

	api := r.Group("/api")
	api.POST("/auth/token", authHandler.Auth)
	api.POST("/auth/refresh", authHandler.Refresh)
}
