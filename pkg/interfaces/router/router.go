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
	userRepoimpl := repoimpl.NewUserRepoImpl(db)
	userUseCase := usecase.NewUserUsecase(userRepoimpl)
	userHandler := handler.NewUserHandler(userUseCase)

	// register the handler
	h.Handle("/api/user", userHandler.HandleSelect())

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setting := struct {
			Frequency int `json:"frequency"`
		}{
			Frequency: 0,
		}
		_ = json.NewEncoder(w).Encode(setting)
	})
}
