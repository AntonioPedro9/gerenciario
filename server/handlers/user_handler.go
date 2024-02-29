package handlers

import (
	"net/http"
	"server/models"
	"server/services"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService}
}

/**
 * CreateUser creates a new user.
 * It accepts a JSON body with the user details.
 * Returns 201 if the user is created successfully.
 * Returns 400 if the request fails to bind to JSON.
 * Returns 500 for internal server errors.
**/
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	if err := uh.userService.CreateUser(&user); err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

/**
 * ListUsers lists all users.
 * Returns 200 along with a list of users.
 * Returns 500 for internal server errors.
**/
func (uh *UserHandler) ListUsers(c *gin.Context) {
	users, err := uh.userService.ListUsers()
	if err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusOK, users)
}

/**
 * UpdateUser updates a user.
 * It accepts a JSON body with the user details.
 * Returns 200 if the user is updated successfully.
 * Returns 400 if the request fails to bind to JSON or no token is provided.
 * Returns 401 if the token is unauthorized.
 *Returns 500 for internal server errors.
**/
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var user models.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	updatedUser, err := uh.userService.UpdateUser(&user, userID)
	if err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

/**
 * DeleteUser deletes a user.
 * It requires a user ID as a path parameter.
 * Returns 204 if the user is deleted successfully.
 * Returns 400 if the user ID fails to parse or no token is provided.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
*/
func (uh *UserHandler) DeleteUser(c *gin.Context) {
	paramUserID := c.Param("id")

	userID, err := uuid.Parse(paramUserID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse user ID"})
		return
	}

	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := uh.userService.DeleteUser(userID, tokenID); err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

/**
 * LoginUser logs in a user.
 * It accepts a JSON body with the user login details.
 * Returns 200 along with a token if the user is logged in successfully.
 * Returns 400 if the request fails to bind to JSON.
 * Returns 500 for internal server errors.
*/
func (uh *UserHandler) LoginUser(c *gin.Context) {
	var loginUserRequest models.LoginUserResquest
	if err := c.ShouldBindJSON(&loginUserRequest); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	tokenString, err := uh.userService.LoginUser(&loginUserRequest)
	if err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
