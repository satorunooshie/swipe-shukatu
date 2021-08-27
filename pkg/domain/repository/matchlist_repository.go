//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"
)

type MatchlistRepository interface {
	Select(ctx context.Context, UID string) ([]*Matchlist, error)
	Insert(ctx context.Context, entity *Matchlist) error
	Update(ctx context.Context, entity *Matchlist) error
	Delete(ctx context.Context, entity *Matchlist) error
}

type Matchlist struct {
	LtdID        int32  `json:"ltd_id"`
	RecruitID    int32  `json:"recruit_id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Reactiontype int32  `json:"reactiontype"`
}
