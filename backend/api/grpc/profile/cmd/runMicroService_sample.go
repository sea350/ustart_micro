package main

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq"
	profileapi "github.com/sea350/ustart_micro/backend/api/grpc/profile"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config profileapi.Config
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
	profileService, err := profileapi.New(&config)
	if err != nil {
		panic(err)
	}

	profileService.Run()
}
