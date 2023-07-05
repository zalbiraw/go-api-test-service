package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/zalbiraw/go-api-test-service/services/rest/posts/helpers"
	"github.com/zalbiraw/go-api-test-service/services/rest/posts/model"
)

const defaultPort = "3102"

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

	log.Printf("connect to http://localhost:%s/ to test API", port)
	log.Fatal(http.ListenAndServe(":"+port, muxer))
}
