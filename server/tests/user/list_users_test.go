package tests

import (
	"net/http"
	"net/http/httptest"
	"server/database"
	"server/handlers"
	"server/repositories"
	"server/services"
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

func TestListUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	r.GET("/users", userHandler.ListUsers)

	t.Run("List users", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users", nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		expectedStatus := http.StatusOK

		if response.Code != expectedStatus {
			t.Errorf("Expected status %d but got %d", expectedStatus, response.Code)
		}
	})

	database.ClearTestDatabase(test_db)
}
