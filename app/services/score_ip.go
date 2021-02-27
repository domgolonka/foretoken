package services

import (
	iputils "github.com/domgolonka/foretoken/pkg/utils/ip"
)

func (i *IP) ScoreIP() (int8, error) {
	scoreCfg := i.app.Config.IP.Score

	var score int8
	score = 0

	if i.Proxy != nil {
		score += scoreCfg.Proxy.Yes
	} else {
		score += scoreCfg.Proxy.No
	}
	if i.Spam != nil {
		if i.Spam.Prefix > 0 {
			if iputils.ParseSubnet(i.IPAddress, i.Spam.IP, i.Spam.Prefix) {
				score += scoreCfg.Spam.Yes
			}
		} else {
			score += scoreCfg.Spam.Yes
		}

	} else {
		score += scoreCfg.Spam.No
	}
	if i.Tor != nil {
		score += scoreCfg.Tor.Yes
	} else {
		score += scoreCfg.Tor.No
	}
	if i.Vpn != nil {
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
