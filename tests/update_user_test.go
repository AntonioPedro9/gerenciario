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

func TestUpdateUser(t *testing.T) {
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

	// Create user to update
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

	// Convert user data to JSON
	jsonUserData, err := json.Marshal(userData)
	if err != nil {
		t.Fatal(err)
	}

	// Create a PUT request to update the user
	req, err := http.NewRequest("PUT", "/users", bytes.NewBuffer(jsonUserData))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()    // Create a ResponseRecorder
	userHandler.UpdateUser(recorder, req) // Call the UpdateUser handler function

	// Check the status code is what we expect
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, recorder.Code)
	}

	updatedUser, err := userService.GetUserById(tempUser.ID)
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
