package uploader

//Config contains all information needed to run uploader
type Config struct {
	Region             string
	S3CredentialID     string
	S3CredentialSecret string
	S3CredentialToken  string
}
