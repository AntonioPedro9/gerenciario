package handlers

import (
	"net/http"
	"server/internals/models"
	services "server/internals/services"
	"server/pkg/errors"
	"server/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobHandlerInterface interface {
	CreateJob(c *gin.Context)
	GetJob(c *gin.Context)
	GetUserJobs(c *gin.Context)
	UpdateJob(c *gin.Context)
	DeleteJob(c *gin.Context)
}

type JobHandler struct {
	jobService services.JobServiceInterface
}

func NewJobHandler(jobService services.JobServiceInterface) *JobHandler {
	return &JobHandler{jobService}
}

func (jh *JobHandler) CreateJob(c *gin.Context) {
	var data models.CreateJobRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		errors.HandleError(c, errors.JsonBindingError)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := jh.jobService.CreateJob(&data, tokenUserId); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (jh *JobHandler) GetJob(c *gin.Context) {
	jobId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	job, err := jh.jobService.GetJob(uint(jobId), tokenUserId)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, job)
}

func (jh *JobHandler) GetUserJobs(c *gin.Context) {
	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	jobs, err := jh.jobService.GetUserJobs(tokenUserId)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (jh *JobHandler) UpdateJob(c *gin.Context) {
	jobId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	var data models.UpdateJobRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		errors.HandleError(c, errors.JsonBindingError)
		return
	}

	if err := jh.jobService.UpdateJob(uint(jobId), tokenUserId, &data); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (jh *JobHandler) DeleteJob(c *gin.Context) {
	jobId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := jh.jobService.DeleteJob(uint(jobId), tokenUserId); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
