package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/zalbiraw/go-api-test-service/helpers"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/graph/model"
	"net/http"
	"os"
)

const defaultPort = "4002"

var posts []*model.Post

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := helpers.LoadPosts()
	if nil != err {
		panic("Unable to load posts.")
	}

	posts := helpers.GetPosts()

	muxer := chi.NewMux()

	muxer.Get("/users/{id}/posts", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "id")

		var userPosts []*model.Post

		for _, post := range posts {
			if post.UserID == userID {
				userPosts = append(userPosts, post)
			}
		}

		jsBytes, _ := json.Marshal(userPosts)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsBytes)
	})

	http.ListenAndServe(":"+port, muxer)
}
