//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"

	"github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	matchlistM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	matchlistR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type MatchlistUseCase interface {
	Select(ctx context.Context, UID string) ([]*matchlistM.Matchlist, error)
	Insert(ctx context.Context, entity *matchlistM.Matchlist) error
	Update(ctx context.Context, entity *matchlistM.Matchlist) error
	Delete(ctx context.Context, entity *matchlistM.Matchlist) error
}

type matchlistUseCase struct {
	matchlistRepository matchlistR.MatchlistRepository
}

// NewMatchlistUsecase
func NewMatchlistUsecase(matchlistR matchlistR.MatchlistRepository) MatchlistUseCase {
	return &matchlistUseCase{
		matchlistRepository: matchlistR,
	}
}

// Select
func (matchlistU *matchlistUseCase) Select(ctx context.Context, UID string) ([]*matchlistM.Matchlist, error) {
	var matchlists []*model.Matchlist
	matchlists, err := matchlistU.matchlistRepository.Select(ctx, UID)
	if err != nil {
		return nil, err
	}
	return matchlists, nil
}

// Insert
func (matchlistU *matchlistUseCase) Insert(ctx context.Context, entity *matchlistM.Matchlist) error {
	err := matchlistU.matchlistRepository.Insert(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Update
func (matchlistU *matchlistUseCase) Update(ctx context.Context, entity *matchlistM.Matchlist) error {
	err := matchlistU.matchlistRepository.Update(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (matchlistU *matchlistUseCase) Delete(ctx context.Context, entity *matchlistM.Matchlist) error {
	err := matchlistU.matchlistRepository.Delete(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
