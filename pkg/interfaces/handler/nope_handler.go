//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/satorunooshie/swipe-shukatu/pkg/dcontext"
	nopeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	nopeU "github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

type NopeHandler interface {
	HandleSelect() http.HandlerFunc
	HandleInsert() http.HandlerFunc
	HandleUpdate() http.HandlerFunc
	HandleDelete() http.HandlerFunc
}

type nopeHandler struct {
	nopeUseCase nopeU.NopeUseCase
}

// NewNopeHandler
func NewNopeHandler(nopeU nopeU.NopeUseCase) NopeHandler {
	return &nopeHandler{
		nopeUseCase: nopeU,
	}
}

// HandleSelect
func (nopeH *nopeHandler) HandleSelect() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		UID := dcontext.GetUIDFromContext(ctx)
		nopes, err := nopeH.nopeUseCase.Select(ctx, UID)
		if err != nil {
			log.Printf("[ERROR] failed to Select from nope: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		res := make([]*nopeResponse, len(nopes))
		for i, l := range nopes {
			var ls nopeResponse
			ls.RecruitID = l.RecruitID
			ls.CreatedAt = l.CreatedAt
			res[i] = &ls
		}
		var respms nopeResponses
		respms.Nopes = res
		jsonresponse, err := json.Marshal(respms)
		if err != nil {
			log.Printf("[ERROR] failed to marshal nopes: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsonresponse)
	}
}

type nopeInsertRequest struct {
	RecruitID int32 `json:"recruit_id"`
}

// HandleInsert
func (nopeH *nopeHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		nopeRequest := new(nopeInsertRequest)
		if err := json.NewDecoder(request.Body).Decode(&nopeRequest); err != nil {
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
		nope := new(nopeM.Nope)
		nope.RecruitID = nopeRequest.RecruitID
		err := nopeH.nopeUseCase.Insert(ctx, nope, UID)
		if err != nil {
			log.Printf("[ERROR] failed to Insert: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	}
}

// HandleUpdate
func (nopeH *nopeHandler) HandleUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleDelete
func (nopeH *nopeHandler) HandleDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// NopeRequest
type nopeRequest struct { // nolint
	// Need to implement field
}

type nopeResponses struct {
	Nopes []*nopeResponse `json:"nopes"`
}

// NopeResponse ...
type nopeResponse struct {
	RecruitID int32     `json:"recruit_id"`
	CreatedAt time.Time `json:"created_at"`
}
