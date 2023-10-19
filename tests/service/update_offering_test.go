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

func TestUpdateService(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	serviceRepository := repositories.NewServiceRepository(test_db)
	serviceService := services.NewServiceService(serviceRepository)
	serviceHandler := handlers.NewServiceHandler(serviceService)

	// create user that will update the service
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

	// create service that will updated
	service := &models.Service{
		Name:        "Service",
		Description: "Service description",
		Duration:    1,
		Price:       50,
		UserID:      userID,
	}
	serviceRepository.Create(service)

	// service model to update request
	updateService := &models.UpdateServiceRequest{
		ID:          1,
		Name:        "New Service",
		Description: "New service description",
		Duration:    2,
		Price:       100,
		UserID:      userID,
	}

	r.PUT("/services", serviceHandler.UpdateService)

	t.Run("Update service", func(t *testing.T) {
		requestBody, _ := json.Marshal(updateService)
		request, _ := http.NewRequest(http.MethodPut, "/services", bytes.NewBuffer(requestBody))
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
