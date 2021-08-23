package spam

import (
	"github.com/domgolonka/foretoken/app/data"
	"github.com/domgolonka/foretoken/lib/scrapers/ip/spam/providers"
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
		iplist, err := provider.List()

		if err != nil {
			p.logger.Errorf("cannot load list of proxy %s err:%s", provider.Name(), err)
			continue
		}

		p.logger.Println(provider.Name(), len(iplist))

		for _, s := range iplist {
			p.createOrIgnore(s.IP, s.Prefix, s.Score, s.Type)
			p.hosts = append(p.hosts, s.ToString())
		}
	}
}
func (p *Spam) createOrIgnore(ip string, prefix uint8, score int, iptype string) bool {
	_, err := p.store.Create(ip, prefix, score, iptype)
	return err == nil
}

func (p *Spam) deleteOld(hour int) (bool, error) {
	return p.store.DeleteOld(hour)
}

func (p *Spam) Run(hours int) {
	go func() {
		_, err := instance.deleteOld(hours + 12)
		if err != nil {
			p.logger.Error(err)
		}
	}()
	go p.load()

}

func (p *Spam) Get() []string {
	return p.hosts

}
func NewSpam(store data.SpamStore, hours int, logger logrus.FieldLogger, feedList []string) *Spam {
	once.Do(func() {
		instance = &Spam{
			logger: logger,
			store:  store,
		}

		logger.Debug("starting IP Spam")
		instance.AddProvider(providers.NewTxtDomains(logger, feedList))
		go instance.Run(hours)

	})
	return instance
}
