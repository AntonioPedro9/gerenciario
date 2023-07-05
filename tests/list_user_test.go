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

	_, err = userService.CreateUser(user1)
	if err != nil {
		t.Fatal(err)
	}

	_, err = userService.CreateUser(user2)
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

	var users []models.User
	
	err = json.Unmarshal(recorder.Body.Bytes(), &users)
	if err != nil {
		t.Fatal(err)
	}

	if len(users) != 2 {
		t.Errorf("Expected %d users but got %d", 2, len(users))
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}
