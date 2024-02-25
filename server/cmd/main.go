package main

import (
	"server/database"
	"server/handlers"
	"server/middlewares"
	"server/repositories"
	"server/services"

	"github.com/gin-contrib/cors"
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
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"localhost"})

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	db := database.ConnectToDatabase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	database.CreateDatabaseTables(db)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", userHandler.CreateUser)
		// userGroup.GET("/list", userHandler.ListUsers)
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
		clientGroup.GET("/list/:userID", middlewares.RequireAuth, clientHandler.ListClients)
		clientGroup.GET("/:clientID", middlewares.RequireAuth, clientHandler.GetClient)
		clientGroup.PUT("/", middlewares.RequireAuth, clientHandler.UpdateClient)
		clientGroup.DELETE("/:clientID", middlewares.RequireAuth, clientHandler.DeleteClient)
	}

	serviceRepository := repositories.NewServiceRepository(db)
	serviceService := services.NewServiceService(serviceRepository)
	serviceHandler := handlers.NewServiceHandler(serviceService)
	serviceGroup := r.Group("/services")
	{
		serviceGroup.POST("/", middlewares.RequireAuth, serviceHandler.CreateService)
		serviceGroup.GET("/list/:userID", middlewares.RequireAuth, serviceHandler.ListServices)
		serviceGroup.GET("/:serviceID", middlewares.RequireAuth, serviceHandler.GetService)
		serviceGroup.PUT("/", middlewares.RequireAuth, serviceHandler.UpdateService)
		serviceGroup.DELETE("/:serviceID", middlewares.RequireAuth, serviceHandler.DeleteService)
	}

	budgetRepository := repositories.NewBudgetRepository(db)
	budgetService := services.NewBudgetService(budgetRepository)
	budgetHandler := handlers.NewBudgetHandler(budgetService)
	budgetGroup := r.Group("/budgets")
	{
		budgetGroup.POST("/", middlewares.RequireAuth, budgetHandler.CreateBudget)
		budgetGroup.GET("/list/:userID", middlewares.RequireAuth, budgetHandler.ListBudgets)
		budgetGroup.GET("/list/services/:budgetID", middlewares.RequireAuth, budgetHandler.GetBudgetServices)
		budgetGroup.DELETE("/:budgetID", middlewares.RequireAuth, budgetHandler.DeleteBudget)
	}

	appointmentRepository := repositories.NewAppointmentRepository(db)
	appointmentService := services.NewAppointmentService(appointmentRepository)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)
	appointmentGroup := r.Group("/appointments")
	{
		appointmentGroup.POST("/", middlewares.RequireAuth, appointmentHandler.CreateAppointment)
		appointmentGroup.GET("/list/:userID", middlewares.RequireAuth, appointmentHandler.ListAppointments)
		appointmentGroup.PUT("/:userID", middlewares.RequireAuth, appointmentHandler.UpdateAppointment)
		appointmentGroup.DELETE("/:budgetID", middlewares.RequireAuth, appointmentHandler.DeleteAppointment)
	}

	r.Run()
}
