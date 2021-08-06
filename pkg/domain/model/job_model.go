package model

import "time"

type Job struct {
	ID        int32
	Name      string
	SortID    int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
