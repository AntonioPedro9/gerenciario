package tests

import (
	"log"
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

	createdUser, err := userService.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	url := "/users/" + strconv.Itoa(createdUser.ID)

	log.Println(url)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	userHandler.DeleteUser(recorder, req)

	if recorder.Code != http.StatusNoContent {
		t.Errorf("Expected status %d but got %d", http.StatusNoContent, recorder.Code)
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}
