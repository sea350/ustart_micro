package main

import (
	"log"
	"net/http"

	"github.com/sea350/ustart_mono/backend/api/rest/auth"
	sAuth "github.com/sea350/ustart_mono/backend/auth"
	"github.com/sea350/ustart_mono/backend/auth/storage"
	"github.com/sea350/ustart_mono/backend/auth/storage/sql"
)

var config = &auth.Config{
	AuthCfg: &sAuth.Config{
		StorageConfig: &storage.Config{
			SQLConfig: &sqlstore.Config{
				DriverName:         "postgres",
				Host:               "localhost",
				Port:               "5432",
				DBName:             "localhost",
				Username:           "postgres",
				Password:           "postgres",
				RegistryTable:      "auth",
				LoginTrackingTable: "logins",
			},
		},
	},
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	restAPI, err := auth.New(config)
	if err != nil {
		panic(err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/", nil)
	http.HandleFunc("/authenticate", restAPI.Authenticate)
	http.HandleFunc("/changepassword", restAPI.ChangePassword)
	http.HandleFunc("/register", restAPI.Register)

	http.ListenAndServe(":5001", nil)
}
