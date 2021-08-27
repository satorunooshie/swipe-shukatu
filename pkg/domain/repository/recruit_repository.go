//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	recruitM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

type RecruitRepository interface {
	Select(ctx context.Context) ([]*recruitM.Recruit, error)
	SelectRecruits(ctx context.Context, Param *Parameters) (*Recruits, error)
	SelectRecruitForMessage(ctx context.Context, rID int32) (*recruitM.Recruit, error)
	Insert(ctx context.Context, entity *recruitM.Recruit) error
	Update(ctx context.Context, entity *recruitM.Recruit) error
	Delete(ctx context.Context, entity *recruitM.Recruit) error
}

type Recruits struct {
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
	Location         int
	JobType          int
	EducationHistory int
	Benefits         int
	MinSalary        int
	MaxSalary        int
	StartingSalary   int
}
