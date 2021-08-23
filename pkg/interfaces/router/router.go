package router

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/satorunooshie/swipe-shukatu/pkg/infrastructure/mysql/repoimpl"
	"github.com/satorunooshie/swipe-shukatu/pkg/interfaces/handler"
	"github.com/satorunooshie/swipe-shukatu/pkg/interfaces/middleware"
	"github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

func Route(h *http.ServeMux, db *sql.DB) {
	// DI

	/* repoimpl */
	// userRepoimpl := repoimpl.NewUserRepoImpl(db)
	jobRepoimpl := repoimpl.NewJobRepoImpl(db)
	ltdRepoimpl := repoimpl.NewLtdRepoImpl(db)
	messageRepoimpl := repoimpl.NewMessageRepoImpl(db)
	recruitRepoimpl := repoimpl.NewRecruitRepoImpl(db)
	likeRepoimpl := repoimpl.NewLikeRepoImpl(db)
	matchlistRepoimpl := repoimpl.NewMatchlistRepoImpl(db)

	/* usecase */
	// userUseCase := usecase.NewUserUsecase(userRepoimpl)
	messageUseCase := usecase.NewMessageUsecase(jobRepoimpl, ltdRepoimpl, messageRepoimpl, recruitRepoimpl)
	likeusecase := usecase.NewLikeUsecase(likeRepoimpl)
	mathclistusecase := usecase.NewMatchlistUsecase(matchlistRepoimpl)

	/* handler */
	// userHandler := handler.NewUserHandler(userUseCase)
	messageHandler := handler.NewMessageHandler(messageUseCase)
	likeHandler := handler.NewLikeHandler(likeusecase)
	matchlisthandler := handler.NewMatchlistHandler(mathclistusecase)

	// register the handler
	// h.Handle("/message/", middleware.Auth(middleware.Get(messageHandler.HandleSelect())))
	h.Handle("/message/", middleware.Get(messageHandler.HandleSelect()))
	h.Handle("/like", middleware.Post(likeHandler.HandleInsert()))
	h.Handle("/match/list", middleware.Get(matchlisthandler.HandleSelect()))

	// this endpoint is for health check
	h.HandleFunc("/health", middleware.Get(func(w http.ResponseWriter, r *http.Request) {
		health := struct {
			Ping string `json:"ping"`
		}{
			Ping: "pong",
		}
		_ = json.NewEncoder(w).Encode(health)
	}))
}
