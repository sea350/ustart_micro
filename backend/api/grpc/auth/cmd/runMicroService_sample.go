package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	authapi "github.com/sea350/ustart_micro/backend/api/grpc/auth"
	"github.com/sea350/ustart_micro/backend/auth"
	"github.com/sea350/ustart_micro/backend/auth/authpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	//Generating api object
	authService, err := auth.New(config.AuthCfg)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	authpb.RegisterAuthServer(srv, authService)
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
