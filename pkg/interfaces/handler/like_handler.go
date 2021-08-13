//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	likeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	mid "github.com/satorunooshie/swipe-shukatu/pkg/interfaces/middleware"
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
		panic("do something")
	}
}

// HandleInsert
func (likeH *likeHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		method := request.Method
		UID := mid.GetUIDFromContext(ctx)
		if method == "POST" {
			defer request.Body.Close()
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				log.Fatal(err)
			}
			var like *likeM.Like
			err = json.Unmarshal(body, &like)
			if err != nil {
				log.Fatal(err)
			}
			like.UID = UID
			err = likeH.likeUseCase.Insert(ctx, like)
			if err != nil {
				log.Fatal(err)
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			} else {
				writer.WriteHeader(http.StatusCreated)
			}
		}
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

// LikeResponse
type LikeResponse struct { // nolint
	// Need to implement field
}
