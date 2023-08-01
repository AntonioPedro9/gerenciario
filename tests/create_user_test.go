package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/database"
	"server/handlers"
	"server/models"
	"server/repositories"
	"server/services"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateUser(t *testing.T) {
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
	user := models.NewUser("John Doe", "john@example.com", "password")

	// Convert user data to JSON
	jsonUserData, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	// Create a POST request to create a user
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUserData))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()    // Create a recorder to record the response
	userHandler.CreateUser(recorder, req) // Call the CreateUser handler function

	// Check if the status code is what we expect
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, recorder.Code)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}