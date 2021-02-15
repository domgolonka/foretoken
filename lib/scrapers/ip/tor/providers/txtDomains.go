package providers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/domgolonka/threatdefender/pkg/utils/ip"
)

type TxtDomains struct {
	tor        string
	torList    []string
	logger     logrus.FieldLogger
	lastUpdate time.Time
}

var speedlist = []string{"https://raw.githubusercontent.com/SecOps-Institute/Tor-IP-Addresses/master/tor-exit-nodes.lst"}

func NewTxtDomains(logger logrus.FieldLogger) *TxtDomains {
	return &TxtDomains{logger: logger}
}
func (*TxtDomains) Name() string {
	return "SecOps-Institute"
}

func (c *TxtDomains) SetProxy(tor string) {
	c.tor = tor
}

func (c *TxtDomains) Load(body []byte) ([]string, error) {
	// don't need to update this more than once a day!
	if time.Now().Unix() >= c.lastUpdate.Unix()+(82800) {
		c.torList = make([]string, 0)
	}

	if len(c.torList) != 0 {
		return c.torList, nil
	}
	allbody := make([]string, len(speedlist))
	if body == nil {
		var err error
		for i := 0; i < len(speedlist); i++ {
			if body, err = c.MakeRequest(speedlist[i]); err != nil {
				return nil, err
			}

			ipv4, err := ip.ParseIps(body)
			if err != nil {
				return nil, err
			}
			allbody = append(allbody, ipv4...)
		}
	}

	c.torList = allbody

	c.lastUpdate = time.Now()

	return c.torList, nil

}
func (c *TxtDomains) MakeRequest(urllist string) ([]byte, error) {
	var body bytes.Buffer

	req, err := http.NewRequest(http.MethodGet, urllist, nil)
	if err != nil {
		return nil, err
	}

	var client = NewClient()
	client.Transport.(*http.Transport).Proxy = http.ProxyFromEnvironment

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

func (c *TxtDomains) List() ([]string, error) {
	return c.Load(nil)
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
