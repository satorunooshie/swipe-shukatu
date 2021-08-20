//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"

	superlikeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	superlikeR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type SuperlikeUseCase interface {
	Select(ctx context.Context, UDI string) ([]*superlikeM.Superlike, error)
	Insert(ctx context.Context, entity *superlikeM.Superlike, UID string) error
	Update(ctx context.Context, entity *superlikeM.Superlike) error
	Delete(ctx context.Context, entity *superlikeM.Superlike) error
}

type superlikeUseCase struct {
	superlikeRepository superlikeR.SuperlikeRepository
}

// NewSuperlikeUsecase
func NewSuperlikeUsecase(superlikeR superlikeR.SuperlikeRepository) SuperlikeUseCase {
	return &superlikeUseCase{
		superlikeRepository: superlikeR,
	}
}

// Select
func (superlikeU *superlikeUseCase) Select(ctx context.Context, UID string) ([]*superlikeM.Superlike, error) {
	superlikes, err := superlikeU.superlikeRepository.Select(ctx, UID)
	if err != nil {
		return nil, err
	}
	return superlikes, nil
}

// Insert
func (superlikeU *superlikeUseCase) Insert(ctx context.Context, entity *superlikeM.Superlike, UID string) error {
	err := superlikeU.superlikeRepository.Insert(ctx, entity, UID)
	if err != nil {
		return err
	}
	return nil
}

// Update
func (superlikeU *superlikeUseCase) Update(ctx context.Context, entity *superlikeM.Superlike) error {
	err := superlikeU.superlikeRepository.Update(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (superlikeU *superlikeUseCase) Delete(ctx context.Context, entity *superlikeM.Superlike) error {
	err := superlikeU.superlikeRepository.Delete(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
