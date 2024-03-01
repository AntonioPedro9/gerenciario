package validations

import (
	"server/models"
	"server/utils"
)

func ValidateCreateJobRequest(job *models.CreateJobRequest) error {
	if len(job.Name) < 2 {
		return utils.InvalidNameError
	}

	if job.Price < 0 {
		return utils.InvalidPriceError
	}

	return nil
}

func ValidateUpdateJobRequest(job *models.UpdateJobRequest) error {
	if job.Name != nil {
		if len(*job.Name) < 2 {
			return utils.InvalidNameError
		}
	}

	if job.Price != nil {
		if *job.Price < 0 {
			return utils.InvalidNameError
		}
	}

	return nil
}