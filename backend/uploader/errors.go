package uploader

import "errors"

var (
	//ErrImproperImport reports when a string imported into a function isn't a parseable base64
	ErrImproperImport = errors.New("The string imported is not in base64 format")
)
