//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"
)

type MatchlistRepository interface {
	Select(ctx context.Context, UID string) ([]*Match, error)
	Insert(ctx context.Context, entity *Match) error
	Update(ctx context.Context, entity *Match) error
	Delete(ctx context.Context, entity *Match) error
}

type Match struct {
	LtdID        int32  `json:"ltd_id"`
	RecruitID    int32  `json:"recruit_id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Reactiontype int32  `json:"reaction_type"`
}
