package middleware

import (
	"log"
	"net/http"
)

func Get(api http.HandlerFunc) http.HandlerFunc {
	return httpMethod(api, http.MethodGet)
}

//nolint
func Post(api http.HandlerFunc) http.HandlerFunc {
	return httpMethod(api, http.MethodPost)
}

//nolint
func Put(api http.HandlerFunc) http.HandlerFunc {
	return httpMethod(api, http.MethodPut)
}

//nolint
func Delete(api http.HandlerFunc) http.HandlerFunc {
	return httpMethod(api, http.MethodDelete)
}

func httpMethod(api http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const client = "http://localhost:3000"
		w.Header().Add("Access-Control-Allow-Origin", client)
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Accept, Origin, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			return
		}
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			if _, err := w.Write([]byte("Method not allowed")); err != nil {
				log.Println(err)
			}
			return
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		api(w, r)
	}
}
