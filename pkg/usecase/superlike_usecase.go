//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"

	"github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	superlikeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	superlikeR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type SuperLikeUseCase interface {
	Select(ctx context.Context, UID string) ([]*superlikeM.SuperLike, error)
	Insert(ctx context.Context, entity *superlikeM.SuperLike, UID string) error
	Update(ctx context.Context, entity *superlikeM.SuperLike) error
	Delete(ctx context.Context, entity *superlikeM.SuperLike) error
}

type superlikeUseCase struct {
	superlikeRepository superlikeR.SuperLikeRepository
}

// NewSuperLikeUsecase
func NewSuperLikeUsecase(superlikeR superlikeR.SuperLikeRepository) SuperLikeUseCase {
	return &superlikeUseCase{
		superlikeRepository: superlikeR,
	}
}

// Select
func (superlikeU *superlikeUseCase) Select(ctx context.Context, UID string) ([]*superlikeM.SuperLike, error) {
	var superlikes []*model.SuperLike
	superlikes, err := superlikeU.superlikeRepository.Select(ctx, UID)
	if err != nil {
		return nil, err
	}
	return superlikes, nil
}

// Insert
func (superlikeU *superlikeUseCase) Insert(ctx context.Context, entity *superlikeM.SuperLike, UID string) error {
	err := superlikeU.superlikeRepository.Insert(ctx, entity, UID)
	if err != nil {
		return err
	}
	return nil
}

// Update
func (superlikeU *superlikeUseCase) Update(ctx context.Context, entity *superlikeM.SuperLike) error {
	err := superlikeU.superlikeRepository.Update(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (superlikeU *superlikeUseCase) Delete(ctx context.Context, entity *superlikeM.SuperLike) error {
	err := superlikeU.superlikeRepository.Delete(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
