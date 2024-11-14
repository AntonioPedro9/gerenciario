package routes

import (
	handlers "server/cmd/api/handlers"
	"server/cmd/api/middlewares"
	"server/internals/repositories"
	services "server/internals/services"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(r *gin.Engine, customerRepository *repositories.CustomerRepository) {
	customerService := services.NewCustomerService(customerRepository)
	customerHandler := handlers.NewCustomerHandler(customerService)

	api := r.Group("/api")
	auth := api.Group("", middlewares.AuthMiddleware())

	auth.POST("/customers", customerHandler.CreateCustomer)
	auth.GET("/customers/:id", customerHandler.GetCustomer)
	auth.GET("/users/:id/customers", customerHandler.GetUserCustomers)
	auth.PUT("/customers/:id", customerHandler.UpdateCustomer)
	auth.DELETE("/customers/:id", customerHandler.DeleteCustomer)
}
