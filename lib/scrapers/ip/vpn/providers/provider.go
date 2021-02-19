package providers

import (
	"net"
	"net/http"
	"time"
)

type Format int

const (
	IPRegex = "^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$"
)
const (
	OPENVPN Format = iota
	HTML
)

//type URLs struct {
//	URL    string
//	Typec  string
//	Format Format
//}

//var HTMLURLs = []URLs{
//
//	{URL: "https://support.goldenfrog.com/hc/en-us/articles/360011055671-What-are-the-VyprVPN-server-addresses-", Typec: "PPTP-L2TP", Format: HTML},
//}

func NewTransport() *http.Transport {
	return &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
			DualStack: true,
		}).DialContext,
		TLSHandshakeTimeout:   3 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     true,
	}
}

func NewClient() *http.Client {
	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: NewTransport(),
	}
}

//type Proto string

//const (
//	udp Proto = "udp"
//	tcp Proto = "tcp"
//)

//type remote struct {
//	ips      []string
//	hostname string
//	port     uint
//	proto    Proto
//}
