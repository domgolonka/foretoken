package tor

import (
	"github.com/domgolonka/threatdefender/app/data"
	"github.com/domgolonka/threatdefender/lib/scrapers/ip/tor/providers"
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
			p.createOrIgnore(hosts[i].IP, hosts[i].Prefix, hosts[i].Score)
		}
	}
}
func (p *Tor) createOrIgnore(ip string, prefix byte, score int) bool {
	_, err := p.store.Create(ip, prefix, score)
	if err != nil {
		logrus.Error(err)
	}
	return err == nil
}

func (p *Tor) run() {
	go p.load()
}

func (p *Tor) Get() (*[]string, error) {
	return p.store.FindAll()

}
func NewTor(store data.TorStore, logger logrus.FieldLogger) *Tor {
	once.Do(func() {
		instance = &Tor{
			logger: logger,
			store:  store,
		}
		logger.Debug("starting Tor")
		instance.AddProvider(providers.NewTxtDomains(logger))
		go instance.run()
	})
	return instance
}
