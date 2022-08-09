package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	muxer := chi.NewMux()

	muxer.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		jsBytes, _ := json.Marshal(fmt.Sprintf(`{"time":"%s"}`, time.Now().String()))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsBytes)
	})

	http.ListenAndServe(":"+port, muxer)
}
