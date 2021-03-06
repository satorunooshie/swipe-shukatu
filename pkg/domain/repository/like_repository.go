//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	likeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type LikeRepository interface {
	Select(ctx context.Context, UID string) ([]*likeM.Like, error)
	Insert(ctx context.Context, entity *likeM.Like, UID string) error
	Update(ctx context.Context, entity *likeM.Like) error
	Delete(ctx context.Context, entity *likeM.Like) error
}
