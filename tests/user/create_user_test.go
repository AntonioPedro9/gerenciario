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
	"server/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// user model to create request
	userID, _ := utils.GenerateUUID()
	user := &models.User{
		ID:       userID,
		Name:     "Jonh Doe",
		Email:    "jonhdoe@email.com",
		Password: "password",
	}

	r.POST("/users", userHandler.CreateUser)

	t.Run("Create user", func(t *testing.T) {
		requestBody, _ := json.Marshal(user)
		request, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		expectedStatus := http.StatusCreated

		if response.Code != expectedStatus {
			t.Errorf("Expected status %d but got %d", expectedStatus, response.Code)
		}
	})

	database.ClearTestDatabase(test_db)
}
