package services

import (
	"github.com/domgolonka/threatdefender/app"
)

func ScoreIP(app *app.App, ip string) (uint8, error) {
	var score uint8
	score = 0
	proxyIP, err := app.ProxyStore.FindByIP(ip)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	spamIP, err := app.SpamStore.FindByIP(ip)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	torIP, err := app.TorStore.FindByIP(ip)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}

	vpnIP, err := app.VpnStore.FindByURL(ip)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	if proxyIP != nil {
		score += 30
	}
	if spamIP != nil {
		score += 50
	}
	if torIP != nil {
		score += 15
	}
	if vpnIP != nil {
		score += 10
	}
	if score > 100 {
		score = 100
	}

	return score, nil

}
