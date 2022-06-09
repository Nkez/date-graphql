package repository

import (
	"context"
	"github.com/Nkez/date-graphql/graph/model"
	grpc2 "github.com/Nkez/date-graphql/internal/transportn/grpc"
	date_protobuf "github.com/Nkez/date-protobuf"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type Event interface {
	Get(ctx context.Context, id string) (*model.Event, error)
	List(ctx context.Context, filter *model.Filter) (*model.EventsList, error)
}

type EventRepository struct {
}

func (e EventRepository) Get(ctx context.Context, id string) (*model.Event, error) {
	client, err := grpc2.NewClient()
	if err != nil {
		return nil, err
	}
	ev, err := client.Get(
		ctx,
		&date_protobuf.GetEvent{
			Id: id,
		},
	)
	if err != nil {
		return nil, err
	}
	return &model.Event{
		ID:       ev.Id,
		Request:  ev.TypeRequest,
		Browser:  ev.Browser,
		Os:       ev.City,
		Device:   ev.Device,
		City:     ev.City,
		Country:  ev.Country,
		CreateAt: ev.CreatedAt.String(),
	}, nil
}

func (e EventRepository) List(ctx context.Context, filter *model.Filter) (*model.EventsList, error) {
	client, err := grpc2.NewClient()
	events, _ := client.List(
		ctx,
		&date_protobuf.FilterEvent{
			PageNumber: &wrappers.UInt64Value{
				Value: uint64(*filter.Page),
			},
			PageSize: &wrappers.UInt64Value{
				Value: uint64(*filter.Size),
			},
		},
	)
	if err != nil {
		return nil, err
	}

	var resp []*model.Event
	for _, ev := range events.Event {
		resp = append(resp, &model.Event{
			ID:       ev.Id,
			Request:  ev.TypeRequest,
			Browser:  ev.Browser,
			Os:       ev.City,
			Device:   ev.Device,
			City:     ev.City,
			Country:  ev.Country,
			CreateAt: ev.CreatedAt.String(),
		})
	}

	return &model.EventsList{Events: resp}, err
}
