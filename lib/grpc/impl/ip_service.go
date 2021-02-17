package impl

import (
	"context"

	"github.com/domgolonka/threatdefender/app/services"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/domgolonka/threatdefender/app"

	"github.com/domgolonka/threatdefender/lib/grpc/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

var _ proto.IPServiceServer = new(ipService)

type ipService struct {
	app *app.App
}

func NewIPService(app *app.App) *ipService { //nolint
	return &ipService{app}
}

func (i ipService) GetProxyList(ctx context.Context, empty *empty.Empty) (*proto.GetProxyListResponse, error) {
	proxies, err := services.ProxyGetDBAll(i.app)
	if err != nil {
		return nil, err
	}
	arr := make([]*proto.Proxy, len(*proxies))
	for i, v := range *proxies {
		arr[i] = &proto.Proxy{
			Id:        uint32(v.ID),
			Ip:        v.IP,
			Port:      v.Port,
			Type:      v.Type,
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
		}
	}

	result := &proto.GetProxyListResponse{
		Proxies: arr,
	}

	return result, nil
}

func (i ipService) GetSpamList(ctx context.Context, empty *empty.Empty) (*proto.GetSpamListResponse, error) {
	spam, err := services.SpamGetDBAll(i.app)
	if err != nil {
		return nil, err
	}

	result := &proto.GetSpamListResponse{
		Spam: *spam,
	}

	return result, nil
}

func (i ipService) GetTorList(ctx context.Context, empty *empty.Empty) (*proto.GetTorListResponse, error) {
	tor, err := services.TorGetDBAll(i.app)
	if err != nil {
		return nil, err
	}

	result := &proto.GetTorListResponse{
		Tor: *tor,
	}

	return result, nil
}

func (i ipService) GetVPNList(ctx context.Context, empty *empty.Empty) (*proto.GetVPNListResponse, error) {
	vpn, err := services.VpnGetDBAll(i.app)
	if err != nil {
		return nil, err
	}

	result := &proto.GetVPNListResponse{
		Vpn: *vpn,
	}

	return result, nil
}
