package handlers

import (
	"net/http"
	"server/models"
	"server/services"
	"server/utils"
	"strconv"

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
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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
	paramJobID := c.Param("jobID")

	parsedJobID, err := strconv.ParseUint(paramJobID, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}
	jobID := uint(parsedJobID)

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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
	paramJobID := c.Param("jobID")

	parsedID, err := strconv.ParseUint(paramJobID, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}
	jobID := uint(parsedID)

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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

