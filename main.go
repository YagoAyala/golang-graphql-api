package main

import (
	"fmt"
	"log"
	"net/http"

	"multiplier/website/app"
	"multiplier/website/generated"
	"multiplier/website/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

func init() {
	_, err := app.InitConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
}

func main() {
	db, err := app.InitDB()
	if err != nil {
		log.Fatal("cannot load database:", err)
	}
	defer db.Close()

	router := chi.NewRouter()

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolver.Resolver{}},
		),
	)

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", app.Config.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", app.Config.ServerPort), router))
}
