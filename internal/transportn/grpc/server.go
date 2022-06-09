package grpc2

import (
	date_protobuf "github.com/Nkez/date-protobuf"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func NewClient() (date_protobuf.EventServiceClient, error) {
	conn, err := grpc.Dial(viper.GetString("grpc"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return date_protobuf.NewEventServiceClient(conn), nil
}
