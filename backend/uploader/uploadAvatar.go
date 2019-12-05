package uploader

import "context"

//UploadAvatar uploads an avatar picture while returning the image link and error by first replacing the old image and then uploading a new image
func (uploader *Uploader) UploadAvatar(ctx context.Context, based64, uuid, oldimagelink string) (string, error) {
	err := uploader.Delete(ctx, oldimagelink)

	if err != nil {
		return ``, err
	}

	newLink, err := uploader.Upload(ctx, based64, uuid)
	return newLink, err
}
