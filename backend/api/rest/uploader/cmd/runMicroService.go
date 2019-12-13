package main

import (
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	uploader "github.com/sea350/ustart_micro/backend/api/rest/uploader"
	logicuploader "github.com/sea350/ustart_micro/backend/uploader"
	storageuploader "github.com/sea350/ustart_micro/backend/uploader/storage"
	awsuploader "github.com/sea350/ustart_micro/backend/uploader/storage/aws"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	config := uploader.Config{
		UploaderCfg: &logicuploader.Config{
			StorageConfig: &storageuploader.Config{
				AWSConfig: &awsuploader.Config{
					Region:       "us-east-2",
					BucketName:   "uploader-testbucket",
					S3CredID:     "AKIAJKRW32O276JJJAFQ",
					S3CredSecret: "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",
				},
			},
		},
		Port: 5001,
	}
	// //Importing configuration from json
	// file, err := os.Open("config.json")
	// if err != nil {
	// 	panic(err)
	// }
	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&config)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", config.UploaderCfg.StorageConfig.AWSConfig.BucketName)
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
