package backend

import (
	"log"
	"net/http"

	authapi "github.com/sea350/ustart_mono/backend/api/rest/auth"
	auth "github.com/sea350/ustart_mono/backend/auth"
)

// Server is a monolithic service providing access to all of UStart's data
type Server struct {
	port    string
	authAPI *authapi.ElasticAuthRESTAPI
}

// New returns a new backend server, given the config object
func New(cfg *Config) (*Server, error) {
	// creates all api servers

	authAPI, err := authapi.New(&authapi.Config{
		EAuthCfg: auth.NewConfig(),
	})
	if err != nil {
		return nil, err
	}

	return &Server{
		authAPI: authAPI,
	}, err
}

// Run starts the backend http server
func (srv *Server) Run() error {
	log.SetPrefix("Backend Server:")
	log.Println("Booting...")

	http.HandleFunc("/auth/register", srv.authAPI.Register)

	log.Printf("Listening on %s\n", srv.port)
	return http.ListenAndServe(":"+srv.port, nil)
}
