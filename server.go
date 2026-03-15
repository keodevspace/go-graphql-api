package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/keodevspace/go-graphql-api/graph"
	"github.com/keodevspace/go-graphql-api/internal"
)

func main() {
	// Inicializa o "Banco de Dados" em memória
	db := internal.NewDB()

	// Injeta o banco no Resolver (Dependency Injection manual)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{Store: db},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("Connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
