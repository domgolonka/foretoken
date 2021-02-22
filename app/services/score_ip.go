package services

import (
	"github.com/domgolonka/threatdefender/app"
)

func ScoreIP(app *app.App, ip string) (uint8, error) {
	scoreCfg := app.Config.IP.Score

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

	vpnIP, err := app.VpnStore.FindByIP(ip)
	if err != nil {
		app.Logger.Error(err)
		return score, err
	}
	if proxyIP != nil {
		score += scoreCfg.Proxy.Yes
	} else {
		score += scoreCfg.Proxy.No
	}
	if spamIP != nil {
		score += scoreCfg.Spam.Yes
	} else {
		score += scoreCfg.Spam.No
	}
	if torIP != nil {
		score += scoreCfg.Tor.Yes
	} else {
		score += scoreCfg.Tor.No
	}
	if vpnIP != nil {
		score += scoreCfg.VPN.Yes
	} else {
		score += scoreCfg.VPN.No
	}
	if score > 100 {
		score = 100
	}

	return score, nil

}
