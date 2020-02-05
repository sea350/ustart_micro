package emailer

import "github.com/aws/aws-sdk-go/service/pinpointemail"

var (
	//ErrCodeMessageRejected Imma put min on blast
	ErrCodeMessageRejected = pinpointemail.ErrCodeMessageRejected

	//ErrCodeMailFromDomainNotVerifiedException cuz he's supposed to have done this shit
	ErrCodeMailFromDomainNotVerifiedException = pinpointemail.ErrCodeMailFromDomainNotVerifiedException

	//ErrCodeConfigurationSetDoesNotExistException now i gotta clean up after his rank, curry lookin and smelling, aint knowin how to code lookin ass
	//ErrCodeConfigurationSetDoesNotExistException = pinpointemail.ErrCodeConfigurationSetDoesNotExistException
	//^ that error doesnt even exist
)
