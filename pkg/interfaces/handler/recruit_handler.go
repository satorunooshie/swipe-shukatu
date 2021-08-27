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
		parameters := new(recruitR.Parameters)
		location, _ := strconv.Atoi(request.FormValue("location"))
		parameters.Location = location
		job_type, _ := strconv.Atoi(request.FormValue("job_type"))
		parameters.JobType = job_type
		education_history, _ := strconv.Atoi(request.FormValue("education_history"))
		parameters.EducationHistory = education_history
		benefits
		ms, err := messageH.messageUseCase.Select(ctx, int32(rID))
		if err != nil {
			log.Printf("[ERROR] failed to fetch messages by recruit ID: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		mrs := make([]*MessageResponse, len(ms))
		for i, m := range ms {
			var mr MessageResponse
			mr.LtdID = m.LtdID
			mr.RecruitID = m.RecruitID
			mr.Name = m.Name
			mr.JobType = m.JobType
			mr.Content = m.Content
			mr.Image = m.Image
			mr.CreatedAt = m.CreatedAt
			mrs[i] = &mr
		}
		var respms MessageResponses
		respms.Messages = mrs
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
	// Need to implement field
}
