package authapi

import (
	"net"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Run runs the api server
func (pAPI *GPRCAPI) Run() {
	listener, err := net.Listen("tcp", ":"+pAPI.port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	authpb.RegisterAuthServer(srv, pAPI.auth)
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
