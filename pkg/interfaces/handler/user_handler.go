//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"net/http"

	userU "github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

type UserHandler interface {
	HandleSelect() http.HandlerFunc
	HandleInsert() http.HandlerFunc
	HandleUpdate() http.HandlerFunc
	HandleDelete() http.HandlerFunc
}

type userHandler struct {
	userUseCase userU.UserUseCase
}

// NewUserHandler
func NewUserHandler(userU userU.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: userU,
	}
}

// HandleSelect
func (userH *userHandler) HandleSelect() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleInsert
func (userH *userHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleUpdate
func (userH *userHandler) HandleUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleDelete
func (userH *userHandler) HandleDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// UserRequest
type UserRequest struct {
	// Need to implement field
}

// UserResponse
type UserResponse struct {
	// Need to implement field
}
