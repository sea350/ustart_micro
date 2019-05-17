package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	profileapi "github.com/sea350/ustart_micro/backend/api/grpc/profile"
	"github.com/sea350/ustart_micro/backend/profile"
	"github.com/sea350/ustart_micro/backend/profile/profilepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	profileService, err := profile.New(config.ProfileCfg)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	profilepb.RegisterProfileServer(srv, profileService)
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
