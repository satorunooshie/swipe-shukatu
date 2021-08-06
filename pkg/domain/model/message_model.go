package model

import "time"

type MType int

const (
	Txt MType = iota + 1
	Remind
	Image
)

type Message struct {
	ID        int32
	UserID    int32
	RecruitID int32
	Type      MType
	Content   string
	ImagePath string
	CreatedAt time.Time
	DeletedAt time.Time
}
