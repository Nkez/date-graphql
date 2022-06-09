package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Nkez/date-graphql/graph"
	"github.com/Nkez/date-graphql/graph/generated"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	if err := ConfigInit(); err != nil {
		log.Fatalf("error instaling configs: %s", err.Error())
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to for GraphQL playground")
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))
}

func ConfigInit() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
