package providers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/domgolonka/foretoken/app/models"

	"github.com/jbowtie/gokogiri"
	scraper "github.com/soluchok/go-cloudflare-scraper"
)

type HidemyName struct {
	proxy      models.Proxy
	proxyList  []models.Proxy
	lastUpdate time.Time
}

func NewHidemyName() *HidemyName {
	return &HidemyName{}
}

func (x *HidemyName) Name() string {
	return "hidemyna.me"
}

func (x *HidemyName) SetProxy(proxy models.Proxy) {
	x.proxy = proxy
}

func (x *HidemyName) MakeRequest() ([]byte, error) {
	transport := NewTransport()

	if x.proxy.IP != "" {
		proxyURL, err := url.Parse(x.proxy.ToString())
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	scraperTransport, _ := scraper.NewTransport(transport)
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: scraperTransport,
	}

	resp, err := client.Get("https://hidemyna.me/en/proxy-list")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	var body bytes.Buffer
	if _, err := io.Copy(&body, resp.Body); err != nil {
		return nil, err
	}

	return body.Bytes(), nil
}

func (x *HidemyName) Load(body []byte) ([]models.Proxy, error) {
	if time.Now().Unix() >= x.lastUpdate.Unix()+(60*20) {
		x.proxyList = make([]models.Proxy, 0)
	}

	if len(x.proxyList) != 0 {
		return x.proxyList, nil
	}

	if body == nil {
		var err error
		if body, err = x.MakeRequest(); err != nil {
			return nil, err
		}
	}

	doc, err := gokogiri.ParseHtml(body)
	if err != nil {
		return nil, err
	}
	defer doc.Free()

	//ips, err := doc.Search(`//td[contains(@class, 'tdl')]`)
	ips, err := doc.Search(`/html/body/div[1]/div[4]/div/div[4]/table/tbody/tr/td[1]`)
	if err != nil {
		return nil, err
	}

	if len(ips) == 0 {
		return nil, errors.New("ip not found")
	}

	x.proxyList = make([]models.Proxy, 0, len(ips))

	for _, ip := range ips {
		port := ip.NextSibling()
		types := ip.NextSibling().NextSibling().NextSibling()
		if ipRegexp.MatchString(ip.Content()) {
			prox := models.Proxy{
				IP:   ip.Content(),
				Port: port.Content(),
				Type: strings.ToLower(types.Content()),
			}
			x.proxyList = append(x.proxyList, prox)
			//x.proxyList = append(x.proxyList, ip.Content()+":"+port.Content())
		}

	}
	x.lastUpdate = time.Now()
	return x.proxyList, nil
}

func (x *HidemyName) List() ([]models.Proxy, error) {
	return x.Load(nil)
}
