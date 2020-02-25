package main

import (
	"log"
	"os"

	"github.com/sea350/ustart_micro/backend"
	"github.com/sea350/ustart_micro/backend/session"
	"github.com/sea350/ustart_micro/backend/session/storage"
	sqlstore "github.com/sea350/ustart_micro/backend/session/storage/sql"
)

func main() {
	log.SetPrefix("Backend Command, ")
	log.Println("Loading config...")
	cfg := &backend.Config{
		Port:               os.Getenv("BACKEND_PORT"),
		AuthAPIAdress:      "http://localhost:" + os.Getenv("USTART_AUTH_PORT"),
		ProfileAPIAdresss:  "http://localhost:" + os.Getenv("USTART_PROF_PORT"),
		ProjectAPIAddress:  "http://localhost:" + os.Getenv("USTART_PROJ_PORT"),
		UploaderAPIAddress: "http://localhost:" + os.Getenv("USTART_UPLOADER_PORT"),
		SessionConfig: session.Config{
			StorageConfig: &storage.Config{
				SQLConfig: &sqlstore.Config{
					DriverName: "postgres",
					Host:       "ustart-postgres-test.crgo1xxhh8o5.us-east-1.rds.amazonaws.com",
					Port:       "5432",
					DBName:     "ustart-postgres-test",
					Username:   "ustart",
					Password:   "imadirtyweeb",
				},
			},
			SessionKey: "session_please",
		},
	}

	log.Println("Creating new backend service from config...")
	srv, err := backend.New(cfg)
	if err != nil {
		panic(err)
	}

	log.Println("Running server...")

	if err := srv.Run(); err != nil {
		log.Printf("Backend server exited with error {%v}\n", err)
		os.Exit(1)
	}

	log.Println("Backend server died peacefully")
}
