package model

import "time"

type Recruit struct {
	ID                      int32
	LtdID                   int32
	JobTypeID               int32
	EducationalBackgroundID int32
	AverageSalary           int32
	StartingSalaray         int32
	CreatedAt               time.Time
	UpdatedAt               time.Time
	IsInvalidAt             time.Time
}
