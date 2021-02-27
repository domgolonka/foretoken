package impl

import (
	"context"

	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/app/services"

	"github.com/domgolonka/foretoken/lib/grpc/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

var _ proto.EmailServiceServer = new(emailService)

type emailService struct {
	app *app.App
}

func NewEmailService(app *app.App) *emailService { //nolint
	return &emailService{app}
}

func (e emailService) GetScore(ctx context.Context, request *proto.EmailRequest) (*proto.GetEmailScoreResponse, error) {
	emailSrv := services.Email{}
	emailSrv.Calculate(e.app, request.GetEmail())

	response, err := emailSrv.ScoreEmail()
	if err != nil {
		return nil, err
	}
	score := &proto.GetEmailScoreResponse{
		Score: uint32(response),
	}
	return score, nil
}

func (e emailService) GetEmail(ctx context.Context, request *proto.EmailRequest) (*proto.GetEmailResponse, error) {
	emailSrv := services.Email{}
	emailSrv.Calculate(e.app, request.GetEmail())

	response, err := emailSrv.EmailService()
	if err != nil {
		return nil, err
	}
	score := &proto.GetEmailResponse{
		Success:    response.Success,
		Valid:      response.Valid,
		Disposable: response.Disposable,
		RecentSpam: response.RecentSpam,
		Free:       response.Free,
		Generic:    response.Generic,
		Score:      uint32(response.Score),
		Domain: &proto.GetEmailResponse_Domain{
			CreatedAt:      response.Domain.CreatedDate,
			ExpirationDate: response.Domain.ExpirationDate,
		},
	}
	return score, nil
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
