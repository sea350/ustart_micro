package emailer

type Config struct {
	AWSRegion        string
	EmailConfig      string
	CharSet          string
	CredentialID     string
	CredentialSecret string
	CredentialToken  string
	AWSEndpoint      string
	SenderAddress    string
}
