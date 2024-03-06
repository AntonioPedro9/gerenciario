package repositories

import (
	"server/models"
	"server/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db}
}

func (jr *JobRepository) Create(job *models.Job) error {
	return jr.db.Create(job).Error
}

func (jr *JobRepository) List(userID uuid.UUID) ([]models.Job, error) {
	var jobs []models.Job

	if err := jr.db.Where("user_id = ?", userID).Order("name").Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (jr *JobRepository) GetJobById(id uint) (*models.Job, error) {
	var job models.Job

	if err := jr.db.Where("id = ?", id).First(&job).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NotFoundError
		}
		return nil, err
	}

	return &job, nil
}

func (jr *JobRepository) Update(job *models.UpdateJobRequest) error {
	err := jr.db.Model(&models.Job{}).Where("id = ?", job.ID).Updates(job).Error
	if err != nil {
		return err
	}
	return nil
}

func (jr *JobRepository) Delete(jobID uint) error {
	job := models.Job{ID: jobID}
	return jr.db.Delete(&job).Error
}
