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

func (jh *JobHandler) CreateJob(c *gin.Context) {
	var job models.CreateJobRequest
	if err := c.ShouldBindJSON(&job); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	if err := jh.jobService.CreateJob(&job); err != nil {
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

func (jh *JobHandler) ListJobs(c *gin.Context) {
	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized action"})
		return
	}

	jobs, err := jh.jobService.ListJobs(userID)
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

func (jh *JobHandler) GetJob(c *gin.Context) {
	jobID, err := utils.GetParamID("jobID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized action"})
		return
	}

	job, err := jh.jobService.GetJob(jobID, userID)
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

func (jh *JobHandler) UpdateJob(c *gin.Context) {
	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized action"})
		return
	}

	var job models.UpdateJobRequest
	if err := c.ShouldBindJSON(&job); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	updatedJob, err := jh.jobService.UpdateJob(&job, userID)
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

func (jh *JobHandler) DeleteJob(c *gin.Context) {
	jobID, err := utils.GetParamID("jobID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized action"})
		return
	}

	if err := jh.jobService.DeleteJob(jobID, userID); err != nil {
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

