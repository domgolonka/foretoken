package dc

import (
	"github.com/domgolonka/threatdefender/app/data"
	"github.com/sirupsen/logrus"

	"reflect"
	"sync"
)

// Datacenter list
var (
	instance *DC
	once     sync.Once
)

type DC struct {
	providers []Provider
	//hosts     chan []string
	store  data.VpnStore
	logger logrus.FieldLogger
}

func (p *DC) isProvider(provider Provider) bool {
	for _, pr := range p.providers {
		if reflect.TypeOf(pr) == reflect.TypeOf(provider) {
			return true
		}
	}
	return false
}
func (p *DC) AddProvider(provider Provider) {
	if !p.isProvider(provider) {
		p.providers = append(p.providers, provider)
	}
}
func (p *DC) load() {
	for _, provider := range p.providers {
		hosts, err := provider.List()

		if err != nil {
			p.logger.Errorln(err)
		}
		p.logger.Println(provider.Name(), len(hosts))
		//p.hosts <- hosts
		for i := 0; i < len(hosts); i++ {
			p.createOrIgnore(hosts[i])
		}
	}
}
func (p *DC) createOrIgnore(dc string) bool {
	_, err := p.store.Create(dc)
	return err == nil
}

func (p *DC) run() {
	go p.load()
}

func (p *DC) Get() (*[]string, error) {
	return p.store.FindAll()

}
func NewDC(store data.VpnStore, logger logrus.FieldLogger) *DC {
	once.Do(func() {
		instance = &DC{
			logger: logger,
			store:  store,
		}
		logger.Debug("starting DC")
		//instance.AddProvider(providers.todo(logger))
		go instance.run()
	})
	return instance
}
