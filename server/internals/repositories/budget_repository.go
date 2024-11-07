package repositories

import (
	"server/internals/models"
	"server/pkg/errors"
	"server/pkg/logs"

	"gorm.io/gorm"
)

type BudgetRepositoryInterface interface {
	Create(data *models.CreateBudgetRequest, tokenUserId uint) error
	GetById(budgetId, tokenUserId uint) (*models.GetBudgetResponse, error)
	GetUserBudgets(tokenUserId uint) ([]models.GetBudgetResponse, error)
	Delete(budgetId, tokenUserId uint) error
}

type BudgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) *BudgetRepository {
	return &BudgetRepository{db}
}

func (br *BudgetRepository) Create(data *models.CreateBudgetRequest, tokenUserId uint) error {
	budget := &models.Budget{
		UserId:         tokenUserId,
		CustomerId:     data.CustomerId,
		TotalPrice:     data.TotalPrice,
		Discount:       data.Discount,
		Date:           data.Date,
		ExpirationDate: data.ExpirationDate,
		Jobs:           data.Jobs,
	}

	result := br.db.Create(budget)

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	return nil
}

func (br *BudgetRepository) GetById(budgetId, tokenUserId uint) (*models.GetBudgetResponse, error) {
	var budget models.Budget

	result := br.db.Preload("Jobs").Where("id = ? AND user_id = ?", budgetId, tokenUserId).First(&budget)

	if result.RowsAffected == 0 {
		return nil, errors.NotFoundError
	}

	if result.Error != nil {
		logs.LogError(result.Error)
		return nil, result.Error
	}

	return &models.GetBudgetResponse{
		Id:             budget.Id,
		UserId:         budget.UserId,
		CustomerId:     budget.CustomerId,
		TotalPrice:     budget.TotalPrice,
		Discount:       budget.Discount,
		Date:           budget.Date,
		ExpirationDate: budget.ExpirationDate,
		Jobs:           budget.Jobs,
	}, nil
}

func (br *BudgetRepository) GetUserBudgets(tokenUserId uint) ([]models.GetBudgetResponse, error) {
	var budgets []models.Budget
	var budgetResponses []models.GetBudgetResponse

	result := br.db.Preload("Jobs").Where("user_id = ?", tokenUserId).Find(&budgets)

	if result.Error != nil {
		logs.LogError(result.Error)
		return nil, result.Error
	}

	for _, budget := range budgets {
		budgetResponses = append(budgetResponses, models.GetBudgetResponse{
			Id:             budget.Id,
			UserId:         budget.UserId,
			CustomerId:     budget.CustomerId,
			TotalPrice:     budget.TotalPrice,
			Discount:       budget.Discount,
			Date:           budget.Date,
			ExpirationDate: budget.ExpirationDate,
			Jobs:           budget.Jobs,
		})
	}

	return budgetResponses, nil
}

func (br *BudgetRepository) Delete(budgetId, tokenUserId uint) error {
	result := br.db.Where("id = ? AND user_id = ?", budgetId, tokenUserId).Delete(&models.Budget{})

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.NotFoundError
	}

	return nil
}
