//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	recruitM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type RecruitRepository interface {
	SelectRecruits(ctx context.Context, Param *Parameters) ([]*Recruit, error)
	SelectRecruitForMessage(ctx context.Context, rID int32) (*recruitM.Recruit, error)
	Insert(ctx context.Context, entity *recruitM.Recruit) error
	Update(ctx context.Context, entity *recruitM.Recruit) error
	Delete(ctx context.Context, entity *recruitM.Recruit) error
}

type Recruit struct {
	LtdID                 int32  `json:"ltd_id"`
	RecruitID             int32  `json:"recruit_id"`
	Name                  string `json:"name"`
	Profile               string `json:"profile"`
	EmployeeNumber        int32  `json:"employee_number"`
	AverageAge            int32  `json:"average_age"`
	JobType               string `json:"job_type"`
	EducationalBackground string `json:"educational_background"`
	AverageSalary         int32  `json:"average_salary"`
	StartingSalary        int32  `json:"starting_salary"`
	Image                 string `json:"image"`
}

type Parameters struct {
	Location         int32
	JobType          int32
	EducationHistory int32
	Benefits         int32
	MinSalary        int32
	MaxSalary        int32
	StartingSalary   int32
}
