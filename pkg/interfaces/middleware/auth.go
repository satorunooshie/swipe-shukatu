package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
)

//nolint
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// opt := option.WithCredentialsFile(os.Getenv("firebase-sdk"))
		opt := option.WithCredentialsFile("./firebase-sdk.json")
		ctx := r.Context()
		if ctx != nil {
			ctx = context.Background()
		}
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			log.Printf("[ERROR] auth::Auth: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		auth, err := app.Auth(ctx)
		if err != nil {
			log.Printf("[ERROR] auth::Auth: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := auth.VerifyIDToken(ctx, idToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			// TODO: Delete debug code instead return json error message
			_, _ = w.Write([]byte("error verifying ID token\n"))
			return
		}
		SetUID(ctx, token.UID)
		log.Printf("[INFO] Verified ID token: %v\n", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
