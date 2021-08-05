package model

import "time"

type Job struct {
	ID      int32
	Name    string
	SortID  int32
	Created time.Time
	Updated time.Time
	Deleted time.Time
}
