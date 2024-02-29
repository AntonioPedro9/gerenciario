package handlers

import (
	"net/http"
	"server/models"
	"server/services"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	customerService *services.CustomerService
}

func NewCustomerHandler(customerService *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService}
}

func (ch *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer models.CreateCustomerRequest
	if err := c.ShouldBindJSON(&customer); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	if err := ch.customerService.CreateCustomer(&customer); err != nil {
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

func (ch *CustomerHandler) ListCustomers(c *gin.Context) {
	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	users, err := ch.customerService.ListCustomers(userID)
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

func (ch *CustomerHandler) GetCustomer(c *gin.Context) {
	paramCustomerID := c.Param("customerID")

	parsedCustomerID, err := strconv.ParseUint(paramCustomerID, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	customerID := uint(parsedCustomerID)

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	customer, err := ch.customerService.GetCustomer(customerID, userID)
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

	c.JSON(http.StatusOK, customer)
}

func (ch *CustomerHandler) UpdateCustomer(c *gin.Context) {
	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var customer models.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&customer); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	updatedCustomer, err := ch.customerService.UpdateCustomer(&customer, userID)
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

	c.JSON(http.StatusOK, updatedCustomer)
}

func (ch *CustomerHandler) DeleteCustomer(c *gin.Context) {
	paramCustomerID := c.Param("customerID")

	parsedID, err := strconv.ParseUint(paramCustomerID, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	customerID := uint(parsedID)

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := ch.customerService.DeleteCustomer(customerID, userID); err != nil {
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
