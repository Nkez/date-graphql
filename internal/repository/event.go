package repository

import (
	"context"
	"github.com/Nkez/date-graphql/graph/model"
	date_protobuf "github.com/Nkez/date-protobuf"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type Event interface {
	Get(ctx context.Context, id string) (*model.Event, error)
	List(ctx context.Context, filter *model.Filter) (*model.EventsList, error)
	GetExel(ctx context.Context, filter *model.Filter) (*model.File, error)
}

type EventRepository struct {
	event date_protobuf.EventServiceClient
}

func NewEventRepository(event date_protobuf.EventServiceClient) *EventRepository {
	return &EventRepository{event: event}
}

func (e *EventRepository) Get(ctx context.Context, id string) (*model.Event, error) {
	ev, err := e.event.Get(ctx,
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
		CreateAt: ev.CreatedAt.AsTime().String(),
	}, nil
}

func (e *EventRepository) List(ctx context.Context, filter *model.Filter) (*model.EventsList, error) {
	events, err := e.event.List(
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
			Os:       ev.Os,
			Device:   ev.Device,
			City:     ev.City,
			Country:  ev.Country,
			CreateAt: ev.CreatedAt.AsTime().String(),
		})
	}

	return &model.EventsList{Events: resp}, err
}
