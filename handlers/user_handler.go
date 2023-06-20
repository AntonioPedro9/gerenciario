package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"server/models"
	"server/services"
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
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.CreateUser(user); err != nil {
		log.Println("Error creating user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uh *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.ListUsers()

	if err != nil {
		log.Println("Error fetching users:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.UpdateUser(user); err != nil {
		log.Println("Error updating user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		log.Println("Error getting user ID:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.DeleteUser(id); err != nil {
		log.Println("Error deleting user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
