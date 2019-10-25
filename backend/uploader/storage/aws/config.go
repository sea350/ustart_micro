package awsstore

// Config is the configuration for the AWS storage client
type Config struct {
	Region        string
	BucketName = "ustart-bucket"
	S3CredID      string
	S3CredSecret  string
	S3Token       string
}

// NewConfig creates a default config struct
func NewConfig() *Config {
	return &Config{
		Region:         "",
		BucketName:     "ustart-bucket",
		S3CredID:       "",
		S3CredSecret:   "",
		S3Token:        ""
	}
}