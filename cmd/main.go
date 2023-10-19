package main

import (
	"server/database"
	"server/handlers"
	"server/middlewares"
	"server/repositories"
	"server/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	// configure logger
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "04/04/2001 15:00",
	})

	// load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	db := database.ConnectToDatabase()
	database.CreateDatabaseTables(db)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/", userHandler.ListUsers)
		userGroup.PUT("/", middlewares.RequireAuth, userHandler.UpdateUser)
		userGroup.DELETE("/:id", middlewares.RequireAuth, userHandler.DeleteUser)
		userGroup.POST("/login", userHandler.LoginUser)
	}

	clientRepository := repositories.NewClientRepository(db)
	clientService := services.NewClientService(clientRepository)
	clientHandler := handlers.NewClientHandler(clientService)
	clientGroup := r.Group("/clients")
	{
		clientGroup.POST("/", middlewares.RequireAuth, clientHandler.CreateClient)
		clientGroup.GET("/:userID", middlewares.RequireAuth, clientHandler.ListClients)
		clientGroup.PUT("/", middlewares.RequireAuth, clientHandler.UpdateClient)
		clientGroup.DELETE("/:clientID", middlewares.RequireAuth, clientHandler.DeleteClient)
	}

	serviceRepository := repositories.NewServiceRepository(db)
	serviceService := services.NewServiceService(serviceRepository)
	serviceHandler := handlers.NewServiceHandler(serviceService)
	serviceGroup := r.Group("/services")
	{
		serviceGroup.POST("/", middlewares.RequireAuth, serviceHandler.CreateService)
		serviceGroup.GET("/:userID", middlewares.RequireAuth, serviceHandler.ListServices)
		serviceGroup.PUT("/", middlewares.RequireAuth, serviceHandler.UpdateService)
		serviceGroup.DELETE("/:serviceID", middlewares.RequireAuth, serviceHandler.DeleteService)
	}

	budgetRepository := repositories.NewBudgetRepository(db)
	budgetService := services.NewBudgetService(budgetRepository)
	budgetHandler := handlers.NewBudgetHandler(budgetService)
	budgetGroup := r.Group("/budgets")
	{
		budgetGroup.POST("/", middlewares.RequireAuth, budgetHandler.CreateBudget)
		budgetGroup.GET("/:userID", middlewares.RequireAuth, budgetHandler.ListBudgets)
		budgetGroup.DELETE("/:budgetID", middlewares.RequireAuth, budgetHandler.DeleteBudget)
	}

	r.Run()
}
