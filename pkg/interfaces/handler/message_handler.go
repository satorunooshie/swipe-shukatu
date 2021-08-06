//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	messageU "github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

type MessageHandler interface {
	HandleSelect() http.HandlerFunc
	HandleInsert() http.HandlerFunc
	HandleUpdate() http.HandlerFunc
	HandleDelete() http.HandlerFunc
}

type messageHandler struct {
	messageUseCase messageU.MessageUseCase
}

// NewMessageHandler
func NewMessageHandler(messageU messageU.MessageUseCase) MessageHandler {
	return &messageHandler{
		messageUseCase: messageU,
	}
}

// HandleSelect
func (messageH *messageHandler) HandleSelect() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		rID, err := strconv.Atoi(strings.TrimPrefix(request.URL.Path, "/message/"))
		if err != nil {
			log.Printf("invalid recruitID: %s", err)
			http.Error(writer, "recruitID is invalid", http.StatusBadRequest)
			return
		}
		ms, err := messageH.messageUseCase.Select(ctx, int32(rID))
		if err != nil {
			log.Printf("[ERROR] failed to fetch messages by recruit ID: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		mrs := make([]*MessageResponse, len(ms))
		for i, m := range ms {
			var mr MessageResponse
			mr.LtdID = m.LtdID
			mr.RecruitID = m.RecruitID
			mr.Name = m.Name
			mr.JobType = m.JobType
			mr.Content = m.Content
			mr.Image = m.Image
			mr.CreatedAt = m.CreatedAt
			mrs[i] = &mr
		}
		var respms MessageResponses
		respms.Messages = mrs
		jrespms, err := json.Marshal(respms)
		if err != nil {
			log.Printf("[ERROR] failed to marshal messages: %v", err.Error())
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jrespms)
	}
}

// HandleInsert
func (messageH *messageHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleUpdate
func (messageH *messageHandler) HandleUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// HandleDelete
func (messageH *messageHandler) HandleDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		panic("do something")
	}
}

// MessageRequest
type MessageRequest struct { // nolint
	// Need to implement field
}

type MessageResponses struct {
	Messages []*MessageResponse `json:"messages"`
}

// MessageResponse
type MessageResponse struct {
	LtdID     int32     `json:"ltd_id"`
	RecruitID int32     `json:"recruit_id"`
	Name      string    `json:"name"`
	JobType   string    `json:"job_type"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}
