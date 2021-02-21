package free

import (
	"github.com/domgolonka/threatdefender/app/data"
	"github.com/domgolonka/threatdefender/lib/scrapers/email/free/providers"
	"github.com/sirupsen/logrus"

	"reflect"
	"sync"
)

var (
	instance *Free
	once     sync.Once
)

type Free struct {
	providers []Provider
	store     data.FreeEmailStore
	logger    logrus.FieldLogger
}

func (p *Free) isProvider(provider Provider) bool {
	for _, pr := range p.providers {
		if reflect.TypeOf(pr) == reflect.TypeOf(provider) {
			return true
		}
	}
	return false
}
func (p *Free) AddProvider(provider Provider) {
	if !p.isProvider(provider) {
		p.providers = append(p.providers, provider)
	}
}
func (p *Free) load() {
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
func (p *Free) createOrIgnore(domain string, score int) bool {
	_, err := p.store.Create(domain, score)
	return err == nil
}

func (p *Free) run() {
	go p.load()
}

func (p *Free) Get() (*[]string, error) {
	return p.store.FindAll()
}

func NewFreeEmail(store data.FreeEmailStore, logger logrus.FieldLogger) *Free {
	once.Do(func() {
		instance = &Free{
			logger: logger,
			store:  store,
		}
		logger.Debug("starting Free Email")
		instance.AddProvider(providers.NewTxtDomains(logger))
		go instance.run()
	})
	return instance
}
