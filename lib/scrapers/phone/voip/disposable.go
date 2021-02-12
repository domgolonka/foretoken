package voip

import (
	"github.com/domgolonka/threatscraper/app/data"
	"github.com/domgolonka/threatscraper/lib/scrapers/phone/voip/providers"
	"github.com/sirupsen/logrus"

	"reflect"
	"sync"
)

var (
	instance *Disposable
	once     sync.Once
)

type Disposable struct {
	providers []Provider
	store     data.DisposableStore
	logger    logrus.FieldLogger
}

func (p *Disposable) isProvider(provider Provider) bool {
	for _, pr := range p.providers {
		if reflect.TypeOf(pr) == reflect.TypeOf(provider) {
			return true
		}
	}
	return false
}
func (p *Disposable) AddProvider(provider Provider) {
	if !p.isProvider(provider) {
		p.providers = append(p.providers, provider)
	}
}
func (p *Disposable) load() {
	for _, provider := range p.providers {

		hosts, err := provider.List()

		if err != nil {
			p.logger.Errorln(err)
		}
		p.logger.Println(provider.Name(), len(hosts))

		for i := 0; i < len(hosts); i++ {
			p.createOrIgnore(hosts[i])
		}
	}
}
func (p *Disposable) createOrIgnore(dis string) bool {
	_, err := p.store.Create(dis)
	if err != nil {
		return false
	}
	return true
}

func (p *Disposable) run() {
	go p.load()
}

func (p *Disposable) Get() (*[]string, error) {

	return p.store.FindAll()

}
func NewDisposable(store data.DisposableStore, logger logrus.FieldLogger) *Disposable {
	once.Do(func() {
		instance = &Disposable{
			logger: logger,
			store:  store,
		}
		logger.Debug("starting Disposable")
		instance.AddProvider(providers.NewTxtDomains(logger))
		go instance.run()
	})
	return instance
}
