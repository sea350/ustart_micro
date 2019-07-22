package storage

//Storage holds the necessary functions
type Storage interface {
	UploadBanner()
	DeleteBanner()
	UploadProfile()
	DeleteProfile()
}
