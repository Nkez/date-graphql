package repository

import (
	"context"
	"github.com/Nkez/date-graphql/graph/model"
	date_protobuf "github.com/Nkez/date-protobuf"
)

type User interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	ListUser(ctx context.Context, filter *model.UserFilter) (*model.UsersList, error)
	CreateUser(ctx context.Context, url string, input *model.CreateUser) (*string, error)
	DeleteUser(ctx context.Context, id string) (*string, error)
	UpdateUser(ctx context.Context, url string, input *model.UserUpdate) (*string, error)
}

type UserRepository struct {
	user date_protobuf.UserServiceClient
}

func NewUserRepository(user date_protobuf.UserServiceClient) *UserRepository {
	return &UserRepository{user: user}
}

func (e *UserRepository) UpdateUser(ctx context.Context, url string, input *model.UserUpdate) (*string, error) {
	_, err := e.user.Update(ctx, &date_protobuf.User{
		Id:          *input.ID,
		FirstName:   *input.FirstName,
		LastName:    *input.LastName,
		Email:       *input.Email,
		UserName:    *input.UserName,
		Country:     *input.Country,
		MobilePhone: *input.MobilePhone,
		Photo:       url,
		Enabled:     *input.Enabled,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (e *UserRepository) DeleteUser(ctx context.Context, id string) (*string, error) {
	_, err := e.user.Delete(ctx, &date_protobuf.GetUser{Id: id})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (e *UserRepository) ListUser(ctx context.Context, filter *model.UserFilter) (*model.UsersList, error) {
	events, err := e.user.List(
		ctx,
		&date_protobuf.Filter{
			FirstName:  *filter.FirstName,
			SecondName: *filter.SecondName,
			Email:      *filter.Email,
			Role:       *filter.Role,
			Size:       int64(*filter.Size),
			Page:       int64(*filter.Page),
			Enabled:    *filter.Enabled,
		},
	)
	if err != nil {
		return nil, err
	}

	var resp []*model.User
	for _, ev := range events.Users {
		resp = append(resp, &model.User{
			ID:          ev.Id,
			FirstName:   ev.FirstName,
			LastName:    ev.LastName,
			Email:       ev.Email,
			UserName:    ev.UserName,
			MobilePhone: ev.Country,
			Country:     ev.MobilePhone,
			Photo:       "",
		})
	}

	return &model.UsersList{Users: resp}, err
}

func (e *UserRepository) GetUser(ctx context.Context, id string) (*model.User, error) {
	u, err := e.user.Get(ctx, &date_protobuf.GetUser{Id: id})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:          u.Id,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		UserName:    u.UserName,
		Country:     u.Country,
		MobilePhone: u.MobilePhone,
		Photo:       u.Photo,
		Enabled:     u.Enabled,
	}, nil
}

func (e *UserRepository) CreateUser(ctx context.Context, url string, input *model.CreateUser) (*string, error) {
	_, err := e.user.Create(ctx, &date_protobuf.CreateUser{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		UserName:    input.UserName,
		Country:     input.Country,
		MobilePhone: input.MobilePhone,
		Photo:       url,
		Enabled:     input.Enabled,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
