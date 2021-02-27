package providers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/domgolonka/foretoken/app/models"

	"github.com/jbowtie/gokogiri"
)

const coolProxyURL = `https://www.cool-proxy.net/proxies/http_proxy_list/sort:score/direction:desc`

type CoolProxy struct {
	proxy      models.Proxy
	proxyList  []models.Proxy
	lastUpdate time.Time
}

func NewCoolProxy() *CoolProxy {
	return &CoolProxy{}
}

func (c *CoolProxy) SetProxy(proxy models.Proxy) {
	c.proxy = proxy
}

func (*CoolProxy) Name() string {
	return "www.cool-proxy.net"
}

func (c *CoolProxy) Load(body []byte) ([]models.Proxy, error) {
	if time.Now().Unix() >= c.lastUpdate.Unix()+(60*20) {
		c.proxyList = make([]models.Proxy, 0)
	}

	if len(c.proxyList) != 0 {
		return c.proxyList, nil
	}

	if body == nil {
		var err error
		if body, err = c.MakeRequest(); err != nil {
			return nil, err
		}
	}

	doc, err := gokogiri.ParseHtml(body)
	if err != nil {
		return nil, err
	}

	defer doc.Free()
	ips, err := doc.Search(`//*[@id="main"]/table/tbody/tr/td[1]`)
	if err != nil {
		return nil, err
	}

	ports, err := doc.Search(`//*[@id="main"]/table/tr/td[2]`)
	if err != nil {
		return nil, err
	}

	if len(ips) == 0 {
		return nil, errors.New("ip not found")
	}

	if len(ips) != len(ports) {
		return nil, errors.New("len port not equal ip")
	}

	r := regexp.MustCompile(`"(.*?[^\\])"`)

	for i, ip := range ips {
		raw := r.FindStringSubmatch(ip.Content())
		if len(raw) != 2 {
			continue
		}

		decoded, err := base64.StdEncoding.DecodeString(string(bytes.Map(rot13, []byte(raw[1]))))
		if err != nil {
			continue
		}
		prox := models.Proxy{
			IP:   string(decoded),
			Port: ports[i].Content(),
			Type: "http", // todo
		}
		c.proxyList = append(c.proxyList, prox)
	}
	c.lastUpdate = time.Now()
	return c.proxyList, nil
}

func (c *CoolProxy) MakeRequest() ([]byte, error) {
	var body bytes.Buffer

	req, err := http.NewRequest(http.MethodGet, coolProxyURL, nil)
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

	return body.Bytes(), err
}

func (c *CoolProxy) List() ([]models.Proxy, error) {
	return c.Load(nil)
}

func rot13(x rune) rune {
	capital := x >= 'A' && x <= 'Z'
	if !capital && (x < 'a' || x > 'z') {
		return x
	}
	x += 13
	if capital && x > 'Z' || !capital && x > 'z' {
		x -= 26
	}
	return x
}
