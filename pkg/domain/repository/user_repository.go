//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	userM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type UserRepository interface {
	Select(ctx context.Context, uuid string) (*userM.User, error)
	Insert(ctx context.Context, entity *userM.User) error
	Update(ctx context.Context, entity *userM.User) error
	Delete(ctx context.Context, entity *userM.User) error
}
