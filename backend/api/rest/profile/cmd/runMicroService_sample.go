package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/sea350/ustart_micro/backend/api/rest/profile"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config profile.Config

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
	restAPI, err := profile.New(&config)
	if err != nil {
		panic(err)
	}

	//Assigning the handler functions to a url
	http.HandleFunc("/", nil)
	http.HandleFunc("/lite", restAPI.Lite)
	http.HandleFunc("/lookup", restAPI.Lookup)
	http.HandleFunc("/pull", restAPI.Pull)
	http.HandleFunc("/register", restAPI.Register)

	//Hear and handle
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
