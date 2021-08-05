package router

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/satorunooshie/swipe-shukatu/pkg/infrastructure/mysql/repoimpl"
	"github.com/satorunooshie/swipe-shukatu/pkg/interfaces/handler"
	"github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

func Route(h *http.ServeMux, db *sql.DB) {
	// DI

	/* user */
	// userRepoimpl := repoimpl.NewUserRepoImpl(db)
	// userUseCase := usecase.NewUserUsecase(userRepoimpl)
	// userHandler := handler.NewUserHandler(userUseCase)

	/* job */
	jobRepoimpl := repoimpl.NewJobRepoImpl(db)

	/* recruit */
	recruitRepoimpl := repoimpl.NewRecruitRepoImpl(db)

	messageRepoimpl := repoimpl.NewMessageRepoImpl(db)
	messageUseCase := usecase.NewMessageUsecase(jobRepoimpl, messageRepoimpl, recruitRepoimpl)
	messageHandler := handler.NewMessageHandler(messageUseCase)

	// register the handler
	h.Handle("/message/", messageHandler.HandleSelect())

	h.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		setting := struct {
			Ping string `json:"ping"`
		}{
			Ping: "pong",
		}
		_ = json.NewEncoder(w).Encode(setting)
	})
}
