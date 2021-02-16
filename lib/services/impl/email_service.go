package impl

import (
	"context"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"

	"github.com/domgolonka/threatdefender/lib/services/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

var _ proto.EmailServiceServer = new(emailService)

type emailService struct {
	app *app.App
}

func NewEmailService(app *app.App) *emailService { //nolint
	return &emailService{app}
}

func (e emailService) GetDisposableList(ctx context.Context, empty *empty.Empty) (*proto.GetDisposableListResponse, error) {
	emails, err := services.DisposableGetDBAll(e.app)
	if err != nil {
		return nil, err
	}

	result := &proto.GetDisposableListResponse{
		Emails: *emails,
	}

	return result, nil
}

func (e emailService) GetGenericList(ctx context.Context, empty *empty.Empty) (*proto.GetGenericListResponse, error) {
	genericList, err := services.GenericGetAll(e.app)
	if err != nil {
		return nil, err
	}

	result := &proto.GetGenericListResponse{
		GenericList: *genericList,
	}

	return result, nil
}

func (e emailService) GetFreeEmailList(ctx context.Context, e2 *empty.Empty) (*proto.GetFreeEmailListResponse, error) {
	freeEmailList, err := services.FreeEmailGetDBAll(e.app)
	if err != nil {
		return nil, err
	}

	result := &proto.GetFreeEmailListResponse{
		Emails: *freeEmailList,
	}

	return result, nil
}
