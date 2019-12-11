package awsstore

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//REVISION exported functions MUST be properly commented, imports corrected

//Upload uploads a file into an s3 bucket
func (upload *Uploader) Upload(ctx context.Context, file *bytes.Reader, filename string) (string, error) {
	log.Println(filename)
	result, err := upload.upl.Upload(&s3manager.UploadInput{
		//REVISION, proper bucket name import
		//Bucket:      aws.String(bucketName),
		Bucket:      aws.String(upload.bucketName),
		Key:         aws.String(filename),
		Body:        file,
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		return ``, err
	}

	return result.Location, err

}
