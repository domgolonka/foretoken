package providers

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
)

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

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	var err error
	if body == nil {
		if body, err = c.MakeRequest(vpngateURL); err != nil {
			c.logger.Error(err)
			return nil, err
		}

		c.hosts = re.FindAllString(string(body), -1)

	}

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
