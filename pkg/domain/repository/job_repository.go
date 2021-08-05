//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	jobM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type JobRepository interface {
	Select(ctx context.Context) ([]*jobM.Job, error)
	Insert(ctx context.Context, entity *jobM.Job) error
	Update(ctx context.Context, entity *jobM.Job) error
	Delete(ctx context.Context, entity *jobM.Job) error
}
