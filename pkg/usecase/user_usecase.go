//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"

	userM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	userR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type UserUseCase interface {
	Select(ctx context.Context, uuid string) (*userM.User, error)
	Insert(ctx context.Context, entity *userM.User) error
	Update(ctx context.Context, entity *userM.User) error
	Delete(ctx context.Context, entity *userM.User) error
}

type userUseCase struct {
	userRepository userR.UserRepository
}

// NewUserUsecase ...
func NewUserUsecase(userR userR.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userR,
	}
}

// Select ...
func (userU *userUseCase) Select(ctx context.Context, uuid string) (*userM.User, error) {
	return userU.userRepository.Select(ctx, uuid)
}

// Insert ...
func (userU *userUseCase) Insert(ctx context.Context, entity *userM.User) error {
	err := userU.userRepository.Insert(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (userU *userUseCase) Update(ctx context.Context, entity *userM.User) error {
	err := userU.userRepository.Update(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (userU *userUseCase) Delete(ctx context.Context, entity *userM.User) error {
	err := userU.userRepository.Delete(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
