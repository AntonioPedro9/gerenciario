package routes

import (
	handlers "server/cmd/api/handlers"
	"server/cmd/api/middlewares"
	"server/internals/repositories"
	services "server/internals/services"

	"github.com/gin-gonic/gin"
)

func SetupBudgetRoutes(r *gin.Engine, budgetRepository *repositories.BudgetRepository, userRepository *repositories.UserRepository) {
	budgetService := services.NewBudgetService(budgetRepository)
	budgetHandler := handlers.NewBudgetHandler(budgetService)

	api := r.Group("/api")
	auth := api.Group("", middlewares.AuthMiddleware(userRepository))

	auth.POST("/budgets", budgetHandler.CreateBudget)
	auth.GET("/budgets/:id", budgetHandler.GetBudget)
	auth.GET("/users/:id/budgets", budgetHandler.GetUserBudgets)
	auth.DELETE("/budgets/:id", budgetHandler.DeleteBudget)
}
