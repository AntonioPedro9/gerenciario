package middlewares_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"server/database"
	"server/middlewares"
	"server/models"
	"server/repositories"
	"server/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestRequireAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(middlewares.RequireAuth)

	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	userID, _ := utils.GenerateUUID()

	data := &models.User{
		ID:       userID,
		Name:     "Jonh Doe",
		Email:    "jonhdoe@email.com",
		Password: "password",
	}

	userRepository.Create(data)

	tokenString, _ := utils.GenerateToken(userID)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Authorized")
	})

	t.Run("Authorized", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		request.AddCookie(&http.Cookie{Name: "Authorization", Value: tokenString})
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		expectedStatus := http.StatusOK

		if response.Code != expectedStatus {
			t.Errorf("Expected status %d, got %d", expectedStatus, response.Code)
		}
	})

	t.Run("Unauthorized", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)

		expectedStatus := http.StatusForbidden

		if response.Code != expectedStatus {
			t.Errorf("Expected status %d, got %d", expectedStatus, response.Code)
		}
	})

	database.ClearTestDatabase(test_db)
}
