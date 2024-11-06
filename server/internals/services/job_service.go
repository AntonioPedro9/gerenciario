package services

import (
	"server/internals/models"
	"server/internals/repositories"
)

type JobServiceInterface interface {
	CreateJob(data *models.CreateJobRequest, tokenUserId uint) error
	GetJob(jobId, tokenUserId uint) (models.GetJobResponse, error)
	GetUserJobs(tokenUserId uint) (models.GetUserJobsResponse, error)
	UpdateJob(jobId, tokenUserId uint, data *models.UpdateJobRequest) error
	DeleteJob(jobId, tokenUserId uint) error
}

type JobService struct {
	jobRepository repositories.JobRepositoryInterface
}

func NewJobService(jobRepository repositories.JobRepositoryInterface) *JobService {
	return &JobService{jobRepository}
}

func (js *JobService) CreateJob(data *models.CreateJobRequest, tokenUserId uint) error {
	if err := validate.Struct(data); err != nil {
		return err
	}
	return js.jobRepository.Create(data, tokenUserId)
}

func (js *JobService) GetJob(jobId, tokenUserId uint) (models.GetJobResponse, error) {
	var emptyJob models.GetJobResponse

	job, err := js.jobRepository.GetById(jobId, tokenUserId)
	if err != nil {
		return emptyJob, err
	}

	return *job, nil
}

func (js *JobService) GetUserJobs(tokenUserId uint) (models.GetUserJobsResponse, error) {
	var emptyJobs models.GetUserJobsResponse

	jobs, err := js.jobRepository.GetUserJobs(tokenUserId)
	if err != nil {
		return emptyJobs, err
	}

	jobsArray := models.GetUserJobsResponse{
		Jobs: jobs,
	}

	return jobsArray, nil
}

func (js *JobService) UpdateJob(jobId, tokenUserId uint, data *models.UpdateJobRequest) error {
	if err := validate.Struct(data); err != nil {
		return err
	}
	return js.jobRepository.Update(jobId, tokenUserId, data)
}

func (js *JobService) DeleteJob(jobId, tokenUserId uint) error {
	return js.jobRepository.Delete(jobId, tokenUserId)
}
