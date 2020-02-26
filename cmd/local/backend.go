package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/sea350/ustart_micro/backend"
	"github.com/sea350/ustart_micro/backend/session"
)

func main() {
	log.SetPrefix("Backend Command, ")
	log.Println("Loading config...")

	var config session.Config
	//Importing configuration from json
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	cfg := &backend.Config{
		Port:               os.Getenv("BACKEND_PORT"),
		AuthAPIAdress:      "http://localhost:" + os.Getenv("USTART_AUTH_PORT"),
		ProfileAPIAdresss:  "http://localhost:" + os.Getenv("USTART_PROF_PORT"),
		ProjectAPIAddress:  "http://localhost:" + os.Getenv("USTART_PROJ_PORT"),
		UploaderAPIAddress: "http://localhost:" + os.Getenv("USTART_UPLOADER_PORT"),
		SessionConfig:      config,
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
