package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/sea350/ustart_micro/backend"
	authapi "github.com/sea350/ustart_micro/backend/api/grpc/auth"
	profileapi "github.com/sea350/ustart_micro/backend/api/grpc/profile"
)

func main() {
	log.SetPrefix("Backend Command, ")
	log.Println("Loading config...")
	cfg := &backend.Config{
		Port:              "5000",
		AuthAPIAdress:     "localhost:5101",
		ProfileAPIAdresss: "localhost:5102",
	}

	log.Println("Setting up auth microservice...")
	var authConfig authapi.Config
	//Importing configuration from json
	authFile, err := os.Open("auth-config.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(authFile).Decode(&authConfig)
	if err != nil {
		panic(err)
	}

	//Generating api object
	authService, err := authapi.New(&authConfig)
	if err != nil {
		panic(err)
	}

	log.Println("Setting up profile microservice...")
	var profConfig profileapi.Config
	//Importing configuration from json
	profFile, err := os.Open("prof-config.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(profFile).Decode(&profConfig)
	if err != nil {
		panic(err)
	}

	//Generating api object
	profService, err := profileapi.New(&profConfig)
	if err != nil {
		panic(err)
	}

	log.Println("Creating new backend service from config...")
	srv, err := backend.New(cfg)
	if err != nil {
		panic(err)
	}

	log.Println("Running auth...")
	go authService.Run()

	log.Println("Running profile...")
	go profService.Run()

	log.Println("Running server...")

	if err := srv.Run(); err != nil {
		log.Printf("Backend server exited with error {%v}\n", err)
		os.Exit(1)
	}

	log.Println("Backend server died peacefully")
}
