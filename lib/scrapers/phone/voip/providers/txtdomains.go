package providers

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var domains = []string{"https://raw.githubusercontent.com/andreis/disposable-email-domains/master/domains.txt",
	"https://raw.githubusercontent.com/martenson/disposable-email-domains/master/disposable_email_blocklist.conf",
	"https://raw.githubusercontent.com/oosalt/disposable-email-domain-list/master/domains.txt",
	"https://raw.githubusercontent.com/disposable/disposable-email-domains/master/domains.txt",
	"https://raw.githubusercontent.com/SoftCreatR/dead-letter-dump/master/blacklist_unhashed.txt",
	"https://raw.githubusercontent.com/Xyborg/disposable-burner-email-providers/master/disposable-domains.txt",
	"https://raw.githubusercontent.com/edwin-zvs/email-providers/master/email-providers.csv"}

type TxtDomains struct {
	hosts      []string
	logger     logrus.FieldLogger
	lastUpdate time.Time
}

func NewTxtDomains(logger logrus.FieldLogger) *TxtDomains {
	logger.Debug("starting TxtDomains")
	return &TxtDomains{
		logger: logger,
	}
}
func (*TxtDomains) Name() string {
	return "text_domain"
}

func (c *TxtDomains) Load(body []byte) ([]string, error) {

	// don't need to update this more than once a day!
	if time.Now().Unix() >= c.lastUpdate.Unix()+(43200) {
		c.hosts = make([]string, 0, 0)
	}

	if len(c.hosts) != 0 {
		return c.hosts, nil
	}
	allbody := make([]byte, 0, len(domains))
	if body == nil {
		var err error
		for i := 0; i < len(domains); i++ {
			c.logger.Debug(domains[i])
			if body, err = c.MakeRequest(domains[i]); err == nil {
				allbody = append(allbody, body...)
			}

		}
	}

	c.hosts = strings.Split(string(allbody), "\n")

	c.lastUpdate = time.Now()

	return c.hosts, nil

}
func (c *TxtDomains) MakeRequest(urllist string) ([]byte, error) {
	var body bytes.Buffer

	req, err := http.NewRequest(http.MethodGet, urllist, nil)
	if err != nil {
		return nil, err
	}

	var client = NewClient()

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

func (c *TxtDomains) List() ([]string, error) {
	return c.Load(nil)
}
