//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	ltdM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type LtdRepository interface {
	Select(ctx context.Context) ([]*ltdM.Ltd, error)
	SelectLtdNameByID(ctx context.Context, ltdID int32) (string, error)
	Insert(ctx context.Context, entity *ltdM.Ltd) error
	Update(ctx context.Context, entity *ltdM.Ltd) error
	Delete(ctx context.Context, entity *ltdM.Ltd) error
}
