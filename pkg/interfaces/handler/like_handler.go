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
		likes, err := likeH.likeUseCase.Select(ctx, UID) //likes []*model.Like
		if err != nil {
			log.Printf("[ERROR] failed to Select from like: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		res := make([]*LikeResponse, len(likes))
		for i, l := range likes {
			var ls LikeResponse
			ls.RecruitID = l.RecruitID
			ls.CreatedAt = l.CreatedAt
			res[i] = &ls
		}
		var respms LikeResponses
		respms.Likes = res
		jsonresponse, err := json.Marshal(respms)
		if err != nil {
			log.Printf("[ERROR] failed to marshal likes: %v", err.Error())
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
		like := new(likeM.Like)
		if err := json.NewDecoder(request.Body).Decode(&like); err != nil {
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
			log.Printf("[ERROR] failed to GetUID: ")
			http.Error(writer, "Could not get UID", http.StatusInternalServerError)
			return
		}
		err := likeH.likeUseCase.Insert(ctx, like, UID)
		if err != nil {
			log.Printf("[ERROR] failed to Insert: %v", err.Error())
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
