package services

import (
	"fmt"
	"net"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/lib/services/impl"
	"github.com/domgolonka/threatdefender/lib/services/proto"
	"google.golang.org/grpc"
)

func ServeRPC(app *app.App, ch chan bool, port int) {
	s := grpc.NewServer()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		app.Logger.Panic(err)
	}
	proto.RegisterEmailServiceServer(s, impl.EmailService)
	proto.RegisterIPServiceServer(s, impl.IPService)

	if err = s.Serve(l); err != nil {
		ch <- false
		app.Logger.Panic(err)
	}
}
