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

	offeringRepository := repositories.NewOfferingRepository(db)
	offeringService := services.NewOfferingService(offeringRepository)
	offeringHandler := handlers.NewOfferingHandler(offeringService)
	offeringGroup := r.Group("/offerings")
	{
		offeringGroup.POST("/", middlewares.RequireAuth, offeringHandler.CreateOffering)
		offeringGroup.GET("/:userID", middlewares.RequireAuth, offeringHandler.ListOfferings)
		offeringGroup.PUT("/", middlewares.RequireAuth, offeringHandler.UpdateOffering)
		offeringGroup.DELETE("/:offeringID", middlewares.RequireAuth, offeringHandler.DeleteOffering)
	}

	r.Run()
}
