package providers

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/domgolonka/threatdefender/app/models"

	"github.com/domgolonka/threatdefender/app/entity"
	"github.com/sirupsen/logrus"
)

//var OpenVPNURLs = []URLs{
//	{URL: "https://www.ipvanish.com/software/configs/configs.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://downloads.nordcdn.com/configs/archives/servers/ovpn.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://s3-us-west-1.amazonaws.com/heartbleed/linux/linux-files.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://www.privateinternetaccess.com/openvpn/openvpn-ip.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://www.privateinternetaccess.com/openvpn/openvpn-tcp.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://s3.amazonaws.com/tunnelbear/linux/openvpn.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://torguard.net/downloads/OpenVPN-UDP.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://torguard.net/downloads/OpenVPN-TCP.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://vpn.hidemyass.com/vpn-config/vpn-configs.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://support.vyprvpn.com/hc/article_attachments/360052617332/Vypr_OpenVPN_20200320.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://s3-us-west-1.amazonaws.com/heartbleed/windows/New+OVPN+Files.zip", Typec: "OpenVPN", Format: OPENVPN},
//	//{URL: "http://www.digibit.tv/certs/Certificates.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://www.limevpn.com/downloads/OpenVPN-Config-1194.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://raw.githubusercontent.com/en1gmascr1pts/vpnconfigs/master/Windscribe.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://raw.githubusercontent.com/en1gmascr1pts/vpnconfigs/master/VPNUnlimited.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://raw.githubusercontent.com/en1gmascr1pts/vpnconfigs/master/SmartDNSProxy.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://raw.githubusercontent.com/en1gmascr1pts/vpnconfigs/master/ProtonVPN.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://raw.githubusercontent.com/en1gmascr1pts/vpnconfigs/master/SmartyDNS.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://raw.githubusercontent.com/en1gmascr1pts/vpnconfigs/master/ExpressVPN.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://raw.githubusercontent.com/en1gmascr1pts/vpnconfigs/master/ibVPN.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://raw.githubusercontent.com/en1gmascr1pts/vpnconfigs/master/AirVPN.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://raw.githubusercontent.com/en1gmascr1pts/vpnconfigs/master/MullvadVPN.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://monstervpn.tech/ovpn_configuration.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://www.goldenfrog.com/openvpn/VyprVPNOpenVPNFiles.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://freevpnme.b-cdn.net/FreeVPN.me-OpenVPN-Bundle-July-2020.zip", Typec: "OpenVPN", Format: OPENVPN},
//	{URL: "https://github.com/cryptostorm/cryptostorm_client_configuration_files/archive/master.zip", Typec: "OpenVPN", Format: OPENVPN},
//}

type OpenVpn struct {
	logger logrus.FieldLogger
}

func NewOpenVpn(logger logrus.FieldLogger) *OpenVpn {
	logger.Debug("starting OpenVPN")
	return &OpenVpn{
		logger,
	}
}
func (*OpenVpn) Name() string {
	return "openvpn"
}

func (c *OpenVpn) Download(src *entity.Feed) ([]models.Vpn, error) {
	hosts := []models.Vpn{}
	c.logger.Debug("starting Download for " + src.URL)

	if src.Format == "OPENVPN" {
		res, err := http.Get(src.URL)
		if err != nil {
			c.logger.Error(err)
		}
		defer func() {
			err = res.Body.Close()
			if err != nil {
				c.logger.Error(err)
			}
		}()
		d, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.logger.Error(err)
		}

		r, err := zip.NewReader(bytes.NewReader(d), int64(len(d)))
		if err != nil {
			return nil, err
		}

		for _, f := range r.File {
			if filepath.Ext(f.Name) == ".ovpn" && !strings.HasPrefix(f.Name, "__MACOSX") {
				var buf bytes.Buffer

				rc, err := f.Open()
				if err != nil {
					c.logger.Error(err)
				}

				reRemote := regexp.MustCompile(`remote (\S+)`)

				_, err = io.Copy(&buf, rc) //nolint
				if err != nil {
					return nil, err
				}

				rc.Close()
				domainTmp := reRemote.FindStringSubmatch(buf.String())
				vpn := models.Vpn{
					IP:   domainTmp[1],
					Type: "openvpn",
				}
				hosts = append(hosts, vpn)
			}
		}
	}

	return hosts, nil

}

func (c *OpenVpn) List() ([]models.Vpn, error) {
	hosts := []models.Vpn{}
	f := entity.Feed{
		Logger: c.logger,
	}
	feed, err := f.ReadFile("ip_openvpn.json")
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(feed); i++ {
		host, err := c.Download(feed[i])
		if err != nil {
			return hosts, err
		}
		hosts = append(hosts, host...)
	}

	return hosts, nil
}

// Parses a 'remote' option into a struct
//func getRemote(line string) (remote, error) {
//	rmt := remote{}
//
//	fields := strings.Fields(line)
//	if len(fields) < 2 {
//		return rmt, errors.New("unknown remote option")
//	}
//	isIP, err := regexp.MatchString(IPRegex, fields[1])
//	if err != nil {
//		return rmt, err
//	}
//
//	var ips []net.IP
//	// Lookup ip address if remote is not an IP
//	if !isIP {
//		rmt.hostname = fields[1]
//		ip4, err := net.LookupIP(fields[1])
//		if err != nil {
//			return rmt, err
//		}
//
//		for _, ip := range ip4 {
//			if ip.To4() != nil {
//				ips = append(ips, ip)
//			}
//		}
//	} else {
//		ips = append(ips, net.ParseIP(fields[1]))
//	}
//	if len(ips) == 0 {
//		return remote{}, errors.New("can't resolve domain name")
//	}
//
//	for _, ip := range ips {
//		rmt.ips = append(rmt.ips, ip.String())
//	}
//
//	// port is provided in remote option
//	if len(fields) >= 3 {
//		port, err := strconv.ParseUint(fields[2], 10, 32)
//		if err != nil {
//			return remote{}, err
//		}
//		rmt.port = uint(port)
//	}
//
//	// proto is provided in remote option
//	if len(fields) >= 4 {
//		rmt.proto = getProto(fields[3])
//		if rmt.proto == "" {
//			return remote{}, errors.New("unknown protocol")
//		}
//	}
//	return rmt, nil
//}

//func getProto(p string) Proto {
//	switch p {
//	case "udp":
//		return udp
//	case "tcp":
//		return tcp
//	default:
//		return ""
//	}
//}
