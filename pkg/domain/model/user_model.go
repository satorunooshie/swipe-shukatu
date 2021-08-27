package model

import (
	"time"
)

type User struct {
	UUID             string
	RegisterMethodID int32
	Gender           int32
	GraduateYear     int32
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}
