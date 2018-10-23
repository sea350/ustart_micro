package backend

import (
	"log"
	"net/http"

	"github.com/sea350/ustart_mono/backend/api/rest/auth"
)

// Server is a monolithic service providing access to all of UStart's data
type Server struct {
	port    string
	authSrv *auth.ElasticAuthRESTAPI
}

// New returns a new backend server, given the config object
func New(cfg *Config) (*Server, error) {
	// creates all api servers

	authSrv, err := auth.New(&auth.Config{})
	if err != nil {
		return nil, err
	}

	return &Server{
		authSrv: authSrv,
	}, err
}

// Run starts the backend http server
func (srv *Server) Run() error {
	log.SetPrefix("Backend Server:")
	log.Println("Booting...")

	http.HandleFunc("/auth/register", srv.authSrv.Register)

	log.Printf("Listening on %s\n", srv.port)
	return http.ListenAndServe(":"+srv.port, nil)
}
