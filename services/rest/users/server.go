package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/helpers"
)

const defaultPort = "3101"

var users []*model.User

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := helpers.LoadUsers()
	if nil != err {
		panic("Unable to load users.")
	}

	users := helpers.GetUsers()

	muxer := chi.NewMux()

	muxer.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		jsBytes, _ := json.Marshal(users)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsBytes)
	})

	muxer.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		userIDString := chi.URLParam(r, "id")
		userID, _ := strconv.Atoi(userIDString)

		jsBytes, _ := json.Marshal(users[userID-1])
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsBytes)
	})

	log.Printf("connect to http://localhost:%s/ to test API", port)
	log.Fatal(http.ListenAndServe(":"+port, muxer))
}
