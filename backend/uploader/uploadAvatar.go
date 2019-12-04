package uploader

func (uploader *Uploader) UploadAvatar(based64, uuid, oldimagelink string) (string, error) {
	err = Delete(oldimagelink)

	if err != nil {
		return ``, err
	}

	newLink, err = Upload(based64, uuid)
	return newLink, err
}
