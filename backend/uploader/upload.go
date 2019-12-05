package uploader

import (
	"bytes"
	"context"
	"encoding/base64"
	"strings"
)

//Upload uploads a profile picture while returning the image link
func (uploader *Uploader) Upload(ctx context.Context, based64 string, uploaderID string) (string, error) {
	var arr []string
	i := strings.Index(based64, ",")
	if i < 0 {
		return ``, nil
	}
	arr = strings.Split(based64, `,`)

	dec, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		return ``, err
	}

	r := bytes.NewReader(dec)

	url, err := uploader.stor.Upload(ctx, r, uploaderID)

	if err != nil {
		return ``, err
	}

	return url, nil
}
