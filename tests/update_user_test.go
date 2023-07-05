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
	db, err := database.CreateTestDatabaseConnection()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	cleanup := func() {
		_, err := db.Exec("DELETE FROM users")
		if err != nil {
			t.Fatal(err)
		}
	}
	defer cleanup()

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

	_, err = userService.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	editedUser := &models.User{
		ID:       user.ID,
		Name:     "John Smith",
		Email:    "john@example.com",
		Password: "newpassword",
	}

	editedUserJSON, err := json.Marshal(editedUser)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/users", bytes.NewBuffer(editedUserJSON))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	userHandler.UpdateUser(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, recorder.Code)
	}

	updatedUser, err := userService.GetUserById(user.ID)
	if err != nil {
		t.Fatal(err)
	}

	if updatedUser.Name != editedUser.Name {
		t.Errorf("Expected name %s but got %s", editedUser.Name, updatedUser.Name)
	}

	if updatedUser.Password != editedUser.Password {
		t.Errorf("Expected password %s but got %s", editedUser.Password, updatedUser.Password)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}