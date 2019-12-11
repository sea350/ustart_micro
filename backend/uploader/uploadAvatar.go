package uploader

import (
	"context"
	"log"

	"github.com/sea350/ustart_micro/backend/uploader/uploaderpb"
)

//UploadAvatar uploads an avatar picture while returning the image link and error by first replacing the old image and then uploading a new image
func (uploader *Uploader) UploadAvatar(ctx context.Context, req *uploaderpb.UploadAvatarRequest) (*uploaderpb.UploadAvatarResponse, error) {
	if req.OldImageLink != "" {
		err := uploader.Delete(ctx, req.OldImageLink)
		if err != nil {
			return nil, err
		}
	}
	log.Println("Hi")

	newLink, err := uploader.Upload(ctx, req.Base64, req.UUID)

	return &uploaderpb.UploadAvatarResponse{NewLink: newLink}, err
}
