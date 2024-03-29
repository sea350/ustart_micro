package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

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
		log.Panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Displaying config data")
	log.Println("DB Host: ", config.AuthCfg.StorageConfig.SQLConfig.Host)

	portString := os.Getenv("USTART_AUTH_PORT")
	if portString != "" {
		port, err := strconv.Atoi(portString)
		if err != nil {
			log.Panic(err)
		}
		config.Port = port
	}

	//Generating api object
	authService, err := authapi.New(&config)
	if err != nil {
		log.Panic(err)
	}

	authService.Run()

}
