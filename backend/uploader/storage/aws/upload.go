package awsstore

import(
	"bytes"
	"encoding/base64"
	"strings"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (upload *AWSStore) Upload(ctx context.Context, based64, uploaderID string)(string,error) {
	var arr []string
	i := strings.Index(based64, ",")
	if i < 0 {
		return ``, ErrImproperImport
	}
	arr = strings.Split(based64, `,`)

	dec, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		return ``, err
	}

	r := bytes.NewReader(dec)

	result, err := upload.upl.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(uploaderID + ".png"),
		Body:        r,
		ContentType: aws.String("image/png"),
	})
	
	if err != nil {
		return ``, err
	}

	return result.Location, err
	

}