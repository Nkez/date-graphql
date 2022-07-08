package usecases

import (
	"context"
	"fmt"
	"github.com/Nkez/date-graphql/graph/model"
	"github.com/Nkez/date-graphql/internal/repository"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

type EventI interface {
	Get(ctx context.Context, id string) (*model.Event, error)
	List(ctx context.Context, filter *model.Filter) (*model.EventsList, error)
	GetExel(ctx context.Context, filter *model.Filter) (*model.File, error)
}

type Event struct {
	eventEventRepository *repository.EventRepository
	s3Client             *s3manager.Uploader
}

func NewEvent(eventEventRepository *repository.EventRepository, s3Client *s3manager.Uploader) *Event {
	return &Event{eventEventRepository: eventEventRepository, s3Client: s3Client}
}

func (e *Event) Get(ctx context.Context, id string) (*model.Event, error) {
	return e.eventEventRepository.Get(ctx, id)
}

func (e *Event) List(ctx context.Context, filter *model.Filter) (*model.EventsList, error) {
	return e.eventEventRepository.List(ctx, filter)
}

func (e *Event) GetExel(ctx context.Context, filter *model.Filter) (*model.File, error) {
	var fileStruct model.File
	events, err := e.eventEventRepository.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	categories := map[string]string{
		"B1": "ID", "C1": "Type Request", "D1": "Browser", "E1": "OS", "F1": "Device", "G1": "City", "H1": "Country", "I1": "Create At"}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for i := 0; i < len(events.Events); i++ {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), i+1)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), events.Events[i].ID)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), events.Events[i].Request)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), events.Events[i].Browser)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), events.Events[i].Os)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+2), events.Events[i].Device)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+2), events.Events[i].City)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+2), events.Events[i].Country)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+2), events.Events[i].CreateAt)
	}
	s, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	output, err := e.s3Client.Upload(&s3manager.UploadInput{
		ACL:    aws.String(s3.ObjectCannedACLPublicRead),
		Body:   s,
		Bucket: aws.String(viper.GetString("bucket.exel")),
		Key:    aws.String("tmp" + "/" + "events" + time.Now().String() + ".xlsx"),
	})
	if err != nil {
		return nil, err
	}
	fileStruct.URL = output.Location
	fmt.Println(fileStruct.URL)
	return &fileStruct, nil
}
