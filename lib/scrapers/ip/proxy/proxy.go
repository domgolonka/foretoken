package proxy

import (
	"crypto/rand"
	"math/big"
	"reflect"
	"sync"
	"time"

	"github.com/domgolonka/foretoken/app/models"

	"github.com/domgolonka/foretoken/app/data"

	"github.com/domgolonka/foretoken/lib/scrapers/ip/proxy/providers"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

var (
	instance  *ProxyGenerator
	usedProxy sync.Map
	once      sync.Once
)

type Verify func(proxy models.Proxy) bool

type ProxyGenerator struct { //nolint
	lastValidProxy models.Proxy
	cache          *cache.Cache
	logger         logrus.FieldLogger
	VerifyFn       Verify
	store          data.ProxyStore
	providers      []Provider
	proxy          chan models.Proxy
	job            chan models.Proxy
}

func (p *ProxyGenerator) isProvider(provider Provider) bool {
	for _, pr := range p.providers {
		if reflect.TypeOf(pr) == reflect.TypeOf(provider) {
			return true
		}
	}
	return false
}

func (p *ProxyGenerator) AddProvider(provider Provider) {
	if !p.isProvider(provider) {
		p.providers = append(p.providers, provider)
	}
}

func shuffle(vals []models.Proxy) {

	//r := rand.New(rand.NewSource(time.Now().Unix())) //nolint
	for len(vals) > 0 {
		n := len(vals)
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
		if err != nil {
			panic(err)
		}
		randIndex := nBig.Int64()
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func (p *ProxyGenerator) createOrIgnore(ip, port, ptype string) bool {
	_, err := p.store.Create(ip, port, ptype)
	return err == nil
}

func (p *ProxyGenerator) load() {
	for _, provider := range p.providers {
		usedProxy.Store(p.lastValidProxy, time.Now().Hour())
		provider.SetProxy(p.lastValidProxy)

		ips, err := provider.List()
		if err != nil {
			p.lastValidProxy = models.Proxy{}
			p.logger.Errorf("cannot load list of proxy %s err:%s", provider.Name(), err)
			continue
		}

		// p.logger.Println(provider.Name(), len(ips))

		usedProxy.Range(func(key, value interface{}) bool {
			if value.(int) != time.Now().Hour() {
				usedProxy.Delete(key)
			}
			return true
		})

		// p.logger.Debugf("provider %s found ips %d", provider.Name(), len(ips))
		shuffle(ips)
		for _, proxy := range ips {
			p.createOrIgnore(proxy.IP, proxy.Port, proxy.Type)
		}
	}
}

func (p *ProxyGenerator) GetLast() models.Proxy {
	proxy := <-p.proxy
	_, ok := usedProxy.Load(proxy)
	if !ok {
		p.lastValidProxy = proxy
	}
	return proxy
}

func (p *ProxyGenerator) Count() int {

	return p.cache.ItemCount()
}
func (p *ProxyGenerator) Get() []string {
	items := make([]string, 0, len(p.cache.Items()))
	for k := range p.cache.Items() {
		items = append(items, k)
	}
	return items
}

func (p *ProxyGenerator) verifyWithCache(proxy models.Proxy) bool {
	val, found := p.cache.Get(proxy.ToString())
	if found {
		return val.(bool)
	}
	res := p.VerifyFn(proxy)
	p.cache.Set(proxy.ToString(), res, cache.DefaultExpiration)
	return res
}

func (p *ProxyGenerator) do(proxy models.Proxy) {
	if p.verifyWithCache(proxy) {
		p.proxy <- proxy
	}
}

func (p *ProxyGenerator) worker() {
	for proxy := range p.job {
		p.do(proxy)
	}
}

func (p *ProxyGenerator) deleteOld(hour int) (bool, error) {
	return p.store.DeleteOld(hour)
}

func (p *ProxyGenerator) Run(workers int, hours int) {
	go func() {
		_, err := instance.deleteOld(hours + 12)
		if err != nil {
			p.logger.Error(err)
		}
	}()
	go p.load()

	for w := 1; w <= workers; w++ {
		go p.worker()
	}
}

func New(store data.ProxyStore, workers int, cacheminutes time.Duration, hours int, logger logrus.FieldLogger, feedList []string) *ProxyGenerator {
	once.Do(func() {
		instance = &ProxyGenerator{
			cache:    cache.New(cacheminutes*time.Minute, 5*time.Minute),
			VerifyFn: verifyProxy,
			proxy:    make(chan models.Proxy),
			store:    store,
			logger:   logger,
			job:      make(chan models.Proxy, 100),
		}
		// add providers to generator
		instance.AddProvider(providers.NewFreeProxyList())
		instance.AddProvider(providers.NewXseoIn())
		instance.AddProvider(providers.NewProxyList())
		instance.AddProvider(providers.NewTxtDomains(logger, feedList))
		instance.AddProvider(providers.NewHidemyName())
		instance.AddProvider(providers.NewCoolProxy())
		instance.AddProvider(providers.NewPubProxy())
		// run workers
		go instance.Run(workers, hours)

	})
	return instance
}
