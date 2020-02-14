package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/sea350/ustart_micro/backend/api/rest/profile"
	//prof "github.com/sea350/ustart_micro/backend/profile"
	//"github.com/sea350/ustart_micro/backend/profile/storage"
	//elasticstore "github.com/sea350/ustart_micro/backend/profile/storage/elastic"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config profile.Config

	// config = profile.Config{
	// 	Port: 5002,
	// 	ProfCfg: &prof.Config{
	// 		StorageConfig: &storage.Config{
	// 			ElasticConfig: &elasticstore.Config{
	// 				ElasticAddr: "localhost:9200",
	// 				EType:       "test-profile_data",
	// 				EIndex:      "PROFILE",
	// 			},
	// 		},
	// 		DefaultAvatar: "INSERT URL HERE",
	// 		DefaultBanner: "INSERT URL HERE",
	// 	},
	// }
	// bite, _ := json.Marshal(config)
	// fmt.Println(string(bite))

	//Importing configuration from json
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	portString := os.Getenv("USTART_PROF_PORT")
	if portString != "" {
		port, err := strconv.Atoi(portString)
		if err != nil {
			panic(err)
		}
		config.Port = port
	}

	//Generating api object
	restAPI, err := profile.New(&config)
	if err != nil {
		panic(err)
	}

	//Assigning the handler functions to a url
	http.HandleFunc("/", restAPI.Ping)
	http.HandleFunc("/lite", restAPI.Lite)
	http.HandleFunc("/lookup", restAPI.Lookup)
	http.HandleFunc("/pull", restAPI.Pull)
	http.HandleFunc("/register", restAPI.Register)

	//Hear and handle
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
