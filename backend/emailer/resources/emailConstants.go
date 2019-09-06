package resources

const (

	//SenderAddress ...USTART NoReply Address
	SenderAddress = "NoReply@ustart.today"

	//VerifSubject ... Subject for verification emails
	VerifSubject = "Your Verification Email"

	//PassResetSubject ... Subject for forgot password emails
	PassResetSubject = "Your Password Reset Link"

	//AWSRegion ... AWS region
	AWSRegion = "us-east-1"

	//EmailConfig ... Email Config Set
	EmailConfig = "ConfigSet"

	//Charset ... UTF-8
	CharSet = "UTF-8"

	//The credentials needed to access AWS Pinpoint

	//CredID ... ID
	CredID = "AKIAR75EUAUYFAON4EH6" //FILL THIS IN
	//CredSecret Secret
	CredSecret = "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" //FILL THIS IN
	//CredToken ... Token
	CredToken = ""

	//AWSEndpoint ... endpoint for pinpoint
	AWSEndpoint = "pinpoint.us-east-1.amazonaws.com/v1"

	//BodyHTML ...constant HTML part of email
	BodyHTML = `<html>
  
		<head></head>
		<body>
			<h1>Amazon Pinpoint Test (SDK for JavaScript in Node.js)</h1>
			<p>This email was sent with
			<a href=''>the Amazon Pinpoint Email API</a> using the
			<a href=''>
				AWS SDK for JavaScript in Node.js</a>.</p>
		</body>
		</html>`

	//BodyText ... text version for email browsers that do not allow html
	BodyText = `We received a request to activate your U·START Account. 
					We would love to assist you! 
					Simply click the button below to verify your account
					VERIFY ACCOUNT 
					a new account`
)
