package spamemail

import (
	"github.com/domgolonka/foretoken/app/data"
	"github.com/domgolonka/foretoken/lib/scrapers/email/spam/providers"
	"github.com/sirupsen/logrus"

	"reflect"
	"sync"
)

var (
	instance *SpamEmail
	once     sync.Once
)

type SpamEmail struct {
	providers []Provider
	store     data.SpamEmailStore
	logger    logrus.FieldLogger
}

func (p *SpamEmail) isProvider(provider Provider) bool {
	for _, pr := range p.providers {
		if reflect.TypeOf(pr) == reflect.TypeOf(provider) {
			return true
		}
	}
	return false
}
func (p *SpamEmail) AddProvider(provider Provider) {
	if !p.isProvider(provider) {
		p.providers = append(p.providers, provider)
	}
}
func (p *SpamEmail) load() {
	for _, provider := range p.providers {
		hosts, err := provider.List()

		if err != nil {
			p.logger.Errorln(err)
		}
		p.logger.Println(provider.Name(), len(hosts))

		for i := 0; i < len(hosts); i++ {
			p.createOrIgnore(hosts[i].Domain, hosts[i].Score)
		}
	}
}
func (p *SpamEmail) createOrIgnore(domain string, score int) bool {
	_, err := p.store.Create(domain, score)
	return err == nil
}

func (p *SpamEmail) run() {
	go p.load()
}

func (p *SpamEmail) Get() (*[]string, error) {
	return p.store.FindAll()
}

func NewSpamEmail(store data.SpamEmailStore, logger logrus.FieldLogger, feedList []string) *SpamEmail {
	once.Do(func() {
		instance = &SpamEmail{
			logger: logger,
			store:  store,
		}
		logger.Debug("starting SpamEmail")
		instance.AddProvider(providers.NewTxtDomains(logger, feedList))
		go instance.run()
	})
	return instance
}
