package services

import (
	"server/models"
	"server/repositories"
	"server/utils"
	"server/utils/validations"

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

	err := validations.ValidateCreateJobRequest(job)
	if err != nil {
		return err
	}

	validJob := &models.Job{
		Name:        utils.CapitalizeText(job.Name),
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

func (js *JobService) UpdateJob(job *models.UpdateJobRequest, tokenID uuid.UUID) (*models.Job, error) {
	if job.UserID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	existingJob, err := js.jobRepository.GetJobById(job.ID)
	if err != nil {
		return nil, err
	}
	if existingJob == nil {
		return nil, utils.NotFoundError
	}

	err = validations.ValidateUpdateJobRequest(job)
	if err != nil {
		return nil, err
	}

	if job.Name != nil {
		capitalizedName := utils.CapitalizeText(*job.Name)
		job.Name = &capitalizedName
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
