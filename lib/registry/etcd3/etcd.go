package etcd3

import (
	"context"
	"encoding/json"

	"github.com/coreos/etcd/clientv3"

	"sync"
	"time"

	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/domgolonka/foretoken/lib/registry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/grpclog"
)

type Registrar struct {
	sync.RWMutex
	config   *Config
	etcd3    *clientv3.Client
	canceler map[string]context.CancelFunc
	logger   logrus.FieldLogger
}

type Config struct {
	Etcd      clientv3.Config
	Directory string
	TTL       time.Duration
}

func New(config *Config, logger logrus.FieldLogger) (*Registrar, error) {
	client, err := clientv3.New(config.Etcd)
	if err != nil {
		return nil, err
	}
	return &Registrar{
		etcd3:    client,
		config:   config,
		canceler: make(map[string]context.CancelFunc),
		logger:   logger,
	}, nil

}

func (r *Registrar) Register(registry *registry.Info) error {
	val, err := json.Marshal(registry)
	if err != nil {
		return err
	}
	value := string(val)
	key := r.config.Directory + "/" + registry.Name + "/" + registry.Version + "/" + registry.ID

	ctx, cancel := context.WithCancel(context.Background())
	r.Lock()
	r.canceler[registry.ID] = cancel
	r.Unlock()
	insertFunc := func() error {
		resp, err := r.etcd3.Grant(ctx, int64(r.config.TTL/time.Second))
		if err != nil {
			r.logger.Errorf("[Register] %v\n", err.Error())
			return err
		}
		_, err = r.etcd3.Get(ctx, key)
		if err != nil {
			if err == rpctypes.ErrKeyNotFound {
				if _, err := r.etcd3.Put(ctx, key, value, clientv3.WithLease(resp.ID)); err != nil {
					grpclog.Infof("etcd3: set key '%s' with ttl to etcd3 failed: %s", key, err.Error())
				}
			} else {
				grpclog.Infof("etcd3: key '%s' connect to etcd3 failed: %s", key, err.Error())
			}
			return err
		}
		// refresh set to true for not notifying the watcher
		if _, err := r.etcd3.Put(ctx, key, value, clientv3.WithLease(resp.ID)); err != nil {
			grpclog.Infof("etcd3: refresh key '%s' with ttl to etcd3 failed: %s", key, err.Error())
			return err
		}
		return nil
	}

	err = insertFunc()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(r.config.TTL / 5)
	for {
		select {
		case <-ticker.C:
			err = insertFunc()
			if err != nil {
				grpclog.Errorf("[Register] %v\n", err.Error())
			}
		case <-ctx.Done():
			ticker.Stop()
			if _, err := r.etcd3.Delete(context.Background(), key); err != nil {
				grpclog.Infof("Unregister '%s' failed: %s", key, err.Error())
			}
			return nil
		}
	}

}

func (r *Registrar) Unregister(registry *registry.Info) error {
	r.RLock()
	cancel, ok := r.canceler[registry.ID]
	r.RUnlock()

	if ok {
		cancel()
	}
	return nil
}
func (r *Registrar) Close() {
	err := r.etcd3.Close()
	if err != nil {
		r.logger.Errorf("[Close] %v\n", err.Error())
	}
}
