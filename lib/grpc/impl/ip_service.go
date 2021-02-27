package impl

import (
	"context"

	"github.com/domgolonka/foretoken/app/services"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/domgolonka/foretoken/app"

	"github.com/domgolonka/foretoken/lib/grpc/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

var _ proto.IPServiceServer = new(ipService)

type ipService struct {
	app *app.App
}

func NewIPService(app *app.App) *ipService { //nolint
	return &ipService{app}
}

func (i ipService) GetScore(ctx context.Context, request *proto.IPRequest) (*proto.GetIPScoreResponse, error) {
	ipSrv := services.IP{}
	ipSrv.Calculate(i.app, request.GetIp())

	response, err := ipSrv.ScoreIP()
	if err != nil {
		return nil, err
	}
	score := &proto.GetIPScoreResponse{
		Score: uint32(response),
	}

	return score, nil

}

func (i ipService) GetIP(ctx context.Context, request *proto.IPRequest) (*proto.GetIPResponse, error) {
	ipSrv := services.IP{}
	ipSrv.Calculate(i.app, request.GetIp())

	response, err := ipSrv.IPService()
	if err != nil {
		return nil, err
	}
	score := &proto.GetIPResponse{
		Success:      response.Success,
		Proxy:        response.Proxy,
		ISP:          response.ISP,
		Organization: response.Organization,
		ASN:          uint32(response.ASN),
		Hostname:     response.Hostname,
		CountryCode:  response.CountryCode,
		City:         response.City,
		PostalCode:   response.PostalCode,
		Latitude:     float32(response.Latitude),
		Longitude:    float32(response.Longitude),
		Timezone:     response.Timezone,
		Vpn:          response.Vpn,
		Tor:          response.Tor,
		RecentAbuse:  response.RecentAbuse,
		Score:        uint32(response.Score),
	}

	return score, nil
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
