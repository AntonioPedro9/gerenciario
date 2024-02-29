package handlers

import (
	"net/http"
	"server/models"
	"server/services"
	"server/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type BudgetHandler struct {
	budgetService *services.BudgetService
}

func NewBudgetHandler(budgetService *services.BudgetService) *BudgetHandler {
	return &BudgetHandler{budgetService}
}

/** 
 * Creates a new budget.
 * It accepts a JSON body with the budget details.
 * Returns 201 if the budget is created successfully.
 * Returns 400 if the request fails to bind to JSON.
 * Returns 500 for internal server errors.
**/
func (bh *BudgetHandler) CreateBudget(c *gin.Context) {
	var budget models.CreateBudgetRequest
	if err := c.ShouldBindJSON(&budget); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	if err := bh.budgetService.CreateBudget(&budget); err != nil {
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
 * Lists all budget for a user.
 * It extracts userID from JWT token.
 * Returns 200 along with a list of budget.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (bh *BudgetHandler) ListBudgets(c *gin.Context) {
	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	budgets, err := bh.budgetService.ListBudgets(userID)
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

	c.JSON(http.StatusOK, budgets)
}

/** 
 * Gets a budget by ID.
 * It requires budgetID as a path parameter and extracts userID from JWT token.
 * Returns 200 along with the budget details.
 * Returns 400 if the budget ID fails to parse.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (bh *BudgetHandler) GetBudget(c *gin.Context) {
	budgetID, err := utils.GetParamID("budgetID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	budget, err := bh.budgetService.GetBudget(budgetID, userID)
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

	c.JSON(http.StatusOK, budget)
}

/** 
 * Gets all jobs of a budget.
 * It requires budgetID as a path parameter and extracts userID from JWT token.
 * Returns 200 along with the budget jobs.
 * Returns 400 if the budget ID fails to parse.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (bh *BudgetHandler) GetBudgetJobs(c *gin.Context) {
	budgetID, err := utils.GetParamID("budgetID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	budgetServices, err := bh.budgetService.GetBudgetJobs(budgetID, userID)
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

	c.JSON(http.StatusOK, budgetServices)
}

/** 
 * Deletes a budget.
 * It requires budgetID as a path parameter and extracts userID from JWT token.
 * Returns 204 if the budget is deleted successfully.
 * Returns 400 if the budget ID fails to parse.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (bh *BudgetHandler) DeleteBudget(c *gin.Context) {
	budgetID, err := utils.GetParamID("budgetID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	if err := bh.budgetService.DeleteBudget(budgetID, userID); err != nil {
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
