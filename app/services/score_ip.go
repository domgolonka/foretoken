package services

import (
	"github.com/domgolonka/threatdefender/app"
	iputils "github.com/domgolonka/threatdefender/pkg/utils/ip"
)

func ScoreIP(app *app.App, ip string) (int8, error) {
	scoreCfg := app.Config.IP.Score

	var score int8
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
		if spamIP.Prefix > 0 {
			if iputils.ParseSubnet(ip, spamIP.IP, spamIP.Prefix) {
				score += scoreCfg.Spam.Yes
			}
		} else {
			score += scoreCfg.Spam.Yes
		}

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
	} else if score < 0 {
		score = 0
	}

	return score, nil

}