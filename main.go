package main

import (
	"net/http"

	"server/database"
	"server/handlers"
	"server/repositories"
	"server/services"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	err = database.CreateDatabaseTables(db)
	if err != nil {
		log.Fatal("Error creating database tables: ", err)
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
	log.Fatal(http.ListenAndServe(":8080", router))
}
