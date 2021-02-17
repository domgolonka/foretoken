package impl

import (
	"context"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"

	"github.com/domgolonka/threatdefender/lib/grpc/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

var _ proto.EmailServiceServer = new(emailService)

type emailService struct {
	app *app.App
}

func NewEmailService(app *app.App) *emailService { //nolint
	return &emailService{app}
}

func (e emailService) GetDisposableList(ctx context.Context, empty *empty.Empty) (*proto.GetEmailListResponse, error) {
	emails, err := services.DisposableGetDBAll(e.app)
	if err != nil {
		return nil, err
	}

	result := &proto.GetEmailListResponse{
		Emails: *emails,
	}

	return result, nil
}

func (e emailService) GetGenericList(ctx context.Context, empty *empty.Empty) (*proto.GetEmailListResponse, error) {
	genericList, err := services.GenericGetAll(e.app)
	if err != nil {
		return nil, err
	}

	result := &proto.GetEmailListResponse{
		Emails: *genericList,
	}

	return result, nil
}

func (e emailService) GetSpamList(ctx context.Context, e2 *empty.Empty) (*proto.GetEmailListResponse, error) {
	spamList, err := services.SpamEmailGetDBAll(e.app)
	if err != nil {
		return nil, err
	}

	result := &proto.GetEmailListResponse{
		Emails: *spamList,
	}

	return result, nil
}

func (e emailService) GetFreeEmailList(ctx context.Context, e2 *empty.Empty) (*proto.GetEmailListResponse, error) {
	freeEmailList, err := services.FreeEmailGetDBAll(e.app)
	if err != nil {
		return nil, err
	}

	result := &proto.GetEmailListResponse{
		Emails: *freeEmailList,
	}

	return result, nil
}
