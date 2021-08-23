//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	matchlistM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type MatchlistRepository interface {
	Select(ctx context.Context, UID string) ([]*matchlistM.Matchlist, error)
	Insert(ctx context.Context, entity *matchlistM.Matchlist) error
	Update(ctx context.Context, entity *matchlistM.Matchlist) error
	Delete(ctx context.Context, entity *matchlistM.Matchlist) error
}
