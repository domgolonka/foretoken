package providers

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var OpenVPNUrls = []Urls{
	{Url: "https://www.ipvanish.com/software/configs/configs.zip", Typec: "OpenVPN", Format: OPENVPN},
	{Url: "https://downloads.nordcdn.com/configs/archives/servers/ovpn.zip", Typec: "OpenVPN", Format: OPENVPN},
	{Url: "https://s3-us-west-1.amazonaws.com/heartbleed/linux/linux-files.zip", Typec: "OpenVPN", Format: OPENVPN},
	{Url: "https://www.privateinternetaccess.com/openvpn/openvpn-ip.zip", Typec: "OpenVPN", Format: OPENVPN},
	{Url: "https://www.privateinternetaccess.com/openvpn/openvpn-tcp.zip", Typec: "OpenVPN", Format: OPENVPN},
	{Url: "https://s3.amazonaws.com/tunnelbear/linux/openvpn.zip", Typec: "OpenVPN", Format: OPENVPN},
	{Url: "https://torguard.net/downloads/OpenVPN-UDP.zip", Typec: "OpenVPN", Format: OPENVPN},
	{Url: "https://torguard.net/downloads/OpenVPN-TCP.zip", Typec: "OpenVPN", Format: OPENVPN},
	{Url: "https://vpn.hidemyass.com/vpn-config/vpn-configs.zip", Typec: "OpenVPN", Format: OPENVPN},
}

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

func (c *OpenVpn) Download(src Urls) ([]string, error) {
	hosts := []string{}
	c.logger.Debug("starting Download for " + src.Url)

	if src.Format == OPENVPN {

		res, err := http.Get(src.Url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		d, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var filenames []string

		r, err := zip.NewReader(bytes.NewReader(d), int64(len(d)))
		if err != nil {
			return filenames, err
		}

		for _, f := range r.File {

			if filepath.Ext(f.Name) == ".ovpn" && !strings.HasPrefix(f.Name, "__MACOSX") {
				var buf bytes.Buffer

				rc, err := f.Open()
				if err != nil {
					log.Fatal(err)
				}

				reRemote := regexp.MustCompile(`remote (\S+)`)

				_, err = io.Copy(&buf, rc)
				if err != nil {
					return nil, err
				}

				rc.Close()
				domainTmp := reRemote.FindStringSubmatch(string(buf.String()))
				var domainName = domainTmp[1]

				hosts = append(hosts, domainName)

			}

		}

	}

	return hosts, nil

}

func (c *OpenVpn) List() ([]string, error) {
	hosts := []string{}

	for i := 0; i < len(OpenVPNUrls); i++ {

		host, err := c.Download(OpenVPNUrls[i])
		if err != nil {
			return hosts, err
		}
		hosts = append(hosts, host...)
	}

	return hosts, nil
}

// Parses a 'remote' option into a struct
func getRemote(line string) (remote, error) {
	rmt := remote{}

	fields := strings.Fields(line)
	if len(fields) < 2 {
		return rmt, errors.New("unknown remote option")
	}
	isIP, err := regexp.MatchString(IPRegex, fields[1])
	if err != nil {
		return rmt, err
	}

	var ips []net.IP
	// Lookup ip address if remote is not an IP
	if !isIP {
		rmt.hostname = fields[1]
		ip4, err := net.LookupIP(fields[1])
		if err != nil {
			return rmt, err
		}

		for _, ip := range ip4 {
			if ip.To4() != nil {
				ips = append(ips, ip)
			}
		}
	} else {
		ips = append(ips, net.ParseIP(fields[1]))
	}
	if len(ips) == 0 {
		return remote{}, errors.New("can't resolve domain name")
	}

	for _, ip := range ips {
		rmt.ips = append(rmt.ips, ip.String())
	}

	// port is provided in remote option
	if len(fields) >= 3 {
		port, err := strconv.ParseUint(fields[2], 10, 32)
		if err != nil {
			return remote{}, err
		}
		rmt.port = uint(port)
	}

	// proto is provided in remote option
	if len(fields) >= 4 {
		rmt.proto = getProto(fields[3])
		if rmt.proto == "" {
			return remote{}, errors.New("unknown protocol")
		}
	}
	return rmt, nil
}

func getProto(p string) Proto {
	switch p {
	case "udp":
		return udp
	case "tcp":
		return tcp
	default:
		return ""
	}
}
