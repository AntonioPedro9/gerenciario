package repositories

import (
	"server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BudgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) *BudgetRepository {
	return &BudgetRepository{db}
}

func (br *BudgetRepository) Create(budget *models.Budget) error {
	return br.db.Create(budget).Error
}

func (br *BudgetRepository) CreateBudgetService(budgetService *models.BudgetService) error {
	return br.db.Create(budgetService).Error
}

func (br *BudgetRepository) List(userID uuid.UUID) ([]models.Budget, error) {
	var budgets []models.Budget

	if err := br.db.Where("user_id = ?", userID).Preload("BudgetServices").Find(&budgets).Error; err != nil {
		return nil, err
	}

	return budgets, nil
}

func (br *BudgetRepository) GetBudgetById(id uint) (*models.Budget, error) {
	var budget models.Budget

	if err := br.db.Where("id = ?", id).First(&budget).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &budget, nil
}

func (br *BudgetRepository) Delete(budgetID uint) error {
	budget := models.Budget{ID: budgetID}
	return br.db.Delete(&budget).Error
}
