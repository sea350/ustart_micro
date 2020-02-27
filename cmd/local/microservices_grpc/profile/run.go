package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

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
		log.Panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Panic(err)
	}

	portString := os.Getenv("USTART_PROF_PORT")
	if portString != "" {
		port, err := strconv.Atoi(portString)
		if err != nil {
			log.Panic(err)
		}
		config.Port = port
	}

	log.Println("Displaying config data")
	log.Println("DB Host: ", config.ProfileCfg.StorageConfig.ElasticConfig.ElasticAddr)

	//Generating api object
	profileService, err := profileapi.New(&config)
	if err != nil {
		log.Panic(err)
	}

	profileService.Run()
}
