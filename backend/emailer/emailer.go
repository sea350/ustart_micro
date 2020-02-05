package emailer

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
)

//Emailer is an implementation of the auth service defined in service.proto
type Emailer struct {
	timeFormat    string
	pinpointEmail *pinpointemail.PinpointEmail
}

//New creates a new instance of emailer
func New(cfg *Config) (*Emailer, error) {

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(cfg.AWSRegion),
			Credentials: credentials.NewStaticCredentials(
				cfg.CredentialID,
				cfg.CredentialSecret,
				cfg.CredentialToken)})

	if err != nil {
		return nil, err
	}

	// Create a Pinpoint Email session.
	svc := pinpointemail.New(sess)

	return &Emailer{
		timeFormat:    "RFC3339",
		pinpointEmail: svc,
	}, nil

}
