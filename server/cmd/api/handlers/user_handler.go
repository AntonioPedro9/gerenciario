package handlers

import (
	"net/http"
	"server/internals/models"
	services "server/internals/services"
	"server/pkg/errors"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUserData(c *gin.Context)
	UpdateUserPassword(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserHandler struct {
	userService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) *UserHandler {
	return &UserHandler{userService}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var data models.CreateUserRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		errors.HandleError(c, errors.JsonBindingError)
		return
	}

	if err := uh.userService.CreateUser(&data); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	user, err := uh.userService.GetUser(tokenUserId)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) UpdateUserData(c *gin.Context) {
	var data models.UpdateUserDataRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		errors.HandleError(c, errors.JsonBindingError)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := uh.userService.UpdateUserData(&data, tokenUserId); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (uh *UserHandler) UpdateUserPassword(c *gin.Context) {
	var data models.UpdateUserPasswordRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		errors.HandleError(c, errors.JsonBindingError)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := uh.userService.UpdateUserPassword(&data, tokenUserId); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := uh.userService.DeleteUser(tokenUserId); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
