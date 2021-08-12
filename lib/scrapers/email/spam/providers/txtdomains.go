package providers

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/domgolonka/foretoken/app/models"

	"github.com/domgolonka/foretoken/app/entity"

	"github.com/sirupsen/logrus"
)

type TxtDomains struct {
	iplist     []models.SpamEmail
	logger     logrus.FieldLogger
	lastUpdate time.Time
	feedList   []string
}

func NewTxtDomains(logger logrus.FieldLogger, feedList []string) *TxtDomains {
	logger.Debug("starting TxtDomains")
	return &TxtDomains{
		logger:   logger,
		feedList: feedList,
	}
}
func (*TxtDomains) Name() string {
	return "text_domain"
}

func (c *TxtDomains) Load(body []byte) ([]models.SpamEmail, error) {

	// don't need to update this more than once a day!
	if time.Now().Unix() >= c.lastUpdate.Unix()+(82800) {
		c.iplist = make([]models.SpamEmail, 0)
	}

	f := entity.Feed{
		Logger: c.logger,
	}
	feed, err := f.ReadFile(c.feedList...)
	if err != nil {
		return nil, err
	}
	ips := make(map[string]entity.DomainAnalysis)
	for _, activeFeed := range feed {
		c.logger.Infof("[INFO] Importing data feed %s", activeFeed.Name)
		feedResultsDomains, err := activeFeed.FetchString()
		if err == nil {
			for k, e := range feedResultsDomains { // k is the ip string,  e is the
				if _, ok := ips[k]; ok {
					ip := ips[k]
					ip.Type = e.Type
					ip.Score += e.Score
					ip.Lists = append(ip.Lists, e.Lists[0])
					ips[k] = ip
				} else {
					ips[k] = e
				}

				freeEmail := models.SpamEmail{
					Domain: strings.ToLower(ips[k].Domain),
				}
				c.iplist = append(c.iplist, freeEmail)

			}

			c.logger.Infof("[INFO] Imported %d domains from data feed %d", len(feedResultsDomains),
				len(feedResultsDomains), activeFeed.Name)
		} else {
			c.logger.Errorf("[ERROR] Importing data feed %s\n failed : %v", activeFeed.Name, err)
		}
	}

	c.lastUpdate = time.Now()
	return c.iplist, nil

}
func (c *TxtDomains) MakeRequest(urlList string) ([]byte, error) {
	var body bytes.Buffer

	req, err := http.NewRequest(http.MethodGet, urlList, nil)
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

func (c *TxtDomains) List() ([]models.SpamEmail, error) {
	return c.Load(nil)
}
