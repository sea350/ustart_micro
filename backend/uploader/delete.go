package uploader

import (
	"context"
	urlPackage "net/url"
	"strings"
)

//Delete deletes a profile picture while returning the image link
func (uploader *Uploader) Delete(ctx context.Context, url string) error {
	splt := strings.Split(url, "/")
	key := splt[len(splt)-1]

	key, _ = urlPackage.QueryUnescape(key)

	err := uploader.stor.Delete(ctx, key)

	return err
}
