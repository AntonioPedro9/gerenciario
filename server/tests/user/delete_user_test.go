package tests

import (
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

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// setup layers
	test_db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(test_db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// create a user to be deleted
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

	r.DELETE("/users/:id", userHandler.DeleteUser)

	t.Run("Delete user", func(t *testing.T) {
		requestEndPoint := "/users/" + userID.String()
		request, _ := http.NewRequest(http.MethodDelete, requestEndPoint, nil)
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
