package routes

import (
	handlers "server/cmd/api/handlers"
	"server/cmd/api/middlewares"
	"server/internals/repositories"
	services "server/internals/services"

	"github.com/gin-gonic/gin"
)

func SetupJobRoutes(r *gin.Engine, jobRepository *repositories.JobRepository) {
	jobService := services.NewJobService(jobRepository)
	jobHandler := handlers.NewJobHandler(jobService)

	api := r.Group("/api")
	auth := api.Group("", middlewares.AuthMiddleware())

	auth.POST("/jobs", jobHandler.CreateJob)
	auth.GET("/jobs/:id", jobHandler.GetJob)
	auth.GET("/users/:id/jobs", jobHandler.GetUserJobs)
	auth.PUT("/jobs/:id", jobHandler.UpdateJob)
	auth.DELETE("/jobs/:id", jobHandler.DeleteJob)
}
