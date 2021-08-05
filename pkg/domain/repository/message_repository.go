//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	messageM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type MessageRepository interface {
	Select(ctx context.Context, rID int32) ([]*messageM.Message, error)
	Insert(ctx context.Context, entity *messageM.Message) error
	Update(ctx context.Context, entity *messageM.Message) error
	Delete(ctx context.Context, entity *messageM.Message) error
}
