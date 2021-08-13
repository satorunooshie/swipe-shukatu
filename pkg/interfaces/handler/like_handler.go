//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/satorunooshie/swipe-shukatu/pkg/dcontext"
	likeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	likeU "github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

type LikeHandler interface {
	HandleSelect() http.HandlerFunc
	HandleInsert() http.HandlerFunc
	HandleUpdate() http.HandlerFunc
	HandleDelete() http.HandlerFunc
}

type likeHandler struct {
	likeUseCase likeU.LikeUseCase
}

// NewLikeHandler
func NewLikeHandler(likeU likeU.LikeUseCase) LikeHandler {
	return &likeHandler{
		likeUseCase: likeU,
	}
}

// HandleSelect
func (likeH *likeHandler) HandleSelect() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		UID := dcontext.GetUIDFromContext(ctx)
		lk, err := likeH.likeUseCase.Select(ctx, UID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		lks := make([]*LikeResponse, len(lk))
		for i, l := range lk {
			var ls LikeResponse
			ls.RecruitID = l.RecruitID
			ls.CreatedAt = l.CreatedAt
			lks[i] = &ls
		}
		var respms LikeResponses
		respms.Likes = lks
		jsonresponse, err := json.Marshal(respms)
		if err != nil {
			log.Printf("[ERROR] failed to marshal messages: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsonresponse)
	}
}

// HandleInsert
func (likeH *likeHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		var like *likeM.Like
		if err := json.NewDecoder(request.Body).Decode(&like); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO::読み込んだjsonデータのバリデーションがほしい
		ctx := request.Context()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		UID := dcontext.GetUIDFromContext(ctx)
		err := likeH.likeUseCase.Insert(ctx, like, UID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	}
}

type Rid struct {
	RID string
}

// HandleUpdate
func (likeH *likeHandler) HandleUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleDelete
func (likeH *likeHandler) HandleDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// LikeRequest
type LikeRequest struct { // nolint
	// Need to implement field
}

type LikeResponses struct {
	Likes []*LikeResponse `json:"likes"`
}

// LikeResponse ...
type LikeResponse struct {
	RecruitID int32     `json:"recruit_id"`
	CreatedAt time.Time `json:"created_at"`
}
