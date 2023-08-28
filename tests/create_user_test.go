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

func TestCreateUser(t *testing.T) {
	cleanup := func() {
		initializers.TestDB.Exec("DELETE FROM users")
	}
	defer cleanup()

	userRepository := repositories.NewUserRepository(initializers.TestDB)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	user := models.NewUser("John Doe", "john@example.com", "password")

	jsonUserData, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUserData))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	userHandler.CreateUser(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, recorder.Code)
	}
}
