package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/sea350/ustart_micro/backend/api/rest/project"
	//prof "github.com/sea350/ustart_micro/backend/project"
	//"github.com/sea350/ustart_micro/backend/project/storage"
	//elasticstore "github.com/sea350/ustart_micro/backend/project/storage/elastic"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config project.Config

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

	//Generating api object
	restAPI, err := project.New(&config)
	if err != nil {
		panic(err)
	}

	//Assigning the handler functions to a url
	http.HandleFunc("/block", restAPI.Block)
	http.HandleFunc("/lookup", restAPI.Lookup)
	http.HandleFunc("/pull", restAPI.Pull)
	http.HandleFunc("/follow", restAPI.Follow)

	//Hear and handle
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
