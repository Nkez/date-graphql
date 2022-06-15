//go:generate go run github.com/99designs/gqlgen generate

package graph

import "github.com/Nkez/date-graphql/internal/usecases"

type Resolver struct {
	eventUseCase *usecases.Event
}

func NewResolver(eventUseCase *usecases.Event) *Resolver {
	return &Resolver{eventUseCase: eventUseCase}
}
