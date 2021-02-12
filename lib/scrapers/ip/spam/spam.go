package spam

import (
	"github.com/domgolonka/threatscraper/app/data"
	"github.com/domgolonka/threatscraper/lib/scrapers/ip/spam/providers"
	"github.com/sirupsen/logrus"

	"reflect"
	"sync"
)

var (
	instance *Spam
	once     sync.Once
)

type Spam struct {
	providers []Provider
	store     data.SpamStore
	hosts     []string
	logger    logrus.FieldLogger
}

func (p *Spam) isProvider(provider Provider) bool {
	for _, pr := range p.providers {
		if reflect.TypeOf(pr) == reflect.TypeOf(provider) {
			return true
		}
	}
	return false
}
func (p *Spam) AddProvider(provider Provider) {
	if !p.isProvider(provider) {
		p.providers = append(p.providers, provider)
	}
}
func (p *Spam) load() {
	for _, provider := range p.providers {
		iplist, subnetlist, err := provider.List()

		if err != nil {
			p.logger.Errorf("cannot load list of proxy %s err:%s", provider.Name(), err)
			continue
		}

		p.logger.Println(provider.Name(), len(iplist))

		p.hosts = append(p.hosts, subnetlist...)
		//p.hosts <- hosts
		for _, s := range p.hosts {
			p.createOrIgnore(s, true)
		}
	}
}
func (p *Spam) createOrIgnore(dis string, sub bool) bool {
	_, err := p.store.Create(dis, sub)
	return err == nil
}

func (p *Spam) run() {
	go p.load()
}

func (p *Spam) Get() []string {
	return p.hosts

}
func NewSpam(store data.SpamStore, logger logrus.FieldLogger) *Spam {
	once.Do(func() {
		instance = &Spam{
			logger: logger,
			store:  store,
		}
		logger.Debug("starting Spam")
		instance.AddProvider(providers.NewTxtDomains(logger))
		go instance.run()
	})
	return instance
}
