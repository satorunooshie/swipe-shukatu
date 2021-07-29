package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setting := struct {
			Frequency int `json:"frequency"`
		}{
			Frequency: 0,
		}
		_ = json.NewEncoder(w).Encode(setting)
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
