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

	/* middleware */
	m := middleware.NewAuth(db)

	/* repoimpl */
	// userRepoimpl := repoimpl.NewUserRepoImpl(db)
	jobRepoimpl := repoimpl.NewJobRepoImpl(db)
	ltdRepoimpl := repoimpl.NewLtdRepoImpl(db)
	messageRepoimpl := repoimpl.NewMessageRepoImpl(db)
	recruitRepoimpl := repoimpl.NewRecruitRepoImpl(db)
	likeRepoimpl := repoimpl.NewLikeRepoImpl(db)
	superlikeRepoimpl := repoimpl.NewSuperlikeRepoImpl(db)
	nopeRepoimpl := repoimpl.NewNopeRepoImpl(db)

	/* usecase */
	// userUseCase := usecase.NewUserUsecase(userRepoimpl)
	messageUseCase := usecase.NewMessageUsecase(jobRepoimpl, ltdRepoimpl, messageRepoimpl, recruitRepoimpl)
	likeusecase := usecase.NewLikeUsecase(likeRepoimpl)
	superlikeUsecase := usecase.NewSuperlikeUsecase(superlikeRepoimpl)
	nopeUsecase := usecase.NewNopeUsecase(nopeRepoimpl)

	/* handler */
	// userHandler := handler.NewUserHandler(userUseCase)
	messageHandler := handler.NewMessageHandler(messageUseCase)
	likeHandler := handler.NewLikeHandler(likeusecase)
	superlikeHandler := handler.NewSuperlikeHandler(superlikeUsecase)
	nopeHandler := handler.NewNopeHandler(nopeUsecase)

	// register the handler
	// h.Handle("/message", m.Auth(middleware.Get(messageHandler.HandleSelect())))
	h.Handle("/message", middleware.Get(messageHandler.HandleSelect()))
	h.Handle("/like", middleware.Post(likeHandler.HandleInsert()))
	h.Handle("/superlike", middleware.Post(superlikeHandler.HandleInsert()))
	h.Handle("/nope", middleware.Post(nopeHandler.HandleInsert()))

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
