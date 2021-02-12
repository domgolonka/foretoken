package providers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/jbowtie/gokogiri"
)

var (
	portParamsRegexp = regexp.MustCompile(`([a-z]=\d;){10}`)
	portRegexp       = regexp.MustCompile(`(\+[a-z]){2,4}`)
	ipRegexp         = regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}`)
)

type XseoIn struct {
	proxyList  []string
	lastUpdate time.Time
	proxy      string
}

func NewXseoIn() *XseoIn {
	return &XseoIn{}
}

func (x *XseoIn) SetProxy(proxy string) {
	x.proxy = proxy
}

func (*XseoIn) Name() string {
	return "xseo.in"
}

func (x *XseoIn) MakeRequest() ([]byte, error) {
	postData := strings.NewReader(`submit=%D0%9F%D0%BE%D0%BA%D0%B0%D0%B7%D0%B0%D1%82%D1%8C+%D0%BF%D0%BE+150+%D0%BF%D1%80%D0%BE%D0%BA%D1%81%D0%B8+%D0%BD%D0%B0+%D1%81%D1%82%D1%80%D0%B0%D0%BD%D0%B8%D1%86%D0%B5`)
	req, err := http.NewRequest("POST", "http://xseo.in/proxylist", postData)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Origin", "http://xseo.in")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8,uk;q=0.6,ru;q=0.4")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Referer", "http://xseo.in/proxylist")
	req.Header.Set("Connection", "keep-alive")
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

func (x *XseoIn) DecodeParamsToMap(params string) map[byte]byte {
	if len(params) < 39 {
		return nil
	}
	return map[byte]byte{
		params[0]:  params[2],  //0
		params[4]:  params[6],  //1
		params[8]:  params[10], //2
		params[12]: params[14], //3
		params[16]: params[18], //4
		params[20]: params[22], //5
		params[24]: params[26], //6
		params[28]: params[30], //7
		params[32]: params[34], //8
		params[36]: params[38], //9
	}
}

func (x *XseoIn) DecodePort(decodeParams map[byte]byte, encryptedData string) []byte {
	if len(encryptedData) == 8 {
		return []byte{decodeParams[encryptedData[1]], decodeParams[encryptedData[3]], decodeParams[encryptedData[5]], decodeParams[encryptedData[7]]}
	}
	if len(encryptedData) == 4 {
		return []byte{decodeParams[encryptedData[1]], decodeParams[encryptedData[3]]}
	}
	return nil
}

func (x *XseoIn) Load(body []byte) ([]string, error) {
	if time.Now().Unix() >= x.lastUpdate.Unix()+(60*20) {
		x.proxyList = make([]string, 0)
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

	decodeParams := x.DecodeParamsToMap(portParamsRegexp.FindString(string(body)))
	if decodeParams == nil {
		return nil, errors.New("decodeParams can not be <nil>")
	}

	doc, err := gokogiri.ParseHtml(body)
	if err != nil {
		return nil, err
	}
	defer doc.Free()

	ips, err := doc.Search(`//tr[contains(@class, 'cls8') or contains(@class ,'cls81')]`)
	if err != nil {
		return nil, err
	}

	if len(ips) == 0 {
		return nil, errors.New("ip not found")
	}

	x.proxyList = make([]string, 0, len(ips))

	for _, ip := range ips {
		data := strings.Split(ip.Content(), ":")
		if len(data) > 1 {
			if ipRegexp.MatchString(data[0]) {
				if port := x.DecodePort(decodeParams, portRegexp.FindString(data[1])); port != nil {
					x.proxyList = append(x.proxyList, data[0]+":"+string(port))
				}
			}
		}
	}
	x.lastUpdate = time.Now()
	return x.proxyList, nil
}

func (x *XseoIn) List() ([]string, error) {
	return x.Load(nil)
}
