package repositories

import (
	"server/models"

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

	if err := jr.db.Where("user_id = ?", userID).Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (jr *JobRepository) GetJobById(id uint) (*models.Job, error) {
	var job models.Job

	if err := jr.db.Where("id = ?", id).First(&job).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &job, nil
}

func (jr *JobRepository) Update(job *models.UpdateJobRequest) (*models.Job, error) {
	updateData := make(map[string]interface{})

	if job.Name != nil {
		updateData["name"] = *job.Name
	}

	if job.Description != nil {
		updateData["description"] = *job.Description
	}

	if job.Duration != nil {
		updateData["duration"] = *job.Duration
	}

	if job.Price != nil {
		updateData["price"] = *job.Price
	}

	err := jr.db.Model(&models.Job{}).Where("id = ?", job.ID).Updates(updateData).Error
	if err != nil {
		return nil, err
	}

	updatedJob := &models.Job{}
	err = jr.db.Where("id = ?", job.ID).First(updatedJob).Error
	if err != nil {
		return nil, err
	}

	return updatedJob, nil
}

func (jr *JobRepository) Delete(jobID uint) error {
	job := models.Job{ID: jobID}
	return jr.db.Delete(&job).Error
}
