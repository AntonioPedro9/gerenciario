package services

import (
	"server/models"
	"server/repositories"
	"server/utils"

	"github.com/google/uuid"
)

type JobService struct {
	jobRepository *repositories.JobRepository
}

func NewJobService(jobRepository *repositories.JobRepository) *JobService {
	return &JobService{jobRepository}
}

func (js *JobService) CreateJob(job *models.CreateJobRequest) error {
	if !utils.IsValidName(job.Name) {
		return utils.InvalidNameError
	}

	if job.Duration < 0 {
		return utils.InvalidDurationError
	}

	if job.Price < 0 {
		return utils.InvalidPriceError
	}

	validJob := &models.Job{
		Name:        utils.CapitalizeText(job.Name),
		Description: utils.CapitalizeText(job.Description),
		Duration:    job.Duration,
		Price:       job.Price,
		UserID:      job.UserID,
	}

	return js.jobRepository.Create(validJob)
}

func (js *JobService) ListJobs(userID uuid.UUID) ([]models.Job, error) {
	return js.jobRepository.List(userID)
}

func (js *JobService) GetJob(jobID uint, tokenID uuid.UUID) (*models.Job, error) {
	job, err := js.jobRepository.GetJobById(jobID)
	if err != nil {
		return nil, err
	}

	if job.UserID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	return job, nil
}

func (js *JobService) UpdateJob(job *models.UpdateJobRequest, tokenID uuid.UUID) (*models.Job, error) {
	if job.UserID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	if job.Name != nil && !utils.IsValidName(*job.Name) {
		return nil, utils.InvalidNameError
	}

	if job.Duration != nil && *job.Duration <= 0 {
		return nil, utils.InvalidDurationError
	}

	if job.Price != nil && *job.Price <= 0 {
		return nil, utils.InvalidPriceError
	}

	existingJob, err := js.jobRepository.GetJobById(job.ID)
	if err != nil {
		return nil, err
	}
	if existingJob == nil {
		return nil, utils.NotFoundError
	}

	if job.Name != nil {
		capitalizedName := utils.CapitalizeText(*job.Name)
		job.Name = &capitalizedName
	}

	if job.Description != nil {
		capitalizedDescription := utils.CapitalizeText(*job.Description)
		job.Description = &capitalizedDescription
	}

	updatedJob, err := js.jobRepository.Update(job)
	if err != nil {
		return nil, err
	}

	return updatedJob, nil
}

func (js *JobService) DeleteJob(jobID uint, tokenID uuid.UUID) error {
	existingJob, err := js.jobRepository.GetJobById(jobID)
	if err != nil {
		return err
	}

	if existingJob.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	return js.jobRepository.Delete(jobID)
}
