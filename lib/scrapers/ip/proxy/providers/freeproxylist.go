package providers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/domgolonka/threatdefender/app/models"

	"github.com/jbowtie/gokogiri"
)

type FreeProxyList struct {
	proxy      models.Proxy
	proxyList  []models.Proxy
	lastUpdate time.Time
}

func NewFreeProxyList() *FreeProxyList {
	return &FreeProxyList{}
}

func (x *FreeProxyList) SetProxy(proxy models.Proxy) {
	x.proxy = proxy
}

func (*FreeProxyList) Name() string {
	return "free-proxy-list.net"
}

func (x *FreeProxyList) MakeRequest() ([]byte, error) {
	req, err := http.NewRequest("GET", "https://free-proxy-list.net/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Language", "en-US,en;q=0.8,uk;q=0.6,ru;q=0.4")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Authority", "free-proxy-list.net")
	req.Header.Set("Referer", "https://free-proxy-list.net/web-proxy.html")

	var client = NewClient()
	if x.proxy.IP != "" {
		proxyURL, err := url.Parse("http://" + x.proxy.IP + ":" + x.proxy.Port)
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

func (x *FreeProxyList) Load(body []byte) ([]models.Proxy, error) {
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
	//*[@id="proxylisttable"]/tbody/tr[1]/td[1]
	ips, err := doc.Search(`//*[@id="proxylisttable"]/tbody/tr/td[1]`)
	if err != nil {
		return nil, err
	}
	//*[@id="proxylisttable"]/tbody/tr[1]/td[2]
	ports, err := doc.Search(`//*[@id="proxylisttable"]/tbody/tr/td[2]`)
	if err != nil {
		return nil, err
	}

	if len(ips) == 0 {
		return nil, errors.New("ip not found")
	}

	if len(ips) != len(ports) {
		return nil, errors.New("len port not equal ip")
	}

	x.proxyList = make([]models.Proxy, 0, len(ips))

	for i, ip := range ips {
		prox := models.Proxy{
			IP:   ip.Content(),
			Port: ports[i].Content(),
			Type: "http", // todo
		}
		x.proxyList = append(x.proxyList, prox)
		//x.proxyList = append(x.proxyList, ip.Content()+":"+ports[i].Content())
	}

	x.lastUpdate = time.Now()
	return x.proxyList, nil
}

func (x *FreeProxyList) List() ([]models.Proxy, error) {
	return x.Load(nil)
}
