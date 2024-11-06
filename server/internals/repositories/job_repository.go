package repositories

import (
	"server/internals/models"
	"server/pkg/errors"
	"server/pkg/logs"

	"gorm.io/gorm"
)

type JobRepositoryInterface interface {
	Create(data *models.CreateJobRequest, tokenUserId uint) error
	GetById(jobId, tokenUserId uint) (*models.GetJobResponse, error)
	GetUserJobs(tokenUserId uint) ([]models.GetJobResponse, error)
	Update(jobId, tokenUserId uint, data *models.UpdateJobRequest) error
	Delete(jobId, tokenUserId uint) error
}

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db}
}

func (jr *JobRepository) Create(data *models.CreateJobRequest, tokenUserId uint) error {
	job := &models.Job{
		UserId:         tokenUserId,
		Name:           data.Name,
		Description:    data.Description,
		Duration:       data.Duration,
		AfterSalesDays: data.AfterSalesDays,
		Price:          float32(data.Price),
	}

	result := jr.db.Create(job)

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	return nil
}

func (jr *JobRepository) GetById(jobId, tokenUserId uint) (*models.GetJobResponse, error) {
	var job models.Job

	result := jr.db.Where("id = ? AND user_id = ?", jobId, tokenUserId).First(&job)

	if result.Error != nil {
		logs.LogError(result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.NotFoundError
	}

	return &models.GetJobResponse{
		Id:             job.Id,
		UserId:         job.UserId,
		Name:           job.Name,
		Description:    job.Description,
		Duration:       job.Duration,
		AfterSalesDays: job.AfterSalesDays,
		Price:          job.Price,
	}, nil
}

func (jr *JobRepository) GetUserJobs(tokenUserId uint) ([]models.GetJobResponse, error) {
	var jobs []models.Job
	var jobResponse []models.GetJobResponse

	result := jr.db.Where("user_id = ?", tokenUserId).Find(&jobs)

	if result.Error != nil {
		logs.LogError(result.Error)
		return nil, result.Error
	}

	for _, job := range jobs {
		jobResponse = append(jobResponse, models.GetJobResponse{
			Id:             job.Id,
			UserId:         job.UserId,
			Name:           job.Name,
			Description:    job.Description,
			Duration:       job.Duration,
			AfterSalesDays: job.AfterSalesDays,
			Price:          job.Price,
		})
	}

	return jobResponse, nil
}

func (jr *JobRepository) Update(jobId, tokenUserId uint, data *models.UpdateJobRequest) error {
	job := &models.Job{
		Name:           data.Name,
		Description:    data.Description,
		Duration:       data.Duration,
		AfterSalesDays: data.AfterSalesDays,
		Price:          float32(data.Price),
	}

	result := jr.db.Model(&models.Job{}).Where("id = ? AND user_id = ?", jobId, tokenUserId).Updates(job)

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.NotFoundError
	}

	return nil
}

func (jr *JobRepository) Delete(jobId, tokenUserId uint) error {
	result := jr.db.Where("id = ? AND user_id = ?", jobId, tokenUserId).Delete(&models.Job{})

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.NotFoundError
	}

	return nil
}
