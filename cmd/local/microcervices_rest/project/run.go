package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/sea350/ustart_micro/backend/api/rest/project"
	// proj "github.com/sea350/ustart_micro/backend/project"
	// auth "github.com/sea350/ustart_micro/backend/project/auth"
	// authstrg "github.com/sea350/ustart_micro/backend/project/auth/storage"
	// authsqlstrg "github.com/sea350/ustart_micro/backend/project/auth/storage/sql"
	// "github.com/sea350/ustart_micro/backend/project/storage"
	// elasticstore "github.com/sea350/ustart_micro/backend/project/storage/elastic"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config project.Config

	// config = project.Config{
	// 	Port: 5002,
	// 	ProjCfg: &proj.Config{
	// 		StorageConfig: &storage.Config{
	// 			ElasticConfig: &elasticstore.Config{
	// 				ElasticAddr: "http://localhost:9200/",
	// 				EType:       "test-project_data",
	// 				EIndex:      "project",
	// 			},
	// 		},
	// 		AuthConfig: &auth.Config{
	// 			StorageConfig: &authstrg.Config{
	// 				SQLConfig: &authsqlstrg.Config{
	// 					DriverName:       "postgres",
	// 					Host:             "localhost",
	// 					Port:             "5432",
	// 					DBName:           "test",
	// 					Username:         "postgres",
	// 					Password:         "password",
	// 					RoleTableName:    "project_roles",
	// 					RequestTableName: "project_requests",
	// 					MemberTableName:  "project_members",
	// 				},
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

	portString := os.Getenv("USTART_PROJ_PORT")
	if portString != "" {
		port, err := strconv.Atoi(portString)
		if err != nil {
			panic(err)
		}
		config.Port = port
	}

	//Generating api object
	restAPI, err := project.New(&config)
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
