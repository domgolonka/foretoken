package providers

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type PubProxy struct {
	proxy      string
	proxyList  []string
	lastUpdate time.Time
}

func NewPubProxy() *PubProxy {
	return &PubProxy{}
}

func (*PubProxy) Name() string {
	return "pubproxy.com"
}

func (x *PubProxy) SetProxy(proxy string) {
	x.proxy = proxy
}

func (x *PubProxy) MakeRequest() ([]byte, error) {
	var client = NewClient()
	if x.proxy != "" {
		proxyURL, err := url.Parse("http://" + x.proxy)
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

func (x *PubProxy) Load() ([]string, error) {
	if time.Now().Unix() >= x.lastUpdate.Unix()+(60*20) {
		x.proxyList = make([]string, 0, 0)
	}

	if len(x.proxyList) != 0 {
		return x.proxyList, nil
	}

	body, err := x.MakeRequest()
	if err != nil {
		return nil, err
	}

	x.proxyList = strings.Split(string(body), "\n")
	if len(x.proxyList) != 20 {
		x.proxyList = make([]string, 0, 0)
	} else {
		x.lastUpdate = time.Now()
	}

	return x.proxyList, nil
}

func (x *PubProxy) List() ([]string, error) {
	return x.Load()
}
