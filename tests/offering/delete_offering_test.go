package middlewares_test

import (
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

func TestDeleteOffering(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	offeringRepository := repositories.NewOfferingRepository(test_db)
	offeringService := services.NewOfferingService(offeringRepository)
	offeringHandler := handlers.NewOfferingHandler(offeringService)

	// create user that will delete the offering
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

	// create offering that will deleted
	offering := &models.Offering{
		Name:        "Offering",
		Description: "Offering description",
		Duration:    1,
		UserID:      userID,
	}
	offeringRepository.Create(offering)

	r.DELETE("/offerings/:offeringID", offeringHandler.DeleteOffering)

	t.Run("Delete offering", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodDelete, "/offerings/1", nil)
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
