package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/services"
	"strings"

	log "github.com/sirupsen/logrus"
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
		log.Error("Failed to decode user data:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := uh.userService.CreateUser(user); err != nil {
		log.Error("Failed to create user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uh *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.ListUsers()
	if err != nil {
		log.Error("Failed to fetch users:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Error("Failed to decode user data:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.UpdateUser(user); err != nil {
		log.Error("Failed to update user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	id := urlParts[len(urlParts)-1]

	if err := uh.userService.DeleteUser(id); err != nil {
		log.Error("Failed to delete user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
