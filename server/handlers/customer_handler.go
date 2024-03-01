package handlers

import (
	"net/http"
	"server/models"
	"server/services"
	"server/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	customerService *services.CustomerService
}

func NewCustomerHandler(customerService *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService}
}

/** 
 * Creates a new customer.
 * It accepts a JSON body with the customer details.
 * Returns 201 if the customer is created successfully.
 * Returns 400 if the request fails to bind to JSON.
 * Returns 401 if token userID does not match request userID
 * Returns 500 for internal server errors.
**/
func (ch *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer models.CreateCustomerRequest
	if err := c.ShouldBindJSON(&customer); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	if err := ch.customerService.CreateCustomer(&customer, tokenID); err != nil {
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
 * Lists all customers for a user.
 * It extracts userID from JWT token.
 * Returns 200 along with a list of customers.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (ch *CustomerHandler) ListCustomers(c *gin.Context) {
	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	users, err := ch.customerService.ListCustomers(tokenID)
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
 * Gets a customer by ID.
 * It requires customerID as a path parameter and extracts userID from JWT token.
 * Returns 200 along with the customer details.
 * Returns 400 if the customer ID fails to parse.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (ch *CustomerHandler) GetCustomer(c *gin.Context) {
	customerID, err := utils.GetParamID("customerID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	customer, err := ch.customerService.GetCustomer(customerID, tokenID)
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

/** 
 * Updates a customer.
 * It accepts a JSON body with the customer details and extracts userID from JWT token.
 * Returns 200 if the customer is updated successfully.
 * Returns 400 if the request fails to bind to JSON.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (ch *CustomerHandler) UpdateCustomer(c *gin.Context) {
	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	var customer models.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&customer); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	updatedCustomer, err := ch.customerService.UpdateCustomer(&customer, tokenID)
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

/** 
 * Deletes a customer.
 * It requires customerID as a path parameter and extracts userID from JWT token.
 * Returns 204 if the customer is deleted successfully.
 * Returns 400 if the customer ID fails to parse.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (ch *CustomerHandler) DeleteCustomer(c *gin.Context) {
	customerID, err := utils.GetParamID("customerID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	if err := ch.customerService.DeleteCustomer(customerID, tokenID); err != nil {
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
