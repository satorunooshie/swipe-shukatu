//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"

	ltdM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	ltdR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type LtdUseCase interface {
	Select(ctx context.Context) ([]*ltdM.Ltd, error)
	Insert(ctx context.Context, entity *ltdM.Ltd) error
	Update(ctx context.Context, entity *ltdM.Ltd) error
	Delete(ctx context.Context, entity *ltdM.Ltd) error
}

type ltdUseCase struct {
	ltdRepository ltdR.LtdRepository
}

// NewLtdUsecase
func NewLtdUsecase(ltdR ltdR.LtdRepository) LtdUseCase {
	return &ltdUseCase{
		ltdRepository: ltdR,
	}
}

// Select
func (ltdU *ltdUseCase) Select(ctx context.Context) ([]*ltdM.Ltd, error) {
	ltds, err := ltdU.ltdRepository.Select(ctx)
	if err != nil {
		return nil, err
	}
	return ltds, nil
}

// Insert
func (ltdU *ltdUseCase) Insert(ctx context.Context, entity *ltdM.Ltd) error {
	err := ltdU.ltdRepository.Insert(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Update
func (ltdU *ltdUseCase) Update(ctx context.Context, entity *ltdM.Ltd) error {
	err := ltdU.ltdRepository.Update(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (ltdU *ltdUseCase) Delete(ctx context.Context, entity *ltdM.Ltd) error {
	err := ltdU.ltdRepository.Delete(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
