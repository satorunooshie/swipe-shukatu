package model

import "time"

type Recruit struct {
	ID                      int32
	LtdID                   int32
	JobTypeID               int32
	EducationalBackgroundID int32
	AvgSalary               int32
	StartingSalaray         int32
	Created                 time.Time
	Updated                 time.Time
	IsInvalidAt             time.Time
}
