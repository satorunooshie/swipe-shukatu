package model

type Matchlist struct {
	LtdID        int32  `json:"ltd_id"`
	RecruitID    int32  `json:"recruit_id"`
	Name         string `json:"name"`
	JobType      string `json:"job_type"`
	Image        string `json:"image"`
	Reactiontype int8   `json:"reactiontype"`
}