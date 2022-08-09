package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

const defaultPort = "3100"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	muxer := chi.NewMux()

	muxer.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf(`{"time":"%s"}`, time.Now().String())
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(str))
	})

	log.Printf("connect to http://localhost:%s/ to test API", port)
	log.Fatal(http.ListenAndServe(":"+port, muxer))
}
