package services

import (
	"errors"
	"net"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/entity"
	iputils "github.com/domgolonka/threatdefender/pkg/utils/ip"
)

func IPService(app *app.App, ipaddress string) (*entity.IPAddressResponse, error) {
	ip, _, err := net.ParseCIDR(ipaddress)
	if err != nil {
		return nil, errors.New("not a real IP address")
	}
	ipresponse := &entity.IPAddressResponse{
		Proxy:       false,
		Tor:         false,
		Vpn:         false,
		RecentAbuse: false,
		Score:       0,
	}

	if app.Maxmind != nil {
		err = app.Maxmind.GetIPdata(ip)
		if err != nil {
			app.Logger.Error(err)
		}
	}

	proxy, err := app.ProxyStore.FindByIP(ipaddress)
	if err != nil {
		app.Logger.Error(err)
	}
	if proxy != nil {
		ipresponse.Proxy = true
	}
	tor, err := app.TorStore.FindByIP(ipaddress)
	if err != nil {
		app.Logger.Error(err)
	}
	if tor != nil {
		ipresponse.Tor = true
	}
	vpn, err := app.VpnStore.FindByIP(ipaddress)
	if err != nil {
		app.Logger.Error(err)
	}
	if vpn != nil {
		ipresponse.Vpn = true
	}
	spam, err := app.SpamStore.FindByIP(ipaddress)
	if err != nil {
		app.Logger.Error(err)
	}
	if spam != nil {
		if spam.Prefix > 0 {
			if iputils.ParseSubnet(ipaddress, spam.IP, spam.Prefix) {
				ipresponse.RecentAbuse = true
			}
		} else {
			ipresponse.RecentAbuse = true
		}

	}
	score, err := ScoreIP(app, ipaddress)
	if err != nil {
		app.Logger.Error(err)
	} else {
		ipresponse.Score = score
	}
	return ipresponse, nil

}
