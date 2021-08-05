package model

import "time"

type MTyp int

const (
	Txt MTyp = iota + 1
	Remind
	ImgTyp
)

type Message struct {
	ID        int32
	UserID    int32
	RecruitID int32
	Type      MTyp
	Content   string
	ImgPath   string
	Created   time.Time
	Deleted   time.Time
}
