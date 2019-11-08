package awsstore

import (
	"bytes"
	"context"
	"log"
	urlPackage "net/url"
	"strings"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (upload *Uploader) Delete(ctx context.Context, key string) error {
	input := upload.upl.Delete(&s3.DeleteObjectInput{
		Bucket: aws.String(upload.bucketName),
		Key:    aws.String(key)
	})
​
	_, err := upload.upl.DeleteObject(input)
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
​
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.Println("Debug text: " + result.Location)
​
	return nil

}
