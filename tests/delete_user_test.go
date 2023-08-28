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

func TestDeleteUser(t *testing.T) {
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

	url := "/users/" + tempUser.ID

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	userHandler.DeleteUser(recorder, req)

	if recorder.Code != http.StatusNoContent {
		t.Errorf("Expected status %d but got %d", http.StatusNoContent, recorder.Code)
	}
}
