package providers

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/domgolonka/threatdefender/app/models"
)

type PubProxy struct {
	proxy      models.Proxy
	proxyList  []models.Proxy
	lastUpdate time.Time
}

func NewPubProxy() *PubProxy {
	return &PubProxy{}
}

func (*PubProxy) Name() string {
	return "pubproxy.com"
}

func (x *PubProxy) SetProxy(proxy models.Proxy) {
	x.proxy = proxy
}

func (x *PubProxy) MakeRequest() ([]byte, error) {
	var client = NewClient()
	if x.proxy.IP != "" {
		proxyURL, err := url.Parse(x.proxy.ToString())
		if err != nil {
			return nil, err
		}
		client.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
	} else {
		client.Transport.(*http.Transport).Proxy = http.ProxyFromEnvironment
	}

	resp, err := client.Get("http://pubproxy.com/api/proxy?limit=20&format=txt&type=http")
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

func (x *PubProxy) Load() ([]models.Proxy, error) {
	if time.Now().Unix() >= x.lastUpdate.Unix()+(60*20) {
		x.proxyList = make([]models.Proxy, 0)
	}

	if len(x.proxyList) != 0 {
		return x.proxyList, nil
	}

	body, err := x.MakeRequest()
	if err != nil {
		return nil, err
	}

	proxies := strings.Split(string(body), "\n")
	for _, s := range proxies {
		proxy := strings.Split(s, ":")
		prox := models.Proxy{
			IP:   proxy[0],
			Port: proxy[1],
			Type: "http",
		}
		x.proxyList = append(x.proxyList, prox)
	}

	x.lastUpdate = time.Now()

	return x.proxyList, nil
}

func (x *PubProxy) List() ([]models.Proxy, error) {
	return x.Load()
}
