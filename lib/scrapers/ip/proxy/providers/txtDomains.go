package providers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/domgolonka/foretoken/app/entity"
	"github.com/domgolonka/foretoken/app/models"

	"github.com/domgolonka/foretoken/pkg/utils/ip"
)

type TxtDomains struct {
	proxy      models.Proxy
	proxyList  []models.Proxy
	lastUpdate time.Time
	logger     logrus.FieldLogger
}

type Feed struct {
	Name string
	URL  string
	Type string
}

func NewTxtDomains(logger logrus.FieldLogger) *TxtDomains {
	return &TxtDomains{logger: logger}
}
func (*TxtDomains) Name() string {
	return "txt_domain_proxy"
}

func (c *TxtDomains) SetProxy(proxy models.Proxy) {
	c.proxy = proxy
}

func (c *TxtDomains) Load(body []byte) ([]models.Proxy, error) {
	// don't need to update this more than once a day!
	if time.Now().Unix() >= c.lastUpdate.Unix()+(82800) {
		c.proxyList = make([]models.Proxy, 0)
	}
	f := entity.Feed{
		Logger: c.logger,
	}
	feed, err := f.ReadFile("ip_proxy.json")
	if err != nil {
		return nil, err
	}

	if len(c.proxyList) != 0 {
		return c.proxyList, nil
	}
	if body == nil {
		for i := 0; i < len(feed); i++ {
			expressions, err := feed[i].GetExpressions()
			if err != nil {
				return nil, err
			}
			if body, err = c.MakeRequest(feed[i].URL); err != nil {
				return nil, err
			}

			ipv4 := ip.ParseIPs(body, expressions)
			for _, s := range ipv4 {
				var prox models.Proxy
				if strings.Contains(s, ":") {
					proxy := strings.Split(s, ":")
					prox = models.Proxy{
						IP:   proxy[0],
						Port: proxy[1],
						Type: feed[i].Type,
					}
				} else {
					prox = models.Proxy{
						IP:   s,
						Type: feed[i].Type,
					}
				}

				c.proxyList = append(c.proxyList, prox)
			}
		}
	}

	c.lastUpdate = time.Now()

	return c.proxyList, nil

}
func (c *TxtDomains) MakeRequest(urllist string) ([]byte, error) {
	var body bytes.Buffer

	req, err := http.NewRequest(http.MethodGet, urllist, nil)
	if err != nil {
		return nil, err
	}

	var client = NewClient()
	if c.proxy.IP != "" {
		proxyURL, err := url.Parse(c.proxy.ToString())
		if err != nil {
			return nil, err
		}
		client.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
	} else {
		client.Transport.(*http.Transport).Proxy = http.ProxyFromEnvironment
	}

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

func (c *TxtDomains) List() ([]models.Proxy, error) {
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
