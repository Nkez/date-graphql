package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/Nkez/date-graphql/graph/generated"
	"github.com/Nkez/date-graphql/graph/model"
)

func (r *queryResolver) GetEvent(ctx context.Context, id string) (*model.Event, error) {
	return r.eventUseCase.Get(ctx, id)
}

func (r *queryResolver) ListEvents(ctx context.Context, input model.Filter) (*model.EventsList, error) {
	return r.eventUseCase.List(ctx, &input)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
