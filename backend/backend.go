package backend

import (
	"log"
	"net/http"

	"github.com/sea350/ustart_micro/backend/uploader/uploaderpb"

	"github.com/sea350/ustart_micro/backend/session"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
	"github.com/sea350/ustart_micro/backend/profile/profilepb"
	"github.com/sea350/ustart_micro/backend/project/projectpb"
	"google.golang.org/grpc"
)

// Server is a centrialized service providing access to all of UStart's microservices
type Server struct {
	port           string
	authClient     *authpb.AuthClient
	profileClient  *profilepb.ProfileClient
	sesh           *session.Session
	uploaderClient *uploaderpb.UploaderClient
	projectClient  *projectpb.ProjectClient
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
	autC := authpb.NewAuthClient(authConn)
	server.authClient = &autC

	//then profile
	profileConn, err := grpc.Dial(cfg.ProfileAPIAdresss, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	prfC := profilepb.NewProfileClient(profileConn)
	server.profileClient = &prfC

	//next session
	server.sesh, err = session.New(&cfg.SessionConfig)
	if err != nil {
		panic(err)
	}

	//after that, uploader
	uploaderConn, err := grpc.Dial(cfg.UploaderAPIAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	uplC := uploaderpb.NewUploaderClient(uploaderConn)
	server.uploaderClient = &uplC

	//furthermore, project
	projConn, err := grpc.Dial(cfg.ProjectAPIAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	prjC := projectpb.NewProjectClient(projConn)
	server.projectClient = &prjC

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
	http.HandleFunc("/Profile/ChangeBasicData", srv.ChangeProfileDataHTTP)
	http.HandleFunc("/Profile/ChangeIcon", srv.ProfileChangeIconHTTP)
	http.HandleFunc("/Profile/ChangeBanner", srv.ProfileChangeBannerHTTP)
	http.HandleFunc("/Profile/ChangeTags", srv.ProfileChangeTagsHTTP)
	http.HandleFunc("/Profile/ToggleAvailable", srv.ProfileToggleAvailableHTTP)
	http.HandleFunc("/Profile/ToggleProjectVisible", srv.ProfileToggleProjectVisibleHTTP)

	http.HandleFunc("/Project/New", srv.ProjectNewHTTP)
	http.HandleFunc("/Project/GetJoinRequests", srv.GetJoinersHTTP)
	http.HandleFunc("/Project/GetRoles", srv.GetRolesHTTP)
	http.HandleFunc("/Project/ApproveMember", srv.ApproveMemberHTTP)
	http.HandleFunc("/Project/RejectJoinRequest", srv.RejectJoinHTTP)
	http.HandleFunc("/Project/RemoveMember", srv.RemoveMemberHTTP)

	log.Printf("Listening on %s\n", srv.port)
	return http.ListenAndServe(":"+srv.port, nil)
}
