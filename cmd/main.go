package main

import (
	"net/http"
	"os"
	"server/handlers"
	"server/initializers"
	"server/repositories"
	"server/services"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Configure logger
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "02/01/2006 15:04",
	})

	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.CreateDatabaseTables()
}

func main() {
	userRepository := repositories.NewUserRepository(initializers.DB)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router := mux.NewRouter()

	router.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", userHandler.ListUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", userHandler.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/users/login", userHandler.Login).Methods(http.MethodPost)

	port := os.Getenv("PORT")

	log.Info("Server started on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
