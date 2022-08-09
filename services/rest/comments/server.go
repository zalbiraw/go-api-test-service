package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/zalbiraw/go-api-test-service/helpers"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/comments/graph/model"
)

const defaultPort = "4003"

var comments []*model.Comment

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := helpers.LoadComments()
	if nil != err {
		panic("Unable to load comments.")
	}

	comments := helpers.GetComments()

	muxer := chi.NewMux()

	muxer.Get("/posts/{id}/comments", func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "id")

		var postComments []*model.Comment

		for _, comment := range comments {
			if comment.PostID == postID {
				postComments = append(postComments, comment)
			}
		}

		jsBytes, _ := json.Marshal(postComments)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsBytes)
	})

	http.ListenAndServe(":"+port, muxer)
}
