package zookeeper

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/domgolonka/foretoken/lib/registry"
	"github.com/samuel/go-zookeeper/zk"
	"google.golang.org/grpc/grpclog"
)

type Config struct {
	ZkServers []string
	Directory string
	Timeout   time.Duration
}

type Registrar struct {
	sync.RWMutex
	conf     *Config
	conn     *zk.Conn
	canceler map[string]context.CancelFunc
	logger   logrus.FieldLogger
}

func New(conf *Config, logger logrus.FieldLogger) (*Registrar, error) {
	reg := &Registrar{
		conf:     conf,
		canceler: make(map[string]context.CancelFunc),
		logger:   logger,
	}
	c, err := connect(conf.ZkServers, conf.Timeout)
	if err != nil {
		return nil, err
	}
	reg.conn = c
	return reg, nil
}

func connect(zkServers []string, sessionTimeout time.Duration) (*zk.Conn, error) {
	c, event, err := zk.Connect(zkServers, sessionTimeout)
	if err != nil {
		return nil, err
	}
	timer := time.NewTimer(time.Second * 10)
	for {
		select {
		case e := <-event:
			if e.State == zk.StateConnected {
				return c, nil
			}
		case <-timer.C:
			return nil, errors.New("connect zk timeout")
		}
	}
}

// create node one by one
// zk not support "mkdir -p"
func (r *Registrar) register(path string, nodeInfo string) error {
	znodes := strings.Split(path, "/")
	var onepath string
	for i, znode := range znodes {
		if znode == "" {
			continue
		}
		onepath = onepath + "/" + znode
		exists, _, _ := r.conn.Exists(onepath)
		if exists {
			continue
		}
		var err error
		if i != len(znodes)-1 {
			err = createTempNode(r.conn, onepath)
		} else {
			err = createTemporaryNode(r.conn, onepath, nodeInfo)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Registrar) Register(service *registry.Info) error {
	path := r.conf.Directory + "/" + service.Name + "/" + service.Version + "/" + service.ID
	data, _ := json.Marshal(service)
	err := r.register(path, string(data))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	r.Lock()
	r.canceler[service.ID] = cancel
	r.Unlock()

	r.keepalive(ctx, path, string(data))

	return err
}

func (r *Registrar) keepalive(ctx context.Context, path, value string) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if r.conn.State() != zk.StateHasSession {
				err := r.register(path, value)
				if err != nil {
					grpclog.Errorf("Registrar register error, %v\n", err.Error())
				}
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
	r.conn.Close()
	r.conn = nil
}

// create temporary node
func createTemporaryNode(conn *zk.Conn, path string, nodeInfo string) error {
	_, err := conn.Create(path, []byte(nodeInfo), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return err
}

// create node
func createTempNode(conn *zk.Conn, path string) error {
	_, err := conn.Create(path, nil, 0, zk.WorldACL(zk.PermAll))
	return err
}
