package profileapi

import (
	"net"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
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
	profilepb.RegisterProfileServer(srv, pAPI.prof)
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
