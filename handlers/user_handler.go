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
		log.Println("Failed to decode user data: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.CreateUser(user); err != nil {
		log.Println("Failed to create user: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("User created: ", user.ID)
	w.WriteHeader(http.StatusCreated)
}

func (uh *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.ListUsers()
	if err != nil {
		log.Println("Failed to fetch users: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Fetched", len(users), "users")
	json.NewEncoder(w).Encode(users)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println("Failed to decode user data: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.UpdateUser(user); err != nil {
		log.Println("Failed to update user: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("User updated: ", user.ID)
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Invalid user ID: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.DeleteUser(id); err != nil {
		log.Println("Failed to delete user: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("User deleted: ", id)
	w.WriteHeader(http.StatusNoContent)
}
