package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"server/handlers"
	"server/initializers"
	"server/models"
	"server/repositories"
	"server/services"
)

func init() {
	initializers.ConnectToTestDatabase()
}

func TestLoginAndCreateUser(t *testing.T) {
	cleanup := func() {
		initializers.TestDB.Exec("DELETE FROM users")
	}
	defer cleanup()

	userRepository := repositories.NewUserRepository(initializers.TestDB)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	user := models.NewUser("John Doe", "john@example.com", "password")
	userRepository.Create(user)

	loginData := map[string]string{
		"email":    "john@example.com",
		"password": "password",
	}

	loginJSON, err := json.Marshal(loginData)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/users/login", bytes.NewBuffer(loginJSON))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	userHandler.Login(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, recorder.Code)
	}

	var response map[string]string

	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	_, tokenExists := response["token"]
	if !tokenExists {
		t.Errorf("Token not found in the response")
	}
}
