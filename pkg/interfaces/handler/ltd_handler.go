//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"net/http"

	ltdU "github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

type LtdHandler interface {
	HandleSelect() http.HandlerFunc
	HandleInsert() http.HandlerFunc
	HandleUpdate() http.HandlerFunc
	HandleDelete() http.HandlerFunc
}

type ltdHandler struct {
	ltdUseCase ltdU.LtdUseCase
}

// NewLtdHandler
func NewLtdHandler(ltdU ltdU.LtdUseCase) LtdHandler {
	return &ltdHandler{
		ltdUseCase: ltdU,
	}
}

// HandleSelect
func (ltdH *ltdHandler) HandleSelect() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleInsert
func (ltdH *ltdHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleUpdate
func (ltdH *ltdHandler) HandleUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleDelete
func (ltdH *ltdHandler) HandleDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// LtdRequest
type LtdRequest struct {
	// Need to implement field
}

// LtdResponse
type LtdResponse struct {
	// Need to implement field
}
