package model

import "time"

type Like struct {
	UID        string    `json:"-"`
	RecruitID  int32     `json:"recruit_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateddAt time.Time `json:"updated_at"`
}
