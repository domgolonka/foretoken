package services

import (
	"net"

	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/lib/grpc/impl"
	"github.com/domgolonka/foretoken/lib/grpc/proto"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func ServeRPC(app *app.App, ch chan bool, address string) {
	s := grpc.NewServer()

	l, err := net.Listen("tcp", address)
	if err != nil {
		app.Logger.Panic(err)
	}
	proto.RegisterEmailServiceServer(s, impl.EmailService)
	proto.RegisterIPServiceServer(s, impl.IPService)

	grpc_prometheus.Register(s)

	if err = s.Serve(l); err != nil {
		ch <- false
		app.Logger.Panic(err)
	}
}
