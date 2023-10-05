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

func TestListOfferings(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	offeringRepository := repositories.NewOfferingRepository(test_db)
	offeringService := services.NewOfferingService(offeringRepository)
	offeringHandler := handlers.NewOfferingHandler(offeringService)

	// create user that will list the offerings
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

	r.GET("/offerings/:userID", offeringHandler.ListOfferings)

	t.Run("List offerings", func(t *testing.T) {
		requestEndPoint := "/offerings/" + userID.String()
		request, _ := http.NewRequest(http.MethodGet, requestEndPoint, nil)
		request.AddCookie(&http.Cookie{Name: "Authorization", Value: tokenString})

		w := httptest.NewRecorder()
		r.ServeHTTP(w, request)

		expectedStatus := http.StatusOK

		if w.Code != expectedStatus {
			t.Errorf("Expected status %d but got %d", expectedStatus, w.Code)
		}
	})

	database.ClearTestDatabase(test_db)
}
