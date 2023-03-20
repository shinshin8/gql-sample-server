package main

import (
	"database/sql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinshin8/gql-sample-server/graph"
	"github.com/shinshin8/gql-sample-server/graph/services"
	"github.com/shinshin8/gql-sample-server/internal"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "gql", "gql", "mysql", "3305", "gql"))

	if err != nil {
		log.Fatal(err)
	}

	service := services.New(db)

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{
		Srv: service,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
