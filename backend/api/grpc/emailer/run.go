package emailerapi

import (
	"net"

	"github.com/sea350/ustart_micro/backend/emailer/emailerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Run runs the api server
func (eAPI *GRPCAPI) Run() {
	listener, err := net.Listen("tcp", ":"+eAPI.port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	emailerpb.RegisterAuthServer(srv, eAPI.auth)
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
