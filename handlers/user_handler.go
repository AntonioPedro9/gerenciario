package handlers

import (
	"net/http"
	"server/models"
	"server/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService}
}

// CreateUser godoc
//	@Summary		Create a new user
//	@Description	Create a new user with the provided data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.CreateUserRequest	true	"User data to create"
//	@Success		201		{}			null
//	@Failure		400		{object}	gin.H	"Failed to bind JSON request"
//	@Failure		500		{object}	gin.H	"Internal Server Error"
//	@Router			/users [post]
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	if err := uh.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// ListUsers godoc
//	@Summary		List users
//	@Description	Get a list of all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.User
//	@Failure		500	{object}	gin.H	"Internal Server Error"
//	@Router			/users [get]
func (uh *UserHandler) ListUsers(c *gin.Context) {
	users, err := uh.userService.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
//	@Summary		Update a user
//	@Description	Update an existing user with the provided data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UpdateUserRequest	true	"User data to update"
//	@Success		204		{}			null
//	@Failure		400		{object}	gin.H	"Failed to bind JSON request"
//	@Failure		500		{object}	gin.H	"Internal Server Error"
//	@Router			/users [put]
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var user models.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	if err := uh.userService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// DeleteUser godoc
//	@Summary		Delete a user
//	@Description	Delete a user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		204	{}			null
//	@Failure		400	{object}	gin.H	"Invalid ID"
//	@Failure		500	{object}	gin.H	"Internal Server Error"
//	@Router			/users/{id} [delete]
func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := uh.userService.DeleteUser(parsedID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// LoginUser godoc
//	@Summary		Login as a user
//	@Description	Login with user credentials and receive an authentication token
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.LoginUserResquest	true	"User login data"
//	@Success		200		{string}	string						"token"
//	@Failure		400		{object}	gin.H						"Failed to bind JSON request"
//	@Failure		500		{object}	gin.H						"Internal Server Error"
//	@Router			/users/login [post]
func (uh *UserHandler) LoginUser(c *gin.Context) {
	var loginUserRequest models.LoginUserResquest
	if err := c.ShouldBindJSON(&loginUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	tokenString, err := uh.userService.LoginUser(&loginUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
