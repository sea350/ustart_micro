package awsstore

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

//Delete deletes the files using the key retrieved from the file's url
func (upload *Uploader) Delete(ctx context.Context, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(upload.bucketName),
		Key:    aws.String(key),
	}
	_, err := upload.svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err.Error())
		}
		return err
	}
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.Println("Debug text: " + result.Location)
	return nil

}
