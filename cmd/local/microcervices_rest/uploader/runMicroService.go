package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	uploader "github.com/sea350/ustart_micro/backend/api/rest/uploader"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config uploader.Config

	// config := uploader.Config{
	// 	UploaderCfg: &logicuploader.Config{
	// 		StorageConfig: &storageuploader.Config{
	// 			AWSConfig: &awsuploader.Config{
	// 				Region:       "us-east-2",
	// 				BucketName:   "uploader-testbucket",
	// 				S3CredID:     "AKIAJKRW32O276JJJAFQ",
	// 				S3CredSecret: "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",
	// 			},
	// 		},
	// 	},
	// 	Port: 5001,
	// }
	// bite, _ := json.Marshal(config)
	// fmt.Println(string(bite))

	// //Importing configuration from json
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
	restAPI, err := uploader.New(&config)
	if err != nil {
		panic(err)
	}

	//Assigning the handler functions to a url
	// http.HandleFunc("/", nil)
	http.HandleFunc("/uploadAvatar", restAPI.UploadAvatar)
	http.HandleFunc("/uploadBanner", restAPI.UploadBanner)

	//Hear and handle
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
