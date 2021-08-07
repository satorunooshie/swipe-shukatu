package model

import "time"

type User struct {
	UUID                 int32
	registered_method_id int32
	gender               int32
	graduate_year        int32
	CreatedAt            time.Time
	UpdatedAt            time.Time
	deleted_at           time.Time
}
