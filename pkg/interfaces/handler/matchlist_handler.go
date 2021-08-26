//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/satorunooshie/swipe-shukatu/pkg/dcontext"
	matchlistU "github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

type MatchlistHandler interface {
	HandleSelect() http.HandlerFunc
	HandleInsert() http.HandlerFunc
	HandleUpdate() http.HandlerFunc
	HandleDelete() http.HandlerFunc
}

type matchlistHandler struct {
	matchlistUseCase matchlistU.MatchlistUseCase
}

// NewMatchlistHandler
func NewMatchlistHandler(matchlistU matchlistU.MatchlistUseCase) MatchlistHandler {
	return &matchlistHandler{
		matchlistUseCase: matchlistU,
	}
}

// HandleSelect
func (matchlistH *matchlistHandler) HandleSelect() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		UID := dcontext.GetUIDFromContext(ctx)
		if UID == "" {
			log.Printf("[ERROR] failed to GetUID: ")
			http.Error(writer, "Could not get UID", http.StatusBadRequest)
			return
		}
		matchlists, err := matchlistH.matchlistUseCase.Select(ctx, UID)
		// TODO: cache
		if err != nil {
			log.Printf("[ERROR] at matchlistUseCase.Select: failed to get matchlist: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		res := make([]*matchlistResponse, len(matchlists))
		for i, l := range matchlists {
			var ml matchlistResponse
			ml.LtdID = l.LtdID
			ml.RecruitID = l.RecruitID
			ml.Name = l.Name
			ml.Image = l.Image
			ml.Reactiontype = l.Reactiontype
			res[i] = &ml
		}
		var respms matchlistResponses
		respms.Matchlists = res
		jsonresponse, err := json.Marshal(respms)
		if err != nil {
			log.Printf("[ERROR] at matchlist_handler.HandleSelect: failed to marshal matchlists: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsonresponse)
	}
}

// HandleInsert
func (matchlistH *matchlistHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleUpdate
func (matchlistH *matchlistHandler) HandleUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleDelete
func (matchlistH *matchlistHandler) HandleDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

type matchlistResponses struct {
	Matchlists []*matchlistResponse `json:"matchlists"`
}

// MatchlistResponse
type matchlistResponse struct {
	LtdID        int32  `json:"ltd_id"`
	RecruitID    int32  `json:"recruit_id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Reactiontype int8   `json:"reactiontype"`
}
