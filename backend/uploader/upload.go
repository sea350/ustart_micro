package uploader

import (
	"bytes"
	"encoding/base64"
	"strings"
)

//Upload uploads a profile picture while returning the image link
func (uploader *Uploader) Upload(based64 string, uploaderID string) (string, error) {
	
	url, err = uploader.storage.upload(based64, uploaderID)

	if err != nil {
		return ``, err
	}

	return url, nil
}
