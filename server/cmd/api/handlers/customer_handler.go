package handlers

import (
	"net/http"
	"server/internals/models"
	services "server/internals/services"
	"server/pkg/errors"
	"server/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandlerInterface interface {
	CreateCustomer(c *gin.Context)
	GetCustomer(c *gin.Context)
	GetUserCustomers(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

type CustomerHandler struct {
	customerService services.CustomerServiceInterface
}

func NewCustomerHandler(customerService services.CustomerServiceInterface) *CustomerHandler {
	return &CustomerHandler{customerService}
}

func (ch *CustomerHandler) CreateCustomer(c *gin.Context) {
	var data models.CreateCustomerRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		errors.HandleError(c, errors.JsonBindingError)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := ch.customerService.CreateCustomer(&data, tokenUserId); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (ch *CustomerHandler) GetCustomer(c *gin.Context) {
	customerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	customer, err := ch.customerService.GetCustomer(uint(customerId), tokenUserId)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (ch *CustomerHandler) GetUserCustomers(c *gin.Context) {
	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	customers, err := ch.customerService.GetUserCustomers(tokenUserId)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, customers)
}

func (ch *CustomerHandler) UpdateCustomer(c *gin.Context) {
	customerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	var data models.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		errors.HandleError(c, errors.JsonBindingError)
		return
	}

	if err := ch.customerService.UpdateCustomer(uint(customerId), tokenUserId, &data); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (ch *CustomerHandler) DeleteCustomer(c *gin.Context) {
	customerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := ch.customerService.DeleteCustomer(uint(customerId), tokenUserId); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
