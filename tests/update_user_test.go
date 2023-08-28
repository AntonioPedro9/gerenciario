package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/handlers"
	"server/initializers"
	"server/models"
	"server/repositories"
	"server/services"
	"testing"
)

func init() {
	initializers.ConnectToTestDatabase()
}

func TestUpdateUser(t *testing.T) {
	cleanup := func() {
		initializers.TestDB.Exec("DELETE FROM users")
	}
	defer cleanup()
	
	userRepository := repositories.NewUserRepository(initializers.TestDB)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	user := models.NewUser("John Doe", "john@example.com", "password")

	tempUser, err := userService.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	userData := &models.UpdateUserRequest{
		ID:       tempUser.ID,
		Name:     "John Smith",
		Email:    "johnsmith@example.com",
		Password: "newpassword",
	}

	jsonUserData, err := json.Marshal(userData)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/users", bytes.NewBuffer(jsonUserData))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	userHandler.UpdateUser(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, recorder.Code)
	}
}
