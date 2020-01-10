package awsstore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//Uploader is a struct that manages uploads for files and images
type Uploader struct {
	//REVISION, inconsistant struct naming
	upl *s3manager.Uploader
	svc *s3.S3
	//REVISION, now stores bucket name
	bucketName string
}

//New creates a new Uploader based on the inserted config
func New(cfg *Config) (*Uploader, error) {
	// if cfg.useDummy
	//REVISION, inconsistent field names and field refrences
	//sesh := session.Must(session.NewSession(&aws.Config{Region: aws.String(cfg.Region), Credentials: credentials.NewStaticCredentials(cfg.S3CredentialID, cfg.S3CredentialSecret, cfg.S3CredentialToken)}))
	sesh := session.Must(session.NewSession(&aws.Config{Region: aws.String(cfg.Region), Credentials: credentials.NewStaticCredentials(cfg.S3CredID, cfg.S3CredSecret, cfg.S3Token)}))
	uploader := s3manager.NewUploader(sesh)
	svc := s3.New(sesh)

	upload := &Uploader{
		upl: uploader,
		svc: svc,
		//REVISION: passed bucket name from cfg to struct
		bucketName: cfg.BucketName,
	}

	return upload, nil

}
