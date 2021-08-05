//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"net/http"

	jobU "github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

type JobHandler interface {
	HandleSelect() http.HandlerFunc
	HandleInsert() http.HandlerFunc
	HandleUpdate() http.HandlerFunc
	HandleDelete() http.HandlerFunc
}

type jobHandler struct {
	jobUseCase jobU.JobUseCase
}

// NewJobHandler
func NewJobHandler(jobU jobU.JobUseCase) JobHandler {
	return &jobHandler{
		jobUseCase: jobU,
	}
}

// HandleSelect
func (jobH *jobHandler) HandleSelect() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleInsert
func (jobH *jobHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleUpdate
func (jobH *jobHandler) HandleUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleDelete
func (jobH *jobHandler) HandleDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// JobRequest
type JobRequest struct { // nolint
	// Need to implement field
}

// JobResponse
type JobResponse struct { // nolint
	// Need to implement field
}
