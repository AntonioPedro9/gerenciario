package main

import (
	"log"
	"net/http"

	"server/database"
	"server/handlers"
	"server/repositories"
	"server/services"

	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", userHandler.ListUsers).Methods(http.MethodGet)
	r.HandleFunc("/users", userHandler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
