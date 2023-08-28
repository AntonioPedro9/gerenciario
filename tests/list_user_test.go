package tests

import (
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

func TestListUsers(t *testing.T) {
	cleanup := func() {
		initializers.TestDB.Exec("DELETE FROM users")
	}
	defer cleanup()

	userRepository := repositories.NewUserRepository(initializers.TestDB)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	user := models.NewUser("John Doe", "john@example.com", "password")

	_, err := userService.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	userHandler.ListUsers(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, recorder.Code)
	}
}
