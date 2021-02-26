package services

import (
	"net"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/entity"
	iputils "github.com/domgolonka/threatdefender/pkg/utils/ip"
)

func IPService(app *app.App, ipaddress string) (*entity.IPAddressResponse, error) {

	ipresponse := &entity.IPAddressResponse{
		Proxy:       false,
		Tor:         false,
		Vpn:         false,
		RecentAbuse: false,
		CountryCode: "",
		Timezone:    "",
		City:        "",
		PostalCode:  "",
		Score:       0,
	}

	if app.Maxmind != nil {
		ip := net.ParseIP(ipaddress)
		// country
		country, err := app.Maxmind.GetCountry(ip)
		if err != nil {
			app.Logger.Error(err)
		}
		ipresponse.CountryCode = country
		postal, timezone, city, lat, long, err := app.Maxmind.GetCityData(ip)
		if err != nil {
			app.Logger.Error(err)
		}
		ipresponse.PostalCode = postal
		ipresponse.City = city
		ipresponse.Timezone = timezone
		ipresponse.Longitude = long
		ipresponse.Latitude = lat
		organization, asn, err := app.Maxmind.GetASN(ip)
		if err != nil {
			app.Logger.Error(err)
		}
		ipresponse.ASN = asn
		ipresponse.Organization = organization
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
	ipresponse.Success = true
	return ipresponse, nil

}
