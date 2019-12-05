package uploader

import "context"

//UploadBanner uploads an banner picture while returning the image link and error by first replacing the old image and then uploading a new image
func (uploader *Uploader) UploadBanner(ctx context.Context, based64, uuid, oldimagelink string) (string, error) {

	err := uploader.Delete(ctx, oldimagelink)

	if err != nil {
		return ``, err
	}

	newLink, err := uploader.Upload(ctx, based64, uuid)
	return newLink, err
}
