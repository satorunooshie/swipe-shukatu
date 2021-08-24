//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/satorunooshie/swipe-shukatu/pkg/dcontext"
	superlikeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	superlikeU "github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

type SuperlikeHandler interface {
	HandleSelect() http.HandlerFunc
	HandleInsert() http.HandlerFunc
	HandleUpdate() http.HandlerFunc
	HandleDelete() http.HandlerFunc
}

type superlikeHandler struct {
	superlikeUseCase superlikeU.SuperlikeUseCase
}

// NewSuperlikeHandler
func NewSuperlikeHandler(superlikeU superlikeU.SuperlikeUseCase) SuperlikeHandler {
	return &superlikeHandler{
		superlikeUseCase: superlikeU,
	}
}

// HandleSelect
func (superlikeH *superlikeHandler) HandleSelect() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		UID := dcontext.GetUIDFromContext(ctx)
		superlikes, err := superlikeH.superlikeUseCase.Select(ctx, UID)
		if err != nil {
			log.Printf("[ERROR] failed to Select from superlike: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		res := make([]*superlikeInsertResponse, len(superlikes))
		for i, l := range superlikes {
			var ls superlikeInsertResponse
			ls.RecruitID = l.RecruitID
			ls.CreatedAt = l.CreatedAt
			res[i] = &ls
		}
		var respms superlikeInsertResponses
		respms.Superlikes = res
		jsonresponse, err := json.Marshal(respms)
		if err != nil {
			log.Printf("[ERROR] failed to marshal superlikes: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsonresponse)
	}
}

type superlikeInsertRequest struct {
	RecruitID int32 `json:"recruit_id"`
}

// HandleInsert
func (superlikeH *superlikeHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		SuperlikeRequest := new(superlikeInsertRequest)
		if err := json.NewDecoder(request.Body).Decode(&SuperlikeRequest); err != nil {
			log.Printf("[ERROR] failed to JsonDecord: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO::読み込んだjsonデータのバリデーションがほしい
		ctx := request.Context()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		UID := dcontext.GetUIDFromContext(ctx)
		if UID == "" {
			log.Printf("[INFO] failed to GetUID: ")
			http.Error(writer, "Could not get UID", http.StatusBadRequest)
			return
		}
		superlike := new(superlikeM.Superlike)
		superlike.RecruitID = SuperlikeRequest.RecruitID
		err := superlikeH.superlikeUseCase.Insert(ctx, superlike, UID)
		if err != nil {
			log.Printf("[ERROR] failed to Insert: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	}
}

// HandleUpdate
func (superlikeH *superlikeHandler) HandleUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleDelete
func (superlikeH *superlikeHandler) HandleDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// SuperlikeRequest
type superlikeRequest struct { // nolint
	// Need to implement field
}

type superlikeInsertResponses struct {
	Superlikes []*superlikeInsertResponse `json:"superlikes"`
}

// SuperlikeResponse ...
type superlikeInsertResponse struct {
	RecruitID int32     `json:"recruit_id"`
	CreatedAt time.Time `json:"created_at"`
}
