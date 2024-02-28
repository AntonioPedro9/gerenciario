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

func (br *BudgetRepository) CreateBudgetJob(budgetJob *models.BudgetJob) error {
	return br.db.Create(budgetJob).Error
}

func (br *BudgetRepository) List(userID uuid.UUID) ([]models.Budget, error) {
	var budgets []models.Budget

	if err := br.db.Where("user_id = ?", userID).Find(&budgets).Error; err != nil {
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

func (br *BudgetRepository) GetBudgetJobs(budgetID uint) ([]models.Job, error) {
	var budgetJobs []models.BudgetJob
	var jobs []models.Job

	if err := br.db.Where("budget_id = ?", budgetID).Find(&budgetJobs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	for _, budgetJob := range budgetJobs {
		var job models.Job
		
		if err := br.db.Where("id = ?", budgetJob.JobID).First(&job).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (br *BudgetRepository) Delete(budgetID uint) error {
	budget := models.Budget{ID: budgetID}
	return br.db.Delete(&budget).Error
}
