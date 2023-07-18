package tests

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	_ "github.com/lib/pq"

	"server/database"
	"server/handlers"
	"server/models"
	"server/repositories"
	"server/services"
)

func TestDeleteUser(t *testing.T) {
	// Create a test database connection
	db, err := database.CreateTestDatabaseConnection()
	if err != nil {
		t.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	// Create a cleanup function to delete all users after each test
	cleanup := func() {
		_, err := db.Exec("DELETE FROM users")
		if err != nil {
			t.Fatal(err)
		}
	}
	defer cleanup()

	// Create a transaction to run the test inside
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	user := &models.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}

	// Create user to delete
	createdUser, err := userService.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	url := "/users/" + strconv.Itoa(createdUser.ID)

	// Create a DELETE request to delete the user
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()    // Create a recorder to record the response
	userHandler.DeleteUser(recorder, req) // Call the DeleteUser handler function

	// Check if the status code is what we expect
	if recorder.Code != http.StatusNoContent {
		t.Errorf("Expected status %d but got %d", http.StatusNoContent, recorder.Code)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}
