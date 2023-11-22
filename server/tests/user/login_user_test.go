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

func TestLoginUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// user model to create request
	userID, _ := utils.GenerateUUID()
	hashedPassword, _ := utils.HashPassword("password")
	user := &models.User{
		ID:       userID,
		Name:     "Jonh Doe",
		Email:    "jonhdoe@email.com",
		Password: hashedPassword,
	}
	userRepository.Create(user)

	loginUser := &models.LoginUserResquest{
		Email:    "jonhdoe@email.com",
		Password: "password",
	}

	r.POST("/users/login", userHandler.LoginUser)

	t.Run("Login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(loginUser)
		request, _ := http.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		// check if the token is set in the cookies
		cookies := response.Result().Cookies()

		var token string

		for _, cookie := range cookies {
			if cookie.Name == "Authorization" {
				token = cookie.Value
				break
			}
		}

		if token == "" {
			t.Fatal("Token not found in cookies")
		}
	})

	database.ClearTestDatabase(test_db)
}
