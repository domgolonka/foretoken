package providers

import (
	"bytes"
	"net/http"
	"time"

	"github.com/domgolonka/foretoken/app/models"

	"github.com/domgolonka/foretoken/app/entity"

	"github.com/sirupsen/logrus"
)

type TxtDomains struct {
	iplist     []models.Tor
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

func (c *TxtDomains) Load(body []byte) ([]models.Tor, error) {

	// don't need to update this more than once a day!
	if time.Now().Unix() >= c.lastUpdate.Unix()+(82800) {
		c.iplist = make([]models.Tor, 0)
	}

	f := entity.Feed{
		Logger: c.logger,
	}
	feed, err := f.ReadFile(c.feedList...)
	if err != nil {
		return nil, err
	}
	ips := make(map[string]entity.IPAnalysis)
	subnets := make(map[string]entity.SUBNETAnalysis)
	for _, activeFeed := range feed {
		c.logger.Infof("[INFO] Importing data feed %s", activeFeed.Name)
		feedResultsIPs, feedResultsSubnets, err := activeFeed.FetchIP()
		if err == nil {
			for k, e := range feedResultsIPs { // k is the ip string,  e is the
				if _, ok := ips[k]; ok {
					ip := ips[k]
					ip.Type = e.Type
					ip.Score += e.Score
					ip.Lists = append(ip.Lists, e.Lists[0])
					ips[k] = ip
				} else {
					ips[k] = e
				}

				spam := models.Tor{
					IP:    ips[k].IP,
					Score: ips[k].Score,
					Type:  ips[k].Type,
				}
				c.iplist = append(c.iplist, spam)

			}
			for k, e := range feedResultsSubnets {
				if _, ok := subnets[k]; ok {
					subnet := subnets[k]
					subnet.Type = e.Type
					subnet.Score += e.Score
					subnet.Lists = append(subnet.Lists, e.Lists[0])
					subnets[k] = subnet
				} else {
					subnets[k] = e
				}
				spam := models.Tor{
					IP:     subnets[k].IP,
					Prefix: subnets[k].PrefixLength,
					Score:  subnets[k].Score,
					Type:   subnets[k].Type,
				}
				c.iplist = append(c.iplist, spam)
			}
			c.logger.Infof("[INFO] Imported %d ips and %d subnets from data feed %s", len(feedResultsIPs),
				len(feedResultsSubnets), activeFeed.Name)
		} else {
			c.logger.Errorf("[ERROR] Importing data feed %s\n failed : %s", activeFeed.Name, err)
		}
	}

	c.lastUpdate = time.Now()
	return c.iplist, nil

}
func (*TxtDomains) MakeRequest(urlList string) ([]byte, error) {
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

func (c *TxtDomains) List() ([]models.Tor, error) {
	return c.Load(nil)
}
