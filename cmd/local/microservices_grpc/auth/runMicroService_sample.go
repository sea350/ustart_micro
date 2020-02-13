package main

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq"
	authapi "github.com/sea350/ustart_micro/backend/api/grpc/auth"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config authapi.Config
	//Importing configuration from json
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	//Generating api object
	authService, err := authapi.New(&config)
	if err != nil {
		panic(err)
	}

	authService.Run()

}