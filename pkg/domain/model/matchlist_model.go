package model

type Matchlist struct {
	LtdID        int32  `json:"ltd_id"`
	RecruitID    int32  `json:"recruit_id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Reactiontype int32  `json:"reactiontype"`
}
