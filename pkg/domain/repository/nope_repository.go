//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	nopeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type NopeRepository interface {
	Select(ctx context.Context, UID string) ([]*nopeM.Nope, error)
	Insert(ctx context.Context, entity *nopeM.Nope, UID string) error
	Update(ctx context.Context, entity *nopeM.Nope) error
	Delete(ctx context.Context, entity *nopeM.Nope) error
}
