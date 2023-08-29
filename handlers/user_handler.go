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
	userData := &models.CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(userData); err != nil {
		log.Error("Failed to decode user data:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser := models.NewUser(userData.Name, userData.Email, userData.Password)

	if _, err := uh.userService.CreateUser(newUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uh *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Error("Failed to encode users:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userData := &models.UpdateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(userData); err != nil {
		log.Error("Failed to decode user data:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.userService.UpdateUser(userData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	id := urlParts[len(urlParts)-1]

	if err := uh.userService.DeleteUser(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginData := &models.LoginUserResquest{}

	if err := json.NewDecoder(r.Body).Decode(loginData); err != nil {
		log.Error("Failed to decode login data:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := uh.userService.Login(loginData.Email, loginData.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	response := map[string]string{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error("Failed to encode response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
