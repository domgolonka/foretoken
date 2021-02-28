package services

import (
	"time"

	"github.com/domgolonka/foretoken/lib/registry/zookeeper"

	"github.com/coreos/etcd/clientv3"
	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/lib/registry"
	"github.com/domgolonka/foretoken/lib/registry/consul"
	"github.com/domgolonka/foretoken/lib/registry/etcd3"
	consulapi "github.com/hashicorp/consul/api"
)

type Registrar struct {
	Service         *registry.Info
	Etc3Registrar   *etcd3.Registrar
	ConsulRegistrar *consul.Registrar
	ZKRegistrar     *zookeeper.Registrar
}

func startEtcd(app *app.App) *Registrar {
	etcdConfg := clientv3.Config{
		Endpoints: []string{app.Config.ServiceDiscovery.Endpoint},
	}

	service := &registry.Info{
		ID:      app.Config.ServiceDiscovery.NodeID,
		Name:    "foretoken",
		Version: "1.0",
		Address: app.Config.GRPCAddress,
	}

	registrar, err := etcd3.New(
		&etcd3.Config{
			Etcd:      etcdConfg,
			Directory: "/assets/storage",
			TTL:       10 * time.Second,
		}, app.Logger)
	if err != nil {
		app.Logger.Panic(err)
		return nil
	}

	return &Registrar{
		Service:       service,
		Etc3Registrar: registrar,
	}

}

func startConsul(app *app.App) *Registrar {
	consulConfg := &consulapi.Config{
		Address: app.Config.ServiceDiscovery.Endpoint,
	}

	service := &registry.Info{
		ID:      app.Config.ServiceDiscovery.NodeID,
		Name:    "foretoken",
		Version: "1.0",
		Address: app.Config.GRPCAddress,
	}

	registrar, err := consul.New(
		&consul.Config{
			ConsulCfg: consulConfg,
			TTL:       5,
		}, app.Logger)
	if err != nil {
		app.Logger.Panic(err)
		return nil
	}

	return &Registrar{
		Service:         service,
		ConsulRegistrar: registrar,
	}

}

func startZookeeper(app *app.App) *Registrar {

	service := &registry.Info{
		ID:      app.Config.ServiceDiscovery.NodeID,
		Name:    "foretoken",
		Version: "1.0",
		Address: app.Config.GRPCAddress,
	}

	registrar, err := zookeeper.New(
		&zookeeper.Config{
			ZkServers: []string{app.Config.ServiceDiscovery.Endpoint},
			Directory: "/assets/storage",
			Timeout:   time.Second,
		}, app.Logger)
	if err != nil {
		app.Logger.Panic(err)
		return nil
	}

	return &Registrar{
		Service:     service,
		ZKRegistrar: registrar,
	}

}
