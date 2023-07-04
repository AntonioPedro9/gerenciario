package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"server/models"
	"server/services"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println("Failed to decode user data:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.CreateUser(user); err != nil {
		log.Println("Failed to create user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uh *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.ListUsers()
	if err != nil {
		log.Println("Failed to fetch users:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println("Failed to decode user data:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.UpdateUser(user); err != nil {
		log.Println("Failed to update user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		log.Println("User ID not provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid user ID:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.DeleteUser(convertedId); err != nil {
		log.Println("Failed to delete user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
