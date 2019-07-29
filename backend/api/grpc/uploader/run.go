package uploaderapi

import (
	"net"

	"github.com/sea350/ustart_micro/backend/uploader/uploaderpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Run runs the api server
func (uAPI *GRPCAPI) Run() {
	listener, err := net.Listen("tcp", ":"+uAPI.port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	uploaderpb.RegisterUploaderServer(srv, uAPI.upload)
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
