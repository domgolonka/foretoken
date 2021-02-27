package tor

import (
	"github.com/domgolonka/foretoken/app/data"
	"github.com/domgolonka/foretoken/lib/scrapers/ip/tor/providers"
	"github.com/sirupsen/logrus"

	"reflect"
	"sync"
)

var (
	instance *Tor
	once     sync.Once
)

type Tor struct {
	providers []Provider
	store     data.TorStore
	logger    logrus.FieldLogger
}

func (p *Tor) isProvider(provider Provider) bool {
	for _, pr := range p.providers {
		if reflect.TypeOf(pr) == reflect.TypeOf(provider) {
			return true
		}
	}
	return false
}
func (p *Tor) AddProvider(provider Provider) {
	if !p.isProvider(provider) {
		p.providers = append(p.providers, provider)
	}
}
func (p *Tor) load() {
	for _, provider := range p.providers {
		hosts, err := provider.List()

		if err != nil {
			p.logger.Errorln(err)
		}
		p.logger.Println(provider.Name(), len(hosts))
		for i := 0; i < len(hosts); i++ {
			p.createOrIgnore(hosts[i].IP, hosts[i].Prefix, hosts[i].Type, hosts[i].Score)
		}
	}
}
func (p *Tor) createOrIgnore(ip string, prefix byte, iptype string, score int) bool {
	_, err := p.store.Create(ip, prefix, iptype, score)
	return err == nil
}

func (p *Tor) Run(hours int) {
	go func() {
		_, err := instance.deleteOld(hours + 12)
		if err != nil {
			p.logger.Error(err)
		}
	}()
	go p.load()

}

func (p *Tor) deleteOld(hour int) (bool, error) {
	return p.store.DeleteOld(hour)
}

func (p *Tor) Get() (*[]string, error) {
	return p.store.FindAllIPs()

}
func NewTor(store data.TorStore, hours int, logger logrus.FieldLogger) *Tor {
	once.Do(func() {
		instance = &Tor{
			logger: logger,
			store:  store,
		}
		logger.Debug("starting Tor")
		instance.AddProvider(providers.NewTxtDomains(logger))
		go instance.Run(hours)
	})
	return instance
}
