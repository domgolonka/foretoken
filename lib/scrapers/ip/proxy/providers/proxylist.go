package providers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type ProxyList struct {
	proxy      string
	proxyList  []string
	lastUpdate time.Time
}

var proxyRegexp = regexp.MustCompile(`Proxy\(\'([\w\d=+]+)\'\)`)

func NewProxyList() *ProxyList {
	return &ProxyList{}
}

func (x *ProxyList) SetProxy(proxy string) {
	x.proxy = proxy
}

func (*ProxyList) Name() string {
	return "proxy-list.org"
}

func (x *ProxyList) MakeRequest(page int) ([]byte, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("http://proxy-list.org/english/index.php?p=%d", page), nil)
	if err != nil {
		return nil, err
	}
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

	resp, err := client.Do(req)
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

func (x *ProxyList) Load() ([]string, error) {
	if time.Now().Unix() >= x.lastUpdate.Unix()+(60*20) {
		x.proxyList = make([]string, 0, 0)
	}

	if len(x.proxyList) != 0 {
		return x.proxyList, nil
	}

	for i := 1; i <= 10; i++ {
		body, err := x.MakeRequest(i)
		if err != nil {
			return nil, err
		}

		for _, d := range proxyRegexp.FindAllString(string(body), -1) {
			data, err := base64.StdEncoding.DecodeString(d[7 : len(d)-2])
			if err != nil {
				continue
			}
			x.proxyList = append(x.proxyList, string(data))
		}
	}

	x.lastUpdate = time.Now()
	return x.proxyList, nil
}

func (x *ProxyList) List() ([]string, error) {
	return x.Load()
}
