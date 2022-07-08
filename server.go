package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Nkez/date-graphql/graph"
	"github.com/Nkez/date-graphql/graph/generated"
	"github.com/Nkez/date-graphql/internal/repository"
	"github.com/Nkez/date-graphql/internal/usecases"
	date_protobuf "github.com/Nkez/date-protobuf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	if err := ConfigInit(); err != nil {
		log.Fatalf("error instaling configs: %s", err.Error())
	}
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(viper.GetString("s3.accessKeyID"), viper.GetString("s3.secretAccessKey"), ""),
		Endpoint:         aws.String(viper.GetString("s3.endpoint")),
		Region:           aws.String(viper.GetString("s3.region")),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	sess := session.Must(session.NewSession(s3Config))
	awsClient := s3manager.NewUploader(sess)

	eventConn, err := grpc.Dial(viper.GetString("grpc.event"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error grpc connect : %s", err.Error())
	}
	userConn, err := grpc.Dial(viper.GetString("grpc.user"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error grpc connect : %s", err.Error())
	}
	email := repository.NewSMTPEmailRepository(viper.GetString("smtp.email"), viper.GetString("smtp.password"),
		viper.GetString("smtp.host"), viper.GetInt("smtp.port"))
	user := date_protobuf.NewUserServiceClient(userConn)
	event := date_protobuf.NewEventServiceClient(eventConn)
	eventRepos := repository.NewEventRepository(event)
	userRepos := repository.NewUserRepository(user)
	useCaseEvent := usecases.NewEvent(eventRepos, awsClient)
	useCaseUser := usecases.NewUser(userRepos, awsClient, email)
	resolver := graph.NewResolver(useCaseEvent, useCaseUser)
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
