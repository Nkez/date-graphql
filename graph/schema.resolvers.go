package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Nkez/date-graphql/graph/generated"
	"github.com/Nkez/date-graphql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*string, error) {
	return r.userUseCase.CreateUser(ctx, &input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, update model.UserUpdate) (*string, error) {
	return r.userUseCase.UpdateUser(ctx, &update)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*string, error) {
	return r.userUseCase.DeleteUser(ctx, id)
}

func (r *queryResolver) GetEvent(ctx context.Context, id string) (*model.Event, error) {
	return r.eventUseCase.Get(ctx, id)
}

func (r *queryResolver) ListEvents(ctx context.Context, input model.Filter) (*model.EventsList, error) {
	return r.eventUseCase.List(ctx, &input)
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	return r.userUseCase.GetUser(ctx, id)
}

func (r *queryResolver) ListUsers(ctx context.Context, input model.UserFilter) (*model.UsersList, error) {
	return r.userUseCase.ListUsers(ctx, &input)
}

func (r *queryResolver) GetExel(ctx context.Context, input model.Filter) (*model.File, error) {
	return r.eventUseCase.GetExel(ctx, &input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
