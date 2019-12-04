package uploader

import (
	"bytes"
	"encoding/base64"
	"strings"
	"github.com/sea350/ustart_micro/backend/uploader/storage"
)

//Delete deletes a profile picture while returning the image link
func (uploader *Uploader) Delete(url string) error {
	splt := strings.Split(url, "/")
	key := splt[len(splt)-1]
​
	key, _ = urlPackage.QueryUnescape(key)

	err = storage.delete(key)

	return err
}
