package uploader

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//Uploader is a struct that manages uploads for files and images
type Uploader struct {
	upl *s3manager.Uploader
}

//New creates a new Uploader based on the inserted config
func New(cfg *Config) (*Uploader, error) {
	// if cfg.useDummy
	sesh := session.Must(session.NewSession(&aws.Config{Region: aws.String(cfg.Region), Credentials: credentials.NewStaticCredentials(cfg.S3CredentialID, cfg.S3CredentialSecret, cfg.S3CredentialToken)}))

	uploader := s3manager.NewUploader(sesh)

	upl := &Uploader{
		upl: uploader,
	}
	return upl, nil
}
