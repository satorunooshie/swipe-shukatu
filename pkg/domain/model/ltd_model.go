package model

import "time"

type Ltd struct {
	ID             int32
	Name           string
	Profile        string
	EmployeeNumber int32
	AverageAge     int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
