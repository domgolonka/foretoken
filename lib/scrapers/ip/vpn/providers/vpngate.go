package providers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/sirupsen/logrus"
)

type vpngatecsv struct {
	HostName                string `csv:"HostName"`
	IP                      string `csv:"IP"`
	Score                   string `csv:"Score"`
	Ping                    string `csv:"Ping"`
	Speed                   string `csv:"Speed"`
	CountryLong             string `csv:"CountryLong"`
	CountryShort            string `csv:"CountryShort"`
	NumVpnSessions          string `csv:"NumVpnSessions"`
	Uptime                  string `csv:"Uptime"`
	TotalUsers              string `csv:"TotalUsers"`
	TotalTraffic            string `csv:"TotalTraffic"`
	LogType                 string `csv:"LogType"`
	Operator                string `csv:"Operator"`
	Message                 string `csv:"Message"`
	OpenVPNConfigDataBase64 string `csv:"OpenVPN_ConfigData_Base64"`
}

type VPNGate struct {
	hosts      []string
	logger     logrus.FieldLogger
	lastUpdate time.Time
}

const vpngateURL = "https://www.vpngate.net/api/iphone/"

func NewVPNGate(logger logrus.FieldLogger) *VPNGate {
	logger.Debug("starting SlickVPN")
	return &VPNGate{
		logger: logger,
	}
}
func (*VPNGate) Name() string {
	return "vpngate"
}

func (c *VPNGate) Load(body []byte) ([]string, error) {
	// don't need to update this more than once a day!
	if time.Now().Unix() >= c.lastUpdate.Unix()+(82800) {
		c.hosts = make([]string, 0)
	}

	if len(c.hosts) != 0 {
		return c.hosts, nil
	}
	var ips []*vpngatecsv
	var err error
	if body == nil {
		if body, err = c.MakeRequest(vpngateURL); err != nil {
			c.logger.Error(err)
			return nil, err
		}

		if err := gocsv.UnmarshalBytes(body, &ips); err != nil { // Load clients from file
			c.logger.Error(err)
		}
	}
	for _, client := range ips {
		c.logger.Error(client.IP)
	}

	//c.hosts = strings.Split(string(allbody), "\n")

	c.lastUpdate = time.Now()

	return c.hosts, nil

}
func (c *VPNGate) MakeRequest(urllist string) ([]byte, error) {
	var body bytes.Buffer

	req, err := http.NewRequest(http.MethodGet, urllist, nil)
	if err != nil {
		return nil, err
	}

	var client = NewClient()

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	cut, ok := skip(body.Bytes(), 1)
	if !ok {
		return nil, fmt.Errorf("less than %d lines", 2)
	}

	return cut, err
}

func (c *VPNGate) List() ([]string, error) {
	return c.Load(nil)
}
