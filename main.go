package main

import (
	"net/http"
	"os"
	"server/database"
	"server/handlers"
	"server/repositories"
	"server/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Create database connection
	db, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	// Create database tables
	err = database.CreateDatabaseTables(db)
	if err != nil {
		log.Fatal("Error creating database tables: ", err)
	}

	// Load environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router := mux.NewRouter()

	router.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", userHandler.ListUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", userHandler.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)

	log.Info("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}
