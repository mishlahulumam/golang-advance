package service

import (
	"context"
	"fmt"
	"golang-advance/assignment-1/entity"
)

type ISubmissionService interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) error
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
	GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.Submission, error)
}

type ISubmissionRepository interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
	GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.Submission, error)
}

type submissionService struct {
	submissionRepo ISubmissionRepository
}

func NewSubmissionService(submissionRepo ISubmissionRepository) ISubmissionService {
	return &submissionService{submissionRepo: submissionRepo}
}

func (s *submissionService) CreateSubmission(ctx context.Context, submission *entity.Submission) error {
	score, category, definition := calculateProfileRiskFromAnswers(submission.Answers)
	fmt.Println(score, category, definition)

	panic("implement me")
}

func (s *submissionService) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	panic("implement me")
}

func (s *submissionService) GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error) {
	panic("implement me")
}

func (s *submissionService) DeleteSubmission(ctx context.Context, id int) error {
	panic("implement me")
}

func (s *submissionService) GetAllSubmissions(ctx context.Context) ([]entity.Submission, error) {
	panic("implement me")
}

func calculateProfileRiskFromAnswers(answers []entity.Answer) (score int, category entity.ProfileRiskCategory, definition string) {
	return score, category, definition
}
