package handler

import (
	"golang-advance/assignment-1/service"

	"github.com/gin-gonic/gin"
)

type ISubmissionHandler interface {
	CreateSubmission(c *gin.Context)
	GetSubmission(c *gin.Context)
	GetAllSubmissions(c *gin.Context)
	DeleteSubmission(c *gin.Context)
}

func NewSubmissionHandler(submissionService service.ISubmissionService) ISubmissionHandler {
	return &SubmissionHandler{
		submissionService: submissionService,
	}
}

type SubmissionHandler struct {
	submissionService service.ISubmissionService
}

func (s *SubmissionHandler) CreateSubmission(c *gin.Context) {
	panic("implement me")
}

func (s *SubmissionHandler) GetSubmission(c *gin.Context) {
	panic("implement me")
}

func (s *SubmissionHandler) DeleteSubmission(c *gin.Context) {
	panic("implement me")
}

func (s *SubmissionHandler) GetAllSubmissions(c *gin.Context) {
	panic("implement me")
}
