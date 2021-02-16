package providers

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var domains = []string{"https://raw.githubusercontent.com/dpup/freemail/main/data/free.txt"}

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
		c.hosts = make([]string, 0)
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

func (c *TxtDomains) List() ([]string, error) {
	return c.Load(nil)
}
