package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Nkez/date-graphql/graph"
	"github.com/Nkez/date-graphql/graph/generated"
	"github.com/Nkez/date-graphql/internal/repository"
	"github.com/Nkez/date-graphql/internal/usecases"
	date_protobuf "github.com/Nkez/date-protobuf"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	if err := ConfigInit(); err != nil {
		log.Fatalf("error instaling configs: %s", err.Error())
	}
	dateConn, err := grpc.Dial(viper.GetString("grpc"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error grpc connect : %s", err.Error())
	}
	client := date_protobuf.NewEventServiceClient(dateConn)
	repos := repository.NewEventRepository(client)
	useCase := usecases.NewEvent(repos)
	resolver := graph.NewResolver(useCase)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))
}

func ConfigInit() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
