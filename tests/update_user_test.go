package tests

import (
	"bytes"
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

func TestUpdateUser(t *testing.T) {
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

	user := &models.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password",
	}

	// Create user to update
	_, err = userService.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	userData := &models.User{
		ID:       user.ID,
		Name:     "John Smith",
		Email:    "john@example.com",
		Password: "newpassword",
	}

	// Convert user data to JSON
	userDataJSON, err := json.Marshal(userData)
	if err != nil {
		t.Fatal(err)
	}

	// Create a PUT request to update the user
	req, err := http.NewRequest("PUT", "/users", bytes.NewBuffer(userDataJSON))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()    // Create a ResponseRecorder
	userHandler.UpdateUser(recorder, req) // Call the UpdateUser handler function

	// Check the status code is what we expect
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, recorder.Code)
	}

	updatedUser, err := userService.GetUserById(user.ID)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the name was updated
	if updatedUser.Name != userData.Name {
		t.Errorf("Expected name %s but got %s", userData.Name, updatedUser.Name)
	}

	// Check if the email was updated
	if updatedUser.Password != userData.Password {
		t.Errorf("Expected password %s but got %s", userData.Password, updatedUser.Password)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}
