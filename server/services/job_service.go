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

func (js *JobService) CreateJob(job *models.CreateJobRequest, tokenID uuid.UUID) error {
	if job.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	formattedName, err := utils.FormatName(job.Name)
	if err != nil {
		return err
	}
	if job.Price < 0 {
		return utils.InvalidPriceError
	}

	validJob := &models.Job{
		Name:        formattedName,
		Description: job.Description,
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

func (js *JobService) UpdateJob(job *models.UpdateJobRequest, tokenID uuid.UUID) error {
	if job.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	existingJob, err := js.jobRepository.GetJobById(job.ID)
	if err != nil {
		return err
	}
	if existingJob == nil {
		return utils.NotFoundError
	}

	if job.Name != nil {
		formattedName, err := utils.FormatName(*job.Name)
		if err != nil {
			return err
		}
		job.Name = &formattedName
	}

	if job.Price != nil && *job.Price < 0 {
		return utils.InvalidPriceError
	}

	return js.jobRepository.Update(job)
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
