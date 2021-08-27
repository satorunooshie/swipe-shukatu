//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	superlikeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type SuperlikeRepository interface {
	Select(ctx context.Context, UID string) ([]*superlikeM.Superlike, error)
	Insert(ctx context.Context, entity *superlikeM.Superlike) error
	Update(ctx context.Context, entity *superlikeM.Superlike) error
	Delete(ctx context.Context, entity *superlikeM.Superlike) error
}
