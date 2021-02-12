package providers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/jbowtie/gokogiri"
)

var (
	_portRegexp = regexp.MustCompile(`\d+`)
)

type FreeProxyListNet struct {
	proxy      string
	proxyList  []string
	lastUpdate time.Time
}

func NewFreeProxyListNet() *FreeProxyListNet {
	return &FreeProxyListNet{}
}

func (x *FreeProxyListNet) SetProxy(proxy string) {
	x.proxy = proxy
}

func (*FreeProxyListNet) Name() string {
	return "www.freeproxylists.net"
}

func (x *FreeProxyListNet) MakeRequest() ([]byte, error) {
	req, err := http.NewRequest("GET", "http://www.freeproxylists.net/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Language", "en-US,en;q=0.8,uk;q=0.6,ru;q=0.4")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")

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

func (x *FreeProxyListNet) IPDecode(ipstr string) (string, error) {
	data, err := url.PathUnescape(ipstr[10 : len(ipstr)-2])
	if err != nil {
		return "", err
	}
	return ipRegexp.FindString(data), nil
}

func (x *FreeProxyListNet) Load(body []byte) ([]string, error) {
	if time.Now().Unix() >= x.lastUpdate.Unix()+(60*20) {
		x.proxyList = make([]string, 0, 0)
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
	ips, err := doc.Search(`//tr[contains(@class, 'Odd') or contains(@class ,'Even')]/td[1]`)
	if err != nil {
		return nil, err
	}
	_proxyList := make([]string, 0)
	portList := make([]string, 0)

	if len(ips) == 0 {
		return nil, errors.New("ips not found")
	}

	for _, ip := range ips {
		ipaddr, err := x.IPDecode(ip.Content())
		if err != nil || ipaddr == "" {
			continue
		}
		_proxyList = append(_proxyList, ipaddr)
	}

	ports, err := doc.Search(`//tr[contains(@class, 'Odd') or contains(@class ,'Even')]/td[2]`)
	if err != nil {
		return nil, err
	}

	if len(ports) == 0 {
		return nil, errors.New("ports not found")
	}

	for _, port := range ports {
		if _portRegexp.MatchString(port.Content()) {
			portList = append(portList, port.Content())
		}
	}

	if len(portList) != len(_proxyList) {
		return nil, errors.New("len port not equal ip")
	}

	x.proxyList = make([]string, 0, len(_proxyList))
	for i, proxy := range _proxyList {
		x.proxyList = append(x.proxyList, proxy+":"+portList[i])
	}

	x.lastUpdate = time.Now()
	return x.proxyList, nil
}

func (x *FreeProxyListNet) List() ([]string, error) {
	return x.Load(nil)
}
