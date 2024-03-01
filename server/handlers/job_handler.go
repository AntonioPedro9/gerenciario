package handlers

import (
	"net/http"
	"server/models"
	"server/services"
	"server/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type JobHandler struct {
	jobService *services.JobService
}

func NewJobHandler(jobService *services.JobService) *JobHandler {
	return &JobHandler{jobService}
}

/** 
 * Creates a new job.
 * It accepts a JSON body with the job details.
 * Returns 201 if the job is created successfully.
 * Returns 400 if the request fails to bind to JSON.
 * Returns 401 if token userID does not match request userID
 * Returns 500 for internal server errors.
**/
func (jh *JobHandler) CreateJob(c *gin.Context) {
	var job models.CreateJobRequest
	if err := c.ShouldBindJSON(&job); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	if err := jh.jobService.CreateJob(&job, tokenID); err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

/** 
 * Lists all jobs for a user.
 * It extracts userID from JWT token.
 * Returns 200 along with a list of jobs.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (jh *JobHandler) ListJobs(c *gin.Context) {
	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	jobs, err := jh.jobService.ListJobs(tokenID)
	if err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusOK, jobs)
}

/** 
 * Gets a job by ID.
 * It requires jobID as a path parameter and extracts userID from JWT token.
 * Returns 200 along with the job details.
 * Returns 400 if the job ID fails to parse.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (jh *JobHandler) GetJob(c *gin.Context) {
	jobID, err := utils.GetParamID("jobID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	job, err := jh.jobService.GetJob(jobID, tokenID)
	if err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusOK, job)
}

/** 
 * Updates a job.
 * It accepts a JSON body with the job details and extracts userID from JWT token.
 * Returns 200 if the job is updated successfully.
 * Returns 400 if the request fails to bind to JSON.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (jh *JobHandler) UpdateJob(c *gin.Context) {
	var job models.UpdateJobRequest
	if err := c.ShouldBindJSON(&job); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	updatedJob, err := jh.jobService.UpdateJob(&job, tokenID)
	if err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusOK, updatedJob)
}

/** 
 * Deletes a job.
 * It requires jobID as a path parameter and extracts userID from JWT token.
 * Returns 204 if the job is deleted successfully.
 * Returns 400 if the job ID fails to parse.
 * Returns 401 if the token is unauthorized.
 * Returns 500 for internal server errors.
**/
func (jh *JobHandler) DeleteJob(c *gin.Context) {
	jobID, err := utils.GetParamID("jobID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	tokenID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
		return
	}

	if err := jh.jobService.DeleteJob(jobID, tokenID); err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
