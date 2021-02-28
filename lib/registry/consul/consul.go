package consul

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/domgolonka/foretoken/lib/registry"
	consul "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/grpclog"
)

type Registrar struct {
	sync.RWMutex
	client   *consul.Client
	cfg      *Config
	canceler map[string]context.CancelFunc
	logger   logrus.FieldLogger
}

type Config struct {
	ConsulCfg *consul.Config
	TTL       int //ttl seconds
}

func New(cfg *Config, logger logrus.FieldLogger) (*Registrar, error) {
	c, err := consul.NewClient(cfg.ConsulCfg)
	if err != nil {
		return nil, err
	}
	return &Registrar{
		canceler: make(map[string]context.CancelFunc),
		client:   c,
		cfg:      cfg,
		logger:   logger,
	}, nil
}

func (r *Registrar) Register(service *registry.Info) error {
	// register service
	metadata, err := json.Marshal(service.Metadata)
	if err != nil {
		return err
	}
	tags := make([]string, 0)
	tags = append(tags, string(metadata))

	register := func() error {
		regis := &consul.AgentServiceRegistration{
			ID:      service.ID,
			Name:    service.Name + ":" + service.Version,
			Address: service.Address,
			Tags:    tags,
			Check: &consul.AgentServiceCheck{
				TTL:                            fmt.Sprintf("%ds", r.cfg.TTL),
				Status:                         consul.HealthPassing,
				DeregisterCriticalServiceAfter: "1m",
			}}
		err := r.client.Agent().ServiceRegister(regis)
		if err != nil {
			r.logger.Errorf("register service to consul error: %s", err.Error())
			return err
		}
		return nil
	}

	err = register()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())

	r.Lock()
	r.canceler[service.ID] = cancel
	r.Unlock()

	keepAliveTicker := time.NewTicker(time.Duration(r.cfg.TTL) * time.Second / 5)
	registerTicker := time.NewTicker(time.Minute)

	for {
		select {
		case <-ctx.Done():
			keepAliveTicker.Stop()
			registerTicker.Stop()
			err = r.client.Agent().ServiceDeregister(service.ID)
			if err != nil {
				grpclog.Errorf("service deregister consul error: %s", err.Error())
			}
			return nil
		case <-keepAliveTicker.C:
			err := r.client.Agent().PassTTL("service:"+service.ID, "")
			if err != nil {
				grpclog.Infof("consul registry check %v.\n", err)
			}
		case <-registerTicker.C:
			err = register()
			if err != nil {
				grpclog.Infof("consul register service error: %v.\n", err)
			}
		}
	}

}

func (r *Registrar) Unregister(service *registry.Info) error {
	r.RLock()
	cancel, ok := r.canceler[service.ID]
	r.RUnlock()

	if ok {
		cancel()
	}
	return nil
}

func (r *Registrar) Close() {
}
