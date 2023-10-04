package middlewares_test

import (
	"bytes"
	"encoding/json"
	"log"
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
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestCreateClient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	clientRepository := repositories.NewClientRepository(test_db)
	clientService := services.NewClientService(clientRepository)
	clientHandler := handlers.NewClientHandler(clientService)

	// create user that will create the client
	userID, _ := utils.GenerateUUID()
	user := &models.User{
		ID:       userID,
		Name:     "Jonh Doe",
		Email:    "jonhdoe@email.com",
		Password: "password",
	}
	userRepository.Create(user)

	// generate jwt token to authorize action
	tokenString, _ := utils.GenerateToken(userID)

	// client model to create request
	client := &models.CreateClientRequest{
		CPF:    "12345678910",
		Name:   "Jonh Doe",
		Email:  "jonhdoe@email.com",
		Phone:  "(11) 9 2233-4455",
		UserID: userID,
	}

	r.POST("/clients", clientHandler.CreateClient)

	t.Run("Create client", func(t *testing.T) {
		requestBody, _ := json.Marshal(client)
		request, _ := http.NewRequest(http.MethodPost, "/clients", bytes.NewBuffer(requestBody))
		request.AddCookie(&http.Cookie{Name: "Authorization", Value: tokenString})

		w := httptest.NewRecorder()
		r.ServeHTTP(w, request)

		expectedStatus := http.StatusCreated

		if w.Code != expectedStatus {
			t.Errorf("Expected status %d but got %d", expectedStatus, w.Code)
		}
	})

	database.ClearTestDatabase(test_db)
}
