package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/sea350/ustart_micro/backend/api/rest/auth"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config auth.Config
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
	restAPI, err := auth.New(&config)
	if err != nil {
		panic(err)
	}

	//Assigning the handler functions to a url
	http.HandleFunc("/", nil)
	http.HandleFunc("/authenticate", restAPI.Authenticate)
	http.HandleFunc("/changePassword", restAPI.ChangePassword)
	http.HandleFunc("/register", restAPI.Register)
	http.HandleFunc("/lookup", restAPI.Lookup)
	http.HandleFunc("/verify", restAPI.Verify)
	http.HandleFunc("/reverify", restAPI.Reverify)
	http.HandleFunc("/recoverPassword", restAPI.RecoverPassword)

	//Hear and handle
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
