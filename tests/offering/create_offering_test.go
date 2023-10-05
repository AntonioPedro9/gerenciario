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

func TestCreateOffering(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	offeringRepository := repositories.NewOfferingRepository(test_db)
	offeringService := services.NewOfferingService(offeringRepository)
	offeringHandler := handlers.NewOfferingHandler(offeringService)

	// create user that will create the offering
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

	// offering model to create request
	offering := &models.CreateOfferingRequest{
		Name:        "Offering",
		Description: "Offering description",
		Duration:    1,
		UserID:      userID,
	}

	r.POST("/offerings", offeringHandler.CreateOffering)

	t.Run("Create offering", func(t *testing.T) {
		requestBody, _ := json.Marshal(offering)
		request, _ := http.NewRequest(http.MethodPost, "/offerings", bytes.NewBuffer(requestBody))
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
