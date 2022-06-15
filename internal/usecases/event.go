package usecases

import (
	"context"
	"github.com/Nkez/date-graphql/graph/model"
	"github.com/Nkez/date-graphql/internal/repository"
)

type EventI interface {
	Get(ctx context.Context, id string) (*model.Event, error)
	List(ctx context.Context, filter *model.Filter) (*model.EventsList, error)
}

type Event struct {
	eventEventRepository *repository.EventRepository
}

func NewEvent(eventEventRepository *repository.EventRepository) *Event {
	return &Event{eventEventRepository: eventEventRepository}
}

func (e *Event) Get(ctx context.Context, id string) (*model.Event, error) {
	return e.eventEventRepository.Get(ctx, id)
}

func (e *Event) List(ctx context.Context, filter *model.Filter) (*model.EventsList, error) {
	return e.eventEventRepository.List(ctx, filter)
}
