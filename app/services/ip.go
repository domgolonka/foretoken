package services

import (
	"net"

	"github.com/domgolonka/foretoken/app/models"

	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/app/entity"
	iputils "github.com/domgolonka/foretoken/pkg/utils/ip"
)

type IP struct {
	app          *app.App
	IPAddress    string
	Proxy        *models.Proxy
	Tor          *models.Tor
	Vpn          *models.Vpn
	Spam         *models.Spam
	CountryCode  string
	Timezone     string
	City         string
	PostalCode   string
	ASN          uint
	Organization string
	Longitude    float64
	Latitude     float64
}

func (i *IP) IPService() (*entity.IPAddressResponse, error) {

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

	score, err := i.ScoreIP()
	if err != nil {
		i.app.Logger.Error(err)
	} else {
		ipresponse.Score = score
	}
	if i.Proxy != nil {
		ipresponse.Proxy = true
	}
	if i.Tor != nil {
		ipresponse.Tor = true
	}
	if i.Vpn != nil {
		ipresponse.Vpn = true
	}
	if i.Spam != nil {
		ipresponse.RecentAbuse = true
	}
	ipresponse.CountryCode = i.CountryCode
	ipresponse.City = i.City
	ipresponse.PostalCode = i.PostalCode
	ipresponse.Timezone = i.Timezone
	ipresponse.Success = true
	return ipresponse, nil

}

func (i *IP) Calculate(app *app.App, ipaddress string) {
	i.app = app
	i.IPAddress = ipaddress
	vpn, err := app.VpnStore.FindByIP(ipaddress)
	if err != nil {
		app.Logger.Error(err)
	}
	if vpn != nil {
		i.Vpn = vpn
	}
	spam, err := app.SpamStore.FindByIP(ipaddress)
	if err != nil {
		app.Logger.Error(err)
	}
	if spam != nil {
		if spam.Prefix > 0 {
			if iputils.ParseSubnet(ipaddress, spam.IP, spam.Prefix) {
				i.Spam = spam
			}
		} else {
			i.Spam = spam
		}

	}

	if app.Maxmind != nil {
		ip := net.ParseIP(ipaddress)
		// country
		country, err := app.Maxmind.GetCountry(ip)
		if err != nil {
			app.Logger.Error(err)
		}
		i.CountryCode = country
		postal, timezone, city, lat, long, err := app.Maxmind.GetCityData(ip)
		if err != nil {
			app.Logger.Error(err)
		}
		i.PostalCode = postal
		i.City = city
		i.Timezone = timezone
		i.Longitude = long
		i.Latitude = lat
		organization, asn, err := app.Maxmind.GetASN(ip)
		if err != nil {
			app.Logger.Error(err)
		}
		i.ASN = asn
		i.Organization = organization
	}

	proxy, err := app.ProxyStore.FindByIP(ipaddress)
	if err != nil {
		app.Logger.Error(err)
	}
	if proxy != nil {
		i.Proxy = proxy
	}
	tor, err := app.TorStore.FindByIP(ipaddress)
	if err != nil {
		app.Logger.Error(err)
	}
	if tor != nil {
		i.Tor = tor
	}

}
