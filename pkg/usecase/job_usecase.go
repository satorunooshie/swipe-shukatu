//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"

	jobM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	jobR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type JobUseCase interface {
	Select(ctx context.Context) ([]*jobM.Job, error)
	Insert(ctx context.Context, entity *jobM.Job) error
	Update(ctx context.Context, entity *jobM.Job) error
	Delete(ctx context.Context, entity *jobM.Job) error
}

type jobUseCase struct {
	jobRepository jobR.JobRepository
}

// NewJobUsecase
func NewJobUsecase(jobR jobR.JobRepository) JobUseCase {
	return &jobUseCase{
		jobRepository: jobR,
	}
}

// Select
func (jobU *jobUseCase) Select(ctx context.Context) ([]*jobM.Job, error) {
	jobs, err := jobU.jobRepository.Select(ctx)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

// Insert
func (jobU *jobUseCase) Insert(ctx context.Context, entity *jobM.Job) error {
	err := jobU.jobRepository.Insert(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Update
func (jobU *jobUseCase) Update(ctx context.Context, entity *jobM.Job) error {
	err := jobU.jobRepository.Update(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (jobU *jobUseCase) Delete(ctx context.Context, entity *jobM.Job) error {
	err := jobU.jobRepository.Delete(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
