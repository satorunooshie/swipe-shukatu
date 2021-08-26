package middleware

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go"

	"github.com/satorunooshie/swipe-shukatu/pkg/dcontext"
	"github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	"github.com/satorunooshie/swipe-shukatu/pkg/infrastructure/mysql/repoimpl"
	"github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

type Auth struct {
	db *sql.DB
}

func NewAuth(db *sql.DB) *Auth {
	return &Auth{
		db: db,
	}
}

//nolint
func (a *Auth) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		if idToken == "" {
			log.Printf("[INFO] auth::Auth: %v\n", errors.New("header is not set"))
			w.WriteHeader(http.StatusBadRequest)
			// TODO: Delete debug code instead return json error message
			_, _ = w.Write([]byte("empty token\n"))
			return
		}

		path, err := getFilePath()
		if err != nil {
			log.Printf("[ERROR] auth::Auth::getFilePath: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		opt := option.WithCredentialsFile(path)

		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			log.Printf("[ERROR] auth::Auth::firebase.NewApp: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		auth, err := app.Auth(ctx)
		if err != nil {
			log.Printf("[ERROR] auth::Auth::app.Auth: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token, err := auth.VerifyIDToken(ctx, idToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			// TODO: Delete debug code instead return json error message
			log.Println("error verifying ID token")
			return
		}
		uri := repoimpl.NewUserRepoImpl(a.db)
		uc := usecase.NewUserUsecase(uri)
		user, err := uc.Select(ctx, token.UID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO: Delete debug code instead return json error message
			log.Printf("error select user: %v\n", err)
		}
		if user == nil {
			ent := &model.User{
				UUID:             token.UID,
				RegisterMethodID: 1, // Google
			}
			if err := uc.Insert(ctx, ent); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				// TODO: Delete debug code instead return json error message
				log.Printf("error insert user: %v\n", err)
				return
			}
		}
		log.Printf("[INFO] Verified ID token: %v\n", token)
		dcontext.SetUID(ctx, token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// NOTE: 実行パスが違うのでトリミング
func getFilePath() (string, error) {
	const (
		dirname  = "swipe-shukatu"
		filename = "firebase-sdk.json"
	)
	cp, err := os.Getwd()
	if err != nil {
		return "", err
	}
	idx := strings.LastIndex(cp, dirname)
	return cp[:(idx+len(dirname))] + "/" + filename, nil
}
