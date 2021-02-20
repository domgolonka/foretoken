package providers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/domgolonka/threatdefender/app/models"

	"github.com/domgolonka/threatdefender/pkg/utils/ip"

	"github.com/domgolonka/threatdefender/app/entity"
	"github.com/sirupsen/logrus"
)

//var ervsfreevps = []string{"https://raw.githubusercontent.com/ejrv/VPNs/master/vpn-ipv4.txt",
//	"https://raw.githubusercontent.com/ejrv/VPNs/master/vpn-ipv6.txt"}

type TxtDomains struct {
	hosts      []models.Vpn
	logger     logrus.FieldLogger
	lastUpdate time.Time
}

func NewTxtDomains(logger logrus.FieldLogger) *TxtDomains {
	logger.Debug("starting VPN TxtDomains")
	return &TxtDomains{
		logger: logger,
	}
}
func (*TxtDomains) Name() string {
	return "vpn_txt_domains"
}

func (c *TxtDomains) Load(body []byte) ([]models.Vpn, error) {
	// don't need to update this more than once a day!
	if time.Now().Unix() >= c.lastUpdate.Unix()+(82800) {
		c.hosts = make([]models.Vpn, 0)
	}

	f := entity.Feed{}
	feed, err := f.ReadFile("ip_vpn.json")
	if err != nil {
		return nil, err
	}

	if len(c.hosts) != 0 {
		return c.hosts, nil
	}
	if body == nil {
		for i := 0; i < len(feed); i++ {
			expressions, err := feed[i].GetExpressions()
			if err != nil {
				return nil, err
			}
			if body, err = c.MakeRequest(feed[i].URL); err == nil {

				ips := ip.ParseIPs(body, expressions)
				for _, a := range ips {
					vpn := models.Vpn{
						URL:    a,
						Source: feed[i].Name,
					}
					c.hosts = append(c.hosts, vpn)
				}
			}
		}
	}
	c.lastUpdate = time.Now()
	return c.hosts, nil

}
func (c *TxtDomains) MakeRequest(urllist string) ([]byte, error) {
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

	cut, ok := skip(body.Bytes(), 2)
	if !ok {
		return nil, fmt.Errorf("less than %d lines", 2)
	}

	return cut, err
}

func skip(b []byte, n int) ([]byte, bool) {
	for ; n > 0; n-- {
		if len(b) == 0 {
			return nil, false
		}
		x := bytes.IndexByte(b, '\n')
		if x < 0 {
			x = len(b)
		} else {
			x++
		}
		b = b[x:]
	}
	return b, true
}

func (c *TxtDomains) List() ([]models.Vpn, error) {
	return c.Load(nil)
}
