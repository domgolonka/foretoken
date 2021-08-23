package services

import (
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/lib/grpc/impl"
	"github.com/domgolonka/foretoken/lib/grpc/proto"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

const (
	CONSUL    = "consul"
	ETCD3     = "etcd3"
	ZOOKEEPER = "zookeeper"
)

func ServeRPC(app *app.App, ch chan bool) {
	s := grpc.NewServer()

	l, err := net.Listen("tcp", app.Config.GRPCAddress)
	if err != nil {
		app.Logger.Panic(err)
	}
	proto.RegisterEmailServiceServer(s, impl.EmailSrv)
	proto.RegisterIPServiceServer(s, impl.IPSrv)

	grpc_prometheus.Register(s)

	if err = s.Serve(l); err != nil {
		ch <- false
		app.Logger.Panic(err)
	}

	regist := &Registrar{}
	wg := sync.WaitGroup{}

	switch srv := app.Config.ServiceDiscovery.Service; srv {
	case ETCD3:
		regist = startEtcd(app)

		wg.Add(1)
		go func() {
			err = regist.Etc3Registrar.Register(regist.Service)
			if err != nil {
				ch <- false
				app.Logger.Error("cannot register etcd3: %s", err.Error())
			}
			wg.Done()
		}()
	case CONSUL:
		regist = startConsul(app)

		wg.Add(1)
		go func() {
			err = regist.ConsulRegistrar.Register(regist.Service)
			if err != nil {
				ch <- false
				app.Logger.Error("cannot register consul: %s", err.Error())
			}
			wg.Done()
		}()
	case ZOOKEEPER:
		regist = startZookeeper(app)

		wg.Add(1)
		go func() {
			err = regist.ZKRegistrar.Register(regist.Service)
			if err != nil {
				ch <- false
				app.Logger.Error("cannot register zookeeper: %s", err.Error())
			}
			wg.Done()
		}()
	default:
		ch <- false
		app.Logger.Panic("service is unknown type")
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	if app.Config.ServiceDiscovery.Service == ETCD3 {
		err = regist.Etc3Registrar.Unregister(regist.Service)
		if err != nil {
			ch <- false
			app.Logger.Error("cannot unregister etcd3: %s", err.Error())
		}
	} else if app.Config.ServiceDiscovery.Service == CONSUL {
		err = regist.ConsulRegistrar.Unregister(regist.Service)
		if err != nil {
			ch <- false
			app.Logger.Error("cannot unregister consul: %s", err.Error())
		}
	} else if app.Config.ServiceDiscovery.Service == ZOOKEEPER {
		err = regist.ZKRegistrar.Unregister(regist.Service)
		if err != nil {
			ch <- false
			app.Logger.Error("cannot unregister zookeeper: %s", err.Error())
		}
	}

	s.Stop()
	wg.Wait()
}
