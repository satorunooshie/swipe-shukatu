//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"net/http"

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
		panic("do something")
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
