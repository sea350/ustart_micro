package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	projectapi "github.com/sea350/ustart_micro/backend/api/grpc/project"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config projectapi.Config
	//Importing configuration from json
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	portString := os.Getenv("USTART_PROJ_PORT")
	if portString != "" {
		port, err := strconv.Atoi(portString)
		if err != nil {
			panic(err)
		}
		config.Port = port
	}

	//Generating api object
	projectService, err := projectapi.New(&config)
	if err != nil {
		panic(err)
	}

	projectService.Run()
}
