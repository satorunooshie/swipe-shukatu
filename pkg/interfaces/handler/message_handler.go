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

	"github.com/satorunooshie/swipe-shukatu/pkg/dcontext"
	messageM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
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

type messageInsertRequest struct {
	MessageType int32  `json:"message_type"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	ExecuteAt   string `json:"execute_at"`
}

// HandleInsert
func (messageH *messageHandler) HandleInsert() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		rID, err := strconv.Atoi(strings.TrimPrefix(request.URL.Path, "/message/"))
		if err != nil {
			log.Printf("invalid recruitID: %s", err)
			http.Error(writer, "recruitID is invalid", http.StatusBadRequest)
			return
		}
		messageRequest := new(messageInsertRequest)
		if err := json.NewDecoder(request.Body).Decode(&messageRequest); err != nil {
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
			log.Printf("[INFO] failed to GetUID")
			http.Error(writer, "Could not get UID", http.StatusBadRequest)
			return
		}
		message := new(messageM.Message)
		message.UserID = UID
		message.RecruitID = int32(rID)
		message.Type = messageM.MType(messageRequest.MessageType)
		switch message.Type {
		case 1:
			message.Content = messageRequest.Content
			if err := messageH.messageUseCase.InsertMessage(ctx, message); err != nil {
				log.Printf("[ERROR] failed to InsertMessage: %v", err.Error())
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		case 2:
			message.ImagePath = messageRequest.Image
			if err := messageH.messageUseCase.InsertMessage(ctx, message); err != nil {
				log.Printf("[ERROR] failed to InsertImagePath: %v", err.Error())
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		case 3:
			message.Content = messageRequest.Content
			str := messageRequest.ExecuteAt
			form := "2006-01-02T15:04:05+09:00"
			t, _ := time.Parse(form, str)
			if err := messageH.messageUseCase.InsertRemind(ctx, message, t); err != nil {
				log.Printf("[ERROR] failed to InsertRemind: %v", err.Error())
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			log.Printf("[INFO] message_type is undefind")
			http.Error(writer, "message_type must be 1,2,3", http.StatusBadRequest)
			return
		}

		writer.WriteHeader(http.StatusCreated)
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

// MessageResponse ...
type MessageResponses struct {
	Messages []*MessageResponse `json:"messages"`
}

// MessageResponse ...
type MessageResponse struct {
	LtdID     int32     `json:"ltd_id"`
	RecruitID int32     `json:"recruit_id"`
	Name      string    `json:"name"`
	JobType   string    `json:"job_type"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}
