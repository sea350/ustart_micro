package backend

import (
	"log"
	"net/http"

	"github.com/sea350/ustart_micro/backend/session"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
	"github.com/sea350/ustart_micro/backend/profile/profilepb"
	"google.golang.org/grpc"
)

// Server is a centrialized service providing access to all of UStart's microservices
type Server struct {
	port          string
	authClient    *authpb.AuthClient
	profileClient *profilepb.ProfileClient
	sesh          *session.Session
}

// New returns a new backend server, given the config object
func New(cfg *Config) (*Server, error) {
	server := &Server{port: cfg.Port}

	// creates all api clients
	//first auth
	authConn, err := grpc.Dial(cfg.AuthAPIAdress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	authClient := authpb.NewAuthClient(authConn)
	server.authClient = &authClient

	//then profile
	profileConn, err := grpc.Dial(cfg.ProfileAPIAdresss, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	profileClient := profilepb.NewProfileClient(profileConn)
	server.profileClient = &profileClient

	//next session
	server.sesh, err = session.New(&cfg.SessionConfig)
	if err != nil {
		panic(err)
	}

	return server, nil

}

// Run starts the backend http server
func (srv *Server) Run() error {
	log.SetPrefix("Backend Server:")
	log.Println("Booting...")

	http.HandleFunc("/", nil)

	http.HandleFunc("/Registration/Register", srv.SignupHTTP)

	http.HandleFunc("/Authentication/Login", srv.LoginHTTP)
	http.HandleFunc("/Authentication/Logoout", nil)

	http.HandleFunc("/Profile/UserPage", srv.PublicProfileHTTP)

	log.Printf("Listening on %s\n", srv.port)
	return http.ListenAndServe(":"+srv.port, nil)
}
