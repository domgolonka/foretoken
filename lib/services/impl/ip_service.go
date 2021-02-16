package impl

import (
	"context"

	"github.com/domgolonka/threatdefender/lib/services/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

var _ proto.IPServiceServer = new(ipService)

type ipService struct {
}

func NewIPService() *ipService {
	return &ipService{}
}

func (i ipService) GetProxyList(ctx context.Context, empty *empty.Empty) (*proto.GetProxyListResponse, error) {
	panic("implement me")
}

func (i ipService) GetSpamList(ctx context.Context, empty *empty.Empty) (*proto.GetSpamListResponse, error) {
	panic("implement me")
}

func (i ipService) GetTorList(ctx context.Context, empty *empty.Empty) (*proto.GetTorListResponse, error) {
	panic("implement me")
}

func (i ipService) GetVPNList(ctx context.Context, empty *empty.Empty) (*proto.GetVPNListResponse, error) {
	panic("implement me")
}
