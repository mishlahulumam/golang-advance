package postgres_gorm

import (
	"context"
	"golang-advance/assignment-1/entity"
	"golang-advance/assignment-1/service"
)

type submissionRepository struct {
	db GormDBIface
}

func NewSubmissionRepository(db GormDBIface) service.ISubmissionRepository {
	return &submissionRepository{db: db}
}

func (s *submissionRepository) CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error) {
	panic("implement me")
}

func (s *submissionRepository) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	panic("implement me")
}

func (s *submissionRepository) GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error) {
	panic("implement me")
}

func (s *submissionRepository) DeleteSubmission(ctx context.Context, id int) error {
	panic("implement me")
}

func (s *submissionRepository) GetAllSubmissions(ctx context.Context) ([]entity.Submission, error) {
	panic("implement me")
}
