package usecases

import (
	"context"
	"github.com/Nkez/date-graphql/graph/model"
	"github.com/Nkez/date-graphql/internal/repository"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/martian/log"
	"github.com/spf13/viper"
)

type UserI interface {
	ListUsers(ctx context.Context, filter *model.UserFilter) (*model.UsersList, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.CreateUser) (*string, error)
	DeleteUser(ctx context.Context, id string) (*string, error)
	UpdateUser(ctx context.Context, input *model.UserUpdate) (*string, error)
}

type User struct {
	SMTPEmailRepository *repository.SMTPEmailRepository
	userUserRepository  *repository.UserRepository
	s3Client            *s3manager.Uploader
}

func NewUser(userUserRepository *repository.UserRepository, s3Client *s3manager.Uploader, SMTPEmailRepository *repository.SMTPEmailRepository) *User {
	return &User{userUserRepository: userUserRepository, s3Client: s3Client, SMTPEmailRepository: SMTPEmailRepository}
}

func (u *User) DeleteUser(ctx context.Context, id string) (*string, error) {
	return u.userUserRepository.DeleteUser(ctx, id)
}

func (u *User) ListUsers(ctx context.Context, filter *model.UserFilter) (*model.UsersList, error) {
	return u.userUserRepository.ListUser(ctx, filter)
}

func (u *User) GetUser(ctx context.Context, id string) (*model.User, error) {
	return u.userUserRepository.GetUser(ctx, id)
}

func (u *User) CreateUser(ctx context.Context, input *model.CreateUser) (*string, error) {
	output, err := u.s3Client.UploadWithContext(ctx, &s3manager.UploadInput{
		ACL:    aws.String(s3.ObjectCannedACLPublicRead),
		Body:   input.Photo.File,
		Bucket: aws.String(viper.GetString("bucket.photo")),
		Key:    aws.String(input.Photo.Filename),
	})
	if err != nil {
		return nil, err
	}
	url := output.Location
	_, err = u.userUserRepository.CreateUser(ctx, url, input)
	if err != nil {
		return nil, err
	}
	go func() {
		err = u.SMTPEmailRepository.Send(ctx, input.Email)
		if err != nil {
			log.Infof("Error with send email")
		}
	}()
	return &url, nil
}

func (u *User) UpdateUser(ctx context.Context, input *model.UserUpdate) (*string, error) {
	output, err := u.s3Client.UploadWithContext(ctx, &s3manager.UploadInput{
		ACL:    aws.String(s3.ObjectCannedACLPublicRead),
		Body:   input.Photo.File,
		Bucket: aws.String(viper.GetString("bucket.photo")),
		Key:    aws.String(input.Photo.Filename),
	})
	if err != nil {
		return nil, err
	}
	url := output.Location
	_, err = u.userUserRepository.UpdateUser(ctx, url, input)
	if err != nil {
		return nil, err
	}
	return &url, nil
}
