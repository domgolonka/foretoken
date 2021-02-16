package proxy

import (
	"math/rand"
	"reflect"
	"sync"
	"time"

	"github.com/domgolonka/threatdefender/app/data"

	"github.com/domgolonka/threatdefender/lib/scrapers/ip/proxy/providers"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

var (
	instance  *ProxyGenerator
	usedProxy sync.Map
	once      sync.Once
)

type Verify func(proxy string) bool

type ProxyGenerator struct { //nolint
	lastValidProxy string
	cache          *cache.Cache
	logger         logrus.FieldLogger
	VerifyFn       Verify
	store          data.ProxyStore
	providers      []Provider
	proxy          chan string
	job            chan string
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

func shuffle(vals []string) {
	r := rand.New(rand.NewSource(time.Now().Unix())) //nolint
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func (p *ProxyGenerator) createOrIgnore(dis string, ptype string) bool {
	_, err := p.store.Create(dis, ptype)
	return err == nil
}

func (p *ProxyGenerator) load() {
	for {
		for _, provider := range p.providers {
			usedProxy.Store(p.lastValidProxy, time.Now().Hour())
			provider.SetProxy(p.lastValidProxy)

			ips, err := provider.List()
			if err != nil {
				p.lastValidProxy = ""
				p.logger.Errorf("cannot load list of proxy %s err:%s", provider.Name(), err)
				continue
			}

			//p.logger.Println(provider.Name(), len(ips))

			usedProxy.Range(func(key, value interface{}) bool {
				if value.(int) != time.Now().Hour() {
					usedProxy.Delete(key)
				}
				return true
			})

			//p.logger.Debugf("provider %s found ips %d", provider.Name(), len(ips))
			shuffle(ips)
			for _, proxy := range ips {
				//p.job <- proxy
				p.createOrIgnore(proxy, "ipv4")
			}
		}
	}
}

func (p *ProxyGenerator) GetLast() string {
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
	var items []string
	for k := range p.cache.Items() {
		items = append(items, k)
	}
	return items
}

func (p *ProxyGenerator) verifyWithCache(proxy string) bool {
	val, found := p.cache.Get(proxy)
	if found {
		return val.(bool)
	}
	res := p.VerifyFn(proxy)
	p.cache.Set(proxy, res, cache.DefaultExpiration)
	return res
}

func (p *ProxyGenerator) do(proxy string) {
	if p.verifyWithCache(proxy) {
		p.proxy <- proxy
	}
}

func (p *ProxyGenerator) worker() {
	for proxy := range p.job {
		p.do(proxy)
	}
}

func (p *ProxyGenerator) run(workers int) {
	go p.load()

	for w := 1; w <= workers; w++ {
		go p.worker()
	}
}

func New(store data.ProxyStore, workers int, cacheminutes time.Duration, logger logrus.FieldLogger) *ProxyGenerator {
	once.Do(func() {
		instance = &ProxyGenerator{
			cache:    cache.New(cacheminutes*time.Minute, 5*time.Minute),
			VerifyFn: verifyProxy,
			proxy:    make(chan string),
			store:    store,
			logger:   logger,
			job:      make(chan string, 100),
		}

		//add providers to generator
		instance.AddProvider(providers.NewFreeProxyList())
		instance.AddProvider(providers.NewXseoIn())
		instance.AddProvider(providers.NewFreeProxyListNet())
		instance.AddProvider(providers.NewProxyList())
		instance.AddProvider(providers.NewTxtDomains())

		//instance.AddProvider(providers.NewHidemyName())
		instance.AddProvider(providers.NewCoolProxy())
		//instance.AddProvider(providers.NewProxyTech())
		instance.AddProvider(providers.NewPubProxy())
		//run workers
		go instance.run(workers)
	})
	return instance
}
