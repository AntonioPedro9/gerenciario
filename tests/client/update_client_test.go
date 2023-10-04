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

func TestUpdateClient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	clientRepository := repositories.NewClientRepository(test_db)
	clientService := services.NewClientService(clientRepository)
	clientHandler := handlers.NewClientHandler(clientService)

	// create user that will update the client
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

	// create client that will updated
	client := &models.Client{
		CPF:    "12345678910",
		Name:   "Jonh Doe",
		Email:  "jonhdoe@email.com",
		Phone:  "(11) 9 2233-4455",
		UserID: userID,
	}
	clientRepository.Create(client)

	// client model to update request
	updateClient := &models.UpdateClientRequest{
		ID:     1,
		CPF:    "12345678910",
		Name:   "Jonh Doe",
		Email:  "jonhdoe@email.com",
		Phone:  "(11) 9 2233-4455",
		UserID: userID,
	}

	r.PUT("/clients", clientHandler.UpdateClient)

	t.Run("Update client", func(t *testing.T) {
		requestBody, _ := json.Marshal(updateClient)
		request, _ := http.NewRequest(http.MethodPut, "/clients", bytes.NewBuffer(requestBody))
		request.AddCookie(&http.Cookie{Name: "Authorization", Value: tokenString})

		w := httptest.NewRecorder()
		r.ServeHTTP(w, request)

		expectedStatus := http.StatusNoContent

		if w.Code != expectedStatus {
			t.Errorf("Expected status %d but got %d", expectedStatus, w.Code)
		}
	})

	database.ClearTestDatabase(test_db)
}
