package emailer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
)

func (emailer *Emailer) SendMail(ctx context.Context, req *emailerpb.EmailerRequest) (*emailerpb.SendMailResponse, error) {

	sendTo := Req.To[0]

	input := &pinpointemail.SendEmailInput{
		Destination: &pinpointemail.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(sendTo),
			},
		},
		Message: &pinpointemail.Message{
			Body: &pinpointemail.Body{
				Html: &pinpointemail.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(req.Body),
				},
				Text: &pinpointemail.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(),
				},
			},
			Subject: &pinpointemail.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(req.Subject),
			},
		},
		Source: aws.String(SenderAddress),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := emailer.pinpointEmail.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ErrCodeMessageRejected:
				fmt.Println(ErrCodeMessageRejected, aerr.Error())
			case ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return nil, false

		return &emailerpb.EmailerResponse{Success: true}, nil

	}

}

/*
	svc := pinpointemailiface.New(sess)
	pinpointEmail := new aws.PinpointEmail()


	params := {
		FromEmailAddress: resources.SenderAddress,
		Destination:{
			ToAddresses: []string{emailReq.To},
		},
		Content:{
			Simple:{
				Body: {
					Html: {
						Data: resources.BodyHTML,
						Charset: resources.CharSet,

					},
					Subject:{
						Data: resources.VerifSubject,
						Charset: resources.CharSet
					}
				}
			},
			ConfigurationSetName: resources.EmailConfig,
			EmailTags:[]
		}
	}


	ret := make(map[string]interface{})

	resp, err := rapi.auth.Authenticate(regCtx, emailReq)
	if resp != nil {
		ret["response"] = resp
	} else {
		ret["response"] = ""
	}
	if err != nil {
		ret["error"] = err.Error()
	} else {
		ret["error"] = ""
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Println("Problem marshaling return data", err)
	}*/
