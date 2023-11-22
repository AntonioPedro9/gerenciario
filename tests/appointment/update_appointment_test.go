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
	"time"

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

func TestUpdateAppointment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	appointmentRepository := repositories.NewAppointmentRepository(test_db)
	appointmentService := services.NewAppointmentService(appointmentRepository)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)

	// create user that will update the appointment
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

	// create appointment that will be updated
	appointment := &models.Appointment{
		Date:     time.Now(),
		BudgetID: 1,
		UserID:   userID,
	}
	appointmentRepository.Create(appointment)

	// appointment model to update request
	date := time.Now()

	updateAppointment := &models.UpdateAppointmentRequest{
		ID:     1,
		Date: &date,
		UserID: userID,
	}

	r.PUT("/appointments", appointmentHandler.UpdateAppointment)

	t.Run("Update appointment", func(t *testing.T) {
		requestBody, _ := json.Marshal(updateAppointment)
		request, _ := http.NewRequest(http.MethodPut, "/appointments", bytes.NewBuffer(requestBody))
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
