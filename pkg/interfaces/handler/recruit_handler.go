//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	recruitR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
	recruitU "github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

type RecruitHandler interface {
	HandleSelect() http.HandlerFunc
	HandleInsert() http.HandlerFunc
	HandleUpdate() http.HandlerFunc
	HandleDelete() http.HandlerFunc
}

type recruitHandler struct {
	recruitUseCase recruitU.RecruitUseCase
}

// NewRecruitHandler
func NewRecruitHandler(recruitU recruitU.RecruitUseCase) RecruitHandler {
	return &recruitHandler{
		recruitUseCase: recruitU,
	}
}

// HandleSelect
func (recruitH *recruitHandler) HandleSelect() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		location, _ := strconv.Atoi(request.FormValue("location"))
		jobType, _ := strconv.Atoi(request.FormValue("job_type"))
		educationHistory, _ := strconv.Atoi(request.FormValue("education_history"))
		benefits, _ := strconv.Atoi(request.FormValue("benefits"))
		minSalary, _ := strconv.Atoi(request.FormValue("min_salary"))
		maxSalary, _ := strconv.Atoi(request.FormValue("max_salary"))
		startingSalary, _ := strconv.Atoi(request.FormValue("starting_salary"))
		parameters := new(recruitR.Parameters)
		parameters.Location = int32(location)
		parameters.JobType = int32(jobType)
		parameters.EducationHistory = int32(educationHistory)
		parameters.Benefits = int32(benefits)
		parameters.MinSalary = int32(minSalary)
		parameters.MaxSalary = int32(maxSalary)
		parameters.StartingSalary = int32(startingSalary)
		rs, err := recruitH.recruitUseCase.SelectRecruits(ctx, parameters)
		if err != nil {
			log.Printf("[ERROR] failed to fetch recruit: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		recs := make([]*recruitR.Recruit, len(rs))
		copy(recs, rs)
		var respms RecruitResponse
		respms.Recruits = recs
		jrespms, err := json.Marshal(respms)
		if err != nil {
			log.Printf("[ERROR] failed to marshal messages: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jrespms)
	}
}

// HandleInsert
func (recruitH *recruitHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleUpdate
func (recruitH *recruitHandler) HandleUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleDelete
func (recruitH *recruitHandler) HandleDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// RecruitRequest
type RecruitRequest struct { // nolint
	// Need to implement field
}

// RecruitResponse
type RecruitResponse struct { // nolint
	Recruits []*recruitR.Recruit `json:"recruits"`
}
