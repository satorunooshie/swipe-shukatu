//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	recruitM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type RecruitRepository interface {
	Select(ctx context.Context) ([]*recruitM.Recruit, error)
	SelectRecruitForMessage(ctx context.Context, rID int32) (*recruitM.Recruit, error)
	Insert(ctx context.Context, entity *recruitM.Recruit) error
	Update(ctx context.Context, entity *recruitM.Recruit) error
	Delete(ctx context.Context, entity *recruitM.Recruit) error
}
