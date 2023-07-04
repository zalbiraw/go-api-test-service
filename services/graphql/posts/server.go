//go:generate go run -mod=mod github.com/99designs/gqlgen
package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/zalbiraw/go-api-test-service/services/graphql/posts/graph"
	"github.com/zalbiraw/go-api-test-service/services/graphql/posts/helpers"
)

const defaultPort = "4102"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := helpers.LoadPosts()

	if nil != err {
		panic("Unable to load posts.")
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
