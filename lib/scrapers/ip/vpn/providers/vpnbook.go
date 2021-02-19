package providers

//
//import (
//	"archive/zip"
//	"bytes"
//	"fmt"
//	"io"
//	"io/ioutil"
//	"net/http"
//	"path/filepath"
//	"regexp"
//	"strings"
//	"time"
//
//	"github.com/gocolly/colly/v2"
//	"github.com/sirupsen/logrus"
//)
//
//type VPNBook struct {
//	hosts      []string
//	logger     logrus.FieldLogger
//	lastUpdate time.Time
//}
//
//func NewVPNBook(logger logrus.FieldLogger) *VPNBook {
//	logger.Debug("starting VPNBook")
//	return &VPNBook{
//		logger: logger,
//	}
//}
//func (*VPNBook) Name() string {
//	return "vpn_book"
//}
//
//func (c *VPNBook) Load(body []byte) ([]string, error) {
//	// don't need to update this more than once a day!
//	if time.Now().Unix() >= c.lastUpdate.Unix()+(82800) {
//		c.hosts = make([]string, 0)
//	}
//
//	if len(c.hosts) != 0 {
//		return c.hosts, nil
//	}
//	allbody := make([]byte, 0, len(ervsfreevps))
//	if body == nil {
//		var err error
//		for i := 0; i < len(ervsfreevps); i++ {
//			c.logger.Debug(ervsfreevps[i])
//			if body, err = c.MakeRequest(ervsfreevps[i]); err != nil {
//				allbody = append(allbody, body...)
//			}
//
//		}
//	}
//
//	c.hosts = strings.Split(string(allbody), "\n")
//
//	c.lastUpdate = time.Now()
//
//	return c.hosts, nil
//
//}
//
//func (c *VPNBook) Download(src URLs) ([]string, error) {
//	hosts := []string{}
//	c.logger.Debug("starting Download for " + src.URL)
//
//	if src.Format == OPENVPN {
//		res, err := http.Get(src.URL)
//		if err != nil {
//			c.logger.Fatal(err)
//		}
//		defer res.Body.Close()
//		d, err := ioutil.ReadAll(res.Body)
//		if err != nil {
//			c.logger.Fatal(err)
//		}
//
//		var filenames []string
//
//		r, err := zip.NewReader(bytes.NewReader(d), int64(len(d)))
//		if err != nil {
//			return filenames, err
//		}
//
//		for _, f := range r.File {
//			if filepath.Ext(f.Name) == ".ovpn" && !strings.HasPrefix(f.Name, "__MACOSX") {
//				var buf bytes.Buffer
//
//				rc, err := f.Open()
//				if err != nil {
//					c.logger.Fatal(err)
//				}
//
//				reRemote := regexp.MustCompile(`remote (\S+)`)
//
//				_, err = io.Copy(&buf, rc) //nolint
//				if err != nil {
//					return nil, err
//				}
//
//				rc.Close()
//				domainTmp := reRemote.FindStringSubmatch(string(buf.String()))
//				var domainName = domainTmp[1]
//
//				hosts = append(hosts, domainName)
//			}
//		}
//	}
//
//	return hosts, nil
//
//}
//
//func (c *VPNBook) MakeRequest(urllist string) ([]byte, error) {
//	var body bytes.Buffer
//
//	req, err := http.NewRequest(http.MethodGet, urllist, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var client = NewClient()
//
//	resp, err := client.Do(req)
//	if resp != nil {
//		defer resp.Body.Close()
//	}
//
//	if err != nil {
//		return nil, err
//	}
//
//	_, err = body.ReadFrom(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	cut, ok := skip(body.Bytes(), 2)
//	if !ok {
//		return nil, fmt.Errorf("less than %d lines", 2)
//	}
//
//	return cut, err
//}
//
//func (c *VPNBook) List() ([]string, error) {
//	hosts := []string{}
//	links := []string{}
//	// todo
//	col := colly.NewCollector()
//	col.OnHTML("#openvpn", func(e *colly.HTMLElement) {
//		link := e.Request.AbsoluteURL(e.Attr("href"))
//		if link != "" {
//			links = append(links, link)
//		}
//
//	})
//	for i := 0; i < len(links); i++ {
//		host, err := c.Download(OpenVPNURLs[i])
//		if err != nil {
//			return hosts, err
//		}
//		hosts = append(hosts, host...)
//	}
//
//	return hosts, nil
//}
