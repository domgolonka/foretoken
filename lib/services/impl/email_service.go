package impl

import (
	"context"
	"github.com/domgolonka/threatdefender/lib/services/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

var _ proto.EmailServiceServer = new(emailService)

type emailService struct {
}

func NewRegistryService() *emailService {
	return &emailService{}
}

func (e emailService) GetDisposableList(ctx context.Context, empty *empty.Empty) (*proto.GetDisposableListResponse, error) {
	panic("implement me")
}

func (e emailService) GetGenericList(ctx context.Context, empty *empty.Empty) (*proto.GetGenericListResponse, error) {
	panic("implement me")
}
