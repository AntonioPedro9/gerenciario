package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"

	"server/database"
	"server/handlers"
	"server/models"
	"server/repositories"
	"server/services"
)

func TestListUsers(t *testing.T) {
	// Create a test database connection
	db, err := database.CreateTestDatabaseConnection()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
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

	user1 := &models.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}

	user2 := &models.User{
		Name:     "Jane Doe",
		Email:    "jane@example.com",
		Password: "password",
	}

	// Create users to list
	_, err = userService.CreateUser(user1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = userService.CreateUser(user2)
	if err != nil {
		t.Fatal(err)
	}

	// Create a GET request to list users
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()   // Create a recorder to record the response
	userHandler.ListUsers(recorder, req) // Call the ListUsers handler function

	// Check the status code is what we expect
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, recorder.Code)
	}

	var users []models.User

	// Convert the response body to a slice of users
	err = json.Unmarshal(recorder.Body.Bytes(), &users)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the number of users is what we expect
	if len(users) != 2 {
		t.Errorf("Expected %d users but got %d", 2, len(users))
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}
