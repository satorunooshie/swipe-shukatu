//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"

	likeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	likeR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type LikeUseCase interface {
	Select(ctx context.Context) ([]*likeM.Like, error)
	Insert(ctx context.Context, entity *likeM.Like) error
	Update(ctx context.Context, entity *likeM.Like) error
	Delete(ctx context.Context, entity *likeM.Like) error
}

type likeUseCase struct {
	likeRepository likeR.LikeRepository
}

// NewLikeUsecase
func NewLikeUsecase(likeR likeR.LikeRepository) LikeUseCase {
	return &likeUseCase{
		likeRepository: likeR,
	}
}

// Select
func (likeU *likeUseCase) Select(ctx context.Context) ([]*likeM.Like, error) {
	likes, err := likeU.likeRepository.Select(ctx)
	if err != nil {
		return nil, err
	}
	return likes, nil
}

// Insert
func (likeU *likeUseCase) Insert(ctx context.Context, entity *likeM.Like) error {
	err := likeU.likeRepository.Insert(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Update
func (likeU *likeUseCase) Update(ctx context.Context, entity *likeM.Like) error {
	err := likeU.likeRepository.Update(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (likeU *likeUseCase) Delete(ctx context.Context, entity *likeM.Like) error {
	err := likeU.likeRepository.Delete(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
