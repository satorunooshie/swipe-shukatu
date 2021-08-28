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
	recruitRepoimpl := repoimpl.NewRecruitRepoImpl(db)
	jobRepoimpl := repoimpl.NewJobRepoImpl(db)
	ltdRepoimpl := repoimpl.NewLtdRepoImpl(db)
	messageRepoimpl := repoimpl.NewMessageRepoImpl(db)
	likeRepoimpl := repoimpl.NewLikeRepoImpl(db)
	matchlistRepoimpl := repoimpl.NewMatchlistRepoImpl(db)
	superlikeRepoimpl := repoimpl.NewSuperlikeRepoImpl(db)
	nopeRepoimpl := repoimpl.NewNopeRepoImpl(db)

	/* usecase */
	// userUseCase := usecase.NewUserUsecase(userRepoimpl)
	recruitUseCase := usecase.NewRecruitUsecase(recruitRepoimpl)
	messageUseCase := usecase.NewMessageUsecase(jobRepoimpl, ltdRepoimpl, messageRepoimpl, recruitRepoimpl)
	likeusecase := usecase.NewLikeUsecase(likeRepoimpl)
	mathclistusecase := usecase.NewMatchlistUsecase(matchlistRepoimpl)
	superlikeUsecase := usecase.NewSuperlikeUsecase(superlikeRepoimpl)
	nopeUsecase := usecase.NewNopeUsecase(nopeRepoimpl)

	/* handler */
	// userHandler := handler.NewUserHandler(userUseCase)
	recruitHandler := handler.NewRecruitHandler(recruitUseCase)
	messageHandler := handler.NewMessageHandler(messageUseCase)
	likeHandler := handler.NewLikeHandler(likeusecase)
	matchlisthandler := handler.NewMatchlistHandler(mathclistusecase)
	superlikeHandler := handler.NewSuperlikeHandler(superlikeUsecase)
	nopeHandler := handler.NewNopeHandler(nopeUsecase)

	// register the handler
	h.Handle("/recruits", middleware.Get(recruitHandler.HandleSelect()))
	h.Handle("/message", middleware.Get(m.Auth(messageHandler.HandleSelect())))
	h.Handle("/message/", middleware.Post(m.Auth(messageHandler.HandleInsert())))
	h.Handle("/like", middleware.Post(m.Auth(likeHandler.HandleInsert())))
	h.Handle("/match/list", middleware.Get(m.Auth(matchlisthandler.HandleSelect())))
	h.Handle("/superlike", middleware.Post(m.Auth(superlikeHandler.HandleInsert())))
	h.Handle("/nope", middleware.Post(m.Auth(nopeHandler.HandleInsert())))

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
