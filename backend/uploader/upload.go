package uploader

import (
	"bytes"
	"encoding/base64"
	"strings"
	"github.com/sea350/ustart_micro/backend/uploader/storage"

)

//Upload uploads a profile picture while returning the image link
func (uploader *Uploader) Upload(based64 string, uploaderID string) (string, error) {
	
	url, err = storage.upload(based64, uploaderID)

	if err != nil {
		return ``, err
	}

	return url, nil
}
