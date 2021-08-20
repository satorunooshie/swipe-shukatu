//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"

	"github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	nopeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	nopeR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type NopeUseCase interface {
	Select(ctx context.Context, UID string) ([]*nopeM.Nope, error)
	Insert(ctx context.Context, entity *nopeM.Nope, UID string) error
	Update(ctx context.Context, entity *nopeM.Nope) error
	Delete(ctx context.Context, entity *nopeM.Nope) error
}

type nopeUseCase struct {
	nopeRepository nopeR.NopeRepository
}

// NewNopeUsecase
func NewNopeUsecase(nopeR nopeR.NopeRepository) NopeUseCase {
	return &nopeUseCase{
		nopeRepository: nopeR,
	}
}

// Select
func (nopeU *nopeUseCase) Select(ctx context.Context, UID string) ([]*nopeM.Nope, error) {
	var nopes []*model.Nope
	nopes, err := nopeU.nopeRepository.Select(ctx, UID)
	if err != nil {
		return nil, err
	}
	return nopes, nil
}

// Insert
func (nopeU *nopeUseCase) Insert(ctx context.Context, entity *nopeM.Nope, UID string) error {
	err := nopeU.nopeRepository.Insert(ctx, entity, UID)
	if err != nil {
		return err
	}
	return nil
}

// Update
func (nopeU *nopeUseCase) Update(ctx context.Context, entity *nopeM.Nope) error {
	err := nopeU.nopeRepository.Update(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (nopeU *nopeUseCase) Delete(ctx context.Context, entity *nopeM.Nope) error {
	err := nopeU.nopeRepository.Delete(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
