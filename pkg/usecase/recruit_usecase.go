//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"

	recruitM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	recruitR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type RecruitUseCase interface {
	Select(ctx context.Context) ([]*recruitM.Recruit, error)
	Insert(ctx context.Context, entity *recruitM.Recruit) error
	Update(ctx context.Context, entity *recruitM.Recruit) error
	Delete(ctx context.Context, entity *recruitM.Recruit) error
}

type recruitUseCase struct {
	recruitRepository recruitR.RecruitRepository
}

// NewRecruitUsecase
func NewRecruitUsecase(recruitR recruitR.RecruitRepository) RecruitUseCase {
	return &recruitUseCase{
		recruitRepository: recruitR,
	}
}

// Select
func (recruitU *recruitUseCase) Select(ctx context.Context) ([]*recruitM.Recruit, error) {
	recruits, err := recruitU.recruitRepository.Select(ctx)
	if err != nil {
		return nil, err
	}
	return recruits, nil
}

// Insert
func (recruitU *recruitUseCase) Insert(ctx context.Context, entity *recruitM.Recruit) error {
	err := recruitU.recruitRepository.Insert(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Update
func (recruitU *recruitUseCase) Update(ctx context.Context, entity *recruitM.Recruit) error {
	err := recruitU.recruitRepository.Update(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (recruitU *recruitUseCase) Delete(ctx context.Context, entity *recruitM.Recruit) error {
	err := recruitU.recruitRepository.Delete(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
