package main

import (
	"log"
	"net/http"

	"github.com/sea350/ustart_mono/backend/api/rest/auth"
	sAuth "github.com/sea350/ustart_mono/backend/auth"
	"github.com/sea350/ustart_mono/backend/auth/storage"
	"github.com/sea350/ustart_mono/backend/auth/storage/sql"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	restAPI, err := auth.New(&auth.Config{
		AuthCfg: &sAuth.Config{
			StorageConfig: &storage.Config{
				SQLConfig: &sqlstore.Config{
					DriverName:         "postgresql",
					Host:               "localhost",
					Port:               "5432",
					DBName:             "localhost",
					Username:           "admin",
					Password:           "password",
					RegistryTable:      "auth",
					LoginTrackingTable: "logins",
				},
			},
		},
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")
	http.HandleFunc("/", nil)
	http.HandleFunc("/authenticate", restAPI.Authenticate)
	http.HandleFunc("/changepassword", restAPI.ChangePassword)
	http.HandleFunc("/register", restAPI.Register)

	http.ListenAndServe(":5001", nil)
}
