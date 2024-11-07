package services

import (
	"server/internals/models"
	"server/internals/repositories"
)

type BudgetServiceInterface interface {
	CreateBudget(data *models.CreateBudgetRequest, tokenUserId uint) error
	GetBudget(budgetId, tokenUserId uint) (models.GetBudgetResponse, error)
	GetUserBudgets(tokenUserId uint) (models.GetUserBudgetsResponse, error)
	DeleteBudget(budgetId, tokenUserId uint) error
}

type BudgetService struct {
	budgetRepository repositories.BudgetRepositoryInterface
}

func NewBudgetService(budgetRepository repositories.BudgetRepositoryInterface) *BudgetService {
	return &BudgetService{budgetRepository}
}

func (bs *BudgetService) CreateBudget(data *models.CreateBudgetRequest, tokenUserId uint) error {
	if err := validate.Struct(data); err != nil {
		return err
	}
	return bs.budgetRepository.Create(data, tokenUserId)
}

func (bs *BudgetService) GetBudget(budgetId, tokenUserId uint) (models.GetBudgetResponse, error) {
	var emptyBudget models.GetBudgetResponse

	budget, err := bs.budgetRepository.GetById(budgetId, tokenUserId)
	if err != nil {
		return emptyBudget, err
	}

	return *budget, nil
}

func (bs *BudgetService) GetUserBudgets(tokenUserId uint) (models.GetUserBudgetsResponse, error) {
	var emptyBudgets models.GetUserBudgetsResponse

	budgets, err := bs.budgetRepository.GetUserBudgets(tokenUserId)
	if err != nil {
		return emptyBudgets, err
	}

	budgetsArray := models.GetUserBudgetsResponse{
		Budgets: budgets,
	}

	return budgetsArray, nil
}

func (bs *BudgetService) DeleteBudget(budgetId, tokenUserId uint) error {
	return bs.budgetRepository.Delete(budgetId, tokenUserId)
}
